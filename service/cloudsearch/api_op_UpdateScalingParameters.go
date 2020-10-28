// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudsearch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudsearch/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Configures scaling parameters for a domain. A domain's scaling parameters
// specify the desired search instance type and replication count. Amazon
// CloudSearch will still automatically scale your domain based on the volume of
// data and traffic, but not below the desired instance type and replication count.
// If the Multi-AZ option is enabled, these values control the resources used per
// Availability Zone. For more information, see Configuring Scaling Options
// (http://docs.aws.amazon.com/cloudsearch/latest/developerguide/configuring-scaling-options.html)
// in the Amazon CloudSearch Developer Guide.
func (c *Client) UpdateScalingParameters(ctx context.Context, params *UpdateScalingParametersInput, optFns ...func(*Options)) (*UpdateScalingParametersOutput, error) {
	if params == nil {
		params = &UpdateScalingParametersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateScalingParameters", params, optFns, addOperationUpdateScalingParametersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateScalingParametersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the UpdateScalingParameters operation. Specifies
// the name of the domain you want to update and the scaling parameters you want to
// configure.
type UpdateScalingParametersInput struct {

	// A string that represents the name of a domain. Domain names are unique across
	// the domains owned by an account within an AWS region. Domain names start with a
	// letter or number and can contain the following characters: a-z (lowercase), 0-9,
	// and - (hyphen).
	//
	// This member is required.
	DomainName *string

	// The desired instance type and desired number of replicas of each index
	// partition.
	//
	// This member is required.
	ScalingParameters *types.ScalingParameters
}

// The result of a UpdateScalingParameters request. Contains the status of the
// newly-configured scaling parameters.
type UpdateScalingParametersOutput struct {

	// The status and configuration of a search domain's scaling parameters.
	//
	// This member is required.
	ScalingParameters *types.ScalingParametersStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateScalingParametersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsquery_serializeOpUpdateScalingParameters{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsquery_deserializeOpUpdateScalingParameters{})
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
	if err := addOpUpdateScalingParametersValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opUpdateScalingParameters(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opUpdateScalingParameters(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudsearch",
		OperationName: "UpdateScalingParameters",
	}
}
