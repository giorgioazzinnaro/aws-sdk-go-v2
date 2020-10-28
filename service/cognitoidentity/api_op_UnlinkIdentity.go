// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentity

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Unlinks a federated identity from an existing account. Unlinked logins will be
// considered new identities next time they are seen. Removing the last linked
// login will make this identity inaccessible. This is a public API. You do not
// need any credentials to call this API.
func (c *Client) UnlinkIdentity(ctx context.Context, params *UnlinkIdentityInput, optFns ...func(*Options)) (*UnlinkIdentityOutput, error) {
	if params == nil {
		params = &UnlinkIdentityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UnlinkIdentity", params, optFns, addOperationUnlinkIdentityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UnlinkIdentityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input to the UnlinkIdentity action.
type UnlinkIdentityInput struct {

	// A unique identifier in the format REGION:GUID.
	//
	// This member is required.
	IdentityId *string

	// A set of optional name-value pairs that map provider names to provider tokens.
	//
	// This member is required.
	Logins map[string]*string

	// Provider names to unlink from this identity.
	//
	// This member is required.
	LoginsToRemove []*string
}

type UnlinkIdentityOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUnlinkIdentityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpUnlinkIdentity{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpUnlinkIdentity{})
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
	if err := addRetryMiddlewares(stack, options); err != nil {
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
	if err := addOpUnlinkIdentityValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opUnlinkIdentity(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opUnlinkIdentity(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-identity",
		OperationName: "UnlinkIdentity",
	}
}
