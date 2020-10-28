// Code generated by smithy-go-codegen DO NOT EDIT.

package sns

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the endpoints and endpoint attributes for devices in a supported push
// notification service, such as GCM (Firebase Cloud Messaging) and APNS. The
// results for ListEndpointsByPlatformApplication are paginated and return a
// limited list of endpoints, up to 100. If additional records are available after
// the first page results, then a NextToken string will be returned. To receive the
// next page, you call ListEndpointsByPlatformApplication again using the NextToken
// string received from the previous call. When there are no more records to
// return, NextToken will be null. For more information, see Using Amazon SNS
// Mobile Push Notifications
// (https://docs.aws.amazon.com/sns/latest/dg/SNSMobilePush.html). This action is
// throttled at 30 transactions per second (TPS).
func (c *Client) ListEndpointsByPlatformApplication(ctx context.Context, params *ListEndpointsByPlatformApplicationInput, optFns ...func(*Options)) (*ListEndpointsByPlatformApplicationOutput, error) {
	if params == nil {
		params = &ListEndpointsByPlatformApplicationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEndpointsByPlatformApplication", params, optFns, addOperationListEndpointsByPlatformApplicationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEndpointsByPlatformApplicationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input for ListEndpointsByPlatformApplication action.
type ListEndpointsByPlatformApplicationInput struct {

	// PlatformApplicationArn for ListEndpointsByPlatformApplicationInput action.
	//
	// This member is required.
	PlatformApplicationArn *string

	// NextToken string is used when calling ListEndpointsByPlatformApplication action
	// to retrieve additional records that are available after the first page results.
	NextToken *string
}

// Response for ListEndpointsByPlatformApplication action.
type ListEndpointsByPlatformApplicationOutput struct {

	// Endpoints returned for ListEndpointsByPlatformApplication action.
	Endpoints []*types.Endpoint

	// NextToken string is returned when calling ListEndpointsByPlatformApplication
	// action if additional records are available after the first page results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListEndpointsByPlatformApplicationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsquery_serializeOpListEndpointsByPlatformApplication{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsquery_deserializeOpListEndpointsByPlatformApplication{})
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
	if err := addOpListEndpointsByPlatformApplicationValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opListEndpointsByPlatformApplication(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opListEndpointsByPlatformApplication(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sns",
		OperationName: "ListEndpointsByPlatformApplication",
	}
}
