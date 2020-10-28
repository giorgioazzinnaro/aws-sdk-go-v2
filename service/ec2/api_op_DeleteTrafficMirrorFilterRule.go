// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the specified Traffic Mirror rule.
func (c *Client) DeleteTrafficMirrorFilterRule(ctx context.Context, params *DeleteTrafficMirrorFilterRuleInput, optFns ...func(*Options)) (*DeleteTrafficMirrorFilterRuleOutput, error) {
	if params == nil {
		params = &DeleteTrafficMirrorFilterRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteTrafficMirrorFilterRule", params, optFns, addOperationDeleteTrafficMirrorFilterRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteTrafficMirrorFilterRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteTrafficMirrorFilterRuleInput struct {

	// The ID of the Traffic Mirror rule.
	//
	// This member is required.
	TrafficMirrorFilterRuleId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type DeleteTrafficMirrorFilterRuleOutput struct {

	// The ID of the deleted Traffic Mirror rule.
	TrafficMirrorFilterRuleId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteTrafficMirrorFilterRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsEc2query_serializeOpDeleteTrafficMirrorFilterRule{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsEc2query_deserializeOpDeleteTrafficMirrorFilterRule{})
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
	if err := addOpDeleteTrafficMirrorFilterRuleValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDeleteTrafficMirrorFilterRule(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDeleteTrafficMirrorFilterRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DeleteTrafficMirrorFilterRule",
	}
}
