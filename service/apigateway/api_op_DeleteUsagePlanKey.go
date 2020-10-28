// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a usage plan key and remove the underlying API key from the associated
// usage plan.
func (c *Client) DeleteUsagePlanKey(ctx context.Context, params *DeleteUsagePlanKeyInput, optFns ...func(*Options)) (*DeleteUsagePlanKeyOutput, error) {
	if params == nil {
		params = &DeleteUsagePlanKeyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteUsagePlanKey", params, optFns, addOperationDeleteUsagePlanKeyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteUsagePlanKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The DELETE request to delete a usage plan key and remove the underlying API key
// from the associated usage plan.
type DeleteUsagePlanKeyInput struct {

	// [Required] The Id of the UsagePlanKey resource to be deleted.
	//
	// This member is required.
	KeyId *string

	// [Required] The Id of the UsagePlan resource representing the usage plan
	// containing the to-be-deleted UsagePlanKey resource representing a plan customer.
	//
	// This member is required.
	UsagePlanId *string

	Name *string

	Template *bool

	TemplateSkipList []*string

	Title *string
}

type DeleteUsagePlanKeyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteUsagePlanKeyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpDeleteUsagePlanKey{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpDeleteUsagePlanKey{})
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
	if err := addOpDeleteUsagePlanKeyValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDeleteUsagePlanKey(options.Region)); err != nil {
		return err
	}
	if err := addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err := addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err := addAcceptHeader(stack); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDeleteUsagePlanKey(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "DeleteUsagePlanKey",
	}
}
