// Code generated by smithy-go-codegen DO NOT EDIT.

package route53resolver

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes the association between a specified Resolver rule and a specified VPC.
// If you disassociate a Resolver rule from a VPC, Resolver stops forwarding DNS
// queries for the domain name that you specified in the Resolver rule.
func (c *Client) DisassociateResolverRule(ctx context.Context, params *DisassociateResolverRuleInput, optFns ...func(*Options)) (*DisassociateResolverRuleOutput, error) {
	if params == nil {
		params = &DisassociateResolverRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateResolverRule", params, optFns, addOperationDisassociateResolverRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateResolverRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateResolverRuleInput struct {

	// The ID of the Resolver rule that you want to disassociate from the specified
	// VPC.
	//
	// This member is required.
	ResolverRuleId *string

	// The ID of the VPC that you want to disassociate the Resolver rule from.
	//
	// This member is required.
	VPCId *string
}

type DisassociateResolverRuleOutput struct {

	// Information about the DisassociateResolverRule request, including the status of
	// the request.
	ResolverRuleAssociation *types.ResolverRuleAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDisassociateResolverRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDisassociateResolverRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisassociateResolverRule{}, middleware.After)
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
	addOpDisassociateResolverRuleValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateResolverRule(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDisassociateResolverRule(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53resolver",
		OperationName: "DisassociateResolverRule",
	}
}
