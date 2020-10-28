// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Starts the replication task assessment for unsupported data types in the source
// database.
func (c *Client) StartReplicationTaskAssessment(ctx context.Context, params *StartReplicationTaskAssessmentInput, optFns ...func(*Options)) (*StartReplicationTaskAssessmentOutput, error) {
	if params == nil {
		params = &StartReplicationTaskAssessmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartReplicationTaskAssessment", params, optFns, addOperationStartReplicationTaskAssessmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartReplicationTaskAssessmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type StartReplicationTaskAssessmentInput struct {

	// The Amazon Resource Name (ARN) of the replication task.
	//
	// This member is required.
	ReplicationTaskArn *string
}

//
type StartReplicationTaskAssessmentOutput struct {

	// The assessed replication task.
	ReplicationTask *types.ReplicationTask

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStartReplicationTaskAssessmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpStartReplicationTaskAssessment{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpStartReplicationTaskAssessment{})
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
	if err := addOpStartReplicationTaskAssessmentValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opStartReplicationTaskAssessment(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opStartReplicationTaskAssessment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dms",
		OperationName: "StartReplicationTaskAssessment",
	}
}
