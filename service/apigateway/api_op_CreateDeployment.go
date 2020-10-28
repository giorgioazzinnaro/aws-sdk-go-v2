// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Creates a Deployment resource, which makes a specified RestApi callable over the
// internet.
func (c *Client) CreateDeployment(ctx context.Context, params *CreateDeploymentInput, optFns ...func(*Options)) (*CreateDeploymentOutput, error) {
	if params == nil {
		params = &CreateDeploymentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDeployment", params, optFns, addOperationCreateDeploymentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDeploymentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Requests API Gateway to create a Deployment resource.
type CreateDeploymentInput struct {

	// [Required] The string identifier of the associated RestApi.
	//
	// This member is required.
	RestApiId *string

	// Enables a cache cluster for the Stage resource specified in the input.
	CacheClusterEnabled *bool

	// Specifies the cache cluster size for the Stage resource specified in the input,
	// if a cache cluster is enabled.
	CacheClusterSize types.CacheClusterSize

	// The input configuration for the canary deployment when the deployment is a
	// canary release deployment.
	CanarySettings *types.DeploymentCanarySettings

	// The description for the Deployment resource to create.
	Description *string

	Name *string

	// The description of the Stage resource for the Deployment resource to create.
	StageDescription *string

	// The name of the Stage resource for the Deployment resource to create.
	StageName *string

	Template *bool

	TemplateSkipList []*string

	Title *string

	// Specifies whether active tracing with X-ray is enabled for the Stage.
	TracingEnabled *bool

	// A map that defines the stage variables for the Stage resource that is associated
	// with the new deployment. Variable names can have alphanumeric and underscore
	// characters, and the values must match [A-Za-z0-9-._~:/?#&=,]+.
	Variables map[string]*string
}

// An immutable representation of a RestApi resource that can be called by users
// using Stages. A deployment must be associated with a Stage for it to be callable
// over the Internet. To create a deployment, call POST on the Deployments resource
// of a RestApi. To view, update, or delete a deployment, call GET, PATCH, or
// DELETE on the specified deployment resource
// (/restapis/{restapi_id}/deployments/{deployment_id}). RestApi, Deployments,
// Stage, AWS CLI
// (https://docs.aws.amazon.com/cli/latest/reference/apigateway/get-deployment.html),
// AWS SDKs (https://aws.amazon.com/tools/)
type CreateDeploymentOutput struct {

	// A summary of the RestApi at the date and time that the deployment resource was
	// created.
	ApiSummary map[string]map[string]*types.MethodSnapshot

	// The date and time that the deployment resource was created.
	CreatedDate *time.Time

	// The description for the deployment resource.
	Description *string

	// The identifier for the deployment resource.
	Id *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateDeploymentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpCreateDeployment{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpCreateDeployment{})
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
	if err := addOpCreateDeploymentValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opCreateDeployment(options.Region)); err != nil {
		return err
	}
	if err := addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err := addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err := addAcceptHeader(stack); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateDeployment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "CreateDeployment",
	}
}
