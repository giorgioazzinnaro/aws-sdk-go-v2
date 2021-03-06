// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Searches for analyses that belong to the user specified in the filter.
func (c *Client) SearchAnalyses(ctx context.Context, params *SearchAnalysesInput, optFns ...func(*Options)) (*SearchAnalysesOutput, error) {
	if params == nil {
		params = &SearchAnalysesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchAnalyses", params, optFns, addOperationSearchAnalysesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchAnalysesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchAnalysesInput struct {

	// The ID of the AWS account that contains the analyses that you're searching for.
	//
	// This member is required.
	AwsAccountId *string

	// The structure for the search filters that you want to apply to your search.
	//
	// This member is required.
	Filters []*types.AnalysisSearchFilter

	// The maximum number of results to return.
	MaxResults *int32

	// A pagination token that can be used in a subsequent request.
	NextToken *string
}

type SearchAnalysesOutput struct {

	// Metadata describing the analyses that you searched for.
	AnalysisSummaryList []*types.AnalysisSummary

	// A pagination token that can be used in a subsequent request.
	NextToken *string

	// The AWS request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationSearchAnalysesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSearchAnalyses{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSearchAnalyses{}, middleware.After)
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
	addOpSearchAnalysesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSearchAnalyses(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opSearchAnalyses(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "SearchAnalyses",
	}
}
