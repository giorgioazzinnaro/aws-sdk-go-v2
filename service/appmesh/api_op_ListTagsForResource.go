// Code generated by smithy-go-codegen DO NOT EDIT.

package appmesh

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appmesh/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// List the tags for an App Mesh resource.
func (c *Client) ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) {
	if params == nil {
		params = &ListTagsForResourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTagsForResource", params, optFns, addOperationListTagsForResourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTagsForResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type ListTagsForResourceInput struct {

	// The Amazon Resource Name (ARN) that identifies the resource to list the tags
	// for.
	//
	// This member is required.
	ResourceArn *string

	// The maximum number of tag results returned by ListTagsForResource in paginated
	// output. When this parameter is used, ListTagsForResource returns only limit
	// results in a single page along with a nextToken response element. You can see
	// the remaining results of the initial request by sending another
	// ListTagsForResource request with the returned nextToken value. This value can be
	// between 1 and 100. If you don't use this parameter, ListTagsForResource returns
	// up to 100 results and a nextToken value if applicable.
	Limit *int32

	// The nextToken value returned from a previous paginated ListTagsForResource
	// request where limit was used and the results exceeded the value of that
	// parameter. Pagination continues from the end of the previous results that
	// returned the nextToken value.
	NextToken *string
}

//
type ListTagsForResourceOutput struct {

	// The tags for the resource.
	//
	// This member is required.
	Tags []*types.TagRef

	// The nextToken value to include in a future ListTagsForResource request. When the
	// results of a ListTagsForResource request exceed limit, you can use this value to
	// retrieve the next page of results. This value is null when there are no more
	// results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListTagsForResourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpListTagsForResource{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpListTagsForResource{})
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
	if err := addOpListTagsForResourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opListTagsForResource(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opListTagsForResource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appmesh",
		OperationName: "ListTagsForResource",
	}
}
