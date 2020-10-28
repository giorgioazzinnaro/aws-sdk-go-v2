// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/directoryservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Shares a specified directory (DirectoryId) in your AWS account (directory owner)
// with another AWS account (directory consumer). With this operation you can use
// your directory from any AWS account and from any Amazon VPC within an AWS
// Region. When you share your AWS Managed Microsoft AD directory, AWS Directory
// Service creates a shared directory in the directory consumer account. This
// shared directory contains the metadata to provide access to the directory within
// the directory owner account. The shared directory is visible in all VPCs in the
// directory consumer account. The ShareMethod parameter determines whether the
// specified directory can be shared between AWS accounts inside the same AWS
// organization (ORGANIZATIONS). It also determines whether you can share the
// directory with any other AWS account either inside or outside of the
// organization (HANDSHAKE). The ShareNotes parameter is only used when HANDSHAKE
// is called, which sends a directory sharing request to the directory consumer.
func (c *Client) ShareDirectory(ctx context.Context, params *ShareDirectoryInput, optFns ...func(*Options)) (*ShareDirectoryOutput, error) {
	if params == nil {
		params = &ShareDirectoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ShareDirectory", params, optFns, addOperationShareDirectoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ShareDirectoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ShareDirectoryInput struct {

	// Identifier of the AWS Managed Microsoft AD directory that you want to share with
	// other AWS accounts.
	//
	// This member is required.
	DirectoryId *string

	// The method used when sharing a directory to determine whether the directory
	// should be shared within your AWS organization (ORGANIZATIONS) or with any AWS
	// account by sending a directory sharing request (HANDSHAKE).
	//
	// This member is required.
	ShareMethod types.ShareMethod

	// Identifier for the directory consumer account with whom the directory is to be
	// shared.
	//
	// This member is required.
	ShareTarget *types.ShareTarget

	// A directory share request that is sent by the directory owner to the directory
	// consumer. The request includes a typed message to help the directory consumer
	// administrator determine whether to approve or reject the share invitation.
	ShareNotes *string
}

type ShareDirectoryOutput struct {

	// Identifier of the directory that is stored in the directory consumer account
	// that is shared from the specified directory (DirectoryId).
	SharedDirectoryId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationShareDirectoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpShareDirectory{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpShareDirectory{})
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
	if err := addOpShareDirectoryValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opShareDirectory(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opShareDirectory(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "ShareDirectory",
	}
}
