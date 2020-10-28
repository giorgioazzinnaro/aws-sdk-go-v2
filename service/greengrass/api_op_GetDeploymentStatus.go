// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrass

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the status of a deployment.
func (c *Client) GetDeploymentStatus(ctx context.Context, params *GetDeploymentStatusInput, optFns ...func(*Options)) (*GetDeploymentStatusOutput, error) {
	if params == nil {
		params = &GetDeploymentStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDeploymentStatus", params, optFns, addOperationGetDeploymentStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDeploymentStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDeploymentStatusInput struct {

	// The ID of the deployment.
	//
	// This member is required.
	DeploymentId *string

	// The ID of the Greengrass group.
	//
	// This member is required.
	GroupId *string
}

type GetDeploymentStatusOutput struct {

	// The status of the deployment: ''InProgress'', ''Building'', ''Success'', or
	// ''Failure''.
	DeploymentStatus *string

	// The type of the deployment.
	DeploymentType types.DeploymentType

	// Error details
	ErrorDetails []*types.ErrorDetail

	// Error message
	ErrorMessage *string

	// The time, in milliseconds since the epoch, when the deployment status was
	// updated.
	UpdatedAt *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetDeploymentStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpGetDeploymentStatus{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpGetDeploymentStatus{})
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
	if err := addOpGetDeploymentStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opGetDeploymentStatus(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opGetDeploymentStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "GetDeploymentStatus",
	}
}
