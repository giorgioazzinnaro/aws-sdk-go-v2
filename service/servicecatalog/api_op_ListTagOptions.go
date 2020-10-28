// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the specified TagOptions or all TagOptions.
func (c *Client) ListTagOptions(ctx context.Context, params *ListTagOptionsInput, optFns ...func(*Options)) (*ListTagOptionsOutput, error) {
	if params == nil {
		params = &ListTagOptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTagOptions", params, optFns, addOperationListTagOptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTagOptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTagOptionsInput struct {

	// The search filters. If no search filters are specified, the output includes all
	// TagOptions.
	Filters *types.ListTagOptionsFilters

	// The maximum number of items to return with this call.
	PageSize *int32

	// The page token for the next set of results. To retrieve the first set of
	// results, use null.
	PageToken *string
}

type ListTagOptionsOutput struct {

	// The page token for the next set of results. To retrieve the first set of
	// results, use null.
	PageToken *string

	// Information about the TagOptions.
	TagOptionDetails []*types.TagOptionDetail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListTagOptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpListTagOptions{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpListTagOptions{})
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
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opListTagOptions(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opListTagOptions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "ListTagOptions",
	}
}
