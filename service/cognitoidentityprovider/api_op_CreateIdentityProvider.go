// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an identity provider for a user pool.
func (c *Client) CreateIdentityProvider(ctx context.Context, params *CreateIdentityProviderInput, optFns ...func(*Options)) (*CreateIdentityProviderOutput, error) {
	if params == nil {
		params = &CreateIdentityProviderInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateIdentityProvider", params, optFns, addOperationCreateIdentityProviderMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateIdentityProviderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateIdentityProviderInput struct {

	// The identity provider details. The following list describes the provider detail
	// keys for each identity provider type.
	//
	//     * For Google and Login with Amazon:
	//
	//
	// * client_id
	//
	//         * client_secret
	//
	//         * authorize_scopes
	//
	//     * For
	// Facebook:
	//
	//         * client_id
	//
	//         * client_secret
	//
	//         *
	// authorize_scopes
	//
	//         * api_version
	//
	//     * For Sign in with Apple:
	//
	//
	// * client_id
	//
	//         * team_id
	//
	//         * key_id
	//
	//         * private_key
	//
	//
	// * authorize_scopes
	//
	//     * For OIDC providers:
	//
	//         * client_id
	//
	//         *
	// client_secret
	//
	//         * attributes_request_method
	//
	//         * oidc_issuer
	//
	//
	// * authorize_scopes
	//
	//         * authorize_url if not available from discovery URL
	// specified by oidc_issuer key
	//
	//         * token_url if not available from
	// discovery URL specified by oidc_issuer key
	//
	//         * attributes_url if not
	// available from discovery URL specified by oidc_issuer key
	//
	//         * jwks_uri if
	// not available from discovery URL specified by oidc_issuer key
	//
	//     * For SAML
	// providers:
	//
	//         * MetadataFile OR MetadataURL
	//
	//         * IDPSignout optional
	//
	// This member is required.
	ProviderDetails map[string]*string

	// The identity provider name.
	//
	// This member is required.
	ProviderName *string

	// The identity provider type.
	//
	// This member is required.
	ProviderType types.IdentityProviderTypeType

	// The user pool ID.
	//
	// This member is required.
	UserPoolId *string

	// A mapping of identity provider attributes to standard and custom user pool
	// attributes.
	AttributeMapping map[string]*string

	// A list of identity provider identifiers.
	IdpIdentifiers []*string
}

type CreateIdentityProviderOutput struct {

	// The newly created identity provider object.
	//
	// This member is required.
	IdentityProvider *types.IdentityProviderType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateIdentityProviderMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpCreateIdentityProvider{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpCreateIdentityProvider{})
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
	if err := addOpCreateIdentityProviderValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opCreateIdentityProvider(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opCreateIdentityProvider(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "CreateIdentityProvider",
	}
}
