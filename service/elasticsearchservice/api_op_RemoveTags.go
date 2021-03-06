// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticsearchservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes the specified set of tags from the specified Elasticsearch domain.
func (c *Client) RemoveTags(ctx context.Context, params *RemoveTagsInput, optFns ...func(*Options)) (*RemoveTagsOutput, error) {
	if params == nil {
		params = &RemoveTagsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RemoveTags", params, optFns, addOperationRemoveTagsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RemoveTagsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the RemoveTags operation. Specify the ARN for
// the Elasticsearch domain from which you want to remove the specified TagKey.
type RemoveTagsInput struct {

	// Specifies the ARN for the Elasticsearch domain from which you want to delete the
	// specified tags.
	//
	// This member is required.
	ARN *string

	// Specifies the TagKey list which you want to remove from the Elasticsearch
	// domain.
	//
	// This member is required.
	TagKeys []*string
}

type RemoveTagsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRemoveTagsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpRemoveTags{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpRemoveTags{}, middleware.After)
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
	addOpRemoveTagsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRemoveTags(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opRemoveTags(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "RemoveTags",
	}
}
