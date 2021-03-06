// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns contact list metadata. It does not return any information about the
// contacts present in the list.
func (c *Client) GetContactList(ctx context.Context, params *GetContactListInput, optFns ...func(*Options)) (*GetContactListOutput, error) {
	if params == nil {
		params = &GetContactListInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetContactList", params, optFns, addOperationGetContactListMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetContactListOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetContactListInput struct {

	// The name of the contact list.
	//
	// This member is required.
	ContactListName *string
}

type GetContactListOutput struct {

	// The name of the contact list.
	ContactListName *string

	// A timestamp noting when the contact list was created.
	CreatedTimestamp *time.Time

	// A description of what the contact list is about.
	Description *string

	// A timestamp noting the last time the contact list was updated.
	LastUpdatedTimestamp *time.Time

	// The tags associated with a contact list.
	Tags []*types.Tag

	// An interest group, theme, or label within a list. A contact list can have
	// multiple topics.
	Topics []*types.Topic

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetContactListMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetContactList{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetContactList{}, middleware.After)
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
	addOpGetContactListValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetContactList(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetContactList(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "GetContactList",
	}
}
