// Code generated by smithy-go-codegen DO NOT EDIT.

package sns

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Sets the attributes for an endpoint for a device on one of the supported push
// notification services, such as GCM (Firebase Cloud Messaging) and APNS. For more
// information, see Using Amazon SNS Mobile Push Notifications
// (https://docs.aws.amazon.com/sns/latest/dg/SNSMobilePush.html).
func (c *Client) SetEndpointAttributes(ctx context.Context, params *SetEndpointAttributesInput, optFns ...func(*Options)) (*SetEndpointAttributesOutput, error) {
	if params == nil {
		params = &SetEndpointAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetEndpointAttributes", params, optFns, addOperationSetEndpointAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetEndpointAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input for SetEndpointAttributes action.
type SetEndpointAttributesInput struct {

	// A map of the endpoint attributes. Attributes in this map include the
	// following:
	//
	// * CustomUserData – arbitrary user data to associate with the
	// endpoint. Amazon SNS does not use this data. The data must be in UTF-8 format
	// and less than 2KB.
	//
	// * Enabled – flag that enables/disables delivery to the
	// endpoint. Amazon SNS will set this to false when a notification service
	// indicates to Amazon SNS that the endpoint is invalid. Users can set it back to
	// true, typically after updating Token.
	//
	// * Token – device token, also referred to
	// as a registration id, for an app and mobile device. This is returned from the
	// notification service when an app and mobile device are registered with the
	// notification service.
	//
	// This member is required.
	Attributes map[string]*string

	// EndpointArn used for SetEndpointAttributes action.
	//
	// This member is required.
	EndpointArn *string
}

type SetEndpointAttributesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationSetEndpointAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpSetEndpointAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpSetEndpointAttributes{}, middleware.After)
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
	addOpSetEndpointAttributesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSetEndpointAttributes(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opSetEndpointAttributes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sns",
		OperationName: "SetEndpointAttributes",
	}
}
