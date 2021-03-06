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

// Lists the subscription filters for the specified log group. You can list all the
// subscription filters or filter the results by prefix. The results are
// ASCII-sorted by filter name.
func (c *Client) DescribeSubscriptionFilters(ctx context.Context, params *DescribeSubscriptionFiltersInput, optFns ...func(*Options)) (*DescribeSubscriptionFiltersOutput, error) {
	if params == nil {
		params = &DescribeSubscriptionFiltersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeSubscriptionFilters", params, optFns, addOperationDescribeSubscriptionFiltersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeSubscriptionFiltersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeSubscriptionFiltersInput struct {

	// The name of the log group.
	//
	// This member is required.
	LogGroupName *string

	// The prefix to match. If you don't specify a value, no prefix filter is applied.
	FilterNamePrefix *string

	// The maximum number of items returned. If you don't specify a value, the default
	// is up to 50 items.
	Limit *int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string
}

type DescribeSubscriptionFiltersOutput struct {

	// The token for the next set of items to return. The token expires after 24 hours.
	NextToken *string

	// The subscription filters.
	SubscriptionFilters []*types.SubscriptionFilter

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeSubscriptionFiltersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeSubscriptionFilters{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeSubscriptionFilters{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpDescribeSubscriptionFiltersValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSubscriptionFilters(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeSubscriptionFilters(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "logs",
		OperationName: "DescribeSubscriptionFilters",
	}
}
