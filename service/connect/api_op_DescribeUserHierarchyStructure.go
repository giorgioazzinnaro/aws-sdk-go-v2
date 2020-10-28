// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the hierarchy structure of the specified Amazon Connect instance.
func (c *Client) DescribeUserHierarchyStructure(ctx context.Context, params *DescribeUserHierarchyStructureInput, optFns ...func(*Options)) (*DescribeUserHierarchyStructureOutput, error) {
	if params == nil {
		params = &DescribeUserHierarchyStructureInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeUserHierarchyStructure", params, optFns, addOperationDescribeUserHierarchyStructureMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeUserHierarchyStructureOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeUserHierarchyStructureInput struct {

	// The identifier of the Amazon Connect instance.
	//
	// This member is required.
	InstanceId *string
}

type DescribeUserHierarchyStructureOutput struct {

	// Information about the hierarchy structure.
	HierarchyStructure *types.HierarchyStructure

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeUserHierarchyStructureMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpDescribeUserHierarchyStructure{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpDescribeUserHierarchyStructure{})
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
	if err := addOpDescribeUserHierarchyStructureValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDescribeUserHierarchyStructure(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDescribeUserHierarchyStructure(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "DescribeUserHierarchyStructure",
	}
}
