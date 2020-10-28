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

// Describes the task sets in the specified cluster and service. This is used when
// a service uses the EXTERNAL deployment controller type. For more information,
// see Amazon ECS Deployment Types
// (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html)
// in the Amazon Elastic Container Service Developer Guide.
func (c *Client) DescribeTaskSets(ctx context.Context, params *DescribeTaskSetsInput, optFns ...func(*Options)) (*DescribeTaskSetsOutput, error) {
	if params == nil {
		params = &DescribeTaskSetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTaskSets", params, optFns, addOperationDescribeTaskSetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTaskSetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTaskSetsInput struct {

	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts the
	// service that the task sets exist in.
	//
	// This member is required.
	Cluster *string

	// The short name or full Amazon Resource Name (ARN) of the service that the task
	// sets exist in.
	//
	// This member is required.
	Service *string

	// Specifies whether to see the resource tags for the task set. If TAGS is
	// specified, the tags are included in the response. If this field is omitted, tags
	// are not included in the response.
	Include []types.TaskSetField

	// The ID or full Amazon Resource Name (ARN) of task sets to describe.
	TaskSets []*string
}

type DescribeTaskSetsOutput struct {

	// Any failures associated with the call.
	Failures []*types.Failure

	// The list of task sets described.
	TaskSets []*types.TaskSet

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeTaskSetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpDescribeTaskSets{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpDescribeTaskSets{})
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
	if err := addOpDescribeTaskSetsValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDescribeTaskSets(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDescribeTaskSets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecs",
		OperationName: "DescribeTaskSets",
	}
}
