// Code generated by smithy-go-codegen DO NOT EDIT.

package forecast

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/forecast/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of forecast export jobs created using the CreateForecastExportJob
// operation. For each forecast export job, this operation returns a summary of its
// properties, including its Amazon Resource Name (ARN). To retrieve the complete
// set of properties, use the ARN with the DescribeForecastExportJob operation. You
// can filter the list using an array of Filter objects.
func (c *Client) ListForecastExportJobs(ctx context.Context, params *ListForecastExportJobsInput, optFns ...func(*Options)) (*ListForecastExportJobsOutput, error) {
	if params == nil {
		params = &ListForecastExportJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListForecastExportJobs", params, optFns, addOperationListForecastExportJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListForecastExportJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListForecastExportJobsInput struct {

	// An array of filters. For each filter, you provide a condition and a match
	// statement. The condition is either IS or IS_NOT, which specifies whether to
	// include or exclude the forecast export jobs that match the statement from the
	// list, respectively. The match statement consists of a key and a value. Filter
	// properties
	//
	//     * Condition - The condition to apply. Valid values are IS and
	// IS_NOT. To include the forecast export jobs that match the statement, specify
	// IS. To exclude matching forecast export jobs, specify IS_NOT.
	//
	//     * Key - The
	// name of the parameter to filter on. Valid values are ForecastArn and Status.
	//
	//
	// * Value - The value to match.
	//
	// For example, to list all jobs that export a
	// forecast named electricityforecast, specify the following filter: "Filters": [ {
	// "Condition": "IS", "Key": "ForecastArn", "Value":
	// "arn:aws:forecast:us-west-2::forecast/electricityforecast" } ]
	Filters []*types.Filter

	// The number of items to return in the response.
	MaxResults *int32

	// If the result of the previous request was truncated, the response includes a
	// NextToken. To retrieve the next set of results, use the token in the next
	// request. Tokens expire after 24 hours.
	NextToken *string
}

type ListForecastExportJobsOutput struct {

	// An array of objects that summarize each export job's properties.
	ForecastExportJobs []*types.ForecastExportJobSummary

	// If the response is truncated, Amazon Forecast returns this token. To retrieve
	// the next set of results, use the token in the next request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListForecastExportJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpListForecastExportJobs{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpListForecastExportJobs{})
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
	if err := addOpListForecastExportJobsValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opListForecastExportJobs(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opListForecastExportJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "forecast",
		OperationName: "ListForecastExportJobs",
	}
}
