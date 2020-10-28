// Code generated by smithy-go-codegen DO NOT EDIT.

package serverlessapplicationrepository

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/serverlessapplicationrepository/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Sets the permission policy for an application. For the list of actions supported
// for this operation, see Application Permissions
// (https://docs.aws.amazon.com/serverlessrepo/latest/devguide/access-control-resource-based.html#application-permissions)
// .
func (c *Client) PutApplicationPolicy(ctx context.Context, params *PutApplicationPolicyInput, optFns ...func(*Options)) (*PutApplicationPolicyOutput, error) {
	if params == nil {
		params = &PutApplicationPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutApplicationPolicy", params, optFns, addOperationPutApplicationPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutApplicationPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutApplicationPolicyInput struct {

	// The Amazon Resource Name (ARN) of the application.
	//
	// This member is required.
	ApplicationId *string

	// An array of policy statements applied to the application.
	//
	// This member is required.
	Statements []*types.ApplicationPolicyStatement
}

type PutApplicationPolicyOutput struct {

	// An array of policy statements applied to the application.
	Statements []*types.ApplicationPolicyStatement

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutApplicationPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpPutApplicationPolicy{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpPutApplicationPolicy{})
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
	if err := addOpPutApplicationPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opPutApplicationPolicy(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opPutApplicationPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "serverlessrepo",
		OperationName: "PutApplicationPolicy",
	}
}
