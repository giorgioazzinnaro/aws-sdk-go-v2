// Code generated by smithy-go-codegen DO NOT EDIT.

package budgets

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/budgets/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes a budget action history detail.
func (c *Client) DescribeBudgetActionHistories(ctx context.Context, params *DescribeBudgetActionHistoriesInput, optFns ...func(*Options)) (*DescribeBudgetActionHistoriesOutput, error) {
	if params == nil {
		params = &DescribeBudgetActionHistoriesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeBudgetActionHistories", params, optFns, addOperationDescribeBudgetActionHistoriesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeBudgetActionHistoriesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeBudgetActionHistoriesInput struct {

	// The account ID of the user. It should be a 12-digit number.
	//
	// This member is required.
	AccountId *string

	// A system-generated universally unique identifier (UUID) for the action.
	//
	// This member is required.
	ActionId *string

	// A string that represents the budget name. The ":" and "\" characters aren't
	// allowed.
	//
	// This member is required.
	BudgetName *string

	// An integer that represents how many entries a paginated response contains. The
	// maximum is 100.
	MaxResults *int32

	// A generic string.
	NextToken *string

	// The period of time that is covered by a budget. The period has a start date and
	// an end date. The start date must come before the end date. There are no
	// restrictions on the end date.
	TimePeriod *types.TimePeriod
}

type DescribeBudgetActionHistoriesOutput struct {

	// The historical record of the budget action resource.
	//
	// This member is required.
	ActionHistories []*types.ActionHistory

	// A generic string.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeBudgetActionHistoriesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpDescribeBudgetActionHistories{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpDescribeBudgetActionHistories{})
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
	if err := addOpDescribeBudgetActionHistoriesValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDescribeBudgetActionHistories(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDescribeBudgetActionHistories(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "budgets",
		OperationName: "DescribeBudgetActionHistories",
	}
}
