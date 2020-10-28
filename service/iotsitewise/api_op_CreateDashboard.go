// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a dashboard in an AWS IoT SiteWise Monitor project.
func (c *Client) CreateDashboard(ctx context.Context, params *CreateDashboardInput, optFns ...func(*Options)) (*CreateDashboardOutput, error) {
	if params == nil {
		params = &CreateDashboardInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDashboard", params, optFns, addOperationCreateDashboardMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDashboardOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDashboardInput struct {

	// The dashboard definition specified in a JSON literal. For detailed information,
	// see Creating dashboards (CLI)
	// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/create-dashboards-using-aws-cli.html)
	// in the AWS IoT SiteWise User Guide.
	//
	// This member is required.
	DashboardDefinition *string

	// A friendly name for the dashboard.
	//
	// This member is required.
	DashboardName *string

	// The ID of the project in which to create the dashboard.
	//
	// This member is required.
	ProjectId *string

	// A unique case-sensitive identifier that you can provide to ensure the
	// idempotency of the request. Don't reuse this client token if a new idempotent
	// request is required.
	ClientToken *string

	// A description for the dashboard.
	DashboardDescription *string

	// A list of key-value pairs that contain metadata for the dashboard. For more
	// information, see Tagging your AWS IoT SiteWise resources
	// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/tag-resources.html)
	// in the AWS IoT SiteWise User Guide.
	Tags map[string]*string
}

type CreateDashboardOutput struct {

	// The ARN
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of
	// the dashboard, which has the following format.
	// arn:${Partition}:iotsitewise:${Region}:${Account}:dashboard/${DashboardId}
	//
	// This member is required.
	DashboardArn *string

	// The ID of the dashboard.
	//
	// This member is required.
	DashboardId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateDashboardMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpCreateDashboard{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpCreateDashboard{})
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
	if err := addEndpointPrefix_opCreateDashboardMiddleware(stack); err != nil {
		return err
	}
	if err := addIdempotencyToken_opCreateDashboardMiddleware(stack, options); err != nil {
		return err
	}
	if err := addOpCreateDashboardValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opCreateDashboard(options.Region)); err != nil {
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

type endpointPrefix_opCreateDashboardMiddleware struct {
}

func (*endpointPrefix_opCreateDashboardMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opCreateDashboardMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.HostPrefix = "monitor."

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opCreateDashboardMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert("OperationSerializer", middleware.Before, &endpointPrefix_opCreateDashboardMiddleware{})
}

type idempotencyToken_initializeOpCreateDashboard struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateDashboard) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateDashboard) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateDashboardInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateDashboardInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateDashboardMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(middleware.Before, &idempotencyToken_initializeOpCreateDashboard{tokenProvider: cfg.IdempotencyTokenProvider})
}

func newServiceMetadataMiddleware_opCreateDashboard(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotsitewise",
		OperationName: "CreateDashboard",
	}
}
