// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates one or more values for an SSM document.
func (c *Client) UpdateDocument(ctx context.Context, params *UpdateDocumentInput, optFns ...func(*Options)) (*UpdateDocumentOutput, error) {
	if params == nil {
		params = &UpdateDocumentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDocument", params, optFns, addOperationUpdateDocumentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDocumentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDocumentInput struct {

	// A valid JSON or YAML string.
	//
	// This member is required.
	Content *string

	// The name of the document that you want to update.
	//
	// This member is required.
	Name *string

	// A list of key and value pairs that describe attachments to a version of a
	// document.
	Attachments []*types.AttachmentsSource

	// Specify the document format for the new document version. Systems Manager
	// supports JSON and YAML documents. JSON is the default format.
	DocumentFormat types.DocumentFormat

	// (Required) The latest version of the document that you want to update. The
	// latest document version can be specified using the $LATEST variable or by the
	// version number. Updating a previous version of a document is not supported.
	DocumentVersion *string

	// Specify a new target type for the document.
	TargetType *string

	// An optional field specifying the version of the artifact you are updating with
	// the document. For example, "Release 12, Update 6". This value is unique across
	// all versions of a document, and cannot be changed.
	VersionName *string
}

type UpdateDocumentOutput struct {

	// A description of the document that was updated.
	DocumentDescription *types.DocumentDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateDocumentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpUpdateDocument{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpUpdateDocument{})
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
	if err := addOpUpdateDocumentValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opUpdateDocument(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opUpdateDocument(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "UpdateDocument",
	}
}
