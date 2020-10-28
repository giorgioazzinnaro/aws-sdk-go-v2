// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationdiscoveryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/applicationdiscoveryservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists exports as specified by ID. All continuous exports associated with your
// user account can be listed if you call DescribeContinuousExports as is without
// passing any parameters.
func (c *Client) DescribeContinuousExports(ctx context.Context, params *DescribeContinuousExportsInput, optFns ...func(*Options)) (*DescribeContinuousExportsOutput, error) {
	if params == nil {
		params = &DescribeContinuousExportsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeContinuousExports", params, optFns, addOperationDescribeContinuousExportsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeContinuousExportsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeContinuousExportsInput struct {

	// The unique IDs assigned to the exports.
	ExportIds []*string

	// A number between 1 and 100 specifying the maximum number of continuous export
	// descriptions returned.
	MaxResults *int32

	// The token from the previous call to DescribeExportTasks.
	NextToken *string
}

type DescribeContinuousExportsOutput struct {

	// A list of continuous export descriptions.
	Descriptions []*types.ContinuousExportDescription

	// The token from the previous call to DescribeExportTasks.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeContinuousExportsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpDescribeContinuousExports{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpDescribeContinuousExports{})
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
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDescribeContinuousExports(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDescribeContinuousExports(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "discovery",
		OperationName: "DescribeContinuousExports",
	}
}
