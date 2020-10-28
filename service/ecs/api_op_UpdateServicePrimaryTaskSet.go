// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modifies which task set in a service is the primary task set. Any parameters
// that are updated on the primary task set in a service will transition to the
// service. This is used when a service uses the EXTERNAL deployment controller
// type. For more information, see Amazon ECS Deployment Types
// (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html)
// in the Amazon Elastic Container Service Developer Guide.
func (c *Client) UpdateServicePrimaryTaskSet(ctx context.Context, params *UpdateServicePrimaryTaskSetInput, optFns ...func(*Options)) (*UpdateServicePrimaryTaskSetOutput, error) {
	if params == nil {
		params = &UpdateServicePrimaryTaskSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateServicePrimaryTaskSet", params, optFns, addOperationUpdateServicePrimaryTaskSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateServicePrimaryTaskSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateServicePrimaryTaskSetInput struct {

	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts the
	// service that the task set exists in.
	//
	// This member is required.
	Cluster *string

	// The short name or full Amazon Resource Name (ARN) of the task set to set as the
	// primary task set in the deployment.
	//
	// This member is required.
	PrimaryTaskSet *string

	// The short name or full Amazon Resource Name (ARN) of the service that the task
	// set exists in.
	//
	// This member is required.
	Service *string
}

type UpdateServicePrimaryTaskSetOutput struct {

	// Information about a set of Amazon ECS tasks in either an AWS CodeDeploy or an
	// EXTERNAL deployment. An Amazon ECS task set includes details such as the desired
	// number of tasks, how many tasks are running, and whether the task set serves
	// production traffic.
	TaskSet *types.TaskSet

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateServicePrimaryTaskSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpUpdateServicePrimaryTaskSet{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpUpdateServicePrimaryTaskSet{})
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
	if err := addOpUpdateServicePrimaryTaskSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opUpdateServicePrimaryTaskSet(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opUpdateServicePrimaryTaskSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecs",
		OperationName: "UpdateServicePrimaryTaskSet",
	}
}
