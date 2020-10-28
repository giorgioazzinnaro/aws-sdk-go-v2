// Code generated by smithy-go-codegen DO NOT EDIT.

package servicequotas

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicequotas/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Requests a list of the changes to specific service quotas. This command provides
// additional granularity over the ListRequestedServiceQuotaChangeHistory command.
// Once a quota change request has reached CASE_CLOSED, APPROVED, or DENIED, the
// history has been kept for 90 days.
func (c *Client) ListRequestedServiceQuotaChangeHistoryByQuota(ctx context.Context, params *ListRequestedServiceQuotaChangeHistoryByQuotaInput, optFns ...func(*Options)) (*ListRequestedServiceQuotaChangeHistoryByQuotaOutput, error) {
	if params == nil {
		params = &ListRequestedServiceQuotaChangeHistoryByQuotaInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRequestedServiceQuotaChangeHistoryByQuota", params, optFns, addOperationListRequestedServiceQuotaChangeHistoryByQuotaMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRequestedServiceQuotaChangeHistoryByQuotaOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRequestedServiceQuotaChangeHistoryByQuotaInput struct {

	// Specifies the service quota that you want to use
	//
	// This member is required.
	QuotaCode *string

	// Specifies the service that you want to use.
	//
	// This member is required.
	ServiceCode *string

	// (Optional) Limits the number of results that you want to include in the
	// response. If you don't include this parameter, the response defaults to a value
	// that's specific to the operation. If additional items exist beyond the specified
	// maximum, the NextToken element is present and has a value (isn't null). Include
	// that value as the NextToken request parameter in the call to the operation to
	// get the next part of the results. You should check NextToken after every
	// operation to ensure that you receive all of the results.
	MaxResults *int32

	// (Optional) Use this parameter in a request if you receive a NextToken response
	// in a previous request that indicates that there's more output available. In a
	// subsequent call, set it to the value of the previous call's NextToken response
	// to indicate where the output should continue from.
	NextToken *string

	// Specifies the status value of the quota increase request.
	Status types.RequestStatus
}

type ListRequestedServiceQuotaChangeHistoryByQuotaOutput struct {

	// If present in the response, this value indicates there's more output available
	// that what's included in the current response. This can occur even when the
	// response includes no values at all, such as when you ask for a filtered view of
	// a very long list. Use this value in the NextToken request parameter in a
	// subsequent call to the operation to continue processing and get the next part of
	// the output. You should repeat this until the NextToken response element comes
	// back empty (as null).
	NextToken *string

	// Returns a list of service quota requests.
	RequestedQuotas []*types.RequestedServiceQuotaChange

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListRequestedServiceQuotaChangeHistoryByQuotaMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpListRequestedServiceQuotaChangeHistoryByQuota{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpListRequestedServiceQuotaChangeHistoryByQuota{})
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
	if err := addOpListRequestedServiceQuotaChangeHistoryByQuotaValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opListRequestedServiceQuotaChangeHistoryByQuota(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opListRequestedServiceQuotaChangeHistoryByQuota(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicequotas",
		OperationName: "ListRequestedServiceQuotaChangeHistoryByQuota",
	}
}
