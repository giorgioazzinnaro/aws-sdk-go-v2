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

// Returns a list of DBClusterParameterGroup descriptions. If a
// DBClusterParameterGroupName parameter is specified, the list will contain only
// the description of the specified DB cluster parameter group. For more
// information on Amazon Aurora, see  What Is Amazon Aurora?
// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html)
// in the Amazon Aurora User Guide. This action only applies to Aurora DB clusters.
func (c *Client) DescribeDBClusterParameterGroups(ctx context.Context, params *DescribeDBClusterParameterGroupsInput, optFns ...func(*Options)) (*DescribeDBClusterParameterGroupsOutput, error) {
	if params == nil {
		params = &DescribeDBClusterParameterGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDBClusterParameterGroups", params, optFns, addOperationDescribeDBClusterParameterGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDBClusterParameterGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DescribeDBClusterParameterGroupsInput struct {

	// The name of a specific DB cluster parameter group to return details for.
	// Constraints:
	//
	// * If supplied, must match the name of an existing
	// DBClusterParameterGroup.
	DBClusterParameterGroupName *string

	// This parameter isn't currently supported.
	Filters []*types.Filter

	// An optional pagination token provided by a previous
	// DescribeDBClusterParameterGroups request. If this parameter is specified, the
	// response includes only records beyond the marker, up to the value specified by
	// MaxRecords.
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so you can retrieve the remaining results. Default: 100
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int32
}

//
type DescribeDBClusterParameterGroupsOutput struct {

	// A list of DB cluster parameter groups.
	DBClusterParameterGroups []*types.DBClusterParameterGroup

	// An optional pagination token provided by a previous
	// DescribeDBClusterParameterGroups request. If this parameter is specified, the
	// response includes only records beyond the marker, up to the value specified by
	// MaxRecords.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeDBClusterParameterGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeDBClusterParameterGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeDBClusterParameterGroups{}, middleware.After)
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
	addOpDescribeDBClusterParameterGroupsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDBClusterParameterGroups(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeDBClusterParameterGroups(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeDBClusterParameterGroups",
	}
}
