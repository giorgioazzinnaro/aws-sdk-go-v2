// Code generated by smithy-go-codegen DO NOT EDIT.

package codebuild

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codebuild/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// For an existing AWS CodeBuild build project that has its source code stored in a
// GitHub or Bitbucket repository, enables AWS CodeBuild to start rebuilding the
// source code every time a code change is pushed to the repository. If you enable
// webhooks for an AWS CodeBuild project, and the project is used as a build step
// in AWS CodePipeline, then two identical builds are created for each commit. One
// build is triggered through webhooks, and one through AWS CodePipeline. Because
// billing is on a per-build basis, you are billed for both builds. Therefore, if
// you are using AWS CodePipeline, we recommend that you disable webhooks in AWS
// CodeBuild. In the AWS CodeBuild console, clear the Webhook box. For more
// information, see step 5 in Change a Build Project's Settings
// (https://docs.aws.amazon.com/codebuild/latest/userguide/change-project.html#change-project-console).
func (c *Client) CreateWebhook(ctx context.Context, params *CreateWebhookInput, optFns ...func(*Options)) (*CreateWebhookOutput, error) {
	if params == nil {
		params = &CreateWebhookInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateWebhook", params, optFns, addOperationCreateWebhookMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateWebhookOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateWebhookInput struct {

	// The name of the AWS CodeBuild project.
	//
	// This member is required.
	ProjectName *string

	// A regular expression used to determine which repository branches are built when
	// a webhook is triggered. If the name of a branch matches the regular expression,
	// then it is built. If branchFilter is empty, then all branches are built. It is
	// recommended that you use filterGroups instead of branchFilter.
	BranchFilter *string

	// Specifies the type of build this webhook will trigger.
	BuildType types.WebhookBuildType

	// An array of arrays of WebhookFilter objects used to determine which webhooks are
	// triggered. At least one WebhookFilter in the array must specify EVENT as its
	// type. For a build to be triggered, at least one filter group in the filterGroups
	// array must pass. For a filter group to pass, each of its filters must pass.
	FilterGroups [][]*types.WebhookFilter
}

type CreateWebhookOutput struct {

	// Information about a webhook that connects repository events to a build project
	// in AWS CodeBuild.
	Webhook *types.Webhook

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateWebhookMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpCreateWebhook{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpCreateWebhook{})
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
	if err := addOpCreateWebhookValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opCreateWebhook(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opCreateWebhook(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codebuild",
		OperationName: "CreateWebhook",
	}
}
