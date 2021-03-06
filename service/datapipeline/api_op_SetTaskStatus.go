// Code generated by smithy-go-codegen DO NOT EDIT.

package datapipeline

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datapipeline/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Task runners call SetTaskStatus to notify AWS Data Pipeline that a task is
// completed and provide information about the final status. A task runner makes
// this call regardless of whether the task was sucessful. A task runner does not
// need to call SetTaskStatus for tasks that are canceled by the web service during
// a call to ReportTaskProgress.
func (c *Client) SetTaskStatus(ctx context.Context, params *SetTaskStatusInput, optFns ...func(*Options)) (*SetTaskStatusOutput, error) {
	if params == nil {
		params = &SetTaskStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetTaskStatus", params, optFns, addOperationSetTaskStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetTaskStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for SetTaskStatus.
type SetTaskStatusInput struct {

	// The ID of the task assigned to the task runner. This value is provided in the
	// response for PollForTask.
	//
	// This member is required.
	TaskId *string

	// If FINISHED, the task successfully completed. If FAILED, the task ended
	// unsuccessfully. Preconditions use false.
	//
	// This member is required.
	TaskStatus types.TaskStatus

	// If an error occurred during the task, this value specifies the error code. This
	// value is set on the physical attempt object. It is used to display error
	// information to the user. It should not start with string "Service_" which is
	// reserved by the system.
	ErrorId *string

	// If an error occurred during the task, this value specifies a text description of
	// the error. This value is set on the physical attempt object. It is used to
	// display error information to the user. The web service does not parse this
	// value.
	ErrorMessage *string

	// If an error occurred during the task, this value specifies the stack trace
	// associated with the error. This value is set on the physical attempt object. It
	// is used to display error information to the user. The web service does not parse
	// this value.
	ErrorStackTrace *string
}

// Contains the output of SetTaskStatus.
type SetTaskStatusOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationSetTaskStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSetTaskStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSetTaskStatus{}, middleware.After)
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
	addOpSetTaskStatusValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSetTaskStatus(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opSetTaskStatus(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datapipeline",
		OperationName: "SetTaskStatus",
	}
}
