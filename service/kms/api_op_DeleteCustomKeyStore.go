// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a custom key store
// (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html).
// This operation does not delete the AWS CloudHSM cluster that is associated with
// the custom key store, or affect any users or keys in the cluster. The custom key
// store that you delete cannot contain any AWS KMS customer master keys (CMKs)
// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#master_keys).
// Before deleting the key store, verify that you will never need to use any of the
// CMKs in the key store for any cryptographic operations
// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#cryptographic-operations).
// Then, use ScheduleKeyDeletion to delete the AWS KMS customer master keys (CMKs)
// from the key store. When the scheduled waiting period expires, the
// ScheduleKeyDeletion operation deletes the CMKs. Then it makes a best effort to
// delete the key material from the associated cluster. However, you might need to
// manually delete the orphaned key material
// (https://docs.aws.amazon.com/kms/latest/developerguide/fix-keystore.html#fix-keystore-orphaned-key)
// from the cluster and its backups. After all CMKs are deleted from AWS KMS, use
// DisconnectCustomKeyStore to disconnect the key store from AWS KMS. Then, you can
// delete the custom key store. Instead of deleting the custom key store, consider
// using DisconnectCustomKeyStore to disconnect it from AWS KMS. While the key
// store is disconnected, you cannot create or use the CMKs in the key store. But,
// you do not need to delete CMKs and you can reconnect a disconnected custom key
// store at any time. If the operation succeeds, it returns a JSON object with no
// properties. This operation is part of the Custom Key Store feature
// (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html)
// feature in AWS KMS, which combines the convenience and extensive integration of
// AWS KMS with the isolation and control of a single-tenant key store.
func (c *Client) DeleteCustomKeyStore(ctx context.Context, params *DeleteCustomKeyStoreInput, optFns ...func(*Options)) (*DeleteCustomKeyStoreOutput, error) {
	if params == nil {
		params = &DeleteCustomKeyStoreInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteCustomKeyStore", params, optFns, addOperationDeleteCustomKeyStoreMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteCustomKeyStoreOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteCustomKeyStoreInput struct {

	// Enter the ID of the custom key store you want to delete. To find the ID of a
	// custom key store, use the DescribeCustomKeyStores operation.
	//
	// This member is required.
	CustomKeyStoreId *string
}

type DeleteCustomKeyStoreOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteCustomKeyStoreMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpDeleteCustomKeyStore{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpDeleteCustomKeyStore{})
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
	if err := addOpDeleteCustomKeyStoreValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDeleteCustomKeyStore(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDeleteCustomKeyStore(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "DeleteCustomKeyStore",
	}
}
