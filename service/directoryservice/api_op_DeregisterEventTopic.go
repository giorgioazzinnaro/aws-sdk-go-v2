// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes the specified directory as a publisher to the specified SNS topic.
func (c *Client) DeregisterEventTopic(ctx context.Context, params *DeregisterEventTopicInput, optFns ...func(*Options)) (*DeregisterEventTopicOutput, error) {
	if params == nil {
		params = &DeregisterEventTopicInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeregisterEventTopic", params, optFns, addOperationDeregisterEventTopicMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeregisterEventTopicOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Removes the specified directory as a publisher to the specified SNS topic.
type DeregisterEventTopicInput struct {

	// The Directory ID to remove as a publisher. This directory will no longer send
	// messages to the specified SNS topic.
	//
	// This member is required.
	DirectoryId *string

	// The name of the SNS topic from which to remove the directory as a publisher.
	//
	// This member is required.
	TopicName *string
}

// The result of a DeregisterEventTopic request.
type DeregisterEventTopicOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeregisterEventTopicMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpDeregisterEventTopic{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpDeregisterEventTopic{})
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
	if err := addOpDeregisterEventTopicValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDeregisterEventTopic(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDeregisterEventTopic(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "DeregisterEventTopic",
	}
}
