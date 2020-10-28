// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Provides the sending limits for the Amazon SES account. You can execute this
// operation no more than once per second.
func (c *Client) GetSendQuota(ctx context.Context, params *GetSendQuotaInput, optFns ...func(*Options)) (*GetSendQuotaOutput, error) {
	if params == nil {
		params = &GetSendQuotaInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSendQuota", params, optFns, addOperationGetSendQuotaMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSendQuotaOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSendQuotaInput struct {
}

// Represents your Amazon SES daily sending quota, maximum send rate, and the
// number of emails you have sent in the last 24 hours.
type GetSendQuotaOutput struct {

	// The maximum number of emails the user is allowed to send in a 24-hour interval.
	// A value of -1 signifies an unlimited quota.
	Max24HourSend *float64

	// The maximum number of emails that Amazon SES can accept from the user's account
	// per second. The rate at which Amazon SES accepts the user's messages might be
	// less than the maximum send rate.
	MaxSendRate *float64

	// The number of emails sent during the previous 24 hours.
	SentLast24Hours *float64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetSendQuotaMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsquery_serializeOpGetSendQuota{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsquery_deserializeOpGetSendQuota{})
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
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opGetSendQuota(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opGetSendQuota(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "GetSendQuota",
	}
}
