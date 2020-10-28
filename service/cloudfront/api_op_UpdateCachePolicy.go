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

// Updates a cache policy configuration. When you update a cache policy
// configuration, all the fields are updated with the values provided in the
// request. You cannot update some fields independent of others. To update a cache
// policy configuration:
//
//     * Use GetCachePolicyConfig to get the current
// configuration.
//
//     * Locally modify the fields in the cache policy
// configuration that you want to update.
//
//     * Call UpdateCachePolicy by
// providing the entire cache policy configuration, including the fields that you
// modified and those that you didn’t.
func (c *Client) UpdateCachePolicy(ctx context.Context, params *UpdateCachePolicyInput, optFns ...func(*Options)) (*UpdateCachePolicyOutput, error) {
	if params == nil {
		params = &UpdateCachePolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateCachePolicy", params, optFns, addOperationUpdateCachePolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateCachePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateCachePolicyInput struct {

	// A cache policy configuration.
	//
	// This member is required.
	CachePolicyConfig *types.CachePolicyConfig

	// The unique identifier for the cache policy that you are updating. The identifier
	// is returned in a cache behavior’s CachePolicyId field in the response to
	// GetDistributionConfig.
	//
	// This member is required.
	Id *string

	// The version of the cache policy that you are updating. The version is returned
	// in the cache policy’s ETag field in the response to GetCachePolicyConfig.
	IfMatch *string
}

type UpdateCachePolicyOutput struct {

	// A cache policy.
	CachePolicy *types.CachePolicy

	// The current version of the cache policy.
	ETag *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateCachePolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestxml_serializeOpUpdateCachePolicy{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestxml_deserializeOpUpdateCachePolicy{})
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
	if err := addOpUpdateCachePolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opUpdateCachePolicy(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opUpdateCachePolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "UpdateCachePolicy",
	}
}
