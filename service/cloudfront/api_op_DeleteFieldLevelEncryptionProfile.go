// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Remove a field-level encryption profile.
func (c *Client) DeleteFieldLevelEncryptionProfile(ctx context.Context, params *DeleteFieldLevelEncryptionProfileInput, optFns ...func(*Options)) (*DeleteFieldLevelEncryptionProfileOutput, error) {
	if params == nil {
		params = &DeleteFieldLevelEncryptionProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteFieldLevelEncryptionProfile", params, optFns, addOperationDeleteFieldLevelEncryptionProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteFieldLevelEncryptionProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteFieldLevelEncryptionProfileInput struct {

	// Request the ID of the profile you want to delete from CloudFront.
	//
	// This member is required.
	Id *string

	// The value of the ETag header that you received when retrieving the profile to
	// delete. For example: E2QWRUHAPOMQZL.
	IfMatch *string
}

type DeleteFieldLevelEncryptionProfileOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteFieldLevelEncryptionProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestxml_serializeOpDeleteFieldLevelEncryptionProfile{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestxml_deserializeOpDeleteFieldLevelEncryptionProfile{})
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
	if err := addOpDeleteFieldLevelEncryptionProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDeleteFieldLevelEncryptionProfile(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDeleteFieldLevelEncryptionProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "DeleteFieldLevelEncryptionProfile",
	}
}
