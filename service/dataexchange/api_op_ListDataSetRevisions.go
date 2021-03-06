// Code generated by smithy-go-codegen DO NOT EDIT.

package dataexchange

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/dataexchange/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This operation lists a data set's revisions sorted by CreatedAt in descending
// order.
func (c *Client) ListDataSetRevisions(ctx context.Context, params *ListDataSetRevisionsInput, optFns ...func(*Options)) (*ListDataSetRevisionsOutput, error) {
	if params == nil {
		params = &ListDataSetRevisionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDataSetRevisions", params, optFns, addOperationListDataSetRevisionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDataSetRevisionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDataSetRevisionsInput struct {

	// The unique identifier for a data set.
	//
	// This member is required.
	DataSetId *string

	// The maximum number of results returned by a single call.
	MaxResults *int32

	// The token value retrieved from a previous call to access the next page of
	// results.
	NextToken *string
}

type ListDataSetRevisionsOutput struct {

	// The token value retrieved from a previous call to access the next page of
	// results.
	NextToken *string

	// The asset objects listed by the request.
	Revisions []*types.RevisionEntry

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListDataSetRevisionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDataSetRevisions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDataSetRevisions{}, middleware.After)
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
	addOpListDataSetRevisionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListDataSetRevisions(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListDataSetRevisions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dataexchange",
		OperationName: "ListDataSetRevisions",
	}
}
