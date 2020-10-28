// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Changes information about an ApiKey resource.
func (c *Client) UpdateApiKey(ctx context.Context, params *UpdateApiKeyInput, optFns ...func(*Options)) (*UpdateApiKeyOutput, error) {
	if params == nil {
		params = &UpdateApiKeyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateApiKey", params, optFns, addOperationUpdateApiKeyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateApiKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to change information about an ApiKey resource.
type UpdateApiKeyInput struct {

	// [Required] The identifier of the ApiKey resource to be updated.
	//
	// This member is required.
	ApiKey *string

	Name *string

	// A list of update operations to be applied to the specified resource and in the
	// order specified in this list.
	PatchOperations []*types.PatchOperation

	Template *bool

	TemplateSkipList []*string

	Title *string
}

// A resource that can be distributed to callers for executing Method resources
// that require an API key. API keys can be mapped to any Stage on any RestApi,
// which indicates that the callers with the API key can make requests to that
// stage. Use API Keys
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-api-keys.html)
type UpdateApiKeyOutput struct {

	// The timestamp when the API Key was created.
	CreatedDate *time.Time

	// An AWS Marketplace customer identifier , when integrating with the AWS SaaS
	// Marketplace.
	CustomerId *string

	// The description of the API Key.
	Description *string

	// Specifies whether the API Key can be used by callers.
	Enabled *bool

	// The identifier of the API Key.
	Id *string

	// The timestamp when the API Key was last updated.
	LastUpdatedDate *time.Time

	// The name of the API Key.
	Name *string

	// A list of Stage resources that are associated with the ApiKey resource.
	StageKeys []*string

	// The collection of tags. Each tag element is associated with a given resource.
	Tags map[string]*string

	// The value of the API Key.
	Value *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateApiKeyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpUpdateApiKey{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpUpdateApiKey{})
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
	if err := addOpUpdateApiKeyValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opUpdateApiKey(options.Region)); err != nil {
		return err
	}
	if err := addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err := addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err := addAcceptHeader(stack); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opUpdateApiKey(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "UpdateApiKey",
	}
}
