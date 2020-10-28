// Code generated by smithy-go-codegen DO NOT EDIT.

package codedeploy

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deregisters an on-premises instance.
func (c *Client) DeregisterOnPremisesInstance(ctx context.Context, params *DeregisterOnPremisesInstanceInput, optFns ...func(*Options)) (*DeregisterOnPremisesInstanceOutput, error) {
	if params == nil {
		params = &DeregisterOnPremisesInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeregisterOnPremisesInstance", params, optFns, addOperationDeregisterOnPremisesInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeregisterOnPremisesInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a DeregisterOnPremisesInstance operation.
type DeregisterOnPremisesInstanceInput struct {

	// The name of the on-premises instance to deregister.
	//
	// This member is required.
	InstanceName *string
}

type DeregisterOnPremisesInstanceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeregisterOnPremisesInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpDeregisterOnPremisesInstance{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpDeregisterOnPremisesInstance{})
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
	if err := addOpDeregisterOnPremisesInstanceValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDeregisterOnPremisesInstance(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDeregisterOnPremisesInstance(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codedeploy",
		OperationName: "DeregisterOnPremisesInstance",
	}
}
