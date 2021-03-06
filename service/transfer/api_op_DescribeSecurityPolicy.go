// Code generated by smithy-go-codegen DO NOT EDIT.

package transfer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/transfer/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the security policy that is attached to your file transfer
// protocol-enabled server. The response contains a description of the security
// policy's properties. For more information about security policies, see Working
// with security policies
// (https://docs.aws.amazon.com/transfer/latest/userguide/security-policies.html).
func (c *Client) DescribeSecurityPolicy(ctx context.Context, params *DescribeSecurityPolicyInput, optFns ...func(*Options)) (*DescribeSecurityPolicyOutput, error) {
	if params == nil {
		params = &DescribeSecurityPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeSecurityPolicy", params, optFns, addOperationDescribeSecurityPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeSecurityPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeSecurityPolicyInput struct {

	// Specifies the name of the security policy that is attached to the server.
	//
	// This member is required.
	SecurityPolicyName *string
}

type DescribeSecurityPolicyOutput struct {

	// An array containing the properties of the security policy.
	//
	// This member is required.
	SecurityPolicy *types.DescribedSecurityPolicy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeSecurityPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeSecurityPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeSecurityPolicy{}, middleware.After)
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
	addOpDescribeSecurityPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSecurityPolicy(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeSecurityPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transfer",
		OperationName: "DescribeSecurityPolicy",
	}
}
