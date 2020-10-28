// Code generated by smithy-go-codegen DO NOT EDIT.

package directconnect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes one or more association proposals for connection between a virtual
// private gateway or transit gateway and a Direct Connect gateway.
func (c *Client) DescribeDirectConnectGatewayAssociationProposals(ctx context.Context, params *DescribeDirectConnectGatewayAssociationProposalsInput, optFns ...func(*Options)) (*DescribeDirectConnectGatewayAssociationProposalsOutput, error) {
	if params == nil {
		params = &DescribeDirectConnectGatewayAssociationProposalsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDirectConnectGatewayAssociationProposals", params, optFns, addOperationDescribeDirectConnectGatewayAssociationProposalsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDirectConnectGatewayAssociationProposalsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDirectConnectGatewayAssociationProposalsInput struct {

	// The ID of the associated gateway.
	AssociatedGatewayId *string

	// The ID of the Direct Connect gateway.
	DirectConnectGatewayId *string

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value. If
	// MaxResults is given a value larger than 100, only 100 results are returned.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	// The ID of the proposal.
	ProposalId *string
}

type DescribeDirectConnectGatewayAssociationProposalsOutput struct {

	// Describes the Direct Connect gateway association proposals.
	DirectConnectGatewayAssociationProposals []*types.DirectConnectGatewayAssociationProposal

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeDirectConnectGatewayAssociationProposalsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpDescribeDirectConnectGatewayAssociationProposals{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpDescribeDirectConnectGatewayAssociationProposals{})
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
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDescribeDirectConnectGatewayAssociationProposals(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDescribeDirectConnectGatewayAssociationProposals(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "directconnect",
		OperationName: "DescribeDirectConnectGatewayAssociationProposals",
	}
}
