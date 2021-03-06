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

// Returns intent information as follows:
//
// * If you specify the nameContains field,
// returns the $LATEST version of all intents that contain the specified string.
//
// *
// If you don't specify the nameContains field, returns information about the
// $LATEST version of all intents.
//
// The operation requires permission for the
// lex:GetIntents action.
func (c *Client) GetIntents(ctx context.Context, params *GetIntentsInput, optFns ...func(*Options)) (*GetIntentsOutput, error) {
	if params == nil {
		params = &GetIntentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetIntents", params, optFns, addOperationGetIntentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetIntentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetIntentsInput struct {

	// The maximum number of intents to return in the response. The default is 10.
	MaxResults *int32

	// Substring to match in intent names. An intent will be returned if any part of
	// its name matches the substring. For example, "xyz" matches both "xyzabc" and
	// "abcxyz."
	NameContains *string

	// A pagination token that fetches the next page of intents. If the response to
	// this API call is truncated, Amazon Lex returns a pagination token in the
	// response. To fetch the next page of intents, specify the pagination token in the
	// next request.
	NextToken *string
}

type GetIntentsOutput struct {

	// An array of Intent objects. For more information, see PutBot.
	Intents []*types.IntentMetadata

	// If the response is truncated, the response includes a pagination token that you
	// can specify in your next request to fetch the next page of intents.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetIntentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetIntents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetIntents{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetIntents(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetIntents(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "GetIntents",
	}
}
