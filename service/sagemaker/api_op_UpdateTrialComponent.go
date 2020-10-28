// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Updates one or more properties of a trial component.
func (c *Client) UpdateTrialComponent(ctx context.Context, params *UpdateTrialComponentInput, optFns ...func(*Options)) (*UpdateTrialComponentOutput, error) {
	if params == nil {
		params = &UpdateTrialComponentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateTrialComponent", params, optFns, addOperationUpdateTrialComponentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateTrialComponentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateTrialComponentInput struct {

	// The name of the component to update.
	//
	// This member is required.
	TrialComponentName *string

	// The name of the component as displayed. The name doesn't need to be unique. If
	// DisplayName isn't specified, TrialComponentName is displayed.
	DisplayName *string

	// When the component ended.
	EndTime *time.Time

	// Replaces all of the component's input artifacts with the specified artifacts.
	InputArtifacts map[string]*types.TrialComponentArtifact

	// The input artifacts to remove from the component.
	InputArtifactsToRemove []*string

	// Replaces all of the component's output artifacts with the specified artifacts.
	OutputArtifacts map[string]*types.TrialComponentArtifact

	// The output artifacts to remove from the component.
	OutputArtifactsToRemove []*string

	// Replaces all of the component's hyperparameters with the specified
	// hyperparameters.
	Parameters map[string]*types.TrialComponentParameterValue

	// The hyperparameters to remove from the component.
	ParametersToRemove []*string

	// When the component started.
	StartTime *time.Time

	// The new status of the component.
	Status *types.TrialComponentStatus
}

type UpdateTrialComponentOutput struct {

	// The Amazon Resource Name (ARN) of the trial component.
	TrialComponentArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateTrialComponentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpUpdateTrialComponent{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpUpdateTrialComponent{})
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
	if err := addOpUpdateTrialComponentValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opUpdateTrialComponent(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opUpdateTrialComponent(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "UpdateTrialComponent",
	}
}
