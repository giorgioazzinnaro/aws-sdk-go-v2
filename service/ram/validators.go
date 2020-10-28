// Code generated by smithy-go-codegen DO NOT EDIT.

package ram

import (
	"context"
	"fmt"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
)

type validateOpAcceptResourceShareInvitation struct {
}

func (*validateOpAcceptResourceShareInvitation) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpAcceptResourceShareInvitation) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*AcceptResourceShareInvitationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpAcceptResourceShareInvitationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpAssociateResourceShare struct {
}

func (*validateOpAssociateResourceShare) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpAssociateResourceShare) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*AssociateResourceShareInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpAssociateResourceShareInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpAssociateResourceSharePermission struct {
}

func (*validateOpAssociateResourceSharePermission) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpAssociateResourceSharePermission) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*AssociateResourceSharePermissionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpAssociateResourceSharePermissionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateResourceShare struct {
}

func (*validateOpCreateResourceShare) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateResourceShare) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateResourceShareInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateResourceShareInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteResourceShare struct {
}

func (*validateOpDeleteResourceShare) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteResourceShare) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteResourceShareInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteResourceShareInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDisassociateResourceShare struct {
}

func (*validateOpDisassociateResourceShare) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDisassociateResourceShare) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DisassociateResourceShareInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDisassociateResourceShareInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDisassociateResourceSharePermission struct {
}

func (*validateOpDisassociateResourceSharePermission) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDisassociateResourceSharePermission) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DisassociateResourceSharePermissionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDisassociateResourceSharePermissionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetPermission struct {
}

func (*validateOpGetPermission) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetPermission) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetPermissionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetPermissionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetResourcePolicies struct {
}

func (*validateOpGetResourcePolicies) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetResourcePolicies) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetResourcePoliciesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetResourcePoliciesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetResourceShareAssociations struct {
}

func (*validateOpGetResourceShareAssociations) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetResourceShareAssociations) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetResourceShareAssociationsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetResourceShareAssociationsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetResourceShares struct {
}

func (*validateOpGetResourceShares) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetResourceShares) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetResourceSharesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetResourceSharesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListPendingInvitationResources struct {
}

func (*validateOpListPendingInvitationResources) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListPendingInvitationResources) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListPendingInvitationResourcesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListPendingInvitationResourcesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListPrincipals struct {
}

func (*validateOpListPrincipals) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListPrincipals) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListPrincipalsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListPrincipalsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListResourceSharePermissions struct {
}

func (*validateOpListResourceSharePermissions) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListResourceSharePermissions) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListResourceSharePermissionsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListResourceSharePermissionsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListResources struct {
}

func (*validateOpListResources) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListResources) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListResourcesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListResourcesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPromoteResourceShareCreatedFromPolicy struct {
}

func (*validateOpPromoteResourceShareCreatedFromPolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPromoteResourceShareCreatedFromPolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PromoteResourceShareCreatedFromPolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPromoteResourceShareCreatedFromPolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpRejectResourceShareInvitation struct {
}

func (*validateOpRejectResourceShareInvitation) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpRejectResourceShareInvitation) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*RejectResourceShareInvitationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpRejectResourceShareInvitationInput(input); err != nil {
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

type validateOpUpdateResourceShare struct {
}

func (*validateOpUpdateResourceShare) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateResourceShare) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateResourceShareInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateResourceShareInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpAcceptResourceShareInvitationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpAcceptResourceShareInvitation{})
}

func addOpAssociateResourceShareValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpAssociateResourceShare{})
}

func addOpAssociateResourceSharePermissionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpAssociateResourceSharePermission{})
}

func addOpCreateResourceShareValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpCreateResourceShare{})
}

func addOpDeleteResourceShareValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpDeleteResourceShare{})
}

func addOpDisassociateResourceShareValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpDisassociateResourceShare{})
}

func addOpDisassociateResourceSharePermissionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpDisassociateResourceSharePermission{})
}

func addOpGetPermissionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpGetPermission{})
}

func addOpGetResourcePoliciesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpGetResourcePolicies{})
}

func addOpGetResourceShareAssociationsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpGetResourceShareAssociations{})
}

func addOpGetResourceSharesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpGetResourceShares{})
}

func addOpListPendingInvitationResourcesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpListPendingInvitationResources{})
}

func addOpListPrincipalsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpListPrincipals{})
}

func addOpListResourceSharePermissionsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpListResourceSharePermissions{})
}

func addOpListResourcesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpListResources{})
}

func addOpPromoteResourceShareCreatedFromPolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpPromoteResourceShareCreatedFromPolicy{})
}

func addOpRejectResourceShareInvitationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpRejectResourceShareInvitation{})
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpTagResource{})
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpUntagResource{})
}

func addOpUpdateResourceShareValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpUpdateResourceShare{})
}

func validateOpAcceptResourceShareInvitationInput(v *AcceptResourceShareInvitationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AcceptResourceShareInvitationInput"}
	if v.ResourceShareInvitationArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceShareInvitationArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpAssociateResourceShareInput(v *AssociateResourceShareInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AssociateResourceShareInput"}
	if v.ResourceShareArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceShareArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpAssociateResourceSharePermissionInput(v *AssociateResourceSharePermissionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AssociateResourceSharePermissionInput"}
	if v.ResourceShareArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceShareArn"))
	}
	if v.PermissionArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PermissionArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateResourceShareInput(v *CreateResourceShareInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateResourceShareInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteResourceShareInput(v *DeleteResourceShareInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteResourceShareInput"}
	if v.ResourceShareArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceShareArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDisassociateResourceShareInput(v *DisassociateResourceShareInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DisassociateResourceShareInput"}
	if v.ResourceShareArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceShareArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDisassociateResourceSharePermissionInput(v *DisassociateResourceSharePermissionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DisassociateResourceSharePermissionInput"}
	if v.ResourceShareArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceShareArn"))
	}
	if v.PermissionArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PermissionArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetPermissionInput(v *GetPermissionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetPermissionInput"}
	if v.PermissionArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PermissionArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetResourcePoliciesInput(v *GetResourcePoliciesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetResourcePoliciesInput"}
	if v.ResourceArns == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArns"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetResourceShareAssociationsInput(v *GetResourceShareAssociationsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetResourceShareAssociationsInput"}
	if len(v.AssociationType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("AssociationType"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetResourceSharesInput(v *GetResourceSharesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetResourceSharesInput"}
	if len(v.ResourceOwner) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceOwner"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListPendingInvitationResourcesInput(v *ListPendingInvitationResourcesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListPendingInvitationResourcesInput"}
	if v.ResourceShareInvitationArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceShareInvitationArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListPrincipalsInput(v *ListPrincipalsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListPrincipalsInput"}
	if len(v.ResourceOwner) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceOwner"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListResourceSharePermissionsInput(v *ListResourceSharePermissionsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListResourceSharePermissionsInput"}
	if v.ResourceShareArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceShareArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListResourcesInput(v *ListResourcesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListResourcesInput"}
	if len(v.ResourceOwner) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceOwner"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPromoteResourceShareCreatedFromPolicyInput(v *PromoteResourceShareCreatedFromPolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PromoteResourceShareCreatedFromPolicyInput"}
	if v.ResourceShareArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceShareArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpRejectResourceShareInvitationInput(v *RejectResourceShareInvitationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RejectResourceShareInvitationInput"}
	if v.ResourceShareInvitationArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceShareInvitationArn"))
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
	if v.ResourceShareArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceShareArn"))
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
	if v.ResourceShareArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceShareArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateResourceShareInput(v *UpdateResourceShareInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateResourceShareInput"}
	if v.ResourceShareArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceShareArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
