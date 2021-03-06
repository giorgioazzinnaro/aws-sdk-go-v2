// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Reorders the receipt rules within a receipt rule set. All of the rules in the
// rule set must be represented in this request. That is, this API will return an
// error if the reorder request doesn't explicitly position all of the rules. For
// information about managing receipt rule sets, see the Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-managing-receipt-rule-sets.html).
// You can execute this operation no more than once per second.
func (c *Client) ReorderReceiptRuleSet(ctx context.Context, params *ReorderReceiptRuleSetInput, optFns ...func(*Options)) (*ReorderReceiptRuleSetOutput, error) {
	if params == nil {
		params = &ReorderReceiptRuleSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ReorderReceiptRuleSet", params, optFns, addOperationReorderReceiptRuleSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ReorderReceiptRuleSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to reorder the receipt rules within a receipt rule set. You
// use receipt rule sets to receive email with Amazon SES. For more information,
// see the Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-concepts.html).
type ReorderReceiptRuleSetInput struct {

	// A list of the specified receipt rule set's receipt rules in the order that you
	// want to put them.
	//
	// This member is required.
	RuleNames []*string

	// The name of the receipt rule set to reorder.
	//
	// This member is required.
	RuleSetName *string
}

// An empty element returned on a successful request.
type ReorderReceiptRuleSetOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationReorderReceiptRuleSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpReorderReceiptRuleSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpReorderReceiptRuleSet{}, middleware.After)
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
	addOpReorderReceiptRuleSetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opReorderReceiptRuleSet(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opReorderReceiptRuleSet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "ReorderReceiptRuleSet",
	}
}
