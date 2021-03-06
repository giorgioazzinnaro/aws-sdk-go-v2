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

// Retrieves the metadata for a given workflow run.
func (c *Client) GetWorkflowRun(ctx context.Context, params *GetWorkflowRunInput, optFns ...func(*Options)) (*GetWorkflowRunOutput, error) {
	if params == nil {
		params = &GetWorkflowRunInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetWorkflowRun", params, optFns, addOperationGetWorkflowRunMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetWorkflowRunOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetWorkflowRunInput struct {

	// Name of the workflow being run.
	//
	// This member is required.
	Name *string

	// The ID of the workflow run.
	//
	// This member is required.
	RunId *string

	// Specifies whether to include the workflow graph in response or not.
	IncludeGraph *bool
}

type GetWorkflowRunOutput struct {

	// The requested workflow run metadata.
	Run *types.WorkflowRun

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetWorkflowRunMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetWorkflowRun{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetWorkflowRun{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpGetWorkflowRunValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetWorkflowRun(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetWorkflowRun(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "GetWorkflowRun",
	}
}
