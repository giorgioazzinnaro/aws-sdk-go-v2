// Code generated by smithy-go-codegen DO NOT EDIT.

package glacier

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	glaciercust "github.com/aws/aws-sdk-go-v2/service/glacier/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/glacier/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This operation lists jobs for a vault, including jobs that are in-progress and
// jobs that have recently finished. The List Job operation returns a list of these
// jobs sorted by job initiation time. Amazon Glacier retains recently completed
// jobs for a period before deleting them; however, it eventually removes completed
// jobs. The output of completed jobs can be retrieved. Retaining completed jobs
// for a period of time after they have completed enables you to get a job output
// in the event you miss the job completion notification or your first attempt to
// download it fails. For example, suppose you start an archive retrieval job to
// download an archive. After the job completes, you start to download the archive
// but encounter a network error. In this scenario, you can retry and download the
// archive while the job exists. The List Jobs operation supports pagination. You
// should always check the response Marker field. If there are no more jobs to
// list, the Marker field is set to null. If there are more jobs to list, the
// Marker field is set to a non-null value, which you can use to continue the
// pagination of the list. To return a list of jobs that begins at a specific job,
// set the marker request parameter to the Marker value for that job that you
// obtained from a previous List Jobs request. You can set a maximum limit for the
// number of jobs returned in the response by specifying the limit parameter in the
// request. The default limit is 50. The number of jobs returned might be fewer
// than the limit, but the number of returned jobs never exceeds the limit.
// Additionally, you can filter the jobs list returned by specifying the optional
// statuscode parameter or completed parameter, or both. Using the statuscode
// parameter, you can specify to return only jobs that match either the InProgress,
// Succeeded, or Failed status. Using the completed parameter, you can specify to
// return only jobs that were completed (true) or jobs that were not completed
// (false). For more information about using this operation, see the documentation
// for the underlying REST API List Jobs
// (https://docs.aws.amazon.com/amazonglacier/latest/dev/api-jobs-get.html).
func (c *Client) ListJobs(ctx context.Context, params *ListJobsInput, optFns ...func(*Options)) (*ListJobsOutput, error) {
	if params == nil {
		params = &ListJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListJobs", params, optFns, addOperationListJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Provides options for retrieving a job list for an Amazon S3 Glacier vault.
type ListJobsInput struct {

	// The AccountId value is the AWS account ID of the account that owns the vault.
	// You can either specify an AWS account ID or optionally a single '-' (hyphen), in
	// which case Amazon S3 Glacier uses the AWS account ID associated with the
	// credentials used to sign the request. If you use an account ID, do not include
	// any hyphens ('-') in the ID.
	//
	// This member is required.
	AccountId *string

	// The name of the vault.
	//
	// This member is required.
	VaultName *string

	// The state of the jobs to return. You can specify true or false.
	Completed *string

	// The maximum number of jobs to be returned. The default limit is 50. The number
	// of jobs returned might be fewer than the specified limit, but the number of
	// returned jobs never exceeds the limit.
	Limit *string

	// An opaque string used for pagination. This value specifies the job at which the
	// listing of jobs should begin. Get the marker value from a previous List Jobs
	// response. You only need to include the marker if you are continuing the
	// pagination of results started in a previous List Jobs request.
	Marker *string

	// The type of job status to return. You can specify the following values:
	// InProgress, Succeeded, or Failed.
	Statuscode *string
}

// Contains the Amazon S3 Glacier response to your request.
type ListJobsOutput struct {

	// A list of job objects. Each job object contains metadata describing the job.
	JobList []*types.GlacierJobDescription

	// An opaque string used for pagination that specifies the job at which the listing
	// of jobs should begin. You get the marker value from a previous List Jobs
	// response. You only need to include the marker if you are continuing the
	// pagination of the results started in a previous List Jobs request.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpListJobs{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpListJobs{})
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
	if err := addOpListJobsValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opListJobs(options.Region)); err != nil {
		return err
	}
	if err := addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err := addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err := glaciercust.AddTreeHashMiddleware(stack); err != nil {
		return err
	}
	if err := glaciercust.AddGlacierAPIVersionMiddleware(stack, ServiceAPIVersion); err != nil {
		return err
	}
	if err := glaciercust.AddDefaultAccountIDMiddleware(stack, setDefaultAccountID); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opListJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glacier",
		OperationName: "ListJobs",
	}
}
