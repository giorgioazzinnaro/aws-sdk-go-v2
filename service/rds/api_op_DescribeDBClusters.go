// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about provisioned Aurora DB clusters. This API supports
// pagination. For more information on Amazon Aurora, see  What Is Amazon Aurora?
// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html)
// in the Amazon Aurora User Guide. This operation can also return information for
// Amazon Neptune DB instances and Amazon DocumentDB instances.
func (c *Client) DescribeDBClusters(ctx context.Context, params *DescribeDBClustersInput, optFns ...func(*Options)) (*DescribeDBClustersOutput, error) {
	if params == nil {
		params = &DescribeDBClustersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDBClusters", params, optFns, addOperationDescribeDBClustersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDBClustersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DescribeDBClustersInput struct {

	// The user-supplied DB cluster identifier. If this parameter is specified,
	// information from only the specific DB cluster is returned. This parameter isn't
	// case-sensitive. Constraints:
	//
	//     * If supplied, must match an existing
	// DBClusterIdentifier.
	DBClusterIdentifier *string

	// A filter that specifies one or more DB clusters to describe. Supported
	// filters:
	//
	//     * db-cluster-id - Accepts DB cluster identifiers and DB cluster
	// Amazon Resource Names (ARNs). The results list will only include information
	// about the DB clusters identified by these ARNs.
	Filters []*types.Filter

	// Optional Boolean parameter that specifies whether the output includes
	// information about clusters shared from other AWS accounts.
	IncludeShared *bool

	// An optional pagination token provided by a previous DescribeDBClusters request.
	// If this parameter is specified, the response includes only records beyond the
	// marker, up to the value specified by MaxRecords.
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so you can retrieve the remaining results. Default: 100
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int32
}

// Contains the result of a successful invocation of the DescribeDBClusters action.
type DescribeDBClustersOutput struct {

	// Contains a list of DB clusters for the user.
	DBClusters []*types.DBCluster

	// A pagination token that can be used in a later DescribeDBClusters request.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeDBClustersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsquery_serializeOpDescribeDBClusters{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsquery_deserializeOpDescribeDBClusters{})
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
	if err := addOpDescribeDBClustersValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDescribeDBClusters(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDescribeDBClusters(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeDBClusters",
	}
}
