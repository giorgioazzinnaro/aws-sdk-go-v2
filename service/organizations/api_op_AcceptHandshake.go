// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Sends a response to the originator of a handshake agreeing to the action
// proposed by the handshake request. This operation can be called only by the
// following principals when they also have the relevant IAM permissions:
//
//     *
// Invitation to join or Approve all features request handshakes: only a principal
// from the member account. The user who calls the API for an invitation to join
// must have the organizations:AcceptHandshake permission. If you enabled all
// features in the organization, the user must also have the
// iam:CreateServiceLinkedRole permission so that AWS Organizations can create the
// required service-linked role named AWSServiceRoleForOrganizations. For more
// information, see AWS Organizations and Service-Linked Roles
// (http://docs.aws.amazon.com/organizations/latest/userguide/orgs_integration_services.html#orgs_integration_service-linked-roles)
// in the AWS Organizations User Guide.
//
//     * Enable all features final
// confirmation handshake: only a principal from the management account. For more
// information about invitations, see Inviting an AWS Account to Join Your
// Organization
// (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_accounts_invites.html)
// in the AWS Organizations User Guide. For more information about requests to
// enable all features in the organization, see Enabling All Features in Your
// Organization
// (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_org_support-all-features.html)
// in the AWS Organizations User Guide.
//
// After you accept a handshake, it continues
// to appear in the results of relevant APIs for only 30 days. After that, it's
// deleted.
func (c *Client) AcceptHandshake(ctx context.Context, params *AcceptHandshakeInput, optFns ...func(*Options)) (*AcceptHandshakeOutput, error) {
	if params == nil {
		params = &AcceptHandshakeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AcceptHandshake", params, optFns, addOperationAcceptHandshakeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AcceptHandshakeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AcceptHandshakeInput struct {

	// The unique identifier (ID) of the handshake that you want to accept. The regex
	// pattern (http://wikipedia.org/wiki/regex) for handshake ID string requires "h-"
	// followed by from 8 to 32 lowercase letters or digits.
	//
	// This member is required.
	HandshakeId *string
}

type AcceptHandshakeOutput struct {

	// A structure that contains details about the accepted handshake.
	Handshake *types.Handshake

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAcceptHandshakeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpAcceptHandshake{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpAcceptHandshake{})
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
	if err := addOpAcceptHandshakeValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opAcceptHandshake(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opAcceptHandshake(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "organizations",
		OperationName: "AcceptHandshake",
	}
}
