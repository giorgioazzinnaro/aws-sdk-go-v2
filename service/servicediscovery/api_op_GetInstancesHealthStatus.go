// Code generated by smithy-go-codegen DO NOT EDIT.

package servicediscovery

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicediscovery/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the current health status (Healthy, Unhealthy, or Unknown) of one or more
// instances that are associated with a specified service. There is a brief delay
// between when you register an instance and when the health status for the
// instance is available.
func (c *Client) GetInstancesHealthStatus(ctx context.Context, params *GetInstancesHealthStatusInput, optFns ...func(*Options)) (*GetInstancesHealthStatusOutput, error) {
	if params == nil {
		params = &GetInstancesHealthStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetInstancesHealthStatus", params, optFns, addOperationGetInstancesHealthStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetInstancesHealthStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetInstancesHealthStatusInput struct {

	// The ID of the service that the instance is associated with.
	//
	// This member is required.
	ServiceId *string

	// An array that contains the IDs of all the instances that you want to get the
	// health status for. If you omit Instances, AWS Cloud Map returns the health
	// status for all the instances that are associated with the specified service. To
	// get the IDs for the instances that you've registered by using a specified
	// service, submit a ListInstances
	// (https://docs.aws.amazon.com/cloud-map/latest/api/API_ListInstances.html)
	// request.
	Instances []*string

	// The maximum number of instances that you want AWS Cloud Map to return in the
	// response to a GetInstancesHealthStatus request. If you don't specify a value for
	// MaxResults, AWS Cloud Map returns up to 100 instances.
	MaxResults *int32

	// For the first GetInstancesHealthStatus request, omit this value. If more than
	// MaxResults instances match the specified criteria, you can submit another
	// GetInstancesHealthStatus request to get the next group of results. Specify the
	// value of NextToken from the previous response in the next request.
	NextToken *string
}

type GetInstancesHealthStatusOutput struct {

	// If more than MaxResults instances match the specified criteria, you can submit
	// another GetInstancesHealthStatus request to get the next group of results.
	// Specify the value of NextToken from the previous response in the next request.
	NextToken *string

	// A complex type that contains the IDs and the health status of the instances that
	// you specified in the GetInstancesHealthStatus request.
	Status map[string]types.HealthStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetInstancesHealthStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpGetInstancesHealthStatus{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpGetInstancesHealthStatus{})
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
	if err := addOpGetInstancesHealthStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opGetInstancesHealthStatus(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opGetInstancesHealthStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicediscovery",
		OperationName: "GetInstancesHealthStatus",
	}
}
