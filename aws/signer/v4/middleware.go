package v4

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"

	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	v4Internal "github.com/aws/aws-sdk-go-v2/aws/signer/internal/v4"
	"github.com/aws/aws-sdk-go-v2/internal/sdk"
	"github.com/awslabs/smithy-go/middleware"
	smithyHTTP "github.com/awslabs/smithy-go/transport/http"
)

// HashComputationError indicates an error occurred while computing the signing hash
type HashComputationError struct {
	Err error
}

// Error is the error message
func (e *HashComputationError) Error() string {
	return fmt.Sprintf("failed to compute payload hash: %v", e.Err)
}

// Unwrap returns the underlying error if one is set
func (e *HashComputationError) Unwrap() error {
	return e.Err
}

// SigningError indicates an error condition occurred while performing SigV4 signing
type SigningError struct {
	Err error
}

func (e *SigningError) Error() string {
	return fmt.Sprintf("failed to sign request: %v", e.Err)
}

// Unwrap returns the underlying error cause
func (e *SigningError) Unwrap() error {
	return e.Err
}

// UnsignedPayloadMiddleware sets the SigV4 request payload hash to unsigned
type UnsignedPayloadMiddleware struct{}

// ID returns the UnsignedPayloadMiddleware identifier
func (m *UnsignedPayloadMiddleware) ID() string {
	return "SigV4UnsignedPayloadMiddleware"
}

// HandleFinalize sets the payload hash to be an unsigned payload
func (m *UnsignedPayloadMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	ctx = SetPayloadHash(ctx, v4Internal.UnsignedPayload)
	return next.HandleFinalize(ctx, in)
}

// ComputePayloadSHA256Middleware computes sha256 payload hash to sign
type ComputePayloadSHA256Middleware struct{}

// ID is the middleware name
func (m *ComputePayloadSHA256Middleware) ID() string {
	return "ComputePayloadSHA256Middleware"
}

// HandleFinalize compute the payload hash for the request payload
func (m *ComputePayloadSHA256Middleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	req, ok := in.Request.(*smithyHTTP.Request)
	if !ok {
		return out, metadata, &HashComputationError{Err: fmt.Errorf("unexpected request middleware type %T", in.Request)}
	}

	hash := sha256.New()
	if stream := req.GetStream(); stream != nil {
		_, err = io.Copy(hash, stream)
		if err != nil {
			return out, metadata, &HashComputationError{Err: fmt.Errorf("failed to compute payload hash, %w", err)}
		}

		if err := req.RewindStream(); err != nil {
			return out, metadata, &HashComputationError{Err: fmt.Errorf("failed to seek body to start, %w", err)}
		}
	}

	ctx = SetPayloadHash(ctx, hex.EncodeToString(hash.Sum(nil)))

	return next.HandleFinalize(ctx, in)
}

// SignHTTPRequestMiddleware is a `FinalizeMiddleware` implementation for SigV4 HTTP Signing
type SignHTTPRequestMiddleware struct {
	signer HTTPSigner
}

// NewSignHTTPRequestMiddleware constructs a SignHTTPRequestMiddleware using the given Signer for signing requests
func NewSignHTTPRequestMiddleware(signer HTTPSigner) *SignHTTPRequestMiddleware {
	return &SignHTTPRequestMiddleware{signer: signer}
}

// ID is the SignHTTPRequestMiddleware identifier
func (s *SignHTTPRequestMiddleware) ID() string {
	return "SigV4SignHTTPRequestMiddleware"
}

// HandleFinalize will take the provided input and sign the request using the SigV4 authentication scheme
func (s *SignHTTPRequestMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	req, ok := in.Request.(*smithyHTTP.Request)
	if !ok {
		return out, metadata, &SigningError{Err: fmt.Errorf("unexpected request middleware type %T", in.Request)}
	}

	signingName, signingRegion := awsmiddleware.GetSigningName(ctx), awsmiddleware.GetSigningRegion(ctx)
	payloadHash := GetPayloadHash(ctx)
	if len(payloadHash) == 0 {
		return out, metadata, &SigningError{Err: fmt.Errorf("computed payload hash missing from context")}
	}

	err = s.signer.SignHTTP(ctx, req.Request, payloadHash, signingName, signingRegion, sdk.NowTime())
	if err != nil {
		return out, metadata, &SigningError{Err: fmt.Errorf("failed to sign http request, %w", err)}
	}

	return next.HandleFinalize(ctx, in)
}

type payloadHashKey struct{}

// GetPayloadHash retrieves the payload hash to use for signing
func GetPayloadHash(ctx context.Context) (v string) {
	v, _ = ctx.Value(payloadHashKey{}).(string)
	return v
}

// SetPayloadHash sets the payload hash to be used for signing the request
func SetPayloadHash(ctx context.Context, hash string) context.Context {
	ctx = context.WithValue(ctx, payloadHashKey{}, hash)
	return ctx
}

// AddUnsignedPayloadMiddleware adds UnsignedPayloadMiddleware to the operation middleware stack
func AddUnsignedPayloadMiddleware(stack *middleware.Stack) {
	unsignedPayloadMiddleware := UnsignedPayloadMiddleware{}
	stack.Finalize.Add(&unsignedPayloadMiddleware, middleware.After)
}

// HTTPSignerMiddlewareConfig interface for HTTP signer middleware config
type HTTPSignerMiddlewareConfig interface {
	GetSigner() HTTPSigner
}

// HTTPSignerMiddlewares represent the middleware's for HTTPSigner
type HTTPSignerMiddlewares struct {
	ComputePayloadSHA256Middleware *ComputePayloadSHA256Middleware
	SignHTTPRequestMiddleware      *SignHTTPRequestMiddleware
}

// AddHTTPSignerMiddlewares adds HTTP signer middleware's to operation stack
func AddHTTPSignerMiddlewares(stack *middleware.Stack, cfg HTTPSignerMiddlewareConfig, optFns ...func(*HTTPSignerMiddlewares)) {
	m := HTTPSignerMiddlewares {
		ComputePayloadSHA256Middleware: &ComputePayloadSHA256Middleware{},
		SignHTTPRequestMiddleware:      NewSignHTTPRequestMiddleware(cfg.GetSigner()),
	}
	for _, fn := range optFns {
		fn(&m)
	}

	stack.Finalize.Add(m.ComputePayloadSHA256Middleware, middleware.Before)
	stack.Finalize.Add(m.SignHTTPRequestMiddleware, middleware.After)
}
