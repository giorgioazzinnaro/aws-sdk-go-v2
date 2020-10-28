// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes an alias for the specified bot. You can't delete an alias that is used
// in the association between a bot and a messaging channel. If an alias is used in
// a channel association, the DeleteBot operation returns a ResourceInUseException
// exception that includes a reference to the channel association that refers to
// the bot. You can remove the reference to the alias by deleting the channel
// association. If you get the same exception again, delete the referring
// association until the DeleteBotAlias operation is successful.
func (c *Client) DeleteBotAlias(ctx context.Context, params *DeleteBotAliasInput, optFns ...func(*Options)) (*DeleteBotAliasOutput, error) {
	if params == nil {
		params = &DeleteBotAliasInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteBotAlias", params, optFns, addOperationDeleteBotAliasMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteBotAliasOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteBotAliasInput struct {

	// The name of the bot that the alias points to.
	//
	// This member is required.
	BotName *string

	// The name of the alias to delete. The name is case sensitive.
	//
	// This member is required.
	Name *string
}

type DeleteBotAliasOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteBotAliasMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpDeleteBotAlias{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpDeleteBotAlias{})
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
	if err := addOpDeleteBotAliasValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDeleteBotAlias(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDeleteBotAlias(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "DeleteBotAlias",
	}
}
