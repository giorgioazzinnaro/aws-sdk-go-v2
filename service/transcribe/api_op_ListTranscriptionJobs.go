// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribe

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/transcribe/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists transcription jobs with the specified status.
func (c *Client) ListTranscriptionJobs(ctx context.Context, params *ListTranscriptionJobsInput, optFns ...func(*Options)) (*ListTranscriptionJobsOutput, error) {
	if params == nil {
		params = &ListTranscriptionJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTranscriptionJobs", params, optFns, addOperationListTranscriptionJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTranscriptionJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTranscriptionJobsInput struct {

	// When specified, the jobs returned in the list are limited to jobs whose name
	// contains the specified string.
	JobNameContains *string

	// The maximum number of jobs to return in the response. If there are fewer results
	// in the list, this response contains only the actual results.
	MaxResults *int32

	// If the result of the previous request to ListTranscriptionJobs was truncated,
	// include the NextToken to fetch the next set of jobs.
	NextToken *string

	// When specified, returns only transcription jobs with the specified status. Jobs
	// are ordered by creation date, with the newest jobs returned first. If you don’t
	// specify a status, Amazon Transcribe returns all transcription jobs ordered by
	// creation date.
	Status types.TranscriptionJobStatus
}

type ListTranscriptionJobsOutput struct {

	// The ListTranscriptionJobs operation returns a page of jobs at a time. The
	// maximum size of the page is set by the MaxResults parameter. If there are more
	// jobs in the list than the page size, Amazon Transcribe returns the NextPage
	// token. Include the token in the next request to the ListTranscriptionJobs
	// operation to return in the next page of jobs.
	NextToken *string

	// The requested status of the jobs returned.
	Status types.TranscriptionJobStatus

	// A list of objects containing summary information for a transcription job.
	TranscriptionJobSummaries []*types.TranscriptionJobSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListTranscriptionJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpListTranscriptionJobs{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpListTranscriptionJobs{})
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
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opListTranscriptionJobs(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opListTranscriptionJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transcribe",
		OperationName: "ListTranscriptionJobs",
	}
}
