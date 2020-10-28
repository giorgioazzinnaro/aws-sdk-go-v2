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

// Returns a list of organization config rules. Only a master account and a
// delegated administrator account can call this API. When calling this API with a
// delegated administrator, you must ensure AWS Organizations
// ListDelegatedAdministrator permissions are added.  When you specify the limit
// and the next token, you receive a paginated response. Limit and next token are
// not applicable if you specify organization config rule names. It is only
// applicable, when you request all the organization config rules.
func (c *Client) DescribeOrganizationConfigRules(ctx context.Context, params *DescribeOrganizationConfigRulesInput, optFns ...func(*Options)) (*DescribeOrganizationConfigRulesOutput, error) {
	if params == nil {
		params = &DescribeOrganizationConfigRulesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeOrganizationConfigRules", params, optFns, addOperationDescribeOrganizationConfigRulesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeOrganizationConfigRulesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeOrganizationConfigRulesInput struct {

	// The maximum number of organization config rules returned on each page. If you do
	// no specify a number, AWS Config uses the default. The default is 100.
	Limit *int32

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	// The names of organization config rules for which you want details. If you do not
	// specify any names, AWS Config returns details for all your organization config
	// rules.
	OrganizationConfigRuleNames []*string
}

type DescribeOrganizationConfigRulesOutput struct {

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	// Returns a list of OrganizationConfigRule objects.
	OrganizationConfigRules []*types.OrganizationConfigRule

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeOrganizationConfigRulesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpDescribeOrganizationConfigRules{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpDescribeOrganizationConfigRules{})
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
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDescribeOrganizationConfigRules(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDescribeOrganizationConfigRules(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "DescribeOrganizationConfigRules",
	}
}
