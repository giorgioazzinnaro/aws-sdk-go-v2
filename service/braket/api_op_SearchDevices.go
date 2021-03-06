// Code generated by smithy-go-codegen DO NOT EDIT.

package braket

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/braket/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Searches for devices using the specified filters.
func (c *Client) SearchDevices(ctx context.Context, params *SearchDevicesInput, optFns ...func(*Options)) (*SearchDevicesOutput, error) {
	if params == nil {
		params = &SearchDevicesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchDevices", params, optFns, addOperationSearchDevicesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchDevicesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchDevicesInput struct {

	// The filter values to use to search for a device.
	//
	// This member is required.
	Filters []*types.SearchDevicesFilter

	// The maximum number of results to return in the response.
	MaxResults *int32

	// A token used for pagination of results returned in the response. Use the token
	// returned from the previous request continue results where the previous request
	// ended.
	NextToken *string
}

type SearchDevicesOutput struct {

	// An array of DeviceSummary objects for devices that match the specified filter
	// values.
	//
	// This member is required.
	Devices []*types.DeviceSummary

	// A token used for pagination of results, or null if there are no additional
	// results. Use the token value in a subsequent request to continue results where
	// the previous request ended.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationSearchDevicesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSearchDevices{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSearchDevices{}, middleware.After)
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
	addOpSearchDevicesValidationMiddleware(stack)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}
