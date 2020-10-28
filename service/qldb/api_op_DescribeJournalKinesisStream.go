// Code generated by smithy-go-codegen DO NOT EDIT.

package qldb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/qldb/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns detailed information about a given Amazon QLDB journal stream. The
// output includes the Amazon Resource Name (ARN), stream name, current status,
// creation time, and the parameters of your original stream creation request.
func (c *Client) DescribeJournalKinesisStream(ctx context.Context, params *DescribeJournalKinesisStreamInput, optFns ...func(*Options)) (*DescribeJournalKinesisStreamOutput, error) {
	if params == nil {
		params = &DescribeJournalKinesisStreamInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeJournalKinesisStream", params, optFns, addOperationDescribeJournalKinesisStreamMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeJournalKinesisStreamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeJournalKinesisStreamInput struct {

	// The name of the ledger.
	//
	// This member is required.
	LedgerName *string

	// The unique ID that QLDB assigns to each QLDB journal stream.
	//
	// This member is required.
	StreamId *string
}

type DescribeJournalKinesisStreamOutput struct {

	// Information about the QLDB journal stream returned by a DescribeJournalS3Export
	// request.
	Stream *types.JournalKinesisStreamDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeJournalKinesisStreamMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpDescribeJournalKinesisStream{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpDescribeJournalKinesisStream{})
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
	if err := addOpDescribeJournalKinesisStreamValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDescribeJournalKinesisStream(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDescribeJournalKinesisStream(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "qldb",
		OperationName: "DescribeJournalKinesisStream",
	}
}
