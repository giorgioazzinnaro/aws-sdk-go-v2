// Code generated by smithy-go-codegen DO NOT EDIT.

package inspector

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/inspector/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the assessment templates that are specified by the ARNs of the
// assessment templates.
func (c *Client) DescribeAssessmentTemplates(ctx context.Context, params *DescribeAssessmentTemplatesInput, optFns ...func(*Options)) (*DescribeAssessmentTemplatesOutput, error) {
	if params == nil {
		params = &DescribeAssessmentTemplatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAssessmentTemplates", params, optFns, addOperationDescribeAssessmentTemplatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAssessmentTemplatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAssessmentTemplatesInput struct {
	AssessmentTemplateArns []*string
}

type DescribeAssessmentTemplatesOutput struct {

	// Information about the assessment templates.
	//
	// This member is required.
	AssessmentTemplates []*types.AssessmentTemplate

	// Assessment template details that cannot be described. An error code is provided
	// for each failed item.
	//
	// This member is required.
	FailedItems map[string]*types.FailedItemDetails

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeAssessmentTemplatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpDescribeAssessmentTemplates{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpDescribeAssessmentTemplates{})
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
	if err := addOpDescribeAssessmentTemplatesValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDescribeAssessmentTemplates(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDescribeAssessmentTemplates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "inspector",
		OperationName: "DescribeAssessmentTemplates",
	}
}
