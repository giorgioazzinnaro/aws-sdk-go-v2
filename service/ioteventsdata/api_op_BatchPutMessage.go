// Code generated by smithy-go-codegen DO NOT EDIT.

package ioteventsdata

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ioteventsdata/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Sends a set of messages to the AWS IoT Events system. Each message payload is
// transformed into the input you specify ("inputName") and ingested into any
// detectors that monitor that input. If multiple messages are sent, the order in
// which the messages are processed isn't guaranteed. To guarantee ordering, you
// must send messages one at a time and wait for a successful response.
func (c *Client) BatchPutMessage(ctx context.Context, params *BatchPutMessageInput, optFns ...func(*Options)) (*BatchPutMessageOutput, error) {
	if params == nil {
		params = &BatchPutMessageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchPutMessage", params, optFns, addOperationBatchPutMessageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchPutMessageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchPutMessageInput struct {

	// The list of messages to send. Each message has the following format: '{
	// "messageId": "string", "inputName": "string", "payload": "string"}'
	//
	// This member is required.
	Messages []*types.Message
}

type BatchPutMessageOutput struct {

	// A list of any errors encountered when sending the messages.
	BatchPutMessageErrorEntries []*types.BatchPutMessageErrorEntry

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationBatchPutMessageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpBatchPutMessage{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpBatchPutMessage{})
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
	if err := addOpBatchPutMessageValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opBatchPutMessage(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opBatchPutMessage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ioteventsdata",
		OperationName: "BatchPutMessage",
	}
}
