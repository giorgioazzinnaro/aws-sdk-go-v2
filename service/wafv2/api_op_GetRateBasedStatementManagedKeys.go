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
// Retrieves the keys that are currently blocked by a rate-based rule. The maximum
// number of managed keys that can be blocked for a single rate-based rule is
// 10,000. If more than 10,000 addresses exceed the rate limit, those with the
// highest rates are blocked.
func (c *Client) GetRateBasedStatementManagedKeys(ctx context.Context, params *GetRateBasedStatementManagedKeysInput, optFns ...func(*Options)) (*GetRateBasedStatementManagedKeysOutput, error) {
	if params == nil {
		params = &GetRateBasedStatementManagedKeysInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRateBasedStatementManagedKeys", params, optFns, addOperationGetRateBasedStatementManagedKeysMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRateBasedStatementManagedKeysOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRateBasedStatementManagedKeysInput struct {

	// The name of the rate-based rule to get the keys for.
	//
	// This member is required.
	RuleName *string

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

	// The unique identifier for the Web ACL. This ID is returned in the responses to
	// create and list commands. You provide it to operations like update and delete.
	//
	// This member is required.
	WebACLId *string

	// The name of the Web ACL. You cannot change the name of a Web ACL after you
	// create it.
	//
	// This member is required.
	WebACLName *string
}

type GetRateBasedStatementManagedKeysOutput struct {

	// The keys that are of Internet Protocol version 4 (IPv4).
	ManagedKeysIPV4 *types.RateBasedStatementManagedKeysIPSet

	// The keys that are of Internet Protocol version 6 (IPv6).
	ManagedKeysIPV6 *types.RateBasedStatementManagedKeysIPSet

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetRateBasedStatementManagedKeysMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpGetRateBasedStatementManagedKeys{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpGetRateBasedStatementManagedKeys{})
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
	if err := addOpGetRateBasedStatementManagedKeysValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opGetRateBasedStatementManagedKeys(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opGetRateBasedStatementManagedKeys(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "wafv2",
		OperationName: "GetRateBasedStatementManagedKeys",
	}
}
