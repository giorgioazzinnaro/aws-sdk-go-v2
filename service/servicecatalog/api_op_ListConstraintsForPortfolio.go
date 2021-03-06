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

// Lists the constraints for the specified portfolio and product.
func (c *Client) ListConstraintsForPortfolio(ctx context.Context, params *ListConstraintsForPortfolioInput, optFns ...func(*Options)) (*ListConstraintsForPortfolioOutput, error) {
	if params == nil {
		params = &ListConstraintsForPortfolioInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListConstraintsForPortfolio", params, optFns, addOperationListConstraintsForPortfolioMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListConstraintsForPortfolioOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListConstraintsForPortfolioInput struct {

	// The portfolio identifier.
	//
	// This member is required.
	PortfolioId *string

	// The language code.
	//
	// * en - English (default)
	//
	// * jp - Japanese
	//
	// * zh - Chinese
	AcceptLanguage *string

	// The maximum number of items to return with this call.
	PageSize *int32

	// The page token for the next set of results. To retrieve the first set of
	// results, use null.
	PageToken *string

	// The product identifier.
	ProductId *string
}

type ListConstraintsForPortfolioOutput struct {

	// Information about the constraints.
	ConstraintDetails []*types.ConstraintDetail

	// The page token to use to retrieve the next set of results. If there are no
	// additional results, this value is null.
	NextPageToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListConstraintsForPortfolioMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListConstraintsForPortfolio{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListConstraintsForPortfolio{}, middleware.After)
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
	addOpListConstraintsForPortfolioValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListConstraintsForPortfolio(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListConstraintsForPortfolio(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "ListConstraintsForPortfolio",
	}
}
