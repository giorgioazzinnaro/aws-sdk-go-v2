// Code generated by smithy-go-codegen DO NOT EDIT.

package ssoadmin

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the inline policy from a specified permission set.
func (c *Client) DeleteInlinePolicyFromPermissionSet(ctx context.Context, params *DeleteInlinePolicyFromPermissionSetInput, optFns ...func(*Options)) (*DeleteInlinePolicyFromPermissionSetOutput, error) {
	if params == nil {
		params = &DeleteInlinePolicyFromPermissionSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteInlinePolicyFromPermissionSet", params, optFns, addOperationDeleteInlinePolicyFromPermissionSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteInlinePolicyFromPermissionSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteInlinePolicyFromPermissionSetInput struct {

	// The ARN of the SSO instance under which the operation will be executed. For more
	// information about ARNs, see Amazon Resource Names (ARNs) and AWS Service
	// Namespaces in the AWS General Reference.
	//
	// This member is required.
	InstanceArn *string

	// The ARN of the permission set that will be used to remove access.
	//
	// This member is required.
	PermissionSetArn *string
}

type DeleteInlinePolicyFromPermissionSetOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteInlinePolicyFromPermissionSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpDeleteInlinePolicyFromPermissionSet{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpDeleteInlinePolicyFromPermissionSet{})
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
	if err := addOpDeleteInlinePolicyFromPermissionSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDeleteInlinePolicyFromPermissionSet(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDeleteInlinePolicyFromPermissionSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sso",
		OperationName: "DeleteInlinePolicyFromPermissionSet",
	}
}
