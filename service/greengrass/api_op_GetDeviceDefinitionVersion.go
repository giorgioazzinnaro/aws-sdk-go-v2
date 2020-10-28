// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrass

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves information about a device definition version.
func (c *Client) GetDeviceDefinitionVersion(ctx context.Context, params *GetDeviceDefinitionVersionInput, optFns ...func(*Options)) (*GetDeviceDefinitionVersionOutput, error) {
	if params == nil {
		params = &GetDeviceDefinitionVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDeviceDefinitionVersion", params, optFns, addOperationGetDeviceDefinitionVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDeviceDefinitionVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDeviceDefinitionVersionInput struct {

	// The ID of the device definition.
	//
	// This member is required.
	DeviceDefinitionId *string

	// The ID of the device definition version. This value maps to the ''Version''
	// property of the corresponding ''VersionInformation'' object, which is returned
	// by ''ListDeviceDefinitionVersions'' requests. If the version is the last one
	// that was associated with a device definition, the value also maps to the
	// ''LatestVersion'' property of the corresponding ''DefinitionInformation''
	// object.
	//
	// This member is required.
	DeviceDefinitionVersionId *string

	// The token for the next set of results, or ''null'' if there are no additional
	// results.
	NextToken *string
}

type GetDeviceDefinitionVersionOutput struct {

	// The ARN of the device definition version.
	Arn *string

	// The time, in milliseconds since the epoch, when the device definition version
	// was created.
	CreationTimestamp *string

	// Information about the device definition version.
	Definition *types.DeviceDefinitionVersion

	// The ID of the device definition version.
	Id *string

	// The token for the next set of results, or ''null'' if there are no additional
	// results.
	NextToken *string

	// The version of the device definition version.
	Version *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetDeviceDefinitionVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpGetDeviceDefinitionVersion{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpGetDeviceDefinitionVersion{})
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
	if err := addOpGetDeviceDefinitionVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opGetDeviceDefinitionVersion(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opGetDeviceDefinitionVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "GetDeviceDefinitionVersion",
	}
}
