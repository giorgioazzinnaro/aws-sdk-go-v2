// Code generated by smithy-go-codegen DO NOT EDIT.

package mturk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The AssociateQualificationWithWorker operation gives a Worker a Qualification.
// AssociateQualificationWithWorker does not require that the Worker submit a
// Qualification request. It gives the Qualification directly to the Worker. You
// can only assign a Qualification of a Qualification type that you created (using
// the CreateQualificationType operation). Note: AssociateQualificationWithWorker
// does not affect any pending Qualification requests for the Qualification by the
// Worker. If you assign a Qualification to a Worker, then later grant a
// Qualification request made by the Worker, the granting of the request may modify
// the Qualification score. To resolve a pending Qualification request without
// affecting the Qualification the Worker already has, reject the request with the
// RejectQualificationRequest operation.
func (c *Client) AssociateQualificationWithWorker(ctx context.Context, params *AssociateQualificationWithWorkerInput, optFns ...func(*Options)) (*AssociateQualificationWithWorkerOutput, error) {
	if params == nil {
		params = &AssociateQualificationWithWorkerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateQualificationWithWorker", params, optFns, addOperationAssociateQualificationWithWorkerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateQualificationWithWorkerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateQualificationWithWorkerInput struct {

	// The ID of the Qualification type to use for the assigned Qualification.
	//
	// This member is required.
	QualificationTypeId *string

	// The ID of the Worker to whom the Qualification is being assigned. Worker IDs are
	// included with submitted HIT assignments and Qualification requests.
	//
	// This member is required.
	WorkerId *string

	// The value of the Qualification to assign.
	IntegerValue *int32

	// Specifies whether to send a notification email message to the Worker saying that
	// the qualification was assigned to the Worker. Note: this is true by default.
	SendNotification *bool
}

type AssociateQualificationWithWorkerOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAssociateQualificationWithWorkerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpAssociateQualificationWithWorker{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpAssociateQualificationWithWorker{})
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
	if err := addOpAssociateQualificationWithWorkerValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opAssociateQualificationWithWorker(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opAssociateQualificationWithWorker(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mturk-requester",
		OperationName: "AssociateQualificationWithWorker",
	}
}
