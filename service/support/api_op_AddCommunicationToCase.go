// Code generated by smithy-go-codegen DO NOT EDIT.

package support

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds additional customer communication to an AWS Support case. Use the caseId
// parameter to identify the case to which to add communication. You can list a set
// of email addresses to copy on the communication by using the ccEmailAddresses
// parameter. The communicationBody value contains the text of the communication.
//
//
// * You must have a Business or Enterprise support plan to use the AWS Support
// API.
//
//     * If you call the AWS Support API from an account that does not have a
// Business or Enterprise support plan, the SubscriptionRequiredException error
// message appears. For information about changing your support plan, see AWS
// Support (http://aws.amazon.com/premiumsupport/).
func (c *Client) AddCommunicationToCase(ctx context.Context, params *AddCommunicationToCaseInput, optFns ...func(*Options)) (*AddCommunicationToCaseOutput, error) {
	if params == nil {
		params = &AddCommunicationToCaseInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AddCommunicationToCase", params, optFns, addOperationAddCommunicationToCaseMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AddCommunicationToCaseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AddCommunicationToCaseInput struct {

	// The body of an email communication to add to the support case.
	//
	// This member is required.
	CommunicationBody *string

	// The ID of a set of one or more attachments for the communication to add to the
	// case. Create the set by calling AddAttachmentsToSet
	AttachmentSetId *string

	// The AWS Support case ID requested or returned in the call. The case ID is an
	// alphanumeric string formatted as shown in this example:
	// case-12345678910-2013-c4c1d2bf33c5cf47
	CaseId *string

	// The email addresses in the CC line of an email to be added to the support case.
	CcEmailAddresses []*string
}

// The result of the AddCommunicationToCase operation.
type AddCommunicationToCaseOutput struct {

	// True if AddCommunicationToCase succeeds. Otherwise, returns an error.
	Result *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAddCommunicationToCaseMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpAddCommunicationToCase{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpAddCommunicationToCase{})
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
	if err := addOpAddCommunicationToCaseValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opAddCommunicationToCase(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opAddCommunicationToCase(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "support",
		OperationName: "AddCommunicationToCase",
	}
}
