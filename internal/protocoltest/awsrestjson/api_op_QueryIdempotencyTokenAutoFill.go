// Code generated by smithy-go-codegen DO NOT EDIT.
package awsrestjson

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Automatically adds idempotency tokens.
func (c *Client) QueryIdempotencyTokenAutoFill(ctx context.Context, params *QueryIdempotencyTokenAutoFillInput, optFns ...func(*Options)) (*QueryIdempotencyTokenAutoFillOutput, error) {
	stack := middleware.NewStack("QueryIdempotencyTokenAutoFill", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	awsmiddleware.AddResolveServiceEndpointMiddleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addIdempotencyToken_opQueryIdempotencyTokenAutoFillMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opQueryIdempotencyTokenAutoFill(options.Region), middleware.Before)
	addawsRestjson1_serdeOpQueryIdempotencyTokenAutoFillMiddlewares(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     c.ServiceID(),
			OperationName: "QueryIdempotencyTokenAutoFill",
			Err:           err,
		}
	}
	out := result.(*QueryIdempotencyTokenAutoFillOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type QueryIdempotencyTokenAutoFillInput struct {
	Token *string
}

type QueryIdempotencyTokenAutoFillOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpQueryIdempotencyTokenAutoFillMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add("&awsRestjson1_serializeOpQueryIdempotencyTokenAutoFill{}", middleware.After)
	stack.Deserialize.Add("&awsRestjson1_deserializeOpQueryIdempotencyTokenAutoFill{}", middleware.After)
}

type idempotencyToken_initializeOpQueryIdempotencyTokenAutoFill struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpQueryIdempotencyTokenAutoFill) ID() string {
	return "idempotencyToken_initializeOpQueryIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpQueryIdempotencyTokenAutoFill) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}
	input, ok := in.Parameters.(*QueryIdempotencyTokenAutoFillInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *QueryIdempotencyTokenAutoFillInput ")
	}
	if input.Token == nil {
		t, err := m.tokenProvider.GetToken()
		if err != nil {
			return out, metadata, err
		}
		input.Token = &t
	}
	return next.HandleInitialize(ctx, in)
}

func addIdempotencyToken_opQueryIdempotencyTokenAutoFillMiddleware(stack *middleware.Stack, cfg IdempotencyTokenProvider) {
	stack.Initialize.Add(&idempotencyToken_initializeOpQueryIdempotencyTokenAutoFill{cfg}, middleware.After)
}

func newServiceMetadataMiddleware_opQueryIdempotencyTokenAutoFill(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:         region,
		ServiceName:    "Rest Json Protocol",
		ServiceID:      "restjsonprotocol",
		EndpointPrefix: "restjsonprotocol",
		OperationName:  "QueryIdempotencyTokenAutoFill",
	}
}
