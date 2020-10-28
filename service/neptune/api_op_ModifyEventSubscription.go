// Code generated by smithy-go-codegen DO NOT EDIT.

package neptune

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modifies an existing event notification subscription. Note that you can't modify
// the source identifiers using this call; to change source identifiers for a
// subscription, use the AddSourceIdentifierToSubscription and
// RemoveSourceIdentifierFromSubscription calls. You can see a list of the event
// categories for a given SourceType by using the DescribeEventCategories action.
func (c *Client) ModifyEventSubscription(ctx context.Context, params *ModifyEventSubscriptionInput, optFns ...func(*Options)) (*ModifyEventSubscriptionOutput, error) {
	if params == nil {
		params = &ModifyEventSubscriptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyEventSubscription", params, optFns, addOperationModifyEventSubscriptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyEventSubscriptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyEventSubscriptionInput struct {

	// The name of the event notification subscription.
	//
	// This member is required.
	SubscriptionName *string

	// A Boolean value; set to true to activate the subscription.
	Enabled *bool

	// A list of event categories for a SourceType that you want to subscribe to. You
	// can see a list of the categories for a given SourceType by using the
	// DescribeEventCategories action.
	EventCategories []*string

	// The Amazon Resource Name (ARN) of the SNS topic created for event notification.
	// The ARN is created by Amazon SNS when you create a topic and subscribe to it.
	SnsTopicArn *string

	// The type of source that is generating the events. For example, if you want to be
	// notified of events generated by a DB instance, you would set this parameter to
	// db-instance. if this value is not specified, all events are returned. Valid
	// values: db-instance | db-parameter-group | db-security-group | db-snapshot
	SourceType *string
}

type ModifyEventSubscriptionOutput struct {

	// Contains the results of a successful invocation of the
	// DescribeEventSubscriptions action.
	EventSubscription *types.EventSubscription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationModifyEventSubscriptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsquery_serializeOpModifyEventSubscription{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsquery_deserializeOpModifyEventSubscription{})
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
	if err := addOpModifyEventSubscriptionValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opModifyEventSubscription(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opModifyEventSubscription(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "ModifyEventSubscription",
	}
}
