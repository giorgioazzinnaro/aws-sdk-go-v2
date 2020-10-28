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

// Lists member users in a group.
func (c *Client) ListGroupMemberships(ctx context.Context, params *ListGroupMembershipsInput, optFns ...func(*Options)) (*ListGroupMembershipsOutput, error) {
	if params == nil {
		params = &ListGroupMembershipsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListGroupMemberships", params, optFns, addOperationListGroupMembershipsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListGroupMembershipsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListGroupMembershipsInput struct {

	// The ID for the AWS account that the group is in. Currently, you use the ID for
	// the AWS account that contains your Amazon QuickSight account.
	//
	// This member is required.
	AwsAccountId *string

	// The name of the group that you want to see a membership list of.
	//
	// This member is required.
	GroupName *string

	// The namespace. Currently, you should set this to default.
	//
	// This member is required.
	Namespace *string

	// The maximum number of results to return from this request.
	MaxResults *int32

	// A pagination token that can be used in a subsequent request.
	NextToken *string
}

type ListGroupMembershipsOutput struct {

	// The list of the members of the group.
	GroupMemberList []*types.GroupMember

	// A pagination token that can be used in a subsequent request.
	NextToken *string

	// The AWS request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListGroupMembershipsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpListGroupMemberships{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpListGroupMemberships{})
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
	if err := addOpListGroupMembershipsValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opListGroupMemberships(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opListGroupMemberships(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "ListGroupMemberships",
	}
}
