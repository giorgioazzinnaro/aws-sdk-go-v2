// Code generated by smithy-go-codegen DO NOT EDIT.

package support

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/support/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the results for the AWS Trusted Advisor check summaries for the check
// IDs that you specified. You can get the check IDs by calling the
// DescribeTrustedAdvisorChecks operation. The response contains an array of
// TrustedAdvisorCheckSummary objects.
//
// * You must have a Business or Enterprise
// support plan to use the AWS Support API.
//
// * If you call the AWS Support API from
// an account that does not have a Business or Enterprise support plan, the
// SubscriptionRequiredException error message appears. For information about
// changing your support plan, see AWS Support
// (http://aws.amazon.com/premiumsupport/).
func (c *Client) DescribeTrustedAdvisorCheckSummaries(ctx context.Context, params *DescribeTrustedAdvisorCheckSummariesInput, optFns ...func(*Options)) (*DescribeTrustedAdvisorCheckSummariesOutput, error) {
	if params == nil {
		params = &DescribeTrustedAdvisorCheckSummariesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTrustedAdvisorCheckSummaries", params, optFns, addOperationDescribeTrustedAdvisorCheckSummariesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTrustedAdvisorCheckSummariesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTrustedAdvisorCheckSummariesInput struct {

	// The IDs of the Trusted Advisor checks.
	//
	// This member is required.
	CheckIds []*string
}

// The summaries of the Trusted Advisor checks returned by the
// DescribeTrustedAdvisorCheckSummaries operation.
type DescribeTrustedAdvisorCheckSummariesOutput struct {

	// The summary information for the requested Trusted Advisor checks.
	//
	// This member is required.
	Summaries []*types.TrustedAdvisorCheckSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeTrustedAdvisorCheckSummariesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeTrustedAdvisorCheckSummaries{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeTrustedAdvisorCheckSummaries{}, middleware.After)
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
	addOpDescribeTrustedAdvisorCheckSummariesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTrustedAdvisorCheckSummaries(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeTrustedAdvisorCheckSummaries(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "support",
		OperationName: "DescribeTrustedAdvisorCheckSummaries",
	}
}
