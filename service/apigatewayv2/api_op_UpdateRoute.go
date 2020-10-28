// Code generated by smithy-go-codegen DO NOT EDIT.

package apigatewayv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates a Route.
func (c *Client) UpdateRoute(ctx context.Context, params *UpdateRouteInput, optFns ...func(*Options)) (*UpdateRouteOutput, error) {
	if params == nil {
		params = &UpdateRouteInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateRoute", params, optFns, addOperationUpdateRouteMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateRouteOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Updates a Route.
type UpdateRouteInput struct {

	// The API identifier.
	//
	// This member is required.
	ApiId *string

	// The route ID.
	//
	// This member is required.
	RouteId *string

	// Specifies whether an API key is required for the route. Supported only for
	// WebSocket APIs.
	ApiKeyRequired *bool

	// The authorization scopes supported by this route.
	AuthorizationScopes []*string

	// The authorization type for the route. For WebSocket APIs, valid values are NONE
	// for open access, AWS_IAM for using AWS IAM permissions, and CUSTOM for using a
	// Lambda authorizer For HTTP APIs, valid values are NONE for open access, JWT for
	// using JSON Web Tokens, AWS_IAM for using AWS IAM permissions, and CUSTOM for
	// using a Lambda authorizer.
	AuthorizationType types.AuthorizationType

	// The identifier of the Authorizer resource to be associated with this route. The
	// authorizer identifier is generated by API Gateway when you created the
	// authorizer.
	AuthorizerId *string

	// The model selection expression for the route. Supported only for WebSocket APIs.
	ModelSelectionExpression *string

	// The operation name for the route.
	OperationName *string

	// The request models for the route. Supported only for WebSocket APIs.
	RequestModels map[string]*string

	// The request parameters for the route. Supported only for WebSocket APIs.
	RequestParameters map[string]*types.ParameterConstraints

	// The route key for the route.
	RouteKey *string

	// The route response selection expression for the route. Supported only for
	// WebSocket APIs.
	RouteResponseSelectionExpression *string

	// The target for the route.
	Target *string
}

type UpdateRouteOutput struct {

	// Specifies whether a route is managed by API Gateway. If you created an API using
	// quick create, the $default route is managed by API Gateway. You can't modify the
	// $default route key.
	ApiGatewayManaged *bool

	// Specifies whether an API key is required for this route. Supported only for
	// WebSocket APIs.
	ApiKeyRequired *bool

	// A list of authorization scopes configured on a route. The scopes are used with a
	// JWT authorizer to authorize the method invocation. The authorization works by
	// matching the route scopes against the scopes parsed from the access token in the
	// incoming request. The method invocation is authorized if any route scope matches
	// a claimed scope in the access token. Otherwise, the invocation is not
	// authorized. When the route scope is configured, the client must provide an
	// access token instead of an identity token for authorization purposes.
	AuthorizationScopes []*string

	// The authorization type for the route. For WebSocket APIs, valid values are NONE
	// for open access, AWS_IAM for using AWS IAM permissions, and CUSTOM for using a
	// Lambda authorizer For HTTP APIs, valid values are NONE for open access, JWT for
	// using JSON Web Tokens, AWS_IAM for using AWS IAM permissions, and CUSTOM for
	// using a Lambda authorizer.
	AuthorizationType types.AuthorizationType

	// The identifier of the Authorizer resource to be associated with this route. The
	// authorizer identifier is generated by API Gateway when you created the
	// authorizer.
	AuthorizerId *string

	// The model selection expression for the route. Supported only for WebSocket APIs.
	ModelSelectionExpression *string

	// The operation name for the route.
	OperationName *string

	// The request models for the route. Supported only for WebSocket APIs.
	RequestModels map[string]*string

	// The request parameters for the route. Supported only for WebSocket APIs.
	RequestParameters map[string]*types.ParameterConstraints

	// The route ID.
	RouteId *string

	// The route key for the route.
	RouteKey *string

	// The route response selection expression for the route. Supported only for
	// WebSocket APIs.
	RouteResponseSelectionExpression *string

	// The target for the route.
	Target *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateRouteMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpUpdateRoute{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpUpdateRoute{})
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
	if err := addOpUpdateRouteValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opUpdateRoute(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opUpdateRoute(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "UpdateRoute",
	}
}
