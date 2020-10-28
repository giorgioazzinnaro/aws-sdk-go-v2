// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticinference

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticinference/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes information over a provided set of accelerators belonging to an
// account.
func (c *Client) DescribeAccelerators(ctx context.Context, params *DescribeAcceleratorsInput, optFns ...func(*Options)) (*DescribeAcceleratorsOutput, error) {
	if params == nil {
		params = &DescribeAcceleratorsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAccelerators", params, optFns, addOperationDescribeAcceleratorsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAcceleratorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAcceleratorsInput struct {

	// The IDs of the accelerators to describe.
	AcceleratorIds []*string

	// One or more filters. Filter names and values are case-sensitive. Valid filter
	// names are: accelerator-types: can provide a list of accelerator type names to
	// filter for. instance-id: can provide a list of EC2 instance ids to filter for.
	Filters []*types.Filter

	// The total number of items to return in the command's output. If the total number
	// of items available is more than the value specified, a NextToken is provided in
	// the command's output. To resume pagination, provide the NextToken value in the
	// starting-token argument of a subsequent command. Do not use the NextToken
	// response element directly outside of the AWS CLI.
	MaxResults *int32

	// A token to specify where to start paginating. This is the NextToken from a
	// previously truncated response.
	NextToken *string
}

type DescribeAcceleratorsOutput struct {

	// The details of the Elastic Inference Accelerators.
	AcceleratorSet []*types.ElasticInferenceAccelerator

	// A token to specify where to start paginating. This is the NextToken from a
	// previously truncated response.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeAcceleratorsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpDescribeAccelerators{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpDescribeAccelerators{})
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
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDescribeAccelerators(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDescribeAccelerators(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elastic-inference",
		OperationName: "DescribeAccelerators",
	}
}
