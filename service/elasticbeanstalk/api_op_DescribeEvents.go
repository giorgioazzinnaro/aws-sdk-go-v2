// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns list of event descriptions matching criteria up to the last 6 weeks.
// This action returns the most recent 1,000 events from the specified NextToken.
func (c *Client) DescribeEvents(ctx context.Context, params *DescribeEventsInput, optFns ...func(*Options)) (*DescribeEventsOutput, error) {
	if params == nil {
		params = &DescribeEventsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEvents", params, optFns, addOperationDescribeEventsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEventsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request to retrieve a list of events for an environment.
type DescribeEventsInput struct {

	// If specified, AWS Elastic Beanstalk restricts the returned descriptions to
	// include only those associated with this application.
	ApplicationName *string

	// If specified, AWS Elastic Beanstalk restricts the returned descriptions to those
	// that occur up to, but not including, the EndTime.
	EndTime *time.Time

	// If specified, AWS Elastic Beanstalk restricts the returned descriptions to those
	// associated with this environment.
	EnvironmentId *string

	// If specified, AWS Elastic Beanstalk restricts the returned descriptions to those
	// associated with this environment.
	EnvironmentName *string

	// Specifies the maximum number of events that can be returned, beginning with the
	// most recent event.
	MaxRecords *int32

	// Pagination token. If specified, the events return the next batch of results.
	NextToken *string

	// The ARN of a custom platform version. If specified, AWS Elastic Beanstalk
	// restricts the returned descriptions to those associated with this custom
	// platform version.
	PlatformArn *string

	// If specified, AWS Elastic Beanstalk restricts the described events to include
	// only those associated with this request ID.
	RequestId *string

	// If specified, limits the events returned from this call to include only those
	// with the specified severity or higher.
	Severity types.EventSeverity

	// If specified, AWS Elastic Beanstalk restricts the returned descriptions to those
	// that occur on or after this time.
	StartTime *time.Time

	// If specified, AWS Elastic Beanstalk restricts the returned descriptions to those
	// that are associated with this environment configuration.
	TemplateName *string

	// If specified, AWS Elastic Beanstalk restricts the returned descriptions to those
	// associated with this application version.
	VersionLabel *string
}

// Result message wrapping a list of event descriptions.
type DescribeEventsOutput struct {

	// A list of EventDescription.
	Events []*types.EventDescription

	// If returned, this indicates that there are more results to obtain. Use this
	// token in the next DescribeEvents call to get the next batch of events.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeEventsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsquery_serializeOpDescribeEvents{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsquery_deserializeOpDescribeEvents{})
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
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDescribeEvents(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDescribeEvents(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticbeanstalk",
		OperationName: "DescribeEvents",
	}
}
