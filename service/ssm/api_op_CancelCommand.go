// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Attempts to cancel the command specified by the Command ID. There is no
// guarantee that the command will be terminated and the underlying process
// stopped.
func (c *Client) CancelCommand(ctx context.Context, params *CancelCommandInput, optFns ...func(*Options)) (*CancelCommandOutput, error) {
	if params == nil {
		params = &CancelCommandInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CancelCommand", params, optFns, addOperationCancelCommandMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CancelCommandOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type CancelCommandInput struct {

	// The ID of the command you want to cancel.
	//
	// This member is required.
	CommandId *string

	// (Optional) A list of instance IDs on which you want to cancel the command. If
	// not provided, the command is canceled on every instance on which it was
	// requested.
	InstanceIds []*string
}

// Whether or not the command was successfully canceled. There is no guarantee that
// a request can be canceled.
type CancelCommandOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCancelCommandMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCancelCommand{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCancelCommand{}, middleware.After)
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
	addOpCancelCommandValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCancelCommand(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCancelCommand(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "CancelCommand",
	}
}
