// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a NAT gateway in the specified public subnet. This action creates a
// network interface in the specified subnet with a private IP address from the IP
// address range of the subnet. Internet-bound traffic from a private subnet can be
// routed to the NAT gateway, therefore enabling instances in the private subnet to
// connect to the internet. For more information, see NAT Gateways
// (https://docs.aws.amazon.com/vpc/latest/userguide/vpc-nat-gateway.html) in the
// Amazon Virtual Private Cloud User Guide.
func (c *Client) CreateNatGateway(ctx context.Context, params *CreateNatGatewayInput, optFns ...func(*Options)) (*CreateNatGatewayOutput, error) {
	if params == nil {
		params = &CreateNatGatewayInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateNatGateway", params, optFns, addOperationCreateNatGatewayMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateNatGatewayOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateNatGatewayInput struct {

	// The allocation ID of an Elastic IP address to associate with the NAT gateway. If
	// the Elastic IP address is associated with another resource, you must first
	// disassociate it.
	//
	// This member is required.
	AllocationId *string

	// The subnet in which to create the NAT gateway.
	//
	// This member is required.
	SubnetId *string

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. For more information, see How to Ensure Idempotency
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html).
	// Constraint: Maximum 64 ASCII characters.
	ClientToken *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The tags to assign to the NAT gateway.
	TagSpecifications []*types.TagSpecification
}

type CreateNatGatewayOutput struct {

	// Unique, case-sensitive identifier to ensure the idempotency of the request. Only
	// returned if a client token was provided in the request.
	ClientToken *string

	// Information about the NAT gateway.
	NatGateway *types.NatGateway

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateNatGatewayMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsEc2query_serializeOpCreateNatGateway{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsEc2query_deserializeOpCreateNatGateway{})
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
	if err := addIdempotencyToken_opCreateNatGatewayMiddleware(stack, options); err != nil {
		return err
	}
	if err := addOpCreateNatGatewayValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opCreateNatGateway(options.Region)); err != nil {
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

type idempotencyToken_initializeOpCreateNatGateway struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateNatGateway) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateNatGateway) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateNatGatewayInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateNatGatewayInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateNatGatewayMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(middleware.Before, &idempotencyToken_initializeOpCreateNatGateway{tokenProvider: cfg.IdempotencyTokenProvider})
}

func newServiceMetadataMiddleware_opCreateNatGateway(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateNatGateway",
	}
}
