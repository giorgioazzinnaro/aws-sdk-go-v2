// Code generated by smithy-go-codegen DO NOT EDIT.

package robomaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/robomaker/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Describes a world generation job.
func (c *Client) DescribeWorldGenerationJob(ctx context.Context, params *DescribeWorldGenerationJobInput, optFns ...func(*Options)) (*DescribeWorldGenerationJobOutput, error) {
	if params == nil {
		params = &DescribeWorldGenerationJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeWorldGenerationJob", params, optFns, addOperationDescribeWorldGenerationJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeWorldGenerationJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeWorldGenerationJobInput struct {

	// The Amazon Resource Name (arn) of the world generation job to describe.
	//
	// This member is required.
	Job *string
}

type DescribeWorldGenerationJobOutput struct {

	// The Amazon Resource Name (ARN) of the world generation job.
	Arn *string

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request.
	ClientRequestToken *string

	// The time, in milliseconds since the epoch, when the world generation job was
	// created.
	CreatedAt *time.Time

	// The failure code of the world generation job if it failed: InternalServiceError
	// Internal service error. LimitExceeded The requested resource exceeds the maximum
	// number allowed, or the number of concurrent stream requests exceeds the maximum
	// number allowed. ResourceNotFound The specified resource could not be found.
	// RequestThrottled The request was throttled. InvalidInput An input parameter in
	// the request is not valid.
	FailureCode types.WorldGenerationJobErrorCode

	// The reason why the world generation job failed.
	FailureReason *string

	// Summary information about finished worlds.
	FinishedWorldsSummary *types.FinishedWorldsSummary

	// The status of the world generation job: Pending The world generation job request
	// is pending. Running The world generation job is running. Completed The world
	// generation job completed. Failed The world generation job failed. See
	// failureCode for more information. PartialFailed Some worlds did not generate.
	// Canceled The world generation job was cancelled. Canceling The world generation
	// job is being cancelled.
	Status types.WorldGenerationJobStatus

	// A map that contains tag keys and tag values that are attached to the world
	// generation job.
	Tags map[string]*string

	// The Amazon Resource Name (arn) of the world template.
	Template *string

	// Information about the world count.
	WorldCount *types.WorldCount

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeWorldGenerationJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpDescribeWorldGenerationJob{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpDescribeWorldGenerationJob{})
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
	if err := addOpDescribeWorldGenerationJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDescribeWorldGenerationJob(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDescribeWorldGenerationJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "robomaker",
		OperationName: "DescribeWorldGenerationJob",
	}
}
