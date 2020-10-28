// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates a channel.
func (c *Client) UpdateChannel(ctx context.Context, params *UpdateChannelInput, optFns ...func(*Options)) (*UpdateChannelOutput, error) {
	if params == nil {
		params = &UpdateChannelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateChannel", params, optFns, addOperationUpdateChannelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateChannelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to update a channel.
type UpdateChannelInput struct {

	// channel ID
	//
	// This member is required.
	ChannelId *string

	// Specification of CDI inputs for this channel
	CdiInputSpecification *types.CdiInputSpecification

	// A list of output destinations for this channel.
	Destinations []*types.OutputDestination

	// The encoder settings for this channel.
	EncoderSettings *types.EncoderSettings

	// Placeholder documentation for __listOfInputAttachment
	InputAttachments []*types.InputAttachment

	// Specification of network and file inputs for this channel
	InputSpecification *types.InputSpecification

	// The log level to write to CloudWatch Logs.
	LogLevel types.LogLevel

	// The name of the channel.
	Name *string

	// An optional Amazon Resource Name (ARN) of the role to assume when running the
	// Channel. If you do not specify this on an update call but the role was
	// previously set that role will be removed.
	RoleArn *string
}

// Placeholder documentation for UpdateChannelResponse
type UpdateChannelOutput struct {

	// Placeholder documentation for Channel
	Channel *types.Channel

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateChannelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpUpdateChannel{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpUpdateChannel{})
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
	if err := addOpUpdateChannelValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opUpdateChannel(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opUpdateChannel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "medialive",
		OperationName: "UpdateChannel",
	}
}
