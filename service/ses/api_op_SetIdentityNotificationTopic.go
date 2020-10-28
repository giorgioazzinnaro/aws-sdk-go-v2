// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Sets an Amazon Simple Notification Service (Amazon SNS) topic to use when
// delivering notifications. When you use this operation, you specify a verified
// identity, such as an email address or domain. When you send an email that uses
// the chosen identity in the Source field, Amazon SES sends notifications to the
// topic you specified. You can send bounce, complaint, or delivery notifications
// (or any combination of the three) to the Amazon SNS topic that you specify. You
// can execute this operation no more than once per second. For more information
// about feedback notification, see the Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/notifications.html).
func (c *Client) SetIdentityNotificationTopic(ctx context.Context, params *SetIdentityNotificationTopicInput, optFns ...func(*Options)) (*SetIdentityNotificationTopicOutput, error) {
	if params == nil {
		params = &SetIdentityNotificationTopicInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetIdentityNotificationTopic", params, optFns, addOperationSetIdentityNotificationTopicMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetIdentityNotificationTopicOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to specify the Amazon SNS topic to which Amazon SES will
// publish bounce, complaint, or delivery notifications for emails sent with that
// identity as the Source. For information about Amazon SES notifications, see the
// Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/notifications-via-sns.html).
type SetIdentityNotificationTopicInput struct {

	// The identity (email address or domain) that you want to set the Amazon SNS topic
	// for. You can only specify a verified identity for this parameter. You can
	// specify an identity by using its name or by using its Amazon Resource Name
	// (ARN). The following examples are all valid identities: sender@example.com,
	// example.com, arn:aws:ses:us-east-1:123456789012:identity/example.com.
	//
	// This member is required.
	Identity *string

	// The type of notifications that will be published to the specified Amazon SNS
	// topic.
	//
	// This member is required.
	NotificationType types.NotificationType

	// The Amazon Resource Name (ARN) of the Amazon SNS topic. If the parameter is
	// omitted from the request or a null value is passed, SnsTopic is cleared and
	// publishing is disabled.
	SnsTopic *string
}

// An empty element returned on a successful request.
type SetIdentityNotificationTopicOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationSetIdentityNotificationTopicMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsquery_serializeOpSetIdentityNotificationTopic{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsquery_deserializeOpSetIdentityNotificationTopic{})
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
	if err := addOpSetIdentityNotificationTopicValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opSetIdentityNotificationTopic(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opSetIdentityNotificationTopic(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "SetIdentityNotificationTopic",
	}
}
