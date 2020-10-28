// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a replication subnet group given a list of the subnet IDs in a VPC.
func (c *Client) CreateReplicationSubnetGroup(ctx context.Context, params *CreateReplicationSubnetGroupInput, optFns ...func(*Options)) (*CreateReplicationSubnetGroupOutput, error) {
	if params == nil {
		params = &CreateReplicationSubnetGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateReplicationSubnetGroup", params, optFns, addOperationCreateReplicationSubnetGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateReplicationSubnetGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type CreateReplicationSubnetGroupInput struct {

	// The description for the subnet group.
	//
	// This member is required.
	ReplicationSubnetGroupDescription *string

	// The name for the replication subnet group. This value is stored as a lowercase
	// string. Constraints: Must contain no more than 255 alphanumeric characters,
	// periods, spaces, underscores, or hyphens. Must not be "default". Example:
	// mySubnetgroup
	//
	// This member is required.
	ReplicationSubnetGroupIdentifier *string

	// One or more subnet IDs to be assigned to the subnet group.
	//
	// This member is required.
	SubnetIds []*string

	// One or more tags to be assigned to the subnet group.
	Tags []*types.Tag
}

//
type CreateReplicationSubnetGroupOutput struct {

	// The replication subnet group that was created.
	ReplicationSubnetGroup *types.ReplicationSubnetGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateReplicationSubnetGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpCreateReplicationSubnetGroup{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpCreateReplicationSubnetGroup{})
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
	if err := addOpCreateReplicationSubnetGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opCreateReplicationSubnetGroup(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opCreateReplicationSubnetGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dms",
		OperationName: "CreateReplicationSubnetGroup",
	}
}
