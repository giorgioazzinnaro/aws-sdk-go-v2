// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationdiscoveryservice

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/applicationdiscoveryservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
)

type validateOpAssociateConfigurationItemsToApplication struct {
}

func (*validateOpAssociateConfigurationItemsToApplication) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpAssociateConfigurationItemsToApplication) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*AssociateConfigurationItemsToApplicationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpAssociateConfigurationItemsToApplicationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpBatchDeleteImportData struct {
}

func (*validateOpBatchDeleteImportData) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpBatchDeleteImportData) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*BatchDeleteImportDataInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpBatchDeleteImportDataInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateApplication struct {
}

func (*validateOpCreateApplication) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateApplication) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateApplicationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateApplicationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateTags struct {
}

func (*validateOpCreateTags) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateTags) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateTagsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateTagsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteApplications struct {
}

func (*validateOpDeleteApplications) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteApplications) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteApplicationsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteApplicationsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteTags struct {
}

func (*validateOpDeleteTags) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteTags) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteTagsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteTagsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeAgents struct {
}

func (*validateOpDescribeAgents) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeAgents) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeAgentsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeAgentsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeConfigurations struct {
}

func (*validateOpDescribeConfigurations) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeConfigurations) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeConfigurationsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeConfigurationsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeExportTasks struct {
}

func (*validateOpDescribeExportTasks) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeExportTasks) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeExportTasksInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeExportTasksInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeTags struct {
}

func (*validateOpDescribeTags) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeTags) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeTagsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeTagsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDisassociateConfigurationItemsFromApplication struct {
}

func (*validateOpDisassociateConfigurationItemsFromApplication) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDisassociateConfigurationItemsFromApplication) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DisassociateConfigurationItemsFromApplicationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDisassociateConfigurationItemsFromApplicationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListConfigurations struct {
}

func (*validateOpListConfigurations) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListConfigurations) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListConfigurationsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListConfigurationsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListServerNeighbors struct {
}

func (*validateOpListServerNeighbors) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListServerNeighbors) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListServerNeighborsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListServerNeighborsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartDataCollectionByAgentIds struct {
}

func (*validateOpStartDataCollectionByAgentIds) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartDataCollectionByAgentIds) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartDataCollectionByAgentIdsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartDataCollectionByAgentIdsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartExportTask struct {
}

func (*validateOpStartExportTask) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartExportTask) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartExportTaskInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartExportTaskInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartImportTask struct {
}

func (*validateOpStartImportTask) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartImportTask) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartImportTaskInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartImportTaskInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStopContinuousExport struct {
}

func (*validateOpStopContinuousExport) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStopContinuousExport) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StopContinuousExportInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStopContinuousExportInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStopDataCollectionByAgentIds struct {
}

func (*validateOpStopDataCollectionByAgentIds) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStopDataCollectionByAgentIds) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StopDataCollectionByAgentIdsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStopDataCollectionByAgentIdsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateApplication struct {
}

func (*validateOpUpdateApplication) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateApplication) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateApplicationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateApplicationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpAssociateConfigurationItemsToApplicationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpAssociateConfigurationItemsToApplication{})
}

func addOpBatchDeleteImportDataValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpBatchDeleteImportData{})
}

func addOpCreateApplicationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpCreateApplication{})
}

func addOpCreateTagsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpCreateTags{})
}

func addOpDeleteApplicationsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpDeleteApplications{})
}

func addOpDeleteTagsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpDeleteTags{})
}

func addOpDescribeAgentsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpDescribeAgents{})
}

func addOpDescribeConfigurationsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpDescribeConfigurations{})
}

func addOpDescribeExportTasksValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpDescribeExportTasks{})
}

func addOpDescribeTagsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpDescribeTags{})
}

func addOpDisassociateConfigurationItemsFromApplicationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpDisassociateConfigurationItemsFromApplication{})
}

func addOpListConfigurationsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpListConfigurations{})
}

func addOpListServerNeighborsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpListServerNeighbors{})
}

func addOpStartDataCollectionByAgentIdsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpStartDataCollectionByAgentIds{})
}

func addOpStartExportTaskValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpStartExportTask{})
}

func addOpStartImportTaskValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpStartImportTask{})
}

func addOpStopContinuousExportValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpStopContinuousExport{})
}

func addOpStopDataCollectionByAgentIdsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpStopDataCollectionByAgentIds{})
}

func addOpUpdateApplicationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpUpdateApplication{})
}

func validateExportFilter(v *types.ExportFilter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ExportFilter"}
	if v.Values == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Values"))
	}
	if v.Condition == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Condition"))
	}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateExportFilters(v []*types.ExportFilter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ExportFilters"}
	for i := range v {
		if err := validateExportFilter(v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateFilter(v *types.Filter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Filter"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.Values == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Values"))
	}
	if v.Condition == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Condition"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateFilters(v []*types.Filter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Filters"}
	for i := range v {
		if err := validateFilter(v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOrderByElement(v *types.OrderByElement) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "OrderByElement"}
	if v.FieldName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FieldName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOrderByList(v []*types.OrderByElement) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "OrderByList"}
	for i := range v {
		if err := validateOrderByElement(v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
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

func validateTagFilter(v *types.TagFilter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagFilter"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.Values == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Values"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTagFilters(v []*types.TagFilter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagFilters"}
	for i := range v {
		if err := validateTagFilter(v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTagSet(v []*types.Tag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagSet"}
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

func validateOpAssociateConfigurationItemsToApplicationInput(v *AssociateConfigurationItemsToApplicationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AssociateConfigurationItemsToApplicationInput"}
	if v.ApplicationConfigurationId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ApplicationConfigurationId"))
	}
	if v.ConfigurationIds == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ConfigurationIds"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpBatchDeleteImportDataInput(v *BatchDeleteImportDataInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "BatchDeleteImportDataInput"}
	if v.ImportTaskIds == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ImportTaskIds"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateApplicationInput(v *CreateApplicationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateApplicationInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateTagsInput(v *CreateTagsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateTagsInput"}
	if v.ConfigurationIds == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ConfigurationIds"))
	}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	} else if v.Tags != nil {
		if err := validateTagSet(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteApplicationsInput(v *DeleteApplicationsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteApplicationsInput"}
	if v.ConfigurationIds == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ConfigurationIds"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteTagsInput(v *DeleteTagsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteTagsInput"}
	if v.Tags != nil {
		if err := validateTagSet(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
	}
	if v.ConfigurationIds == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ConfigurationIds"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeAgentsInput(v *DescribeAgentsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeAgentsInput"}
	if v.Filters != nil {
		if err := validateFilters(v.Filters); err != nil {
			invalidParams.AddNested("Filters", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeConfigurationsInput(v *DescribeConfigurationsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeConfigurationsInput"}
	if v.ConfigurationIds == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ConfigurationIds"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeExportTasksInput(v *DescribeExportTasksInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeExportTasksInput"}
	if v.Filters != nil {
		if err := validateExportFilters(v.Filters); err != nil {
			invalidParams.AddNested("Filters", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeTagsInput(v *DescribeTagsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeTagsInput"}
	if v.Filters != nil {
		if err := validateTagFilters(v.Filters); err != nil {
			invalidParams.AddNested("Filters", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDisassociateConfigurationItemsFromApplicationInput(v *DisassociateConfigurationItemsFromApplicationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DisassociateConfigurationItemsFromApplicationInput"}
	if v.ApplicationConfigurationId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ApplicationConfigurationId"))
	}
	if v.ConfigurationIds == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ConfigurationIds"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListConfigurationsInput(v *ListConfigurationsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListConfigurationsInput"}
	if v.OrderBy != nil {
		if err := validateOrderByList(v.OrderBy); err != nil {
			invalidParams.AddNested("OrderBy", err.(smithy.InvalidParamsError))
		}
	}
	if len(v.ConfigurationType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ConfigurationType"))
	}
	if v.Filters != nil {
		if err := validateFilters(v.Filters); err != nil {
			invalidParams.AddNested("Filters", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListServerNeighborsInput(v *ListServerNeighborsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListServerNeighborsInput"}
	if v.ConfigurationId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ConfigurationId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartDataCollectionByAgentIdsInput(v *StartDataCollectionByAgentIdsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartDataCollectionByAgentIdsInput"}
	if v.AgentIds == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AgentIds"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartExportTaskInput(v *StartExportTaskInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartExportTaskInput"}
	if v.Filters != nil {
		if err := validateExportFilters(v.Filters); err != nil {
			invalidParams.AddNested("Filters", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartImportTaskInput(v *StartImportTaskInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartImportTaskInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.ImportUrl == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ImportUrl"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStopContinuousExportInput(v *StopContinuousExportInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StopContinuousExportInput"}
	if v.ExportId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ExportId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStopDataCollectionByAgentIdsInput(v *StopDataCollectionByAgentIdsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StopDataCollectionByAgentIdsInput"}
	if v.AgentIds == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AgentIds"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateApplicationInput(v *UpdateApplicationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateApplicationInput"}
	if v.ConfigurationId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ConfigurationId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
