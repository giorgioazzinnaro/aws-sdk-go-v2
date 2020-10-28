// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes a specified database from a Data Catalog. After completing this
// operation, you no longer have access to the tables (and all table versions and
// partitions that might belong to the tables) and the user-defined functions in
// the deleted database. AWS Glue deletes these "orphaned" resources asynchronously
// in a timely manner, at the discretion of the service. To ensure the immediate
// deletion of all related resources, before calling DeleteDatabase, use
// DeleteTableVersion or BatchDeleteTableVersion, DeletePartition or
// BatchDeletePartition, DeleteUserDefinedFunction, and DeleteTable or
// BatchDeleteTable, to delete any resources that belong to the database.
func (c *Client) DeleteDatabase(ctx context.Context, params *DeleteDatabaseInput, optFns ...func(*Options)) (*DeleteDatabaseOutput, error) {
	if params == nil {
		params = &DeleteDatabaseInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteDatabase", params, optFns, addOperationDeleteDatabaseMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteDatabaseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteDatabaseInput struct {

	// The name of the database to delete. For Hive compatibility, this must be all
	// lowercase.
	//
	// This member is required.
	Name *string

	// The ID of the Data Catalog in which the database resides. If none is provided,
	// the AWS account ID is used by default.
	CatalogId *string
}

type DeleteDatabaseOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteDatabaseMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpDeleteDatabase{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpDeleteDatabase{})
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
	if err := addOpDeleteDatabaseValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDeleteDatabase(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDeleteDatabase(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "DeleteDatabase",
	}
}
