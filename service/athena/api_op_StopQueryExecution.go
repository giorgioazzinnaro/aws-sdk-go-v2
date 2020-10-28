// Code generated by smithy-go-codegen DO NOT EDIT.

package athena

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Stops a query execution. Requires you to have access to the workgroup in which
// the query ran. For code samples using the AWS SDK for Java, see Examples and
// Code Samples (http://docs.aws.amazon.com/athena/latest/ug/code-samples.html) in
// the Amazon Athena User Guide.
func (c *Client) StopQueryExecution(ctx context.Context, params *StopQueryExecutionInput, optFns ...func(*Options)) (*StopQueryExecutionOutput, error) {
	if params == nil {
		params = &StopQueryExecutionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StopQueryExecution", params, optFns, addOperationStopQueryExecutionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StopQueryExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopQueryExecutionInput struct {

	// The unique ID of the query execution to stop.
	//
	// This member is required.
	QueryExecutionId *string
}

type StopQueryExecutionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStopQueryExecutionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpStopQueryExecution{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpStopQueryExecution{})
	if err != nil {
		return err
	}
	if err := awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err := smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err := addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err := v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err := addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err := addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err := awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
		return err
	}
	if err := addClientUserAgent(stack); err != nil {
		return err
	}
	if err := smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err := smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err := addIdempotencyToken_opStopQueryExecutionMiddleware(stack, options); err != nil {
		return err
	}
	if err := addOpStopQueryExecutionValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opStopQueryExecution(options.Region)); err != nil {
		return err
	}
	if err := addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err := addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	return nil
}

type idempotencyToken_initializeOpStopQueryExecution struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStopQueryExecution) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStopQueryExecution) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StopQueryExecutionInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StopQueryExecutionInput ")
	}

	if input.QueryExecutionId == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.QueryExecutionId = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opStopQueryExecutionMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(middleware.Before, &idempotencyToken_initializeOpStopQueryExecution{tokenProvider: cfg.IdempotencyTokenProvider})
}

func newServiceMetadataMiddleware_opStopQueryExecution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "athena",
		OperationName: "StopQueryExecution",
	}
}
