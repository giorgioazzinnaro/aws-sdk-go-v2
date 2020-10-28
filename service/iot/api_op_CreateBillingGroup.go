// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a billing group.
func (c *Client) CreateBillingGroup(ctx context.Context, params *CreateBillingGroupInput, optFns ...func(*Options)) (*CreateBillingGroupOutput, error) {
	if params == nil {
		params = &CreateBillingGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateBillingGroup", params, optFns, addOperationCreateBillingGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateBillingGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateBillingGroupInput struct {

	// The name you wish to give to the billing group.
	//
	// This member is required.
	BillingGroupName *string

	// The properties of the billing group.
	BillingGroupProperties *types.BillingGroupProperties

	// Metadata which can be used to manage the billing group.
	Tags []*types.Tag
}

type CreateBillingGroupOutput struct {

	// The ARN of the billing group.
	BillingGroupArn *string

	// The ID of the billing group.
	BillingGroupId *string

	// The name you gave to the billing group.
	BillingGroupName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateBillingGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpCreateBillingGroup{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpCreateBillingGroup{})
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
	if err := addOpCreateBillingGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opCreateBillingGroup(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opCreateBillingGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "CreateBillingGroup",
	}
}
