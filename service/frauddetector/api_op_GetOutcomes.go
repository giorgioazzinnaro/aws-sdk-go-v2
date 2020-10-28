// Code generated by smithy-go-codegen DO NOT EDIT.

package frauddetector

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets one or more outcomes. This is a paginated API. If you provide a null
// maxResults, this actions retrieves a maximum of 100 records per page. If you
// provide a maxResults, the value must be between 50 and 100. To get the next page
// results, provide the pagination token from the GetOutcomesResult as part of your
// request. A null pagination token fetches the records from the beginning.
func (c *Client) GetOutcomes(ctx context.Context, params *GetOutcomesInput, optFns ...func(*Options)) (*GetOutcomesOutput, error) {
	if params == nil {
		params = &GetOutcomesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetOutcomes", params, optFns, addOperationGetOutcomesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetOutcomesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetOutcomesInput struct {

	// The maximum number of objects to return for the request.
	MaxResults *int32

	// The name of the outcome or outcomes to get.
	Name *string

	// The next page token for the request.
	NextToken *string
}

type GetOutcomesOutput struct {

	// The next page token for subsequent requests.
	NextToken *string

	// The outcomes.
	Outcomes []*types.Outcome

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetOutcomesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpGetOutcomes{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpGetOutcomes{})
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
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opGetOutcomes(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opGetOutcomes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "frauddetector",
		OperationName: "GetOutcomes",
	}
}
