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
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpPutEventsConfiguration{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpPutEventsConfiguration{})
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
	if err := addOpPutEventsConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opPutEventsConfiguration(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opPutEventsConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "PutEventsConfiguration",
	}
}
