// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Starts asynchronous detection of faces in a stored video. Amazon Rekognition
// Video can detect faces in a video stored in an Amazon S3 bucket. Use Video to
// specify the bucket name and the filename of the video. StartFaceDetection
// returns a job identifier (JobId) that you use to get the results of the
// operation. When face detection is finished, Amazon Rekognition Video publishes a
// completion status to the Amazon Simple Notification Service topic that you
// specify in NotificationChannel. To get the results of the face detection
// operation, first check that the status value published to the Amazon SNS topic
// is SUCCEEDED. If so, call GetFaceDetection and pass the job identifier (JobId)
// from the initial call to StartFaceDetection. For more information, see Detecting
// Faces in a Stored Video in the Amazon Rekognition Developer Guide.
func (c *Client) StartFaceDetection(ctx context.Context, params *StartFaceDetectionInput, optFns ...func(*Options)) (*StartFaceDetectionOutput, error) {
	if params == nil {
		params = &StartFaceDetectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartFaceDetection", params, optFns, addOperationStartFaceDetectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartFaceDetectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartFaceDetectionInput struct {

	// The video in which you want to detect faces. The video must be stored in an
	// Amazon S3 bucket.
	//
	// This member is required.
	Video *types.Video

	// Idempotent token used to identify the start request. If you use the same token
	// with multiple StartFaceDetection requests, the same JobId is returned. Use
	// ClientRequestToken to prevent the same job from being accidently started more
	// than once.
	ClientRequestToken *string

	// The face attributes you want returned. DEFAULT - The following subset of facial
	// attributes are returned: BoundingBox, Confidence, Pose, Quality and Landmarks.
	// ALL - All facial attributes are returned.
	FaceAttributes types.FaceAttributes

	// An identifier you specify that's returned in the completion notification that's
	// published to your Amazon Simple Notification Service topic. For example, you can
	// use JobTag to group related jobs and identify them in the completion
	// notification.
	JobTag *string

	// The ARN of the Amazon SNS topic to which you want Amazon Rekognition Video to
	// publish the completion status of the face detection operation.
	NotificationChannel *types.NotificationChannel
}

type StartFaceDetectionOutput struct {

	// The identifier for the face detection job. Use JobId to identify the job in a
	// subsequent call to GetFaceDetection.
	JobId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStartFaceDetectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpStartFaceDetection{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpStartFaceDetection{})
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
	if err := addOpStartFaceDetectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opStartFaceDetection(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opStartFaceDetection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "StartFaceDetection",
	}
}
