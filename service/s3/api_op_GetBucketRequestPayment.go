// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the request payment configuration of a bucket. To use this version of
// the operation, you must be the bucket owner. For more information, see Requester
// Pays Buckets
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/RequesterPaysBuckets.html). The
// following operations are related to GetBucketRequestPayment:
//
//     * ListObjects
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListObjects.html)
func (c *Client) GetBucketRequestPayment(ctx context.Context, params *GetBucketRequestPaymentInput, optFns ...func(*Options)) (*GetBucketRequestPaymentOutput, error) {
	if params == nil {
		params = &GetBucketRequestPaymentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBucketRequestPayment", params, optFns, addOperationGetBucketRequestPaymentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBucketRequestPaymentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBucketRequestPaymentInput struct {

	// The name of the bucket for which to get the payment request configuration
	//
	// This member is required.
	Bucket *string

	// The account id of the expected bucket owner. If the bucket is owned by a
	// different account, the request will fail with an HTTP 403 (Access Denied) error.
	ExpectedBucketOwner *string
}

type GetBucketRequestPaymentOutput struct {

	// Specifies who pays for the download and request fees.
	Payer types.Payer

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetBucketRequestPaymentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestxml_serializeOpGetBucketRequestPayment{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestxml_deserializeOpGetBucketRequestPayment{})
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
	if err := addOpGetBucketRequestPaymentValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opGetBucketRequestPayment(options.Region)); err != nil {
		return err
	}
	if err := addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err := addUpdateEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err := addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err := v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err := disableAcceptEncodingGzip(stack); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetBucketRequestPayment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "GetBucketRequestPayment",
	}
}
