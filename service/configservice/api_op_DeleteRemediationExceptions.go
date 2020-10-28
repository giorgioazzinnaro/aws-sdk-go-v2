// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes one or more remediation exceptions mentioned in the resource keys. AWS
// Config generates a remediation exception when a problem occurs executing a
// remediation action to a specific resource. Remediation exceptions blocks
// auto-remediation until the exception is cleared.
func (c *Client) DeleteRemediationExceptions(ctx context.Context, params *DeleteRemediationExceptionsInput, optFns ...func(*Options)) (*DeleteRemediationExceptionsOutput, error) {
	if params == nil {
		params = &DeleteRemediationExceptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteRemediationExceptions", params, optFns, addOperationDeleteRemediationExceptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteRemediationExceptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteRemediationExceptionsInput struct {

	// The name of the AWS Config rule for which you want to delete remediation
	// exception configuration.
	//
	// This member is required.
	ConfigRuleName *string

	// An exception list of resource exception keys to be processed with the current
	// request. AWS Config adds exception for each resource key. For example, AWS
	// Config adds 3 exceptions for 3 resource keys.
	//
	// This member is required.
	ResourceKeys []*types.RemediationExceptionResourceKey
}

type DeleteRemediationExceptionsOutput struct {

	// Returns a list of failed delete remediation exceptions batch objects. Each
	// object in the batch consists of a list of failed items and failure messages.
	FailedBatches []*types.FailedDeleteRemediationExceptionsBatch

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteRemediationExceptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpDeleteRemediationExceptions{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpDeleteRemediationExceptions{})
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
	if err := addOpDeleteRemediationExceptionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDeleteRemediationExceptions(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDeleteRemediationExceptions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "DeleteRemediationExceptions",
	}
}
