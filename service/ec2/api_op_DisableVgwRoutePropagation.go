// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disables a virtual private gateway (VGW) from propagating routes to a specified
// route table of a VPC.
func (c *Client) DisableVgwRoutePropagation(ctx context.Context, params *DisableVgwRoutePropagationInput, optFns ...func(*Options)) (*DisableVgwRoutePropagationOutput, error) {
	if params == nil {
		params = &DisableVgwRoutePropagationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisableVgwRoutePropagation", params, optFns, addOperationDisableVgwRoutePropagationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisableVgwRoutePropagationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DisableVgwRoutePropagation.
type DisableVgwRoutePropagationInput struct {

	// The ID of the virtual private gateway.
	//
	// This member is required.
	GatewayId *string

	// The ID of the route table.
	//
	// This member is required.
	RouteTableId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type DisableVgwRoutePropagationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDisableVgwRoutePropagationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsEc2query_serializeOpDisableVgwRoutePropagation{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsEc2query_deserializeOpDisableVgwRoutePropagation{})
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
	if err := addOpDisableVgwRoutePropagationValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDisableVgwRoutePropagation(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDisableVgwRoutePropagation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DisableVgwRoutePropagation",
	}
}
