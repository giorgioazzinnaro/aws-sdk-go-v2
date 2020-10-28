// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Detaches the specified principal from the specified thing. A principal can be
// X.509 certificates, IAM users, groups, and roles, Amazon Cognito identities or
// federated identities. This call is asynchronous. It might take several seconds
// for the detachment to propagate.
func (c *Client) DetachThingPrincipal(ctx context.Context, params *DetachThingPrincipalInput, optFns ...func(*Options)) (*DetachThingPrincipalOutput, error) {
	if params == nil {
		params = &DetachThingPrincipalInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DetachThingPrincipal", params, optFns, addOperationDetachThingPrincipalMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DetachThingPrincipalOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the DetachThingPrincipal operation.
type DetachThingPrincipalInput struct {

	// If the principal is a certificate, this value must be ARN of the certificate. If
	// the principal is an Amazon Cognito identity, this value must be the ID of the
	// Amazon Cognito identity.
	//
	// This member is required.
	Principal *string

	// The name of the thing.
	//
	// This member is required.
	ThingName *string
}

// The output from the DetachThingPrincipal operation.
type DetachThingPrincipalOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDetachThingPrincipalMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpDetachThingPrincipal{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpDetachThingPrincipal{})
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
	if err := addOpDetachThingPrincipalValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDetachThingPrincipal(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDetachThingPrincipal(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "DetachThingPrincipal",
	}
}
