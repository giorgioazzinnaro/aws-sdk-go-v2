// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpoint

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates an existing message template for messages that are sent through the
// voice channel.
func (c *Client) UpdateVoiceTemplate(ctx context.Context, params *UpdateVoiceTemplateInput, optFns ...func(*Options)) (*UpdateVoiceTemplateOutput, error) {
	if params == nil {
		params = &UpdateVoiceTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateVoiceTemplate", params, optFns, addOperationUpdateVoiceTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateVoiceTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateVoiceTemplateInput struct {

	// The name of the message template. A template name must start with an
	// alphanumeric character and can contain a maximum of 128 characters. The
	// characters can be alphanumeric characters, underscores (_), or hyphens (-).
	// Template names are case sensitive.
	//
	// This member is required.
	TemplateName *string

	// Specifies the content and settings for a message template that can be used in
	// messages that are sent through the voice channel.
	//
	// This member is required.
	VoiceTemplateRequest *types.VoiceTemplateRequest

	// Specifies whether to save the updates as a new version of the message template.
	// Valid values are: true, save the updates as a new version; and, false, save the
	// updates to (overwrite) the latest existing version of the template. If you don't
	// specify a value for this parameter, Amazon Pinpoint saves the updates to
	// (overwrites) the latest existing version of the template. If you specify a value
	// of true for this parameter, don't specify a value for the version parameter.
	// Otherwise, an error will occur.
	CreateNewVersion *bool

	// The unique identifier for the version of the message template to update,
	// retrieve information about, or delete. To retrieve identifiers and other
	// information for all the versions of a template, use the Template Versions
	// resource. If specified, this value must match the identifier for an existing
	// template version. If specified for an update operation, this value must match
	// the identifier for the latest existing version of the template. This restriction
	// helps ensure that race conditions don't occur. If you don't specify a value for
	// this parameter, Amazon Pinpoint does the following:
	//
	//     * For a get operation,
	// retrieves information about the active version of the template.
	//
	//     * For an
	// update operation, saves the updates to (overwrites) the latest existing version
	// of the template, if the create-new-version parameter isn't used or is set to
	// false.
	//
	//     * For a delete operation, deletes the template, including all
	// versions of the template.
	Version *string
}

type UpdateVoiceTemplateOutput struct {

	// Provides information about an API request or response.
	//
	// This member is required.
	MessageBody *types.MessageBody

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateVoiceTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpUpdateVoiceTemplate{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpUpdateVoiceTemplate{})
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
	if err := addOpUpdateVoiceTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opUpdateVoiceTemplate(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opUpdateVoiceTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mobiletargeting",
		OperationName: "UpdateVoiceTemplate",
	}
}
