// Code generated by smithy-go-codegen DO NOT EDIT.

package efs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modifies the set of security groups in effect for a mount target. When you
// create a mount target, Amazon EFS also creates a new network interface. For more
// information, see CreateMountTarget. This operation replaces the security groups
// in effect for the network interface associated with a mount target, with the
// SecurityGroups provided in the request. This operation requires that the network
// interface of the mount target has been created and the lifecycle state of the
// mount target is not deleted. The operation requires permissions for the
// following actions:
//
//     * elasticfilesystem:ModifyMountTargetSecurityGroups
// action on the mount target's file system.
//
//     *
// ec2:ModifyNetworkInterfaceAttribute action on the mount target's network
// interface.
func (c *Client) ModifyMountTargetSecurityGroups(ctx context.Context, params *ModifyMountTargetSecurityGroupsInput, optFns ...func(*Options)) (*ModifyMountTargetSecurityGroupsOutput, error) {
	if params == nil {
		params = &ModifyMountTargetSecurityGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyMountTargetSecurityGroups", params, optFns, addOperationModifyMountTargetSecurityGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyMountTargetSecurityGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type ModifyMountTargetSecurityGroupsInput struct {

	// The ID of the mount target whose security groups you want to modify.
	//
	// This member is required.
	MountTargetId *string

	// An array of up to five VPC security group IDs.
	SecurityGroups []*string
}

type ModifyMountTargetSecurityGroupsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationModifyMountTargetSecurityGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpModifyMountTargetSecurityGroups{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpModifyMountTargetSecurityGroups{})
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
	if err := addOpModifyMountTargetSecurityGroupsValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opModifyMountTargetSecurityGroups(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opModifyMountTargetSecurityGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticfilesystem",
		OperationName: "ModifyMountTargetSecurityGroups",
	}
}
