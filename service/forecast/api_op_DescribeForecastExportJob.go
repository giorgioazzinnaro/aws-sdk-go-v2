// Code generated by smithy-go-codegen DO NOT EDIT.

package forecast

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/forecast/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Describes a forecast export job created using the CreateForecastExportJob
// operation. In addition to listing the properties provided by the user in the
// CreateForecastExportJob request, this operation lists the following
// properties:
//
//     * CreationTime
//
//     * LastModificationTime
//
//     * Status
//
//     *
// Message - If an error occurred, information about the error.
func (c *Client) DescribeForecastExportJob(ctx context.Context, params *DescribeForecastExportJobInput, optFns ...func(*Options)) (*DescribeForecastExportJobOutput, error) {
	if params == nil {
		params = &DescribeForecastExportJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeForecastExportJob", params, optFns, addOperationDescribeForecastExportJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeForecastExportJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeForecastExportJobInput struct {

	// The Amazon Resource Name (ARN) of the forecast export job.
	//
	// This member is required.
	ForecastExportJobArn *string
}

type DescribeForecastExportJobOutput struct {

	// When the forecast export job was created.
	CreationTime *time.Time

	// The path to the Amazon Simple Storage Service (Amazon S3) bucket where the
	// forecast is exported.
	Destination *types.DataDestination

	// The Amazon Resource Name (ARN) of the exported forecast.
	ForecastArn *string

	// The ARN of the forecast export job.
	ForecastExportJobArn *string

	// The name of the forecast export job.
	ForecastExportJobName *string

	// When the last successful export job finished.
	LastModificationTime *time.Time

	// If an error occurred, an informational message about the error.
	Message *string

	// The status of the forecast export job. States include:
	//
	//     * ACTIVE
	//
	//     *
	// CREATE_PENDING, CREATE_IN_PROGRESS, CREATE_FAILED
	//
	//     * DELETE_PENDING,
	// DELETE_IN_PROGRESS, DELETE_FAILED
	//
	// The Status of the forecast export job must be
	// ACTIVE before you can access the forecast in your S3 bucket.
	Status *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeForecastExportJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpDescribeForecastExportJob{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpDescribeForecastExportJob{})
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
	if err := addOpDescribeForecastExportJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDescribeForecastExportJob(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDescribeForecastExportJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "forecast",
		OperationName: "DescribeForecastExportJob",
	}
}
