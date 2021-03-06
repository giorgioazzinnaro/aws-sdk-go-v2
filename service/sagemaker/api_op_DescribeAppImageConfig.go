// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Describes an AppImageConfig.
func (c *Client) DescribeAppImageConfig(ctx context.Context, params *DescribeAppImageConfigInput, optFns ...func(*Options)) (*DescribeAppImageConfigOutput, error) {
	if params == nil {
		params = &DescribeAppImageConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAppImageConfig", params, optFns, addOperationDescribeAppImageConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAppImageConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAppImageConfigInput struct {

	// The name of the AppImageConfig to describe.
	//
	// This member is required.
	AppImageConfigName *string
}

type DescribeAppImageConfigOutput struct {

	// The Amazon Resource Name (ARN) of the AppImageConfig.
	AppImageConfigArn *string

	// The name of the AppImageConfig.
	AppImageConfigName *string

	// When the AppImageConfig was created.
	CreationTime *time.Time

	// The KernelGateway app.
	KernelGatewayImageConfig *types.KernelGatewayImageConfig

	// When the AppImageConfig was last modified.
	LastModifiedTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeAppImageConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeAppImageConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeAppImageConfig{}, middleware.After)
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
	addOpDescribeAppImageConfigValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAppImageConfig(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeAppImageConfig(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DescribeAppImageConfig",
	}
}
