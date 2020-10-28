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

// Returns an array of all Amazon QLDB journal stream descriptors for a given
// ledger. The output of each stream descriptor includes the same details that are
// returned by DescribeJournalKinesisStream. This action returns a maximum of
// MaxResults items. It is paginated so that you can retrieve all the items by
// calling ListJournalKinesisStreamsForLedger multiple times.
func (c *Client) ListJournalKinesisStreamsForLedger(ctx context.Context, params *ListJournalKinesisStreamsForLedgerInput, optFns ...func(*Options)) (*ListJournalKinesisStreamsForLedgerOutput, error) {
	if params == nil {
		params = &ListJournalKinesisStreamsForLedgerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListJournalKinesisStreamsForLedger", params, optFns, addOperationListJournalKinesisStreamsForLedgerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListJournalKinesisStreamsForLedgerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListJournalKinesisStreamsForLedgerInput struct {

	// The name of the ledger.
	//
	// This member is required.
	LedgerName *string

	// The maximum number of results to return in a single
	// ListJournalKinesisStreamsForLedger request. (The actual number of results
	// returned might be fewer.)
	MaxResults *int32

	// A pagination token, indicating that you want to retrieve the next page of
	// results. If you received a value for NextToken in the response from a previous
	// ListJournalKinesisStreamsForLedger call, you should use that value as input
	// here.
	NextToken *string
}

type ListJournalKinesisStreamsForLedgerOutput struct {

	// * If NextToken is empty, the last page of results has been processed and there
	// are no more results to be retrieved.
	//
	//     * If NextToken is not empty, more
	// results are available. To retrieve the next page of results, use the value of
	// NextToken in a subsequent ListJournalKinesisStreamsForLedger call.
	NextToken *string

	// The array of QLDB journal stream descriptors that are associated with the given
	// ledger.
	Streams []*types.JournalKinesisStreamDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListJournalKinesisStreamsForLedgerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpListJournalKinesisStreamsForLedger{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpListJournalKinesisStreamsForLedger{})
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
	if err := addOpListJournalKinesisStreamsForLedgerValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opListJournalKinesisStreamsForLedger(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opListJournalKinesisStreamsForLedger(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "qldb",
		OperationName: "ListJournalKinesisStreamsForLedger",
	}
}
