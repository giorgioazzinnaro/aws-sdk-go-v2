// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes an existing MethodResponse resource.
func (c *Client) DeleteMethodResponse(ctx context.Context, params *DeleteMethodResponseInput, optFns ...func(*Options)) (*DeleteMethodResponseOutput, error) {
	if params == nil {
		params = &DeleteMethodResponseInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteMethodResponse", params, optFns, addOperationDeleteMethodResponseMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteMethodResponseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to delete an existing MethodResponse resource.
type DeleteMethodResponseInput struct {

	// [Required] The HTTP verb of the Method resource.
	//
	// This member is required.
	HttpMethod *string

	// [Required] The Resource identifier for the MethodResponse resource.
	//
	// This member is required.
	ResourceId *string

	// [Required] The string identifier of the associated RestApi.
	//
	// This member is required.
	RestApiId *string

	// [Required] The status code identifier for the MethodResponse resource.
	//
	// This member is required.
	StatusCode *string

	Name *string

	Template *bool

	TemplateSkipList []*string

	Title *string
}

type DeleteMethodResponseOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteMethodResponseMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpDeleteMethodResponse{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpDeleteMethodResponse{})
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
	if err := addOpDeleteMethodResponseValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDeleteMethodResponse(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDeleteMethodResponse(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "DeleteMethodResponse",
	}
}
