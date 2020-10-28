// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new virtual MFA device for the AWS account. After creating the virtual
// MFA, use EnableMFADevice to attach the MFA device to an IAM user. For more
// information about creating and working with virtual MFA devices, go to Using a
// Virtual MFA Device
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_VirtualMFA.html) in the
// IAM User Guide. The number and size of IAM resources in an AWS account are
// limited. For more information, see IAM and STS Quotas
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html) in
// the IAM User Guide. The seed information contained in the QR code and the Base32
// string should be treated like any other secret access information. In other
// words, protect the seed information as you would your AWS access keys or your
// passwords. After you provision your virtual device, you should ensure that the
// information is destroyed following secure procedures.
func (c *Client) CreateVirtualMFADevice(ctx context.Context, params *CreateVirtualMFADeviceInput, optFns ...func(*Options)) (*CreateVirtualMFADeviceOutput, error) {
	if params == nil {
		params = &CreateVirtualMFADeviceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateVirtualMFADevice", params, optFns, addOperationCreateVirtualMFADeviceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateVirtualMFADeviceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateVirtualMFADeviceInput struct {

	// The name of the virtual MFA device. Use with path to uniquely identify a virtual
	// MFA device. This parameter allows (through its regex pattern
	// (http://wikipedia.org/wiki/regex)) a string of characters consisting of upper
	// and lowercase alphanumeric characters with no spaces. You can also include any
	// of the following characters: _+=,.@-
	//
	// This member is required.
	VirtualMFADeviceName *string

	// The path for the virtual MFA device. For more information about paths, see IAM
	// Identifiers
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the
	// IAM User Guide. This parameter is optional. If it is not included, it defaults
	// to a slash (/). This parameter allows (through its regex pattern
	// (http://wikipedia.org/wiki/regex)) a string of characters consisting of either a
	// forward slash (/) by itself or a string that must begin and end with forward
	// slashes. In addition, it can contain any ASCII character from the ! (\u0021)
	// through the DEL character (\u007F), including most punctuation characters,
	// digits, and upper and lowercased letters.
	Path *string
}

// Contains the response to a successful CreateVirtualMFADevice request.
type CreateVirtualMFADeviceOutput struct {

	// A structure containing details about the new virtual MFA device.
	//
	// This member is required.
	VirtualMFADevice *types.VirtualMFADevice

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateVirtualMFADeviceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsquery_serializeOpCreateVirtualMFADevice{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsquery_deserializeOpCreateVirtualMFADevice{})
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
	if err := addOpCreateVirtualMFADeviceValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opCreateVirtualMFADevice(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opCreateVirtualMFADevice(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "CreateVirtualMFADevice",
	}
}
