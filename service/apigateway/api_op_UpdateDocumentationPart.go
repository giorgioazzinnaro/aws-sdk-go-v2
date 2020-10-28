// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

func (c *Client) UpdateDocumentationPart(ctx context.Context, params *UpdateDocumentationPartInput, optFns ...func(*Options)) (*UpdateDocumentationPartOutput, error) {
	if params == nil {
		params = &UpdateDocumentationPartInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDocumentationPart", params, optFns, addOperationUpdateDocumentationPartMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDocumentationPartOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Updates an existing documentation part of a given API.
type UpdateDocumentationPartInput struct {

	// [Required] The identifier of the to-be-updated documentation part.
	//
	// This member is required.
	DocumentationPartId *string

	// [Required] The string identifier of the associated RestApi.
	//
	// This member is required.
	RestApiId *string

	Name *string

	// A list of update operations to be applied to the specified resource and in the
	// order specified in this list.
	PatchOperations []*types.PatchOperation

	Template *bool

	TemplateSkipList []*string

	Title *string
}

// A documentation part for a targeted API entity. A documentation part consists of
// a content map (properties) and a target (location). The target specifies an API
// entity to which the documentation content applies. The supported API entity
// types are API, AUTHORIZER, MODEL, RESOURCE, METHOD, PATH_PARAMETER,
// QUERY_PARAMETER, REQUEST_HEADER, REQUEST_BODY, RESPONSE, RESPONSE_HEADER, and
// RESPONSE_BODY. Valid location fields depend on the API entity type. All valid
// fields are not required. The content map is a JSON string of API-specific
// key-value pairs. Although an API can use any shape for the content map, only the
// OpenAPI-compliant documentation fields will be injected into the associated API
// entity definition in the exported OpenAPI definition file. Documenting an API
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-documenting-api.html),
// DocumentationParts
type UpdateDocumentationPartOutput struct {

	// The DocumentationPart identifier, generated by API Gateway when the
	// DocumentationPart is created.
	Id *string

	// The location of the API entity to which the documentation applies. Valid fields
	// depend on the targeted API entity type. All the valid location fields are not
	// required. If not explicitly specified, a valid location field is treated as a
	// wildcard and associated documentation content may be inherited by matching
	// entities, unless overridden.
	Location *types.DocumentationPartLocation

	// A content map of API-specific key-value pairs describing the targeted API
	// entity. The map must be encoded as a JSON string, e.g., "{ \"description\":
	// \"The API does ...\" }". Only OpenAPI-compliant documentation-related fields
	// from the properties map are exported and, hence, published as part of the API
	// entity definitions, while the original documentation parts are exported in a
	// OpenAPI extension of x-amazon-apigateway-documentation.
	Properties *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateDocumentationPartMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpUpdateDocumentationPart{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpUpdateDocumentationPart{})
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
	if err := addOpUpdateDocumentationPartValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opUpdateDocumentationPart(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opUpdateDocumentationPart(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "UpdateDocumentationPart",
	}
}
