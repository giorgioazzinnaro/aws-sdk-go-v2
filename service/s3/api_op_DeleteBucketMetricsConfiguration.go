// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a metrics configuration for the Amazon CloudWatch request metrics
// (specified by the metrics configuration ID) from the bucket. Note that this
// doesn't include the daily storage metrics. To use this operation, you must have
// permissions to perform the s3:PutMetricsConfiguration action. The bucket owner
// has this permission by default. The bucket owner can grant this permission to
// others. For more information about permissions, see Permissions Related to
// Bucket Subresource Operations
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources)
// and Managing Access Permissions to Your Amazon S3 Resources
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-access-control.html). For
// information about CloudWatch request metrics for Amazon S3, see Monitoring
// Metrics with Amazon CloudWatch
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/cloudwatch-monitoring.html).
// The following operations are related to DeleteBucketMetricsConfiguration:
//
// *
// GetBucketMetricsConfiguration
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketMetricsConfiguration.html)
//
// *
// PutBucketMetricsConfiguration
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketMetricsConfiguration.html)
//
// *
// ListBucketMetricsConfigurations
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListBucketMetricsConfigurations.html)
//
// *
// Monitoring Metrics with Amazon CloudWatch
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/cloudwatch-monitoring.html)
func (c *Client) DeleteBucketMetricsConfiguration(ctx context.Context, params *DeleteBucketMetricsConfigurationInput, optFns ...func(*Options)) (*DeleteBucketMetricsConfigurationOutput, error) {
	if params == nil {
		params = &DeleteBucketMetricsConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteBucketMetricsConfiguration", params, optFns, addOperationDeleteBucketMetricsConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteBucketMetricsConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteBucketMetricsConfigurationInput struct {

	// The name of the bucket containing the metrics configuration to delete.
	//
	// This member is required.
	Bucket *string

	// The ID used to identify the metrics configuration.
	//
	// This member is required.
	Id *string

	// The account id of the expected bucket owner. If the bucket is owned by a
	// different account, the request will fail with an HTTP 403 (Access Denied) error.
	ExpectedBucketOwner *string
}

type DeleteBucketMetricsConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteBucketMetricsConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpDeleteBucketMetricsConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpDeleteBucketMetricsConfiguration{}, middleware.After)
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
	addOpDeleteBucketMetricsConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteBucketMetricsConfiguration(options.Region), middleware.Before)
	addMetadataRetrieverMiddleware(stack)
	addUpdateEndpointMiddleware(stack, options)
	addResponseErrorMiddleware(stack)
	v4.AddContentSHA256HeaderMiddleware(stack)
	disableAcceptEncodingGzip(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteBucketMetricsConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "DeleteBucketMetricsConfiguration",
	}
}
