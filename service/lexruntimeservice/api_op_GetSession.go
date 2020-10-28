// Code generated by smithy-go-codegen DO NOT EDIT.

package lexruntimeservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lexruntimeservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns session information for a specified bot, alias, and user ID.
func (c *Client) GetSession(ctx context.Context, params *GetSessionInput, optFns ...func(*Options)) (*GetSessionOutput, error) {
	if params == nil {
		params = &GetSessionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSession", params, optFns, addOperationGetSessionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSessionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSessionInput struct {

	// The alias in use for the bot that contains the session data.
	//
	// This member is required.
	BotAlias *string

	// The name of the bot that contains the session data.
	//
	// This member is required.
	BotName *string

	// The ID of the client application user. Amazon Lex uses this to identify a user's
	// conversation with your bot.
	//
	// This member is required.
	UserId *string

	// A string used to filter the intents returned in the recentIntentSummaryView
	// structure. When you specify a filter, only intents with their checkpointLabel
	// field set to that string are returned.
	CheckpointLabelFilter *string
}

type GetSessionOutput struct {

	// Describes the current state of the bot.
	DialogAction *types.DialogAction

	// An array of information about the intents used in the session. The array can
	// contain a maximum of three summaries. If more than three intents are used in the
	// session, the recentIntentSummaryView operation contains information about the
	// last three intents used. If you set the checkpointLabelFilter parameter in the
	// request, the array contains only the intents with the specified label.
	RecentIntentSummaryView []*types.IntentSummary

	// Map of key/value pairs representing the session-specific context information. It
	// contains application information passed between Amazon Lex and a client
	// application.
	SessionAttributes map[string]*string

	// A unique identifier for the session.
	SessionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetSessionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpGetSession{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpGetSession{})
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
	if err := addOpGetSessionValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opGetSession(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opGetSession(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "GetSession",
	}
}
