// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworks

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/opsworks/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Registers instances that were created outside of AWS OpsWorks Stacks with a
// specified stack. We do not recommend using this action to register instances.
// The complete registration operation includes two tasks: installing the AWS
// OpsWorks Stacks agent on the instance, and registering the instance with the
// stack. RegisterInstance handles only the second step. You should instead use the
// AWS CLI register command, which performs the entire registration operation. For
// more information, see  Registering an Instance with an AWS OpsWorks Stacks Stack
// (https://docs.aws.amazon.com/opsworks/latest/userguide/registered-instances-register.html).
// Registered instances have the same requirements as instances that are created by
// using the CreateInstance API. For example, registered instances must be running
// a supported Linux-based operating system, and they must have a supported
// instance type. For more information about requirements for instances that you
// want to register, see  Preparing the Instance
// (https://docs.aws.amazon.com/opsworks/latest/userguide/registered-instances-register-registering-preparer.html).
// Required Permissions: To use this action, an IAM user must have a Manage
// permissions level for the stack or an attached policy that explicitly grants
// permissions. For more information on user permissions, see Managing User
// Permissions
// (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
func (c *Client) RegisterInstance(ctx context.Context, params *RegisterInstanceInput, optFns ...func(*Options)) (*RegisterInstanceOutput, error) {
	if params == nil {
		params = &RegisterInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterInstance", params, optFns, addOperationRegisterInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterInstanceInput struct {

	// The ID of the stack that the instance is to be registered with.
	//
	// This member is required.
	StackId *string

	// The instance's hostname.
	Hostname *string

	// An InstanceIdentity object that contains the instance's identity.
	InstanceIdentity *types.InstanceIdentity

	// The instance's private IP address.
	PrivateIp *string

	// The instance's public IP address.
	PublicIp *string

	// The instances public RSA key. This key is used to encrypt communication between
	// the instance and the service.
	RsaPublicKey *string

	// The instances public RSA key fingerprint.
	RsaPublicKeyFingerprint *string
}

// Contains the response to a RegisterInstanceResult request.
type RegisterInstanceOutput struct {

	// The registered instance's AWS OpsWorks Stacks ID.
	InstanceId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRegisterInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpRegisterInstance{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpRegisterInstance{})
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
	if err := addOpRegisterInstanceValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opRegisterInstance(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opRegisterInstance(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks",
		OperationName: "RegisterInstance",
	}
}
