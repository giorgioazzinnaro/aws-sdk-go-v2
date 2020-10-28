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

// Provides a list of a trials component's properties.
func (c *Client) DescribeTrialComponent(ctx context.Context, params *DescribeTrialComponentInput, optFns ...func(*Options)) (*DescribeTrialComponentOutput, error) {
	if params == nil {
		params = &DescribeTrialComponentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTrialComponent", params, optFns, addOperationDescribeTrialComponentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTrialComponentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTrialComponentInput struct {

	// The name of the trial component to describe.
	//
	// This member is required.
	TrialComponentName *string
}

type DescribeTrialComponentOutput struct {

	// Who created the component.
	CreatedBy *types.UserContext

	// When the component was created.
	CreationTime *time.Time

	// The name of the component as displayed. If DisplayName isn't specified,
	// TrialComponentName is displayed.
	DisplayName *string

	// When the component ended.
	EndTime *time.Time

	// The input artifacts of the component.
	InputArtifacts map[string]*types.TrialComponentArtifact

	// Who last modified the component.
	LastModifiedBy *types.UserContext

	// When the component was last modified.
	LastModifiedTime *time.Time

	// The metrics for the component.
	Metrics []*types.TrialComponentMetricSummary

	// The output artifacts of the component.
	OutputArtifacts map[string]*types.TrialComponentArtifact

	// The hyperparameters of the component.
	Parameters map[string]*types.TrialComponentParameterValue

	// The Amazon Resource Name (ARN) of the source and, optionally, the job type.
	Source *types.TrialComponentSource

	// When the component started.
	StartTime *time.Time

	// The status of the component. States include:
	//
	//     * InProgress
	//
	//     *
	// Completed
	//
	//     * Failed
	Status *types.TrialComponentStatus

	// The Amazon Resource Name (ARN) of the trial component.
	TrialComponentArn *string

	// The name of the trial component.
	TrialComponentName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeTrialComponentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpDescribeTrialComponent{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpDescribeTrialComponent{})
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
	if err := addOpDescribeTrialComponentValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDescribeTrialComponent(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDescribeTrialComponent(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DescribeTrialComponent",
	}
}
