// Code generated by smithy-go-codegen DO NOT EDIT.

package robomaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/robomaker/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists world generator jobs.
func (c *Client) ListWorldGenerationJobs(ctx context.Context, params *ListWorldGenerationJobsInput, optFns ...func(*Options)) (*ListWorldGenerationJobsOutput, error) {
	if params == nil {
		params = &ListWorldGenerationJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListWorldGenerationJobs", params, optFns, addOperationListWorldGenerationJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListWorldGenerationJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListWorldGenerationJobsInput struct {

	// Optional filters to limit results. You can use status and templateId.
	Filters []*types.Filter

	// When this parameter is used, ListWorldGeneratorJobs only returns maxResults
	// results in a single page along with a nextToken response element. The remaining
	// results of the initial request can be seen by sending another
	// ListWorldGeneratorJobs request with the returned nextToken value. This value can
	// be between 1 and 100. If this parameter is not used, then ListWorldGeneratorJobs
	// returns up to 100 results and a nextToken value if applicable.
	MaxResults *int32

	// If the previous paginated request did not return all of the remaining results,
	// the response object's nextToken parameter value is set to a token. To retrieve
	// the next set of results, call ListWorldGenerationJobsRequest again and assign
	// that token to the request object's nextToken parameter. If there are no
	// remaining results, the previous response object's NextToken parameter is set to
	// null.
	NextToken *string
}

type ListWorldGenerationJobsOutput struct {

	// Summary information for world generator jobs.
	//
	// This member is required.
	WorldGenerationJobSummaries []*types.WorldGenerationJobSummary

	// If the previous paginated request did not return all of the remaining results,
	// the response object's nextToken parameter value is set to a token. To retrieve
	// the next set of results, call ListWorldGeneratorJobsRequest again and assign
	// that token to the request object's nextToken parameter. If there are no
	// remaining results, the previous response object's NextToken parameter is set to
	// null.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListWorldGenerationJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListWorldGenerationJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListWorldGenerationJobs{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListWorldGenerationJobs(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListWorldGenerationJobs(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "robomaker",
		OperationName: "ListWorldGenerationJobs",
	}
}
