// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Associates multiple self-service actions with provisioning artifacts.
func (c *Client) BatchAssociateServiceActionWithProvisioningArtifact(ctx context.Context, params *BatchAssociateServiceActionWithProvisioningArtifactInput, optFns ...func(*Options)) (*BatchAssociateServiceActionWithProvisioningArtifactOutput, error) {
	if params == nil {
		params = &BatchAssociateServiceActionWithProvisioningArtifactInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchAssociateServiceActionWithProvisioningArtifact", params, optFns, addOperationBatchAssociateServiceActionWithProvisioningArtifactMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchAssociateServiceActionWithProvisioningArtifactOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchAssociateServiceActionWithProvisioningArtifactInput struct {

	// One or more associations, each consisting of the Action ID, the Product ID, and
	// the Provisioning Artifact ID.
	//
	// This member is required.
	ServiceActionAssociations []*types.ServiceActionAssociation

	// The language code.
	//
	// * en - English (default)
	//
	// * jp - Japanese
	//
	// * zh - Chinese
	AcceptLanguage *string
}

type BatchAssociateServiceActionWithProvisioningArtifactOutput struct {

	// An object that contains a list of errors, along with information to help you
	// identify the self-service action.
	FailedServiceActionAssociations []*types.FailedServiceActionAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationBatchAssociateServiceActionWithProvisioningArtifactMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpBatchAssociateServiceActionWithProvisioningArtifact{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchAssociateServiceActionWithProvisioningArtifact{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpBatchAssociateServiceActionWithProvisioningArtifactValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchAssociateServiceActionWithProvisioningArtifact(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opBatchAssociateServiceActionWithProvisioningArtifact(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "BatchAssociateServiceActionWithProvisioningArtifact",
	}
}
