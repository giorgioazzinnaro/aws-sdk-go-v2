// Code generated by smithy-go-codegen DO NOT EDIT.

package costexplorer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a cost anomaly monitor.
func (c *Client) DeleteAnomalyMonitor(ctx context.Context, params *DeleteAnomalyMonitorInput, optFns ...func(*Options)) (*DeleteAnomalyMonitorOutput, error) {
	if params == nil {
		params = &DeleteAnomalyMonitorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteAnomalyMonitor", params, optFns, addOperationDeleteAnomalyMonitorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteAnomalyMonitorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteAnomalyMonitorInput struct {

	// The unique identifier of the cost anomaly monitor that you want to delete.
	//
	// This member is required.
	MonitorArn *string
}

type DeleteAnomalyMonitorOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteAnomalyMonitorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteAnomalyMonitor{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteAnomalyMonitor{}, middleware.After)
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
	addOpDeleteAnomalyMonitorValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteAnomalyMonitor(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteAnomalyMonitor(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ce",
		OperationName: "DeleteAnomalyMonitor",
	}
}
