// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancing

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the specified listeners from the specified load balancer.
func (c *Client) DeleteLoadBalancerListeners(ctx context.Context, params *DeleteLoadBalancerListenersInput, optFns ...func(*Options)) (*DeleteLoadBalancerListenersOutput, error) {
	if params == nil {
		params = &DeleteLoadBalancerListenersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteLoadBalancerListeners", params, optFns, addOperationDeleteLoadBalancerListenersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteLoadBalancerListenersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DeleteLoadBalancerListeners.
type DeleteLoadBalancerListenersInput struct {

	// The name of the load balancer.
	//
	// This member is required.
	LoadBalancerName *string

	// The client port numbers of the listeners.
	//
	// This member is required.
	LoadBalancerPorts []*int32
}

// Contains the output of DeleteLoadBalancerListeners.
type DeleteLoadBalancerListenersOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteLoadBalancerListenersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsquery_serializeOpDeleteLoadBalancerListeners{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsquery_deserializeOpDeleteLoadBalancerListeners{})
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
	if err := addOpDeleteLoadBalancerListenersValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDeleteLoadBalancerListeners(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDeleteLoadBalancerListeners(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticloadbalancing",
		OperationName: "DeleteLoadBalancerListeners",
	}
}
