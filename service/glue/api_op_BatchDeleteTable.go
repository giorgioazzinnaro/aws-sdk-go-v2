// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes multiple tables at once. After completing this operation, you no longer
// have access to the table versions and partitions that belong to the deleted
// table. AWS Glue deletes these "orphaned" resources asynchronously in a timely
// manner, at the discretion of the service. To ensure the immediate deletion of
// all related resources, before calling BatchDeleteTable, use DeleteTableVersion
// or BatchDeleteTableVersion, and DeletePartition or BatchDeletePartition, to
// delete any resources that belong to the table.
func (c *Client) BatchDeleteTable(ctx context.Context, params *BatchDeleteTableInput, optFns ...func(*Options)) (*BatchDeleteTableOutput, error) {
	if params == nil {
		params = &BatchDeleteTableInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchDeleteTable", params, optFns, addOperationBatchDeleteTableMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchDeleteTableOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchDeleteTableInput struct {

	// The name of the catalog database in which the tables to delete reside. For Hive
	// compatibility, this name is entirely lowercase.
	//
	// This member is required.
	DatabaseName *string

	// A list of the table to delete.
	//
	// This member is required.
	TablesToDelete []*string

	// The ID of the Data Catalog where the table resides. If none is provided, the AWS
	// account ID is used by default.
	CatalogId *string
}

type BatchDeleteTableOutput struct {

	// A list of errors encountered in attempting to delete the specified tables.
	Errors []*types.TableError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationBatchDeleteTableMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpBatchDeleteTable{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpBatchDeleteTable{})
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
	if err := addOpBatchDeleteTableValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opBatchDeleteTable(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opBatchDeleteTable(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "BatchDeleteTable",
	}
}
