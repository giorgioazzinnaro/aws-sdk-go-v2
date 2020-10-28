// Code generated by smithy-go-codegen DO NOT EDIT.

package firehose

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/firehose/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
)

type validateOpCreateDeliveryStream struct {
}

func (*validateOpCreateDeliveryStream) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateDeliveryStream) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateDeliveryStreamInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateDeliveryStreamInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteDeliveryStream struct {
}

func (*validateOpDeleteDeliveryStream) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteDeliveryStream) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteDeliveryStreamInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteDeliveryStreamInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeDeliveryStream struct {
}

func (*validateOpDescribeDeliveryStream) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeDeliveryStream) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeDeliveryStreamInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeDeliveryStreamInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListTagsForDeliveryStream struct {
}

func (*validateOpListTagsForDeliveryStream) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListTagsForDeliveryStream) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListTagsForDeliveryStreamInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListTagsForDeliveryStreamInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPutRecordBatch struct {
}

func (*validateOpPutRecordBatch) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutRecordBatch) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutRecordBatchInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutRecordBatchInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPutRecord struct {
}

func (*validateOpPutRecord) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutRecord) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutRecordInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutRecordInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartDeliveryStreamEncryption struct {
}

func (*validateOpStartDeliveryStreamEncryption) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartDeliveryStreamEncryption) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartDeliveryStreamEncryptionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartDeliveryStreamEncryptionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStopDeliveryStreamEncryption struct {
}

func (*validateOpStopDeliveryStreamEncryption) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStopDeliveryStreamEncryption) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StopDeliveryStreamEncryptionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStopDeliveryStreamEncryptionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpTagDeliveryStream struct {
}

func (*validateOpTagDeliveryStream) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpTagDeliveryStream) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*TagDeliveryStreamInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpTagDeliveryStreamInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUntagDeliveryStream struct {
}

func (*validateOpUntagDeliveryStream) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUntagDeliveryStream) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UntagDeliveryStreamInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUntagDeliveryStreamInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateDestination struct {
}

func (*validateOpUpdateDestination) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateDestination) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateDestinationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateDestinationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCreateDeliveryStreamValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpCreateDeliveryStream{})
}

func addOpDeleteDeliveryStreamValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpDeleteDeliveryStream{})
}

func addOpDescribeDeliveryStreamValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpDescribeDeliveryStream{})
}

func addOpListTagsForDeliveryStreamValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpListTagsForDeliveryStream{})
}

func addOpPutRecordBatchValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpPutRecordBatch{})
}

func addOpPutRecordValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpPutRecord{})
}

func addOpStartDeliveryStreamEncryptionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpStartDeliveryStreamEncryption{})
}

func addOpStopDeliveryStreamEncryptionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpStopDeliveryStreamEncryption{})
}

func addOpTagDeliveryStreamValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpTagDeliveryStream{})
}

func addOpUntagDeliveryStreamValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpUntagDeliveryStream{})
}

func addOpUpdateDestinationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(middleware.After, &validateOpUpdateDestination{})
}

func validateCopyCommand(v *types.CopyCommand) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CopyCommand"}
	if v.DataTableName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DataTableName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateDeliveryStreamEncryptionConfigurationInput(v *types.DeliveryStreamEncryptionConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeliveryStreamEncryptionConfigurationInput"}
	if len(v.KeyType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("KeyType"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateElasticsearchDestinationConfiguration(v *types.ElasticsearchDestinationConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ElasticsearchDestinationConfiguration"}
	if v.ProcessingConfiguration != nil {
		if err := validateProcessingConfiguration(v.ProcessingConfiguration); err != nil {
			invalidParams.AddNested("ProcessingConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if v.S3Configuration == nil {
		invalidParams.Add(smithy.NewErrParamRequired("S3Configuration"))
	} else if v.S3Configuration != nil {
		if err := validateS3DestinationConfiguration(v.S3Configuration); err != nil {
			invalidParams.AddNested("S3Configuration", err.(smithy.InvalidParamsError))
		}
	}
	if v.VpcConfiguration != nil {
		if err := validateVpcConfiguration(v.VpcConfiguration); err != nil {
			invalidParams.AddNested("VpcConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if v.IndexName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IndexName"))
	}
	if v.RoleARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RoleARN"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateElasticsearchDestinationUpdate(v *types.ElasticsearchDestinationUpdate) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ElasticsearchDestinationUpdate"}
	if v.ProcessingConfiguration != nil {
		if err := validateProcessingConfiguration(v.ProcessingConfiguration); err != nil {
			invalidParams.AddNested("ProcessingConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if v.S3Update != nil {
		if err := validateS3DestinationUpdate(v.S3Update); err != nil {
			invalidParams.AddNested("S3Update", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateEncryptionConfiguration(v *types.EncryptionConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "EncryptionConfiguration"}
	if v.KMSEncryptionConfig != nil {
		if err := validateKMSEncryptionConfig(v.KMSEncryptionConfig); err != nil {
			invalidParams.AddNested("KMSEncryptionConfig", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateExtendedS3DestinationConfiguration(v *types.ExtendedS3DestinationConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ExtendedS3DestinationConfiguration"}
	if v.ProcessingConfiguration != nil {
		if err := validateProcessingConfiguration(v.ProcessingConfiguration); err != nil {
			invalidParams.AddNested("ProcessingConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if v.EncryptionConfiguration != nil {
		if err := validateEncryptionConfiguration(v.EncryptionConfiguration); err != nil {
			invalidParams.AddNested("EncryptionConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if v.RoleARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RoleARN"))
	}
	if v.BucketARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BucketARN"))
	}
	if v.S3BackupConfiguration != nil {
		if err := validateS3DestinationConfiguration(v.S3BackupConfiguration); err != nil {
			invalidParams.AddNested("S3BackupConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateExtendedS3DestinationUpdate(v *types.ExtendedS3DestinationUpdate) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ExtendedS3DestinationUpdate"}
	if v.ProcessingConfiguration != nil {
		if err := validateProcessingConfiguration(v.ProcessingConfiguration); err != nil {
			invalidParams.AddNested("ProcessingConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if v.S3BackupUpdate != nil {
		if err := validateS3DestinationUpdate(v.S3BackupUpdate); err != nil {
			invalidParams.AddNested("S3BackupUpdate", err.(smithy.InvalidParamsError))
		}
	}
	if v.EncryptionConfiguration != nil {
		if err := validateEncryptionConfiguration(v.EncryptionConfiguration); err != nil {
			invalidParams.AddNested("EncryptionConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateHttpEndpointCommonAttribute(v *types.HttpEndpointCommonAttribute) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "HttpEndpointCommonAttribute"}
	if v.AttributeName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AttributeName"))
	}
	if v.AttributeValue == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AttributeValue"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateHttpEndpointCommonAttributesList(v []*types.HttpEndpointCommonAttribute) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "HttpEndpointCommonAttributesList"}
	for i := range v {
		if err := validateHttpEndpointCommonAttribute(v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateHttpEndpointConfiguration(v *types.HttpEndpointConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "HttpEndpointConfiguration"}
	if v.Url == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Url"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateHttpEndpointDestinationConfiguration(v *types.HttpEndpointDestinationConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "HttpEndpointDestinationConfiguration"}
	if v.ProcessingConfiguration != nil {
		if err := validateProcessingConfiguration(v.ProcessingConfiguration); err != nil {
			invalidParams.AddNested("ProcessingConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if v.S3Configuration == nil {
		invalidParams.Add(smithy.NewErrParamRequired("S3Configuration"))
	} else if v.S3Configuration != nil {
		if err := validateS3DestinationConfiguration(v.S3Configuration); err != nil {
			invalidParams.AddNested("S3Configuration", err.(smithy.InvalidParamsError))
		}
	}
	if v.RequestConfiguration != nil {
		if err := validateHttpEndpointRequestConfiguration(v.RequestConfiguration); err != nil {
			invalidParams.AddNested("RequestConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if v.EndpointConfiguration == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EndpointConfiguration"))
	} else if v.EndpointConfiguration != nil {
		if err := validateHttpEndpointConfiguration(v.EndpointConfiguration); err != nil {
			invalidParams.AddNested("EndpointConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateHttpEndpointDestinationUpdate(v *types.HttpEndpointDestinationUpdate) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "HttpEndpointDestinationUpdate"}
	if v.ProcessingConfiguration != nil {
		if err := validateProcessingConfiguration(v.ProcessingConfiguration); err != nil {
			invalidParams.AddNested("ProcessingConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if v.EndpointConfiguration != nil {
		if err := validateHttpEndpointConfiguration(v.EndpointConfiguration); err != nil {
			invalidParams.AddNested("EndpointConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if v.RequestConfiguration != nil {
		if err := validateHttpEndpointRequestConfiguration(v.RequestConfiguration); err != nil {
			invalidParams.AddNested("RequestConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if v.S3Update != nil {
		if err := validateS3DestinationUpdate(v.S3Update); err != nil {
			invalidParams.AddNested("S3Update", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateHttpEndpointRequestConfiguration(v *types.HttpEndpointRequestConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "HttpEndpointRequestConfiguration"}
	if v.CommonAttributes != nil {
		if err := validateHttpEndpointCommonAttributesList(v.CommonAttributes); err != nil {
			invalidParams.AddNested("CommonAttributes", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateKinesisStreamSourceConfiguration(v *types.KinesisStreamSourceConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "KinesisStreamSourceConfiguration"}
	if v.KinesisStreamARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("KinesisStreamARN"))
	}
	if v.RoleARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RoleARN"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateKMSEncryptionConfig(v *types.KMSEncryptionConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "KMSEncryptionConfig"}
	if v.AWSKMSKeyARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AWSKMSKeyARN"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateProcessingConfiguration(v *types.ProcessingConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ProcessingConfiguration"}
	if v.Processors != nil {
		if err := validateProcessorList(v.Processors); err != nil {
			invalidParams.AddNested("Processors", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateProcessor(v *types.Processor) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Processor"}
	if len(v.Type) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Type"))
	}
	if v.Parameters != nil {
		if err := validateProcessorParameterList(v.Parameters); err != nil {
			invalidParams.AddNested("Parameters", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateProcessorList(v []*types.Processor) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ProcessorList"}
	for i := range v {
		if err := validateProcessor(v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateProcessorParameter(v *types.ProcessorParameter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ProcessorParameter"}
	if len(v.ParameterName) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ParameterName"))
	}
	if v.ParameterValue == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ParameterValue"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateProcessorParameterList(v []*types.ProcessorParameter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ProcessorParameterList"}
	for i := range v {
		if err := validateProcessorParameter(v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validatePutRecordBatchRequestEntryList(v []*types.Record) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutRecordBatchRequestEntryList"}
	for i := range v {
		if err := validateRecord(v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateRecord(v *types.Record) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Record"}
	if v.Data == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Data"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateRedshiftDestinationConfiguration(v *types.RedshiftDestinationConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RedshiftDestinationConfiguration"}
	if v.CopyCommand == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CopyCommand"))
	} else if v.CopyCommand != nil {
		if err := validateCopyCommand(v.CopyCommand); err != nil {
			invalidParams.AddNested("CopyCommand", err.(smithy.InvalidParamsError))
		}
	}
	if v.ClusterJDBCURL == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClusterJDBCURL"))
	}
	if v.Password == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Password"))
	}
	if v.Username == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Username"))
	}
	if v.S3Configuration == nil {
		invalidParams.Add(smithy.NewErrParamRequired("S3Configuration"))
	} else if v.S3Configuration != nil {
		if err := validateS3DestinationConfiguration(v.S3Configuration); err != nil {
			invalidParams.AddNested("S3Configuration", err.(smithy.InvalidParamsError))
		}
	}
	if v.S3BackupConfiguration != nil {
		if err := validateS3DestinationConfiguration(v.S3BackupConfiguration); err != nil {
			invalidParams.AddNested("S3BackupConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if v.RoleARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RoleARN"))
	}
	if v.ProcessingConfiguration != nil {
		if err := validateProcessingConfiguration(v.ProcessingConfiguration); err != nil {
			invalidParams.AddNested("ProcessingConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateRedshiftDestinationUpdate(v *types.RedshiftDestinationUpdate) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RedshiftDestinationUpdate"}
	if v.CopyCommand != nil {
		if err := validateCopyCommand(v.CopyCommand); err != nil {
			invalidParams.AddNested("CopyCommand", err.(smithy.InvalidParamsError))
		}
	}
	if v.S3BackupUpdate != nil {
		if err := validateS3DestinationUpdate(v.S3BackupUpdate); err != nil {
			invalidParams.AddNested("S3BackupUpdate", err.(smithy.InvalidParamsError))
		}
	}
	if v.S3Update != nil {
		if err := validateS3DestinationUpdate(v.S3Update); err != nil {
			invalidParams.AddNested("S3Update", err.(smithy.InvalidParamsError))
		}
	}
	if v.ProcessingConfiguration != nil {
		if err := validateProcessingConfiguration(v.ProcessingConfiguration); err != nil {
			invalidParams.AddNested("ProcessingConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateS3DestinationConfiguration(v *types.S3DestinationConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "S3DestinationConfiguration"}
	if v.EncryptionConfiguration != nil {
		if err := validateEncryptionConfiguration(v.EncryptionConfiguration); err != nil {
			invalidParams.AddNested("EncryptionConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if v.BucketARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BucketARN"))
	}
	if v.RoleARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RoleARN"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateS3DestinationUpdate(v *types.S3DestinationUpdate) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "S3DestinationUpdate"}
	if v.EncryptionConfiguration != nil {
		if err := validateEncryptionConfiguration(v.EncryptionConfiguration); err != nil {
			invalidParams.AddNested("EncryptionConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateSplunkDestinationConfiguration(v *types.SplunkDestinationConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SplunkDestinationConfiguration"}
	if v.S3Configuration == nil {
		invalidParams.Add(smithy.NewErrParamRequired("S3Configuration"))
	} else if v.S3Configuration != nil {
		if err := validateS3DestinationConfiguration(v.S3Configuration); err != nil {
			invalidParams.AddNested("S3Configuration", err.(smithy.InvalidParamsError))
		}
	}
	if v.HECEndpoint == nil {
		invalidParams.Add(smithy.NewErrParamRequired("HECEndpoint"))
	}
	if v.ProcessingConfiguration != nil {
		if err := validateProcessingConfiguration(v.ProcessingConfiguration); err != nil {
			invalidParams.AddNested("ProcessingConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if v.HECToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("HECToken"))
	}
	if len(v.HECEndpointType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("HECEndpointType"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateSplunkDestinationUpdate(v *types.SplunkDestinationUpdate) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SplunkDestinationUpdate"}
	if v.S3Update != nil {
		if err := validateS3DestinationUpdate(v.S3Update); err != nil {
			invalidParams.AddNested("S3Update", err.(smithy.InvalidParamsError))
		}
	}
	if v.ProcessingConfiguration != nil {
		if err := validateProcessingConfiguration(v.ProcessingConfiguration); err != nil {
			invalidParams.AddNested("ProcessingConfiguration", err.(smithy.InvalidParamsError))
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
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTagDeliveryStreamInputTagList(v []*types.Tag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagDeliveryStreamInputTagList"}
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

func validateVpcConfiguration(v *types.VpcConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "VpcConfiguration"}
	if v.SecurityGroupIds == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SecurityGroupIds"))
	}
	if v.SubnetIds == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SubnetIds"))
	}
	if v.RoleARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RoleARN"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateDeliveryStreamInput(v *CreateDeliveryStreamInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateDeliveryStreamInput"}
	if v.S3DestinationConfiguration != nil {
		if err := validateS3DestinationConfiguration(v.S3DestinationConfiguration); err != nil {
			invalidParams.AddNested("S3DestinationConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if v.DeliveryStreamEncryptionConfigurationInput != nil {
		if err := validateDeliveryStreamEncryptionConfigurationInput(v.DeliveryStreamEncryptionConfigurationInput); err != nil {
			invalidParams.AddNested("DeliveryStreamEncryptionConfigurationInput", err.(smithy.InvalidParamsError))
		}
	}
	if v.DeliveryStreamName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DeliveryStreamName"))
	}
	if v.KinesisStreamSourceConfiguration != nil {
		if err := validateKinesisStreamSourceConfiguration(v.KinesisStreamSourceConfiguration); err != nil {
			invalidParams.AddNested("KinesisStreamSourceConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if v.ExtendedS3DestinationConfiguration != nil {
		if err := validateExtendedS3DestinationConfiguration(v.ExtendedS3DestinationConfiguration); err != nil {
			invalidParams.AddNested("ExtendedS3DestinationConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if v.SplunkDestinationConfiguration != nil {
		if err := validateSplunkDestinationConfiguration(v.SplunkDestinationConfiguration); err != nil {
			invalidParams.AddNested("SplunkDestinationConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if v.Tags != nil {
		if err := validateTagDeliveryStreamInputTagList(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
	}
	if v.HttpEndpointDestinationConfiguration != nil {
		if err := validateHttpEndpointDestinationConfiguration(v.HttpEndpointDestinationConfiguration); err != nil {
			invalidParams.AddNested("HttpEndpointDestinationConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if v.RedshiftDestinationConfiguration != nil {
		if err := validateRedshiftDestinationConfiguration(v.RedshiftDestinationConfiguration); err != nil {
			invalidParams.AddNested("RedshiftDestinationConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if v.ElasticsearchDestinationConfiguration != nil {
		if err := validateElasticsearchDestinationConfiguration(v.ElasticsearchDestinationConfiguration); err != nil {
			invalidParams.AddNested("ElasticsearchDestinationConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteDeliveryStreamInput(v *DeleteDeliveryStreamInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteDeliveryStreamInput"}
	if v.DeliveryStreamName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DeliveryStreamName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeDeliveryStreamInput(v *DescribeDeliveryStreamInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeDeliveryStreamInput"}
	if v.DeliveryStreamName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DeliveryStreamName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListTagsForDeliveryStreamInput(v *ListTagsForDeliveryStreamInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListTagsForDeliveryStreamInput"}
	if v.DeliveryStreamName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DeliveryStreamName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPutRecordBatchInput(v *PutRecordBatchInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutRecordBatchInput"}
	if v.DeliveryStreamName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DeliveryStreamName"))
	}
	if v.Records == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Records"))
	} else if v.Records != nil {
		if err := validatePutRecordBatchRequestEntryList(v.Records); err != nil {
			invalidParams.AddNested("Records", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPutRecordInput(v *PutRecordInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutRecordInput"}
	if v.Record == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Record"))
	} else if v.Record != nil {
		if err := validateRecord(v.Record); err != nil {
			invalidParams.AddNested("Record", err.(smithy.InvalidParamsError))
		}
	}
	if v.DeliveryStreamName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DeliveryStreamName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartDeliveryStreamEncryptionInput(v *StartDeliveryStreamEncryptionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartDeliveryStreamEncryptionInput"}
	if v.DeliveryStreamName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DeliveryStreamName"))
	}
	if v.DeliveryStreamEncryptionConfigurationInput != nil {
		if err := validateDeliveryStreamEncryptionConfigurationInput(v.DeliveryStreamEncryptionConfigurationInput); err != nil {
			invalidParams.AddNested("DeliveryStreamEncryptionConfigurationInput", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStopDeliveryStreamEncryptionInput(v *StopDeliveryStreamEncryptionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StopDeliveryStreamEncryptionInput"}
	if v.DeliveryStreamName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DeliveryStreamName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpTagDeliveryStreamInput(v *TagDeliveryStreamInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagDeliveryStreamInput"}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	} else if v.Tags != nil {
		if err := validateTagDeliveryStreamInputTagList(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
	}
	if v.DeliveryStreamName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DeliveryStreamName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUntagDeliveryStreamInput(v *UntagDeliveryStreamInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UntagDeliveryStreamInput"}
	if v.DeliveryStreamName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DeliveryStreamName"))
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

func validateOpUpdateDestinationInput(v *UpdateDestinationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateDestinationInput"}
	if v.DestinationId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DestinationId"))
	}
	if v.S3DestinationUpdate != nil {
		if err := validateS3DestinationUpdate(v.S3DestinationUpdate); err != nil {
			invalidParams.AddNested("S3DestinationUpdate", err.(smithy.InvalidParamsError))
		}
	}
	if v.HttpEndpointDestinationUpdate != nil {
		if err := validateHttpEndpointDestinationUpdate(v.HttpEndpointDestinationUpdate); err != nil {
			invalidParams.AddNested("HttpEndpointDestinationUpdate", err.(smithy.InvalidParamsError))
		}
	}
	if v.DeliveryStreamName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DeliveryStreamName"))
	}
	if v.RedshiftDestinationUpdate != nil {
		if err := validateRedshiftDestinationUpdate(v.RedshiftDestinationUpdate); err != nil {
			invalidParams.AddNested("RedshiftDestinationUpdate", err.(smithy.InvalidParamsError))
		}
	}
	if v.CurrentDeliveryStreamVersionId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CurrentDeliveryStreamVersionId"))
	}
	if v.ExtendedS3DestinationUpdate != nil {
		if err := validateExtendedS3DestinationUpdate(v.ExtendedS3DestinationUpdate); err != nil {
			invalidParams.AddNested("ExtendedS3DestinationUpdate", err.(smithy.InvalidParamsError))
		}
	}
	if v.ElasticsearchDestinationUpdate != nil {
		if err := validateElasticsearchDestinationUpdate(v.ElasticsearchDestinationUpdate); err != nil {
			invalidParams.AddNested("ElasticsearchDestinationUpdate", err.(smithy.InvalidParamsError))
		}
	}
	if v.SplunkDestinationUpdate != nil {
		if err := validateSplunkDestinationUpdate(v.SplunkDestinationUpdate); err != nil {
			invalidParams.AddNested("SplunkDestinationUpdate", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
