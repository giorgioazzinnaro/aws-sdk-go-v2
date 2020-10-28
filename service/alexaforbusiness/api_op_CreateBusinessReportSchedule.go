// Code generated by smithy-go-codegen DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/alexaforbusiness/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a recurring schedule for usage reports to deliver to the specified S3
// location with a specified daily or weekly interval.
func (c *Client) CreateBusinessReportSchedule(ctx context.Context, params *CreateBusinessReportScheduleInput, optFns ...func(*Options)) (*CreateBusinessReportScheduleOutput, error) {
	if params == nil {
		params = &CreateBusinessReportScheduleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateBusinessReportSchedule", params, optFns, addOperationCreateBusinessReportScheduleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateBusinessReportScheduleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateBusinessReportScheduleInput struct {

	// The content range of the reports.
	//
	// This member is required.
	ContentRange *types.BusinessReportContentRange

	// The format of the generated report (individual CSV files or zipped files of
	// individual files).
	//
	// This member is required.
	Format types.BusinessReportFormat

	// The client request token.
	ClientRequestToken *string

	// The recurrence of the reports. If this isn't specified, the report will only be
	// delivered one time when the API is called.
	Recurrence *types.BusinessReportRecurrence

	// The S3 bucket name of the output reports. If this isn't specified, the report
	// can be retrieved from a download link by calling ListBusinessReportSchedule.
	S3BucketName *string

	// The S3 key where the report is delivered.
	S3KeyPrefix *string

	// The name identifier of the schedule.
	ScheduleName *string

	// The tags for the business report schedule.
	Tags []*types.Tag
}

type CreateBusinessReportScheduleOutput struct {

	// The ARN of the business report schedule.
	ScheduleArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateBusinessReportScheduleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpCreateBusinessReportSchedule{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpCreateBusinessReportSchedule{})
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
	if err := addIdempotencyToken_opCreateBusinessReportScheduleMiddleware(stack, options); err != nil {
		return err
	}
	if err := addOpCreateBusinessReportScheduleValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opCreateBusinessReportSchedule(options.Region)); err != nil {
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

type idempotencyToken_initializeOpCreateBusinessReportSchedule struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateBusinessReportSchedule) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateBusinessReportSchedule) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateBusinessReportScheduleInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateBusinessReportScheduleInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateBusinessReportScheduleMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(middleware.Before, &idempotencyToken_initializeOpCreateBusinessReportSchedule{tokenProvider: cfg.IdempotencyTokenProvider})
}

func newServiceMetadataMiddleware_opCreateBusinessReportSchedule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "CreateBusinessReportSchedule",
	}
}
