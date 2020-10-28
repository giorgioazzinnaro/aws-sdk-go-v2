// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Get the information about an origin access identity.
func (c *Client) GetCloudFrontOriginAccessIdentity(ctx context.Context, params *GetCloudFrontOriginAccessIdentityInput, optFns ...func(*Options)) (*GetCloudFrontOriginAccessIdentityOutput, error) {
	if params == nil {
		params = &GetCloudFrontOriginAccessIdentityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCloudFrontOriginAccessIdentity", params, optFns, addOperationGetCloudFrontOriginAccessIdentityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCloudFrontOriginAccessIdentityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request to get an origin access identity's information.
type GetCloudFrontOriginAccessIdentityInput struct {

	// The identity's ID.
	//
	// This member is required.
	Id *string
}

// The returned result of the corresponding request.
type GetCloudFrontOriginAccessIdentityOutput struct {

	// The origin access identity's information.
	CloudFrontOriginAccessIdentity *types.CloudFrontOriginAccessIdentity

	// The current version of the origin access identity's information. For example:
	// E2QWRUHAPOMQZL.
	ETag *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetCloudFrontOriginAccessIdentityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestxml_serializeOpGetCloudFrontOriginAccessIdentity{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestxml_deserializeOpGetCloudFrontOriginAccessIdentity{})
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
	if err := addOpGetCloudFrontOriginAccessIdentityValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opGetCloudFrontOriginAccessIdentity(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opGetCloudFrontOriginAccessIdentity(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "GetCloudFrontOriginAccessIdentity",
	}
}
