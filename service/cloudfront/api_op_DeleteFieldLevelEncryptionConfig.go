// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Remove a field-level encryption configuration.
func (c *Client) DeleteFieldLevelEncryptionConfig(ctx context.Context, params *DeleteFieldLevelEncryptionConfigInput, optFns ...func(*Options)) (*DeleteFieldLevelEncryptionConfigOutput, error) {
	if params == nil {
		params = &DeleteFieldLevelEncryptionConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteFieldLevelEncryptionConfig", params, optFns, addOperationDeleteFieldLevelEncryptionConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteFieldLevelEncryptionConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteFieldLevelEncryptionConfigInput struct {

	// The ID of the configuration you want to delete from CloudFront.
	//
	// This member is required.
	Id *string

	// The value of the ETag header that you received when retrieving the configuration
	// identity to delete. For example: E2QWRUHAPOMQZL.
	IfMatch *string
}

type DeleteFieldLevelEncryptionConfigOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteFieldLevelEncryptionConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpDeleteFieldLevelEncryptionConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpDeleteFieldLevelEncryptionConfig{}, middleware.After)
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
	addOpDeleteFieldLevelEncryptionConfigValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteFieldLevelEncryptionConfig(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteFieldLevelEncryptionConfig(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "DeleteFieldLevelEncryptionConfig",
	}
}
