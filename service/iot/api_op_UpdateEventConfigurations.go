// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the event configurations.
func (c *Client) UpdateEventConfigurations(ctx context.Context, params *UpdateEventConfigurationsInput, optFns ...func(*Options)) (*UpdateEventConfigurationsOutput, error) {
	if params == nil {
		params = &UpdateEventConfigurationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateEventConfigurations", params, optFns, addOperationUpdateEventConfigurationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateEventConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateEventConfigurationsInput struct {

	// The new event configuration values.
	EventConfigurations map[string]*types.Configuration
}

type UpdateEventConfigurationsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateEventConfigurationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpUpdateEventConfigurations{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpUpdateEventConfigurations{})
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
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opUpdateEventConfigurations(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opUpdateEventConfigurations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "UpdateEventConfigurations",
	}
}
