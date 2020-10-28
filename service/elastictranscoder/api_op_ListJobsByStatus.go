// Code generated by smithy-go-codegen DO NOT EDIT.

package elastictranscoder

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The ListJobsByStatus operation gets a list of jobs that have a specified status.
// The response body contains one element for each job that satisfies the search
// criteria.
func (c *Client) ListJobsByStatus(ctx context.Context, params *ListJobsByStatusInput, optFns ...func(*Options)) (*ListJobsByStatusOutput, error) {
	if params == nil {
		params = &ListJobsByStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListJobsByStatus", params, optFns, addOperationListJobsByStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListJobsByStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The ListJobsByStatusRequest structure.
type ListJobsByStatusInput struct {

	// To get information about all of the jobs associated with the current AWS account
	// that have a given status, specify the following status: Submitted, Progressing,
	// Complete, Canceled, or Error.
	//
	// This member is required.
	Status *string

	// To list jobs in chronological order by the date and time that they were
	// submitted, enter true. To list jobs in reverse chronological order, enter false.
	Ascending *string

	// When Elastic Transcoder returns more than one page of results, use pageToken in
	// subsequent GET requests to get each successive page of results.
	PageToken *string
}

// The ListJobsByStatusResponse structure.
type ListJobsByStatusOutput struct {

	// An array of Job objects that have the specified status.
	Jobs []*types.Job

	// A value that you use to access the second and subsequent pages of results, if
	// any. When the jobs in the specified pipeline fit on one page or when you've
	// reached the last page of results, the value of NextPageToken is null.
	NextPageToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListJobsByStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpListJobsByStatus{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpListJobsByStatus{})
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
	if err := addOpListJobsByStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opListJobsByStatus(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opListJobsByStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elastictranscoder",
		OperationName: "ListJobsByStatus",
	}
}
