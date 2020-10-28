// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrass

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates a logger definition.
func (c *Client) UpdateLoggerDefinition(ctx context.Context, params *UpdateLoggerDefinitionInput, optFns ...func(*Options)) (*UpdateLoggerDefinitionOutput, error) {
	if params == nil {
		params = &UpdateLoggerDefinitionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateLoggerDefinition", params, optFns, addOperationUpdateLoggerDefinitionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateLoggerDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateLoggerDefinitionInput struct {

	// The ID of the logger definition.
	//
	// This member is required.
	LoggerDefinitionId *string

	// The name of the definition.
	Name *string
}

type UpdateLoggerDefinitionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateLoggerDefinitionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpUpdateLoggerDefinition{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpUpdateLoggerDefinition{})
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
	if err := addOpUpdateLoggerDefinitionValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opUpdateLoggerDefinition(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opUpdateLoggerDefinition(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "UpdateLoggerDefinition",
	}
}
