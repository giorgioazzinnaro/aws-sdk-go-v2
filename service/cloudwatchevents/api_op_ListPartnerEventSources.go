// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchevents

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchevents/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// An SaaS partner can use this operation to list all the partner event source
// names that they have created. This operation is not used by AWS customers.
func (c *Client) ListPartnerEventSources(ctx context.Context, params *ListPartnerEventSourcesInput, optFns ...func(*Options)) (*ListPartnerEventSourcesOutput, error) {
	if params == nil {
		params = &ListPartnerEventSourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPartnerEventSources", params, optFns, addOperationListPartnerEventSourcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPartnerEventSourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPartnerEventSourcesInput struct {

	// If you specify this, the results are limited to only those partner event sources
	// that start with the string you specify.
	//
	// This member is required.
	NamePrefix *string

	// pecifying this limits the number of results returned by this operation. The
	// operation also returns a NextToken which you can use in a subsequent operation
	// to retrieve the next set of results.
	Limit *int32

	// The token returned by a previous call to this operation. Specifying this
	// retrieves the next set of results.
	NextToken *string
}

type ListPartnerEventSourcesOutput struct {

	// A token you can use in a subsequent operation to retrieve the next set of
	// results.
	NextToken *string

	// The list of partner event sources returned by the operation.
	PartnerEventSources []*types.PartnerEventSource

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListPartnerEventSourcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpListPartnerEventSources{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpListPartnerEventSources{})
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
	if err := addOpListPartnerEventSourcesValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opListPartnerEventSources(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opListPartnerEventSources(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "events",
		OperationName: "ListPartnerEventSources",
	}
}
