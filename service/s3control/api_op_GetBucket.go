// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"strings"
	"time"
)

// Gets an Amazon S3 on Outposts bucket. For more information, see  Using Amazon S3
// on Outposts (https://docs.aws.amazon.com/AmazonS3/latest/dev/S3onOutposts.html)
// in the Amazon Simple Storage Service Developer Guide. The following actions are
// related to GetBucket for Amazon S3 on Outposts:
//
//     * PutObject
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObject.html)
//
//     *
// CreateBucket
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API__control_CreateBucket.html)
//
//
// * DeleteBucket
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API__control_DeleteBucket.html)
func (c *Client) GetBucket(ctx context.Context, params *GetBucketInput, optFns ...func(*Options)) (*GetBucketOutput, error) {
	if params == nil {
		params = &GetBucketInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBucket", params, optFns, addOperationGetBucketMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBucketOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBucketInput struct {

	// The AWS account ID of the Outposts bucket.
	//
	// This member is required.
	AccountId *string

	// The ARN of the bucket. For Amazon S3 on Outposts specify the ARN of the bucket
	// accessed in the format arn:aws:s3-outposts:::outpost//bucket/. For example, to
	// access the bucket reports through outpost my-outpost owned by account
	// 123456789012 in Region us-west-2, use the URL encoding of
	// arn:aws:s3-outposts:us-west-2:123456789012:outpost/my-outpost/bucket/reports.
	// The value must be URL encoded.
	//
	// This member is required.
	Bucket *string
}

type GetBucketOutput struct {

	// The Outposts bucket requested.
	Bucket *string

	// The creation date of the Outposts bucket.
	CreationDate *time.Time

	//
	PublicAccessBlockEnabled *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetBucketMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestxml_serializeOpGetBucket{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestxml_deserializeOpGetBucket{})
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
	if err := addEndpointPrefix_opGetBucketMiddleware(stack); err != nil {
		return err
	}
	if err := addOpGetBucketValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opGetBucket(options.Region)); err != nil {
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
	return nil
}

type endpointPrefix_opGetBucketMiddleware struct {
}

func (*endpointPrefix_opGetBucketMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGetBucketMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	input, ok := in.Parameters.(*GetBucketInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input type %T", in.Parameters)
	}

	var prefix strings.Builder
	if input.AccountId == nil {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so may not be nil")}
	} else if !smithyhttp.ValidHostLabel(*input.AccountId) {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so must match \"[a-zA-Z0-9-]{1,63}\", but was \"%s\"", *input.AccountId)}
	} else {
		prefix.WriteString(*input.AccountId)
	}
	prefix.WriteString(".")
	req.HostPrefix = prefix.String()

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opGetBucketMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert("OperationSerializer", middleware.Before, &endpointPrefix_opGetBucketMiddleware{})
}

func newServiceMetadataMiddleware_opGetBucket(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "GetBucket",
	}
}
