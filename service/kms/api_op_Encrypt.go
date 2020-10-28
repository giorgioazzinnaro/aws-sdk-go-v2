// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Encrypts plaintext into ciphertext by using a customer master key (CMK). The
// Encrypt operation has two primary use cases:
//
//     * You can encrypt small
// amounts of arbitrary data, such as a personal identifier or database password,
// or other sensitive information.
//
//     * You can use the Encrypt operation to move
// encrypted data from one AWS Region to another. For example, in Region A,
// generate a data key and use the plaintext key to encrypt your data. Then, in
// Region A, use the Encrypt operation to encrypt the plaintext data key under a
// CMK in Region B. Now, you can move the encrypted data and the encrypted data key
// to Region B. When necessary, you can decrypt the encrypted data key and the
// encrypted data entirely within in Region B.
//
// You don't need to use the Encrypt
// operation to encrypt a data key. The GenerateDataKey and GenerateDataKeyPair
// operations return a plaintext data key and an encrypted copy of that data key.
// When you encrypt data, you must specify a symmetric or asymmetric CMK to use in
// the encryption operation. The CMK must have a KeyUsage value of ENCRYPT_DECRYPT.
// To find the KeyUsage of a CMK, use the DescribeKey operation. If you use a
// symmetric CMK, you can use an encryption context to add additional security to
// your encryption operation. If you specify an EncryptionContext when encrypting
// data, you must specify the same encryption context (a case-sensitive exact
// match) when decrypting the data. Otherwise, the request to decrypt fails with an
// InvalidCiphertextException. For more information, see Encryption Context
// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context)
// in the AWS Key Management Service Developer Guide. If you specify an asymmetric
// CMK, you must also specify the encryption algorithm. The algorithm must be
// compatible with the CMK type. When you use an asymmetric CMK to encrypt or
// reencrypt data, be sure to record the CMK and encryption algorithm that you
// choose. You will be required to provide the same CMK and encryption algorithm
// when you decrypt the data. If the CMK and algorithm do not match the values used
// to encrypt the data, the decrypt operation fails. You are not required to supply
// the CMK ID and encryption algorithm when you decrypt with symmetric CMKs because
// AWS KMS stores this information in the ciphertext blob. AWS KMS cannot store
// metadata in ciphertext generated with asymmetric keys. The standard format for
// asymmetric key ciphertext does not include configurable fields. The maximum size
// of the data that you can encrypt varies with the type of CMK and the encryption
// algorithm that you choose.
//
//     * Symmetric CMKs
//
//         * SYMMETRIC_DEFAULT:
// 4096 bytes
//
//     * RSA_2048
//
//         * RSAES_OAEP_SHA_1: 214 bytes
//
//         *
// RSAES_OAEP_SHA_256: 190 bytes
//
//     * RSA_3072
//
//         * RSAES_OAEP_SHA_1: 342
// bytes
//
//         * RSAES_OAEP_SHA_256: 318 bytes
//
//     * RSA_4096
//
//         *
// RSAES_OAEP_SHA_1: 470 bytes
//
//         * RSAES_OAEP_SHA_256: 446 bytes
//
// The CMK
// that you use for this operation must be in a compatible key state. For details,
// see How Key State Affects Use of a Customer Master Key
// (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html) in the
// AWS Key Management Service Developer Guide. To perform this operation on a CMK
// in a different AWS account, specify the key ARN or alias ARN in the value of the
// KeyId parameter.
func (c *Client) Encrypt(ctx context.Context, params *EncryptInput, optFns ...func(*Options)) (*EncryptOutput, error) {
	if params == nil {
		params = &EncryptInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "Encrypt", params, optFns, addOperationEncryptMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*EncryptOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EncryptInput struct {

	// A unique identifier for the customer master key (CMK). To specify a CMK, use its
	// key ID, Amazon Resource Name (ARN), alias name, or alias ARN. When using an
	// alias name, prefix it with "alias/". To specify a CMK in a different AWS
	// account, you must use the key ARN or alias ARN. For example:
	//
	//     * Key ID:
	// 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//     * Key ARN:
	// arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//
	// * Alias name: alias/ExampleAlias
	//
	//     * Alias ARN:
	// arn:aws:kms:us-east-2:111122223333:alias/ExampleAlias
	//
	// To get the key ID and key
	// ARN for a CMK, use ListKeys or DescribeKey. To get the alias name and alias ARN,
	// use ListAliases.
	//
	// This member is required.
	KeyId *string

	// Data to be encrypted.
	//
	// This member is required.
	Plaintext []byte

	// Specifies the encryption algorithm that AWS KMS will use to encrypt the
	// plaintext message. The algorithm must be compatible with the CMK that you
	// specify. This parameter is required only for asymmetric CMKs. The default value,
	// SYMMETRIC_DEFAULT, is the algorithm used for symmetric CMKs. If you are using an
	// asymmetric CMK, we recommend RSAES_OAEP_SHA_256.
	EncryptionAlgorithm types.EncryptionAlgorithmSpec

	// Specifies the encryption context that will be used to encrypt the data. An
	// encryption context is valid only for cryptographic operations
	// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#cryptographic-operations)
	// with a symmetric CMK. The standard asymmetric encryption algorithms that AWS KMS
	// uses do not support an encryption context. An encryption context is a collection
	// of non-secret key-value pairs that represents additional authenticated data.
	// When you use an encryption context to encrypt data, you must specify the same
	// (an exact case-sensitive match) encryption context to decrypt the data. An
	// encryption context is optional when encrypting with a symmetric CMK, but it is
	// highly recommended. For more information, see Encryption Context
	// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context)
	// in the AWS Key Management Service Developer Guide.
	EncryptionContext map[string]*string

	// A list of grant tokens. For more information, see Grant Tokens
	// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#grant_token)
	// in the AWS Key Management Service Developer Guide.
	GrantTokens []*string
}

type EncryptOutput struct {

	// The encrypted plaintext. When you use the HTTP API or the AWS CLI, the value is
	// Base64-encoded. Otherwise, it is not Base64-encoded.
	CiphertextBlob []byte

	// The encryption algorithm that was used to encrypt the plaintext.
	EncryptionAlgorithm types.EncryptionAlgorithmSpec

	// The Amazon Resource Name (key ARN
	// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-id-key-ARN))
	// of the CMK that was used to encrypt the plaintext.
	KeyId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationEncryptMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpEncrypt{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpEncrypt{})
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
	if err := addOpEncryptValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opEncrypt(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opEncrypt(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "Encrypt",
	}
}
