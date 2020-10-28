// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This API places an outbound call to a contact, and then initiates the contact
// flow. It performs the actions in the contact flow that's specified (in
// ContactFlowId). Agents are not involved in initiating the outbound API (that is,
// dialing the contact). If the contact flow places an outbound call to a contact,
// and then puts the contact in queue, that's when the call is routed to the agent,
// like any other inbound case. There is a 60 second dialing timeout for this
// operation. If the call is not connected after 60 seconds, it fails. UK numbers
// with a 447 prefix are not allowed by default. Before you can dial these UK
// mobile numbers, you must submit a service quota increase request. For more
// information, see Amazon Connect Service Quotas
// (https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-service-limits.html)
// in the Amazon Connect Administrator Guide.
func (c *Client) StartOutboundVoiceContact(ctx context.Context, params *StartOutboundVoiceContactInput, optFns ...func(*Options)) (*StartOutboundVoiceContactOutput, error) {
	if params == nil {
		params = &StartOutboundVoiceContactInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartOutboundVoiceContact", params, optFns, addOperationStartOutboundVoiceContactMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartOutboundVoiceContactOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartOutboundVoiceContactInput struct {

	// The identifier of the contact flow for the outbound call. To see the
	// ContactFlowId in the Amazon Connect console user interface, on the navigation
	// menu go to Routing, Contact Flows. Choose the contact flow. On the contact flow
	// page, under the name of the contact flow, choose Show additional flow
	// information. The ContactFlowId is the last part of the ARN, shown here in bold:
	// arn:aws:connect:us-west-2:xxxxxxxxxxxx:instance/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/contact-flow/846ec553-a005-41c0-8341-xxxxxxxxxxxx
	//
	// This member is required.
	ContactFlowId *string

	// The phone number of the customer, in E.164 format.
	//
	// This member is required.
	DestinationPhoneNumber *string

	// The identifier of the Amazon Connect instance.
	//
	// This member is required.
	InstanceId *string

	// A custom key-value pair using an attribute map. The attributes are standard
	// Amazon Connect attributes, and can be accessed in contact flows just like any
	// other contact attributes. There can be up to 32,768 UTF-8 bytes across all
	// key-value pairs per contact. Attribute keys can include only alphanumeric, dash,
	// and underscore characters.
	Attributes map[string]*string

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. The token is valid for 7 days after creation. If a contact is
	// already started, the contact ID is returned. If the contact is disconnected, a
	// new contact is started.
	ClientToken *string

	// The queue for the call. If you specify a queue, the phone displayed for caller
	// ID is the phone number specified in the queue. If you do not specify a queue,
	// the queue defined in the contact flow is used. If you do not specify a queue,
	// you must specify a source phone number.
	QueueId *string

	// The phone number associated with the Amazon Connect instance, in E.164 format.
	// If you do not specify a source phone number, you must specify a queue.
	SourcePhoneNumber *string
}

type StartOutboundVoiceContactOutput struct {

	// The identifier of this contact within the Amazon Connect instance.
	ContactId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStartOutboundVoiceContactMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpStartOutboundVoiceContact{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpStartOutboundVoiceContact{})
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
	if err := addIdempotencyToken_opStartOutboundVoiceContactMiddleware(stack, options); err != nil {
		return err
	}
	if err := addOpStartOutboundVoiceContactValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opStartOutboundVoiceContact(options.Region)); err != nil {
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

type idempotencyToken_initializeOpStartOutboundVoiceContact struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStartOutboundVoiceContact) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStartOutboundVoiceContact) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StartOutboundVoiceContactInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StartOutboundVoiceContactInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opStartOutboundVoiceContactMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(middleware.Before, &idempotencyToken_initializeOpStartOutboundVoiceContact{tokenProvider: cfg.IdempotencyTokenProvider})
}

func newServiceMetadataMiddleware_opStartOutboundVoiceContact(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "StartOutboundVoiceContact",
	}
}
