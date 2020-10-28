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

// Sets the cors configuration for your bucket. If the configuration exists, Amazon
// S3 replaces it. To use this operation, you must be allowed to perform the
// s3:PutBucketCORS action. By default, the bucket owner has this permission and
// can grant it to others. You set this configuration on a bucket so that the
// bucket can service cross-origin requests. For example, you might want to enable
// a request whose origin is http://www.example.com to access your Amazon S3 bucket
// at my.example.bucket.com by using the browser's XMLHttpRequest capability. To
// enable cross-origin resource sharing (CORS) on a bucket, you add the cors
// subresource to the bucket. The cors subresource is an XML document in which you
// configure rules that identify origins and the HTTP methods that can be executed
// on your bucket. The document is limited to 64 KB in size. When Amazon S3
// receives a cross-origin request (or a pre-flight OPTIONS request) against a
// bucket, it evaluates the cors configuration on the bucket and uses the first
// CORSRule rule that matches the incoming browser request to enable a cross-origin
// request. For a rule to match, the following conditions must be met:
//
//     * The
// request's Origin header must match AllowedOrigin elements.
//
//     * The request
// method (for example, GET, PUT, HEAD, and so on) or the
// Access-Control-Request-Method header in case of a pre-flight OPTIONS request
// must be one of the AllowedMethod elements.
//
//     * Every header specified in the
// Access-Control-Request-Headers request header of a pre-flight request must match
// an AllowedHeader element.
//
// For more information about CORS, go to Enabling
// Cross-Origin Resource Sharing
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) in the Amazon Simple
// Storage Service Developer Guide. Related Resources
//
//     * GetBucketCors
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketCors.html)
//
//     *
// DeleteBucketCors
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketCors.html)
//
//
// * RESTOPTIONSobject
// (https://docs.aws.amazon.com/AmazonS3/latest/API/RESTOPTIONSobject.html)
func (c *Client) PutBucketCors(ctx context.Context, params *PutBucketCorsInput, optFns ...func(*Options)) (*PutBucketCorsOutput, error) {
	if params == nil {
		params = &PutBucketCorsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutBucketCors", params, optFns, addOperationPutBucketCorsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutBucketCorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutBucketCorsInput struct {

	// Specifies the bucket impacted by the corsconfiguration.
	//
	// This member is required.
	Bucket *string

	// Describes the cross-origin access configuration for objects in an Amazon S3
	// bucket. For more information, see Enabling Cross-Origin Resource Sharing
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) in the Amazon Simple
	// Storage Service Developer Guide.
	//
	// This member is required.
	CORSConfiguration *types.CORSConfiguration

	// The base64-encoded 128-bit MD5 digest of the data. This header must be used as a
	// message integrity check to verify that the request body was not corrupted in
	// transit. For more information, go to RFC 1864.
	// (http://www.ietf.org/rfc/rfc1864.txt)
	ContentMD5 *string

	// The account id of the expected bucket owner. If the bucket is owned by a
	// different account, the request will fail with an HTTP 403 (Access Denied) error.
	ExpectedBucketOwner *string
}

type PutBucketCorsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutBucketCorsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestxml_serializeOpPutBucketCors{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestxml_deserializeOpPutBucketCors{})
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
	if err := addOpPutBucketCorsValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opPutBucketCors(options.Region)); err != nil {
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
	if err := smithyhttp.AddContentChecksumMiddleware(stack); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opPutBucketCors(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "PutBucketCors",
	}
}
