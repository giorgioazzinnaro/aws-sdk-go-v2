// Code generated by smithy-go-codegen DO NOT EDIT.

package macie2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/macie2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Changes the status of a classification job.
func (c *Client) UpdateClassificationJob(ctx context.Context, params *UpdateClassificationJobInput, optFns ...func(*Options)) (*UpdateClassificationJobOutput, error) {
	if params == nil {
		params = &UpdateClassificationJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateClassificationJob", params, optFns, addOperationUpdateClassificationJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateClassificationJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateClassificationJobInput struct {

	// The unique identifier for the classification job.
	//
	// This member is required.
	JobId *string

	// The new status for the job. Valid values are:
	//
	//     * CANCELLED - Stops the job
	// permanently and cancels it. You can't resume a job after you cancel it. This
	// value is valid only if the job's current status is IDLE, PAUSED, RUNNING, or
	// USER_PAUSED.
	//
	//     * RUNNING - Resumes the job. This value is valid only if the
	// job's current status is USER_PAUSED. If you specify this value, Amazon Macie
	// immediately resumes the job.
	//
	//     * USER_PAUSED - Pauses the job. This value is
	// valid only if the job's current status is IDLE or RUNNING. If you specify this
	// value and the job is currently running, Macie immediately stops running the job.
	// To resume a job after you pause it, change the job's status to RUNNING. If you
	// don't resume a job within 30 days of pausing it, the job expires and Macie
	// cancels it. You can't resume a job after it's cancelled.
	//
	// This member is required.
	JobStatus types.JobStatus
}

type UpdateClassificationJobOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateClassificationJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpUpdateClassificationJob{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpUpdateClassificationJob{})
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
	if err := addOpUpdateClassificationJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opUpdateClassificationJob(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opUpdateClassificationJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "macie2",
		OperationName: "UpdateClassificationJob",
	}
}
