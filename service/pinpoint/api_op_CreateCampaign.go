// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpoint

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new campaign for an application or updates the settings of an existing
// campaign for an application.
func (c *Client) CreateCampaign(ctx context.Context, params *CreateCampaignInput, optFns ...func(*Options)) (*CreateCampaignOutput, error) {
	if params == nil {
		params = &CreateCampaignInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCampaign", params, optFns, addOperationCreateCampaignMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCampaignOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateCampaignInput struct {

	// The unique identifier for the application. This identifier is displayed as the
	// Project ID on the Amazon Pinpoint console.
	//
	// This member is required.
	ApplicationId *string

	// Specifies the configuration and other settings for a campaign.
	//
	// This member is required.
	WriteCampaignRequest *types.WriteCampaignRequest
}

type CreateCampaignOutput struct {

	// Provides information about the status, configuration, and other settings for a
	// campaign.
	//
	// This member is required.
	CampaignResponse *types.CampaignResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateCampaignMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateCampaign{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateCampaign{}, middleware.After)
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
	addOpCreateCampaignValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCampaign(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateCampaign(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mobiletargeting",
		OperationName: "CreateCampaign",
	}
}
