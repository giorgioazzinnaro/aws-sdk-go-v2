// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Registers a target with a maintenance window.
func (c *Client) RegisterTargetWithMaintenanceWindow(ctx context.Context, params *RegisterTargetWithMaintenanceWindowInput, optFns ...func(*Options)) (*RegisterTargetWithMaintenanceWindowOutput, error) {
	if params == nil {
		params = &RegisterTargetWithMaintenanceWindowInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterTargetWithMaintenanceWindow", params, optFns, addOperationRegisterTargetWithMaintenanceWindowMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterTargetWithMaintenanceWindowOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterTargetWithMaintenanceWindowInput struct {

	// The type of target being registered with the maintenance window.
	//
	// This member is required.
	ResourceType types.MaintenanceWindowResourceType

	// The targets to register with the maintenance window. In other words, the
	// instances to run commands on when the maintenance window runs. You can specify
	// targets using instance IDs, resource group names, or tags that have been applied
	// to instances. Example 1: Specify instance IDs
	// Key=InstanceIds,Values=instance-id-1,instance-id-2,instance-id-3  Example 2: Use
	// tag key-pairs applied to instances
	// Key=tag:my-tag-key,Values=my-tag-value-1,my-tag-value-2  Example 3: Use tag-keys
	// applied to instances Key=tag-key,Values=my-tag-key-1,my-tag-key-2  Example 4:
	// Use resource group names Key=resource-groups:Name,Values=resource-group-name
	// Example 5: Use filters for resource group types
	// Key=resource-groups:ResourceTypeFilters,Values=resource-type-1,resource-type-2
	// For Key=resource-groups:ResourceTypeFilters, specify resource types in the
	// following format
	// Key=resource-groups:ResourceTypeFilters,Values=AWS::EC2::INSTANCE,AWS::EC2::VPC
	// For more information about these examples formats, including the best use case
	// for each one, see Examples: Register targets with a maintenance window
	// (https://docs.aws.amazon.com/systems-manager/latest/userguide/mw-cli-tutorial-targets-examples.html)
	// in the AWS Systems Manager User Guide.
	//
	// This member is required.
	Targets []*types.Target

	// The ID of the maintenance window the target should be registered with.
	//
	// This member is required.
	WindowId *string

	// User-provided idempotency token.
	ClientToken *string

	// An optional description for the target.
	Description *string

	// An optional name for the target.
	Name *string

	// User-provided value that will be included in any CloudWatch events raised while
	// running tasks for these targets in this maintenance window.
	OwnerInformation *string
}

type RegisterTargetWithMaintenanceWindowOutput struct {

	// The ID of the target definition in this maintenance window.
	WindowTargetId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRegisterTargetWithMaintenanceWindowMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpRegisterTargetWithMaintenanceWindow{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpRegisterTargetWithMaintenanceWindow{})
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
	if err := addIdempotencyToken_opRegisterTargetWithMaintenanceWindowMiddleware(stack, options); err != nil {
		return err
	}
	if err := addOpRegisterTargetWithMaintenanceWindowValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opRegisterTargetWithMaintenanceWindow(options.Region)); err != nil {
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

type idempotencyToken_initializeOpRegisterTargetWithMaintenanceWindow struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpRegisterTargetWithMaintenanceWindow) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpRegisterTargetWithMaintenanceWindow) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*RegisterTargetWithMaintenanceWindowInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *RegisterTargetWithMaintenanceWindowInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opRegisterTargetWithMaintenanceWindowMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(middleware.Before, &idempotencyToken_initializeOpRegisterTargetWithMaintenanceWindow{tokenProvider: cfg.IdempotencyTokenProvider})
}

func newServiceMetadataMiddleware_opRegisterTargetWithMaintenanceWindow(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "RegisterTargetWithMaintenanceWindow",
	}
}
