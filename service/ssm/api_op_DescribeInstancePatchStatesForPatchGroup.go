// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the high-level patch state for the instances in the specified patch
// group.
func (c *Client) DescribeInstancePatchStatesForPatchGroup(ctx context.Context, params *DescribeInstancePatchStatesForPatchGroupInput, optFns ...func(*Options)) (*DescribeInstancePatchStatesForPatchGroupOutput, error) {
	if params == nil {
		params = &DescribeInstancePatchStatesForPatchGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeInstancePatchStatesForPatchGroup", params, optFns, addOperationDescribeInstancePatchStatesForPatchGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeInstancePatchStatesForPatchGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeInstancePatchStatesForPatchGroupInput struct {

	// The name of the patch group for which the patch state information should be
	// retrieved.
	//
	// This member is required.
	PatchGroup *string

	// Each entry in the array is a structure containing: Key (string between 1 and 200
	// characters) Values (array containing a single string) Type (string "Equal",
	// "NotEqual", "LessThan", "GreaterThan")
	Filters []*types.InstancePatchStateFilter

	// The maximum number of patches to return (per page).
	MaxResults *int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string
}

type DescribeInstancePatchStatesForPatchGroupOutput struct {

	// The high-level patch state for the requested instances.
	InstancePatchStates []*types.InstancePatchState

	// The token to use when requesting the next set of items. If there are no
	// additional items to return, the string is empty.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeInstancePatchStatesForPatchGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpDescribeInstancePatchStatesForPatchGroup{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpDescribeInstancePatchStatesForPatchGroup{})
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
	if err := addOpDescribeInstancePatchStatesForPatchGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDescribeInstancePatchStatesForPatchGroup(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDescribeInstancePatchStatesForPatchGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "DescribeInstancePatchStatesForPatchGroup",
	}
}
