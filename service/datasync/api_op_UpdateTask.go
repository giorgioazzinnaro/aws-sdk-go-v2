// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the metadata associated with a task.
func (c *Client) UpdateTask(ctx context.Context, params *UpdateTaskInput, optFns ...func(*Options)) (*UpdateTaskOutput, error) {
	if params == nil {
		params = &UpdateTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateTask", params, optFns, addOperationUpdateTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// UpdateTaskResponse
type UpdateTaskInput struct {

	// The Amazon Resource Name (ARN) of the resource name of the task to update.
	//
	// This member is required.
	TaskArn *string

	// The Amazon Resource Name (ARN) of the resource name of the CloudWatch LogGroup.
	CloudWatchLogGroupArn *string

	// A list of filter rules that determines which files to exclude from a task. The
	// list should contain a single filter string that consists of the patterns to
	// exclude. The patterns are delimited by "|" (that is, a pipe), for example:
	// "/folder1|/folder2"
	Excludes []*types.FilterRule

	// The name of the task to update.
	Name *string

	// Represents the options that are available to control the behavior of a
	// StartTaskExecution operation. Behavior includes preserving metadata such as user
	// ID (UID), group ID (GID), and file permissions, and also overwriting files in
	// the destination, data integrity verification, and so on. A task has a set of
	// default options associated with it. If you don't specify an option in
	// StartTaskExecution, the default value is used. You can override the defaults
	// options on each task execution by specifying an overriding Options value to
	// StartTaskExecution.
	Options *types.Options

	// Specifies a schedule used to periodically transfer files from a source to a
	// destination location. You can configure your task to execute hourly, daily,
	// weekly or on specific days of the week. You control when in the day or hour you
	// want the task to execute. The time you specify is UTC time. For more
	// information, see task-scheduling.
	Schedule *types.TaskSchedule
}

type UpdateTaskOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpUpdateTask{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpUpdateTask{})
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
	if err := addOpUpdateTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opUpdateTask(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opUpdateTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datasync",
		OperationName: "UpdateTask",
	}
}
