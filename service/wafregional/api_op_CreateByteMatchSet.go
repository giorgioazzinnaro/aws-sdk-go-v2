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
// global use. Creates a ByteMatchSet. You then use UpdateByteMatchSet to identify
// the part of a web request that you want AWS WAF to inspect, such as the values
// of the User-Agent header or the query string. For example, you can create a
// ByteMatchSet that matches any requests with User-Agent headers that contain the
// string BadBot. You can then configure AWS WAF to reject those requests. To
// create and configure a ByteMatchSet, perform the following steps:
//
//     * Use
// GetChangeToken to get the change token that you provide in the ChangeToken
// parameter of a CreateByteMatchSet request.
//
//     * Submit a CreateByteMatchSet
// request.
//
//     * Use GetChangeToken to get the change token that you provide in
// the ChangeToken parameter of an UpdateByteMatchSet request.
//
//     * Submit an
// UpdateByteMatchSet request to specify the part of the request that you want AWS
// WAF to inspect (for example, the header or the URI) and the value that you want
// AWS WAF to watch for.
//
// For more information about how to use the AWS WAF API to
// allow or block HTTP requests, see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/).
func (c *Client) CreateByteMatchSet(ctx context.Context, params *CreateByteMatchSetInput, optFns ...func(*Options)) (*CreateByteMatchSetOutput, error) {
	if params == nil {
		params = &CreateByteMatchSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateByteMatchSet", params, optFns, addOperationCreateByteMatchSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateByteMatchSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateByteMatchSetInput struct {

	// The value returned by the most recent call to GetChangeToken.
	//
	// This member is required.
	ChangeToken *string

	// A friendly name or description of the ByteMatchSet. You can't change Name after
	// you create a ByteMatchSet.
	//
	// This member is required.
	Name *string
}

type CreateByteMatchSetOutput struct {

	// A ByteMatchSet that contains no ByteMatchTuple objects.
	ByteMatchSet *types.ByteMatchSet

	// The ChangeToken that you used to submit the CreateByteMatchSet request. You can
	// also use this value to query the status of the request. For more information,
	// see GetChangeTokenStatus.
	ChangeToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateByteMatchSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpCreateByteMatchSet{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpCreateByteMatchSet{})
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
	if err := addOpCreateByteMatchSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opCreateByteMatchSet(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opCreateByteMatchSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf-regional",
		OperationName: "CreateByteMatchSet",
	}
}
