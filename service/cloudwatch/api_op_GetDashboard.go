// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Displays the details of the dashboard that you specify. To copy an existing
// dashboard, use GetDashboard, and then use the data returned within DashboardBody
// as the template for the new dashboard when you call PutDashboard to create the
// copy.
func (c *Client) GetDashboard(ctx context.Context, params *GetDashboardInput, optFns ...func(*Options)) (*GetDashboardOutput, error) {
	if params == nil {
		params = &GetDashboardInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDashboard", params, optFns, addOperationGetDashboardMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDashboardOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDashboardInput struct {

	// The name of the dashboard to be described.
	//
	// This member is required.
	DashboardName *string
}

type GetDashboardOutput struct {

	// The Amazon Resource Name (ARN) of the dashboard.
	DashboardArn *string

	// The detailed information about the dashboard, including what widgets are
	// included and their location on the dashboard. For more information about the
	// DashboardBody syntax, see Dashboard Body Structure and Syntax
	// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/CloudWatch-Dashboard-Body-Structure.html).
	DashboardBody *string

	// The name of the dashboard.
	DashboardName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetDashboardMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsquery_serializeOpGetDashboard{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsquery_deserializeOpGetDashboard{})
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
	if err := addOpGetDashboardValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opGetDashboard(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opGetDashboard(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "monitoring",
		OperationName: "GetDashboard",
	}
}
