// Code generated by smithy-go-codegen DO NOT EDIT.

package frauddetector

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the status of a model version. You can perform the following status
// updates:
//
//     * Change the TRAINING_COMPLETE status to ACTIVE.
//
//     * Change
// ACTIVEto INACTIVE.
func (c *Client) UpdateModelVersionStatus(ctx context.Context, params *UpdateModelVersionStatusInput, optFns ...func(*Options)) (*UpdateModelVersionStatusOutput, error) {
	if params == nil {
		params = &UpdateModelVersionStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateModelVersionStatus", params, optFns, addOperationUpdateModelVersionStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateModelVersionStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateModelVersionStatusInput struct {

	// The model ID of the model version to update.
	//
	// This member is required.
	ModelId *string

	// The model type.
	//
	// This member is required.
	ModelType types.ModelTypeEnum

	// The model version number.
	//
	// This member is required.
	ModelVersionNumber *string

	// The model version status.
	//
	// This member is required.
	Status types.ModelVersionStatus
}

type UpdateModelVersionStatusOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateModelVersionStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpUpdateModelVersionStatus{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpUpdateModelVersionStatus{})
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
	if err := addOpUpdateModelVersionStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opUpdateModelVersionStatus(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opUpdateModelVersionStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "frauddetector",
		OperationName: "UpdateModelVersionStatus",
	}
}
