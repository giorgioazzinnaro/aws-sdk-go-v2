// Code generated by smithy-go-codegen DO NOT EDIT.

package appstream

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Starts the specified image builder.
func (c *Client) StartImageBuilder(ctx context.Context, params *StartImageBuilderInput, optFns ...func(*Options)) (*StartImageBuilderOutput, error) {
	if params == nil {
		params = &StartImageBuilderInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartImageBuilder", params, optFns, addOperationStartImageBuilderMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartImageBuilderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartImageBuilderInput struct {

	// The name of the image builder.
	//
	// This member is required.
	Name *string

	// The version of the AppStream 2.0 agent to use for this image builder. To use the
	// latest version of the AppStream 2.0 agent, specify [LATEST].
	AppstreamAgentVersion *string
}

type StartImageBuilderOutput struct {

	// Information about the image builder.
	ImageBuilder *types.ImageBuilder

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStartImageBuilderMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartImageBuilder{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartImageBuilder{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpStartImageBuilderValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartImageBuilder(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opStartImageBuilder(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appstream",
		OperationName: "StartImageBuilder",
	}
}
