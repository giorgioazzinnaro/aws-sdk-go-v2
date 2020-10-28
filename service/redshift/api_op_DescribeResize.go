// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about the last resize operation for the specified cluster.
// If no resize operation has ever been initiated for the specified cluster, a HTTP
// 404 error is returned. If a resize operation was initiated and completed, the
// status of the resize remains as SUCCEEDED until the next resize. A resize
// operation can be requested using ModifyCluster and specifying a different number
// or type of nodes for the cluster.
func (c *Client) DescribeResize(ctx context.Context, params *DescribeResizeInput, optFns ...func(*Options)) (*DescribeResizeOutput, error) {
	if params == nil {
		params = &DescribeResizeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeResize", params, optFns, addOperationDescribeResizeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeResizeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DescribeResizeInput struct {

	// The unique identifier of a cluster whose resize progress you are requesting.
	// This parameter is case-sensitive. By default, resize operations for all clusters
	// defined for an AWS account are returned.
	//
	// This member is required.
	ClusterIdentifier *string
}

// Describes the result of a cluster resize operation.
type DescribeResizeOutput struct {

	// The average rate of the resize operation over the last few minutes, measured in
	// megabytes per second. After the resize operation completes, this value shows the
	// average rate of the entire resize operation.
	AvgResizeRateInMegaBytesPerSecond *float64

	// The percent of data transferred from source cluster to target cluster.
	DataTransferProgressPercent *float64

	// The amount of seconds that have elapsed since the resize operation began. After
	// the resize operation completes, this value shows the total actual time, in
	// seconds, for the resize operation.
	ElapsedTimeInSeconds *int64

	// The estimated time remaining, in seconds, until the resize operation is
	// complete. This value is calculated based on the average resize rate and the
	// estimated amount of data remaining to be processed. Once the resize operation is
	// complete, this value will be 0.
	EstimatedTimeToCompletionInSeconds *int64

	// The names of tables that have been completely imported . Valid Values: List of
	// table names.
	ImportTablesCompleted []*string

	// The names of tables that are being currently imported. Valid Values: List of
	// table names.
	ImportTablesInProgress []*string

	// The names of tables that have not been yet imported. Valid Values: List of table
	// names
	ImportTablesNotStarted []*string

	// An optional string to provide additional details about the resize action.
	Message *string

	// While the resize operation is in progress, this value shows the current amount
	// of data, in megabytes, that has been processed so far. When the resize operation
	// is complete, this value shows the total amount of data, in megabytes, on the
	// cluster, which may be more or less than TotalResizeDataInMegaBytes (the
	// estimated total amount of data before resize).
	ProgressInMegaBytes *int64

	// An enum with possible values of ClassicResize and ElasticResize. These values
	// describe the type of resize operation being performed.
	ResizeType *string

	// The status of the resize operation. Valid Values: NONE | IN_PROGRESS | FAILED |
	// SUCCEEDED | CANCELLING
	Status *string

	// The cluster type after the resize operation is complete. Valid Values:
	// multi-node | single-node
	TargetClusterType *string

	// The type of encryption for the cluster after the resize is complete. Possible
	// values are KMS and None.
	TargetEncryptionType *string

	// The node type that the cluster will have after the resize operation is complete.
	TargetNodeType *string

	// The number of nodes that the cluster will have after the resize operation is
	// complete.
	TargetNumberOfNodes *int32

	// The estimated total amount of data, in megabytes, on the cluster before the
	// resize operation began.
	TotalResizeDataInMegaBytes *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeResizeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsquery_serializeOpDescribeResize{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsquery_deserializeOpDescribeResize{})
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
	if err := addOpDescribeResizeValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDescribeResize(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDescribeResize(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "DescribeResize",
	}
}
