// Code generated by smithy-go-codegen DO NOT EDIT.

package robomaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/robomaker/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Describes a simulation job.
func (c *Client) DescribeSimulationJob(ctx context.Context, params *DescribeSimulationJobInput, optFns ...func(*Options)) (*DescribeSimulationJobOutput, error) {
	if params == nil {
		params = &DescribeSimulationJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeSimulationJob", params, optFns, addOperationDescribeSimulationJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeSimulationJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeSimulationJobInput struct {

	// The Amazon Resource Name (ARN) of the simulation job to be described.
	//
	// This member is required.
	Job *string
}

type DescribeSimulationJobOutput struct {

	// The Amazon Resource Name (ARN) of the simulation job.
	Arn *string

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request.
	ClientRequestToken *string

	// Compute information for the simulation job.
	Compute *types.ComputeResponse

	// The data sources for the simulation job.
	DataSources []*types.DataSource

	// The failure behavior for the simulation job.
	FailureBehavior types.FailureBehavior

	// The failure code of the simulation job if it failed: InternalServiceError
	// Internal service error. RobotApplicationCrash Robot application exited
	// abnormally. SimulationApplicationCrash Simulation application exited abnormally.
	// BadPermissionsRobotApplication Robot application bundle could not be downloaded.
	// BadPermissionsSimulationApplication Simulation application bundle could not be
	// downloaded. BadPermissionsS3Output Unable to publish outputs to
	// customer-provided S3 bucket. BadPermissionsCloudwatchLogs Unable to publish logs
	// to customer-provided CloudWatch Logs resource. SubnetIpLimitExceeded Subnet IP
	// limit exceeded. ENILimitExceeded ENI limit exceeded.
	// BadPermissionsUserCredentials Unable to use the Role provided.
	// InvalidBundleRobotApplication Robot bundle cannot be extracted (invalid format,
	// bundling error, or other issue). InvalidBundleSimulationApplication Simulation
	// bundle cannot be extracted (invalid format, bundling error, or other issue).
	// RobotApplicationVersionMismatchedEtag Etag for RobotApplication does not match
	// value during version creation. SimulationApplicationVersionMismatchedEtag Etag
	// for SimulationApplication does not match value during version creation.
	FailureCode types.SimulationJobErrorCode

	// Details about why the simulation job failed. For more information about
	// troubleshooting, see Troubleshooting
	// (https://docs.aws.amazon.com/robomaker/latest/dg/troubleshooting.html).
	FailureReason *string

	// The IAM role that allows the simulation instance to call the AWS APIs that are
	// specified in its associated policies on your behalf.
	IamRole *string

	// The time, in milliseconds since the epoch, when the simulation job was last
	// started.
	LastStartedAt *time.Time

	// The time, in milliseconds since the epoch, when the simulation job was last
	// updated.
	LastUpdatedAt *time.Time

	// The logging configuration.
	LoggingConfig *types.LoggingConfig

	// The maximum job duration in seconds. The value must be 8 days (691,200 seconds)
	// or less.
	MaxJobDurationInSeconds *int64

	// The name of the simulation job.
	Name *string

	// The network interface information for the simulation job.
	NetworkInterface *types.NetworkInterface

	// Location for output files generated by the simulation job.
	OutputLocation *types.OutputLocation

	// A list of robot applications.
	RobotApplications []*types.RobotApplicationConfig

	// A list of simulation applications.
	SimulationApplications []*types.SimulationApplicationConfig

	// The simulation job execution duration in milliseconds.
	SimulationTimeMillis *int64

	// The status of the simulation job.
	Status types.SimulationJobStatus

	// The list of all tags added to the specified simulation job.
	Tags map[string]*string

	// The VPC configuration.
	VpcConfig *types.VPCConfigResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeSimulationJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpDescribeSimulationJob{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpDescribeSimulationJob{})
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
	if err := addOpDescribeSimulationJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDescribeSimulationJob(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDescribeSimulationJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "robomaker",
		OperationName: "DescribeSimulationJob",
	}
}
