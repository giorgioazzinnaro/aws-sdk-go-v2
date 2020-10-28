// Code generated by smithy-go-codegen DO NOT EDIT.

package codeartifact

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Sets a resource policy on a domain that specifies permissions to access it.
func (c *Client) PutDomainPermissionsPolicy(ctx context.Context, params *PutDomainPermissionsPolicyInput, optFns ...func(*Options)) (*PutDomainPermissionsPolicyOutput, error) {
	if params == nil {
		params = &PutDomainPermissionsPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutDomainPermissionsPolicy", params, optFns, addOperationPutDomainPermissionsPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutDomainPermissionsPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutDomainPermissionsPolicyInput struct {

	// The name of the domain on which to set the resource policy.
	//
	// This member is required.
	Domain *string

	// A valid displayable JSON Aspen policy string to be set as the access control
	// resource policy on the provided domain.
	//
	// This member is required.
	PolicyDocument *string

	// The 12-digit account number of the AWS account that owns the domain. It does not
	// include dashes or spaces.
	DomainOwner *string

	// The current revision of the resource policy to be set. This revision is used for
	// optimistic locking, which prevents others from overwriting your changes to the
	// domain's resource policy.
	PolicyRevision *string
}

type PutDomainPermissionsPolicyOutput struct {

	// The resource policy that was set after processing the request.
	Policy *types.ResourcePolicy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutDomainPermissionsPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpPutDomainPermissionsPolicy{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpPutDomainPermissionsPolicy{})
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
	if err := addOpPutDomainPermissionsPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opPutDomainPermissionsPolicy(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opPutDomainPermissionsPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeartifact",
		OperationName: "PutDomainPermissionsPolicy",
	}
}
