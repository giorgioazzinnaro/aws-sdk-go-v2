// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the item that the specified template alias points to. If you provide a
// specific alias, you delete the version of the template that the alias points to.
func (c *Client) DeleteTemplateAlias(ctx context.Context, params *DeleteTemplateAliasInput, optFns ...func(*Options)) (*DeleteTemplateAliasOutput, error) {
	if params == nil {
		params = &DeleteTemplateAliasInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteTemplateAlias", params, optFns, addOperationDeleteTemplateAliasMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteTemplateAliasOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteTemplateAliasInput struct {

	// The name for the template alias. To delete a specific alias, you delete the
	// version that the alias points to. You can specify the alias name, or specify the
	// latest version of the template by providing the keyword $LATEST in the AliasName
	// parameter.
	//
	// This member is required.
	AliasName *string

	// The ID of the AWS account that contains the item to delete.
	//
	// This member is required.
	AwsAccountId *string

	// The ID for the template that the specified alias is for.
	//
	// This member is required.
	TemplateId *string
}

type DeleteTemplateAliasOutput struct {

	// The name for the template alias.
	AliasName *string

	// The Amazon Resource Name (ARN) of the template you want to delete.
	Arn *string

	// The AWS request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status *int32

	// An ID for the template associated with the deletion.
	TemplateId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteTemplateAliasMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpDeleteTemplateAlias{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpDeleteTemplateAlias{})
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
	if err := addOpDeleteTemplateAliasValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDeleteTemplateAlias(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDeleteTemplateAlias(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "DeleteTemplateAlias",
	}
}
