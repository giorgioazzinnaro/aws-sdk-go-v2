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

// Lists the users in the Amazon Cognito user pool.
func (c *Client) ListUsers(ctx context.Context, params *ListUsersInput, optFns ...func(*Options)) (*ListUsersOutput, error) {
	if params == nil {
		params = &ListUsersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListUsers", params, optFns, addOperationListUsersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListUsersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to list users.
type ListUsersInput struct {

	// The user pool ID for the user pool on which the search should be performed.
	//
	// This member is required.
	UserPoolId *string

	// An array of strings, where each string is the name of a user attribute to be
	// returned for each user in the search results. If the array is null, all
	// attributes are returned.
	AttributesToGet []*string

	// A filter string of the form "AttributeName Filter-Type "AttributeValue"".
	// Quotation marks within the filter string must be escaped using the backslash (\)
	// character. For example, "family_name = \"Reddy\"".
	//
	//     * AttributeName: The
	// name of the attribute to search for. You can only search for one attribute at a
	// time.
	//
	//     * Filter-Type: For an exact match, use =, for example, "given_name =
	// \"Jon\"". For a prefix ("starts with") match, use ^=, for example, "given_name
	// ^= \"Jon\"".
	//
	//     * AttributeValue: The attribute value that must be matched for
	// each user.
	//
	// If the filter string is empty, ListUsers returns all users in the
	// user pool. You can only search for the following standard attributes:
	//
	//     *
	// username (case-sensitive)
	//
	//     * email
	//
	//     * phone_number
	//
	//     * name
	//
	//     *
	// given_name
	//
	//     * family_name
	//
	//     * preferred_username
	//
	//     *
	// cognito:user_status (called Status in the Console) (case-insensitive)
	//
	//     *
	// status (called Enabled in the Console) (case-sensitive)
	//
	//     * sub
	//
	// Custom
	// attributes are not searchable. For more information, see Searching for Users
	// Using the ListUsers API
	// (https://docs.aws.amazon.com/cognito/latest/developerguide/how-to-manage-user-accounts.html#cognito-user-pools-searching-for-users-using-listusers-api)
	// and Examples of Using the ListUsers API
	// (https://docs.aws.amazon.com/cognito/latest/developerguide/how-to-manage-user-accounts.html#cognito-user-pools-searching-for-users-listusers-api-examples)
	// in the Amazon Cognito Developer Guide.
	Filter *string

	// Maximum number of users to be returned.
	Limit *int32

	// An identifier that was returned from the previous call to this operation, which
	// can be used to return the next set of items in the list.
	PaginationToken *string
}

// The response from the request to list users.
type ListUsersOutput struct {

	// An identifier that was returned from the previous call to this operation, which
	// can be used to return the next set of items in the list.
	PaginationToken *string

	// The users returned in the request to list users.
	Users []*types.UserType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListUsersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpListUsers{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpListUsers{})
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
	if err := addOpListUsersValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opListUsers(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opListUsers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "ListUsers",
	}
}
