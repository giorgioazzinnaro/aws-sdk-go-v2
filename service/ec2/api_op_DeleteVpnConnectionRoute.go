// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the specified static route associated with a VPN connection between an
// existing virtual private gateway and a VPN customer gateway. The static route
// allows traffic to be routed from the virtual private gateway to the VPN customer
// gateway.
func (c *Client) DeleteVpnConnectionRoute(ctx context.Context, params *DeleteVpnConnectionRouteInput, optFns ...func(*Options)) (*DeleteVpnConnectionRouteOutput, error) {
	if params == nil {
		params = &DeleteVpnConnectionRouteInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteVpnConnectionRoute", params, optFns, addOperationDeleteVpnConnectionRouteMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteVpnConnectionRouteOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DeleteVpnConnectionRoute.
type DeleteVpnConnectionRouteInput struct {

	// The CIDR block associated with the local subnet of the customer network.
	//
	// This member is required.
	DestinationCidrBlock *string

	// The ID of the VPN connection.
	//
	// This member is required.
	VpnConnectionId *string
}

type DeleteVpnConnectionRouteOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteVpnConnectionRouteMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsEc2query_serializeOpDeleteVpnConnectionRoute{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsEc2query_deserializeOpDeleteVpnConnectionRoute{})
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
	if err := addOpDeleteVpnConnectionRouteValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDeleteVpnConnectionRoute(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDeleteVpnConnectionRoute(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DeleteVpnConnectionRoute",
	}
}
