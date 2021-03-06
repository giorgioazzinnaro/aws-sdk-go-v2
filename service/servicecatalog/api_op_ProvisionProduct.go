// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Provisions the specified product. A provisioned product is a resourced instance
// of a product. For example, provisioning a product based on a CloudFormation
// template launches a CloudFormation stack and its underlying resources. You can
// check the status of this request using DescribeRecord. If the request contains a
// tag key with an empty list of values, there is a tag conflict for that key. Do
// not include conflicted keys as tags, or this causes the error "Parameter
// validation failed: Missing required parameter in Tags[N]:Value".
func (c *Client) ProvisionProduct(ctx context.Context, params *ProvisionProductInput, optFns ...func(*Options)) (*ProvisionProductOutput, error) {
	if params == nil {
		params = &ProvisionProductInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ProvisionProduct", params, optFns, addOperationProvisionProductMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ProvisionProductOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ProvisionProductInput struct {

	// An idempotency token that uniquely identifies the provisioning request.
	//
	// This member is required.
	ProvisionToken *string

	// A user-friendly name for the provisioned product. This value must be unique for
	// the AWS account and cannot be updated after the product is provisioned.
	//
	// This member is required.
	ProvisionedProductName *string

	// The language code.
	//
	// * en - English (default)
	//
	// * jp - Japanese
	//
	// * zh - Chinese
	AcceptLanguage *string

	// Passed to CloudFormation. The SNS topic ARNs to which to publish stack-related
	// events.
	NotificationArns []*string

	// The path identifier of the product. This value is optional if the product has a
	// default path, and required if the product has more than one path. To list the
	// paths for a product, use ListLaunchPaths. You must provide the name or ID, but
	// not both.
	PathId *string

	// The name of the path. You must provide the name or ID, but not both.
	PathName *string

	// The product identifier. You must provide the name or ID, but not both.
	ProductId *string

	// The name of the product. You must provide the name or ID, but not both.
	ProductName *string

	// The identifier of the provisioning artifact. You must provide the name or ID,
	// but not both.
	ProvisioningArtifactId *string

	// The name of the provisioning artifact. You must provide the name or ID, but not
	// both.
	ProvisioningArtifactName *string

	// Parameters specified by the administrator that are required for provisioning the
	// product.
	ProvisioningParameters []*types.ProvisioningParameter

	// An object that contains information about the provisioning preferences for a
	// stack set.
	ProvisioningPreferences *types.ProvisioningPreferences

	// One or more tags.
	Tags []*types.Tag
}

type ProvisionProductOutput struct {

	// Information about the result of provisioning the product.
	RecordDetail *types.RecordDetail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationProvisionProductMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpProvisionProduct{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpProvisionProduct{}, middleware.After)
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
	addIdempotencyToken_opProvisionProductMiddleware(stack, options)
	addOpProvisionProductValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opProvisionProduct(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

type idempotencyToken_initializeOpProvisionProduct struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpProvisionProduct) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpProvisionProduct) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*ProvisionProductInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *ProvisionProductInput ")
	}

	if input.ProvisionToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ProvisionToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opProvisionProductMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpProvisionProduct{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opProvisionProduct(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "ProvisionProduct",
	}
}
