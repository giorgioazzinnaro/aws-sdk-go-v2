// Code generated by smithy-go-codegen DO NOT EDIT.

package iotthingsgraph

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the latest version of the user's namespace and the public version that it
// is tracking.
func (c *Client) DescribeNamespace(ctx context.Context, params *DescribeNamespaceInput, optFns ...func(*Options)) (*DescribeNamespaceOutput, error) {
	if params == nil {
		params = &DescribeNamespaceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeNamespace", params, optFns, addOperationDescribeNamespaceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeNamespaceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeNamespaceInput struct {

	// The name of the user's namespace. Set this to aws to get the public namespace.
	NamespaceName *string
}

type DescribeNamespaceOutput struct {

	// The ARN of the namespace.
	NamespaceArn *string

	// The name of the namespace.
	NamespaceName *string

	// The version of the user's namespace to describe.
	NamespaceVersion *int64

	// The name of the public namespace that the latest namespace version is tracking.
	TrackingNamespaceName *string

	// The version of the public namespace that the latest version is tracking.
	TrackingNamespaceVersion *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeNamespaceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpDescribeNamespace{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpDescribeNamespace{})
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
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDescribeNamespace(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDescribeNamespace(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotthingsgraph",
		OperationName: "DescribeNamespace",
	}
}
