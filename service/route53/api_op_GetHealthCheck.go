// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about a specified health check.
func (c *Client) GetHealthCheck(ctx context.Context, params *GetHealthCheckInput, optFns ...func(*Options)) (*GetHealthCheckOutput, error) {
	if params == nil {
		params = &GetHealthCheckInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetHealthCheck", params, optFns, addOperationGetHealthCheckMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetHealthCheckOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to get information about a specified health check.
type GetHealthCheckInput struct {

	// The identifier that Amazon Route 53 assigned to the health check when you
	// created it. When you add or update a resource record set, you use this value to
	// specify which health check to use. The value can be up to 64 characters long.
	//
	// This member is required.
	HealthCheckId *string
}

// A complex type that contains the response to a GetHealthCheck request.
type GetHealthCheckOutput struct {

	// A complex type that contains information about one health check that is
	// associated with the current AWS account.
	//
	// This member is required.
	HealthCheck *types.HealthCheck

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetHealthCheckMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestxml_serializeOpGetHealthCheck{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestxml_deserializeOpGetHealthCheck{})
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
	if err := addOpGetHealthCheckValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opGetHealthCheck(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opGetHealthCheck(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "GetHealthCheck",
	}
}
