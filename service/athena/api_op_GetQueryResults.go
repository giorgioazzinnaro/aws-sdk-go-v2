// Code generated by smithy-go-codegen DO NOT EDIT.

package athena

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Streams the results of a single query execution specified by QueryExecutionId
// from the Athena query results location in Amazon S3. For more information, see
// Query Results (https://docs.aws.amazon.com/athena/latest/ug/querying.html) in
// the Amazon Athena User Guide. This request does not execute the query but
// returns results. Use StartQueryExecution to run a query. To stream query results
// successfully, the IAM principal with permission to call GetQueryResults also
// must have permissions to the Amazon S3 GetObject action for the Athena query
// results location. IAM principals with permission to the Amazon S3 GetObject
// action for the query results location are able to retrieve query results from
// Amazon S3 even if permission to the GetQueryResults action is denied. To
// restrict user or role access, ensure that Amazon S3 permissions to the Athena
// query location are denied.
func (c *Client) GetQueryResults(ctx context.Context, params *GetQueryResultsInput, optFns ...func(*Options)) (*GetQueryResultsOutput, error) {
	if params == nil {
		params = &GetQueryResultsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetQueryResults", params, optFns, addOperationGetQueryResultsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetQueryResultsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetQueryResultsInput struct {

	// The unique ID of the query execution.
	//
	// This member is required.
	QueryExecutionId *string

	// The maximum number of results (rows) to return in this request.
	MaxResults *int32

	// A token generated by the Athena service that specifies where to continue
	// pagination if a previous request was truncated. To obtain the next set of pages,
	// pass in the NextToken from the response object of the previous page call.
	NextToken *string
}

type GetQueryResultsOutput struct {

	// A token generated by the Athena service that specifies where to continue
	// pagination if a previous request was truncated. To obtain the next set of pages,
	// pass in the NextToken from the response object of the previous page call.
	NextToken *string

	// The results of the query execution.
	ResultSet *types.ResultSet

	// The number of rows inserted with a CREATE TABLE AS SELECT statement.
	UpdateCount *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetQueryResultsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpGetQueryResults{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpGetQueryResults{})
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
	if err := addOpGetQueryResultsValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opGetQueryResults(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opGetQueryResults(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "athena",
		OperationName: "GetQueryResults",
	}
}
