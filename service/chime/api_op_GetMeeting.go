// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the Amazon Chime SDK meeting details for the specified meeting ID. For more
// information about the Amazon Chime SDK, see Using the Amazon Chime SDK
// (https://docs.aws.amazon.com/chime/latest/dg/meetings-sdk.html) in the Amazon
// Chime Developer Guide.
func (c *Client) GetMeeting(ctx context.Context, params *GetMeetingInput, optFns ...func(*Options)) (*GetMeetingOutput, error) {
	if params == nil {
		params = &GetMeetingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMeeting", params, optFns, addOperationGetMeetingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMeetingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMeetingInput struct {

	// The Amazon Chime SDK meeting ID.
	//
	// This member is required.
	MeetingId *string
}

type GetMeetingOutput struct {

	// The Amazon Chime SDK meeting information.
	Meeting *types.Meeting

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetMeetingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpGetMeeting{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpGetMeeting{})
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
	if err := addOpGetMeetingValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opGetMeeting(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opGetMeeting(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "GetMeeting",
	}
}
