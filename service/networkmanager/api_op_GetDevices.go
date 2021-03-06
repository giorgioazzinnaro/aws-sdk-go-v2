// Code generated by smithy-go-codegen DO NOT EDIT.

package networkmanager

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/networkmanager/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about one or more of your devices in a global network.
func (c *Client) GetDevices(ctx context.Context, params *GetDevicesInput, optFns ...func(*Options)) (*GetDevicesOutput, error) {
	if params == nil {
		params = &GetDevicesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDevices", params, optFns, addOperationGetDevicesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDevicesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDevicesInput struct {

	// The ID of the global network.
	//
	// This member is required.
	GlobalNetworkId *string

	// One or more device IDs. The maximum is 10.
	DeviceIds []*string

	// The maximum number of results to return.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	// The ID of the site.
	SiteId *string
}

type GetDevicesOutput struct {

	// The devices.
	Devices []*types.Device

	// The token for the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetDevicesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetDevices{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDevices{}, middleware.After)
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
	addOpGetDevicesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetDevices(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetDevices(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "networkmanager",
		OperationName: "GetDevices",
	}
}
