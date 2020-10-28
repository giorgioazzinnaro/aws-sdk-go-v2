// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the specified disk snapshot. When you make periodic snapshots of a disk,
// the snapshots are incremental, and only the blocks on the device that have
// changed since your last snapshot are saved in the new snapshot. When you delete
// a snapshot, only the data not needed for any other snapshot is removed. So
// regardless of which prior snapshots have been deleted, all active snapshots will
// have access to all the information needed to restore the disk. The delete disk
// snapshot operation supports tag-based access control via resource tags applied
// to the resource identified by disk snapshot name. For more information, see the
// Lightsail Dev Guide
// (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-controlling-access-using-tags).
func (c *Client) DeleteDiskSnapshot(ctx context.Context, params *DeleteDiskSnapshotInput, optFns ...func(*Options)) (*DeleteDiskSnapshotOutput, error) {
	if params == nil {
		params = &DeleteDiskSnapshotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteDiskSnapshot", params, optFns, addOperationDeleteDiskSnapshotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteDiskSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteDiskSnapshotInput struct {

	// The name of the disk snapshot you want to delete (e.g., my-disk-snapshot).
	//
	// This member is required.
	DiskSnapshotName *string
}

type DeleteDiskSnapshotOutput struct {

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request.
	Operations []*types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteDiskSnapshotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpDeleteDiskSnapshot{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpDeleteDiskSnapshot{})
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
	if err := addOpDeleteDiskSnapshotValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDeleteDiskSnapshot(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDeleteDiskSnapshot(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "DeleteDiskSnapshot",
	}
}
