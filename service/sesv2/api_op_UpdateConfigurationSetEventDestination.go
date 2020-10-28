// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Update the configuration of an event destination for a configuration set. Events
// include message sends, deliveries, opens, clicks, bounces, and complaints. Event
// destinations are places that you can send information about these events to. For
// example, you can send event data to Amazon SNS to receive notifications when you
// receive bounces or complaints, or you can use Amazon Kinesis Data Firehose to
// stream data to Amazon S3 for long-term storage.
func (c *Client) UpdateConfigurationSetEventDestination(ctx context.Context, params *UpdateConfigurationSetEventDestinationInput, optFns ...func(*Options)) (*UpdateConfigurationSetEventDestinationOutput, error) {
	if params == nil {
		params = &UpdateConfigurationSetEventDestinationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateConfigurationSetEventDestination", params, optFns, addOperationUpdateConfigurationSetEventDestinationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateConfigurationSetEventDestinationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to change the settings for an event destination for a configuration
// set.
type UpdateConfigurationSetEventDestinationInput struct {

	// The name of the configuration set that contains the event destination that you
	// want to modify.
	//
	// This member is required.
	ConfigurationSetName *string

	// An object that defines the event destination.
	//
	// This member is required.
	EventDestination *types.EventDestinationDefinition

	// The name of the event destination that you want to modify.
	//
	// This member is required.
	EventDestinationName *string
}

// An HTTP 200 response if the request succeeds, or an error message if the request
// fails.
type UpdateConfigurationSetEventDestinationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateConfigurationSetEventDestinationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpUpdateConfigurationSetEventDestination{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpUpdateConfigurationSetEventDestination{})
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
	if err := addOpUpdateConfigurationSetEventDestinationValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opUpdateConfigurationSetEventDestination(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opUpdateConfigurationSetEventDestination(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "UpdateConfigurationSetEventDestination",
	}
}
