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

// Adds termination settings for the specified Amazon Chime Voice Connector. If
// emergency calling is configured for the Amazon Chime Voice Connector, it must be
// deleted prior to turning off termination settings.
func (c *Client) PutVoiceConnectorTermination(ctx context.Context, params *PutVoiceConnectorTerminationInput, optFns ...func(*Options)) (*PutVoiceConnectorTerminationOutput, error) {
	if params == nil {
		params = &PutVoiceConnectorTerminationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutVoiceConnectorTermination", params, optFns, addOperationPutVoiceConnectorTerminationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutVoiceConnectorTerminationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutVoiceConnectorTerminationInput struct {

	// The termination setting details to add.
	//
	// This member is required.
	Termination *types.Termination

	// The Amazon Chime Voice Connector ID.
	//
	// This member is required.
	VoiceConnectorId *string
}

type PutVoiceConnectorTerminationOutput struct {

	// The updated termination setting details.
	Termination *types.Termination

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutVoiceConnectorTerminationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpPutVoiceConnectorTermination{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpPutVoiceConnectorTermination{})
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
	if err := addOpPutVoiceConnectorTerminationValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opPutVoiceConnectorTermination(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opPutVoiceConnectorTermination(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "PutVoiceConnectorTermination",
	}
}
