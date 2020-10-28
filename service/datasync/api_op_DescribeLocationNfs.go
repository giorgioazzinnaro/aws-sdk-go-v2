// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns metadata, such as the path information, about an NFS location.
func (c *Client) DescribeLocationNfs(ctx context.Context, params *DescribeLocationNfsInput, optFns ...func(*Options)) (*DescribeLocationNfsOutput, error) {
	if params == nil {
		params = &DescribeLocationNfsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeLocationNfs", params, optFns, addOperationDescribeLocationNfsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeLocationNfsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// DescribeLocationNfsRequest
type DescribeLocationNfsInput struct {

	// The Amazon Resource Name (ARN) of the NFS location to describe.
	//
	// This member is required.
	LocationArn *string
}

// DescribeLocationNfsResponse
type DescribeLocationNfsOutput struct {

	// The time that the NFS location was created.
	CreationTime *time.Time

	// The Amazon Resource Name (ARN) of the NFS location that was described.
	LocationArn *string

	// The URL of the source NFS location that was described.
	LocationUri *string

	// The NFS mount options that DataSync used to mount your NFS share.
	MountOptions *types.NfsMountOptions

	// A list of Amazon Resource Names (ARNs) of agents to use for a Network File
	// System (NFS) location.
	OnPremConfig *types.OnPremConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeLocationNfsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpDescribeLocationNfs{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpDescribeLocationNfs{})
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
	if err := addOpDescribeLocationNfsValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDescribeLocationNfs(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDescribeLocationNfs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datasync",
		OperationName: "DescribeLocationNfs",
	}
}
