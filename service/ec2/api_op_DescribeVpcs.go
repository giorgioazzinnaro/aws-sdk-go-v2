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

// Describes one or more of your VPCs.
func (c *Client) DescribeVpcs(ctx context.Context, params *DescribeVpcsInput, optFns ...func(*Options)) (*DescribeVpcsOutput, error) {
	if params == nil {
		params = &DescribeVpcsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeVpcs", params, optFns, addOperationDescribeVpcsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeVpcsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeVpcsInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// One or more filters.
	//
	//     * cidr - The primary IPv4 CIDR block of the VPC. The
	// CIDR block you specify must exactly match the VPC's CIDR block for information
	// to be returned for the VPC. Must contain the slash followed by one or two digits
	// (for example, /28).
	//
	//     * cidr-block-association.cidr-block - An IPv4 CIDR
	// block associated with the VPC.
	//
	//     * cidr-block-association.association-id -
	// The association ID for an IPv4 CIDR block associated with the VPC.
	//
	//     *
	// cidr-block-association.state - The state of an IPv4 CIDR block associated with
	// the VPC.
	//
	//     * dhcp-options-id - The ID of a set of DHCP options.
	//
	//     *
	// ipv6-cidr-block-association.ipv6-cidr-block - An IPv6 CIDR block associated with
	// the VPC.
	//
	//     * ipv6-cidr-block-association.ipv6-pool - The ID of the IPv6
	// address pool from which the IPv6 CIDR block is allocated.
	//
	//     *
	// ipv6-cidr-block-association.association-id - The association ID for an IPv6 CIDR
	// block associated with the VPC.
	//
	//     * ipv6-cidr-block-association.state - The
	// state of an IPv6 CIDR block associated with the VPC.
	//
	//     * isDefault -
	// Indicates whether the VPC is the default VPC.
	//
	//     * owner-id - The ID of the
	// AWS account that owns the VPC.
	//
	//     * state - The state of the VPC (pending |
	// available).
	//
	//     * tag: - The key/value combination of a tag assigned to the
	// resource. Use the tag key in the filter name and the tag value as the filter
	// value. For example, to find all resources that have a tag with the key Owner and
	// the value TeamA, specify tag:Owner for the filter name and TeamA for the filter
	// value.
	//
	//     * tag-key - The key of a tag assigned to the resource. Use this
	// filter to find all resources assigned a tag with a specific key, regardless of
	// the tag value.
	//
	//     * vpc-id - The ID of the VPC.
	Filters []*types.Filter

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	// One or more VPC IDs. Default: Describes all your VPCs.
	VpcIds []*string
}

type DescribeVpcsOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Information about one or more VPCs.
	Vpcs []*types.Vpc

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeVpcsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsEc2query_serializeOpDescribeVpcs{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsEc2query_deserializeOpDescribeVpcs{})
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
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDescribeVpcs(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDescribeVpcs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeVpcs",
	}
}
