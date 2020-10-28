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

// Set the user's multi-factor authentication (MFA) method preference, including
// which MFA factors are enabled and if any are preferred. Only one factor can be
// set as preferred. The preferred MFA factor will be used to authenticate a user
// if multiple factors are enabled. If multiple options are enabled and no
// preference is set, a challenge to choose an MFA option will be returned during
// sign in.
func (c *Client) SetUserMFAPreference(ctx context.Context, params *SetUserMFAPreferenceInput, optFns ...func(*Options)) (*SetUserMFAPreferenceOutput, error) {
	if params == nil {
		params = &SetUserMFAPreferenceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetUserMFAPreference", params, optFns, addOperationSetUserMFAPreferenceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetUserMFAPreferenceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SetUserMFAPreferenceInput struct {

	// The access token for the user.
	//
	// This member is required.
	AccessToken *string

	// The SMS text message multi-factor authentication (MFA) settings.
	SMSMfaSettings *types.SMSMfaSettingsType

	// The time-based one-time password software token MFA settings.
	SoftwareTokenMfaSettings *types.SoftwareTokenMfaSettingsType
}

type SetUserMFAPreferenceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationSetUserMFAPreferenceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpSetUserMFAPreference{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpSetUserMFAPreference{})
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
	if err := addOpSetUserMFAPreferenceValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opSetUserMFAPreference(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opSetUserMFAPreference(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "SetUserMFAPreference",
	}
}
