// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the role aliases registered in your account.
func (c *Client) ListRoleAliases(ctx context.Context, params *ListRoleAliasesInput, optFns ...func(*Options)) (*ListRoleAliasesOutput, error) {
	if params == nil {
		params = &ListRoleAliasesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRoleAliases", params, optFns, addOperationListRoleAliasesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRoleAliasesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRoleAliasesInput struct {

	// Return the list of role aliases in ascending alphabetical order.
	AscendingOrder *bool

	// A marker used to get the next set of results.
	Marker *string

	// The maximum number of results to return at one time.
	PageSize *int32
}

type ListRoleAliasesOutput struct {

	// A marker used to get the next set of results.
	NextMarker *string

	// The role aliases.
	RoleAliases []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListRoleAliasesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpListRoleAliases{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpListRoleAliases{})
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
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opListRoleAliases(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opListRoleAliases(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ListRoleAliases",
	}
}
