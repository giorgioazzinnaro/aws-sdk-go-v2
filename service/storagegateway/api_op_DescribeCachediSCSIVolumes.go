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

// Returns a description of the gateway volumes specified in the request. This
// operation is only supported in the cached volume gateway types. The list of
// gateway volumes in the request must be from one gateway. In the response, AWS
// Storage Gateway returns volume information sorted by volume Amazon Resource Name
// (ARN).
func (c *Client) DescribeCachediSCSIVolumes(ctx context.Context, params *DescribeCachediSCSIVolumesInput, optFns ...func(*Options)) (*DescribeCachediSCSIVolumesOutput, error) {
	if params == nil {
		params = &DescribeCachediSCSIVolumesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeCachediSCSIVolumes", params, optFns, addOperationDescribeCachediSCSIVolumesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeCachediSCSIVolumesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeCachediSCSIVolumesInput struct {

	// An array of strings where each string represents the Amazon Resource Name (ARN)
	// of a cached volume. All of the specified cached volumes must be from the same
	// gateway. Use ListVolumes to get volume ARNs for a gateway.
	//
	// This member is required.
	VolumeARNs []*string
}

// A JSON object containing the following fields:
type DescribeCachediSCSIVolumesOutput struct {

	// An array of objects where each object contains metadata about one cached volume.
	CachediSCSIVolumes []*types.CachediSCSIVolume

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeCachediSCSIVolumesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpDescribeCachediSCSIVolumes{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpDescribeCachediSCSIVolumes{})
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
	if err := addOpDescribeCachediSCSIVolumesValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDescribeCachediSCSIVolumes(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDescribeCachediSCSIVolumes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "DescribeCachediSCSIVolumes",
	}
}
