// Code generated by smithy-go-codegen DO NOT EDIT.

package xray

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/xray/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Get an aggregation of service statistics defined by a specific time range.
func (c *Client) GetTimeSeriesServiceStatistics(ctx context.Context, params *GetTimeSeriesServiceStatisticsInput, optFns ...func(*Options)) (*GetTimeSeriesServiceStatisticsOutput, error) {
	if params == nil {
		params = &GetTimeSeriesServiceStatisticsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetTimeSeriesServiceStatistics", params, optFns, addOperationGetTimeSeriesServiceStatisticsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetTimeSeriesServiceStatisticsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTimeSeriesServiceStatisticsInput struct {

	// The end of the time frame for which to aggregate statistics.
	//
	// This member is required.
	EndTime *time.Time

	// The start of the time frame for which to aggregate statistics.
	//
	// This member is required.
	StartTime *time.Time

	// A filter expression defining entities that will be aggregated for statistics.
	// Supports ID, service, and edge functions. If no selector expression is
	// specified, edge statistics are returned.
	EntitySelectorExpression *string

	// The Amazon Resource Name (ARN) of the group for which to pull statistics from.
	GroupARN *string

	// The case-sensitive name of the group for which to pull statistics from.
	GroupName *string

	// Pagination token.
	NextToken *string

	// Aggregation period in seconds.
	Period *int32
}

type GetTimeSeriesServiceStatisticsOutput struct {

	// A flag indicating whether or not a group's filter expression has been
	// consistent, or if a returned aggregation might show statistics from an older
	// version of the group's filter expression.
	ContainsOldGroupVersions *bool

	// Pagination token.
	NextToken *string

	// The collection of statistics.
	TimeSeriesServiceStatistics []*types.TimeSeriesServiceStatistics

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetTimeSeriesServiceStatisticsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpGetTimeSeriesServiceStatistics{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpGetTimeSeriesServiceStatistics{})
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
	if err := addOpGetTimeSeriesServiceStatisticsValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opGetTimeSeriesServiceStatistics(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opGetTimeSeriesServiceStatistics(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "xray",
		OperationName: "GetTimeSeriesServiceStatistics",
	}
}
