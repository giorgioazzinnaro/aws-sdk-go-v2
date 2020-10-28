// Code generated by smithy-go-codegen DO NOT EDIT.

package migrationhubconfig

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/migrationhubconfig/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
)

type validateOpCreateHomeRegionControl struct {
}

func (*validateOpCreateHomeRegionControl) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateHomeRegionControl) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateHomeRegionControlInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateHomeRegionControlInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeHomeRegionControls struct {
}

func (*validateOpDescribeHomeRegionControls) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeHomeRegionControls) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeHomeRegionControlsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeHomeRegionControlsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCreateHomeRegionControlValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpCreateHomeRegionControl{})
}

func addOpDescribeHomeRegionControlsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpDescribeHomeRegionControls{})
}

func validateTarget(v *types.Target) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Target"}
	if len(v.Type) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Type"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateHomeRegionControlInput(v *CreateHomeRegionControlInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateHomeRegionControlInput"}
	if v.HomeRegion == nil {
		invalidParams.Add(smithy.NewErrParamRequired("HomeRegion"))
	}
	if v.Target == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Target"))
	} else if v.Target != nil {
		if err := validateTarget(v.Target); err != nil {
			invalidParams.AddNested("Target", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeHomeRegionControlsInput(v *DescribeHomeRegionControlsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeHomeRegionControlsInput"}
	if v.Target != nil {
		if err := validateTarget(v.Target); err != nil {
			invalidParams.AddNested("Target", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
