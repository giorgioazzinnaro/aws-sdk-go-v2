// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Cancels an export task in progress that is exporting a snapshot to Amazon S3.
// Any data that has already been written to the S3 bucket isn't removed.
func (c *Client) CancelExportTask(ctx context.Context, params *CancelExportTaskInput, optFns ...func(*Options)) (*CancelExportTaskOutput, error) {
	if params == nil {
		params = &CancelExportTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CancelExportTask", params, optFns, addOperationCancelExportTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CancelExportTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CancelExportTaskInput struct {

	// The identifier of the snapshot export task to cancel.
	//
	// This member is required.
	ExportTaskIdentifier *string
}

// Contains the details of a snapshot export to Amazon S3. This data type is used
// as a response element in the DescribeExportTasks action.
type CancelExportTaskOutput struct {

	// The data exported from the snapshot. Valid values are the following:
	//
	//     *
	// database - Export all the data from a specified database.
	//
	//     * database.table
	// table-name - Export a table of the snapshot. This format is valid only for RDS
	// for MySQL, RDS for MariaDB, and Aurora MySQL.
	//
	//     * database.schema schema-name
	// - Export a database schema of the snapshot. This format is valid only for RDS
	// for PostgreSQL and Aurora PostgreSQL.
	//
	//     * database.schema.table table-name -
	// Export a table of the database schema. This format is valid only for RDS for
	// PostgreSQL and Aurora PostgreSQL.
	ExportOnly []*string

	// A unique identifier for the snapshot export task. This ID isn't an identifier
	// for the Amazon S3 bucket where the snapshot is exported to.
	ExportTaskIdentifier *string

	// The reason the export failed, if it failed.
	FailureCause *string

	// The name of the IAM role that is used to write to Amazon S3 when exporting a
	// snapshot.
	IamRoleArn *string

	// The ID of the AWS KMS key that is used to encrypt the snapshot when it's
	// exported to Amazon S3. The KMS key ID is the Amazon Resource Name (ARN), the KMS
	// key identifier, or the KMS key alias for the KMS encryption key. The IAM role
	// used for the snapshot export must have encryption and decryption permissions to
	// use this KMS key.
	KmsKeyId *string

	// The progress of the snapshot export task as a percentage.
	PercentProgress *int32

	// The Amazon S3 bucket that the snapshot is exported to.
	S3Bucket *string

	// The Amazon S3 bucket prefix that is the file name and path of the exported
	// snapshot.
	S3Prefix *string

	// The time that the snapshot was created.
	SnapshotTime *time.Time

	// The Amazon Resource Name (ARN) of the snapshot exported to Amazon S3.
	SourceArn *string

	// The progress status of the export task.
	Status *string

	// The time that the snapshot export task completed.
	TaskEndTime *time.Time

	// The time that the snapshot export task started.
	TaskStartTime *time.Time

	// The total amount of data exported, in gigabytes.
	TotalExtractedDataInGB *int32

	// A warning about the snapshot export task.
	WarningMessage *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCancelExportTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsquery_serializeOpCancelExportTask{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsquery_deserializeOpCancelExportTask{})
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
	if err := addOpCancelExportTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opCancelExportTask(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opCancelExportTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "CancelExportTask",
	}
}
