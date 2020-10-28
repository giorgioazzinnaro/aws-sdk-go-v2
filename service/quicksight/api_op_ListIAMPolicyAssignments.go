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

// Lists IAM policy assignments in the current Amazon QuickSight account.
func (c *Client) ListIAMPolicyAssignments(ctx context.Context, params *ListIAMPolicyAssignmentsInput, optFns ...func(*Options)) (*ListIAMPolicyAssignmentsOutput, error) {
	if params == nil {
		params = &ListIAMPolicyAssignmentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListIAMPolicyAssignments", params, optFns, addOperationListIAMPolicyAssignmentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListIAMPolicyAssignmentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListIAMPolicyAssignmentsInput struct {

	// The ID of the AWS account that contains these IAM policy assignments.
	//
	// This member is required.
	AwsAccountId *string

	// The namespace for the assignments.
	//
	// This member is required.
	Namespace *string

	// The status of the assignments.
	AssignmentStatus types.AssignmentStatus

	// The maximum number of results to be returned per request.
	MaxResults *int32

	// The token for the next set of results, or null if there are no more results.
	NextToken *string
}

type ListIAMPolicyAssignmentsOutput struct {

	// Information describing the IAM policy assignments.
	IAMPolicyAssignments []*types.IAMPolicyAssignmentSummary

	// The token for the next set of results, or null if there are no more results.
	NextToken *string

	// The AWS request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListIAMPolicyAssignmentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpListIAMPolicyAssignments{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpListIAMPolicyAssignments{})
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
	if err := addOpListIAMPolicyAssignmentsValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opListIAMPolicyAssignments(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opListIAMPolicyAssignments(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "ListIAMPolicyAssignments",
	}
}
