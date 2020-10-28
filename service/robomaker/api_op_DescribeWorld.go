// Code generated by smithy-go-codegen DO NOT EDIT.

package robomaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Describes a world.
func (c *Client) DescribeWorld(ctx context.Context, params *DescribeWorldInput, optFns ...func(*Options)) (*DescribeWorldOutput, error) {
	if params == nil {
		params = &DescribeWorldInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeWorld", params, optFns, addOperationDescribeWorldMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeWorldOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeWorldInput struct {

	// The Amazon Resource Name (arn) of the world you want to describe.
	//
	// This member is required.
	World *string
}

type DescribeWorldOutput struct {

	// The Amazon Resource Name (arn) of the world.
	Arn *string

	// The time, in milliseconds since the epoch, when the world was created.
	CreatedAt *time.Time

	// The Amazon Resource Name (arn) of the world generation job that generated the
	// world.
	GenerationJob *string

	// A map that contains tag keys and tag values that are attached to the world.
	Tags map[string]*string

	// The world template.
	Template *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeWorldMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpDescribeWorld{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpDescribeWorld{})
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
	if err := addOpDescribeWorldValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDescribeWorld(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDescribeWorld(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "robomaker",
		OperationName: "DescribeWorld",
	}
}
