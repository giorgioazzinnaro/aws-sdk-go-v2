// Code generated by smithy-go-codegen DO NOT EDIT.

package worklink

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates domain metadata, such as DisplayName.
func (c *Client) UpdateDomainMetadata(ctx context.Context, params *UpdateDomainMetadataInput, optFns ...func(*Options)) (*UpdateDomainMetadataOutput, error) {
	if params == nil {
		params = &UpdateDomainMetadataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDomainMetadata", params, optFns, addOperationUpdateDomainMetadataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDomainMetadataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDomainMetadataInput struct {

	// The name of the domain.
	//
	// This member is required.
	DomainName *string

	// The ARN of the fleet.
	//
	// This member is required.
	FleetArn *string

	// The name to display.
	DisplayName *string
}

type UpdateDomainMetadataOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateDomainMetadataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateDomainMetadata{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateDomainMetadata{}, middleware.After)
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
	addOpUpdateDomainMetadataValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDomainMetadata(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateDomainMetadata(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "worklink",
		OperationName: "UpdateDomainMetadata",
	}
}
