// Code generated by smithy-go-codegen DO NOT EDIT.

package codeguruprofiler

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codeguruprofiler/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists profiling groups.
func (c *Client) ListProfilingGroups(ctx context.Context, params *ListProfilingGroupsInput, optFns ...func(*Options)) (*ListProfilingGroupsOutput, error) {
	if params == nil {
		params = &ListProfilingGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListProfilingGroups", params, optFns, addOperationListProfilingGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListProfilingGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The structure representing the listProfilingGroupsRequest.
type ListProfilingGroupsInput struct {

	// A Boolean value indicating whether to include a description.
	IncludeDescription *bool

	// The maximum number of profiling groups results returned by ListProfilingGroups
	// in paginated output. When this parameter is used, ListProfilingGroups only
	// returns maxResults results in a single page along with a nextToken response
	// element. The remaining results of the initial request can be seen by sending
	// another ListProfilingGroups request with the returned nextToken value.
	MaxResults *int32

	// The nextToken value returned from a previous paginated ListProfilingGroups
	// request where maxResults was used and the results exceeded the value of that
	// parameter. Pagination continues from the end of the previous results that
	// returned the nextToken value. This token should be treated as an opaque
	// identifier that is only used to retrieve the next items in a list and not for
	// other programmatic purposes.
	NextToken *string
}

// The structure representing the listProfilingGroupsResponse.
type ListProfilingGroupsOutput struct {

	// Information about profiling group names.
	//
	// This member is required.
	ProfilingGroupNames []*string

	// The nextToken value to include in a future ListProfilingGroups request. When the
	// results of a ListProfilingGroups request exceed maxResults, this value can be
	// used to retrieve the next page of results. This value is null when there are no
	// more results to return.
	NextToken *string

	// Information about profiling groups.
	ProfilingGroups []*types.ProfilingGroupDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListProfilingGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpListProfilingGroups{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpListProfilingGroups{})
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
	if err := addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err := addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	return nil
}
