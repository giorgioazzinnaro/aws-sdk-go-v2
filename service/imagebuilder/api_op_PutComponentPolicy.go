// Code generated by smithy-go-codegen DO NOT EDIT.

package imagebuilder

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Applies a policy to a component. We recommend that you call the RAM API
// CreateResourceShare
// (https://docs.aws.amazon.com/ram/latest/APIReference/API_CreateResourceShare.html)
// to share resources. If you call the Image Builder API PutComponentPolicy, you
// must also call the RAM API PromoteResourceShareCreatedFromPolicy
// (https://docs.aws.amazon.com/ram/latest/APIReference/API_PromoteResourceShareCreatedFromPolicy.html)
// in order for the resource to be visible to all principals with whom the resource
// is shared.
func (c *Client) PutComponentPolicy(ctx context.Context, params *PutComponentPolicyInput, optFns ...func(*Options)) (*PutComponentPolicyOutput, error) {
	if params == nil {
		params = &PutComponentPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutComponentPolicy", params, optFns, addOperationPutComponentPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutComponentPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutComponentPolicyInput struct {

	// The Amazon Resource Name (ARN) of the component that this policy should be
	// applied to.
	//
	// This member is required.
	ComponentArn *string

	// The policy to apply.
	//
	// This member is required.
	Policy *string
}

type PutComponentPolicyOutput struct {

	// The Amazon Resource Name (ARN) of the component that this policy was applied to.
	ComponentArn *string

	// The request ID that uniquely identifies this request.
	RequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutComponentPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpPutComponentPolicy{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpPutComponentPolicy{})
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
	if err := addOpPutComponentPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opPutComponentPolicy(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opPutComponentPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "imagebuilder",
		OperationName: "PutComponentPolicy",
	}
}
