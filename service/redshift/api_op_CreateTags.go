// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds tags to a cluster. A resource can have up to 50 tags. If you try to create
// more than 50 tags for a resource, you will receive an error and the attempt will
// fail. If you specify a key that already exists for the resource, the value for
// that key will be updated with the new value.
func (c *Client) CreateTags(ctx context.Context, params *CreateTagsInput, optFns ...func(*Options)) (*CreateTagsOutput, error) {
	if params == nil {
		params = &CreateTagsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateTags", params, optFns, addOperationCreateTagsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateTagsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the output from the CreateTags action.
type CreateTagsInput struct {

	// The Amazon Resource Name (ARN) to which you want to add the tag or tags. For
	// example, arn:aws:redshift:us-east-2:123456789:cluster:t1.
	//
	// This member is required.
	ResourceName *string

	// One or more name/value pairs to add as tags to the specified resource. Each tag
	// name is passed in with the parameter Key and the corresponding value is passed
	// in with the parameter Value. The Key and Value parameters are separated by a
	// comma (,). Separate multiple tags with a space. For example, --tags
	// "Key"="owner","Value"="admin" "Key"="environment","Value"="test"
	// "Key"="version","Value"="1.0".
	//
	// This member is required.
	Tags []*types.Tag
}

type CreateTagsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateTagsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsquery_serializeOpCreateTags{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsquery_deserializeOpCreateTags{})
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
	if err := addOpCreateTagsValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opCreateTags(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opCreateTags(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "CreateTags",
	}
}
