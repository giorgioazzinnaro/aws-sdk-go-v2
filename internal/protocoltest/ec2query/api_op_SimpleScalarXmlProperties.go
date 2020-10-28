// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2query

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

func (c *Client) SimpleScalarXmlProperties(ctx context.Context, params *SimpleScalarXmlPropertiesInput, optFns ...func(*Options)) (*SimpleScalarXmlPropertiesOutput, error) {
	if params == nil {
		params = &SimpleScalarXmlPropertiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SimpleScalarXmlProperties", params, optFns, addOperationSimpleScalarXmlPropertiesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SimpleScalarXmlPropertiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SimpleScalarXmlPropertiesInput struct {
}

type SimpleScalarXmlPropertiesOutput struct {
	ByteValue *int8

	DoubleValue *float64

	EmptyStringValue *string

	FalseBooleanValue *bool

	FloatValue *float32

	IntegerValue *int32

	LongValue *int64

	ShortValue *int16

	StringValue *string

	TrueBooleanValue *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationSimpleScalarXmlPropertiesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsEc2query_serializeOpSimpleScalarXmlProperties{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsEc2query_deserializeOpSimpleScalarXmlProperties{})
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
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opSimpleScalarXmlProperties(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opSimpleScalarXmlProperties(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SimpleScalarXmlProperties",
	}
}
