// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehend

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/comprehend/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Stops a sentiment detection job in progress. If the job state is IN_PROGRESS the
// job is marked for termination and put into the STOP_REQUESTED state. If the job
// completes before it can be stopped, it is put into the COMPLETED state;
// otherwise the job is be stopped and put into the STOPPED state. If the job is in
// the COMPLETED or FAILED state when you call the StopDominantLanguageDetectionJob
// operation, the operation returns a 400 Internal Request Exception. When a job is
// stopped, any documents already processed are written to the output location.
func (c *Client) StopSentimentDetectionJob(ctx context.Context, params *StopSentimentDetectionJobInput, optFns ...func(*Options)) (*StopSentimentDetectionJobOutput, error) {
	if params == nil {
		params = &StopSentimentDetectionJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StopSentimentDetectionJob", params, optFns, addOperationStopSentimentDetectionJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StopSentimentDetectionJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopSentimentDetectionJobInput struct {

	// The identifier of the sentiment detection job to stop.
	//
	// This member is required.
	JobId *string
}

type StopSentimentDetectionJobOutput struct {

	// The identifier of the sentiment detection job to stop.
	JobId *string

	// Either STOP_REQUESTED if the job is currently running, or STOPPED if the job was
	// previously stopped with the StopSentimentDetectionJob operation.
	JobStatus types.JobStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStopSentimentDetectionJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStopSentimentDetectionJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStopSentimentDetectionJob{}, middleware.After)
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
	addOpStopSentimentDetectionJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStopSentimentDetectionJob(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opStopSentimentDetectionJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehend",
		OperationName: "StopSentimentDetectionJob",
	}
}
