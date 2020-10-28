// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves information about when the specified access key was last used. The
// information includes the date and time of last use, along with the AWS service
// and Region that were specified in the last request made with that key.
func (c *Client) GetAccessKeyLastUsed(ctx context.Context, params *GetAccessKeyLastUsedInput, optFns ...func(*Options)) (*GetAccessKeyLastUsedOutput, error) {
	if params == nil {
		params = &GetAccessKeyLastUsedInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAccessKeyLastUsed", params, optFns, addOperationGetAccessKeyLastUsedMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAccessKeyLastUsedOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAccessKeyLastUsedInput struct {

	// The identifier of an access key. This parameter allows (through its regex
	// pattern (http://wikipedia.org/wiki/regex)) a string of characters that can
	// consist of any upper or lowercased letter or digit.
	//
	// This member is required.
	AccessKeyId *string
}

// Contains the response to a successful GetAccessKeyLastUsed request. It is also
// returned as a member of the AccessKeyMetaData structure returned by the
// ListAccessKeys action.
type GetAccessKeyLastUsedOutput struct {

	// Contains information about the last time the access key was used.
	AccessKeyLastUsed *types.AccessKeyLastUsed

	// The name of the AWS IAM user that owns this access key.
	UserName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetAccessKeyLastUsedMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsquery_serializeOpGetAccessKeyLastUsed{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsquery_deserializeOpGetAccessKeyLastUsed{})
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
	if err := addOpGetAccessKeyLastUsedValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opGetAccessKeyLastUsed(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opGetAccessKeyLastUsed(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "GetAccessKeyLastUsed",
	}
}
