// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Terminates the specified instance and optionally adjusts the desired group size.
// This call simply makes a termination request. The instance is not terminated
// immediately. When an instance is terminated, the instance status changes to
// terminated. You can't connect to or start an instance after you've terminated
// it. If you do not specify the option to decrement the desired capacity, Amazon
// EC2 Auto Scaling launches instances to replace the ones that are terminated. By
// default, Amazon EC2 Auto Scaling balances instances across all Availability
// Zones. If you decrement the desired capacity, your Auto Scaling group can become
// unbalanced between Availability Zones. Amazon EC2 Auto Scaling tries to
// rebalance the group, and rebalancing might terminate instances in other zones.
// For more information, see Rebalancing Activities
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/auto-scaling-benefits.html#AutoScalingBehavior.InstanceUsage)
// in the Amazon EC2 Auto Scaling User Guide.
func (c *Client) TerminateInstanceInAutoScalingGroup(ctx context.Context, params *TerminateInstanceInAutoScalingGroupInput, optFns ...func(*Options)) (*TerminateInstanceInAutoScalingGroupOutput, error) {
	if params == nil {
		params = &TerminateInstanceInAutoScalingGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "TerminateInstanceInAutoScalingGroup", params, optFns, addOperationTerminateInstanceInAutoScalingGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*TerminateInstanceInAutoScalingGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TerminateInstanceInAutoScalingGroupInput struct {

	// The ID of the instance.
	//
	// This member is required.
	InstanceId *string

	// Indicates whether terminating the instance also decrements the size of the Auto
	// Scaling group.
	//
	// This member is required.
	ShouldDecrementDesiredCapacity *bool
}

type TerminateInstanceInAutoScalingGroupOutput struct {

	// A scaling activity.
	Activity *types.Activity

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationTerminateInstanceInAutoScalingGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsquery_serializeOpTerminateInstanceInAutoScalingGroup{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsquery_deserializeOpTerminateInstanceInAutoScalingGroup{})
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
	if err := addOpTerminateInstanceInAutoScalingGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opTerminateInstanceInAutoScalingGroup(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opTerminateInstanceInAutoScalingGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "TerminateInstanceInAutoScalingGroup",
	}
}
