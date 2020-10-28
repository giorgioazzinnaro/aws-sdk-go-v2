// Code generated by smithy-go-codegen DO NOT EDIT.

package machinelearning

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/machinelearning/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of BatchPrediction operations that match the search criteria in
// the request.
func (c *Client) DescribeBatchPredictions(ctx context.Context, params *DescribeBatchPredictionsInput, optFns ...func(*Options)) (*DescribeBatchPredictionsOutput, error) {
	if params == nil {
		params = &DescribeBatchPredictionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeBatchPredictions", params, optFns, addOperationDescribeBatchPredictionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeBatchPredictionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeBatchPredictionsInput struct {

	// The equal to operator. The BatchPrediction results will have FilterVariable
	// values that exactly match the value specified with EQ.
	EQ *string

	// Use one of the following variables to filter a list of BatchPrediction:
	//
	//     *
	// CreatedAt - Sets the search criteria to the BatchPrediction creation date.
	//
	//
	// * Status - Sets the search criteria to the BatchPrediction status.
	//
	//     * Name -
	// Sets the search criteria to the contents of the BatchPredictionName.
	//
	//     *
	// IAMUser - Sets the search criteria to the user account that invoked the
	// BatchPrediction creation.
	//
	//     * MLModelId - Sets the search criteria to the
	// MLModel used in the BatchPrediction.
	//
	//     * DataSourceId - Sets the search
	// criteria to the DataSource used in the BatchPrediction.
	//
	//     * DataURI - Sets
	// the search criteria to the data file(s) used in the BatchPrediction. The URL can
	// identify either a file or an Amazon Simple Storage Solution (Amazon S3) bucket
	// or directory.
	FilterVariable types.BatchPredictionFilterVariable

	// The greater than or equal to operator. The BatchPrediction results will have
	// FilterVariable values that are greater than or equal to the value specified with
	// GE.
	GE *string

	// The greater than operator. The BatchPrediction results will have FilterVariable
	// values that are greater than the value specified with GT.
	GT *string

	// The less than or equal to operator. The BatchPrediction results will have
	// FilterVariable values that are less than or equal to the value specified with
	// LE.
	LE *string

	// The less than operator. The BatchPrediction results will have FilterVariable
	// values that are less than the value specified with LT.
	LT *string

	// The number of pages of information to include in the result. The range of
	// acceptable values is 1 through 100. The default value is 100.
	Limit *int32

	// The not equal to operator. The BatchPrediction results will have FilterVariable
	// values not equal to the value specified with NE.
	NE *string

	// An ID of the page in the paginated results.
	NextToken *string

	// A string that is found at the beginning of a variable, such as Name or Id. For
	// example, a Batch Prediction operation could have the
	// Name2014-09-09-HolidayGiftMailer. To search for this BatchPrediction, select
	// Name for the FilterVariable and any of the following strings for the Prefix:
	//
	//
	// * 2014-09
	//
	//     * 2014-09-09
	//
	//     * 2014-09-09-Holiday
	Prefix *string

	// A two-value parameter that determines the sequence of the resulting list of
	// MLModels.
	//
	//     * asc - Arranges the list in ascending order (A-Z, 0-9).
	//
	//     *
	// dsc - Arranges the list in descending order (Z-A, 9-0).
	//
	// Results are sorted by
	// FilterVariable.
	SortOrder types.SortOrder
}

// Represents the output of a DescribeBatchPredictions operation. The content is
// essentially a list of BatchPredictions.
type DescribeBatchPredictionsOutput struct {

	// The ID of the next page in the paginated results that indicates at least one
	// more page follows.
	NextToken *string

	// A list of BatchPrediction objects that meet the search criteria.
	Results []*types.BatchPrediction

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeBatchPredictionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpDescribeBatchPredictions{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpDescribeBatchPredictions{})
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
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDescribeBatchPredictions(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDescribeBatchPredictions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "machinelearning",
		OperationName: "DescribeBatchPredictions",
	}
}
