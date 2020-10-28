// Code generated by smithy-go-codegen DO NOT EDIT.

package schemas

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

func (c *Client) ExportSchema(ctx context.Context, params *ExportSchemaInput, optFns ...func(*Options)) (*ExportSchemaOutput, error) {
	if params == nil {
		params = &ExportSchemaInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ExportSchema", params, optFns, addOperationExportSchemaMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ExportSchemaOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ExportSchemaInput struct {

	// The name of the registry.
	//
	// This member is required.
	RegistryName *string

	// The name of the schema.
	//
	// This member is required.
	SchemaName *string

	Type *string

	// Specifying this limits the results to only this schema version.
	SchemaVersion *string
}

type ExportSchemaOutput struct {
	Content *string

	SchemaArn *string

	SchemaName *string

	SchemaVersion *string

	Type *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationExportSchemaMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpExportSchema{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpExportSchema{})
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
	if err := addOpExportSchemaValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opExportSchema(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opExportSchema(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "schemas",
		OperationName: "ExportSchema",
	}
}
