// Code generated by smithy-go-codegen DO NOT EDIT.

package ram

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Enables resource sharing within your AWS Organization. The caller must be the
// master account for the AWS Organization.
func (c *Client) EnableSharingWithAwsOrganization(ctx context.Context, params *EnableSharingWithAwsOrganizationInput, optFns ...func(*Options)) (*EnableSharingWithAwsOrganizationOutput, error) {
	if params == nil {
		params = &EnableSharingWithAwsOrganizationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "EnableSharingWithAwsOrganization", params, optFns, addOperationEnableSharingWithAwsOrganizationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*EnableSharingWithAwsOrganizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EnableSharingWithAwsOrganizationInput struct {
}

type EnableSharingWithAwsOrganizationOutput struct {

	// Indicates whether the request succeeded.
	ReturnValue *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationEnableSharingWithAwsOrganizationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpEnableSharingWithAwsOrganization{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpEnableSharingWithAwsOrganization{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opEnableSharingWithAwsOrganization(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opEnableSharingWithAwsOrganization(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ram",
		OperationName: "EnableSharingWithAwsOrganization",
	}
}
