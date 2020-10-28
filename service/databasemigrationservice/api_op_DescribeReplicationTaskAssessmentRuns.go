// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a paginated list of premigration assessment runs based on filter
// settings. These filter settings can specify a combination of premigration
// assessment runs, migration tasks, replication instances, and assessment run
// status values. This operation doesn't return information about individual
// assessments. For this information, see the
// DescribeReplicationTaskIndividualAssessments operation.
func (c *Client) DescribeReplicationTaskAssessmentRuns(ctx context.Context, params *DescribeReplicationTaskAssessmentRunsInput, optFns ...func(*Options)) (*DescribeReplicationTaskAssessmentRunsOutput, error) {
	if params == nil {
		params = &DescribeReplicationTaskAssessmentRunsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeReplicationTaskAssessmentRuns", params, optFns, addOperationDescribeReplicationTaskAssessmentRunsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeReplicationTaskAssessmentRunsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DescribeReplicationTaskAssessmentRunsInput struct {

	// Filters applied to the premigration assessment runs described in the form of
	// key-value pairs. Valid filter names: replication-task-assessment-run-arn,
	// replication-task-arn, replication-instance-arn, status
	Filters []*types.Filter

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords.
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that the remaining results can be retrieved.
	MaxRecords *int32
}

//
type DescribeReplicationTaskAssessmentRunsOutput struct {

	// A pagination token returned for you to pass to a subsequent request. If you pass
	// this token as the Marker value in a subsequent request, the response includes
	// only records beyond the marker, up to the value specified in the request by
	// MaxRecords.
	Marker *string

	// One or more premigration assessment runs as specified by Filters.
	ReplicationTaskAssessmentRuns []*types.ReplicationTaskAssessmentRun

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeReplicationTaskAssessmentRunsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpDescribeReplicationTaskAssessmentRuns{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpDescribeReplicationTaskAssessmentRuns{})
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
	if err := addOpDescribeReplicationTaskAssessmentRunsValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDescribeReplicationTaskAssessmentRuns(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDescribeReplicationTaskAssessmentRuns(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dms",
		OperationName: "DescribeReplicationTaskAssessmentRuns",
	}
}
