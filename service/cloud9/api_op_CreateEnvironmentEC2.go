// Code generated by smithy-go-codegen DO NOT EDIT.

package cloud9

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloud9/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an AWS Cloud9 development environment, launches an Amazon Elastic
// Compute Cloud (Amazon EC2) instance, and then connects from the instance to the
// environment.
func (c *Client) CreateEnvironmentEC2(ctx context.Context, params *CreateEnvironmentEC2Input, optFns ...func(*Options)) (*CreateEnvironmentEC2Output, error) {
	if params == nil {
		params = &CreateEnvironmentEC2Input{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateEnvironmentEC2", params, optFns, addOperationCreateEnvironmentEC2Middlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateEnvironmentEC2Output)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateEnvironmentEC2Input struct {

	// The type of instance to connect to the environment (for example, t2.micro).
	//
	// This member is required.
	InstanceType *string

	// The name of the environment to create. This name is visible to other AWS IAM
	// users in the same AWS account.
	//
	// This member is required.
	Name *string

	// The number of minutes until the running instance is shut down after the
	// environment has last been used.
	AutomaticStopTimeMinutes *int32

	// A unique, case-sensitive string that helps AWS Cloud9 to ensure this operation
	// completes no more than one time. For more information, see Client Tokens
	// (http://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html)
	// in the Amazon EC2 API Reference.
	ClientRequestToken *string

	// The connection type used for connecting to an Amazon EC2 environment.
	ConnectionType types.ConnectionType

	// The description of the environment to create.
	Description *string

	// The Amazon Resource Name (ARN) of the environment owner. This ARN can be the ARN
	// of any AWS IAM principal. If this value is not specified, the ARN defaults to
	// this environment's creator.
	OwnerArn *string

	// The ID of the subnet in Amazon VPC that AWS Cloud9 will use to communicate with
	// the Amazon EC2 instance.
	SubnetId *string

	// An array of key-value pairs that will be associated with the new AWS Cloud9
	// development environment.
	Tags []*types.Tag
}

type CreateEnvironmentEC2Output struct {

	// The ID of the environment that was created.
	EnvironmentId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateEnvironmentEC2Middlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpCreateEnvironmentEC2{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpCreateEnvironmentEC2{})
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
	if err := addOpCreateEnvironmentEC2ValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opCreateEnvironmentEC2(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opCreateEnvironmentEC2(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloud9",
		OperationName: "CreateEnvironmentEC2",
	}
}
