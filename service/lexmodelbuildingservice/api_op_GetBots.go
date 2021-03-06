// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns bot information as follows:
//
// * If you provide the nameContains field,
// the response includes information for the $LATEST version of all bots whose name
// contains the specified string.
//
// * If you don't specify the nameContains field,
// the operation returns information about the $LATEST version of all of your
// bots.
//
// This operation requires permission for the lex:GetBots action.
func (c *Client) GetBots(ctx context.Context, params *GetBotsInput, optFns ...func(*Options)) (*GetBotsOutput, error) {
	if params == nil {
		params = &GetBotsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBots", params, optFns, addOperationGetBotsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBotsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBotsInput struct {

	// The maximum number of bots to return in the response that the request will
	// return. The default is 10.
	MaxResults *int32

	// Substring to match in bot names. A bot will be returned if any part of its name
	// matches the substring. For example, "xyz" matches both "xyzabc" and "abcxyz."
	NameContains *string

	// A pagination token that fetches the next page of bots. If the response to this
	// call is truncated, Amazon Lex returns a pagination token in the response. To
	// fetch the next page of bots, specify the pagination token in the next request.
	NextToken *string
}

type GetBotsOutput struct {

	// An array of botMetadata objects, with one entry for each bot.
	Bots []*types.BotMetadata

	// If the response is truncated, it includes a pagination token that you can
	// specify in your next request to fetch the next page of bots.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetBotsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetBots{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetBots{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetBots(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetBots(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "GetBots",
	}
}
