// Code generated by smithy-go-codegen DO NOT EDIT.

package iotanalytics

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotanalytics/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the settings of a data store.
func (c *Client) UpdateDatastore(ctx context.Context, params *UpdateDatastoreInput, optFns ...func(*Options)) (*UpdateDatastoreOutput, error) {
	if params == nil {
		params = &UpdateDatastoreInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDatastore", params, optFns, addOperationUpdateDatastoreMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDatastoreOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDatastoreInput struct {

	// The name of the data store to be updated.
	//
	// This member is required.
	DatastoreName *string

	// Where data store data is stored. You may choose one of "serviceManagedS3" or
	// "customerManagedS3" storage. If not specified, the default is
	// "serviceManagedS3". This cannot be changed after the data store is created.
	DatastoreStorage *types.DatastoreStorage

	// How long, in days, message data is kept for the data store. The retention period
	// cannot be updated if the data store's S3 storage is customer-managed.
	RetentionPeriod *types.RetentionPeriod
}

type UpdateDatastoreOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateDatastoreMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpUpdateDatastore{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpUpdateDatastore{})
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
	if err := addOpUpdateDatastoreValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opUpdateDatastore(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opUpdateDatastore(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotanalytics",
		OperationName: "UpdateDatastore",
	}
}
