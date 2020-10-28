// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesis

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the tags for the specified Kinesis data stream. This operation has a limit
// of five transactions per second per account.
func (c *Client) ListTagsForStream(ctx context.Context, params *ListTagsForStreamInput, optFns ...func(*Options)) (*ListTagsForStreamOutput, error) {
	if params == nil {
		params = &ListTagsForStreamInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTagsForStream", params, optFns, addOperationListTagsForStreamMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTagsForStreamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for ListTagsForStream.
type ListTagsForStreamInput struct {

	// The name of the stream.
	//
	// This member is required.
	StreamName *string

	// The key to use as the starting point for the list of tags. If this parameter is
	// set, ListTagsForStream gets all tags that occur after ExclusiveStartTagKey.
	ExclusiveStartTagKey *string

	// The number of tags to return. If this number is less than the total number of
	// tags associated with the stream, HasMoreTags is set to true. To list additional
	// tags, set ExclusiveStartTagKey to the last key in the response.
	Limit *int32
}

// Represents the output for ListTagsForStream.
type ListTagsForStreamOutput struct {

	// If set to true, more tags are available. To request additional tags, set
	// ExclusiveStartTagKey to the key of the last tag returned.
	//
	// This member is required.
	HasMoreTags *bool

	// A list of tags associated with StreamName, starting with the first tag after
	// ExclusiveStartTagKey and up to the specified Limit.
	//
	// This member is required.
	Tags []*types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListTagsForStreamMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpListTagsForStream{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpListTagsForStream{})
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
	if err := addOpListTagsForStreamValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opListTagsForStream(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opListTagsForStream(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesis",
		OperationName: "ListTagsForStream",
	}
}
