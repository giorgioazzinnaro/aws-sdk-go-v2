// Code generated by smithy-go-codegen DO NOT EDIT.

package marketplacecatalog

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the metadata and content of the entity.
func (c *Client) DescribeEntity(ctx context.Context, params *DescribeEntityInput, optFns ...func(*Options)) (*DescribeEntityOutput, error) {
	if params == nil {
		params = &DescribeEntityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEntity", params, optFns, addOperationDescribeEntityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEntityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEntityInput struct {

	// Required. The catalog related to the request. Fixed value: AWSMarketplace
	//
	// This member is required.
	Catalog *string

	// Required. The unique ID of the entity to describe.
	//
	// This member is required.
	EntityId *string
}

type DescribeEntityOutput struct {

	// This stringified JSON object includes the details of the entity.
	Details *string

	// The ARN associated to the unique identifier for the change set referenced in
	// this request.
	EntityArn *string

	// The identifier of the entity, in the format of EntityId@RevisionId.
	EntityIdentifier *string

	// The named type of the entity, in the format of EntityType@Version.
	EntityType *string

	// The last modified date of the entity, in ISO 8601 format (2018-02-27T13:45:22Z).
	LastModifiedDate *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeEntityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpDescribeEntity{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpDescribeEntity{})
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
	if err := addOpDescribeEntityValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDescribeEntity(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDescribeEntity(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "aws-marketplace",
		OperationName: "DescribeEntity",
	}
}
