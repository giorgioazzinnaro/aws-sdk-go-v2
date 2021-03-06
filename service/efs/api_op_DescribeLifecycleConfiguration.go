// Code generated by smithy-go-codegen DO NOT EDIT.

package efs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/efs/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the current LifecycleConfiguration object for the specified Amazon EFS
// file system. EFS lifecycle management uses the LifecycleConfiguration object to
// identify which files to move to the EFS Infrequent Access (IA) storage class.
// For a file system without a LifecycleConfiguration object, the call returns an
// empty array in the response. This operation requires permissions for the
// elasticfilesystem:DescribeLifecycleConfiguration operation.
func (c *Client) DescribeLifecycleConfiguration(ctx context.Context, params *DescribeLifecycleConfigurationInput, optFns ...func(*Options)) (*DescribeLifecycleConfigurationOutput, error) {
	if params == nil {
		params = &DescribeLifecycleConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeLifecycleConfiguration", params, optFns, addOperationDescribeLifecycleConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeLifecycleConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeLifecycleConfigurationInput struct {

	// The ID of the file system whose LifecycleConfiguration object you want to
	// retrieve (String).
	//
	// This member is required.
	FileSystemId *string
}

type DescribeLifecycleConfigurationOutput struct {

	// An array of lifecycle management policies. Currently, EFS supports a maximum of
	// one policy per file system.
	LifecyclePolicies []*types.LifecyclePolicy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeLifecycleConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeLifecycleConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeLifecycleConfiguration{}, middleware.After)
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
	addOpDescribeLifecycleConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeLifecycleConfiguration(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeLifecycleConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticfilesystem",
		OperationName: "DescribeLifecycleConfiguration",
	}
}
