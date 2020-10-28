// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kendra/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Describes an existing Amazon Kendra index
func (c *Client) DescribeIndex(ctx context.Context, params *DescribeIndexInput, optFns ...func(*Options)) (*DescribeIndexOutput, error) {
	if params == nil {
		params = &DescribeIndexInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeIndex", params, optFns, addOperationDescribeIndexMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeIndexOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeIndexInput struct {

	// The name of the index to describe.
	//
	// This member is required.
	Id *string
}

type DescribeIndexOutput struct {

	// For enterprise edtion indexes, you can choose to use additional capacity to meet
	// the needs of your application. This contains the capacity units used for the
	// index. A 0 for the query capacity or the storage capacity indicates that the
	// index is using the default capacity for the index.
	CapacityUnits *types.CapacityUnitsConfiguration

	// The Unix datetime that the index was created.
	CreatedAt *time.Time

	// The description of the index.
	Description *string

	// Configuration settings for any metadata applied to the documents in the index.
	DocumentMetadataConfigurations []*types.DocumentMetadataConfiguration

	// The Amazon Kendra edition used for the index. You decide the edition when you
	// create the index.
	Edition types.IndexEdition

	// When th eStatus field value is FAILED, the ErrorMessage field contains a message
	// that explains why.
	ErrorMessage *string

	// the name of the index.
	Id *string

	// Provides information about the number of FAQ questions and answers and the
	// number of text documents indexed.
	IndexStatistics *types.IndexStatistics

	// The name of the index.
	Name *string

	// The Amazon Resource Name (ARN) of the IAM role that gives Amazon Kendra
	// permission to write to your Amazon Cloudwatch logs.
	RoleArn *string

	// The identifier of the AWS KMS customer master key (CMK) used to encrypt your
	// data. Amazon Kendra doesn't support asymmetric CMKs.
	ServerSideEncryptionConfiguration *types.ServerSideEncryptionConfiguration

	// The current status of the index. When the value is ACTIVE, the index is ready
	// for use. If the Status field value is FAILED, the ErrorMessage field contains a
	// message that explains why.
	Status types.IndexStatus

	// The Unix datetime that the index was last updated.
	UpdatedAt *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeIndexMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpDescribeIndex{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpDescribeIndex{})
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
	if err := addOpDescribeIndexValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDescribeIndex(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDescribeIndex(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kendra",
		OperationName: "DescribeIndex",
	}
}
