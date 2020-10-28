// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchlogs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists log events from the specified log group. You can list all the log events
// or filter the results using a filter pattern, a time range, and the name of the
// log stream. By default, this operation returns as many log events as can fit in
// 1 MB (up to 10,000 log events) or all the events found within the time range
// that you specify. If the results include a token, then there are more log events
// available, and you can get additional results by specifying the token in a
// subsequent call. This operation can return empty results while there are more
// log events available through the token. The returned log events are sorted by
// event timestamp, the timestamp when the event was ingested by CloudWatch Logs,
// and the ID of the PutLogEvents request.
func (c *Client) FilterLogEvents(ctx context.Context, params *FilterLogEventsInput, optFns ...func(*Options)) (*FilterLogEventsOutput, error) {
	if params == nil {
		params = &FilterLogEventsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "FilterLogEvents", params, optFns, addOperationFilterLogEventsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*FilterLogEventsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type FilterLogEventsInput struct {

	// The name of the log group to search.
	//
	// This member is required.
	LogGroupName *string

	// The end of the time range, expressed as the number of milliseconds after Jan 1,
	// 1970 00:00:00 UTC. Events with a timestamp later than this time are not
	// returned.
	EndTime *int64

	// The filter pattern to use. For more information, see Filter and Pattern Syntax
	// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/FilterAndPatternSyntax.html).
	// If not provided, all the events are matched.
	FilterPattern *string

	// If the value is true, the operation makes a best effort to provide responses
	// that contain events from multiple log streams within the log group, interleaved
	// in a single response. If the value is false, all the matched log events in the
	// first log stream are searched first, then those in the next log stream, and so
	// on. The default is false. Important: Starting on June 17, 2019, this parameter
	// is ignored and the value is assumed to be true. The response from this operation
	// always interleaves events from multiple log streams within a log group.
	Interleaved *bool

	// The maximum number of events to return. The default is 10,000 events.
	Limit *int32

	// Filters the results to include only events from log streams that have names
	// starting with this prefix. If you specify a value for both logStreamNamePrefix
	// and logStreamNames, but the value for logStreamNamePrefix does not match any log
	// stream names specified in logStreamNames, the action returns an
	// InvalidParameterException error.
	LogStreamNamePrefix *string

	// Filters the results to only logs from the log streams in this list. If you
	// specify a value for both logStreamNamePrefix and logStreamNames, the action
	// returns an InvalidParameterException error.
	LogStreamNames []*string

	// The token for the next set of events to return. (You received this token from a
	// previous call.)
	NextToken *string

	// The start of the time range, expressed as the number of milliseconds after Jan
	// 1, 1970 00:00:00 UTC. Events with a timestamp before this time are not returned.
	// If you omit startTime and endTime the most recent log events are retrieved, to
	// up 1 MB or 10,000 log events.
	StartTime *int64
}

type FilterLogEventsOutput struct {

	// The matched events.
	Events []*types.FilteredLogEvent

	// The token to use when requesting the next set of items. The token expires after
	// 24 hours.
	NextToken *string

	// IMPORTANT Starting on May 15, 2020, this parameter will be deprecated. This
	// parameter will be an empty list after the deprecation occurs. Indicates which
	// log streams have been searched and whether each has been searched completely.
	SearchedLogStreams []*types.SearchedLogStream

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationFilterLogEventsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpFilterLogEvents{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpFilterLogEvents{})
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
	if err := addOpFilterLogEventsValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opFilterLogEvents(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opFilterLogEvents(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "logs",
		OperationName: "FilterLogEvents",
	}
}
