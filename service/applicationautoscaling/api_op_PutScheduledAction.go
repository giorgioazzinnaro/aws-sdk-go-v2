// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationautoscaling

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Creates or updates a scheduled action for an Application Auto Scaling scalable
// target. Each scalable target is identified by a service namespace, resource ID,
// and scalable dimension. A scheduled action applies to the scalable target
// identified by those three attributes. You cannot create a scheduled action until
// you have registered the resource as a scalable target. When start and end times
// are specified with a recurring schedule using a cron expression or rates, they
// form the boundaries of when the recurring action starts and stops. To update a
// scheduled action, specify the parameters that you want to change. If you don't
// specify start and end times, the old values are deleted. For more information,
// see Scheduled Scaling
// (https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-scheduled-scaling.html)
// in the Application Auto Scaling User Guide. If a scalable target is
// deregistered, the scalable target is no longer available to run scheduled
// actions. Any scheduled actions that were specified for the scalable target are
// deleted.
func (c *Client) PutScheduledAction(ctx context.Context, params *PutScheduledActionInput, optFns ...func(*Options)) (*PutScheduledActionOutput, error) {
	if params == nil {
		params = &PutScheduledActionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutScheduledAction", params, optFns, addOperationPutScheduledActionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutScheduledActionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutScheduledActionInput struct {

	// The identifier of the resource associated with the scheduled action. This string
	// consists of the resource type and unique identifier.
	//
	// * ECS service - The
	// resource type is service and the unique identifier is the cluster name and
	// service name. Example: service/default/sample-webapp.
	//
	// * Spot Fleet request -
	// The resource type is spot-fleet-request and the unique identifier is the Spot
	// Fleet request ID. Example:
	// spot-fleet-request/sfr-73fbd2ce-aa30-494c-8788-1cee4EXAMPLE.
	//
	// * EMR cluster -
	// The resource type is instancegroup and the unique identifier is the cluster ID
	// and instance group ID. Example:
	// instancegroup/j-2EEZNYKUA1NTV/ig-1791Y4E1L8YI0.
	//
	// * AppStream 2.0 fleet - The
	// resource type is fleet and the unique identifier is the fleet name. Example:
	// fleet/sample-fleet.
	//
	// * DynamoDB table - The resource type is table and the
	// unique identifier is the table name. Example: table/my-table.
	//
	// * DynamoDB global
	// secondary index - The resource type is index and the unique identifier is the
	// index name. Example: table/my-table/index/my-table-index.
	//
	// * Aurora DB cluster -
	// The resource type is cluster and the unique identifier is the cluster name.
	// Example: cluster:my-db-cluster.
	//
	// * Amazon SageMaker endpoint variant - The
	// resource type is variant and the unique identifier is the resource ID. Example:
	// endpoint/my-end-point/variant/KMeansClustering.
	//
	// * Custom resources are not
	// supported with a resource type. This parameter must specify the OutputValue from
	// the CloudFormation template stack used to access the resources. The unique
	// identifier is defined by the service provider. More information is available in
	// our GitHub repository
	// (https://github.com/aws/aws-auto-scaling-custom-resource).
	//
	// * Amazon Comprehend
	// document classification endpoint - The resource type and unique identifier are
	// specified using the endpoint ARN. Example:
	// arn:aws:comprehend:us-west-2:123456789012:document-classifier-endpoint/EXAMPLE.
	//
	// *
	// Amazon Comprehend entity recognizer endpoint - The resource type and unique
	// identifier are specified using the endpoint ARN. Example:
	// arn:aws:comprehend:us-west-2:123456789012:entity-recognizer-endpoint/EXAMPLE.
	//
	// *
	// Lambda provisioned concurrency - The resource type is function and the unique
	// identifier is the function name with a function version or alias name suffix
	// that is not $LATEST. Example: function:my-function:prod or
	// function:my-function:1.
	//
	// * Amazon Keyspaces table - The resource type is table
	// and the unique identifier is the table name. Example:
	// keyspace/mykeyspace/table/mytable.
	//
	// * Amazon MSK cluster - The resource type and
	// unique identifier are specified using the cluster ARN. Example:
	// arn:aws:kafka:us-east-1:123456789012:cluster/demo-cluster-1/6357e0b2-0e6a-4b86-a0b4-70df934c2e31-5.
	//
	// This member is required.
	ResourceId *string

	// The scalable dimension. This string consists of the service namespace, resource
	// type, and scaling property.
	//
	// * ecs:service:DesiredCount - The desired task count
	// of an ECS service.
	//
	// * ec2:spot-fleet-request:TargetCapacity - The target
	// capacity of a Spot Fleet request.
	//
	// *
	// elasticmapreduce:instancegroup:InstanceCount - The instance count of an EMR
	// Instance Group.
	//
	// * appstream:fleet:DesiredCapacity - The desired capacity of an
	// AppStream 2.0 fleet.
	//
	// * dynamodb:table:ReadCapacityUnits - The provisioned read
	// capacity for a DynamoDB table.
	//
	// * dynamodb:table:WriteCapacityUnits - The
	// provisioned write capacity for a DynamoDB table.
	//
	// *
	// dynamodb:index:ReadCapacityUnits - The provisioned read capacity for a DynamoDB
	// global secondary index.
	//
	// * dynamodb:index:WriteCapacityUnits - The provisioned
	// write capacity for a DynamoDB global secondary index.
	//
	// *
	// rds:cluster:ReadReplicaCount - The count of Aurora Replicas in an Aurora DB
	// cluster. Available for Aurora MySQL-compatible edition and Aurora
	// PostgreSQL-compatible edition.
	//
	// * sagemaker:variant:DesiredInstanceCount - The
	// number of EC2 instances for an Amazon SageMaker model endpoint variant.
	//
	// *
	// custom-resource:ResourceType:Property - The scalable dimension for a custom
	// resource provided by your own application or service.
	//
	// *
	// comprehend:document-classifier-endpoint:DesiredInferenceUnits - The number of
	// inference units for an Amazon Comprehend document classification endpoint.
	//
	// *
	// comprehend:entity-recognizer-endpoint:DesiredInferenceUnits - The number of
	// inference units for an Amazon Comprehend entity recognizer endpoint.
	//
	// *
	// lambda:function:ProvisionedConcurrency - The provisioned concurrency for a
	// Lambda function.
	//
	// * cassandra:table:ReadCapacityUnits - The provisioned read
	// capacity for an Amazon Keyspaces table.
	//
	// * cassandra:table:WriteCapacityUnits -
	// The provisioned write capacity for an Amazon Keyspaces table.
	//
	// *
	// kafka:broker-storage:VolumeSize - The provisioned volume size (in GiB) for
	// brokers in an Amazon MSK cluster.
	//
	// This member is required.
	ScalableDimension types.ScalableDimension

	// The name of the scheduled action. This name must be unique among all other
	// scheduled actions on the specified scalable target.
	//
	// This member is required.
	ScheduledActionName *string

	// The namespace of the AWS service that provides the resource. For a resource
	// provided by your own application or service, use custom-resource instead.
	//
	// This member is required.
	ServiceNamespace types.ServiceNamespace

	// The date and time for the recurring schedule to end.
	EndTime *time.Time

	// The new minimum and maximum capacity. You can set both values or just one. At
	// the scheduled time, if the current capacity is below the minimum capacity,
	// Application Auto Scaling scales out to the minimum capacity. If the current
	// capacity is above the maximum capacity, Application Auto Scaling scales in to
	// the maximum capacity.
	ScalableTargetAction *types.ScalableTargetAction

	// The schedule for this action. The following formats are supported:
	//
	// * At
	// expressions - "at(yyyy-mm-ddThh:mm:ss)"
	//
	// * Rate expressions - "rate(value
	// unit)"
	//
	// * Cron expressions - "cron(fields)"
	//
	// At expressions are useful for
	// one-time schedules. Specify the time in UTC. For rate expressions, value is a
	// positive integer and unit is minute | minutes | hour | hours | day | days. For
	// more information about cron expressions, see Cron Expressions
	// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html#CronExpressions)
	// in the Amazon CloudWatch Events User Guide. For examples of using these
	// expressions, see Scheduled Scaling
	// (https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-scheduled-scaling.html)
	// in the Application Auto Scaling User Guide.
	Schedule *string

	// The date and time for this scheduled action to start.
	StartTime *time.Time
}

type PutScheduledActionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutScheduledActionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutScheduledAction{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutScheduledAction{}, middleware.After)
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
	addOpPutScheduledActionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutScheduledAction(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opPutScheduledAction(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "application-autoscaling",
		OperationName: "PutScheduledAction",
	}
}
