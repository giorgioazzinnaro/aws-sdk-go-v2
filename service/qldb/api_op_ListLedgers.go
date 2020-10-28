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

// Returns an array of ledger summaries that are associated with the current AWS
// account and Region. This action returns a maximum of 100 items and is paginated
// so that you can retrieve all the items by calling ListLedgers multiple times.
func (c *Client) ListLedgers(ctx context.Context, params *ListLedgersInput, optFns ...func(*Options)) (*ListLedgersOutput, error) {
	if params == nil {
		params = &ListLedgersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListLedgers", params, optFns, addOperationListLedgersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListLedgersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListLedgersInput struct {

	// The maximum number of results to return in a single ListLedgers request. (The
	// actual number of results returned might be fewer.)
	MaxResults *int32

	// A pagination token, indicating that you want to retrieve the next page of
	// results. If you received a value for NextToken in the response from a previous
	// ListLedgers call, then you should use that value as input here.
	NextToken *string
}

type ListLedgersOutput struct {

	// The array of ledger summaries that are associated with the current AWS account
	// and Region.
	Ledgers []*types.LedgerSummary

	// A pagination token, indicating whether there are more results available:
	//
	//     *
	// If NextToken is empty, then the last page of results has been processed and
	// there are no more results to be retrieved.
	//
	//     * If NextToken is not empty,
	// then there are more results available. To retrieve the next page of results, use
	// the value of NextToken in a subsequent ListLedgers call.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListLedgersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpListLedgers{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpListLedgers{})
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
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opListLedgers(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opListLedgers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "qldb",
		OperationName: "ListLedgers",
	}
}
