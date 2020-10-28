// Code generated by smithy-go-codegen DO NOT EDIT.

package costandusagereportservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/costandusagereportservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new report using the description that you provide.
func (c *Client) PutReportDefinition(ctx context.Context, params *PutReportDefinitionInput, optFns ...func(*Options)) (*PutReportDefinitionOutput, error) {
	if params == nil {
		params = &PutReportDefinitionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutReportDefinition", params, optFns, addOperationPutReportDefinitionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutReportDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Creates a Cost and Usage Report.
type PutReportDefinitionInput struct {

	// Represents the output of the PutReportDefinition operation. The content consists
	// of the detailed metadata and data file information.
	//
	// This member is required.
	ReportDefinition *types.ReportDefinition
}

// If the action is successful, the service sends back an HTTP 200 response with an
// empty HTTP body.
type PutReportDefinitionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutReportDefinitionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpPutReportDefinition{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpPutReportDefinition{})
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
	if err := addOpPutReportDefinitionValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opPutReportDefinition(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opPutReportDefinition(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cur",
		OperationName: "PutReportDefinition",
	}
}
