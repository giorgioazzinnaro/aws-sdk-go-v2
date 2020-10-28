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

// Update the properties of a repository.
func (c *Client) UpdateRepository(ctx context.Context, params *UpdateRepositoryInput, optFns ...func(*Options)) (*UpdateRepositoryOutput, error) {
	if params == nil {
		params = &UpdateRepositoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateRepository", params, optFns, addOperationUpdateRepositoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateRepositoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateRepositoryInput struct {

	// The name of the domain associated with the repository to update.
	//
	// This member is required.
	Domain *string

	// The name of the repository to update.
	//
	// This member is required.
	Repository *string

	// An updated repository description.
	Description *string

	// The 12-digit account number of the AWS account that owns the domain. It does not
	// include dashes or spaces.
	DomainOwner *string

	// A list of upstream repositories to associate with the repository. The order of
	// the upstream repositories in the list determines their priority order when AWS
	// CodeArtifact looks for a requested package version. For more information, see
	// Working with upstream repositories
	// (https://docs.aws.amazon.com/codeartifact/latest/ug/repos-upstream.html).
	Upstreams []*types.UpstreamRepository
}

type UpdateRepositoryOutput struct {

	// The updated repository.
	Repository *types.RepositoryDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateRepositoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpUpdateRepository{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpUpdateRepository{})
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
	if err := addOpUpdateRepositoryValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opUpdateRepository(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opUpdateRepository(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeartifact",
		OperationName: "UpdateRepository",
	}
}
