// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the structure of an approval rule created specifically for a pull
// request. For example, you can change the number of required approvers and the
// approval pool for approvers.
func (c *Client) UpdatePullRequestApprovalRuleContent(ctx context.Context, params *UpdatePullRequestApprovalRuleContentInput, optFns ...func(*Options)) (*UpdatePullRequestApprovalRuleContentOutput, error) {
	if params == nil {
		params = &UpdatePullRequestApprovalRuleContentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdatePullRequestApprovalRuleContent", params, optFns, addOperationUpdatePullRequestApprovalRuleContentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdatePullRequestApprovalRuleContentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdatePullRequestApprovalRuleContentInput struct {

	// The name of the approval rule you want to update.
	//
	// This member is required.
	ApprovalRuleName *string

	// The updated content for the approval rule. When you update the content of the
	// approval rule, you can specify approvers in an approval pool in one of two
	// ways:
	//
	//     * CodeCommitApprovers: This option only requires an AWS account and a
	// resource. It can be used for both IAM users and federated access users whose
	// name matches the provided resource name. This is a very powerful option that
	// offers a great deal of flexibility. For example, if you specify the AWS account
	// 123456789012 and Mary_Major, all of the following are counted as approvals
	// coming from that user:
	//
	//         * An IAM user in the account
	// (arn:aws:iam::123456789012:user/Mary_Major)
	//
	//         * A federated user
	// identified in IAM as Mary_Major
	// (arn:aws:sts::123456789012:federated-user/Mary_Major)
	//
	//     This option does not
	// recognize an active session of someone assuming the role of CodeCommitReview
	// with a role session name of Mary_Major
	// (arn:aws:sts::123456789012:assumed-role/CodeCommitReview/Mary_Major) unless you
	// include a wildcard (*Mary_Major).
	//
	//     * Fully qualified ARN: This option allows
	// you to specify the fully qualified Amazon Resource Name (ARN) of the IAM user or
	// role.
	//
	// For more information about IAM ARNs, wildcards, and formats, see IAM
	// Identifiers
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html) in
	// the IAM User Guide.
	//
	// This member is required.
	NewRuleContent *string

	// The system-generated ID of the pull request.
	//
	// This member is required.
	PullRequestId *string

	// The SHA-256 hash signature for the content of the approval rule. You can
	// retrieve this information by using GetPullRequest.
	ExistingRuleContentSha256 *string
}

type UpdatePullRequestApprovalRuleContentOutput struct {

	// Information about the updated approval rule.
	//
	// This member is required.
	ApprovalRule *types.ApprovalRule

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdatePullRequestApprovalRuleContentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpUpdatePullRequestApprovalRuleContent{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpUpdatePullRequestApprovalRuleContent{})
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
	if err := addOpUpdatePullRequestApprovalRuleContentValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opUpdatePullRequestApprovalRuleContent(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opUpdatePullRequestApprovalRuleContent(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "UpdatePullRequestApprovalRuleContent",
	}
}
