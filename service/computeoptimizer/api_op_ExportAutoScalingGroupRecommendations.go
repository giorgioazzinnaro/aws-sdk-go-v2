// Code generated by smithy-go-codegen DO NOT EDIT.

package computeoptimizer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/computeoptimizer/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Exports optimization recommendations for Auto Scaling groups. Recommendations
// are exported in a comma-separated values (.csv) file, and its metadata in a
// JavaScript Object Notation (.json) file, to an existing Amazon Simple Storage
// Service (Amazon S3) bucket that you specify. For more information, see Exporting
// Recommendations
// (https://docs.aws.amazon.com/compute-optimizer/latest/ug/exporting-recommendations.html)
// in the Compute Optimizer User Guide. You can have only one Auto Scaling group
// export job in progress per AWS Region.
func (c *Client) ExportAutoScalingGroupRecommendations(ctx context.Context, params *ExportAutoScalingGroupRecommendationsInput, optFns ...func(*Options)) (*ExportAutoScalingGroupRecommendationsOutput, error) {
	if params == nil {
		params = &ExportAutoScalingGroupRecommendationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ExportAutoScalingGroupRecommendations", params, optFns, addOperationExportAutoScalingGroupRecommendationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ExportAutoScalingGroupRecommendationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ExportAutoScalingGroupRecommendationsInput struct {

	// An object to specify the destination Amazon Simple Storage Service (Amazon S3)
	// bucket name and key prefix for the export job. You must create the destination
	// Amazon S3 bucket for your recommendations export before you create the export
	// job. Compute Optimizer does not create the S3 bucket for you. After you create
	// the S3 bucket, ensure that it has the required permission policy to allow
	// Compute Optimizer to write the export file to it. If you plan to specify an
	// object prefix when you create the export job, you must include the object prefix
	// in the policy that you add to the S3 bucket. For more information, see Amazon S3
	// Bucket Policy for Compute Optimizer
	// (https://docs.aws.amazon.com/compute-optimizer/latest/ug/create-s3-bucket-policy-for-compute-optimizer.html)
	// in the Compute Optimizer user guide.
	//
	// This member is required.
	S3DestinationConfig *types.S3DestinationConfig

	// The IDs of the AWS accounts for which to export Auto Scaling group
	// recommendations. If your account is the master account of an organization, use
	// this parameter to specify the member accounts for which you want to export
	// recommendations. This parameter cannot be specified together with the include
	// member accounts parameter. The parameters are mutually exclusive.
	// Recommendations for member accounts are not included in the export if this
	// parameter, or the include member accounts parameter, is omitted. You can specify
	// multiple account IDs per request.
	AccountIds []*string

	// The recommendations data to include in the export file. For more information
	// about the fields that can be exported, see Exported files
	// (https://docs.aws.amazon.com/compute-optimizer/latest/ug/exporting-recommendations.html#exported-files)
	// in the Compute Optimizer User Guide.
	FieldsToExport []types.ExportableAutoScalingGroupField

	// The format of the export file. The only export file format currently supported
	// is Csv.
	FileFormat types.FileFormat

	// An array of objects that describe a filter to export a more specific set of Auto
	// Scaling group recommendations.
	Filters []*types.Filter

	// Indicates whether to include recommendations for resources in all member
	// accounts of the organization if your account is the master account of an
	// organization. The member accounts must also be opted in to Compute Optimizer.
	// Recommendations for member accounts of the organization are not included in the
	// export file if this parameter is omitted. This parameter cannot be specified
	// together with the account IDs parameter. The parameters are mutually exclusive.
	// Recommendations for member accounts are not included in the export if this
	// parameter, or the account IDs parameter, is omitted.
	IncludeMemberAccounts *bool
}

type ExportAutoScalingGroupRecommendationsOutput struct {

	// The identification number of the export job. Use the
	// DescribeRecommendationExportJobs action, and specify the job ID to view the
	// status of an export job.
	JobId *string

	// An object that describes the destination Amazon S3 bucket of a recommendations
	// export file.
	S3Destination *types.S3Destination

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationExportAutoScalingGroupRecommendationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson10_serializeOpExportAutoScalingGroupRecommendations{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson10_deserializeOpExportAutoScalingGroupRecommendations{})
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
	if err := addOpExportAutoScalingGroupRecommendationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opExportAutoScalingGroupRecommendations(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opExportAutoScalingGroupRecommendations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "compute-optimizer",
		OperationName: "ExportAutoScalingGroupRecommendations",
	}
}
