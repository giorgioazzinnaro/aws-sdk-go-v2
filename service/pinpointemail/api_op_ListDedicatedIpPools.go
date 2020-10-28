// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointemail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// List all of the dedicated IP pools that exist in your Amazon Pinpoint account in
// the current AWS Region.
func (c *Client) ListDedicatedIpPools(ctx context.Context, params *ListDedicatedIpPoolsInput, optFns ...func(*Options)) (*ListDedicatedIpPoolsOutput, error) {
	if params == nil {
		params = &ListDedicatedIpPoolsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDedicatedIpPools", params, optFns, addOperationListDedicatedIpPoolsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDedicatedIpPoolsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to obtain a list of dedicated IP pools.
type ListDedicatedIpPoolsInput struct {

	// A token returned from a previous call to ListDedicatedIpPools to indicate the
	// position in the list of dedicated IP pools.
	NextToken *string

	// The number of results to show in a single call to ListDedicatedIpPools. If the
	// number of results is larger than the number you specified in this parameter,
	// then the response includes a NextToken element, which you can use to obtain
	// additional results.
	PageSize *int32
}

// A list of dedicated IP pools.
type ListDedicatedIpPoolsOutput struct {

	// A list of all of the dedicated IP pools that are associated with your Amazon
	// Pinpoint account.
	DedicatedIpPools []*string

	// A token that indicates that there are additional IP pools to list. To view
	// additional IP pools, issue another request to ListDedicatedIpPools, passing this
	// token in the NextToken parameter.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListDedicatedIpPoolsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpListDedicatedIpPools{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpListDedicatedIpPools{})
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
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opListDedicatedIpPools(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opListDedicatedIpPools(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "ListDedicatedIpPools",
	}
}
