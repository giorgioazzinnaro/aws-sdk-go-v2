// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds one or more instance groups to a running cluster.
func (c *Client) AddInstanceGroups(ctx context.Context, params *AddInstanceGroupsInput, optFns ...func(*Options)) (*AddInstanceGroupsOutput, error) {
	if params == nil {
		params = &AddInstanceGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AddInstanceGroups", params, optFns, addOperationAddInstanceGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AddInstanceGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input to an AddInstanceGroups call.
type AddInstanceGroupsInput struct {

	// Instance groups to add.
	//
	// This member is required.
	InstanceGroups []*types.InstanceGroupConfig

	// Job flow in which to add the instance groups.
	//
	// This member is required.
	JobFlowId *string
}

// Output from an AddInstanceGroups call.
type AddInstanceGroupsOutput struct {

	// The Amazon Resource Name of the cluster.
	ClusterArn *string

	// Instance group IDs of the newly created instance groups.
	InstanceGroupIds []*string

	// The job flow ID in which the instance groups are added.
	JobFlowId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAddInstanceGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpAddInstanceGroups{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpAddInstanceGroups{})
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
	if err := addOpAddInstanceGroupsValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opAddInstanceGroups(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opAddInstanceGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticmapreduce",
		OperationName: "AddInstanceGroups",
	}
}
