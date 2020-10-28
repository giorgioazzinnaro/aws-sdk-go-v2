// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets a list of distribution IDs for distributions that have a cache behavior
// that’s associated with the specified cache policy. You can optionally specify
// the maximum number of items to receive in the response. If the total number of
// items in the list exceeds the maximum that you specify, or the default maximum,
// the response is paginated. To get the next page of items, send a subsequent
// request that specifies the NextMarker value from the current response as the
// Marker value in the subsequent request.
func (c *Client) ListDistributionsByCachePolicyId(ctx context.Context, params *ListDistributionsByCachePolicyIdInput, optFns ...func(*Options)) (*ListDistributionsByCachePolicyIdOutput, error) {
	if params == nil {
		params = &ListDistributionsByCachePolicyIdInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDistributionsByCachePolicyId", params, optFns, addOperationListDistributionsByCachePolicyIdMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDistributionsByCachePolicyIdOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDistributionsByCachePolicyIdInput struct {

	// The ID of the cache policy whose associated distribution IDs you want to list.
	//
	// This member is required.
	CachePolicyId *string

	// Use this field when paginating results to indicate where to begin in your list
	// of distribution IDs. The response includes distribution IDs in the list that
	// occur after the marker. To get the next page of the list, set this field’s value
	// to the value of NextMarker from the current page’s response.
	Marker *string

	// The maximum number of distribution IDs that you want in the response.
	MaxItems *string
}

type ListDistributionsByCachePolicyIdOutput struct {

	// A list of distribution IDs.
	DistributionIdList *types.DistributionIdList

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListDistributionsByCachePolicyIdMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestxml_serializeOpListDistributionsByCachePolicyId{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestxml_deserializeOpListDistributionsByCachePolicyId{})
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
	if err := addOpListDistributionsByCachePolicyIdValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opListDistributionsByCachePolicyId(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opListDistributionsByCachePolicyId(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "ListDistributionsByCachePolicyId",
	}
}
