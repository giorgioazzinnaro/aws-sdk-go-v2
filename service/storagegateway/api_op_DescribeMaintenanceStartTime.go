// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns your gateway's weekly maintenance start time including the day and time
// of the week. Note that values are in terms of the gateway's time zone.
func (c *Client) DescribeMaintenanceStartTime(ctx context.Context, params *DescribeMaintenanceStartTimeInput, optFns ...func(*Options)) (*DescribeMaintenanceStartTimeOutput, error) {
	if params == nil {
		params = &DescribeMaintenanceStartTimeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeMaintenanceStartTime", params, optFns, addOperationDescribeMaintenanceStartTimeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeMaintenanceStartTimeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A JSON object containing the Amazon Resource Name (ARN) of the gateway.
type DescribeMaintenanceStartTimeInput struct {

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation to
	// return a list of gateways for your account and AWS Region.
	//
	// This member is required.
	GatewayARN *string
}

// A JSON object containing the following fields:
//
//     *
// DescribeMaintenanceStartTimeOutput$DayOfMonth
//
//     *
// DescribeMaintenanceStartTimeOutput$DayOfWeek
//
//     *
// DescribeMaintenanceStartTimeOutput$HourOfDay
//
//     *
// DescribeMaintenanceStartTimeOutput$MinuteOfHour
//
//     *
// DescribeMaintenanceStartTimeOutput$Timezone
type DescribeMaintenanceStartTimeOutput struct {

	// The day of the month component of the maintenance start time represented as an
	// ordinal number from 1 to 28, where 1 represents the first day of the month and
	// 28 represents the last day of the month.
	DayOfMonth *int32

	// An ordinal number between 0 and 6 that represents the day of the week, where 0
	// represents Sunday and 6 represents Saturday. The day of week is in the time zone
	// of the gateway.
	DayOfWeek *int32

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation to
	// return a list of gateways for your account and AWS Region.
	GatewayARN *string

	// The hour component of the maintenance start time represented as hh, where hh is
	// the hour (0 to 23). The hour of the day is in the time zone of the gateway.
	HourOfDay *int32

	// The minute component of the maintenance start time represented as mm, where mm
	// is the minute (0 to 59). The minute of the hour is in the time zone of the
	// gateway.
	MinuteOfHour *int32

	// A value that indicates the time zone that is set for the gateway. The start time
	// and day of week specified should be in the time zone of the gateway.
	Timezone *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeMaintenanceStartTimeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpDescribeMaintenanceStartTime{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpDescribeMaintenanceStartTime{})
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
	if err := addOpDescribeMaintenanceStartTimeValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDescribeMaintenanceStartTime(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDescribeMaintenanceStartTime(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "DescribeMaintenanceStartTime",
	}
}
