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

// Returns the default encryption configuration for an Amazon S3 bucket. For
// information about the Amazon S3 default encryption feature, see Amazon S3
// Default Bucket Encryption
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html). To use
// this operation, you must have permission to perform the
// s3:GetEncryptionConfiguration action. The bucket owner has this permission by
// default. The bucket owner can grant this permission to others. For more
// information about permissions, see Permissions Related to Bucket Subresource
// Operations
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources)
// and Managing Access Permissions to Your Amazon S3 Resources
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-access-control.html). The
// following operations are related to GetBucketEncryption:
//
//     *
// PutBucketEncryption
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketEncryption.html)
//
//
// * DeleteBucketEncryption
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketEncryption.html)
func (c *Client) GetBucketEncryption(ctx context.Context, params *GetBucketEncryptionInput, optFns ...func(*Options)) (*GetBucketEncryptionOutput, error) {
	if params == nil {
		params = &GetBucketEncryptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBucketEncryption", params, optFns, addOperationGetBucketEncryptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBucketEncryptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBucketEncryptionInput struct {

	// The name of the bucket from which the server-side encryption configuration is
	// retrieved.
	//
	// This member is required.
	Bucket *string

	// The account id of the expected bucket owner. If the bucket is owned by a
	// different account, the request will fail with an HTTP 403 (Access Denied) error.
	ExpectedBucketOwner *string
}

type GetBucketEncryptionOutput struct {

	// Specifies the default server-side-encryption configuration.
	ServerSideEncryptionConfiguration *types.ServerSideEncryptionConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetBucketEncryptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestxml_serializeOpGetBucketEncryption{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestxml_deserializeOpGetBucketEncryption{})
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
	if err := addOpGetBucketEncryptionValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opGetBucketEncryption(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opGetBucketEncryption(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "GetBucketEncryption",
	}
}
