// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the password policy settings for the AWS account.
//
//     * This operation
// does not support partial updates. No parameters are required, but if you do not
// specify a parameter, that parameter's value reverts to its default value. See
// the Request Parameters section for each parameter's default value. Also note
// that some parameters do not allow the default parameter to be explicitly set.
// Instead, to invoke the default value, do not include that parameter when you
// invoke the operation.
//
// For more information about using a password policy, see
// Managing an IAM Password Policy
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_ManagingPasswordPolicies.html)
// in the IAM User Guide.
func (c *Client) UpdateAccountPasswordPolicy(ctx context.Context, params *UpdateAccountPasswordPolicyInput, optFns ...func(*Options)) (*UpdateAccountPasswordPolicyOutput, error) {
	if params == nil {
		params = &UpdateAccountPasswordPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateAccountPasswordPolicy", params, optFns, addOperationUpdateAccountPasswordPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAccountPasswordPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAccountPasswordPolicyInput struct {

	// Allows all IAM users in your account to use the AWS Management Console to change
	// their own passwords. For more information, see Letting IAM Users Change Their
	// Own Passwords
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/HowToPwdIAMUser.html) in the
	// IAM User Guide. If you do not specify a value for this parameter, then the
	// operation uses the default value of false. The result is that IAM users in the
	// account do not automatically have permissions to change their own password.
	AllowUsersToChangePassword *bool

	// Prevents IAM users from setting a new password after their password has expired.
	// The IAM user cannot be accessed until an administrator resets the password. If
	// you do not specify a value for this parameter, then the operation uses the
	// default value of false. The result is that IAM users can change their passwords
	// after they expire and continue to sign in as the user.
	HardExpiry *bool

	// The number of days that an IAM user password is valid. If you do not specify a
	// value for this parameter, then the operation uses the default value of 0. The
	// result is that IAM user passwords never expire.
	MaxPasswordAge *int32

	// The minimum number of characters allowed in an IAM user password. If you do not
	// specify a value for this parameter, then the operation uses the default value of
	// 6.
	MinimumPasswordLength *int32

	// Specifies the number of previous passwords that IAM users are prevented from
	// reusing. If you do not specify a value for this parameter, then the operation
	// uses the default value of 0. The result is that IAM users are not prevented from
	// reusing previous passwords.
	PasswordReusePrevention *int32

	// Specifies whether IAM user passwords must contain at least one lowercase
	// character from the ISO basic Latin alphabet (a to z). If you do not specify a
	// value for this parameter, then the operation uses the default value of false.
	// The result is that passwords do not require at least one lowercase character.
	RequireLowercaseCharacters *bool

	// Specifies whether IAM user passwords must contain at least one numeric character
	// (0 to 9). If you do not specify a value for this parameter, then the operation
	// uses the default value of false. The result is that passwords do not require at
	// least one numeric character.
	RequireNumbers *bool

	// Specifies whether IAM user passwords must contain at least one of the following
	// non-alphanumeric characters: ! @ # $ % ^ & * ( ) _ + - = [ ] { } | ' If you do
	// not specify a value for this parameter, then the operation uses the default
	// value of false. The result is that passwords do not require at least one symbol
	// character.
	RequireSymbols *bool

	// Specifies whether IAM user passwords must contain at least one uppercase
	// character from the ISO basic Latin alphabet (A to Z). If you do not specify a
	// value for this parameter, then the operation uses the default value of false.
	// The result is that passwords do not require at least one uppercase character.
	RequireUppercaseCharacters *bool
}

type UpdateAccountPasswordPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateAccountPasswordPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsquery_serializeOpUpdateAccountPasswordPolicy{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsquery_deserializeOpUpdateAccountPasswordPolicy{})
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
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opUpdateAccountPasswordPolicy(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opUpdateAccountPasswordPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "UpdateAccountPasswordPolicy",
	}
}
