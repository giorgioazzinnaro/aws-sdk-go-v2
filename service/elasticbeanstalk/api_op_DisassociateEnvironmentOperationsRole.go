// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disassociate the operations role from an environment. After this call is made,
// Elastic Beanstalk uses the caller's permissions for permissions to downstream
// services during subsequent calls acting on this environment. For more
// information, see Operations roles
// (https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/iam-operationsrole.html)
// in the AWS Elastic Beanstalk Developer Guide.
func (c *Client) DisassociateEnvironmentOperationsRole(ctx context.Context, params *DisassociateEnvironmentOperationsRoleInput, optFns ...func(*Options)) (*DisassociateEnvironmentOperationsRoleOutput, error) {
	if params == nil {
		params = &DisassociateEnvironmentOperationsRoleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateEnvironmentOperationsRole", params, optFns, addOperationDisassociateEnvironmentOperationsRoleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateEnvironmentOperationsRoleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request to disassociate the operations role from an environment.
type DisassociateEnvironmentOperationsRoleInput struct {

	// The name of the environment from which to disassociate the operations role.
	//
	// This member is required.
	EnvironmentName *string
}

type DisassociateEnvironmentOperationsRoleOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDisassociateEnvironmentOperationsRoleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsquery_serializeOpDisassociateEnvironmentOperationsRole{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsquery_deserializeOpDisassociateEnvironmentOperationsRole{})
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
	if err := addOpDisassociateEnvironmentOperationsRoleValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDisassociateEnvironmentOperationsRole(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDisassociateEnvironmentOperationsRole(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticbeanstalk",
		OperationName: "DisassociateEnvironmentOperationsRole",
	}
}
