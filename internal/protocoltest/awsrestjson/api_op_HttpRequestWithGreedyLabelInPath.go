// Code generated by smithy-go-codegen DO NOT EDIT.

package awsrestjson

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

func (c *Client) HttpRequestWithGreedyLabelInPath(ctx context.Context, params *HttpRequestWithGreedyLabelInPathInput, optFns ...func(*Options)) (*HttpRequestWithGreedyLabelInPathOutput, error) {
	if params == nil {
		params = &HttpRequestWithGreedyLabelInPathInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "HttpRequestWithGreedyLabelInPath", params, optFns, addOperationHttpRequestWithGreedyLabelInPathMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*HttpRequestWithGreedyLabelInPathOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type HttpRequestWithGreedyLabelInPathInput struct {
	Baz *string

	Foo *string
}

type HttpRequestWithGreedyLabelInPathOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationHttpRequestWithGreedyLabelInPathMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpHttpRequestWithGreedyLabelInPath{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpHttpRequestWithGreedyLabelInPath{})
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
	if err := addRetryMiddlewares(stack, options); err != nil {
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
	if err := addOpHttpRequestWithGreedyLabelInPathValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opHttpRequestWithGreedyLabelInPath(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opHttpRequestWithGreedyLabelInPath(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "HttpRequestWithGreedyLabelInPath",
	}
}
