// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3control/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"strings"
)

// Returns a list of the access points currently associated with the specified
// bucket. You can retrieve up to 1000 access points per call. If the specified
// bucket has more than 1,000 access points (or the number specified in maxResults,
// whichever is less), the response will include a continuation token that you can
// use to list the additional access points. All Amazon S3 on Outposts REST API
// requests for this action require an additional parameter of outpost-id to be
// passed with the request and an S3 on Outposts endpoint hostname prefix instead
// of s3-control. For an example of the request syntax for Amazon S3 on Outposts
// that uses the S3 on Outposts endpoint hostname prefix and the outpost-id derived
// using the access point ARN, see the  Example
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API__control_GetAccessPoint.html#API_control_GetAccessPoint_Examples)
// section below. The following actions are related to ListAccessPoints:
//
//     *
// CreateAccessPoint
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateAccessPoint.html)
//
//
// * DeleteAccessPoint
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteAccessPoint.html)
//
//
// * GetAccessPoint
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetAccessPoint.html)
func (c *Client) ListAccessPoints(ctx context.Context, params *ListAccessPointsInput, optFns ...func(*Options)) (*ListAccessPointsOutput, error) {
	if params == nil {
		params = &ListAccessPointsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAccessPoints", params, optFns, addOperationListAccessPointsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAccessPointsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAccessPointsInput struct {

	// The AWS account ID for owner of the bucket whose access points you want to list.
	//
	// This member is required.
	AccountId *string

	// The name of the bucket whose associated access points you want to list. For
	// Amazon S3 on Outposts specify the ARN of the bucket accessed in the format
	// arn:aws:s3-outposts:::outpost//bucket/. For example, to access the bucket
	// reports through outpost my-outpost owned by account 123456789012 in Region
	// us-west-2, use the URL encoding of
	// arn:aws:s3-outposts:us-west-2:123456789012:outpost/my-outpost/bucket/reports.
	// The value must be URL encoded.
	Bucket *string

	// The maximum number of access points that you want to include in the list. If the
	// specified bucket has more than this number of access points, then the response
	// will include a continuation token in the NextToken field that you can use to
	// retrieve the next page of access points.
	MaxResults *int32

	// A continuation token. If a previous call to ListAccessPoints returned a
	// continuation token in the NextToken field, then providing that value here causes
	// Amazon S3 to retrieve the next page of results.
	NextToken *string
}

type ListAccessPointsOutput struct {

	// Contains identification and configuration information for one or more access
	// points associated with the specified bucket.
	AccessPointList []*types.AccessPoint

	// If the specified bucket has more access points than can be returned in one call
	// to this API, this field contains a continuation token that you can provide in
	// subsequent calls to this API to retrieve additional access points.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListAccessPointsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestxml_serializeOpListAccessPoints{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestxml_deserializeOpListAccessPoints{})
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
	if err := addEndpointPrefix_opListAccessPointsMiddleware(stack); err != nil {
		return err
	}
	if err := addOpListAccessPointsValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opListAccessPoints(options.Region)); err != nil {
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

type endpointPrefix_opListAccessPointsMiddleware struct {
}

func (*endpointPrefix_opListAccessPointsMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListAccessPointsMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	input, ok := in.Parameters.(*ListAccessPointsInput)
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
func addEndpointPrefix_opListAccessPointsMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert("OperationSerializer", middleware.Before, &endpointPrefix_opListAccessPointsMiddleware{})
}

func newServiceMetadataMiddleware_opListAccessPoints(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "ListAccessPoints",
	}
}
