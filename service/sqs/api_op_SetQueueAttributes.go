// Code generated by smithy-go-codegen DO NOT EDIT.

package sqs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Sets the value of one or more queue attributes. When you change a queue's
// attributes, the change can take up to 60 seconds for most of the attributes to
// propagate throughout the Amazon SQS system. Changes made to the
// MessageRetentionPeriod attribute can take up to 15 minutes.
//
//     * In the
// future, new attributes might be added. If you write code that calls this action,
// we recommend that you structure your code so that it can handle new attributes
// gracefully.
//
//     * Cross-account permissions don't apply to this action. For
// more information, see Grant Cross-Account Permissions to a Role and a User Name
// (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-customer-managed-policy-examples.html#grant-cross-account-permissions-to-role-and-user-name)
// in the Amazon Simple Queue Service Developer Guide.
//
//     * To remove the ability
// to change queue permissions, you must deny permission to the AddPermission,
// RemovePermission, and SetQueueAttributes actions in your IAM policy.
func (c *Client) SetQueueAttributes(ctx context.Context, params *SetQueueAttributesInput, optFns ...func(*Options)) (*SetQueueAttributesOutput, error) {
	if params == nil {
		params = &SetQueueAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetQueueAttributes", params, optFns, addOperationSetQueueAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetQueueAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type SetQueueAttributesInput struct {

	// A map of attributes to set. The following lists the names, descriptions, and
	// values of the special request parameters that the SetQueueAttributes action
	// uses:
	//
	//     * DelaySeconds – The length of time, in seconds, for which the
	// delivery of all messages in the queue is delayed. Valid values: An integer from
	// 0 to 900 (15 minutes). Default: 0.
	//
	//     * MaximumMessageSize – The limit of how
	// many bytes a message can contain before Amazon SQS rejects it. Valid values: An
	// integer from 1,024 bytes (1 KiB) up to 262,144 bytes (256 KiB). Default: 262,144
	// (256 KiB).
	//
	//     * MessageRetentionPeriod – The length of time, in seconds, for
	// which Amazon SQS retains a message. Valid values: An integer representing
	// seconds, from 60 (1 minute) to 1,209,600 (14 days). Default: 345,600 (4 days).
	//
	//
	// * Policy – The queue's policy. A valid AWS policy. For more information about
	// policy structure, see Overview of AWS IAM Policies
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/PoliciesOverview.html) in the
	// Amazon IAM User Guide.
	//
	//     * ReceiveMessageWaitTimeSeconds – The length of
	// time, in seconds, for which a ReceiveMessage action waits for a message to
	// arrive. Valid values: An integer from 0 to 20 (seconds). Default: 0.
	//
	//     *
	// RedrivePolicy – The string that includes the parameters for the dead-letter
	// queue functionality of the source queue as a JSON object. For more information
	// about the redrive policy and dead-letter queues, see Using Amazon SQS
	// Dead-Letter Queues
	// (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-dead-letter-queues.html)
	// in the Amazon Simple Queue Service Developer Guide.
	//
	//         *
	// deadLetterTargetArn – The Amazon Resource Name (ARN) of the dead-letter queue to
	// which Amazon SQS moves messages after the value of maxReceiveCount is
	// exceeded.
	//
	//         * maxReceiveCount – The number of times a message is
	// delivered to the source queue before being moved to the dead-letter queue. When
	// the ReceiveCount for a message exceeds the maxReceiveCount for a queue, Amazon
	// SQS moves the message to the dead-letter-queue.
	//
	//     The dead-letter queue of a
	// FIFO queue must also be a FIFO queue. Similarly, the dead-letter queue of a
	// standard queue must also be a standard queue.
	//
	//     * VisibilityTimeout – The
	// visibility timeout for the queue, in seconds. Valid values: An integer from 0 to
	// 43,200 (12 hours). Default: 30. For more information about the visibility
	// timeout, see Visibility Timeout
	// (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-visibility-timeout.html)
	// in the Amazon Simple Queue Service Developer Guide.
	//
	// The following attributes
	// apply only to server-side-encryption
	// (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html):
	//
	//
	// * KmsMasterKeyId – The ID of an AWS-managed customer master key (CMK) for Amazon
	// SQS or a custom CMK. For more information, see Key Terms
	// (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html#sqs-sse-key-terms).
	// While the alias of the AWS-managed CMK for Amazon SQS is always alias/aws/sqs,
	// the alias of a custom CMK can, for example, be alias/MyAlias . For more
	// examples, see KeyId
	// (https://docs.aws.amazon.com/kms/latest/APIReference/API_DescribeKey.html#API_DescribeKey_RequestParameters)
	// in the AWS Key Management Service API Reference.
	//
	//     *
	// KmsDataKeyReusePeriodSeconds – The length of time, in seconds, for which Amazon
	// SQS can reuse a data key
	// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#data-keys)
	// to encrypt or decrypt messages before calling AWS KMS again. An integer
	// representing seconds, between 60 seconds (1 minute) and 86,400 seconds (24
	// hours). Default: 300 (5 minutes). A shorter time period provides better security
	// but results in more calls to KMS which might incur charges after Free Tier. For
	// more information, see How Does the Data Key Reuse Period Work?
	// (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html#sqs-how-does-the-data-key-reuse-period-work).
	//
	// The
	// following attribute applies only to FIFO (first-in-first-out) queues
	// (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html):
	//
	//
	// * ContentBasedDeduplication – Enables content-based deduplication. For more
	// information, see Exactly-Once Processing
	// (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html#FIFO-queues-exactly-once-processing)
	// in the Amazon Simple Queue Service Developer Guide.
	//
	//         * Every message
	// must have a unique MessageDeduplicationId,
	//
	//             * You may provide a
	// MessageDeduplicationId explicitly.
	//
	//             * If you aren't able to provide
	// a MessageDeduplicationId and you enable ContentBasedDeduplication for your
	// queue, Amazon SQS uses a SHA-256 hash to generate the MessageDeduplicationId
	// using the body of the message (but not the attributes of the message).
	//
	//
	// * If you don't provide a MessageDeduplicationId and the queue doesn't have
	// ContentBasedDeduplication set, the action fails with an error.
	//
	//             * If
	// the queue has ContentBasedDeduplication set, your MessageDeduplicationId
	// overrides the generated one.
	//
	//         * When ContentBasedDeduplication is in
	// effect, messages with identical content sent within the deduplication interval
	// are treated as duplicates and only one copy of the message is delivered.
	//
	//
	// * If you send one message with ContentBasedDeduplication enabled and then
	// another message with a MessageDeduplicationId that is the same as the one
	// generated for the first MessageDeduplicationId, the two messages are treated as
	// duplicates and only one copy of the message is delivered.
	//
	// This member is required.
	Attributes map[string]*string

	// The URL of the Amazon SQS queue whose attributes are set. Queue URLs and names
	// are case-sensitive.
	//
	// This member is required.
	QueueUrl *string
}

type SetQueueAttributesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationSetQueueAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsquery_serializeOpSetQueueAttributes{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsquery_deserializeOpSetQueueAttributes{})
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
	if err := addOpSetQueueAttributesValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opSetQueueAttributes(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opSetQueueAttributes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sqs",
		OperationName: "SetQueueAttributes",
	}
}
