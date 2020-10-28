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

// Create a new streaming distribution with tags.
func (c *Client) CreateStreamingDistributionWithTags(ctx context.Context, params *CreateStreamingDistributionWithTagsInput, optFns ...func(*Options)) (*CreateStreamingDistributionWithTagsOutput, error) {
	if params == nil {
		params = &CreateStreamingDistributionWithTagsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateStreamingDistributionWithTags", params, optFns, addOperationCreateStreamingDistributionWithTagsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateStreamingDistributionWithTagsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request to create a new streaming distribution with tags.
type CreateStreamingDistributionWithTagsInput struct {

	// The streaming distribution's configuration information.
	//
	// This member is required.
	StreamingDistributionConfigWithTags *types.StreamingDistributionConfigWithTags
}

// The returned result of the corresponding request.
type CreateStreamingDistributionWithTagsOutput struct {

	// The current version of the distribution created.
	ETag *string

	// The fully qualified URI of the new streaming distribution resource just created.
	Location *string

	// The streaming distribution's information.
	StreamingDistribution *types.StreamingDistribution

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateStreamingDistributionWithTagsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestxml_serializeOpCreateStreamingDistributionWithTags{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestxml_deserializeOpCreateStreamingDistributionWithTags{})
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
	if err := addOpCreateStreamingDistributionWithTagsValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opCreateStreamingDistributionWithTags(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opCreateStreamingDistributionWithTags(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "CreateStreamingDistributionWithTags",
	}
}
