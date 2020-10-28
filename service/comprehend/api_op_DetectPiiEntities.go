// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehend

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/comprehend/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Inspects the input text for entities that contain personally identifiable
// information (PII) and returns information about them.
func (c *Client) DetectPiiEntities(ctx context.Context, params *DetectPiiEntitiesInput, optFns ...func(*Options)) (*DetectPiiEntitiesOutput, error) {
	if params == nil {
		params = &DetectPiiEntitiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DetectPiiEntities", params, optFns, addOperationDetectPiiEntitiesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DetectPiiEntitiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DetectPiiEntitiesInput struct {

	// The language of the input documents.
	//
	// This member is required.
	LanguageCode types.LanguageCode

	// A UTF-8 text string. Each string must contain fewer that 5,000 bytes of UTF-8
	// encoded characters.
	//
	// This member is required.
	Text *string
}

type DetectPiiEntitiesOutput struct {

	// A collection of PII entities identified in the input text. For each entity, the
	// response provides the entity type, where the entity text begins and ends, and
	// the level of confidence that Amazon Comprehend has in the detection.
	Entities []*types.PiiEntity

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDetectPiiEntitiesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpDetectPiiEntities{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpDetectPiiEntities{})
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
	if err := addOpDetectPiiEntitiesValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDetectPiiEntities(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDetectPiiEntities(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehend",
		OperationName: "DetectPiiEntities",
	}
}
