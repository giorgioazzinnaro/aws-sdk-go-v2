// Code generated by smithy-go-codegen DO NOT EDIT.

package restxml

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/restxml/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// The example tests how requests and responses are serialized when there is no
// input or output payload but there are HTTP header bindings.
func (c *Client) InputAndOutputWithHeaders(ctx context.Context, params *InputAndOutputWithHeadersInput, optFns ...func(*Options)) (*InputAndOutputWithHeadersOutput, error) {
	if params == nil {
		params = &InputAndOutputWithHeadersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "InputAndOutputWithHeaders", params, optFns, addOperationInputAndOutputWithHeadersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*InputAndOutputWithHeadersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type InputAndOutputWithHeadersInput struct {
	HeaderBooleanList []*bool

	HeaderByte *int8

	HeaderDouble *float64

	HeaderEnum types.FooEnum

	HeaderEnumList []types.FooEnum

	HeaderFalseBool *bool

	HeaderFloat *float32

	HeaderInteger *int32

	HeaderIntegerList []*int32

	HeaderLong *int64

	HeaderShort *int16

	HeaderString *string

	HeaderStringList []*string

	HeaderStringSet []*string

	HeaderTimestampList []*time.Time

	HeaderTrueBool *bool
}

type InputAndOutputWithHeadersOutput struct {
	HeaderBooleanList []*bool

	HeaderByte *int8

	HeaderDouble *float64

	HeaderEnum types.FooEnum

	HeaderEnumList []types.FooEnum

	HeaderFalseBool *bool

	HeaderFloat *float32

	HeaderInteger *int32

	HeaderIntegerList []*int32

	HeaderLong *int64

	HeaderShort *int16

	HeaderString *string

	HeaderStringList []*string

	HeaderStringSet []*string

	HeaderTimestampList []*time.Time

	HeaderTrueBool *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationInputAndOutputWithHeadersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestxml_serializeOpInputAndOutputWithHeaders{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestxml_deserializeOpInputAndOutputWithHeaders{})
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
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opInputAndOutputWithHeaders(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opInputAndOutputWithHeaders(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "InputAndOutputWithHeaders",
	}
}
