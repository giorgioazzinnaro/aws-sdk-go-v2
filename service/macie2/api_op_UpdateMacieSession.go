// Code generated by smithy-go-codegen DO NOT EDIT.

package macie2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/macie2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Suspends or re-enables an Amazon Macie account, or updates the configuration
// settings for a Macie account.
func (c *Client) UpdateMacieSession(ctx context.Context, params *UpdateMacieSessionInput, optFns ...func(*Options)) (*UpdateMacieSessionOutput, error) {
	if params == nil {
		params = &UpdateMacieSessionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateMacieSession", params, optFns, addOperationUpdateMacieSessionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateMacieSessionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateMacieSessionInput struct {

	// Specifies how often to publish updates to policy findings for the account. This
	// includes publishing updates to AWS Security Hub and Amazon EventBridge (formerly
	// called Amazon CloudWatch Events).
	FindingPublishingFrequency types.FindingPublishingFrequency

	// Specifies whether to change the status of the account. Valid values are:
	// ENABLED, resume all Amazon Macie activities for the account; and, PAUSED,
	// suspend all Macie activities for the account.
	Status types.MacieStatus
}

type UpdateMacieSessionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateMacieSessionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpUpdateMacieSession{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpUpdateMacieSession{})
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
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opUpdateMacieSession(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opUpdateMacieSession(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "macie2",
		OperationName: "UpdateMacieSession",
	}
}
