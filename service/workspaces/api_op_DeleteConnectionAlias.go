// Code generated by smithy-go-codegen DO NOT EDIT.

package workspaces

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the specified connection alias. For more information, see  Cross-Region
// Redirection for Amazon WorkSpaces
// (https://docs.aws.amazon.com/workspaces/latest/adminguide/cross-region-redirection.html).
// If you will no longer be using a fully qualified domain name (FQDN) as the
// registration code for your WorkSpaces users, you must take certain precautions
// to prevent potential security issues. For more information, see  Security
// Considerations if You Stop Using Cross-Region Redirection
// (https://docs.aws.amazon.com/workspaces/latest/adminguide/cross-region-redirection.html#cross-region-redirection-security-considerations).
// To delete a connection alias that has been shared, the shared account must first
// disassociate the connection alias from any directories it has been associated
// with. Then you must unshare the connection alias from the account it has been
// shared with. You can delete a connection alias only after it is no longer shared
// with any accounts or associated with any directories.
func (c *Client) DeleteConnectionAlias(ctx context.Context, params *DeleteConnectionAliasInput, optFns ...func(*Options)) (*DeleteConnectionAliasOutput, error) {
	if params == nil {
		params = &DeleteConnectionAliasInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteConnectionAlias", params, optFns, addOperationDeleteConnectionAliasMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteConnectionAliasOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteConnectionAliasInput struct {

	// The identifier of the connection alias to delete.
	//
	// This member is required.
	AliasId *string
}

type DeleteConnectionAliasOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteConnectionAliasMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpDeleteConnectionAlias{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpDeleteConnectionAlias{})
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
	if err := addOpDeleteConnectionAliasValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDeleteConnectionAlias(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDeleteConnectionAlias(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces",
		OperationName: "DeleteConnectionAlias",
	}
}
