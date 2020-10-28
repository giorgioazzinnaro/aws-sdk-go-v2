// Code generated by smithy-go-codegen DO NOT EDIT.

package codestarconnections

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codestarconnections/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a connection that can then be given to other AWS services like
// CodePipeline so that it can access third-party code repositories. The connection
// is in pending status until the third-party connection handshake is completed
// from the console.
func (c *Client) CreateConnection(ctx context.Context, params *CreateConnectionInput, optFns ...func(*Options)) (*CreateConnectionOutput, error) {
	if params == nil {
		params = &CreateConnectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateConnection", params, optFns, addOperationCreateConnectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateConnectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateConnectionInput struct {

	// The name of the connection to be created. The name must be unique in the calling
	// AWS account.
	//
	// This member is required.
	ConnectionName *string

	// The Amazon Resource Name (ARN) of the host associated with the connection to be
	// created.
	HostArn *string

	// The name of the external provider where your third-party code repository is
	// configured. The valid provider type is Bitbucket.
	ProviderType types.ProviderType

	// The key-value pair to use when tagging the resource.
	Tags []*types.Tag
}

type CreateConnectionOutput struct {

	// The Amazon Resource Name (ARN) of the connection to be created. The ARN is used
	// as the connection reference when the connection is shared between AWS services.
	// The ARN is never reused if the connection is deleted.
	//
	// This member is required.
	ConnectionArn *string

	// Specifies the tags applied to the resource.
	Tags []*types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateConnectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson10_serializeOpCreateConnection{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson10_deserializeOpCreateConnection{})
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
	if err := addOpCreateConnectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opCreateConnection(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opCreateConnection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codestar-connections",
		OperationName: "CreateConnection",
	}
}
