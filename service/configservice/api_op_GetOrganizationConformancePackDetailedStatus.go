// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns detailed status for each member account within an organization for a
// given organization conformance pack. Only a master account and a delegated
// administrator account can call this API. When calling this API with a delegated
// administrator, you must ensure AWS Organizations ListDelegatedAdministrator
// permissions are added.
func (c *Client) GetOrganizationConformancePackDetailedStatus(ctx context.Context, params *GetOrganizationConformancePackDetailedStatusInput, optFns ...func(*Options)) (*GetOrganizationConformancePackDetailedStatusOutput, error) {
	if params == nil {
		params = &GetOrganizationConformancePackDetailedStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetOrganizationConformancePackDetailedStatus", params, optFns, addOperationGetOrganizationConformancePackDetailedStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetOrganizationConformancePackDetailedStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetOrganizationConformancePackDetailedStatusInput struct {

	// The name of organization conformance pack for which you want status details for
	// member accounts.
	//
	// This member is required.
	OrganizationConformancePackName *string

	// An OrganizationResourceDetailedStatusFilters object.
	Filters *types.OrganizationResourceDetailedStatusFilters

	// The maximum number of OrganizationConformancePackDetailedStatuses returned on
	// each page. If you do not specify a number, AWS Config uses the default. The
	// default is 100.
	Limit *int32

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string
}

type GetOrganizationConformancePackDetailedStatusOutput struct {

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	// A list of OrganizationConformancePackDetailedStatus objects.
	OrganizationConformancePackDetailedStatuses []*types.OrganizationConformancePackDetailedStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetOrganizationConformancePackDetailedStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpGetOrganizationConformancePackDetailedStatus{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpGetOrganizationConformancePackDetailedStatus{})
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
	if err := addOpGetOrganizationConformancePackDetailedStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opGetOrganizationConformancePackDetailedStatus(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opGetOrganizationConformancePackDetailedStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "GetOrganizationConformancePackDetailedStatus",
	}
}
