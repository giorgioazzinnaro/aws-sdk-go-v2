// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the comment for a specified hosted zone.
func (c *Client) UpdateHostedZoneComment(ctx context.Context, params *UpdateHostedZoneCommentInput, optFns ...func(*Options)) (*UpdateHostedZoneCommentOutput, error) {
	if params == nil {
		params = &UpdateHostedZoneCommentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateHostedZoneComment", params, optFns, addOperationUpdateHostedZoneCommentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateHostedZoneCommentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to update the comment for a hosted zone.
type UpdateHostedZoneCommentInput struct {

	// The ID for the hosted zone that you want to update the comment for.
	//
	// This member is required.
	Id *string

	// The new comment for the hosted zone. If you don't specify a value for Comment,
	// Amazon Route 53 deletes the existing value of the Comment element, if any.
	Comment *string
}

// A complex type that contains the response to the UpdateHostedZoneComment
// request.
type UpdateHostedZoneCommentOutput struct {

	// A complex type that contains the response to the UpdateHostedZoneComment
	// request.
	//
	// This member is required.
	HostedZone *types.HostedZone

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateHostedZoneCommentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestxml_serializeOpUpdateHostedZoneComment{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestxml_deserializeOpUpdateHostedZoneComment{})
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
	if err := addOpUpdateHostedZoneCommentValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opUpdateHostedZoneComment(options.Region)); err != nil {
		return err
	}
	if err := addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err := addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err := addSanitizeURLMiddleware(stack); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opUpdateHostedZoneComment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "UpdateHostedZoneComment",
	}
}
