// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticinference

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticinference/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the locations in which a given accelerator type or set of types is
// present in a given region.
func (c *Client) DescribeAcceleratorOfferings(ctx context.Context, params *DescribeAcceleratorOfferingsInput, optFns ...func(*Options)) (*DescribeAcceleratorOfferingsOutput, error) {
	if params == nil {
		params = &DescribeAcceleratorOfferingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAcceleratorOfferings", params, optFns, addOperationDescribeAcceleratorOfferingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAcceleratorOfferingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAcceleratorOfferingsInput struct {

	// The location type that you want to describe accelerator type offerings for. It
	// can assume the following values: region: will return the accelerator type
	// offering at the regional level. availability-zone: will return the accelerator
	// type offering at the availability zone level. availability-zone-id: will return
	// the accelerator type offering at the availability zone level returning the
	// availability zone id.
	//
	// This member is required.
	LocationType types.LocationType

	// The list of accelerator types to describe.
	AcceleratorTypes []*string
}

type DescribeAcceleratorOfferingsOutput struct {

	// The list of accelerator type offerings for a specific location.
	AcceleratorTypeOfferings []*types.AcceleratorTypeOffering

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeAcceleratorOfferingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpDescribeAcceleratorOfferings{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpDescribeAcceleratorOfferings{})
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
	if err := addOpDescribeAcceleratorOfferingsValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDescribeAcceleratorOfferings(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDescribeAcceleratorOfferings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elastic-inference",
		OperationName: "DescribeAcceleratorOfferings",
	}
}
