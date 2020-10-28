// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Use this operation to list all private and vendor workforces in an AWS Region.
// Note that you can only have one private workforce per AWS Region.
func (c *Client) ListWorkforces(ctx context.Context, params *ListWorkforcesInput, optFns ...func(*Options)) (*ListWorkforcesOutput, error) {
	if params == nil {
		params = &ListWorkforcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListWorkforces", params, optFns, addOperationListWorkforcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListWorkforcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListWorkforcesInput struct {

	// The maximum number of workforces returned in the response.
	MaxResults *int32

	// A filter you can use to search for workforces using part of the workforce name.
	NameContains *string

	// A token to resume pagination.
	NextToken *string

	// Sort workforces using the workforce name or creation date.
	SortBy types.ListWorkforcesSortByOptions

	// Sort workforces in ascending or descending order.
	SortOrder types.SortOrder
}

type ListWorkforcesOutput struct {

	// A list containing information about your workforce.
	//
	// This member is required.
	Workforces []*types.Workforce

	// A token to resume pagination.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListWorkforcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpListWorkforces{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpListWorkforces{})
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
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opListWorkforces(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opListWorkforces(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListWorkforces",
	}
}
