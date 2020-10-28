// Code generated by smithy-go-codegen DO NOT EDIT.

package transfer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/transfer/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// If the IdentityProviderType of a file transfer protocol-enabled server is
// API_Gateway, tests whether your API Gateway is set up successfully. We highly
// recommend that you call this operation to test your authentication method as
// soon as you create your server. By doing so, you can troubleshoot issues with
// the API Gateway integration to ensure that your users can successfully use the
// service.
func (c *Client) TestIdentityProvider(ctx context.Context, params *TestIdentityProviderInput, optFns ...func(*Options)) (*TestIdentityProviderOutput, error) {
	if params == nil {
		params = &TestIdentityProviderInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "TestIdentityProvider", params, optFns, addOperationTestIdentityProviderMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*TestIdentityProviderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TestIdentityProviderInput struct {

	// A system-assigned identifier for a specific server. That server's user
	// authentication method is tested with a user name and password.
	//
	// This member is required.
	ServerId *string

	// The name of the user account to be tested.
	//
	// This member is required.
	UserName *string

	// The type of file transfer protocol to be tested. The available protocols are:
	//
	//
	// * Secure Shell (SSH) File Transfer Protocol (SFTP)
	//
	//     * File Transfer Protocol
	// Secure (FTPS)
	//
	//     * File Transfer Protocol (FTP)
	ServerProtocol types.Protocol

	// The source IP address of the user account to be tested.
	SourceIp *string

	// The password of the user account to be tested.
	UserPassword *string
}

type TestIdentityProviderOutput struct {

	// The HTTP status code that is the response from your API Gateway.
	//
	// This member is required.
	StatusCode *int32

	// The endpoint of the service used to authenticate a user.
	//
	// This member is required.
	Url *string

	// A message that indicates whether the test was successful or not.
	Message *string

	// The response that is returned from your API Gateway.
	Response *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationTestIdentityProviderMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpTestIdentityProvider{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpTestIdentityProvider{})
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
	if err := addOpTestIdentityProviderValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opTestIdentityProvider(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opTestIdentityProvider(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transfer",
		OperationName: "TestIdentityProvider",
	}
}
