// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves all tags that are assigned to a GameLift resource. Resource tags are
// used to organize AWS resources for a range of purposes. This operation handles
// the permissions necessary to manage tags for the following GameLift resource
// types:
//
//     * Build
//
//     * Script
//
//     * Fleet
//
//     * Alias
//
//     *
// GameSessionQueue
//
//     * MatchmakingConfiguration
//
//     * MatchmakingRuleSet
//
// To
// list tags for a resource, specify the unique ARN value for the resource. Learn
// more Tagging AWS Resources
// (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the AWS
// General Reference  AWS Tagging Strategies
// (http://aws.amazon.com/answers/account-management/aws-tagging-strategies/)
// Related operations
//
//     * TagResource
//
//     * UntagResource
//
//     *
// ListTagsForResource
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

type ListTagsForResourceInput struct {

	// The Amazon Resource Name (ARN
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-arn-format.html)) that is
	// assigned to and uniquely identifies the GameLift resource that you want to
	// retrieve tags for. GameLift resource ARNs are included in the data object for
	// the resource, which can be retrieved by calling a List or Describe operation for
	// the resource type.
	//
	// This member is required.
	ResourceARN *string
}

type ListTagsForResourceOutput struct {

	// The collection of tags that have been assigned to the specified resource.
	Tags []*types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListTagsForResourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpListTagsForResource{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpListTagsForResource{})
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
		SigningName:   "gamelift",
		OperationName: "ListTagsForResource",
	}
}
