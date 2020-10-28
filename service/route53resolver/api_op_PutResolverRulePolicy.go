// Code generated by smithy-go-codegen DO NOT EDIT.

package route53resolver

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Specifies an AWS account that you want to share rules with, the Resolver rules
// that you want to share, and the operations that you want the account to be able
// to perform on those rules.
func (c *Client) PutResolverRulePolicy(ctx context.Context, params *PutResolverRulePolicyInput, optFns ...func(*Options)) (*PutResolverRulePolicyOutput, error) {
	if params == nil {
		params = &PutResolverRulePolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutResolverRulePolicy", params, optFns, addOperationPutResolverRulePolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutResolverRulePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutResolverRulePolicyInput struct {

	// The Amazon Resource Name (ARN) of the account that you want to share rules with.
	//
	// This member is required.
	Arn *string

	// An AWS Identity and Access Management policy statement that lists the rules that
	// you want to share with another AWS account and the operations that you want the
	// account to be able to perform. You can specify the following operations in the
	// Actions section of the statement:
	//
	//     * route53resolver:GetResolverRule
	//
	//     *
	// route53resolver:AssociateResolverRule
	//
	//     *
	// route53resolver:DisassociateResolverRule
	//
	//     *
	// route53resolver:ListResolverRules
	//
	//     *
	// route53resolver:ListResolverRuleAssociations
	//
	// In the Resource section of the
	// statement, you specify the ARNs for the rules that you want to share with the
	// account that you specified in Arn.
	//
	// This member is required.
	ResolverRulePolicy *string
}

// The response to a PutResolverRulePolicy request.
type PutResolverRulePolicyOutput struct {

	// Whether the PutResolverRulePolicy request was successful.
	ReturnValue *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutResolverRulePolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpPutResolverRulePolicy{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpPutResolverRulePolicy{})
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
	if err := addOpPutResolverRulePolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opPutResolverRulePolicy(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opPutResolverRulePolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53resolver",
		OperationName: "PutResolverRulePolicy",
	}
}
