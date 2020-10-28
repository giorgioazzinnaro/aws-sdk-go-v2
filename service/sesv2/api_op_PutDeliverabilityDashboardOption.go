// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Enable or disable the Deliverability dashboard. When you enable the
// Deliverability dashboard, you gain access to reputation, deliverability, and
// other metrics for the domains that you use to send email. You also gain the
// ability to perform predictive inbox placement tests. When you use the
// Deliverability dashboard, you pay a monthly subscription charge, in addition to
// any other fees that you accrue by using Amazon SES and other AWS services. For
// more information about the features and cost of a Deliverability dashboard
// subscription, see Amazon SES Pricing (http://aws.amazon.com/ses/pricing/).
func (c *Client) PutDeliverabilityDashboardOption(ctx context.Context, params *PutDeliverabilityDashboardOptionInput, optFns ...func(*Options)) (*PutDeliverabilityDashboardOptionOutput, error) {
	if params == nil {
		params = &PutDeliverabilityDashboardOptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutDeliverabilityDashboardOption", params, optFns, addOperationPutDeliverabilityDashboardOptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutDeliverabilityDashboardOptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Enable or disable the Deliverability dashboard. When you enable the
// Deliverability dashboard, you gain access to reputation, deliverability, and
// other metrics for the domains that you use to send email using Amazon SES API
// v2. You also gain the ability to perform predictive inbox placement tests. When
// you use the Deliverability dashboard, you pay a monthly subscription charge, in
// addition to any other fees that you accrue by using Amazon SES and other AWS
// services. For more information about the features and cost of a Deliverability
// dashboard subscription, see Amazon Pinpoint Pricing
// (http://aws.amazon.com/pinpoint/pricing/).
type PutDeliverabilityDashboardOptionInput struct {

	// Specifies whether to enable the Deliverability dashboard. To enable the
	// dashboard, set this value to true.
	//
	// This member is required.
	DashboardEnabled *bool

	// An array of objects, one for each verified domain that you use to send email and
	// enabled the Deliverability dashboard for.
	SubscribedDomains []*types.DomainDeliverabilityTrackingOption
}

// A response that indicates whether the Deliverability dashboard is enabled.
type PutDeliverabilityDashboardOptionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutDeliverabilityDashboardOptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpPutDeliverabilityDashboardOption{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpPutDeliverabilityDashboardOption{})
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
	if err := addOpPutDeliverabilityDashboardOptionValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opPutDeliverabilityDashboardOption(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opPutDeliverabilityDashboardOption(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "PutDeliverabilityDashboardOption",
	}
}
