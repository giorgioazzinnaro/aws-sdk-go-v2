// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a key group. You cannot delete a key group that is referenced in a cache
// behavior. First update your distributions to remove the key group from all cache
// behaviors, then delete the key group. To delete a key group, you must provide
// the key group’s identifier and version. To get these values, use ListKeyGroups
// followed by GetKeyGroup or GetKeyGroupConfig.
func (c *Client) DeleteKeyGroup(ctx context.Context, params *DeleteKeyGroupInput, optFns ...func(*Options)) (*DeleteKeyGroupOutput, error) {
	if params == nil {
		params = &DeleteKeyGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteKeyGroup", params, optFns, addOperationDeleteKeyGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteKeyGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteKeyGroupInput struct {

	// The identifier of the key group that you are deleting. To get the identifier,
	// use ListKeyGroups.
	//
	// This member is required.
	Id *string

	// The version of the key group that you are deleting. The version is the key
	// group’s ETag value. To get the ETag, use GetKeyGroup or GetKeyGroupConfig.
	IfMatch *string
}

type DeleteKeyGroupOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteKeyGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpDeleteKeyGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpDeleteKeyGroup{}, middleware.After)
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
	addOpDeleteKeyGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteKeyGroup(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteKeyGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "DeleteKeyGroup",
	}
}
