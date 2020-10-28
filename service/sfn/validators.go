// Code generated by smithy-go-codegen DO NOT EDIT.

package sfn

import (
	"context"
	"fmt"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
)

type validateOpCreateActivity struct {
}

func (*validateOpCreateActivity) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateActivity) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateActivityInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateActivityInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateStateMachine struct {
}

func (*validateOpCreateStateMachine) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateStateMachine) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateStateMachineInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateStateMachineInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteActivity struct {
}

func (*validateOpDeleteActivity) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteActivity) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteActivityInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteActivityInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteStateMachine struct {
}

func (*validateOpDeleteStateMachine) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteStateMachine) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteStateMachineInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteStateMachineInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeActivity struct {
}

func (*validateOpDescribeActivity) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeActivity) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeActivityInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeActivityInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeExecution struct {
}

func (*validateOpDescribeExecution) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeExecution) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeExecutionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeExecutionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeStateMachineForExecution struct {
}

func (*validateOpDescribeStateMachineForExecution) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeStateMachineForExecution) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeStateMachineForExecutionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeStateMachineForExecutionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeStateMachine struct {
}

func (*validateOpDescribeStateMachine) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeStateMachine) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeStateMachineInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeStateMachineInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetActivityTask struct {
}

func (*validateOpGetActivityTask) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetActivityTask) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetActivityTaskInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetActivityTaskInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetExecutionHistory struct {
}

func (*validateOpGetExecutionHistory) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetExecutionHistory) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetExecutionHistoryInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetExecutionHistoryInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListExecutions struct {
}

func (*validateOpListExecutions) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListExecutions) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListExecutionsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListExecutionsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListTagsForResource struct {
}

func (*validateOpListTagsForResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListTagsForResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListTagsForResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListTagsForResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpSendTaskFailure struct {
}

func (*validateOpSendTaskFailure) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpSendTaskFailure) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*SendTaskFailureInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpSendTaskFailureInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpSendTaskHeartbeat struct {
}

func (*validateOpSendTaskHeartbeat) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpSendTaskHeartbeat) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*SendTaskHeartbeatInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpSendTaskHeartbeatInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpSendTaskSuccess struct {
}

func (*validateOpSendTaskSuccess) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpSendTaskSuccess) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*SendTaskSuccessInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpSendTaskSuccessInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartExecution struct {
}

func (*validateOpStartExecution) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartExecution) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartExecutionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartExecutionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStopExecution struct {
}

func (*validateOpStopExecution) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStopExecution) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StopExecutionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStopExecutionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpTagResource struct {
}

func (*validateOpTagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpTagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*TagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpTagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUntagResource struct {
}

func (*validateOpUntagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUntagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UntagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUntagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateStateMachine struct {
}

func (*validateOpUpdateStateMachine) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateStateMachine) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateStateMachineInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateStateMachineInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCreateActivityValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpCreateActivity{})
}

func addOpCreateStateMachineValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpCreateStateMachine{})
}

func addOpDeleteActivityValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpDeleteActivity{})
}

func addOpDeleteStateMachineValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpDeleteStateMachine{})
}

func addOpDescribeActivityValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpDescribeActivity{})
}

func addOpDescribeExecutionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpDescribeExecution{})
}

func addOpDescribeStateMachineForExecutionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpDescribeStateMachineForExecution{})
}

func addOpDescribeStateMachineValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpDescribeStateMachine{})
}

func addOpGetActivityTaskValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpGetActivityTask{})
}

func addOpGetExecutionHistoryValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpGetExecutionHistory{})
}

func addOpListExecutionsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpListExecutions{})
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpListTagsForResource{})
}

func addOpSendTaskFailureValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpSendTaskFailure{})
}

func addOpSendTaskHeartbeatValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpSendTaskHeartbeat{})
}

func addOpSendTaskSuccessValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpSendTaskSuccess{})
}

func addOpStartExecutionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpStartExecution{})
}

func addOpStopExecutionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpStopExecution{})
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpTagResource{})
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpUntagResource{})
}

func addOpUpdateStateMachineValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpUpdateStateMachine{})
}

func validateOpCreateActivityInput(v *CreateActivityInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateActivityInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateStateMachineInput(v *CreateStateMachineInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateStateMachineInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.Definition == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Definition"))
	}
	if v.RoleArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RoleArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteActivityInput(v *DeleteActivityInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteActivityInput"}
	if v.ActivityArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ActivityArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteStateMachineInput(v *DeleteStateMachineInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteStateMachineInput"}
	if v.StateMachineArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StateMachineArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeActivityInput(v *DescribeActivityInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeActivityInput"}
	if v.ActivityArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ActivityArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeExecutionInput(v *DescribeExecutionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeExecutionInput"}
	if v.ExecutionArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ExecutionArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeStateMachineForExecutionInput(v *DescribeStateMachineForExecutionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeStateMachineForExecutionInput"}
	if v.ExecutionArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ExecutionArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeStateMachineInput(v *DescribeStateMachineInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeStateMachineInput"}
	if v.StateMachineArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StateMachineArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetActivityTaskInput(v *GetActivityTaskInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetActivityTaskInput"}
	if v.ActivityArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ActivityArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetExecutionHistoryInput(v *GetExecutionHistoryInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetExecutionHistoryInput"}
	if v.ExecutionArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ExecutionArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListExecutionsInput(v *ListExecutionsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListExecutionsInput"}
	if v.StateMachineArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StateMachineArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListTagsForResourceInput(v *ListTagsForResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListTagsForResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpSendTaskFailureInput(v *SendTaskFailureInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SendTaskFailureInput"}
	if v.TaskToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TaskToken"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpSendTaskHeartbeatInput(v *SendTaskHeartbeatInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SendTaskHeartbeatInput"}
	if v.TaskToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TaskToken"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpSendTaskSuccessInput(v *SendTaskSuccessInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SendTaskSuccessInput"}
	if v.Output == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Output"))
	}
	if v.TaskToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TaskToken"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartExecutionInput(v *StartExecutionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartExecutionInput"}
	if v.StateMachineArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StateMachineArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStopExecutionInput(v *StopExecutionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StopExecutionInput"}
	if v.ExecutionArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ExecutionArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpTagResourceInput(v *TagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagResourceInput"}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUntagResourceInput(v *UntagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UntagResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if v.TagKeys == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TagKeys"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateStateMachineInput(v *UpdateStateMachineInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateStateMachineInput"}
	if v.StateMachineArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StateMachineArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
