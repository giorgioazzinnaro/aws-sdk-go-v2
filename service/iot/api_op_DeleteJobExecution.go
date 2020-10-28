// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a job execution.
func (c *Client) DeleteJobExecution(ctx context.Context, params *DeleteJobExecutionInput, optFns ...func(*Options)) (*DeleteJobExecutionOutput, error) {
	if params == nil {
		params = &DeleteJobExecutionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteJobExecution", params, optFns, addOperationDeleteJobExecutionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteJobExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteJobExecutionInput struct {

	// The ID of the job execution to be deleted. The executionNumber refers to the
	// execution of a particular job on a particular device. Note that once a job
	// execution is deleted, the executionNumber may be reused by IoT, so be sure you
	// get and use the correct value here.
	//
	// This member is required.
	ExecutionNumber *int64

	// The ID of the job whose execution on a particular device will be deleted.
	//
	// This member is required.
	JobId *string

	// The name of the thing whose job execution will be deleted.
	//
	// This member is required.
	ThingName *string

	// (Optional) When true, you can delete a job execution which is "IN_PROGRESS".
	// Otherwise, you can only delete a job execution which is in a terminal state
	// ("SUCCEEDED", "FAILED", "REJECTED", "REMOVED" or "CANCELED") or an exception
	// will occur. The default is false. Deleting a job execution which is
	// "IN_PROGRESS", will cause the device to be unable to access job information or
	// update the job execution status. Use caution and ensure that the device is able
	// to recover to a valid state.
	Force *bool
}

type DeleteJobExecutionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteJobExecutionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpDeleteJobExecution{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpDeleteJobExecution{})
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
	if err := addOpDeleteJobExecutionValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDeleteJobExecution(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDeleteJobExecution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "DeleteJobExecution",
	}
}
