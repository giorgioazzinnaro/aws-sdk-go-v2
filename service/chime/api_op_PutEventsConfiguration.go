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

// Creates an events configuration that allows a bot to receive outgoing events
// sent by Amazon Chime. Choose either an HTTPS endpoint or a Lambda function ARN.
// For more information, see Bot.
func (c *Client) PutEventsConfiguration(ctx context.Context, params *PutEventsConfigurationInput, optFns ...func(*Options)) (*PutEventsConfigurationOutput, error) {
	if params == nil {
		params = &PutEventsConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutEventsConfiguration", params, optFns, addOperationPutEventsConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutEventsConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutEventsConfigurationInput struct {

	// The Amazon Chime account ID.
	//
	// This member is required.
	AccountId *string

	// The bot ID.
	//
	// This member is required.
	BotId *string

	// Lambda function ARN that allows the bot to receive outgoing events.
	LambdaFunctionArn *string

	// HTTPS endpoint that allows the bot to receive outgoing events.
	OutboundEventsHTTPSEndpoint *string
}

type PutEventsConfigurationOutput struct {

	// The configuration that allows a bot to receive outgoing events. Can be either an
	// HTTPS endpoint or a Lambda function ARN.
	EventsConfiguration *types.EventsConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutEventsConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutEventsConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutEventsConfiguration{}, middleware.After)
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
	addOpPutEventsConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutEventsConfiguration(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opPutEventsConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "PutEventsConfiguration",
	}
}
