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

// Retrieves information about a logger definition version.
func (c *Client) GetLoggerDefinitionVersion(ctx context.Context, params *GetLoggerDefinitionVersionInput, optFns ...func(*Options)) (*GetLoggerDefinitionVersionOutput, error) {
	if params == nil {
		params = &GetLoggerDefinitionVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetLoggerDefinitionVersion", params, optFns, addOperationGetLoggerDefinitionVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetLoggerDefinitionVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetLoggerDefinitionVersionInput struct {

	// The ID of the logger definition.
	//
	// This member is required.
	LoggerDefinitionId *string

	// The ID of the logger definition version. This value maps to the ''Version''
	// property of the corresponding ''VersionInformation'' object, which is returned
	// by ''ListLoggerDefinitionVersions'' requests. If the version is the last one
	// that was associated with a logger definition, the value also maps to the
	// ''LatestVersion'' property of the corresponding ''DefinitionInformation''
	// object.
	//
	// This member is required.
	LoggerDefinitionVersionId *string

	// The token for the next set of results, or ''null'' if there are no additional
	// results.
	NextToken *string
}

type GetLoggerDefinitionVersionOutput struct {

	// The ARN of the logger definition version.
	Arn *string

	// The time, in milliseconds since the epoch, when the logger definition version
	// was created.
	CreationTimestamp *string

	// Information about the logger definition version.
	Definition *types.LoggerDefinitionVersion

	// The ID of the logger definition version.
	Id *string

	// The version of the logger definition version.
	Version *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetLoggerDefinitionVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpGetLoggerDefinitionVersion{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpGetLoggerDefinitionVersion{})
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
	if err := addOpGetLoggerDefinitionVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opGetLoggerDefinitionVersion(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opGetLoggerDefinitionVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "GetLoggerDefinitionVersion",
	}
}
