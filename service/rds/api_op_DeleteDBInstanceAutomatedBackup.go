// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes automated backups based on the source instance's DbiResourceId value or
// the restorable instance's resource ID.
func (c *Client) DeleteDBInstanceAutomatedBackup(ctx context.Context, params *DeleteDBInstanceAutomatedBackupInput, optFns ...func(*Options)) (*DeleteDBInstanceAutomatedBackupOutput, error) {
	if params == nil {
		params = &DeleteDBInstanceAutomatedBackupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteDBInstanceAutomatedBackup", params, optFns, addOperationDeleteDBInstanceAutomatedBackupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteDBInstanceAutomatedBackupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Parameter input for the DeleteDBInstanceAutomatedBackup operation.
type DeleteDBInstanceAutomatedBackupInput struct {

	// The identifier for the source DB instance, which can't be changed and which is
	// unique to an AWS Region.
	//
	// This member is required.
	DbiResourceId *string
}

type DeleteDBInstanceAutomatedBackupOutput struct {

	// An automated backup of a DB instance. It it consists of system backups,
	// transaction logs, and the database instance properties that existed at the time
	// you deleted the source instance.
	DBInstanceAutomatedBackup *types.DBInstanceAutomatedBackup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteDBInstanceAutomatedBackupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsquery_serializeOpDeleteDBInstanceAutomatedBackup{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsquery_deserializeOpDeleteDBInstanceAutomatedBackup{})
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
	if err := addOpDeleteDBInstanceAutomatedBackupValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDeleteDBInstanceAutomatedBackup(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDeleteDBInstanceAutomatedBackup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DeleteDBInstanceAutomatedBackup",
	}
}
