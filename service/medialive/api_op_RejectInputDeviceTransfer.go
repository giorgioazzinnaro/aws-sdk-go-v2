// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Reject the transfer of the specified input device to your AWS account.
func (c *Client) RejectInputDeviceTransfer(ctx context.Context, params *RejectInputDeviceTransferInput, optFns ...func(*Options)) (*RejectInputDeviceTransferOutput, error) {
	if params == nil {
		params = &RejectInputDeviceTransferInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RejectInputDeviceTransfer", params, optFns, addOperationRejectInputDeviceTransferMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RejectInputDeviceTransferOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Placeholder documentation for RejectInputDeviceTransferRequest
type RejectInputDeviceTransferInput struct {

	// The unique ID of the input device to reject. For example, hd-123456789abcdef.
	//
	// This member is required.
	InputDeviceId *string
}

// Placeholder documentation for RejectInputDeviceTransferResponse
type RejectInputDeviceTransferOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRejectInputDeviceTransferMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpRejectInputDeviceTransfer{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpRejectInputDeviceTransfer{})
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
	if err := addOpRejectInputDeviceTransferValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opRejectInputDeviceTransfer(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opRejectInputDeviceTransfer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "medialive",
		OperationName: "RejectInputDeviceTransfer",
	}
}
