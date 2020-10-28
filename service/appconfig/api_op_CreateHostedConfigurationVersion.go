// Code generated by smithy-go-codegen DO NOT EDIT.

package appconfig

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Create a new configuration in the AppConfig configuration store.
func (c *Client) CreateHostedConfigurationVersion(ctx context.Context, params *CreateHostedConfigurationVersionInput, optFns ...func(*Options)) (*CreateHostedConfigurationVersionOutput, error) {
	if params == nil {
		params = &CreateHostedConfigurationVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateHostedConfigurationVersion", params, optFns, addOperationCreateHostedConfigurationVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateHostedConfigurationVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateHostedConfigurationVersionInput struct {

	// The application ID.
	//
	// This member is required.
	ApplicationId *string

	// The configuration profile ID.
	//
	// This member is required.
	ConfigurationProfileId *string

	// The content of the configuration or the configuration data.
	//
	// This member is required.
	Content []byte

	// A standard MIME type describing the format of the configuration content. For
	// more information, see Content-Type
	// (https://docs.aws.amazon.com/https:/www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.17).
	//
	// This member is required.
	ContentType *string

	// A description of the configuration.
	Description *string

	// An optional locking token used to prevent race conditions from overwriting
	// configuration updates when creating a new version. To ensure your data is not
	// overwritten when creating multiple hosted configuration versions in rapid
	// succession, specify the version of the latest hosted configuration version.
	LatestVersionNumber *int32
}

type CreateHostedConfigurationVersionOutput struct {

	// The application ID.
	ApplicationId *string

	// The configuration profile ID.
	ConfigurationProfileId *string

	// The content of the configuration or the configuration data.
	Content []byte

	// A standard MIME type describing the format of the configuration content. For
	// more information, see Content-Type
	// (https://docs.aws.amazon.com/https:/www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.17).
	ContentType *string

	// A description of the configuration.
	Description *string

	// The configuration version.
	VersionNumber *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateHostedConfigurationVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpCreateHostedConfigurationVersion{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpCreateHostedConfigurationVersion{})
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
	if err := addOpCreateHostedConfigurationVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opCreateHostedConfigurationVersion(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opCreateHostedConfigurationVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appconfig",
		OperationName: "CreateHostedConfigurationVersion",
	}
}
