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

// Associates the specified attachment with the specified transit gateway route
// table. You can associate only one route table with an attachment.
func (c *Client) AssociateTransitGatewayRouteTable(ctx context.Context, params *AssociateTransitGatewayRouteTableInput, optFns ...func(*Options)) (*AssociateTransitGatewayRouteTableOutput, error) {
	if params == nil {
		params = &AssociateTransitGatewayRouteTableInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateTransitGatewayRouteTable", params, optFns, addOperationAssociateTransitGatewayRouteTableMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateTransitGatewayRouteTableOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateTransitGatewayRouteTableInput struct {

	// The ID of the attachment.
	//
	// This member is required.
	TransitGatewayAttachmentId *string

	// The ID of the transit gateway route table.
	//
	// This member is required.
	TransitGatewayRouteTableId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type AssociateTransitGatewayRouteTableOutput struct {

	// The ID of the association.
	Association *types.TransitGatewayAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAssociateTransitGatewayRouteTableMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsEc2query_serializeOpAssociateTransitGatewayRouteTable{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsEc2query_deserializeOpAssociateTransitGatewayRouteTable{})
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
	if err := addOpAssociateTransitGatewayRouteTableValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opAssociateTransitGatewayRouteTable(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opAssociateTransitGatewayRouteTable(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "AssociateTransitGatewayRouteTable",
	}
}
