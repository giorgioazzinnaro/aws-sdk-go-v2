// Code generated by smithy-go-codegen DO NOT EDIT.

package xray

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/xray/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves a list of traces specified by ID. Each trace is a collection of
// segment documents that originates from a single request. Use GetTraceSummaries
// to get a list of trace IDs.
func (c *Client) BatchGetTraces(ctx context.Context, params *BatchGetTracesInput, optFns ...func(*Options)) (*BatchGetTracesOutput, error) {
	if params == nil {
		params = &BatchGetTracesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchGetTraces", params, optFns, addOperationBatchGetTracesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchGetTracesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchGetTracesInput struct {

	// Specify the trace IDs of requests for which to retrieve segments.
	//
	// This member is required.
	TraceIds []*string

	// Pagination token.
	NextToken *string
}

type BatchGetTracesOutput struct {

	// Pagination token.
	NextToken *string

	// Full traces for the specified requests.
	Traces []*types.Trace

	// Trace IDs of requests that haven't been processed.
	UnprocessedTraceIds []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationBatchGetTracesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpBatchGetTraces{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpBatchGetTraces{})
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
	if err := addOpBatchGetTracesValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opBatchGetTraces(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opBatchGetTraces(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "xray",
		OperationName: "BatchGetTraces",
	}
}
