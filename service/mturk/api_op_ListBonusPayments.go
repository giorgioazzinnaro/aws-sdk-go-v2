// Code generated by smithy-go-codegen DO NOT EDIT.

package mturk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mturk/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The ListBonusPayments operation retrieves the amounts of bonuses you have paid
// to Workers for a given HIT or assignment.
func (c *Client) ListBonusPayments(ctx context.Context, params *ListBonusPaymentsInput, optFns ...func(*Options)) (*ListBonusPaymentsOutput, error) {
	if params == nil {
		params = &ListBonusPaymentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListBonusPayments", params, optFns, addOperationListBonusPaymentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListBonusPaymentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListBonusPaymentsInput struct {

	// The ID of the assignment associated with the bonus payments to retrieve. If
	// specified, only bonus payments for the given assignment are returned. Either the
	// HITId parameter or the AssignmentId parameter must be specified
	AssignmentId *string

	// The ID of the HIT associated with the bonus payments to retrieve. If not
	// specified, all bonus payments for all assignments for the given HIT are
	// returned. Either the HITId parameter or the AssignmentId parameter must be
	// specified
	HITId *string

	MaxResults *int32

	// Pagination token
	NextToken *string
}

type ListBonusPaymentsOutput struct {

	// A successful request to the ListBonusPayments operation returns a list of
	// BonusPayment objects.
	BonusPayments []*types.BonusPayment

	// If the previous response was incomplete (because there is more data to
	// retrieve), Amazon Mechanical Turk returns a pagination token in the response.
	// You can use this pagination token to retrieve the next set of results.
	NextToken *string

	// The number of bonus payments on this page in the filtered results list,
	// equivalent to the number of bonus payments being returned by this call.
	NumResults *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListBonusPaymentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpListBonusPayments{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpListBonusPayments{})
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
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opListBonusPayments(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opListBonusPayments(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mturk-requester",
		OperationName: "ListBonusPayments",
	}
}
