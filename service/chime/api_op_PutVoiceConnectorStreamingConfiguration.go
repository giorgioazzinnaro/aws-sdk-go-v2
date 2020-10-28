// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds a streaming configuration for the specified Amazon Chime Voice Connector.
// The streaming configuration specifies whether media streaming is enabled for
// sending to Amazon Kinesis. It also sets the retention period, in hours, for the
// Amazon Kinesis data.
func (c *Client) PutVoiceConnectorStreamingConfiguration(ctx context.Context, params *PutVoiceConnectorStreamingConfigurationInput, optFns ...func(*Options)) (*PutVoiceConnectorStreamingConfigurationOutput, error) {
	if params == nil {
		params = &PutVoiceConnectorStreamingConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutVoiceConnectorStreamingConfiguration", params, optFns, addOperationPutVoiceConnectorStreamingConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutVoiceConnectorStreamingConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutVoiceConnectorStreamingConfigurationInput struct {

	// The streaming configuration details to add.
	//
	// This member is required.
	StreamingConfiguration *types.StreamingConfiguration

	// The Amazon Chime Voice Connector ID.
	//
	// This member is required.
	VoiceConnectorId *string
}

type PutVoiceConnectorStreamingConfigurationOutput struct {

	// The updated streaming configuration details.
	StreamingConfiguration *types.StreamingConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutVoiceConnectorStreamingConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpPutVoiceConnectorStreamingConfiguration{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpPutVoiceConnectorStreamingConfiguration{})
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
	if err := addOpPutVoiceConnectorStreamingConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opPutVoiceConnectorStreamingConfiguration(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opPutVoiceConnectorStreamingConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "PutVoiceConnectorStreamingConfiguration",
	}
}
