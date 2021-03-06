// Code generated by smithy-go-codegen DO NOT EDIT.

package workmail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workmail/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the mailbox export jobs started for the specified organization within the
// last seven days.
func (c *Client) ListMailboxExportJobs(ctx context.Context, params *ListMailboxExportJobsInput, optFns ...func(*Options)) (*ListMailboxExportJobsOutput, error) {
	if params == nil {
		params = &ListMailboxExportJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListMailboxExportJobs", params, optFns, addOperationListMailboxExportJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListMailboxExportJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListMailboxExportJobsInput struct {

	// The organization ID.
	//
	// This member is required.
	OrganizationId *string

	// The maximum number of results to return in a single call.
	MaxResults *int32

	// The token to use to retrieve the next page of results.
	NextToken *string
}

type ListMailboxExportJobsOutput struct {

	// The mailbox export job details.
	Jobs []*types.MailboxExportJob

	// The token to use to retrieve the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListMailboxExportJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListMailboxExportJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListMailboxExportJobs{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpListMailboxExportJobsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListMailboxExportJobs(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListMailboxExportJobs(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workmail",
		OperationName: "ListMailboxExportJobs",
	}
}
