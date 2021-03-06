// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns details of the update actions
func (c *Client) DescribeUpdateActions(ctx context.Context, params *DescribeUpdateActionsInput, optFns ...func(*Options)) (*DescribeUpdateActionsOutput, error) {
	if params == nil {
		params = &DescribeUpdateActionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeUpdateActions", params, optFns, addOperationDescribeUpdateActionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeUpdateActionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeUpdateActionsInput struct {

	// The cache cluster IDs
	CacheClusterIds []*string

	// The Elasticache engine to which the update applies. Either Redis or Memcached
	Engine *string

	// An optional marker returned from a prior request. Use this marker for pagination
	// of results from this operation. If this parameter is specified, the response
	// includes only records beyond the marker, up to the value specified by
	// MaxRecords.
	Marker *string

	// The maximum number of records to include in the response
	MaxRecords *int32

	// The replication group IDs
	ReplicationGroupIds []*string

	// The unique ID of the service update
	ServiceUpdateName *string

	// The status of the service update
	ServiceUpdateStatus []types.ServiceUpdateStatus

	// The range of time specified to search for service updates that are in available
	// status
	ServiceUpdateTimeRange *types.TimeRangeFilter

	// Dictates whether to include node level update status in the response
	ShowNodeLevelUpdateStatus *bool

	// The status of the update action.
	UpdateActionStatus []types.UpdateActionStatus
}

type DescribeUpdateActionsOutput struct {

	// An optional marker returned from a prior request. Use this marker for pagination
	// of results from this operation. If this parameter is specified, the response
	// includes only records beyond the marker, up to the value specified by
	// MaxRecords.
	Marker *string

	// Returns a list of update actions
	UpdateActions []*types.UpdateAction

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeUpdateActionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeUpdateActions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeUpdateActions{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeUpdateActions(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeUpdateActions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "DescribeUpdateActions",
	}
}
