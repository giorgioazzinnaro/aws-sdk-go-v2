// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/personalize/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a batch inference job. The operation can handle up to 50 million records
// and the input file must be in JSON format. For more information, see
// recommendations-batch.
func (c *Client) CreateBatchInferenceJob(ctx context.Context, params *CreateBatchInferenceJobInput, optFns ...func(*Options)) (*CreateBatchInferenceJobOutput, error) {
	if params == nil {
		params = &CreateBatchInferenceJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateBatchInferenceJob", params, optFns, addOperationCreateBatchInferenceJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateBatchInferenceJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateBatchInferenceJobInput struct {

	// The Amazon S3 path that leads to the input file to base your recommendations on.
	// The input material must be in JSON format.
	//
	// This member is required.
	JobInput *types.BatchInferenceJobInput

	// The name of the batch inference job to create.
	//
	// This member is required.
	JobName *string

	// The path to the Amazon S3 bucket where the job's output will be stored.
	//
	// This member is required.
	JobOutput *types.BatchInferenceJobOutput

	// The ARN of the Amazon Identity and Access Management role that has permissions
	// to read and write to your input and out Amazon S3 buckets respectively.
	//
	// This member is required.
	RoleArn *string

	// The Amazon Resource Name (ARN) of the solution version that will be used to
	// generate the batch inference recommendations.
	//
	// This member is required.
	SolutionVersionArn *string

	// The configuration details of a batch inference job.
	BatchInferenceJobConfig *types.BatchInferenceJobConfig

	// The ARN of the filter to apply to the batch inference job. For more information
	// on using filters, see Using Filters with Amazon Personalize.
	FilterArn *string

	// The number of recommendations to retreive.
	NumResults *int32
}

type CreateBatchInferenceJobOutput struct {

	// The ARN of the batch inference job.
	BatchInferenceJobArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateBatchInferenceJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpCreateBatchInferenceJob{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpCreateBatchInferenceJob{})
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
	if err := addOpCreateBatchInferenceJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opCreateBatchInferenceJob(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opCreateBatchInferenceJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "CreateBatchInferenceJob",
	}
}
