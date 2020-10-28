// Code generated by smithy-go-codegen DO NOT EDIT.

package cloud9

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/cloud9/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
)

type validateOpCreateEnvironmentEC2 struct {
}

func (*validateOpCreateEnvironmentEC2) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateEnvironmentEC2) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateEnvironmentEC2Input)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateEnvironmentEC2Input(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateEnvironmentMembership struct {
}

func (*validateOpCreateEnvironmentMembership) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateEnvironmentMembership) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateEnvironmentMembershipInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateEnvironmentMembershipInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteEnvironment struct {
}

func (*validateOpDeleteEnvironment) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteEnvironment) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteEnvironmentInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteEnvironmentInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteEnvironmentMembership struct {
}

func (*validateOpDeleteEnvironmentMembership) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteEnvironmentMembership) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteEnvironmentMembershipInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteEnvironmentMembershipInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeEnvironments struct {
}

func (*validateOpDescribeEnvironments) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeEnvironments) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeEnvironmentsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeEnvironmentsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeEnvironmentStatus struct {
}

func (*validateOpDescribeEnvironmentStatus) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeEnvironmentStatus) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeEnvironmentStatusInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeEnvironmentStatusInput(input); err != nil {
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

type validateOpUpdateEnvironment struct {
}

func (*validateOpUpdateEnvironment) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateEnvironment) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateEnvironmentInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateEnvironmentInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateEnvironmentMembership struct {
}

func (*validateOpUpdateEnvironmentMembership) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateEnvironmentMembership) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateEnvironmentMembershipInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateEnvironmentMembershipInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCreateEnvironmentEC2ValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpCreateEnvironmentEC2{})
}

func addOpCreateEnvironmentMembershipValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpCreateEnvironmentMembership{})
}

func addOpDeleteEnvironmentValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpDeleteEnvironment{})
}

func addOpDeleteEnvironmentMembershipValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpDeleteEnvironmentMembership{})
}

func addOpDescribeEnvironmentsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpDescribeEnvironments{})
}

func addOpDescribeEnvironmentStatusValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpDescribeEnvironmentStatus{})
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpListTagsForResource{})
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpTagResource{})
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpUntagResource{})
}

func addOpUpdateEnvironmentValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpUpdateEnvironment{})
}

func addOpUpdateEnvironmentMembershipValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpUpdateEnvironmentMembership{})
}

func validateTag(v *types.Tag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Tag"}
	if v.Key == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Key"))
	}
	if v.Value == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Value"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTagList(v []*types.Tag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagList"}
	for i := range v {
		if err := validateTag(v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateEnvironmentEC2Input(v *CreateEnvironmentEC2Input) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateEnvironmentEC2Input"}
	if v.InstanceType == nil {
		invalidParams.Add(smithy.NewErrParamRequired("InstanceType"))
	}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.Tags != nil {
		if err := validateTagList(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateEnvironmentMembershipInput(v *CreateEnvironmentMembershipInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateEnvironmentMembershipInput"}
	if v.UserArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("UserArn"))
	}
	if v.EnvironmentId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EnvironmentId"))
	}
	if len(v.Permissions) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Permissions"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteEnvironmentInput(v *DeleteEnvironmentInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteEnvironmentInput"}
	if v.EnvironmentId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EnvironmentId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteEnvironmentMembershipInput(v *DeleteEnvironmentMembershipInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteEnvironmentMembershipInput"}
	if v.UserArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("UserArn"))
	}
	if v.EnvironmentId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EnvironmentId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeEnvironmentsInput(v *DescribeEnvironmentsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeEnvironmentsInput"}
	if v.EnvironmentIds == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EnvironmentIds"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeEnvironmentStatusInput(v *DescribeEnvironmentStatusInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeEnvironmentStatusInput"}
	if v.EnvironmentId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EnvironmentId"))
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
	if v.ResourceARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceARN"))
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
	} else if v.Tags != nil {
		if err := validateTagList(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
	}
	if v.ResourceARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceARN"))
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
	if v.TagKeys == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TagKeys"))
	}
	if v.ResourceARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceARN"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateEnvironmentInput(v *UpdateEnvironmentInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateEnvironmentInput"}
	if v.EnvironmentId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EnvironmentId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateEnvironmentMembershipInput(v *UpdateEnvironmentMembershipInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateEnvironmentMembershipInput"}
	if len(v.Permissions) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Permissions"))
	}
	if v.EnvironmentId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EnvironmentId"))
	}
	if v.UserArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("UserArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
