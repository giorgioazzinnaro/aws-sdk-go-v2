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

// Lists summary information about the instances that you registered by using a
// specified service.
func (c *Client) ListInstances(ctx context.Context, params *ListInstancesInput, optFns ...func(*Options)) (*ListInstancesOutput, error) {
	if params == nil {
		params = &ListInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListInstances", params, optFns, addOperationListInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListInstancesInput struct {

	// The ID of the service that you want to list instances for.
	//
	// This member is required.
	ServiceId *string

	// The maximum number of instances that you want AWS Cloud Map to return in the
	// response to a ListInstances request. If you don't specify a value for
	// MaxResults, AWS Cloud Map returns up to 100 instances.
	MaxResults *int32

	// For the first ListInstances request, omit this value. If more than MaxResults
	// instances match the specified criteria, you can submit another ListInstances
	// request to get the next group of results. Specify the value of NextToken from
	// the previous response in the next request.
	NextToken *string
}

type ListInstancesOutput struct {

	// Summary information about the instances that are associated with the specified
	// service.
	Instances []*types.InstanceSummary

	// If more than MaxResults instances match the specified criteria, you can submit
	// another ListInstances request to get the next group of results. Specify the
	// value of NextToken from the previous response in the next request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpListInstances{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpListInstances{})
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
	if err := addOpListInstancesValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opListInstances(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opListInstances(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicediscovery",
		OperationName: "ListInstances",
	}
}
