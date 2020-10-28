// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an assignment with one specified IAM policy, identified by its Amazon
// Resource Name (ARN). This policy will be assigned to specified groups or users
// of Amazon QuickSight. The users and groups need to be in the same namespace.
func (c *Client) CreateIAMPolicyAssignment(ctx context.Context, params *CreateIAMPolicyAssignmentInput, optFns ...func(*Options)) (*CreateIAMPolicyAssignmentOutput, error) {
	if params == nil {
		params = &CreateIAMPolicyAssignmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateIAMPolicyAssignment", params, optFns, addOperationCreateIAMPolicyAssignmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateIAMPolicyAssignmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateIAMPolicyAssignmentInput struct {

	// The name of the assignment. It must be unique within an AWS account.
	//
	// This member is required.
	AssignmentName *string

	// The status of the assignment. Possible values are as follows:
	//
	//     * ENABLED -
	// Anything specified in this assignment is used when creating the data source.
	//
	//
	// * DISABLED - This assignment isn't used when creating the data source.
	//
	//     *
	// DRAFT - This assignment is an unfinished draft and isn't used when creating the
	// data source.
	//
	// This member is required.
	AssignmentStatus types.AssignmentStatus

	// The ID of the AWS account where you want to assign an IAM policy to QuickSight
	// users or groups.
	//
	// This member is required.
	AwsAccountId *string

	// The namespace that contains the assignment.
	//
	// This member is required.
	Namespace *string

	// The QuickSight users, groups, or both that you want to assign the policy to.
	Identities map[string][]*string

	// The ARN for the IAM policy to apply to the QuickSight users and groups specified
	// in this assignment.
	PolicyArn *string
}

type CreateIAMPolicyAssignmentOutput struct {

	// The ID for the assignment.
	AssignmentId *string

	// The name of the assignment. This name must be unique within the AWS account.
	AssignmentName *string

	// The status of the assignment. Possible values are as follows:
	//
	//     * ENABLED -
	// Anything specified in this assignment is used when creating the data source.
	//
	//
	// * DISABLED - This assignment isn't used when creating the data source.
	//
	//     *
	// DRAFT - This assignment is an unfinished draft and isn't used when creating the
	// data source.
	AssignmentStatus types.AssignmentStatus

	// The QuickSight users, groups, or both that the IAM policy is assigned to.
	Identities map[string][]*string

	// The ARN for the IAM policy that is applied to the QuickSight users and groups
	// specified in this assignment.
	PolicyArn *string

	// The AWS request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateIAMPolicyAssignmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpCreateIAMPolicyAssignment{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpCreateIAMPolicyAssignment{})
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
	if err := addOpCreateIAMPolicyAssignmentValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opCreateIAMPolicyAssignment(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opCreateIAMPolicyAssignment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "CreateIAMPolicyAssignment",
	}
}
