// Code generated by smithy-go-codegen DO NOT EDIT.

package appstream

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an image builder. An image builder is a virtual machine that is used to
// create an image. The initial state of the builder is PENDING. When it is ready,
// the state is RUNNING.
func (c *Client) CreateImageBuilder(ctx context.Context, params *CreateImageBuilderInput, optFns ...func(*Options)) (*CreateImageBuilderOutput, error) {
	if params == nil {
		params = &CreateImageBuilderInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateImageBuilder", params, optFns, addOperationCreateImageBuilderMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateImageBuilderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateImageBuilderInput struct {

	// The instance type to use when launching the image builder. The following
	// instance types are available:
	//
	//     * stream.standard.medium
	//
	//     *
	// stream.standard.large
	//
	//     * stream.compute.large
	//
	//     * stream.compute.xlarge
	//
	//
	// * stream.compute.2xlarge
	//
	//     * stream.compute.4xlarge
	//
	//     *
	// stream.compute.8xlarge
	//
	//     * stream.memory.large
	//
	//     * stream.memory.xlarge
	//
	//
	// * stream.memory.2xlarge
	//
	//     * stream.memory.4xlarge
	//
	//     *
	// stream.memory.8xlarge
	//
	//     * stream.memory.z1d.large
	//
	//     *
	// stream.memory.z1d.xlarge
	//
	//     * stream.memory.z1d.2xlarge
	//
	//     *
	// stream.memory.z1d.3xlarge
	//
	//     * stream.memory.z1d.6xlarge
	//
	//     *
	// stream.memory.z1d.12xlarge
	//
	//     * stream.graphics-design.large
	//
	//     *
	// stream.graphics-design.xlarge
	//
	//     * stream.graphics-design.2xlarge
	//
	//     *
	// stream.graphics-design.4xlarge
	//
	//     * stream.graphics-desktop.2xlarge
	//
	//     *
	// stream.graphics.g4dn.xlarge
	//
	//     * stream.graphics.g4dn.2xlarge
	//
	//     *
	// stream.graphics.g4dn.4xlarge
	//
	//     * stream.graphics.g4dn.8xlarge
	//
	//     *
	// stream.graphics.g4dn.12xlarge
	//
	//     * stream.graphics.g4dn.16xlarge
	//
	//     *
	// stream.graphics-pro.4xlarge
	//
	//     * stream.graphics-pro.8xlarge
	//
	//     *
	// stream.graphics-pro.16xlarge
	//
	// This member is required.
	InstanceType *string

	// A unique name for the image builder.
	//
	// This member is required.
	Name *string

	// The list of interface VPC endpoint (interface endpoint) objects. Administrators
	// can connect to the image builder only through the specified endpoints.
	AccessEndpoints []*types.AccessEndpoint

	// The version of the AppStream 2.0 agent to use for this image builder. To use the
	// latest version of the AppStream 2.0 agent, specify [LATEST].
	AppstreamAgentVersion *string

	// The description to display.
	Description *string

	// The image builder name to display.
	DisplayName *string

	// The name of the directory and organizational unit (OU) to use to join the image
	// builder to a Microsoft Active Directory domain.
	DomainJoinInfo *types.DomainJoinInfo

	// Enables or disables default internet access for the image builder.
	EnableDefaultInternetAccess *bool

	// The Amazon Resource Name (ARN) of the IAM role to apply to the image builder. To
	// assume a role, the image builder calls the AWS Security Token Service (STS)
	// AssumeRole API operation and passes the ARN of the role to use. The operation
	// creates a new session with temporary credentials. AppStream 2.0 retrieves the
	// temporary credentials and creates the appstream_machine_role credential profile
	// on the instance. For more information, see Using an IAM Role to Grant
	// Permissions to Applications and Scripts Running on AppStream 2.0 Streaming
	// Instances
	// (https://docs.aws.amazon.com/appstream2/latest/developerguide/using-iam-roles-to-grant-permissions-to-applications-scripts-streaming-instances.html)
	// in the Amazon AppStream 2.0 Administration Guide.
	IamRoleArn *string

	// The ARN of the public, private, or shared image to use.
	ImageArn *string

	// The name of the image used to create the image builder.
	ImageName *string

	// The tags to associate with the image builder. A tag is a key-value pair, and the
	// value is optional. For example, Environment=Test. If you do not specify a value,
	// Environment=. Generally allowed characters are: letters, numbers, and spaces
	// representable in UTF-8, and the following special characters: _ . : / = + \ - @
	// If you do not specify a value, the value is set to an empty string. For more
	// information about tags, see Tagging Your Resources
	// (https://docs.aws.amazon.com/appstream2/latest/developerguide/tagging-basic.html)
	// in the Amazon AppStream 2.0 Administration Guide.
	Tags map[string]*string

	// The VPC configuration for the image builder. You can specify only one subnet.
	VpcConfig *types.VpcConfig
}

type CreateImageBuilderOutput struct {

	// Information about the image builder.
	ImageBuilder *types.ImageBuilder

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateImageBuilderMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpCreateImageBuilder{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpCreateImageBuilder{})
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
	if err := addOpCreateImageBuilderValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opCreateImageBuilder(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opCreateImageBuilder(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appstream",
		OperationName: "CreateImageBuilder",
	}
}
