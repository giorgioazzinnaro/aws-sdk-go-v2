// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a given schema. Schemas in a development and published state can only be
// deleted.
func (c *Client) DeleteSchema(ctx context.Context, params *DeleteSchemaInput, optFns ...func(*Options)) (*DeleteSchemaOutput, error) {
	if params == nil {
		params = &DeleteSchemaInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteSchema", params, optFns, addOperationDeleteSchemaMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteSchemaOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteSchemaInput struct {

	// The Amazon Resource Name (ARN) of the development schema. For more information,
	// see arns.
	//
	// This member is required.
	SchemaArn *string
}

type DeleteSchemaOutput struct {

	// The input ARN that is returned as part of the response. For more information,
	// see arns.
	SchemaArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteSchemaMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpDeleteSchema{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpDeleteSchema{})
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
	if err := addOpDeleteSchemaValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDeleteSchema(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDeleteSchema(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "DeleteSchema",
	}
}
