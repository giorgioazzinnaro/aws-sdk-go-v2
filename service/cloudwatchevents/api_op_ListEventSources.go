// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchevents

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchevents/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// You can use this to see all the partner event sources that have been shared with
// your AWS account. For more information about partner event sources, see
// CreateEventBus.
func (c *Client) ListEventSources(ctx context.Context, params *ListEventSourcesInput, optFns ...func(*Options)) (*ListEventSourcesOutput, error) {
	if params == nil {
		params = &ListEventSourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEventSources", params, optFns, addOperationListEventSourcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEventSourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEventSourcesInput struct {

	// Specifying this limits the number of results returned by this operation. The
	// operation also returns a NextToken which you can use in a subsequent operation
	// to retrieve the next set of results.
	Limit *int32

	// Specifying this limits the results to only those partner event sources with
	// names that start with the specified prefix.
	NamePrefix *string

	// The token returned by a previous call to retrieve the next set of results.
	NextToken *string
}

type ListEventSourcesOutput struct {

	// The list of event sources.
	EventSources []*types.EventSource

	// A token you can use in a subsequent operation to retrieve the next set of
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListEventSourcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListEventSources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListEventSources{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListEventSources(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListEventSources(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "events",
		OperationName: "ListEventSources",
	}
}
