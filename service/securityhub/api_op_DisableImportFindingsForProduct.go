// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disables the integration of the specified product with Security Hub. After the
// integration is disabled, findings from that product are no longer sent to
// Security Hub.
func (c *Client) DisableImportFindingsForProduct(ctx context.Context, params *DisableImportFindingsForProductInput, optFns ...func(*Options)) (*DisableImportFindingsForProductOutput, error) {
	if params == nil {
		params = &DisableImportFindingsForProductInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisableImportFindingsForProduct", params, optFns, addOperationDisableImportFindingsForProductMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisableImportFindingsForProductOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisableImportFindingsForProductInput struct {

	// The ARN of the integrated product to disable the integration for.
	//
	// This member is required.
	ProductSubscriptionArn *string
}

type DisableImportFindingsForProductOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDisableImportFindingsForProductMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpDisableImportFindingsForProduct{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpDisableImportFindingsForProduct{})
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
	if err := addOpDisableImportFindingsForProductValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDisableImportFindingsForProduct(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDisableImportFindingsForProduct(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securityhub",
		OperationName: "DisableImportFindingsForProduct",
	}
}
