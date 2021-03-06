// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Modifies the start time for a scheduled Amazon EC2 instance event.
func (c *Client) ModifyInstanceEventStartTime(ctx context.Context, params *ModifyInstanceEventStartTimeInput, optFns ...func(*Options)) (*ModifyInstanceEventStartTimeOutput, error) {
	if params == nil {
		params = &ModifyInstanceEventStartTimeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyInstanceEventStartTime", params, optFns, addOperationModifyInstanceEventStartTimeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyInstanceEventStartTimeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyInstanceEventStartTimeInput struct {

	// The ID of the event whose date and time you are modifying.
	//
	// This member is required.
	InstanceEventId *string

	// The ID of the instance with the scheduled event.
	//
	// This member is required.
	InstanceId *string

	// The new date and time when the event will take place.
	//
	// This member is required.
	NotBefore *time.Time

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type ModifyInstanceEventStartTimeOutput struct {

	// Describes a scheduled event for an instance.
	Event *types.InstanceStatusEvent

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationModifyInstanceEventStartTimeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpModifyInstanceEventStartTime{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpModifyInstanceEventStartTime{}, middleware.After)
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
	addOpModifyInstanceEventStartTimeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opModifyInstanceEventStartTime(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opModifyInstanceEventStartTime(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ModifyInstanceEventStartTime",
	}
}
