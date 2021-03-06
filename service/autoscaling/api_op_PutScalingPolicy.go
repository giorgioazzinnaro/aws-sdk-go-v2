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

// Creates or updates a scaling policy for an Auto Scaling group. For more
// information about using scaling policies to scale your Auto Scaling group, see
// Target Tracking Scaling Policies
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-scaling-target-tracking.html)
// and Step and Simple Scaling Policies
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-scaling-simple-step.html)
// in the Amazon EC2 Auto Scaling User Guide.
func (c *Client) PutScalingPolicy(ctx context.Context, params *PutScalingPolicyInput, optFns ...func(*Options)) (*PutScalingPolicyOutput, error) {
	if params == nil {
		params = &PutScalingPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutScalingPolicy", params, optFns, addOperationPutScalingPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutScalingPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutScalingPolicyInput struct {

	// The name of the Auto Scaling group.
	//
	// This member is required.
	AutoScalingGroupName *string

	// The name of the policy.
	//
	// This member is required.
	PolicyName *string

	// Specifies how the scaling adjustment is interpreted (for example, an absolute
	// number or a percentage). The valid values are ChangeInCapacity, ExactCapacity,
	// and PercentChangeInCapacity. Required if the policy type is StepScaling or
	// SimpleScaling. For more information, see Scaling Adjustment Types
	// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-scaling-simple-step.html#as-scaling-adjustment)
	// in the Amazon EC2 Auto Scaling User Guide.
	AdjustmentType *string

	// The duration of the policy's cooldown period, in seconds. When a cooldown period
	// is specified here, it overrides the default cooldown period defined for the Auto
	// Scaling group. Valid only if the policy type is SimpleScaling. For more
	// information, see Scaling Cooldowns for Amazon EC2 Auto Scaling
	// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/Cooldown.html) in the
	// Amazon EC2 Auto Scaling User Guide.
	Cooldown *int32

	// Indicates whether the scaling policy is enabled or disabled. The default is
	// enabled. For more information, see Disabling a Scaling Policy for an Auto
	// Scaling Group
	// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-enable-disable-scaling-policy.html)
	// in the Amazon EC2 Auto Scaling User Guide.
	Enabled *bool

	// The estimated time, in seconds, until a newly launched instance can contribute
	// to the CloudWatch metrics. If not provided, the default is to use the value from
	// the default cooldown period for the Auto Scaling group. Valid only if the policy
	// type is TargetTrackingScaling or StepScaling.
	EstimatedInstanceWarmup *int32

	// The aggregation type for the CloudWatch metrics. The valid values are Minimum,
	// Maximum, and Average. If the aggregation type is null, the value is treated as
	// Average. Valid only if the policy type is StepScaling.
	MetricAggregationType *string

	// The minimum value to scale by when the adjustment type is
	// PercentChangeInCapacity. For example, suppose that you create a step scaling
	// policy to scale out an Auto Scaling group by 25 percent and you specify a
	// MinAdjustmentMagnitude of 2. If the group has 4 instances and the scaling policy
	// is performed, 25 percent of 4 is 1. However, because you specified a
	// MinAdjustmentMagnitude of 2, Amazon EC2 Auto Scaling scales out the group by 2
	// instances. Valid only if the policy type is StepScaling or SimpleScaling. For
	// more information, see Scaling Adjustment Types
	// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-scaling-simple-step.html#as-scaling-adjustment)
	// in the Amazon EC2 Auto Scaling User Guide. Some Auto Scaling groups use instance
	// weights. In this case, set the MinAdjustmentMagnitude to a value that is at
	// least as large as your largest instance weight.
	MinAdjustmentMagnitude *int32

	// Available for backward compatibility. Use MinAdjustmentMagnitude instead.
	MinAdjustmentStep *int32

	// One of the following policy types:
	//
	// * TargetTrackingScaling
	//
	// * StepScaling
	//
	// *
	// SimpleScaling (default)
	PolicyType *string

	// The amount by which to scale, based on the specified adjustment type. A positive
	// value adds to the current capacity while a negative number removes from the
	// current capacity. For exact capacity, you must specify a positive value.
	// Required if the policy type is SimpleScaling. (Not used with any other policy
	// type.)
	ScalingAdjustment *int32

	// A set of adjustments that enable you to scale based on the size of the alarm
	// breach. Required if the policy type is StepScaling. (Not used with any other
	// policy type.)
	StepAdjustments []*types.StepAdjustment

	// A target tracking scaling policy. Includes support for predefined or customized
	// metrics. The following predefined metrics are available:
	//
	// *
	// ASGAverageCPUUtilization
	//
	// * ASGAverageNetworkIn
	//
	// * ASGAverageNetworkOut
	//
	// *
	// ALBRequestCountPerTarget
	//
	// If you specify ALBRequestCountPerTarget for the
	// metric, you must specify the ResourceLabel parameter with the
	// PredefinedMetricSpecification. For more information, see
	// TargetTrackingConfiguration
	// (https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_TargetTrackingConfiguration.html)
	// in the Amazon EC2 Auto Scaling API Reference. Required if the policy type is
	// TargetTrackingScaling.
	TargetTrackingConfiguration *types.TargetTrackingConfiguration
}

// Contains the output of PutScalingPolicy.
type PutScalingPolicyOutput struct {

	// The CloudWatch alarms created for the target tracking scaling policy.
	Alarms []*types.Alarm

	// The Amazon Resource Name (ARN) of the policy.
	PolicyARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutScalingPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpPutScalingPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpPutScalingPolicy{}, middleware.After)
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
	addOpPutScalingPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutScalingPolicy(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opPutScalingPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "PutScalingPolicy",
	}
}
