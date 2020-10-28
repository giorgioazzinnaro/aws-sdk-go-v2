// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds up to 50 cost allocation tags to the named resource. A cost allocation tag
// is a key-value pair where the key and value are case-sensitive. You can use cost
// allocation tags to categorize and track your AWS costs. When you apply tags to
// your ElastiCache resources, AWS generates a cost allocation report as a
// comma-separated value (CSV) file with your usage and costs aggregated by your
// tags. You can apply tags that represent business categories (such as cost
// centers, application names, or owners) to organize your costs across multiple
// services. For more information, see Using Cost Allocation Tags in Amazon
// ElastiCache
// (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/Tagging.html) in
// the ElastiCache User Guide.
func (c *Client) AddTagsToResource(ctx context.Context, params *AddTagsToResourceInput, optFns ...func(*Options)) (*AddTagsToResourceOutput, error) {
	if params == nil {
		params = &AddTagsToResourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AddTagsToResource", params, optFns, addOperationAddTagsToResourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AddTagsToResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of an AddTagsToResource operation.
type AddTagsToResourceInput struct {

	// The Amazon Resource Name (ARN) of the resource to which the tags are to be
	// added, for example arn:aws:elasticache:us-west-2:0123456789:cluster:myCluster or
	// arn:aws:elasticache:us-west-2:0123456789:snapshot:mySnapshot. ElastiCache
	// resources are cluster and snapshot. For more information about ARNs, see Amazon
	// Resource Names (ARNs) and AWS Service Namespaces
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html).
	//
	// This member is required.
	ResourceName *string

	// A list of cost allocation tags to be added to this resource. A tag is a
	// key-value pair. A tag key must be accompanied by a tag value.
	//
	// This member is required.
	Tags []*types.Tag
}

// Represents the output from the AddTagsToResource, ListTagsForResource, and
// RemoveTagsFromResource operations.
type AddTagsToResourceOutput struct {

	// A list of cost allocation tags as key-value pairs.
	TagList []*types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAddTagsToResourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsquery_serializeOpAddTagsToResource{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsquery_deserializeOpAddTagsToResource{})
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
	if err := addOpAddTagsToResourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opAddTagsToResource(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opAddTagsToResource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "AddTagsToResource",
	}
}
