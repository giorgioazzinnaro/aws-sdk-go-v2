// Code generated by smithy-go-codegen DO NOT EDIT.

package workspaces

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the specified tags from the specified WorkSpaces resource.
func (c *Client) DeleteTags(ctx context.Context, params *DeleteTagsInput, optFns ...func(*Options)) (*DeleteTagsOutput, error) {
	if params == nil {
		params = &DeleteTagsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteTags", params, optFns, addOperationDeleteTagsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteTagsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteTagsInput struct {

	// The identifier of the WorkSpaces resource. The supported resource types are
	// WorkSpaces, registered directories, images, custom bundles, IP access control
	// groups, and connection aliases.
	//
	// This member is required.
	ResourceId *string

	// The tag keys.
	//
	// This member is required.
	TagKeys []*string
}

type DeleteTagsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteTagsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpDeleteTags{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpDeleteTags{})
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
	if err := addOpDeleteTagsValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDeleteTags(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDeleteTags(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces",
		OperationName: "DeleteTags",
	}
}
