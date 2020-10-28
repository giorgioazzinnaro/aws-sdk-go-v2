// Code generated by smithy-go-codegen DO NOT EDIT.

package swf

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/swf/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Used by deciders to get a DecisionTask from the specified decision taskList. A
// decision task may be returned for any open workflow execution that is using the
// specified task list. The task includes a paginated view of the history of the
// workflow execution. The decider should use the workflow type and the history to
// determine how to properly handle the task. This action initiates a long poll,
// where the service holds the HTTP connection open and responds as soon a task
// becomes available. If no decision task is available in the specified task list
// before the timeout of 60 seconds expires, an empty result is returned. An empty
// result, in this context, means that a DecisionTask is returned, but that the
// value of taskToken is an empty string. Deciders should set their client side
// socket timeout to at least 70 seconds (10 seconds higher than the timeout).
// Because the number of workflow history events for a single workflow execution
// might be very large, the result returned might be split up across a number of
// pages. To retrieve subsequent pages, make additional calls to
// PollForDecisionTask using the nextPageToken returned by the initial call. Note
// that you do not call GetWorkflowExecutionHistory with this nextPageToken.
// Instead, call PollForDecisionTask again. Access Control You can use IAM policies
// to control this action's access to Amazon SWF resources as follows:
//
//     * Use a
// Resource element with the domain name to limit the action to only specified
// domains.
//
//     * Use an Action element to allow or deny permission to call this
// action.
//
//     * Constrain the taskList.name parameter by using a Condition
// element with the swf:taskList.name key to allow the action to access only
// certain task lists.
//
// If the caller doesn't have sufficient permissions to invoke
// the action, or the parameter values fall outside the specified constraints, the
// action fails. The associated event attribute's cause parameter is set to
// OPERATION_NOT_PERMITTED. For details and example IAM policies, see Using IAM to
// Manage Access to Amazon SWF Workflows
// (https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html)
// in the Amazon SWF Developer Guide.
func (c *Client) PollForDecisionTask(ctx context.Context, params *PollForDecisionTaskInput, optFns ...func(*Options)) (*PollForDecisionTaskOutput, error) {
	if params == nil {
		params = &PollForDecisionTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PollForDecisionTask", params, optFns, addOperationPollForDecisionTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PollForDecisionTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PollForDecisionTaskInput struct {

	// The name of the domain containing the task lists to poll.
	//
	// This member is required.
	Domain *string

	// Specifies the task list to poll for decision tasks. The specified string must
	// not start or end with whitespace. It must not contain a : (colon), / (slash), |
	// (vertical bar), or any control characters (\u0000-\u001f | \u007f-\u009f). Also,
	// it must not be the literal string arn.
	//
	// This member is required.
	TaskList *types.TaskList

	// Identity of the decider making the request, which is recorded in the
	// DecisionTaskStarted event in the workflow history. This enables diagnostic
	// tracing when problems arise. The form of this identity is user defined.
	Identity *string

	// The maximum number of results that are returned per call. Use nextPageToken to
	// obtain further pages of results. This is an upper limit only; the actual number
	// of results returned per call may be fewer than the specified maximum.
	MaximumPageSize *int32

	// If NextPageToken is returned there are more results available. The value of
	// NextPageToken is a unique pagination token for each page. Make the call again
	// using the returned token to retrieve the next page. Keep all other arguments
	// unchanged. Each pagination token expires after 60 seconds. Using an expired
	// pagination token will return a 400 error: "Specified token has exceeded its
	// maximum lifetime". The configured maximumPageSize determines how many results
	// can be returned in a single call. The nextPageToken returned by this action
	// cannot be used with GetWorkflowExecutionHistory to get the next page. You must
	// call PollForDecisionTask again (with the nextPageToken) to retrieve the next
	// page of history records. Calling PollForDecisionTask with a nextPageToken
	// doesn't return a new decision task.
	NextPageToken *string

	// When set to true, returns the events in reverse order. By default the results
	// are returned in ascending order of the eventTimestamp of the events.
	ReverseOrder *bool
}

// A structure that represents a decision task. Decision tasks are sent to deciders
// in order for them to make decisions.
type PollForDecisionTaskOutput struct {

	// A paginated list of history events of the workflow execution. The decider uses
	// this during the processing of the decision task.
	//
	// This member is required.
	Events []*types.HistoryEvent

	// The ID of the DecisionTaskStarted event recorded in the history.
	//
	// This member is required.
	StartedEventId *int64

	// The opaque string used as a handle on the task. This token is used by workers to
	// communicate progress and response information back to the system about the task.
	//
	// This member is required.
	TaskToken *string

	// The workflow execution for which this decision task was created.
	//
	// This member is required.
	WorkflowExecution *types.WorkflowExecution

	// The type of the workflow execution for which this decision task was created.
	//
	// This member is required.
	WorkflowType *types.WorkflowType

	// If a NextPageToken was returned by a previous call, there are more results
	// available. To retrieve the next page of results, make the call again using the
	// returned token in nextPageToken. Keep all other arguments unchanged. The
	// configured maximumPageSize determines how many results can be returned in a
	// single call.
	NextPageToken *string

	// The ID of the DecisionTaskStarted event of the previous decision task of this
	// workflow execution that was processed by the decider. This can be used to
	// determine the events in the history new since the last decision task received by
	// the decider.
	PreviousStartedEventId *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPollForDecisionTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson10_serializeOpPollForDecisionTask{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson10_deserializeOpPollForDecisionTask{})
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
	if err := addOpPollForDecisionTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opPollForDecisionTask(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opPollForDecisionTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "swf",
		OperationName: "PollForDecisionTask",
	}
}
