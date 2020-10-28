// Code generated by smithy-go-codegen DO NOT EDIT.

package codedeploy

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a deployment group to which application revisions are deployed.
func (c *Client) CreateDeploymentGroup(ctx context.Context, params *CreateDeploymentGroupInput, optFns ...func(*Options)) (*CreateDeploymentGroupOutput, error) {
	if params == nil {
		params = &CreateDeploymentGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDeploymentGroup", params, optFns, addOperationCreateDeploymentGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDeploymentGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a CreateDeploymentGroup operation.
type CreateDeploymentGroupInput struct {

	// The name of an AWS CodeDeploy application associated with the IAM user or AWS
	// account.
	//
	// This member is required.
	ApplicationName *string

	// The name of a new deployment group for the specified application.
	//
	// This member is required.
	DeploymentGroupName *string

	// A service role Amazon Resource Name (ARN) that allows AWS CodeDeploy to act on
	// the user's behalf when interacting with AWS services.
	//
	// This member is required.
	ServiceRoleArn *string

	// Information to add about Amazon CloudWatch alarms when the deployment group is
	// created.
	AlarmConfiguration *types.AlarmConfiguration

	// Configuration information for an automatic rollback that is added when a
	// deployment group is created.
	AutoRollbackConfiguration *types.AutoRollbackConfiguration

	// A list of associated Amazon EC2 Auto Scaling groups.
	AutoScalingGroups []*string

	// Information about blue/green deployment options for a deployment group.
	BlueGreenDeploymentConfiguration *types.BlueGreenDeploymentConfiguration

	// If specified, the deployment configuration name can be either one of the
	// predefined configurations provided with AWS CodeDeploy or a custom deployment
	// configuration that you create by calling the create deployment configuration
	// operation. CodeDeployDefault.OneAtATime is the default deployment configuration.
	// It is used if a configuration isn't specified for the deployment or deployment
	// group. For more information about the predefined deployment configurations in
	// AWS CodeDeploy, see Working with Deployment Configurations in CodeDeploy
	// (https://docs.aws.amazon.com/codedeploy/latest/userguide/deployment-configurations.html)
	// in the AWS CodeDeploy User Guide.
	DeploymentConfigName *string

	// Information about the type of deployment, in-place or blue/green, that you want
	// to run and whether to route deployment traffic behind a load balancer.
	DeploymentStyle *types.DeploymentStyle

	// The Amazon EC2 tags on which to filter. The deployment group includes EC2
	// instances with any of the specified tags. Cannot be used in the same call as
	// ec2TagSet.
	Ec2TagFilters []*types.EC2TagFilter

	// Information about groups of tags applied to EC2 instances. The deployment group
	// includes only EC2 instances identified by all the tag groups. Cannot be used in
	// the same call as ec2TagFilters.
	Ec2TagSet *types.EC2TagSet

	// The target Amazon ECS services in the deployment group. This applies only to
	// deployment groups that use the Amazon ECS compute platform. A target Amazon ECS
	// service is specified as an Amazon ECS cluster and service name pair using the
	// format :.
	EcsServices []*types.ECSService

	// Information about the load balancer used in a deployment.
	LoadBalancerInfo *types.LoadBalancerInfo

	// The on-premises instance tags on which to filter. The deployment group includes
	// on-premises instances with any of the specified tags. Cannot be used in the same
	// call as OnPremisesTagSet.
	OnPremisesInstanceTagFilters []*types.TagFilter

	// Information about groups of tags applied to on-premises instances. The
	// deployment group includes only on-premises instances identified by all of the
	// tag groups. Cannot be used in the same call as onPremisesInstanceTagFilters.
	OnPremisesTagSet *types.OnPremisesTagSet

	// The metadata that you apply to CodeDeploy deployment groups to help you organize
	// and categorize them. Each tag consists of a key and an optional value, both of
	// which you define.
	Tags []*types.Tag

	// Information about triggers to create when the deployment group is created. For
	// examples, see Create a Trigger for an AWS CodeDeploy Event
	// (https://docs.aws.amazon.com/codedeploy/latest/userguide/how-to-notify-sns.html)
	// in the AWS CodeDeploy User Guide.
	TriggerConfigurations []*types.TriggerConfig
}

// Represents the output of a CreateDeploymentGroup operation.
type CreateDeploymentGroupOutput struct {

	// A unique deployment group ID.
	DeploymentGroupId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateDeploymentGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpCreateDeploymentGroup{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpCreateDeploymentGroup{})
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
	if err := addOpCreateDeploymentGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opCreateDeploymentGroup(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opCreateDeploymentGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codedeploy",
		OperationName: "CreateDeploymentGroup",
	}
}
