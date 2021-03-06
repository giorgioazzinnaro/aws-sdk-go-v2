// Code generated by smithy-go-codegen DO NOT EDIT.

package appflow

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Enables you to delete an existing connector profile.
func (c *Client) DeleteConnectorProfile(ctx context.Context, params *DeleteConnectorProfileInput, optFns ...func(*Options)) (*DeleteConnectorProfileOutput, error) {
	if params == nil {
		params = &DeleteConnectorProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteConnectorProfile", params, optFns, addOperationDeleteConnectorProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteConnectorProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteConnectorProfileInput struct {

	// The name of the connector profile. The name is unique for each ConnectorProfile
	// in your account.
	//
	// This member is required.
	ConnectorProfileName *string

	// Indicates whether Amazon AppFlow should delete the profile, even if it is
	// currently in use in one or more flows.
	ForceDelete *bool
}

type DeleteConnectorProfileOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteConnectorProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteConnectorProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteConnectorProfile{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpDeleteConnectorProfileValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteConnectorProfile(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteConnectorProfile(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appflow",
		OperationName: "DeleteConnectorProfile",
	}
}
