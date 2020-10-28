// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns information about the overall health of the specified environment. The
// DescribeEnvironmentHealth operation is only available with AWS Elastic Beanstalk
// Enhanced Health.
func (c *Client) DescribeEnvironmentHealth(ctx context.Context, params *DescribeEnvironmentHealthInput, optFns ...func(*Options)) (*DescribeEnvironmentHealthOutput, error) {
	if params == nil {
		params = &DescribeEnvironmentHealthInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEnvironmentHealth", params, optFns, addOperationDescribeEnvironmentHealthMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEnvironmentHealthOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// See the example below to learn how to create a request body.
type DescribeEnvironmentHealthInput struct {

	// Specify the response elements to return. To retrieve all attributes, set to All.
	// If no attribute names are specified, returns the name of the environment.
	AttributeNames []types.EnvironmentHealthAttribute

	// Specify the environment by ID. You must specify either this or an
	// EnvironmentName, or both.
	EnvironmentId *string

	// Specify the environment by name. You must specify either this or an
	// EnvironmentName, or both.
	EnvironmentName *string
}

// Health details for an AWS Elastic Beanstalk environment.
type DescribeEnvironmentHealthOutput struct {

	// Application request metrics for the environment.
	ApplicationMetrics *types.ApplicationMetrics

	// Descriptions of the data that contributed to the environment's current health
	// status.
	Causes []*string

	// The health color
	// (https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/health-enhanced-status.html)
	// of the environment.
	Color *string

	// The environment's name.
	EnvironmentName *string

	// The health status
	// (https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/health-enhanced-status.html)
	// of the environment. For example, Ok.
	HealthStatus *string

	// Summary health information for the instances in the environment.
	InstancesHealth *types.InstanceHealthSummary

	// The date and time that the health information was retrieved.
	RefreshedAt *time.Time

	// The environment's operational status. Ready, Launching, Updating, Terminating,
	// or Terminated.
	Status types.EnvironmentHealth

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeEnvironmentHealthMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsquery_serializeOpDescribeEnvironmentHealth{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsquery_deserializeOpDescribeEnvironmentHealth{})
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
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDescribeEnvironmentHealth(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDescribeEnvironmentHealth(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticbeanstalk",
		OperationName: "DescribeEnvironmentHealth",
	}
}
