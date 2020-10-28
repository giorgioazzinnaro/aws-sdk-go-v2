// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes an access policy that grants the specified identity access to the
// specified AWS IoT SiteWise Monitor resource. You can use this operation to
// revoke access to an AWS IoT SiteWise Monitor resource.
func (c *Client) DeleteAccessPolicy(ctx context.Context, params *DeleteAccessPolicyInput, optFns ...func(*Options)) (*DeleteAccessPolicyOutput, error) {
	if params == nil {
		params = &DeleteAccessPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteAccessPolicy", params, optFns, addOperationDeleteAccessPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteAccessPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteAccessPolicyInput struct {

	// The ID of the access policy to be deleted.
	//
	// This member is required.
	AccessPolicyId *string

	// A unique case-sensitive identifier that you can provide to ensure the
	// idempotency of the request. Don't reuse this client token if a new idempotent
	// request is required.
	ClientToken *string
}

type DeleteAccessPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteAccessPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpDeleteAccessPolicy{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpDeleteAccessPolicy{})
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
	if err := addEndpointPrefix_opDeleteAccessPolicyMiddleware(stack); err != nil {
		return err
	}
	if err := addIdempotencyToken_opDeleteAccessPolicyMiddleware(stack, options); err != nil {
		return err
	}
	if err := addOpDeleteAccessPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDeleteAccessPolicy(options.Region)); err != nil {
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

type endpointPrefix_opDeleteAccessPolicyMiddleware struct {
}

func (*endpointPrefix_opDeleteAccessPolicyMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opDeleteAccessPolicyMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.HostPrefix = "monitor."

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opDeleteAccessPolicyMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert("OperationSerializer", middleware.Before, &endpointPrefix_opDeleteAccessPolicyMiddleware{})
}

type idempotencyToken_initializeOpDeleteAccessPolicy struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpDeleteAccessPolicy) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpDeleteAccessPolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*DeleteAccessPolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *DeleteAccessPolicyInput ")
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
func addIdempotencyToken_opDeleteAccessPolicyMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(middleware.Before, &idempotencyToken_initializeOpDeleteAccessPolicy{tokenProvider: cfg.IdempotencyTokenProvider})
}

func newServiceMetadataMiddleware_opDeleteAccessPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotsitewise",
		OperationName: "DeleteAccessPolicy",
	}
}
