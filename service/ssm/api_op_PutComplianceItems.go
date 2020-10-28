// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Registers a compliance type and other compliance details on a designated
// resource. This action lets you register custom compliance details with a
// resource. This call overwrites existing compliance information on the resource,
// so you must provide a full list of compliance items each time that you send the
// request. ComplianceType can be one of the following:
//
//     * ExecutionId: The
// execution ID when the patch, association, or custom compliance item was
// applied.
//
//     * ExecutionType: Specify patch, association, or Custom:string.
//
//
// * ExecutionTime. The time the patch, association, or custom compliance item was
// applied to the instance.
//
//     * Id: The patch, association, or custom compliance
// ID.
//
//     * Title: A title.
//
//     * Status: The status of the compliance item. For
// example, approved for patches, or Failed for associations.
//
//     * Severity: A
// patch severity. For example, critical.
//
//     * DocumentName: A SSM document name.
// For example, AWS-RunPatchBaseline.
//
//     * DocumentVersion: An SSM document
// version number. For example, 4.
//
//     * Classification: A patch classification.
// For example, security updates.
//
//     * PatchBaselineId: A patch baseline ID.
//
//
// * PatchSeverity: A patch severity. For example, Critical.
//
//     * PatchState: A
// patch state. For example, InstancesWithFailedPatches.
//
//     * PatchGroup: The
// name of a patch group.
//
//     * InstalledTime: The time the association, patch, or
// custom compliance item was applied to the resource. Specify the time by using
// the following format: yyyy-MM-dd'T'HH:mm:ss'Z'
func (c *Client) PutComplianceItems(ctx context.Context, params *PutComplianceItemsInput, optFns ...func(*Options)) (*PutComplianceItemsOutput, error) {
	if params == nil {
		params = &PutComplianceItemsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutComplianceItems", params, optFns, addOperationPutComplianceItemsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutComplianceItemsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutComplianceItemsInput struct {

	// Specify the compliance type. For example, specify Association (for a State
	// Manager association), Patch, or Custom:string.
	//
	// This member is required.
	ComplianceType *string

	// A summary of the call execution that includes an execution ID, the type of
	// execution (for example, Command), and the date/time of the execution using a
	// datetime object that is saved in the following format: yyyy-MM-dd'T'HH:mm:ss'Z'.
	//
	// This member is required.
	ExecutionSummary *types.ComplianceExecutionSummary

	// Information about the compliance as defined by the resource type. For example,
	// for a patch compliance type, Items includes information about the PatchSeverity,
	// Classification, and so on.
	//
	// This member is required.
	Items []*types.ComplianceItemEntry

	// Specify an ID for this resource. For a managed instance, this is the instance
	// ID.
	//
	// This member is required.
	ResourceId *string

	// Specify the type of resource. ManagedInstance is currently the only supported
	// resource type.
	//
	// This member is required.
	ResourceType *string

	// MD5 or SHA-256 content hash. The content hash is used to determine if existing
	// information should be overwritten or ignored. If the content hashes match, the
	// request to put compliance information is ignored.
	ItemContentHash *string

	// The mode for uploading compliance items. You can specify COMPLETE or PARTIAL. In
	// COMPLETE mode, the system overwrites all existing compliance information for the
	// resource. You must provide a full list of compliance items each time you send
	// the request. In PARTIAL mode, the system overwrites compliance information for a
	// specific association. The association must be configured with SyncCompliance set
	// to MANUAL. By default, all requests use COMPLETE mode. This attribute is only
	// valid for association compliance.
	UploadType types.ComplianceUploadType
}

type PutComplianceItemsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutComplianceItemsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpPutComplianceItems{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpPutComplianceItems{})
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
	if err := addOpPutComplianceItemsValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opPutComplianceItems(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opPutComplianceItems(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "PutComplianceItems",
	}
}
