// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a dataset group. Before you delete a dataset group, you must delete the
// following:
//
// * All associated event trackers.
//
// * All associated solutions.
//
// * All
// datasets in the dataset group.
func (c *Client) DeleteDatasetGroup(ctx context.Context, params *DeleteDatasetGroupInput, optFns ...func(*Options)) (*DeleteDatasetGroupOutput, error) {
	if params == nil {
		params = &DeleteDatasetGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteDatasetGroup", params, optFns, addOperationDeleteDatasetGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteDatasetGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteDatasetGroupInput struct {

	// The ARN of the dataset group to delete.
	//
	// This member is required.
	DatasetGroupArn *string
}

type DeleteDatasetGroupOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteDatasetGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteDatasetGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteDatasetGroup{}, middleware.After)
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
	addOpDeleteDatasetGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDatasetGroup(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteDatasetGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "DeleteDatasetGroup",
	}
}
