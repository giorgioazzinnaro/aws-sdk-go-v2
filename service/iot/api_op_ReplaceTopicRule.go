// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Replaces the rule. You must specify all parameters for the new rule. Creating
// rules is an administrator-level action. Any user who has permission to create
// rules will be able to access data processed by the rule.
func (c *Client) ReplaceTopicRule(ctx context.Context, params *ReplaceTopicRuleInput, optFns ...func(*Options)) (*ReplaceTopicRuleOutput, error) {
	if params == nil {
		params = &ReplaceTopicRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ReplaceTopicRule", params, optFns, addOperationReplaceTopicRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ReplaceTopicRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the ReplaceTopicRule operation.
type ReplaceTopicRuleInput struct {

	// The name of the rule.
	//
	// This member is required.
	RuleName *string

	// The rule payload.
	//
	// This member is required.
	TopicRulePayload *types.TopicRulePayload
}

type ReplaceTopicRuleOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationReplaceTopicRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpReplaceTopicRule{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpReplaceTopicRule{})
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
	if err := addOpReplaceTopicRuleValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opReplaceTopicRule(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opReplaceTopicRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ReplaceTopicRule",
	}
}
