// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Retrieves the encrypted administrator password for a running Windows instance.
// The Windows password is generated at boot by the EC2Config service or EC2Launch
// scripts (Windows Server 2016 and later). This usually only happens the first
// time an instance is launched. For more information, see EC2Config
// (https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/UsingConfig_WinAMI.html)
// and EC2Launch
// (https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/ec2launch.html) in the
// Amazon Elastic Compute Cloud User Guide. For the EC2Config service, the password
// is not generated for rebundled AMIs unless Ec2SetPassword is enabled before
// bundling. The password is encrypted using the key pair that you specified when
// you launched the instance. You must provide the corresponding key pair file.
// When you launch an instance, password generation and encryption may take a few
// minutes. If you try to retrieve the password before it's available, the output
// returns an empty string. We recommend that you wait up to 15 minutes after
// launching an instance before trying to retrieve the generated password.
func (c *Client) GetPasswordData(ctx context.Context, params *GetPasswordDataInput, optFns ...func(*Options)) (*GetPasswordDataOutput, error) {
	if params == nil {
		params = &GetPasswordDataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetPasswordData", params, optFns, addOperationGetPasswordDataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetPasswordDataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetPasswordDataInput struct {

	// The ID of the Windows instance.
	//
	// This member is required.
	InstanceId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type GetPasswordDataOutput struct {

	// The ID of the Windows instance.
	InstanceId *string

	// The password of the instance. Returns an empty string if the password is not
	// available.
	PasswordData *string

	// The time the data was last updated.
	Timestamp *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetPasswordDataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsEc2query_serializeOpGetPasswordData{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsEc2query_deserializeOpGetPasswordData{})
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
	if err := addOpGetPasswordDataValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opGetPasswordData(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opGetPasswordData(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "GetPasswordData",
	}
}
