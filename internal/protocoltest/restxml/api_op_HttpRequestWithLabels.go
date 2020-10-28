// Code generated by smithy-go-codegen DO NOT EDIT.

package restxml

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// The example tests how requests are serialized when there's no input payload but
// there are HTTP labels.
func (c *Client) HttpRequestWithLabels(ctx context.Context, params *HttpRequestWithLabelsInput, optFns ...func(*Options)) (*HttpRequestWithLabelsOutput, error) {
	if params == nil {
		params = &HttpRequestWithLabelsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "HttpRequestWithLabels", params, optFns, addOperationHttpRequestWithLabelsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*HttpRequestWithLabelsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type HttpRequestWithLabelsInput struct {

	// Serialized in the path as true or false.
	//
	// This member is required.
	Boolean *bool

	Double *float64

	Float *float32

	Integer *int32

	Long *int64

	Short *int16

	String_ *string

	// Note that this member has no format, so it's serialized as an RFC 3399
	// date-time.
	//
	// This member is required.
	Timestamp *time.Time
}

type HttpRequestWithLabelsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationHttpRequestWithLabelsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestxml_serializeOpHttpRequestWithLabels{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestxml_deserializeOpHttpRequestWithLabels{})
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
	if err := addOpHttpRequestWithLabelsValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opHttpRequestWithLabels(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opHttpRequestWithLabels(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "HttpRequestWithLabels",
	}
}
