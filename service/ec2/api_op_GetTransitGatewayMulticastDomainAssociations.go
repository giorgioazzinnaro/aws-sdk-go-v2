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

// Gets information about the associations for the transit gateway multicast
// domain.
func (c *Client) GetTransitGatewayMulticastDomainAssociations(ctx context.Context, params *GetTransitGatewayMulticastDomainAssociationsInput, optFns ...func(*Options)) (*GetTransitGatewayMulticastDomainAssociationsOutput, error) {
	if params == nil {
		params = &GetTransitGatewayMulticastDomainAssociationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetTransitGatewayMulticastDomainAssociations", params, optFns, addOperationGetTransitGatewayMulticastDomainAssociationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetTransitGatewayMulticastDomainAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTransitGatewayMulticastDomainAssociationsInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// One or more filters. The possible values are:
	//
	//     * resource-id - The ID of the
	// resource.
	//
	//     * resource-type - The type of resource. The valid value is:
	// vpc.
	//
	//     * state - The state of the subnet association. Valid values are
	// associated | associating | disassociated | disassociating.
	//
	//     * subnet-id -
	// The ID of the subnet.
	//
	//     * transit-gateway-attachment-id - The id of the
	// transit gateway attachment.
	Filters []*types.Filter

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	// The ID of the transit gateway multicast domain.
	TransitGatewayMulticastDomainId *string
}

type GetTransitGatewayMulticastDomainAssociationsOutput struct {

	// Information about the multicast domain associations.
	MulticastDomainAssociations []*types.TransitGatewayMulticastDomainAssociation

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetTransitGatewayMulticastDomainAssociationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsEc2query_serializeOpGetTransitGatewayMulticastDomainAssociations{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsEc2query_deserializeOpGetTransitGatewayMulticastDomainAssociations{})
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
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opGetTransitGatewayMulticastDomainAssociations(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opGetTransitGatewayMulticastDomainAssociations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "GetTransitGatewayMulticastDomainAssociations",
	}
}
