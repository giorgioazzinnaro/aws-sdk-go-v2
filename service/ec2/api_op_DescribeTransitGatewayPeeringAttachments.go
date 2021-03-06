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

// Describes your transit gateway peering attachments.
func (c *Client) DescribeTransitGatewayPeeringAttachments(ctx context.Context, params *DescribeTransitGatewayPeeringAttachmentsInput, optFns ...func(*Options)) (*DescribeTransitGatewayPeeringAttachmentsOutput, error) {
	if params == nil {
		params = &DescribeTransitGatewayPeeringAttachmentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTransitGatewayPeeringAttachments", params, optFns, addOperationDescribeTransitGatewayPeeringAttachmentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTransitGatewayPeeringAttachmentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTransitGatewayPeeringAttachmentsInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// One or more filters. The possible values are:
	//
	// * transit-gateway-attachment-id -
	// The ID of the transit gateway attachment.
	//
	// * local-owner-id - The ID of your AWS
	// account.
	//
	// * remote-owner-id - The ID of the AWS account in the remote Region
	// that owns the transit gateway.
	//
	// * state - The state of the peering attachment.
	// Valid values are available | deleted | deleting | failed | failing |
	// initiatingRequest | modifying | pendingAcceptance | pending | rollingBack |
	// rejected | rejecting).
	//
	// * transit-gateway-id - The ID of the transit gateway.
	Filters []*types.Filter

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	// One or more IDs of the transit gateway peering attachments.
	TransitGatewayAttachmentIds []*string
}

type DescribeTransitGatewayPeeringAttachmentsOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// The transit gateway peering attachments.
	TransitGatewayPeeringAttachments []*types.TransitGatewayPeeringAttachment

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeTransitGatewayPeeringAttachmentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeTransitGatewayPeeringAttachments{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeTransitGatewayPeeringAttachments{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTransitGatewayPeeringAttachments(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeTransitGatewayPeeringAttachments(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeTransitGatewayPeeringAttachments",
	}
}
