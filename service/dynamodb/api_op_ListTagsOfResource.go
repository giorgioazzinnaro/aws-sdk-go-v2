// Code generated by smithy-go-codegen DO NOT EDIT.

package dynamodb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// List all tags on an Amazon DynamoDB resource. You can call ListTagsOfResource up
// to 10 times per second, per account. For an overview on tagging DynamoDB
// resources, see Tagging for DynamoDB
// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Tagging.html)
// in the Amazon DynamoDB Developer Guide.
func (c *Client) ListTagsOfResource(ctx context.Context, params *ListTagsOfResourceInput, optFns ...func(*Options)) (*ListTagsOfResourceOutput, error) {
	if params == nil {
		params = &ListTagsOfResourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTagsOfResource", params, optFns, addOperationListTagsOfResourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTagsOfResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTagsOfResourceInput struct {

	// The Amazon DynamoDB resource with tags to be listed. This value is an Amazon
	// Resource Name (ARN).
	//
	// This member is required.
	ResourceArn *string

	// An optional string that, if supplied, must be copied from the output of a
	// previous call to ListTagOfResource. When provided in this manner, this API
	// fetches the next page of results.
	NextToken *string
}

type ListTagsOfResourceOutput struct {

	// If this value is returned, there are additional results to be displayed. To
	// retrieve them, call ListTagsOfResource again, with NextToken set to this value.
	NextToken *string

	// The tags currently associated with the Amazon DynamoDB resource.
	Tags []*types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListTagsOfResourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson10_serializeOpListTagsOfResource{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson10_deserializeOpListTagsOfResource{})
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
	if err := addOpListTagsOfResourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opListTagsOfResource(options.Region)); err != nil {
		return err
	}
	if err := addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err := addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err := addValidateResponseChecksum(stack, options); err != nil {
		return err
	}
	if err := addAcceptEncodingGzip(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opListTagsOfResource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dynamodb",
		OperationName: "ListTagsOfResource",
	}
}
