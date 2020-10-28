// Code generated by smithy-go-codegen DO NOT EDIT.

package kafka

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kafka/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the cluster with the configuration that is specified in the request
// body.
func (c *Client) UpdateClusterConfiguration(ctx context.Context, params *UpdateClusterConfigurationInput, optFns ...func(*Options)) (*UpdateClusterConfigurationOutput, error) {
	if params == nil {
		params = &UpdateClusterConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateClusterConfiguration", params, optFns, addOperationUpdateClusterConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateClusterConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateClusterConfigurationInput struct {

	// The Amazon Resource Name (ARN) that uniquely identifies the cluster.
	//
	// This member is required.
	ClusterArn *string

	// Represents the configuration that you want MSK to use for the brokers in a
	// cluster.
	//
	// This member is required.
	ConfigurationInfo *types.ConfigurationInfo

	// The version of the cluster that needs to be updated.
	//
	// This member is required.
	CurrentVersion *string
}

type UpdateClusterConfigurationOutput struct {

	// The Amazon Resource Name (ARN) of the cluster.
	ClusterArn *string

	// The Amazon Resource Name (ARN) of the cluster operation.
	ClusterOperationArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateClusterConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpUpdateClusterConfiguration{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpUpdateClusterConfiguration{})
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
	if err := addOpUpdateClusterConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opUpdateClusterConfiguration(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opUpdateClusterConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kafka",
		OperationName: "UpdateClusterConfiguration",
	}
}
