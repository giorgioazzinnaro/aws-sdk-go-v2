// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets a list of all grants for the specified customer master key (CMK). To
// perform this operation on a CMK in a different AWS account, specify the key ARN
// in the value of the KeyId parameter. The GranteePrincipal field in the
// ListGrants response usually contains the user or role designated as the grantee
// principal in the grant. However, when the grantee principal in the grant is an
// AWS service, the GranteePrincipal field contains the service principal
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_principal.html#principal-services),
// which might represent several different grantee principals.
func (c *Client) ListGrants(ctx context.Context, params *ListGrantsInput, optFns ...func(*Options)) (*ListGrantsOutput, error) {
	if params == nil {
		params = &ListGrantsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListGrants", params, optFns, addOperationListGrantsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListGrantsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListGrantsInput struct {

	// A unique identifier for the customer master key (CMK). Specify the key ID or the
	// Amazon Resource Name (ARN) of the CMK. To specify a CMK in a different AWS
	// account, you must use the key ARN. For example:
	//
	//     * Key ID:
	// 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//     * Key ARN:
	// arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// To
	// get the key ID and key ARN for a CMK, use ListKeys or DescribeKey.
	//
	// This member is required.
	KeyId *string

	// Use this parameter to specify the maximum number of items to return. When this
	// value is present, AWS KMS does not return more than the specified number of
	// items, but it might return fewer. This value is optional. If you include a
	// value, it must be between 1 and 100, inclusive. If you do not include a value,
	// it defaults to 50.
	Limit *int32

	// Use this parameter in a subsequent request after you receive a response with
	// truncated results. Set it to the value of NextMarker from the truncated response
	// you just received.
	Marker *string
}

type ListGrantsOutput struct {

	// A list of grants.
	Grants []*types.GrantListEntry

	// When Truncated is true, this element is present and contains the value to use
	// for the Marker parameter in a subsequent request.
	NextMarker *string

	// A flag that indicates whether there are more items in the list. When this value
	// is true, the list in this response is truncated. To get more items, pass the
	// value of the NextMarker element in thisresponse to the Marker parameter in a
	// subsequent request.
	Truncated *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListGrantsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpListGrants{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpListGrants{})
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
	if err := addOpListGrantsValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opListGrants(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opListGrants(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "ListGrants",
	}
}
