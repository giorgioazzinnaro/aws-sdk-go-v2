// Code generated by smithy-go-codegen DO NOT EDIT.

package devicefarm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about the configuration settings for your Amazon Virtual
// Private Cloud (VPC) endpoint.
func (c *Client) GetVPCEConfiguration(ctx context.Context, params *GetVPCEConfigurationInput, optFns ...func(*Options)) (*GetVPCEConfigurationOutput, error) {
	if params == nil {
		params = &GetVPCEConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetVPCEConfiguration", params, optFns, addOperationGetVPCEConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetVPCEConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetVPCEConfigurationInput struct {

	// The Amazon Resource Name (ARN) of the VPC endpoint configuration you want to
	// describe.
	//
	// This member is required.
	Arn *string
}

type GetVPCEConfigurationOutput struct {

	// An object that contains information about your VPC endpoint configuration.
	VpceConfiguration *types.VPCEConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetVPCEConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpGetVPCEConfiguration{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpGetVPCEConfiguration{})
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
	if err := addOpGetVPCEConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opGetVPCEConfiguration(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opGetVPCEConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devicefarm",
		OperationName: "GetVPCEConfiguration",
	}
}
