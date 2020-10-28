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

// Assigns new properties to a user. Parameters you pass modify any or all of the
// following: the home directory, role, and policy for the UserName and ServerId
// you specify. The response returns the ServerId and the UserName for the updated
// user.
func (c *Client) UpdateUser(ctx context.Context, params *UpdateUserInput, optFns ...func(*Options)) (*UpdateUserOutput, error) {
	if params == nil {
		params = &UpdateUserInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateUser", params, optFns, addOperationUpdateUserMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateUserInput struct {

	// A system-assigned unique identifier for a server instance that the user account
	// is assigned to.
	//
	// This member is required.
	ServerId *string

	// A unique string that identifies a user and is associated with a server as
	// specified by the ServerId. This user name must be a minimum of 3 and a maximum
	// of 100 characters long. The following are valid characters: a-z, A-Z, 0-9,
	// underscore '_', hyphen '-', period '.', and at sign '@'. The user name can't
	// start with a hyphen, period, or at sign.
	//
	// This member is required.
	UserName *string

	// Specifies the landing directory (folder) for a user when they log in to the
	// server using their file transfer protocol client. An example is
	// your-Amazon-S3-bucket-name>/home/username.
	HomeDirectory *string

	// Logical directory mappings that specify what Amazon S3 paths and keys should be
	// visible to your user and how you want to make them visible. You will need to
	// specify the "Entry" and "Target" pair, where Entry shows how the path is made
	// visible and Target is the actual Amazon S3 path. If you only specify a target,
	// it will be displayed as is. You will need to also make sure that your IAM role
	// provides access to paths in Target. The following is an example. '[
	// "/bucket2/documentation", { "Entry": "your-personal-report.pdf", "Target":
	// "/bucket3/customized-reports/${transfer:UserName}.pdf" } ]' In most cases, you
	// can use this value instead of the scope-down policy to lock your user down to
	// the designated home directory ("chroot"). To do this, you can set Entry to '/'
	// and set Target to the HomeDirectory parameter value. If the target of a logical
	// directory entry does not exist in Amazon S3, the entry will be ignored. As a
	// workaround, you can use the Amazon S3 API to create 0 byte objects as place
	// holders for your directory. If using the CLI, use the s3api call instead of s3
	// so you can use the put-object operation. For example, you use the following: aws
	// s3api put-object --bucket bucketname --key path/to/folder/. Make sure that the
	// end of the key name ends in a / for it to be considered a folder.
	HomeDirectoryMappings []*types.HomeDirectoryMapEntry

	// The type of landing directory (folder) you want your users' home directory to be
	// when they log into the server. If you set it to PATH, the user will see the
	// absolute Amazon S3 bucket paths as is in their file transfer protocol clients.
	// If you set it LOGICAL, you will need to provide mappings in the
	// HomeDirectoryMappings for how you want to make Amazon S3 paths visible to your
	// users.
	HomeDirectoryType types.HomeDirectoryType

	// Allows you to supply a scope-down policy for your user so you can use the same
	// IAM role across multiple users. The policy scopes down user access to portions
	// of your Amazon S3 bucket. Variables you can use inside this policy include
	// ${Transfer:UserName}, ${Transfer:HomeDirectory}, and ${Transfer:HomeBucket}. For
	// scope-down policies, AWS Transfer Family stores the policy as a JSON blob,
	// instead of the Amazon Resource Name (ARN) of the policy. You save the policy as
	// a JSON blob and pass it in the Policy argument. For an example of a scope-down
	// policy, see Creating a scope-down policy
	// (https://docs.aws.amazon.com/transfer/latest/userguide/users.html#users-policies-scope-down).
	// For more information, see AssumeRole
	// (https://docs.aws.amazon.com/STS/latest/APIReference/API_AssumeRole.html) in the
	// AWS Security Token Service API Reference.
	Policy *string

	// The IAM role that controls your users' access to your Amazon S3 bucket. The
	// policies attached to this role will determine the level of access you want to
	// provide your users when transferring files into and out of your Amazon S3 bucket
	// or buckets. The IAM role should also contain a trust relationship that allows
	// the server to access your resources when servicing your users' transfer
	// requests.
	Role *string
}

// UpdateUserResponse returns the user name and identifier for the request to
// update a user's properties.
type UpdateUserOutput struct {

	// A system-assigned unique identifier for a server instance that the user account
	// is assigned to.
	//
	// This member is required.
	ServerId *string

	// The unique identifier for a user that is assigned to a server instance that was
	// specified in the request.
	//
	// This member is required.
	UserName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateUserMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpUpdateUser{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpUpdateUser{})
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
	if err := addOpUpdateUserValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opUpdateUser(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opUpdateUser(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transfer",
		OperationName: "UpdateUser",
	}
}
