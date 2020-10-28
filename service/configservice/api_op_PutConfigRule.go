// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds or updates an AWS Config rule for evaluating whether your AWS resources
// comply with your desired configurations. You can use this action for custom AWS
// Config rules and AWS managed Config rules. A custom AWS Config rule is a rule
// that you develop and maintain. An AWS managed Config rule is a customizable,
// predefined rule that AWS Config provides. If you are adding a new custom AWS
// Config rule, you must first create the AWS Lambda function that the rule invokes
// to evaluate your resources. When you use the PutConfigRule action to add the
// rule to AWS Config, you must specify the Amazon Resource Name (ARN) that AWS
// Lambda assigns to the function. Specify the ARN for the SourceIdentifier key.
// This key is part of the Source object, which is part of the ConfigRule object.
// If you are adding an AWS managed Config rule, specify the rule's identifier for
// the SourceIdentifier key. To reference AWS managed Config rule identifiers, see
// About AWS Managed Config Rules
// (https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config_use-managed-rules.html).
// For any new rule that you add, specify the ConfigRuleName in the ConfigRule
// object. Do not specify the ConfigRuleArn or the ConfigRuleId. These values are
// generated by AWS Config for new rules. If you are updating a rule that you added
// previously, you can specify the rule by ConfigRuleName, ConfigRuleId, or
// ConfigRuleArn in the ConfigRule data type that you use in this request. The
// maximum number of rules that AWS Config supports is 150. For information about
// requesting a rule limit increase, see AWS Config Limits
// (http://docs.aws.amazon.com/general/latest/gr/aws_service_limits.html#limits_config)
// in the AWS General Reference Guide. For more information about developing and
// using AWS Config rules, see Evaluating AWS Resource Configurations with AWS
// Config
// (https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config.html)
// in the AWS Config Developer Guide.
func (c *Client) PutConfigRule(ctx context.Context, params *PutConfigRuleInput, optFns ...func(*Options)) (*PutConfigRuleOutput, error) {
	if params == nil {
		params = &PutConfigRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutConfigRule", params, optFns, addOperationPutConfigRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutConfigRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutConfigRuleInput struct {

	// The rule that you want to add to your account.
	//
	// This member is required.
	ConfigRule *types.ConfigRule

	// An array of tag object.
	Tags []*types.Tag
}

type PutConfigRuleOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutConfigRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpPutConfigRule{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpPutConfigRule{})
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
	if err := addOpPutConfigRuleValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opPutConfigRule(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opPutConfigRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "PutConfigRule",
	}
}
