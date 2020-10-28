// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworks

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/opsworks/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes Amazon RDS instances. Required Permissions: To use this action, an IAM
// user must have a Show, Deploy, or Manage permissions level for the stack, or an
// attached policy that explicitly grants permissions. For more information about
// user permissions, see Managing User Permissions
// (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
// This call accepts only one resource-identifying parameter.
func (c *Client) DescribeRdsDbInstances(ctx context.Context, params *DescribeRdsDbInstancesInput, optFns ...func(*Options)) (*DescribeRdsDbInstancesOutput, error) {
	if params == nil {
		params = &DescribeRdsDbInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeRdsDbInstances", params, optFns, addOperationDescribeRdsDbInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeRdsDbInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeRdsDbInstancesInput struct {

	// The ID of the stack with which the instances are registered. The operation
	// returns descriptions of all registered Amazon RDS instances.
	//
	// This member is required.
	StackId *string

	// An array containing the ARNs of the instances to be described.
	RdsDbInstanceArns []*string
}

// Contains the response to a DescribeRdsDbInstances request.
type DescribeRdsDbInstancesOutput struct {

	// An a array of RdsDbInstance objects that describe the instances.
	RdsDbInstances []*types.RdsDbInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeRdsDbInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpDescribeRdsDbInstances{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpDescribeRdsDbInstances{})
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
	if err := addOpDescribeRdsDbInstancesValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDescribeRdsDbInstances(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDescribeRdsDbInstances(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks",
		OperationName: "DescribeRdsDbInstances",
	}
}
