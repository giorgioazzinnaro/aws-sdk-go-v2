// Code generated by smithy-go-codegen DO NOT EDIT.

package workdocs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workdocs/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the current user's special folders; the RootFolder and the RecycleBin.
// RootFolder is the root of user's files and folders and RecycleBin is the root of
// recycled items. This is not a valid action for SigV4 (administrative API)
// clients. This action requires an authentication token. To get an authentication
// token, register an application with Amazon WorkDocs. For more information, see
// Authentication and Access Control for User Applications
// (https://docs.aws.amazon.com/workdocs/latest/developerguide/wd-auth-user.html)
// in the Amazon WorkDocs Developer Guide.
func (c *Client) DescribeRootFolders(ctx context.Context, params *DescribeRootFoldersInput, optFns ...func(*Options)) (*DescribeRootFoldersOutput, error) {
	if params == nil {
		params = &DescribeRootFoldersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeRootFolders", params, optFns, addOperationDescribeRootFoldersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeRootFoldersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeRootFoldersInput struct {

	// Amazon WorkDocs authentication token.
	//
	// This member is required.
	AuthenticationToken *string

	// The maximum number of items to return.
	Limit *int32

	// The marker for the next set of results. (You received this marker from a
	// previous call.)
	Marker *string
}

type DescribeRootFoldersOutput struct {

	// The user's special folders.
	Folders []*types.FolderMetadata

	// The marker for the next set of results.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeRootFoldersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpDescribeRootFolders{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpDescribeRootFolders{})
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
	if err := addOpDescribeRootFoldersValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDescribeRootFolders(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDescribeRootFolders(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workdocs",
		OperationName: "DescribeRootFolders",
	}
}
