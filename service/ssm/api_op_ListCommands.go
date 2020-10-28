// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the commands requested by users of the AWS account.
func (c *Client) ListCommands(ctx context.Context, params *ListCommandsInput, optFns ...func(*Options)) (*ListCommandsOutput, error) {
	if params == nil {
		params = &ListCommandsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCommands", params, optFns, addOperationListCommandsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCommandsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCommandsInput struct {

	// (Optional) If provided, lists only the specified command.
	CommandId *string

	// (Optional) One or more filters. Use a filter to return a more specific list of
	// results.
	Filters []*types.CommandFilter

	// (Optional) Lists commands issued against this instance ID. You can't specify an
	// instance ID in the same command that you specify Status = Pending. This is
	// because the command has not reached the instance yet.
	InstanceId *string

	// (Optional) The maximum number of items to return for this call. The call also
	// returns a token that you can specify in a subsequent call to get the next set of
	// results.
	MaxResults *int32

	// (Optional) The token for the next set of items to return. (You received this
	// token from a previous call.)
	NextToken *string
}

type ListCommandsOutput struct {

	// (Optional) The list of commands requested by the user.
	Commands []*types.Command

	// (Optional) The token for the next set of items to return. (You received this
	// token from a previous call.)
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListCommandsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpListCommands{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpListCommands{})
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
	if err := addOpListCommandsValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opListCommands(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opListCommands(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "ListCommands",
	}
}
