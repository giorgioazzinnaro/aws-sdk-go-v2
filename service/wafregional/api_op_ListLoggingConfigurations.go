// Code generated by smithy-go-codegen DO NOT EDIT.

package wafregional

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/wafregional/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This is AWS WAF Classic documentation. For more information, see AWS WAF Classic
// (https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html)
// in the developer guide. For the latest version of AWS WAF, use the AWS WAFV2 API
// and see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html). With
// the latest version, AWS WAF has a single set of endpoints for regional and
// global use. Returns an array of LoggingConfiguration objects.
func (c *Client) ListLoggingConfigurations(ctx context.Context, params *ListLoggingConfigurationsInput, optFns ...func(*Options)) (*ListLoggingConfigurationsOutput, error) {
	if params == nil {
		params = &ListLoggingConfigurationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListLoggingConfigurations", params, optFns, addOperationListLoggingConfigurationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListLoggingConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListLoggingConfigurationsInput struct {

	// Specifies the number of LoggingConfigurations that you want AWS WAF to return
	// for this request. If you have more LoggingConfigurations than the number that
	// you specify for Limit, the response includes a NextMarker value that you can use
	// to get another batch of LoggingConfigurations.
	Limit *int32

	// If you specify a value for Limit and you have more LoggingConfigurations than
	// the value of Limit, AWS WAF returns a NextMarker value in the response that
	// allows you to list another group of LoggingConfigurations. For the second and
	// subsequent ListLoggingConfigurations requests, specify the value of NextMarker
	// from the previous response to get information about another batch of
	// ListLoggingConfigurations.
	NextMarker *string
}

type ListLoggingConfigurationsOutput struct {

	// An array of LoggingConfiguration objects.
	LoggingConfigurations []*types.LoggingConfiguration

	// If you have more LoggingConfigurations than the number that you specified for
	// Limit in the request, the response includes a NextMarker value. To list more
	// LoggingConfigurations, submit another ListLoggingConfigurations request, and
	// specify the NextMarker value from the response in the NextMarker value in the
	// next request.
	NextMarker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListLoggingConfigurationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpListLoggingConfigurations{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpListLoggingConfigurations{})
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
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opListLoggingConfigurations(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opListLoggingConfigurations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf-regional",
		OperationName: "ListLoggingConfigurations",
	}
}
