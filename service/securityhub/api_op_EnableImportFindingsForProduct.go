// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Enables the integration of a partner product with Security Hub. Integrated
// products send findings to Security Hub. When you enable a product integration, a
// permissions policy that grants permission for the product to send findings to
// Security Hub is applied.
func (c *Client) EnableImportFindingsForProduct(ctx context.Context, params *EnableImportFindingsForProductInput, optFns ...func(*Options)) (*EnableImportFindingsForProductOutput, error) {
	if params == nil {
		params = &EnableImportFindingsForProductInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "EnableImportFindingsForProduct", params, optFns, addOperationEnableImportFindingsForProductMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*EnableImportFindingsForProductOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EnableImportFindingsForProductInput struct {

	// The ARN of the product to enable the integration for.
	//
	// This member is required.
	ProductArn *string
}

type EnableImportFindingsForProductOutput struct {

	// The ARN of your subscription to the product to enable integrations for.
	ProductSubscriptionArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationEnableImportFindingsForProductMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpEnableImportFindingsForProduct{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpEnableImportFindingsForProduct{})
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
	if err := addOpEnableImportFindingsForProductValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opEnableImportFindingsForProduct(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opEnableImportFindingsForProduct(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securityhub",
		OperationName: "EnableImportFindingsForProduct",
	}
}
