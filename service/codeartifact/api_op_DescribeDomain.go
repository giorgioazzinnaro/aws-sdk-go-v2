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

// Returns a DomainDescription
// (https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_DomainDescription.html)
// object that contains information about the requested domain.
func (c *Client) DescribeDomain(ctx context.Context, params *DescribeDomainInput, optFns ...func(*Options)) (*DescribeDomainOutput, error) {
	if params == nil {
		params = &DescribeDomainInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDomain", params, optFns, addOperationDescribeDomainMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDomainOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDomainInput struct {

	// A string that specifies the name of the requested domain.
	//
	// This member is required.
	Domain *string

	// The 12-digit account number of the AWS account that owns the domain. It does not
	// include dashes or spaces.
	DomainOwner *string
}

type DescribeDomainOutput struct {

	// Information about a domain. A domain is a container for repositories. When you
	// create a domain, it is empty until you add one or more repositories.
	Domain *types.DomainDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeDomainMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpDescribeDomain{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpDescribeDomain{})
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
	if err := addOpDescribeDomainValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDescribeDomain(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDescribeDomain(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeartifact",
		OperationName: "DescribeDomain",
	}
}
