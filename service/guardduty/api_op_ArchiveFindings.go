// Code generated by smithy-go-codegen DO NOT EDIT.

package guardduty

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Archives GuardDuty findings that are specified by the list of finding IDs. Only
// the master account can archive findings. Member accounts don't have permission
// to archive findings from their accounts.
func (c *Client) ArchiveFindings(ctx context.Context, params *ArchiveFindingsInput, optFns ...func(*Options)) (*ArchiveFindingsOutput, error) {
	if params == nil {
		params = &ArchiveFindingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ArchiveFindings", params, optFns, addOperationArchiveFindingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ArchiveFindingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ArchiveFindingsInput struct {

	// The ID of the detector that specifies the GuardDuty service whose findings you
	// want to archive.
	//
	// This member is required.
	DetectorId *string

	// The IDs of the findings that you want to archive.
	//
	// This member is required.
	FindingIds []*string
}

type ArchiveFindingsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationArchiveFindingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpArchiveFindings{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpArchiveFindings{})
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
	if err := addOpArchiveFindingsValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opArchiveFindings(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opArchiveFindings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "guardduty",
		OperationName: "ArchiveFindings",
	}
}
