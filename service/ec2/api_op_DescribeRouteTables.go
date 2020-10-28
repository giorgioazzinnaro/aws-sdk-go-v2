// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes one or more of your route tables. Each subnet in your VPC must be
// associated with a route table. If a subnet is not explicitly associated with any
// route table, it is implicitly associated with the main route table. This command
// does not return the subnet ID for implicit associations. For more information,
// see Route Tables
// (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Route_Tables.html) in the
// Amazon Virtual Private Cloud User Guide.
func (c *Client) DescribeRouteTables(ctx context.Context, params *DescribeRouteTablesInput, optFns ...func(*Options)) (*DescribeRouteTablesOutput, error) {
	if params == nil {
		params = &DescribeRouteTablesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeRouteTables", params, optFns, addOperationDescribeRouteTablesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeRouteTablesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeRouteTablesInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// One or more filters.
	//
	//     * association.route-table-association-id - The ID of
	// an association ID for the route table.
	//
	//     * association.route-table-id - The
	// ID of the route table involved in the association.
	//
	//     * association.subnet-id
	// - The ID of the subnet involved in the association.
	//
	//     * association.main -
	// Indicates whether the route table is the main route table for the VPC (true |
	// false). Route tables that do not have an association ID are not returned in the
	// response.
	//
	//     * owner-id - The ID of the AWS account that owns the route
	// table.
	//
	//     * route-table-id - The ID of the route table.
	//
	//     *
	// route.destination-cidr-block - The IPv4 CIDR range specified in a route in the
	// table.
	//
	//     * route.destination-ipv6-cidr-block - The IPv6 CIDR range specified
	// in a route in the route table.
	//
	//     * route.destination-prefix-list-id - The ID
	// (prefix) of the AWS service specified in a route in the table.
	//
	//     *
	// route.egress-only-internet-gateway-id - The ID of an egress-only Internet
	// gateway specified in a route in the route table.
	//
	//     * route.gateway-id - The
	// ID of a gateway specified in a route in the table.
	//
	//     * route.instance-id -
	// The ID of an instance specified in a route in the table.
	//
	//     *
	// route.nat-gateway-id - The ID of a NAT gateway.
	//
	//     * route.transit-gateway-id
	// - The ID of a transit gateway.
	//
	//     * route.origin - Describes how the route was
	// created. CreateRouteTable indicates that the route was automatically created
	// when the route table was created; CreateRoute indicates that the route was
	// manually added to the route table; EnableVgwRoutePropagation indicates that the
	// route was propagated by route propagation.
	//
	//     * route.state - The state of a
	// route in the route table (active | blackhole). The blackhole state indicates
	// that the route's target isn't available (for example, the specified gateway
	// isn't attached to the VPC, the specified NAT instance has been terminated, and
	// so on).
	//
	//     * route.vpc-peering-connection-id - The ID of a VPC peering
	// connection specified in a route in the table.
	//
	//     * tag: - The key/value
	// combination of a tag assigned to the resource. Use the tag key in the filter
	// name and the tag value as the filter value. For example, to find all resources
	// that have a tag with the key Owner and the value TeamA, specify tag:Owner for
	// the filter name and TeamA for the filter value.
	//
	//     * tag-key - The key of a
	// tag assigned to the resource. Use this filter to find all resources assigned a
	// tag with a specific key, regardless of the tag value.
	//
	//     * vpc-id - The ID of
	// the VPC for the route table.
	Filters []*types.Filter

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	// One or more route table IDs. Default: Describes all your route tables.
	RouteTableIds []*string
}

// Contains the output of DescribeRouteTables.
type DescribeRouteTablesOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Information about one or more route tables.
	RouteTables []*types.RouteTable

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeRouteTablesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsEc2query_serializeOpDescribeRouteTables{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsEc2query_deserializeOpDescribeRouteTables{})
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
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDescribeRouteTables(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDescribeRouteTables(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeRouteTables",
	}
}
