// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// For Redis engine version 6.04 onwards: Creates a Redis user group. For more
// information, see Using Role Based Access Control (RBAC)
// (http://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/Clusters.RBAC.html)
func (c *Client) CreateUserGroup(ctx context.Context, params *CreateUserGroupInput, optFns ...func(*Options)) (*CreateUserGroupOutput, error) {
	if params == nil {
		params = &CreateUserGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateUserGroup", params, optFns, addOperationCreateUserGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateUserGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateUserGroupInput struct {

	// Must be Redis.
	//
	// This member is required.
	Engine *string

	// The ID of the user group.
	//
	// This member is required.
	UserGroupId *string

	// The list of user IDs that belong to the user group.
	UserIds []*string
}

type CreateUserGroupOutput struct {

	// The Amazon Resource Name (ARN) of the user group.
	ARN *string

	// Must be Redis.
	Engine *string

	// A list of updates being applied to the user groups.
	PendingChanges *types.UserGroupPendingChanges

	// A list of replication groups that the user group can access.
	ReplicationGroups []*string

	// Indicates user group status. Can be "creating", "active", "modifying",
	// "deleting".
	Status *string

	// The ID of the user group.
	UserGroupId *string

	// The list of user IDs that belong to the user group.
	UserIds []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateUserGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsquery_serializeOpCreateUserGroup{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsquery_deserializeOpCreateUserGroup{})
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
	if err := addOpCreateUserGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opCreateUserGroup(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opCreateUserGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "CreateUserGroup",
	}
}
