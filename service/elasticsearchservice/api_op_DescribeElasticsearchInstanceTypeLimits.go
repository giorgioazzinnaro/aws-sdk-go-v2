// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticsearchservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describe Elasticsearch Limits for a given InstanceType and ElasticsearchVersion.
// When modifying existing Domain, specify the DomainName to know what Limits are
// supported for modifying.
func (c *Client) DescribeElasticsearchInstanceTypeLimits(ctx context.Context, params *DescribeElasticsearchInstanceTypeLimitsInput, optFns ...func(*Options)) (*DescribeElasticsearchInstanceTypeLimitsOutput, error) {
	if params == nil {
		params = &DescribeElasticsearchInstanceTypeLimitsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeElasticsearchInstanceTypeLimits", params, optFns, addOperationDescribeElasticsearchInstanceTypeLimitsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeElasticsearchInstanceTypeLimitsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to DescribeElasticsearchInstanceTypeLimits
// operation.
type DescribeElasticsearchInstanceTypeLimitsInput struct {

	// Version of Elasticsearch for which Limits are needed.
	//
	// This member is required.
	ElasticsearchVersion *string

	// The instance type for an Elasticsearch cluster for which Elasticsearch Limits
	// are needed.
	//
	// This member is required.
	InstanceType types.ESPartitionInstanceType

	// DomainName represents the name of the Domain that we are trying to modify. This
	// should be present only if we are querying for Elasticsearch Limits for existing
	// domain.
	DomainName *string
}

// Container for the parameters received from
// DescribeElasticsearchInstanceTypeLimits operation.
type DescribeElasticsearchInstanceTypeLimitsOutput struct {

	// Map of Role of the Instance and Limits that are applicable. Role performed by
	// given Instance in Elasticsearch can be one of the following:
	//
	//     * data: If the
	// given InstanceType is used as data node
	//
	//     * master: If the given InstanceType
	// is used as master node
	//
	//     * ultra_warm: If the given InstanceType is used as
	// warm node
	LimitsByRole map[string]*types.Limits

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeElasticsearchInstanceTypeLimitsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpDescribeElasticsearchInstanceTypeLimits{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpDescribeElasticsearchInstanceTypeLimits{})
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
	if err := addOpDescribeElasticsearchInstanceTypeLimitsValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDescribeElasticsearchInstanceTypeLimits(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDescribeElasticsearchInstanceTypeLimits(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "DescribeElasticsearchInstanceTypeLimits",
	}
}
