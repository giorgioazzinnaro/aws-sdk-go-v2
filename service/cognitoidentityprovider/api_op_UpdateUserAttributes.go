// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Allows a user to update a specific attribute (one at a time).
func (c *Client) UpdateUserAttributes(ctx context.Context, params *UpdateUserAttributesInput, optFns ...func(*Options)) (*UpdateUserAttributesOutput, error) {
	if params == nil {
		params = &UpdateUserAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateUserAttributes", params, optFns, addOperationUpdateUserAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateUserAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to update user attributes.
type UpdateUserAttributesInput struct {

	// The access token for the request to update user attributes.
	//
	// This member is required.
	AccessToken *string

	// An array of name-value pairs representing user attributes. For custom
	// attributes, you must prepend the custom: prefix to the attribute name.
	//
	// This member is required.
	UserAttributes []*types.AttributeType

	// A map of custom key-value pairs that you can provide as input for any custom
	// workflows that this action triggers. You create custom workflows by assigning
	// AWS Lambda functions to user pool triggers. When you use the
	// UpdateUserAttributes API action, Amazon Cognito invokes the function that is
	// assigned to the custom message trigger. When Amazon Cognito invokes this
	// function, it passes a JSON payload, which the function receives as input. This
	// payload contains a clientMetadata attribute, which provides the data that you
	// assigned to the ClientMetadata parameter in your UpdateUserAttributes request.
	// In your function code in AWS Lambda, you can process the clientMetadata value to
	// enhance your workflow for your specific needs. For more information, see
	// Customizing User Pool Workflows with Lambda Triggers
	// (https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools-working-with-aws-lambda-triggers.html)
	// in the Amazon Cognito Developer Guide. Take the following limitations into
	// consideration when you use the ClientMetadata parameter:
	//
	// * Amazon Cognito does
	// not store the ClientMetadata value. This data is available only to AWS Lambda
	// triggers that are assigned to a user pool to support custom workflows. If your
	// user pool configuration does not include triggers, the ClientMetadata parameter
	// serves no purpose.
	//
	// * Amazon Cognito does not validate the ClientMetadata
	// value.
	//
	// * Amazon Cognito does not encrypt the the ClientMetadata value, so don't
	// use it to provide sensitive information.
	ClientMetadata map[string]*string
}

// Represents the response from the server for the request to update user
// attributes.
type UpdateUserAttributesOutput struct {

	// The code delivery details list from the server for the request to update user
	// attributes.
	CodeDeliveryDetailsList []*types.CodeDeliveryDetailsType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateUserAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateUserAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateUserAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	addRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpUpdateUserAttributesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateUserAttributes(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateUserAttributes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateUserAttributes",
	}
}
