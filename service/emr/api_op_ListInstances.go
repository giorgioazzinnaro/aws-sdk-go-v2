// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Provides information for all active EC2 instances and EC2 instances terminated
// in the last 30 days, up to a maximum of 2,000. EC2 instances in any of the
// following states are considered active: AWAITING_FULFILLMENT, PROVISIONING,
// BOOTSTRAPPING, RUNNING.
func (c *Client) ListInstances(ctx context.Context, params *ListInstancesInput, optFns ...func(*Options)) (*ListInstancesOutput, error) {
	if params == nil {
		params = &ListInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListInstances", params, optFns, addOperationListInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// This input determines which instances to list.
type ListInstancesInput struct {

	// The identifier of the cluster for which to list the instances.
	//
	// This member is required.
	ClusterId *string

	// The unique identifier of the instance fleet.
	InstanceFleetId *string

	// The node type of the instance fleet. For example MASTER, CORE, or TASK.
	InstanceFleetType types.InstanceFleetType

	// The identifier of the instance group for which to list the instances.
	InstanceGroupId *string

	// The type of instance group for which to list the instances.
	InstanceGroupTypes []types.InstanceGroupType

	// A list of instance states that will filter the instances returned with this
	// request.
	InstanceStates []types.InstanceState

	// The pagination token that indicates the next set of results to retrieve.
	Marker *string
}

// This output contains the list of instances.
type ListInstancesOutput struct {

	// The list of instances for the cluster and given filters.
	Instances []*types.Instance

	// The pagination token that indicates the next set of results to retrieve.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListInstances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListInstances{}, middleware.After)
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
	addOpListInstancesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListInstances(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListInstances(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticmapreduce",
		OperationName: "ListInstances",
	}
}
