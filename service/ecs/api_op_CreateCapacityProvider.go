// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new capacity provider. Capacity providers are associated with an
// Amazon ECS cluster and are used in capacity provider strategies to facilitate
// cluster auto scaling. Only capacity providers using an Auto Scaling group can be
// created. Amazon ECS tasks on AWS Fargate use the FARGATE and FARGATE_SPOT
// capacity providers which are already created and available to all accounts in
// Regions supported by AWS Fargate.
func (c *Client) CreateCapacityProvider(ctx context.Context, params *CreateCapacityProviderInput, optFns ...func(*Options)) (*CreateCapacityProviderOutput, error) {
	if params == nil {
		params = &CreateCapacityProviderInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCapacityProvider", params, optFns, addOperationCreateCapacityProviderMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCapacityProviderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateCapacityProviderInput struct {

	// The details of the Auto Scaling group for the capacity provider.
	//
	// This member is required.
	AutoScalingGroupProvider *types.AutoScalingGroupProvider

	// The name of the capacity provider. Up to 255 characters are allowed, including
	// letters (upper and lowercase), numbers, underscores, and hyphens. The name
	// cannot be prefixed with "aws", "ecs", or "fargate".
	//
	// This member is required.
	Name *string

	// The metadata that you apply to the capacity provider to help you categorize and
	// organize them. Each tag consists of a key and an optional value, both of which
	// you define. The following basic restrictions apply to tags:
	//
	//     * Maximum
	// number of tags per resource - 50
	//
	//     * For each resource, each tag key must be
	// unique, and each tag key can have only one value.
	//
	//     * Maximum key length -
	// 128 Unicode characters in UTF-8
	//
	//     * Maximum value length - 256 Unicode
	// characters in UTF-8
	//
	//     * If your tagging schema is used across multiple
	// services and resources, remember that other services may have restrictions on
	// allowed characters. Generally allowed characters are: letters, numbers, and
	// spaces representable in UTF-8, and the following characters: + - = . _ : / @.
	//
	//
	// * Tag keys and values are case-sensitive.
	//
	//     * Do not use aws:, AWS:, or any
	// upper or lowercase combination of such as a prefix for either keys or values as
	// it is reserved for AWS use. You cannot edit or delete tag keys or values with
	// this prefix. Tags with this prefix do not count against your tags per resource
	// limit.
	Tags []*types.Tag
}

type CreateCapacityProviderOutput struct {

	// The full description of the new capacity provider.
	CapacityProvider *types.CapacityProvider

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateCapacityProviderMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpCreateCapacityProvider{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpCreateCapacityProvider{})
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
	if err := addOpCreateCapacityProviderValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opCreateCapacityProvider(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opCreateCapacityProvider(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecs",
		OperationName: "CreateCapacityProvider",
	}
}
