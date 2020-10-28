// Code generated by smithy-go-codegen DO NOT EDIT.

package waf

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/waf/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This is AWS WAF Classic documentation. For more information, see AWS WAF Classic
// (https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html)
// in the developer guide. For the latest version of AWS WAF, use the AWS WAFV2 API
// and see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html). With
// the latest version, AWS WAF has a single set of endpoints for regional and
// global use. Inserts or deletes SizeConstraint objects (filters) in a
// SizeConstraintSet. For each SizeConstraint object, you specify the following
// values:
//
//     * Whether to insert or delete the object from the array. If you
// want to change a SizeConstraintSetUpdate object, you delete the existing object
// and add a new one.
//
//     * The part of a web request that you want AWS WAF to
// evaluate, such as the length of a query string or the length of the User-Agent
// header.
//
//     * Whether to perform any transformations on the request, such as
// converting it to lowercase, before checking its length. Note that
// transformations of the request body are not supported because the AWS resource
// forwards only the first 8192 bytes of your request to AWS WAF. You can only
// specify a single type of TextTransformation.
//
//     * A ComparisonOperator used
// for evaluating the selected part of the request against the specified Size, such
// as equals, greater than, less than, and so on.
//
//     * The length, in bytes, that
// you want AWS WAF to watch for in selected part of the request. The length is
// computed after applying the transformation.
//
// For example, you can add a
// SizeConstraintSetUpdate object that matches web requests in which the length of
// the User-Agent header is greater than 100 bytes. You can then configure AWS WAF
// to block those requests. To create and configure a SizeConstraintSet, perform
// the following steps:
//
//     * Create a SizeConstraintSet. For more information,
// see CreateSizeConstraintSet.
//
//     * Use GetChangeToken to get the change token
// that you provide in the ChangeToken parameter of an UpdateSizeConstraintSet
// request.
//
//     * Submit an UpdateSizeConstraintSet request to specify the part of
// the request that you want AWS WAF to inspect (for example, the header or the
// URI) and the value that you want AWS WAF to watch for.
//
// For more information
// about how to use the AWS WAF API to allow or block HTTP requests, see the AWS
// WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/).
func (c *Client) UpdateSizeConstraintSet(ctx context.Context, params *UpdateSizeConstraintSetInput, optFns ...func(*Options)) (*UpdateSizeConstraintSetOutput, error) {
	if params == nil {
		params = &UpdateSizeConstraintSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateSizeConstraintSet", params, optFns, addOperationUpdateSizeConstraintSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateSizeConstraintSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateSizeConstraintSetInput struct {

	// The value returned by the most recent call to GetChangeToken.
	//
	// This member is required.
	ChangeToken *string

	// The SizeConstraintSetId of the SizeConstraintSet that you want to update.
	// SizeConstraintSetId is returned by CreateSizeConstraintSet and by
	// ListSizeConstraintSets.
	//
	// This member is required.
	SizeConstraintSetId *string

	// An array of SizeConstraintSetUpdate objects that you want to insert into or
	// delete from a SizeConstraintSet. For more information, see the applicable data
	// types:
	//
	//     * SizeConstraintSetUpdate: Contains Action and SizeConstraint
	//
	//     *
	// SizeConstraint: Contains FieldToMatch, TextTransformation, ComparisonOperator,
	// and Size
	//
	//     * FieldToMatch: Contains Data and Type
	//
	// This member is required.
	Updates []*types.SizeConstraintSetUpdate
}

type UpdateSizeConstraintSetOutput struct {

	// The ChangeToken that you used to submit the UpdateSizeConstraintSet request. You
	// can also use this value to query the status of the request. For more
	// information, see GetChangeTokenStatus.
	ChangeToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateSizeConstraintSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpUpdateSizeConstraintSet{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpUpdateSizeConstraintSet{})
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
	if err := addOpUpdateSizeConstraintSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opUpdateSizeConstraintSet(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opUpdateSizeConstraintSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf",
		OperationName: "UpdateSizeConstraintSet",
	}
}
