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

// Provides a list of an experiment's properties.
func (c *Client) DescribeExperiment(ctx context.Context, params *DescribeExperimentInput, optFns ...func(*Options)) (*DescribeExperimentOutput, error) {
	if params == nil {
		params = &DescribeExperimentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeExperiment", params, optFns, addOperationDescribeExperimentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeExperimentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeExperimentInput struct {

	// The name of the experiment to describe.
	//
	// This member is required.
	ExperimentName *string
}

type DescribeExperimentOutput struct {

	// Who created the experiment.
	CreatedBy *types.UserContext

	// When the experiment was created.
	CreationTime *time.Time

	// The description of the experiment.
	Description *string

	// The name of the experiment as displayed. If DisplayName isn't specified,
	// ExperimentName is displayed.
	DisplayName *string

	// The Amazon Resource Name (ARN) of the experiment.
	ExperimentArn *string

	// The name of the experiment.
	ExperimentName *string

	// Who last modified the experiment.
	LastModifiedBy *types.UserContext

	// When the experiment was last modified.
	LastModifiedTime *time.Time

	// The ARN of the source and, optionally, the type.
	Source *types.ExperimentSource

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeExperimentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpDescribeExperiment{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpDescribeExperiment{})
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
	if err := addOpDescribeExperimentValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDescribeExperiment(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDescribeExperiment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DescribeExperiment",
	}
}
