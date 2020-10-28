// Code generated by smithy-go-codegen DO NOT EDIT.

package schemas

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/schemas/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// List the discoverers.
func (c *Client) ListDiscoverers(ctx context.Context, params *ListDiscoverersInput, optFns ...func(*Options)) (*ListDiscoverersOutput, error) {
	if params == nil {
		params = &ListDiscoverersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDiscoverers", params, optFns, addOperationListDiscoverersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDiscoverersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDiscoverersInput struct {

	// Specifying this limits the results to only those discoverer IDs that start with
	// the specified prefix.
	DiscovererIdPrefix *string

	Limit *int32

	// The token that specifies the next page of results to return. To request the
	// first page, leave NextToken empty. The token will expire in 24 hours, and cannot
	// be shared with other accounts.
	NextToken *string

	// Specifying this limits the results to only those ARNs that start with the
	// specified prefix.
	SourceArnPrefix *string
}

type ListDiscoverersOutput struct {

	// An array of DiscovererSummary information.
	Discoverers []*types.DiscovererSummary

	// The token that specifies the next page of results to return. To request the
	// first page, leave NextToken empty. The token will expire in 24 hours, and cannot
	// be shared with other accounts.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListDiscoverersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpListDiscoverers{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpListDiscoverers{})
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
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opListDiscoverers(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opListDiscoverers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "schemas",
		OperationName: "ListDiscoverers",
	}
}
