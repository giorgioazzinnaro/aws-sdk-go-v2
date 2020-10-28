// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes the suspension from up to 50 previously suspended users for the
// specified Amazon Chime EnterpriseLWA account. Only users on EnterpriseLWA
// accounts can be unsuspended using this action. For more information about
// different account types, see Managing Your Amazon Chime Accounts
// (https://docs.aws.amazon.com/chime/latest/ag/manage-chime-account.html) in the
// Amazon Chime Administration Guide. Previously suspended users who are
// unsuspended using this action are returned to Registered status. Users who are
// not previously suspended are ignored.
func (c *Client) BatchUnsuspendUser(ctx context.Context, params *BatchUnsuspendUserInput, optFns ...func(*Options)) (*BatchUnsuspendUserOutput, error) {
	if params == nil {
		params = &BatchUnsuspendUserInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchUnsuspendUser", params, optFns, addOperationBatchUnsuspendUserMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchUnsuspendUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchUnsuspendUserInput struct {

	// The Amazon Chime account ID.
	//
	// This member is required.
	AccountId *string

	// The request containing the user IDs to unsuspend.
	//
	// This member is required.
	UserIdList []*string
}

type BatchUnsuspendUserOutput struct {

	// If the BatchUnsuspendUser action fails for one or more of the user IDs in the
	// request, a list of the user IDs is returned, along with error codes and error
	// messages.
	UserErrors []*types.UserError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationBatchUnsuspendUserMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpBatchUnsuspendUser{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpBatchUnsuspendUser{})
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
	if err := addOpBatchUnsuspendUserValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opBatchUnsuspendUser(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opBatchUnsuspendUser(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "BatchUnsuspendUser",
	}
}
