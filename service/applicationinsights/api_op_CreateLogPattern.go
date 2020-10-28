// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationinsights

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/applicationinsights/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds an log pattern to a LogPatternSet.
func (c *Client) CreateLogPattern(ctx context.Context, params *CreateLogPatternInput, optFns ...func(*Options)) (*CreateLogPatternOutput, error) {
	if params == nil {
		params = &CreateLogPatternInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLogPattern", params, optFns, addOperationCreateLogPatternMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLogPatternOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLogPatternInput struct {

	// The log pattern.
	//
	// This member is required.
	Pattern *string

	// The name of the log pattern.
	//
	// This member is required.
	PatternName *string

	// The name of the log pattern set.
	//
	// This member is required.
	PatternSetName *string

	// Rank of the log pattern.
	//
	// This member is required.
	Rank *int32

	// The name of the resource group.
	//
	// This member is required.
	ResourceGroupName *string
}

type CreateLogPatternOutput struct {

	// The successfully created log pattern.
	LogPattern *types.LogPattern

	// The name of the resource group.
	ResourceGroupName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateLogPatternMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpCreateLogPattern{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpCreateLogPattern{})
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
	if err := addOpCreateLogPatternValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opCreateLogPattern(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opCreateLogPattern(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "applicationinsights",
		OperationName: "CreateLogPattern",
	}
}
