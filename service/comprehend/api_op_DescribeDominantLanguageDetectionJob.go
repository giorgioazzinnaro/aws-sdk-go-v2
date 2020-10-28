// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehend

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/comprehend/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the properties associated with a dominant language detection job. Use this
// operation to get the status of a detection job.
func (c *Client) DescribeDominantLanguageDetectionJob(ctx context.Context, params *DescribeDominantLanguageDetectionJobInput, optFns ...func(*Options)) (*DescribeDominantLanguageDetectionJobOutput, error) {
	if params == nil {
		params = &DescribeDominantLanguageDetectionJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDominantLanguageDetectionJob", params, optFns, addOperationDescribeDominantLanguageDetectionJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDominantLanguageDetectionJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDominantLanguageDetectionJobInput struct {

	// The identifier that Amazon Comprehend generated for the job. The operation
	// returns this identifier in its response.
	//
	// This member is required.
	JobId *string
}

type DescribeDominantLanguageDetectionJobOutput struct {

	// An object that contains the properties associated with a dominant language
	// detection job.
	DominantLanguageDetectionJobProperties *types.DominantLanguageDetectionJobProperties

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeDominantLanguageDetectionJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpDescribeDominantLanguageDetectionJob{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpDescribeDominantLanguageDetectionJob{})
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
	if err := addOpDescribeDominantLanguageDetectionJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDescribeDominantLanguageDetectionJob(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDescribeDominantLanguageDetectionJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehend",
		OperationName: "DescribeDominantLanguageDetectionJob",
	}
}
