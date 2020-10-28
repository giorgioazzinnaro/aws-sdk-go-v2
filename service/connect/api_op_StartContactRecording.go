// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This API starts recording the contact when the agent joins the call.
// StartContactRecording is a one-time action. For example, if you use
// StopContactRecording to stop recording an ongoing call, you can't use
// StartContactRecording to restart it. For scenarios where the recording has
// started and you want to suspend and resume it, such as when collecting sensitive
// information (for example, a credit card number), use SuspendContactRecording and
// ResumeContactRecording. You can use this API to override the recording behavior
// configured in the Set recording behavior
// (https://docs.aws.amazon.com/connect/latest/adminguide/set-recording-behavior.html)
// block. Only voice recordings are supported at this time.
func (c *Client) StartContactRecording(ctx context.Context, params *StartContactRecordingInput, optFns ...func(*Options)) (*StartContactRecordingOutput, error) {
	if params == nil {
		params = &StartContactRecordingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartContactRecording", params, optFns, addOperationStartContactRecordingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartContactRecordingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartContactRecordingInput struct {

	// The identifier of the contact.
	//
	// This member is required.
	ContactId *string

	// The identifier of the contact. This is the identifier of the contact associated
	// with the first interaction with the contact center.
	//
	// This member is required.
	InitialContactId *string

	// The identifier of the Amazon Connect instance.
	//
	// This member is required.
	InstanceId *string

	// Who is being recorded.
	//
	// This member is required.
	VoiceRecordingConfiguration *types.VoiceRecordingConfiguration
}

type StartContactRecordingOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStartContactRecordingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpStartContactRecording{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpStartContactRecording{})
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
	if err := addOpStartContactRecordingValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opStartContactRecording(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opStartContactRecording(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "StartContactRecording",
	}
}
