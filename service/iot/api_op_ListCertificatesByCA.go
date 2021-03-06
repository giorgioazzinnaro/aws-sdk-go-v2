// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// List the device certificates signed by the specified CA certificate.
func (c *Client) ListCertificatesByCA(ctx context.Context, params *ListCertificatesByCAInput, optFns ...func(*Options)) (*ListCertificatesByCAOutput, error) {
	if params == nil {
		params = &ListCertificatesByCAInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCertificatesByCA", params, optFns, addOperationListCertificatesByCAMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCertificatesByCAOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input to the ListCertificatesByCA operation.
type ListCertificatesByCAInput struct {

	// The ID of the CA certificate. This operation will list all registered device
	// certificate that were signed by this CA certificate.
	//
	// This member is required.
	CaCertificateId *string

	// Specifies the order for results. If True, the results are returned in ascending
	// order, based on the creation date.
	AscendingOrder *bool

	// The marker for the next set of results.
	Marker *string

	// The result page size.
	PageSize *int32
}

// The output of the ListCertificatesByCA operation.
type ListCertificatesByCAOutput struct {

	// The device certificates signed by the specified CA certificate.
	Certificates []*types.Certificate

	// The marker for the next set of results, or null if there are no additional
	// results.
	NextMarker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListCertificatesByCAMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListCertificatesByCA{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListCertificatesByCA{}, middleware.After)
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
	addOpListCertificatesByCAValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListCertificatesByCA(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListCertificatesByCA(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ListCertificatesByCA",
	}
}
