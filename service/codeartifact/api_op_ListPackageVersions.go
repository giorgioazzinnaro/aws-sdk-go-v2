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

// Returns a list of PackageVersionSummary
// (https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_PackageVersionSummary.html)
// objects for package versions in a repository that match the request parameters.
func (c *Client) ListPackageVersions(ctx context.Context, params *ListPackageVersionsInput, optFns ...func(*Options)) (*ListPackageVersionsOutput, error) {
	if params == nil {
		params = &ListPackageVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPackageVersions", params, optFns, addOperationListPackageVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPackageVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPackageVersionsInput struct {

	// The name of the domain that contains the repository that contains the returned
	// package versions.
	//
	// This member is required.
	Domain *string

	// The format of the returned packages. The valid package types are:
	//
	// * npm: A Node
	// Package Manager (npm) package.
	//
	// * pypi: A Python Package Index (PyPI)
	// package.
	//
	// * maven: A Maven package that contains compiled code in a
	// distributable format, such as a JAR file.
	//
	// This member is required.
	Format types.PackageFormat

	// The name of the package for which you want to return a list of package versions.
	//
	// This member is required.
	Package *string

	// The name of the repository that contains the package.
	//
	// This member is required.
	Repository *string

	// The 12-digit account number of the AWS account that owns the domain. It does not
	// include dashes or spaces.
	DomainOwner *string

	// The maximum number of results to return per page.
	MaxResults *int32

	// The namespace of the package. The package component that specifies its namespace
	// depends on its type. For example:
	//
	// * The namespace of a Maven package is its
	// groupId.
	//
	// * The namespace of an npm package is its scope.
	//
	// * A Python package
	// does not contain a corresponding component, so Python packages do not have a
	// namespace.
	Namespace *string

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	// How to sort the returned list of package versions.
	SortBy types.PackageVersionSortType

	// A string that specifies the status of the package versions to include in the
	// returned list. It can be one of the following:
	//
	// * Published
	//
	// * Unfinished
	//
	// *
	// Unlisted
	//
	// * Archived
	//
	// * Disposed
	Status types.PackageVersionStatus
}

type ListPackageVersionsOutput struct {

	// The default package version to display. This depends on the package format:
	//
	// *
	// For Maven and PyPI packages, it's the most recently published package
	// version.
	//
	// * For npm packages, it's the version referenced by the latest tag. If
	// the latest tag is not set, it's the most recently published package version.
	DefaultDisplayVersion *string

	// A format of the package. Valid package format values are:
	//
	// * npm
	//
	// * pypi
	//
	// *
	// maven
	Format types.PackageFormat

	// The namespace of the package. The package component that specifies its namespace
	// depends on its type. For example:
	//
	// * The namespace of a Maven package is its
	// groupId.
	//
	// * The namespace of an npm package is its scope.
	//
	// * A Python package
	// does not contain a corresponding component, so Python packages do not have a
	// namespace.
	Namespace *string

	// If there are additional results, this is the token for the next set of results.
	NextToken *string

	// The name of the package.
	Package *string

	// The returned list of PackageVersionSummary
	// (https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_PackageVersionSummary.html)
	// objects.
	Versions []*types.PackageVersionSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListPackageVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListPackageVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListPackageVersions{}, middleware.After)
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
	addOpListPackageVersionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListPackageVersions(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListPackageVersions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeartifact",
		OperationName: "ListPackageVersions",
	}
}
