// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Create a configuration set. Configuration sets are groups of rules that you can
// apply to the emails that you send. You apply a configuration set to an email by
// specifying the name of the configuration set when you call the Amazon SES API
// v2. When you apply a configuration set to an email, all of the rules in that
// configuration set are applied to the email.
func (c *Client) CreateConfigurationSet(ctx context.Context, params *CreateConfigurationSetInput, optFns ...func(*Options)) (*CreateConfigurationSetOutput, error) {
	if params == nil {
		params = &CreateConfigurationSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateConfigurationSet", params, optFns, addOperationCreateConfigurationSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateConfigurationSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to create a configuration set.
type CreateConfigurationSetInput struct {

	// The name of the configuration set.
	//
	// This member is required.
	ConfigurationSetName *string

	// An object that defines the dedicated IP pool that is used to send emails that
	// you send using the configuration set.
	DeliveryOptions *types.DeliveryOptions

	// An object that defines whether or not Amazon SES collects reputation metrics for
	// the emails that you send that use the configuration set.
	ReputationOptions *types.ReputationOptions

	// An object that defines whether or not Amazon SES can send email that you send
	// using the configuration set.
	SendingOptions *types.SendingOptions

	// An object that contains information about the suppression list preferences for
	// your account.
	SuppressionOptions *types.SuppressionOptions

	// An array of objects that define the tags (keys and values) that you want to
	// associate with the configuration set.
	Tags []*types.Tag

	// An object that defines the open and click tracking options for emails that you
	// send using the configuration set.
	TrackingOptions *types.TrackingOptions
}

// An HTTP 200 response if the request succeeds, or an error message if the request
// fails.
type CreateConfigurationSetOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateConfigurationSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpCreateConfigurationSet{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpCreateConfigurationSet{})
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
	if err := addOpCreateConfigurationSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opCreateConfigurationSet(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opCreateConfigurationSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "CreateConfigurationSet",
	}
}
