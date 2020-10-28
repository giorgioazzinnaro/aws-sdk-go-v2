// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kendra/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new Amazon Kendra index. Index creation is an asynchronous operation.
// To determine if index creation has completed, check the Status field returned
// from a call to . The Status field is set to ACTIVE when the index is ready to
// use. Once the index is active you can index your documents using the operation
// or using one of the supported data sources.
func (c *Client) CreateIndex(ctx context.Context, params *CreateIndexInput, optFns ...func(*Options)) (*CreateIndexOutput, error) {
	if params == nil {
		params = &CreateIndexInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateIndex", params, optFns, addOperationCreateIndexMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateIndexOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateIndexInput struct {

	// The name for the new index.
	//
	// This member is required.
	Name *string

	// An AWS Identity and Access Management (IAM) role that gives Amazon Kendra
	// permissions to access your Amazon CloudWatch logs and metrics. This is also the
	// role used when you use the BatchPutDocument operation to index documents from an
	// Amazon S3 bucket.
	//
	// This member is required.
	RoleArn *string

	// A token that you provide to identify the request to create an index. Multiple
	// calls to the CreateIndex operation with the same client token will create only
	// one index.
	ClientToken *string

	// A description for the index.
	Description *string

	// The Amazon Kendra edition to use for the index. Choose DEVELOPER_EDITION for
	// indexes intended for development, testing, or proof of concept. Use
	// ENTERPRISE_EDITION for your production databases. Once you set the edition for
	// an index, it can't be changed. The Edition parameter is optional. If you don't
	// supply a value, the default is ENTERPRISE_EDITION.
	Edition types.IndexEdition

	// The identifier of the AWS KMS customer managed key (CMK) to use to encrypt data
	// indexed by Amazon Kendra. Amazon Kendra doesn't support asymmetric CMKs.
	ServerSideEncryptionConfiguration *types.ServerSideEncryptionConfiguration

	// A list of key-value pairs that identify the index. You can use the tags to
	// identify and organize your resources and to control access to resources.
	Tags []*types.Tag
}

type CreateIndexOutput struct {

	// The unique identifier of the index. Use this identifier when you query an index,
	// set up a data source, or index a document.
	Id *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateIndexMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpCreateIndex{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpCreateIndex{})
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
	if err := addIdempotencyToken_opCreateIndexMiddleware(stack, options); err != nil {
		return err
	}
	if err := addOpCreateIndexValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opCreateIndex(options.Region)); err != nil {
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

type idempotencyToken_initializeOpCreateIndex struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateIndex) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateIndex) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateIndexInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateIndexInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateIndexMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(middleware.Before, &idempotencyToken_initializeOpCreateIndex{tokenProvider: cfg.IdempotencyTokenProvider})
}

func newServiceMetadataMiddleware_opCreateIndex(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kendra",
		OperationName: "CreateIndex",
	}
}
