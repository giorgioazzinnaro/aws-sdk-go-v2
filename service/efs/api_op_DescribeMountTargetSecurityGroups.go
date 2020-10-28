// Code generated by smithy-go-codegen DO NOT EDIT.

package efs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the security groups currently in effect for a mount target. This
// operation requires that the network interface of the mount target has been
// created and the lifecycle state of the mount target is not deleted. This
// operation requires permissions for the following actions:
//
//     *
// elasticfilesystem:DescribeMountTargetSecurityGroups action on the mount target's
// file system.
//
//     * ec2:DescribeNetworkInterfaceAttribute action on the mount
// target's network interface.
func (c *Client) DescribeMountTargetSecurityGroups(ctx context.Context, params *DescribeMountTargetSecurityGroupsInput, optFns ...func(*Options)) (*DescribeMountTargetSecurityGroupsOutput, error) {
	if params == nil {
		params = &DescribeMountTargetSecurityGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeMountTargetSecurityGroups", params, optFns, addOperationDescribeMountTargetSecurityGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeMountTargetSecurityGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DescribeMountTargetSecurityGroupsInput struct {

	// The ID of the mount target whose security groups you want to retrieve.
	//
	// This member is required.
	MountTargetId *string
}

type DescribeMountTargetSecurityGroupsOutput struct {

	// An array of security groups.
	//
	// This member is required.
	SecurityGroups []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeMountTargetSecurityGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpDescribeMountTargetSecurityGroups{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpDescribeMountTargetSecurityGroups{})
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
	if err := addOpDescribeMountTargetSecurityGroupsValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDescribeMountTargetSecurityGroups(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDescribeMountTargetSecurityGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticfilesystem",
		OperationName: "DescribeMountTargetSecurityGroups",
	}
}
