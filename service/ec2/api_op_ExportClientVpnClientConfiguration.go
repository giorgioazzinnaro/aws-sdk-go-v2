// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Downloads the contents of the Client VPN endpoint configuration file for the
// specified Client VPN endpoint. The Client VPN endpoint configuration file
// includes the Client VPN endpoint and certificate information clients need to
// establish a connection with the Client VPN endpoint.
func (c *Client) ExportClientVpnClientConfiguration(ctx context.Context, params *ExportClientVpnClientConfigurationInput, optFns ...func(*Options)) (*ExportClientVpnClientConfigurationOutput, error) {
	if params == nil {
		params = &ExportClientVpnClientConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ExportClientVpnClientConfiguration", params, optFns, addOperationExportClientVpnClientConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ExportClientVpnClientConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ExportClientVpnClientConfigurationInput struct {

	// The ID of the Client VPN endpoint.
	//
	// This member is required.
	ClientVpnEndpointId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type ExportClientVpnClientConfigurationOutput struct {

	// The contents of the Client VPN endpoint configuration file.
	ClientConfiguration *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationExportClientVpnClientConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpExportClientVpnClientConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpExportClientVpnClientConfiguration{}, middleware.After)
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
	addOpExportClientVpnClientConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opExportClientVpnClientConfiguration(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opExportClientVpnClientConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ExportClientVpnClientConfiguration",
	}
}
