// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrass

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a version of a Lambda function definition that has already been defined.
func (c *Client) CreateFunctionDefinitionVersion(ctx context.Context, params *CreateFunctionDefinitionVersionInput, optFns ...func(*Options)) (*CreateFunctionDefinitionVersionOutput, error) {
	if params == nil {
		params = &CreateFunctionDefinitionVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateFunctionDefinitionVersion", params, optFns, addOperationCreateFunctionDefinitionVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateFunctionDefinitionVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Information needed to create a function definition version.
type CreateFunctionDefinitionVersionInput struct {

	// The ID of the Lambda function definition.
	//
	// This member is required.
	FunctionDefinitionId *string

	// A client token used to correlate requests and responses.
	AmznClientToken *string

	// The default configuration that applies to all Lambda functions in this function
	// definition version. Individual Lambda functions can override these settings.
	DefaultConfig *types.FunctionDefaultConfig

	// A list of Lambda functions in this function definition version.
	Functions []*types.Function
}

type CreateFunctionDefinitionVersionOutput struct {

	// The ARN of the version.
	Arn *string

	// The time, in milliseconds since the epoch, when the version was created.
	CreationTimestamp *string

	// The ID of the parent definition that the version is associated with.
	Id *string

	// The ID of the version.
	Version *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateFunctionDefinitionVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpCreateFunctionDefinitionVersion{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpCreateFunctionDefinitionVersion{})
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
	if err := addOpCreateFunctionDefinitionVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opCreateFunctionDefinitionVersion(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opCreateFunctionDefinitionVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "CreateFunctionDefinitionVersion",
	}
}
