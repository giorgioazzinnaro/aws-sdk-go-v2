// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentity

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets details about a particular identity pool, including the pool name, ID
// description, creation date, and current number of users. You must use AWS
// Developer credentials to call this API.
func (c *Client) DescribeIdentityPool(ctx context.Context, params *DescribeIdentityPoolInput, optFns ...func(*Options)) (*DescribeIdentityPoolOutput, error) {
	if params == nil {
		params = &DescribeIdentityPoolInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeIdentityPool", params, optFns, addOperationDescribeIdentityPoolMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeIdentityPoolOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input to the DescribeIdentityPool action.
type DescribeIdentityPoolInput struct {

	// An identity pool ID in the format REGION:GUID.
	//
	// This member is required.
	IdentityPoolId *string
}

// An object representing an Amazon Cognito identity pool.
type DescribeIdentityPoolOutput struct {

	// TRUE if the identity pool supports unauthenticated logins.
	//
	// This member is required.
	AllowUnauthenticatedIdentities *bool

	// An identity pool ID in the format REGION:GUID.
	//
	// This member is required.
	IdentityPoolId *string

	// A string that you provide.
	//
	// This member is required.
	IdentityPoolName *string

	// Enables or disables the Basic (Classic) authentication flow. For more
	// information, see Identity Pools (Federated Identities) Authentication Flow
	// (https://docs.aws.amazon.com/cognito/latest/developerguide/authentication-flow.html)
	// in the Amazon Cognito Developer Guide.
	AllowClassicFlow *bool

	// A list representing an Amazon Cognito user pool and its client ID.
	CognitoIdentityProviders []*types.CognitoIdentityProvider

	// The "domain" by which Cognito will refer to your users.
	DeveloperProviderName *string

	// The tags that are assigned to the identity pool. A tag is a label that you can
	// apply to identity pools to categorize and manage them in different ways, such as
	// by purpose, owner, environment, or other criteria.
	IdentityPoolTags map[string]*string

	// A list of OpendID Connect provider ARNs.
	OpenIdConnectProviderARNs []*string

	// An array of Amazon Resource Names (ARNs) of the SAML provider for your identity
	// pool.
	SamlProviderARNs []*string

	// Optional key:value pairs mapping provider names to provider app IDs.
	SupportedLoginProviders map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeIdentityPoolMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpDescribeIdentityPool{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpDescribeIdentityPool{})
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
	if err := addOpDescribeIdentityPoolValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDescribeIdentityPool(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDescribeIdentityPool(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-identity",
		OperationName: "DescribeIdentityPool",
	}
}
