// Code generated by smithy-go-codegen DO NOT EDIT.

package groundstation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/groundstation/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a Config.
func (c *Client) DeleteConfig(ctx context.Context, params *DeleteConfigInput, optFns ...func(*Options)) (*DeleteConfigOutput, error) {
	if params == nil {
		params = &DeleteConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteConfig", params, optFns, addOperationDeleteConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DeleteConfigInput struct {

	// UUID of a Config.
	//
	// This member is required.
	ConfigId *string

	// Type of a Config.
	//
	// This member is required.
	ConfigType types.ConfigCapabilityType
}

//
type DeleteConfigOutput struct {

	// ARN of a Config.
	ConfigArn *string

	// UUID of a Config.
	ConfigId *string

	// Type of a Config.
	ConfigType types.ConfigCapabilityType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpDeleteConfig{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpDeleteConfig{})
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
	if err := addOpDeleteConfigValidationMiddleware(stack); err != nil {
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
