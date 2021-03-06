// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists tags for up to 10 health checks or hosted zones. For information about
// using tags for cost allocation, see Using Cost Allocation Tags
// (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html)
// in the AWS Billing and Cost Management User Guide.
func (c *Client) ListTagsForResources(ctx context.Context, params *ListTagsForResourcesInput, optFns ...func(*Options)) (*ListTagsForResourcesOutput, error) {
	if params == nil {
		params = &ListTagsForResourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTagsForResources", params, optFns, addOperationListTagsForResourcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTagsForResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A complex type that contains information about the health checks or hosted zones
// for which you want to list tags.
type ListTagsForResourcesInput struct {

	// A complex type that contains the ResourceId element for each resource for which
	// you want to get a list of tags.
	//
	// This member is required.
	ResourceIds []*string

	// The type of the resources.
	//
	// * The resource type for health checks is
	// healthcheck.
	//
	// * The resource type for hosted zones is hostedzone.
	//
	// This member is required.
	ResourceType types.TagResourceType
}

// A complex type containing tags for the specified resources.
type ListTagsForResourcesOutput struct {

	// A list of ResourceTagSets containing tags associated with the specified
	// resources.
	//
	// This member is required.
	ResourceTagSets []*types.ResourceTagSet

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListTagsForResourcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpListTagsForResources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListTagsForResources{}, middleware.After)
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
	addOpListTagsForResourcesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListTagsForResources(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListTagsForResources(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "ListTagsForResources",
	}
}
