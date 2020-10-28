// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Retrieves a list of email addresses that are on the suppression list for your
// account.
func (c *Client) ListSuppressedDestinations(ctx context.Context, params *ListSuppressedDestinationsInput, optFns ...func(*Options)) (*ListSuppressedDestinationsOutput, error) {
	if params == nil {
		params = &ListSuppressedDestinationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSuppressedDestinations", params, optFns, addOperationListSuppressedDestinationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSuppressedDestinationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to obtain a list of email destinations that are on the suppression
// list for your account.
type ListSuppressedDestinationsInput struct {

	// Used to filter the list of suppressed email destinations so that it only
	// includes addresses that were added to the list before a specific date. The date
	// that you specify should be in Unix time format.
	EndDate *time.Time

	// A token returned from a previous call to ListSuppressedDestinations to indicate
	// the position in the list of suppressed email addresses.
	NextToken *string

	// The number of results to show in a single call to ListSuppressedDestinations. If
	// the number of results is larger than the number you specified in this parameter,
	// then the response includes a NextToken element, which you can use to obtain
	// additional results.
	PageSize *int32

	// The factors that caused the email address to be added to .
	Reasons []types.SuppressionListReason

	// Used to filter the list of suppressed email destinations so that it only
	// includes addresses that were added to the list after a specific date. The date
	// that you specify should be in Unix time format.
	StartDate *time.Time
}

// A list of suppressed email addresses.
type ListSuppressedDestinationsOutput struct {

	// A token that indicates that there are additional email addresses on the
	// suppression list for your account. To view additional suppressed addresses,
	// issue another request to ListSuppressedDestinations, and pass this token in the
	// NextToken parameter.
	NextToken *string

	// A list of summaries, each containing a summary for a suppressed email
	// destination.
	SuppressedDestinationSummaries []*types.SuppressedDestinationSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListSuppressedDestinationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpListSuppressedDestinations{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpListSuppressedDestinations{})
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
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opListSuppressedDestinations(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opListSuppressedDestinations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "ListSuppressedDestinations",
	}
}
