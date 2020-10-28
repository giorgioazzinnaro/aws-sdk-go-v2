// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the running instances for the specified Spot Fleet.
func (c *Client) DescribeSpotFleetInstances(ctx context.Context, params *DescribeSpotFleetInstancesInput, optFns ...func(*Options)) (*DescribeSpotFleetInstancesOutput, error) {
	if params == nil {
		params = &DescribeSpotFleetInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeSpotFleetInstances", params, optFns, addOperationDescribeSpotFleetInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeSpotFleetInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DescribeSpotFleetInstances.
type DescribeSpotFleetInstancesInput struct {

	// The ID of the Spot Fleet request.
	//
	// This member is required.
	SpotFleetRequestId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The maximum number of results to return in a single call. Specify a value
	// between 1 and 1000. The default value is 1000. To retrieve the remaining
	// results, make another call with the returned NextToken value.
	MaxResults *int32

	// The token for the next set of results.
	NextToken *string
}

// Contains the output of DescribeSpotFleetInstances.
type DescribeSpotFleetInstancesOutput struct {

	// The running instances. This list is refreshed periodically and might be out of
	// date.
	ActiveInstances []*types.ActiveInstance

	// The token required to retrieve the next set of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// The ID of the Spot Fleet request.
	SpotFleetRequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeSpotFleetInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsEc2query_serializeOpDescribeSpotFleetInstances{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsEc2query_deserializeOpDescribeSpotFleetInstances{})
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
	if err := addOpDescribeSpotFleetInstancesValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDescribeSpotFleetInstances(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDescribeSpotFleetInstances(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeSpotFleetInstances",
	}
}
