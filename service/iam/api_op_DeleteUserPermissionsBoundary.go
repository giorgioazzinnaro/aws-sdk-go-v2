// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the permissions boundary for the specified IAM user. Deleting the
// permissions boundary for a user might increase its permissions by allowing the
// user to perform all the actions granted in its permissions policies.
func (c *Client) DeleteUserPermissionsBoundary(ctx context.Context, params *DeleteUserPermissionsBoundaryInput, optFns ...func(*Options)) (*DeleteUserPermissionsBoundaryOutput, error) {
	if params == nil {
		params = &DeleteUserPermissionsBoundaryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteUserPermissionsBoundary", params, optFns, addOperationDeleteUserPermissionsBoundaryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteUserPermissionsBoundaryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteUserPermissionsBoundaryInput struct {

	// The name (friendly name, not ARN) of the IAM user from which you want to remove
	// the permissions boundary.
	//
	// This member is required.
	UserName *string
}

type DeleteUserPermissionsBoundaryOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteUserPermissionsBoundaryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsquery_serializeOpDeleteUserPermissionsBoundary{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsquery_deserializeOpDeleteUserPermissionsBoundary{})
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
	if err := addOpDeleteUserPermissionsBoundaryValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDeleteUserPermissionsBoundary(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDeleteUserPermissionsBoundary(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "DeleteUserPermissionsBoundary",
	}
}
