// Code generated by smithy-go-codegen DO NOT EDIT.

package mturk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The GetFileUploadURL operation generates and returns a temporary URL. You use
// the temporary URL to retrieve a file uploaded by a Worker as an answer to a
// FileUploadAnswer question for a HIT. The temporary URL is generated the instant
// the GetFileUploadURL operation is called, and is valid for 60 seconds. You can
// get a temporary file upload URL any time until the HIT is disposed. After the
// HIT is disposed, any uploaded files are deleted, and cannot be retrieved.
// Pending Deprecation on December 12, 2017. The Answer Specification
//
// structure
// will no longer support the FileUploadAnswer element to be used for the
// QuestionForm data structure. Instead, we recommend that Requesters who want to
// create HITs asking Workers to upload files to use Amazon S3.
func (c *Client) GetFileUploadURL(ctx context.Context, params *GetFileUploadURLInput, optFns ...func(*Options)) (*GetFileUploadURLOutput, error) {
	if params == nil {
		params = &GetFileUploadURLInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetFileUploadURL", params, optFns, addOperationGetFileUploadURLMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetFileUploadURLOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetFileUploadURLInput struct {

	// The ID of the assignment that contains the question with a FileUploadAnswer.
	//
	// This member is required.
	AssignmentId *string

	// The identifier of the question with a FileUploadAnswer, as specified in the
	// QuestionForm of the HIT.
	//
	// This member is required.
	QuestionIdentifier *string
}

type GetFileUploadURLOutput struct {

	// A temporary URL for the file that the Worker uploaded for the answer.
	FileUploadURL *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetFileUploadURLMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetFileUploadURL{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetFileUploadURL{}, middleware.After)
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
	addOpGetFileUploadURLValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetFileUploadURL(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetFileUploadURL(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mturk-requester",
		OperationName: "GetFileUploadURL",
	}
}
