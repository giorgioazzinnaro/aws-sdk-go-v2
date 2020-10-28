// Code generated by smithy-go-codegen DO NOT EDIT.

package workspaces

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workspaces/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Specifies which devices and operating systems users can use to access their
// WorkSpaces. For more information, see  Control Device Access
// (https://docs.aws.amazon.com/workspaces/latest/adminguide/update-directory-details.html#control-device-access).
func (c *Client) ModifyWorkspaceAccessProperties(ctx context.Context, params *ModifyWorkspaceAccessPropertiesInput, optFns ...func(*Options)) (*ModifyWorkspaceAccessPropertiesOutput, error) {
	if params == nil {
		params = &ModifyWorkspaceAccessPropertiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyWorkspaceAccessProperties", params, optFns, addOperationModifyWorkspaceAccessPropertiesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyWorkspaceAccessPropertiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyWorkspaceAccessPropertiesInput struct {

	// The identifier of the directory.
	//
	// This member is required.
	ResourceId *string

	// The device types and operating systems to enable or disable for access.
	//
	// This member is required.
	WorkspaceAccessProperties *types.WorkspaceAccessProperties
}

type ModifyWorkspaceAccessPropertiesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationModifyWorkspaceAccessPropertiesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpModifyWorkspaceAccessProperties{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpModifyWorkspaceAccessProperties{})
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
	if err := addOpModifyWorkspaceAccessPropertiesValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opModifyWorkspaceAccessProperties(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opModifyWorkspaceAccessProperties(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces",
		OperationName: "ModifyWorkspaceAccessProperties",
	}
}
