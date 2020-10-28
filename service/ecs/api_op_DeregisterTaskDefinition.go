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

// Deregisters the specified task definition by family and revision. Upon
// deregistration, the task definition is marked as INACTIVE. Existing tasks and
// services that reference an INACTIVE task definition continue to run without
// disruption. Existing services that reference an INACTIVE task definition can
// still scale up or down by modifying the service's desired count. You cannot use
// an INACTIVE task definition to run new tasks or create new services, and you
// cannot update an existing service to reference an INACTIVE task definition.
// However, there may be up to a 10-minute window following deregistration where
// these restrictions have not yet taken effect. At this time, INACTIVE task
// definitions remain discoverable in your account indefinitely. However, this
// behavior is subject to change in the future, so you should not rely on INACTIVE
// task definitions persisting beyond the lifecycle of any associated tasks and
// services.
func (c *Client) DeregisterTaskDefinition(ctx context.Context, params *DeregisterTaskDefinitionInput, optFns ...func(*Options)) (*DeregisterTaskDefinitionOutput, error) {
	if params == nil {
		params = &DeregisterTaskDefinitionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeregisterTaskDefinition", params, optFns, addOperationDeregisterTaskDefinitionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeregisterTaskDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeregisterTaskDefinitionInput struct {

	// The family and revision (family:revision) or full Amazon Resource Name (ARN) of
	// the task definition to deregister. You must specify a revision.
	//
	// This member is required.
	TaskDefinition *string
}

type DeregisterTaskDefinitionOutput struct {

	// The full description of the deregistered task.
	TaskDefinition *types.TaskDefinition

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeregisterTaskDefinitionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpDeregisterTaskDefinition{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpDeregisterTaskDefinition{})
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
	if err := addOpDeregisterTaskDefinitionValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDeregisterTaskDefinition(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDeregisterTaskDefinition(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecs",
		OperationName: "DeregisterTaskDefinition",
	}
}
