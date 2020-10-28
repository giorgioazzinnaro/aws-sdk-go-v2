// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Publishes a development schema with a major version and a recommended minor
// version.
func (c *Client) PublishSchema(ctx context.Context, params *PublishSchemaInput, optFns ...func(*Options)) (*PublishSchemaOutput, error) {
	if params == nil {
		params = &PublishSchemaInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PublishSchema", params, optFns, addOperationPublishSchemaMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PublishSchemaOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PublishSchemaInput struct {

	// The Amazon Resource Name (ARN) that is associated with the development schema.
	// For more information, see arns.
	//
	// This member is required.
	DevelopmentSchemaArn *string

	// The major version under which the schema will be published. Schemas have both a
	// major and minor version associated with them.
	//
	// This member is required.
	Version *string

	// The minor version under which the schema will be published. This parameter is
	// recommended. Schemas have both a major and minor version associated with them.
	MinorVersion *string

	// The new name under which the schema will be published. If this is not provided,
	// the development schema is considered.
	Name *string
}

type PublishSchemaOutput struct {

	// The ARN that is associated with the published schema. For more information, see
	// arns.
	PublishedSchemaArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPublishSchemaMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpPublishSchema{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpPublishSchema{})
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
	if err := addOpPublishSchemaValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opPublishSchema(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opPublishSchema(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "PublishSchema",
	}
}
