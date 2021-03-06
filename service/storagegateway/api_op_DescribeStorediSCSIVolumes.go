// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the description of the gateway volumes specified in the request. The
// list of gateway volumes in the request must be from one gateway. In the
// response, AWS Storage Gateway returns volume information sorted by volume ARNs.
// This operation is only supported in stored volume gateway type.
func (c *Client) DescribeStorediSCSIVolumes(ctx context.Context, params *DescribeStorediSCSIVolumesInput, optFns ...func(*Options)) (*DescribeStorediSCSIVolumesOutput, error) {
	if params == nil {
		params = &DescribeStorediSCSIVolumesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeStorediSCSIVolumes", params, optFns, addOperationDescribeStorediSCSIVolumesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeStorediSCSIVolumesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A JSON object containing a list of DescribeStorediSCSIVolumesInput$VolumeARNs.
type DescribeStorediSCSIVolumesInput struct {

	// An array of strings where each string represents the Amazon Resource Name (ARN)
	// of a stored volume. All of the specified stored volumes must be from the same
	// gateway. Use ListVolumes to get volume ARNs for a gateway.
	//
	// This member is required.
	VolumeARNs []*string
}

type DescribeStorediSCSIVolumesOutput struct {

	// Describes a single unit of output from DescribeStorediSCSIVolumes. The following
	// fields are returned:
	//
	// * ChapEnabled: Indicates whether mutual CHAP is enabled
	// for the iSCSI target.
	//
	// * LunNumber: The logical disk number.
	//
	// *
	// NetworkInterfaceId: The network interface ID of the stored volume that initiator
	// use to map the stored volume as an iSCSI target.
	//
	// * NetworkInterfacePort: The
	// port used to communicate with iSCSI targets.
	//
	// * PreservedExistingData: Indicates
	// when the stored volume was created, existing data on the underlying local disk
	// was preserved.
	//
	// * SourceSnapshotId: If the stored volume was created from a
	// snapshot, this field contains the snapshot ID used, e.g. snap-1122aabb.
	// Otherwise, this field is not included.
	//
	// * StorediSCSIVolumes: An array of
	// StorediSCSIVolume objects where each object contains metadata about one stored
	// volume.
	//
	// * TargetARN: The Amazon Resource Name (ARN) of the volume target.
	//
	// *
	// VolumeARN: The Amazon Resource Name (ARN) of the stored volume.
	//
	// * VolumeDiskId:
	// The disk ID of the local disk that was specified in the CreateStorediSCSIVolume
	// operation.
	//
	// * VolumeId: The unique identifier of the storage volume, e.g.
	// vol-1122AABB.
	//
	// * VolumeiSCSIAttributes: An VolumeiSCSIAttributes object that
	// represents a collection of iSCSI attributes for one stored volume.
	//
	// *
	// VolumeProgress: Represents the percentage complete if the volume is restoring or
	// bootstrapping that represents the percent of data transferred. This field does
	// not appear in the response if the stored volume is not restoring or
	// bootstrapping.
	//
	// * VolumeSizeInBytes: The size of the volume in bytes.
	//
	// *
	// VolumeStatus: One of the VolumeStatus values that indicates the state of the
	// volume.
	//
	// * VolumeType: One of the enumeration values describing the type of the
	// volume. Currently, only STORED volumes are supported.
	StorediSCSIVolumes []*types.StorediSCSIVolume

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeStorediSCSIVolumesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeStorediSCSIVolumes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeStorediSCSIVolumes{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpDescribeStorediSCSIVolumesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeStorediSCSIVolumes(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeStorediSCSIVolumes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "DescribeStorediSCSIVolumes",
	}
}
