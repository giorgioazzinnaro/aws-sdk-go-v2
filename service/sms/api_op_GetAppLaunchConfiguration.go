// Code generated by smithy-go-codegen DO NOT EDIT.

package sms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sms/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the application launch configuration associated with the specified
// application.
func (c *Client) GetAppLaunchConfiguration(ctx context.Context, params *GetAppLaunchConfigurationInput, optFns ...func(*Options)) (*GetAppLaunchConfigurationOutput, error) {
	if params == nil {
		params = &GetAppLaunchConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAppLaunchConfiguration", params, optFns, addOperationGetAppLaunchConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAppLaunchConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAppLaunchConfigurationInput struct {

	// The ID of the application.
	AppId *string
}

type GetAppLaunchConfigurationOutput struct {

	// The ID of the application.
	AppId *string

	// Indicates whether the application is configured to launch automatically after
	// replication is complete.
	AutoLaunch *bool

	// The name of the service role in the customer's account that AWS CloudFormation
	// uses to launch the application.
	RoleName *string

	// The launch configurations for server groups in this application.
	ServerGroupLaunchConfigurations []*types.ServerGroupLaunchConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetAppLaunchConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpGetAppLaunchConfiguration{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpGetAppLaunchConfiguration{})
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
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opGetAppLaunchConfiguration(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opGetAppLaunchConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sms",
		OperationName: "GetAppLaunchConfiguration",
	}
}
