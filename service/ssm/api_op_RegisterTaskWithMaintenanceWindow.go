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

// Adds a new task to a maintenance window.
func (c *Client) RegisterTaskWithMaintenanceWindow(ctx context.Context, params *RegisterTaskWithMaintenanceWindowInput, optFns ...func(*Options)) (*RegisterTaskWithMaintenanceWindowOutput, error) {
	if params == nil {
		params = &RegisterTaskWithMaintenanceWindowInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterTaskWithMaintenanceWindow", params, optFns, addOperationRegisterTaskWithMaintenanceWindowMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterTaskWithMaintenanceWindowOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterTaskWithMaintenanceWindowInput struct {

	// The maximum number of targets this task can be run for in parallel.
	//
	// This member is required.
	MaxConcurrency *string

	// The maximum number of errors allowed before this task stops being scheduled.
	//
	// This member is required.
	MaxErrors *string

	// The targets (either instances or maintenance window targets). Specify instances
	// using the following format: Key=InstanceIds,Values=, Specify maintenance window
	// targets using the following format: Key=WindowTargetIds;,Values=,
	//
	// This member is required.
	Targets []*types.Target

	// The ARN of the task to run.
	//
	// This member is required.
	TaskArn *string

	// The type of task being registered.
	//
	// This member is required.
	TaskType types.MaintenanceWindowTaskType

	// The ID of the maintenance window the task should be added to.
	//
	// This member is required.
	WindowId *string

	// User-provided idempotency token.
	ClientToken *string

	// An optional description for the task.
	Description *string

	// A structure containing information about an S3 bucket to write instance-level
	// logs to. LoggingInfo has been deprecated. To specify an S3 bucket to contain
	// logs, instead use the OutputS3BucketName and OutputS3KeyPrefix options in the
	// TaskInvocationParameters structure. For information about how Systems Manager
	// handles these options for the supported maintenance window task types, see
	// MaintenanceWindowTaskInvocationParameters.
	LoggingInfo *types.LoggingInfo

	// An optional name for the task.
	Name *string

	// The priority of the task in the maintenance window, the lower the number the
	// higher the priority. Tasks in a maintenance window are scheduled in priority
	// order with tasks that have the same priority scheduled in parallel.
	Priority *int32

	// The ARN of the IAM service role for Systems Manager to assume when running a
	// maintenance window task. If you do not specify a service role ARN, Systems
	// Manager uses your account's service-linked role. If no service-linked role for
	// Systems Manager exists in your account, it is created when you run
	// RegisterTaskWithMaintenanceWindow. For more information, see the following
	// topics in the in the AWS Systems Manager User Guide:
	//
	//     * Using service-linked
	// roles for Systems Manager
	// (https://docs.aws.amazon.com/systems-manager/latest/userguide/using-service-linked-roles.html#slr-permissions)
	//
	//
	// * Should I use a service-linked role or a custom service role to run maintenance
	// window tasks?
	// (https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-maintenance-permissions.html#maintenance-window-tasks-service-role)
	ServiceRoleArn *string

	// The parameters that the task should use during execution. Populate only the
	// fields that match the task type. All other fields should be empty.
	TaskInvocationParameters *types.MaintenanceWindowTaskInvocationParameters

	// The parameters that should be passed to the task when it is run. TaskParameters
	// has been deprecated. To specify parameters to pass to a task when it runs,
	// instead use the Parameters option in the TaskInvocationParameters structure. For
	// information about how Systems Manager handles these options for the supported
	// maintenance window task types, see MaintenanceWindowTaskInvocationParameters.
	TaskParameters map[string]*types.MaintenanceWindowTaskParameterValueExpression
}

type RegisterTaskWithMaintenanceWindowOutput struct {

	// The ID of the task in the maintenance window.
	WindowTaskId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRegisterTaskWithMaintenanceWindowMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpRegisterTaskWithMaintenanceWindow{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpRegisterTaskWithMaintenanceWindow{})
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
	if err := addIdempotencyToken_opRegisterTaskWithMaintenanceWindowMiddleware(stack, options); err != nil {
		return err
	}
	if err := addOpRegisterTaskWithMaintenanceWindowValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opRegisterTaskWithMaintenanceWindow(options.Region)); err != nil {
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

type idempotencyToken_initializeOpRegisterTaskWithMaintenanceWindow struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpRegisterTaskWithMaintenanceWindow) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpRegisterTaskWithMaintenanceWindow) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*RegisterTaskWithMaintenanceWindowInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *RegisterTaskWithMaintenanceWindowInput ")
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
func addIdempotencyToken_opRegisterTaskWithMaintenanceWindowMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(middleware.Before, &idempotencyToken_initializeOpRegisterTaskWithMaintenanceWindow{tokenProvider: cfg.IdempotencyTokenProvider})
}

func newServiceMetadataMiddleware_opRegisterTaskWithMaintenanceWindow(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "RegisterTaskWithMaintenanceWindow",
	}
}
