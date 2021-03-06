// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisvideo

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds one or more tags to a stream. A tag is a key-value pair (the value is
// optional) that you can define and assign to AWS resources. If you specify a tag
// that already exists, the tag value is replaced with the value that you specify
// in the request. For more information, see Using Cost Allocation Tags
// (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html)
// in the AWS Billing and Cost Management User Guide. You must provide either the
// StreamName or the StreamARN. This operation requires permission for the
// KinesisVideo:TagStream action. Kinesis video streams support up to 50 tags.
func (c *Client) TagStream(ctx context.Context, params *TagStreamInput, optFns ...func(*Options)) (*TagStreamOutput, error) {
	if params == nil {
		params = &TagStreamInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "TagStream", params, optFns, addOperationTagStreamMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*TagStreamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TagStreamInput struct {

	// A list of tags to associate with the specified stream. Each tag is a key-value
	// pair (the value is optional).
	//
	// This member is required.
	Tags map[string]*string

	// The Amazon Resource Name (ARN) of the resource that you want to add the tag or
	// tags to.
	StreamARN *string

	// The name of the stream that you want to add the tag or tags to.
	StreamName *string
}

type TagStreamOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationTagStreamMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpTagStream{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpTagStream{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpTagStreamValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opTagStream(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opTagStream(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisvideo",
		OperationName: "TagStream",
	}
}
