// Code generated by smithy-go-codegen DO NOT EDIT.

package docdb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds an attribute and values to, or removes an attribute and values from, a
// manual DB cluster snapshot. To share a manual cluster snapshot with other AWS
// accounts, specify restore as the AttributeName, and use the ValuesToAdd
// parameter to add a list of IDs of the AWS accounts that are authorized to
// restore the manual cluster snapshot. Use the value all to make the manual
// cluster snapshot public, which means that it can be copied or restored by all
// AWS accounts. Do not add the all value for any manual DB cluster snapshots that
// contain private information that you don't want available to all AWS accounts.
// If a manual cluster snapshot is encrypted, it can be shared, but only by
// specifying a list of authorized AWS account IDs for the ValuesToAdd parameter.
// You can't use all as a value for that parameter in this case.
func (c *Client) ModifyDBClusterSnapshotAttribute(ctx context.Context, params *ModifyDBClusterSnapshotAttributeInput, optFns ...func(*Options)) (*ModifyDBClusterSnapshotAttributeOutput, error) {
	if params == nil {
		params = &ModifyDBClusterSnapshotAttributeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyDBClusterSnapshotAttribute", params, optFns, addOperationModifyDBClusterSnapshotAttributeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyDBClusterSnapshotAttributeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input to ModifyDBClusterSnapshotAttribute.
type ModifyDBClusterSnapshotAttributeInput struct {

	// The name of the cluster snapshot attribute to modify. To manage authorization
	// for other AWS accounts to copy or restore a manual cluster snapshot, set this
	// value to restore.
	//
	// This member is required.
	AttributeName *string

	// The identifier for the cluster snapshot to modify the attributes for.
	//
	// This member is required.
	DBClusterSnapshotIdentifier *string

	// A list of cluster snapshot attributes to add to the attribute specified by
	// AttributeName. To authorize other AWS accounts to copy or restore a manual
	// cluster snapshot, set this list to include one or more AWS account IDs. To make
	// the manual cluster snapshot restorable by any AWS account, set it to all. Do not
	// add the all value for any manual cluster snapshots that contain private
	// information that you don't want to be available to all AWS accounts.
	ValuesToAdd []*string

	// A list of cluster snapshot attributes to remove from the attribute specified by
	// AttributeName. To remove authorization for other AWS accounts to copy or restore
	// a manual cluster snapshot, set this list to include one or more AWS account
	// identifiers. To remove authorization for any AWS account to copy or restore the
	// cluster snapshot, set it to all . If you specify all, an AWS account whose
	// account ID is explicitly added to the restore attribute can still copy or
	// restore a manual cluster snapshot.
	ValuesToRemove []*string
}

type ModifyDBClusterSnapshotAttributeOutput struct {

	// Detailed information about the attributes that are associated with a cluster
	// snapshot.
	DBClusterSnapshotAttributesResult *types.DBClusterSnapshotAttributesResult

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationModifyDBClusterSnapshotAttributeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsquery_serializeOpModifyDBClusterSnapshotAttribute{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsquery_deserializeOpModifyDBClusterSnapshotAttribute{})
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
	if err := addOpModifyDBClusterSnapshotAttributeValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opModifyDBClusterSnapshotAttribute(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opModifyDBClusterSnapshotAttribute(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "ModifyDBClusterSnapshotAttribute",
	}
}
