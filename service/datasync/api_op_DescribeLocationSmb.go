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

// Returns metadata, such as the path and user information about an SMB location.
func (c *Client) DescribeLocationSmb(ctx context.Context, params *DescribeLocationSmbInput, optFns ...func(*Options)) (*DescribeLocationSmbOutput, error) {
	if params == nil {
		params = &DescribeLocationSmbInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeLocationSmb", params, optFns, addOperationDescribeLocationSmbMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeLocationSmbOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// DescribeLocationSmbRequest
type DescribeLocationSmbInput struct {

	// The Amazon Resource Name (ARN) of the SMB location to describe.
	//
	// This member is required.
	LocationArn *string
}

// DescribeLocationSmbResponse
type DescribeLocationSmbOutput struct {

	// The Amazon Resource Name (ARN) of the source SMB file system location that is
	// created.
	AgentArns []*string

	// The time that the SMB location was created.
	CreationTime *time.Time

	// The name of the Windows domain that the SMB server belongs to.
	Domain *string

	// The Amazon Resource Name (ARN) of the SMB location that was described.
	LocationArn *string

	// The URL of the source SBM location that was described.
	LocationUri *string

	// The mount options that are available for DataSync to use to access an SMB
	// location.
	MountOptions *types.SmbMountOptions

	// The user who can mount the share, has the permissions to access files and
	// folders in the SMB share.
	User *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeLocationSmbMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpDescribeLocationSmb{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpDescribeLocationSmb{})
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
	if err := addOpDescribeLocationSmbValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDescribeLocationSmb(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDescribeLocationSmb(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datasync",
		OperationName: "DescribeLocationSmb",
	}
}
