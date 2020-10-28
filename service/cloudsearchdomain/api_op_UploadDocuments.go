// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudsearchdomain

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudsearchdomain/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"io"
)

// Posts a batch of documents to a search domain for indexing. A document batch is
// a collection of add and delete operations that represent the documents you want
// to add, update, or delete from your domain. Batches can be described in either
// JSON or XML. Each item that you want Amazon CloudSearch to return as a search
// result (such as a product) is represented as a document. Every document has a
// unique ID and one or more fields that contain the data that you want to search
// and return in results. Individual documents cannot contain more than 1 MB of
// data. The entire batch cannot exceed 5 MB. To get the best possible upload
// performance, group add and delete operations in batches that are close the 5 MB
// limit. Submitting a large volume of single-document batches can overload a
// domain's document service. The endpoint for submitting UploadDocuments requests
// is domain-specific and requires the --endpoint-url option. To get the document
// endpoint for your domain, use the Amazon CloudSearch configuration service
// DescribeDomains action. The endpoints are also available on the domain dashboard
// in the Amazon CloudSearch console. For more information about formatting your
// data for Amazon CloudSearch, see Preparing Your Data
// (http://docs.aws.amazon.com/cloudsearch/latest/developerguide/preparing-data.html)
// in the Amazon CloudSearch Developer Guide. For more information about uploading
// data for indexing, see Uploading Data
// (http://docs.aws.amazon.com/cloudsearch/latest/developerguide/uploading-data.html)
// in the Amazon CloudSearch Developer Guide.
func (c *Client) UploadDocuments(ctx context.Context, params *UploadDocumentsInput, optFns ...func(*Options)) (*UploadDocumentsOutput, error) {
	if params == nil {
		params = &UploadDocumentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UploadDocuments", params, optFns, addOperationUploadDocumentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UploadDocumentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the UploadDocuments request.
type UploadDocumentsInput struct {

	// The format of the batch you are uploading. Amazon CloudSearch supports two
	// document batch formats:
	//
	//     * application/json
	//
	//     * application/xml
	//
	// This member is required.
	ContentType types.ContentType

	// A batch of documents formatted in JSON or HTML.
	//
	// This member is required.
	Documents io.Reader
}

// Contains the response to an UploadDocuments request.
type UploadDocumentsOutput struct {

	// The number of documents that were added to the search domain.
	Adds *int64

	// The number of documents that were deleted from the search domain.
	Deletes *int64

	// The status of an UploadDocumentsRequest.
	Status *string

	// Any warnings returned by the document service about the documents being
	// uploaded.
	Warnings []*types.DocumentServiceWarning

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUploadDocumentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpUploadDocuments{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpUploadDocuments{})
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
	if err := addOpUploadDocumentsValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opUploadDocuments(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opUploadDocuments(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudsearch",
		OperationName: "UploadDocuments",
	}
}
