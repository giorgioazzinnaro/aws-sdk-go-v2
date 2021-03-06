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

// Returns a list of cache subnet group descriptions. If a subnet group name is
// specified, the list contains only the description of that group. This is
// applicable only when you have ElastiCache in VPC setup. All ElastiCache clusters
// now launch in VPC by default.
func (c *Client) DescribeCacheSubnetGroups(ctx context.Context, params *DescribeCacheSubnetGroupsInput, optFns ...func(*Options)) (*DescribeCacheSubnetGroupsOutput, error) {
	if params == nil {
		params = &DescribeCacheSubnetGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeCacheSubnetGroups", params, optFns, addOperationDescribeCacheSubnetGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeCacheSubnetGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a DescribeCacheSubnetGroups operation.
type DescribeCacheSubnetGroupsInput struct {

	// The name of the cache subnet group to return details for.
	CacheSubnetGroupName *string

	// An optional marker returned from a prior request. Use this marker for pagination
	// of results from this operation. If this parameter is specified, the response
	// includes only records beyond the marker, up to the value specified by
	// MaxRecords.
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a marker is included in the response so
	// that the remaining results can be retrieved. Default: 100 Constraints: minimum
	// 20; maximum 100.
	MaxRecords *int32
}

// Represents the output of a DescribeCacheSubnetGroups operation.
type DescribeCacheSubnetGroupsOutput struct {

	// A list of cache subnet groups. Each element in the list contains detailed
	// information about one group.
	CacheSubnetGroups []*types.CacheSubnetGroup

	// Provides an identifier to allow retrieval of paginated results.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeCacheSubnetGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeCacheSubnetGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeCacheSubnetGroups{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeCacheSubnetGroups(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeCacheSubnetGroups(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "DescribeCacheSubnetGroups",
	}
}
