// Code generated by smithy-go-codegen DO NOT EDIT.

package wafv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This is the latest version of AWS WAF, named AWS WAFV2, released in November,
// 2019. For information, including how to migrate your AWS WAF resources from the
// prior release, see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html).
// Returns the web ACL capacity unit (WCU) requirements for a specified scope and
// set of rules. You can use this to check the capacity requirements for the rules
// you want to use in a RuleGroup or WebACL. AWS WAF uses WCUs to calculate and
// control the operating resources that are used to run your rules, rule groups,
// and web ACLs. AWS WAF calculates capacity differently for each rule type, to
// reflect the relative cost of each rule. Simple rules that cost little to run use
// fewer WCUs than more complex rules that use more processing power. Rule group
// capacity is fixed at creation, which helps users plan their web ACL WCU usage
// when they use a rule group. The WCU limit for web ACLs is 1,500.
func (c *Client) CheckCapacity(ctx context.Context, params *CheckCapacityInput, optFns ...func(*Options)) (*CheckCapacityOutput, error) {
	if params == nil {
		params = &CheckCapacityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CheckCapacity", params, optFns, addOperationCheckCapacityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CheckCapacityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CheckCapacityInput struct {

	// An array of Rule that you're configuring to use in a rule group or web ACL.
	//
	// This member is required.
	Rules []*types.Rule

	// Specifies whether this is for an AWS CloudFront distribution or for a regional
	// application. A regional application can be an Application Load Balancer (ALB),
	// an API Gateway REST API, or an AppSync GraphQL API. To work with CloudFront, you
	// must also specify the Region US East (N. Virginia) as follows:
	//
	//     * CLI -
	// Specify the Region when you use the CloudFront scope: --scope=CLOUDFRONT
	// --region=us-east-1.
	//
	//     * API and SDKs - For all calls, use the Region endpoint
	// us-east-1.
	//
	// This member is required.
	Scope types.Scope
}

type CheckCapacityOutput struct {

	// The capacity required by the rules and scope.
	Capacity *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCheckCapacityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpCheckCapacity{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpCheckCapacity{})
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
	if err := addOpCheckCapacityValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opCheckCapacity(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opCheckCapacity(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "wafv2",
		OperationName: "CheckCapacity",
	}
}
