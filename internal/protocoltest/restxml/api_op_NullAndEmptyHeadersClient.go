// Code generated by smithy-go-codegen DO NOT EDIT.
package restxml

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Null and empty headers are not sent over the wire.
func (c *Client) NullAndEmptyHeadersClient(ctx context.Context, params *NullAndEmptyHeadersClientInput, optFns ...func(*Options)) (*NullAndEmptyHeadersClientOutput, error) {
	stack := middleware.NewStack("NullAndEmptyHeadersClient", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	awsmiddleware.AddResolveServiceEndpointMiddleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	stack.Initialize.Add(newServiceMetadataMiddleware_opNullAndEmptyHeadersClient(options.Region), middleware.Before)

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
			OperationName: "NullAndEmptyHeadersClient",
			Err:           err,
		}
	}
	out := result.(*NullAndEmptyHeadersClientOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type NullAndEmptyHeadersClientInput struct {
	A *string
	B *string
	C []*string
}

type NullAndEmptyHeadersClientOutput struct {
	A *string
	B *string
	C []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func newServiceMetadataMiddleware_opNullAndEmptyHeadersClient(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:         region,
		ServiceName:    "Rest Xml Protocol",
		ServiceID:      "restxmlprotocol",
		EndpointPrefix: "restxmlprotocol",
		OperationName:  "NullAndEmptyHeadersClient",
	}
}
