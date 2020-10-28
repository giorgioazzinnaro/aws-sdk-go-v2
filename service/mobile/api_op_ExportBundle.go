// Code generated by smithy-go-codegen DO NOT EDIT.

package mobile

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mobile/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Generates customized software development kit (SDK) and or tool packages used to
// integrate mobile web or mobile app clients with backend AWS resources.
func (c *Client) ExportBundle(ctx context.Context, params *ExportBundleInput, optFns ...func(*Options)) (*ExportBundleOutput, error) {
	if params == nil {
		params = &ExportBundleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ExportBundle", params, optFns, addOperationExportBundleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ExportBundleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request structure used to request generation of custom SDK and tool packages
// required to integrate mobile web or app clients with backed AWS resources.
type ExportBundleInput struct {

	// Unique bundle identifier.
	//
	// This member is required.
	BundleId *string

	// Developer desktop or target application platform.
	Platform types.Platform

	// Unique project identifier.
	ProjectId *string
}

// Result structure which contains link to download custom-generated SDK and tool
// packages used to integrate mobile web or app clients with backed AWS resources.
type ExportBundleOutput struct {

	// URL which contains the custom-generated SDK and tool packages used to integrate
	// the client mobile app or web app with the AWS resources created by the AWS
	// Mobile Hub project.
	DownloadUrl *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationExportBundleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpExportBundle{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpExportBundle{})
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
	if err := addOpExportBundleValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opExportBundle(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opExportBundle(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "awsmobilehubservice",
		OperationName: "ExportBundle",
	}
}
