// Code generated by smithy-go-codegen DO NOT EDIT.

package query

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// This tests how timestamps are serialized, including using the default format of
// date-time and various @timestampFormat trait values.
func (c *Client) XmlTimestamps(ctx context.Context, params *XmlTimestampsInput, optFns ...func(*Options)) (*XmlTimestampsOutput, error) {
	if params == nil {
		params = &XmlTimestampsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "XmlTimestamps", params, optFns, addOperationXmlTimestampsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*XmlTimestampsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type XmlTimestampsInput struct {
}

type XmlTimestampsOutput struct {
	DateTime *time.Time

	EpochSeconds *time.Time

	HttpDate *time.Time

	Normal *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationXmlTimestampsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsquery_serializeOpXmlTimestamps{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsquery_deserializeOpXmlTimestamps{})
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
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opXmlTimestamps(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opXmlTimestamps(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "XmlTimestamps",
	}
}
