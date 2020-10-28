package retry

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddle "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/internal/sdk"
	smithymiddle "github.com/awslabs/smithy-go/middleware"
	"github.com/awslabs/smithy-go/transport/http"
)

// RequestCloner is a function that can take an input request type and clone the request
// for use in a subsequent retry attempt
type RequestCloner func(interface{}) interface{}

type retryMetadata struct {
	AttemptNum       int
	AttemptTime      time.Time
	MaxAttempts      int
	AttemptClockSkew time.Duration
}

type retryMetadataKey struct{}

// Attempt is a Smithy FinalizeMiddleware that handles retry attempts using the provided
// Retryer implementation
type Attempt struct {
	retryer       Retryer
	requestCloner RequestCloner
}

// NewAttemptMiddleware returns a new Attempt
func NewAttemptMiddleware(retryer Retryer, requestCloner RequestCloner) Attempt {
	return Attempt{retryer: retryer, requestCloner: requestCloner}
}

// ID returns the middleware identifier
func (r *Attempt) ID() string {
	return "Retry"
}

// HandleFinalize utilizes the provider Retryer implementation to attempt retries over the next handler
func (r Attempt) HandleFinalize(ctx context.Context, in smithymiddle.FinalizeInput, next smithymiddle.FinalizeHandler) (
	out smithymiddle.FinalizeOutput, metadata smithymiddle.Metadata, err error,
) {
	var attemptNum, retryCount int
	var attemptClockSkew time.Duration

	maxAttempts := r.retryer.MaxAttempts()

	relRetryToken := r.retryer.GetInitialToken()

	for {
		attemptNum++

		attemptInput := in
		attemptInput.Request = r.requestCloner(attemptInput.Request)

		attemptCtx := setRetryMetadata(ctx, retryMetadata{
			AttemptNum:       attemptNum,
			AttemptTime:      sdk.NowTime().UTC(),
			MaxAttempts:      maxAttempts,
			AttemptClockSkew: attemptClockSkew,
		})

		if attemptNum > 1 {
			if rewindable, ok := in.Request.(interface{ RewindStream() error }); ok {
				if err := rewindable.RewindStream(); err != nil {
					return out, metadata, fmt.Errorf("failed to rewind transport stream for retry, %w", err)
				}
			}
		}

		out, metadata, reqErr := next.HandleFinalize(attemptCtx, attemptInput)

		relRetryToken(reqErr)
		if reqErr == nil {
			return out, metadata, nil
		}

		retryable := r.retryer.IsErrorRetryable(reqErr)
		if !retryable {
			return out, metadata, reqErr
		}

		if maxAttempts > 0 && attemptNum >= maxAttempts {
			err = &MaxAttemptsError{
				Attempt: attemptNum,
				Err:     reqErr,
			}
			return out, metadata, err
		}

		relRetryToken, err = r.retryer.GetRetryToken(ctx, reqErr)
		if err != nil {
			return out, metadata, err
		}

		retryDelay, err := r.retryer.RetryDelay(attemptNum, reqErr)
		if err != nil {
			return out, metadata, err
		}

		if err = sdk.SleepWithContext(ctx, retryDelay); err != nil {
			err = &aws.RequestCanceledError{Err: err}
			return out, metadata, err
		}

		responseMetadata := awsmiddle.GetResponseMetadata(metadata)
		attemptClockSkew = responseMetadata.AttemptSkew

		retryCount++
	}
}

// MetricsHeader attaches SDK request metric header for retries to the transport
type MetricsHeader struct{}

// ID returns the middleware identifier
func (r MetricsHeader) ID() string {
	return "RetryMetricsHeader"
}

// HandleFinalize attaches the sdk request metric header to the transport layer
func (r MetricsHeader) HandleFinalize(ctx context.Context, in smithymiddle.FinalizeInput, next smithymiddle.FinalizeHandler) (
	out smithymiddle.FinalizeOutput, metadata smithymiddle.Metadata, err error,
) {
	retryMetadata, ok := getRetryMetadata(ctx)
	if !ok {
		return out, metadata, fmt.Errorf("retry metadata value not found on context")
	}

	const retryMetricHeader = "Amz-Sdk-Request"
	var parts []string

	parts = append(parts, "attempt="+strconv.Itoa(retryMetadata.AttemptNum))
	if retryMetadata.MaxAttempts != 0 {
		parts = append(parts, "max="+strconv.Itoa(retryMetadata.MaxAttempts))
	}

	var ttl time.Time
	if deadline, ok := ctx.Deadline(); ok {
		ttl = deadline
	}

	// Only append the TTL if it can be determined.
	if !ttl.IsZero() && retryMetadata.AttemptClockSkew > 0 {
		const unixTimeFormat = "20060102T150405Z"
		ttl = ttl.Add(retryMetadata.AttemptClockSkew)
		parts = append(parts, "ttl="+ttl.Format(unixTimeFormat))
	}

	switch req := in.Request.(type) {
	case *http.Request:
		req.Header[retryMetricHeader] = append(req.Header[retryMetricHeader][:0], strings.Join(parts, "; "))
	default:
		return out, metadata, fmt.Errorf("unknown transport type %T", req)
	}

	return next.HandleFinalize(ctx, in)
}

// getRetryMetadata retrieves retryMetadata from the context and a bool indicating if it was set
func getRetryMetadata(ctx context.Context) (metadata retryMetadata, ok bool) {
	metadata, ok = ctx.Value(retryMetadataKey{}).(retryMetadata)
	return metadata, ok
}

func setRetryMetadata(ctx context.Context, metadata retryMetadata) context.Context {
	return context.WithValue(ctx, retryMetadataKey{}, metadata)
}

// AddRetryMiddlewaresOptions is the set of options that can be passed to AddRetryMiddlewares for configuring retry
// associated middleware.
type AddRetryMiddlewaresOptions struct {
	Retryer Retryer
}

// AddRetryMiddlewares adds retry middleware to operation middleware stack
func AddRetryMiddlewares(stack *smithymiddle.Stack, options AddRetryMiddlewaresOptions) error {
	attempt := NewAttemptMiddleware(options.Retryer, http.RequestCloner)
	if err := stack.Finalize.Add(smithymiddle.After, &attempt); err != nil {
		return err
	}
	if err := stack.Finalize.Add(smithymiddle.After, &MetricsHeader{}); err != nil {
		return err
	}
	return nil
}
