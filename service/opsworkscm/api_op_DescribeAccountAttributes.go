// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworkscm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/opsworkscm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes your OpsWorks-CM account attributes. This operation is synchronous.
func (c *Client) DescribeAccountAttributes(ctx context.Context, params *DescribeAccountAttributesInput, optFns ...func(*Options)) (*DescribeAccountAttributesOutput, error) {
	if params == nil {
		params = &DescribeAccountAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAccountAttributes", params, optFns, addOperationDescribeAccountAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAccountAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAccountAttributesInput struct {
}

type DescribeAccountAttributesOutput struct {

	// The attributes that are currently set for the account.
	Attributes []*types.AccountAttribute

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeAccountAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeAccountAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeAccountAttributes{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAccountAttributes(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeAccountAttributes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks-cm",
		OperationName: "DescribeAccountAttributes",
	}
}
