// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/directoryservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the status of LDAP security for the specified directory.
func (c *Client) DescribeLDAPSSettings(ctx context.Context, params *DescribeLDAPSSettingsInput, optFns ...func(*Options)) (*DescribeLDAPSSettingsOutput, error) {
	if params == nil {
		params = &DescribeLDAPSSettingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeLDAPSSettings", params, optFns, addOperationDescribeLDAPSSettingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeLDAPSSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeLDAPSSettingsInput struct {

	// The identifier of the directory.
	//
	// This member is required.
	DirectoryId *string

	// Specifies the number of items that should be displayed on one page.
	Limit *int32

	// The type of next token used for pagination.
	NextToken *string

	// The type of LDAP security to enable. Currently only the value Client is
	// supported.
	Type types.LDAPSType
}

type DescribeLDAPSSettingsOutput struct {

	// Information about LDAP security for the specified directory, including status of
	// enablement, state last updated date time, and the reason for the state.
	LDAPSSettingsInfo []*types.LDAPSSettingInfo

	// The next token used to retrieve the LDAPS settings if the number of setting
	// types exceeds page limit and there is another page.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeLDAPSSettingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpDescribeLDAPSSettings{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpDescribeLDAPSSettings{})
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
	if err := addOpDescribeLDAPSSettingsValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDescribeLDAPSSettings(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDescribeLDAPSSettings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "DescribeLDAPSSettings",
	}
}
