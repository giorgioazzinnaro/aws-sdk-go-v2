// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the specified EBS volumes or all of your EBS volumes. If you are
// describing a long list of volumes, we recommend that you paginate the output to
// make the list more manageable. The MaxResults parameter sets the maximum number
// of results returned in a single page. If the list of results exceeds your
// MaxResults value, then that number of results is returned along with a NextToken
// value that can be passed to a subsequent DescribeVolumes request to retrieve the
// remaining results. For more information about EBS volumes, see Amazon EBS
// Volumes (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumes.html) in
// the Amazon Elastic Compute Cloud User Guide.
func (c *Client) DescribeVolumes(ctx context.Context, params *DescribeVolumesInput, optFns ...func(*Options)) (*DescribeVolumesOutput, error) {
	if params == nil {
		params = &DescribeVolumesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeVolumes", params, optFns, addOperationDescribeVolumesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeVolumesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeVolumesInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The filters.
	//
	//     * attachment.attach-time - The time stamp when the attachment
	// initiated.
	//
	//     * attachment.delete-on-termination - Whether the volume is
	// deleted on instance termination.
	//
	//     * attachment.device - The device name
	// specified in the block device mapping (for example, /dev/sda1).
	//
	//     *
	// attachment.instance-id - The ID of the instance the volume is attached to.
	//
	//
	// * attachment.status - The attachment state (attaching | attached | detaching).
	//
	//
	// * availability-zone - The Availability Zone in which the volume was created.
	//
	//
	// * create-time - The time stamp when the volume was created.
	//
	//     * encrypted -
	// Indicates whether the volume is encrypted (true | false)
	//
	//     *
	// multi-attach-enabled - Indicates whether the volume is enabled for Multi-Attach
	// (true | false)
	//
	//     * fast-restored - Indicates whether the volume was created
	// from a snapshot that is enabled for fast snapshot restore (true | false).
	//
	//     *
	// size - The size of the volume, in GiB.
	//
	//     * snapshot-id - The snapshot from
	// which the volume was created.
	//
	//     * status - The state of the volume (creating
	// | available | in-use | deleting | deleted | error).
	//
	//     * tag: - The key/value
	// combination of a tag assigned to the resource. Use the tag key in the filter
	// name and the tag value as the filter value. For example, to find all resources
	// that have a tag with the key Owner and the value TeamA, specify tag:Owner for
	// the filter name and TeamA for the filter value.
	//
	//     * tag-key - The key of a
	// tag assigned to the resource. Use this filter to find all resources assigned a
	// tag with a specific key, regardless of the tag value.
	//
	//     * volume-id - The
	// volume ID.
	//
	//     * volume-type - The Amazon EBS volume type. This can be gp2 for
	// General Purpose SSD, io1 or io2 for Provisioned IOPS SSD, st1 for Throughput
	// Optimized HDD, sc1 for Cold HDD, or standard for Magnetic volumes.
	Filters []*types.Filter

	// The maximum number of volume results returned by DescribeVolumes in paginated
	// output. When this parameter is used, DescribeVolumes only returns MaxResults
	// results in a single page along with a NextToken response element. The remaining
	// results of the initial request can be seen by sending another DescribeVolumes
	// request with the returned NextToken value. This value can be between 5 and 500;
	// if MaxResults is given a value larger than 500, only 500 results are returned.
	// If this parameter is not used, then DescribeVolumes returns all results. You
	// cannot specify this parameter and the volume IDs parameter in the same request.
	MaxResults *int32

	// The NextToken value returned from a previous paginated DescribeVolumes request
	// where MaxResults was used and the results exceeded the value of that parameter.
	// Pagination continues from the end of the previous results that returned the
	// NextToken value. This value is null when there are no more results to return.
	NextToken *string

	// The volume IDs.
	VolumeIds []*string
}

type DescribeVolumesOutput struct {

	// The NextToken value to include in a future DescribeVolumes request. When the
	// results of a DescribeVolumes request exceed MaxResults, this value can be used
	// to retrieve the next page of results. This value is null when there are no more
	// results to return.
	NextToken *string

	// Information about the volumes.
	Volumes []*types.Volume

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeVolumesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsEc2query_serializeOpDescribeVolumes{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsEc2query_deserializeOpDescribeVolumes{})
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
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDescribeVolumes(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDescribeVolumes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeVolumes",
	}
}
