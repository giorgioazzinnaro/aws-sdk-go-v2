// Code generated by smithy-go-codegen DO NOT EDIT.

package migrationhub

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/migrationhub/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists progress update streams associated with the user account making this call.
func (c *Client) ListProgressUpdateStreams(ctx context.Context, params *ListProgressUpdateStreamsInput, optFns ...func(*Options)) (*ListProgressUpdateStreamsOutput, error) {
	if params == nil {
		params = &ListProgressUpdateStreamsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListProgressUpdateStreams", params, optFns, addOperationListProgressUpdateStreamsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListProgressUpdateStreamsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListProgressUpdateStreamsInput struct {

	// Filter to limit the maximum number of results to list per page.
	MaxResults *int32

	// If a NextToken was returned by a previous call, there are more results
	// available. To retrieve the next page of results, make the call again using the
	// returned token in NextToken.
	NextToken *string
}

type ListProgressUpdateStreamsOutput struct {

	// If there are more streams created than the max result, return the next token to
	// be passed to the next call as a bookmark of where to start from.
	NextToken *string

	// List of progress update streams up to the max number of results passed in the
	// input.
	ProgressUpdateStreamSummaryList []*types.ProgressUpdateStreamSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListProgressUpdateStreamsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListProgressUpdateStreams{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListProgressUpdateStreams{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListProgressUpdateStreams(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListProgressUpdateStreams(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mgh",
		OperationName: "ListProgressUpdateStreams",
	}
}
