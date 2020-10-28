// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the existing thing types.
func (c *Client) ListThingTypes(ctx context.Context, params *ListThingTypesInput, optFns ...func(*Options)) (*ListThingTypesOutput, error) {
	if params == nil {
		params = &ListThingTypesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListThingTypes", params, optFns, addOperationListThingTypesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListThingTypesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the ListThingTypes operation.
type ListThingTypesInput struct {

	// The maximum number of results to return in this operation.
	MaxResults *int32

	// The token to retrieve the next set of results.
	NextToken *string

	// The name of the thing type.
	ThingTypeName *string
}

// The output for the ListThingTypes operation.
type ListThingTypesOutput struct {

	// The token for the next set of results. Will not be returned if operation has
	// returned all results.
	NextToken *string

	// The thing types.
	ThingTypes []*types.ThingTypeDefinition

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListThingTypesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpListThingTypes{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpListThingTypes{})
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
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opListThingTypes(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opListThingTypes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ListThingTypes",
	}
}
