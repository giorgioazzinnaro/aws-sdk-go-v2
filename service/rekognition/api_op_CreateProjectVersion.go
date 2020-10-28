// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new version of a model and begins training. Models are managed as part
// of an Amazon Rekognition Custom Labels project. You can specify one training
// dataset and one testing dataset. The response from CreateProjectVersion is an
// Amazon Resource Name (ARN) for the version of the model. Training takes a while
// to complete. You can get the current status by calling DescribeProjectVersions.
// Once training has successfully completed, call DescribeProjectVersions to get
// the training results and evaluate the model. After evaluating the model, you
// start the model by calling StartProjectVersion. This operation requires
// permissions to perform the rekognition:CreateProjectVersion action.
func (c *Client) CreateProjectVersion(ctx context.Context, params *CreateProjectVersionInput, optFns ...func(*Options)) (*CreateProjectVersionOutput, error) {
	if params == nil {
		params = &CreateProjectVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateProjectVersion", params, optFns, addOperationCreateProjectVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateProjectVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateProjectVersionInput struct {

	// The Amazon S3 location to store the results of training.
	//
	// This member is required.
	OutputConfig *types.OutputConfig

	// The ARN of the Amazon Rekognition Custom Labels project that manages the model
	// that you want to train.
	//
	// This member is required.
	ProjectArn *string

	// The dataset to use for testing.
	//
	// This member is required.
	TestingData *types.TestingData

	// The dataset to use for training.
	//
	// This member is required.
	TrainingData *types.TrainingData

	// A name for the version of the model. This value must be unique.
	//
	// This member is required.
	VersionName *string
}

type CreateProjectVersionOutput struct {

	// The ARN of the model version that was created. Use DescribeProjectVersion to get
	// the current status of the training operation.
	ProjectVersionArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateProjectVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpCreateProjectVersion{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpCreateProjectVersion{})
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
	if err := addOpCreateProjectVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opCreateProjectVersion(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opCreateProjectVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "CreateProjectVersion",
	}
}
