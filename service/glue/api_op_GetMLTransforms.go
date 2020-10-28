// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets a sortable, filterable list of existing AWS Glue machine learning
// transforms. Machine learning transforms are a special type of transform that use
// machine learning to learn the details of the transformation to be performed by
// learning from examples provided by humans. These transformations are then saved
// by AWS Glue, and you can retrieve their metadata by calling GetMLTransforms.
func (c *Client) GetMLTransforms(ctx context.Context, params *GetMLTransformsInput, optFns ...func(*Options)) (*GetMLTransformsOutput, error) {
	if params == nil {
		params = &GetMLTransformsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMLTransforms", params, optFns, addOperationGetMLTransformsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMLTransformsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMLTransformsInput struct {

	// The filter transformation criteria.
	Filter *types.TransformFilterCriteria

	// The maximum number of results to return.
	MaxResults *int32

	// A paginated token to offset the results.
	NextToken *string

	// The sorting criteria.
	Sort *types.TransformSortCriteria
}

type GetMLTransformsOutput struct {

	// A list of machine learning transforms.
	//
	// This member is required.
	Transforms []*types.MLTransform

	// A pagination token, if more results are available.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetMLTransformsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpGetMLTransforms{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpGetMLTransforms{})
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
	if err := addOpGetMLTransformsValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opGetMLTransforms(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opGetMLTransforms(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "GetMLTransforms",
	}
}
