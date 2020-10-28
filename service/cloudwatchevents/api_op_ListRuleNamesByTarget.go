// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchevents

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the rules for the specified target. You can see which of the rules in
// Amazon EventBridge can invoke a specific target in your account.
func (c *Client) ListRuleNamesByTarget(ctx context.Context, params *ListRuleNamesByTargetInput, optFns ...func(*Options)) (*ListRuleNamesByTargetOutput, error) {
	if params == nil {
		params = &ListRuleNamesByTargetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRuleNamesByTarget", params, optFns, addOperationListRuleNamesByTargetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRuleNamesByTargetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRuleNamesByTargetInput struct {

	// The Amazon Resource Name (ARN) of the target resource.
	//
	// This member is required.
	TargetArn *string

	// Limits the results to show only the rules associated with the specified event
	// bus.
	EventBusName *string

	// The maximum number of results to return.
	Limit *int32

	// The token returned by a previous call to retrieve the next set of results.
	NextToken *string
}

type ListRuleNamesByTargetOutput struct {

	// Indicates whether there are additional results to retrieve. If there are no more
	// results, the value is null.
	NextToken *string

	// The names of the rules that can invoke the given target.
	RuleNames []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListRuleNamesByTargetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpListRuleNamesByTarget{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpListRuleNamesByTarget{})
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
	if err := addOpListRuleNamesByTargetValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opListRuleNamesByTarget(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opListRuleNamesByTarget(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "events",
		OperationName: "ListRuleNamesByTarget",
	}
}
