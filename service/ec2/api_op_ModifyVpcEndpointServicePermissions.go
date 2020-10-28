// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modifies the permissions for your VPC endpoint service
// (https://docs.aws.amazon.com/vpc/latest/userguide/endpoint-service.html). You
// can add or remove permissions for service consumers (IAM users, IAM roles, and
// AWS accounts) to connect to your endpoint service. If you grant permissions to
// all principals, the service is public. Any users who know the name of a public
// service can send a request to attach an endpoint. If the service does not
// require manual approval, attachments are automatically approved.
func (c *Client) ModifyVpcEndpointServicePermissions(ctx context.Context, params *ModifyVpcEndpointServicePermissionsInput, optFns ...func(*Options)) (*ModifyVpcEndpointServicePermissionsOutput, error) {
	if params == nil {
		params = &ModifyVpcEndpointServicePermissionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyVpcEndpointServicePermissions", params, optFns, addOperationModifyVpcEndpointServicePermissionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyVpcEndpointServicePermissionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyVpcEndpointServicePermissionsInput struct {

	// The ID of the service.
	//
	// This member is required.
	ServiceId *string

	// The Amazon Resource Names (ARN) of one or more principals. Permissions are
	// granted to the principals in this list. To grant permissions to all principals,
	// specify an asterisk (*).
	AddAllowedPrincipals []*string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The Amazon Resource Names (ARN) of one or more principals. Permissions are
	// revoked for principals in this list.
	RemoveAllowedPrincipals []*string
}

type ModifyVpcEndpointServicePermissionsOutput struct {

	// Returns true if the request succeeds; otherwise, it returns an error.
	ReturnValue *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationModifyVpcEndpointServicePermissionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsEc2query_serializeOpModifyVpcEndpointServicePermissions{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsEc2query_deserializeOpModifyVpcEndpointServicePermissions{})
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
	if err := addOpModifyVpcEndpointServicePermissionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opModifyVpcEndpointServicePermissions(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opModifyVpcEndpointServicePermissions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ModifyVpcEndpointServicePermissions",
	}
}
