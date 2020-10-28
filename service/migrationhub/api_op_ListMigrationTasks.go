// Code generated by smithy-go-codegen DO NOT EDIT.

package migrationhub

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/migrationhub/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists all, or filtered by resource name, migration tasks associated with the
// user account making this call. This API has the following traits:
//
//     * Can
// show a summary list of the most recent migration tasks.
//
//     * Can show a
// summary list of migration tasks associated with a given discovered resource.
//
//
// * Lists migration tasks in a paginated interface.
func (c *Client) ListMigrationTasks(ctx context.Context, params *ListMigrationTasksInput, optFns ...func(*Options)) (*ListMigrationTasksOutput, error) {
	if params == nil {
		params = &ListMigrationTasksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListMigrationTasks", params, optFns, addOperationListMigrationTasksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListMigrationTasksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListMigrationTasksInput struct {

	// Value to specify how many results are returned per page.
	MaxResults *int32

	// If a NextToken was returned by a previous call, there are more results
	// available. To retrieve the next page of results, make the call again using the
	// returned token in NextToken.
	NextToken *string

	// Filter migration tasks by discovered resource name.
	ResourceName *string
}

type ListMigrationTasksOutput struct {

	// Lists the migration task's summary which includes: MigrationTaskName,
	// ProgressPercent, ProgressUpdateStream, Status, and the UpdateDateTime for each
	// task.
	MigrationTaskSummaryList []*types.MigrationTaskSummary

	// If there are more migration tasks than the max result, return the next token to
	// be passed to the next call as a bookmark of where to start from.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListMigrationTasksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpListMigrationTasks{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpListMigrationTasks{})
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
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opListMigrationTasks(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opListMigrationTasks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mgh",
		OperationName: "ListMigrationTasks",
	}
}
