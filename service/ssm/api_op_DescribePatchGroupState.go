// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns high-level aggregated patch compliance state for a patch group.
func (c *Client) DescribePatchGroupState(ctx context.Context, params *DescribePatchGroupStateInput, optFns ...func(*Options)) (*DescribePatchGroupStateOutput, error) {
	if params == nil {
		params = &DescribePatchGroupStateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribePatchGroupState", params, optFns, addOperationDescribePatchGroupStateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribePatchGroupStateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribePatchGroupStateInput struct {

	// The name of the patch group whose patch snapshot should be retrieved.
	//
	// This member is required.
	PatchGroup *string
}

type DescribePatchGroupStateOutput struct {

	// The number of instances in the patch group.
	Instances *int32

	// The number of instances with patches from the patch baseline that failed to
	// install.
	InstancesWithFailedPatches *int32

	// The number of instances with patches installed that aren't defined in the patch
	// baseline.
	InstancesWithInstalledOtherPatches *int32

	// The number of instances with installed patches.
	InstancesWithInstalledPatches *int32

	// The number of instances with patches installed by Patch Manager that have not
	// been rebooted after the patch installation. The status of these instances is
	// NON_COMPLIANT.
	InstancesWithInstalledPendingRebootPatches *int32

	// The number of instances with patches installed that are specified in a
	// RejectedPatches list. Patches with a status of INSTALLED_REJECTED were typically
	// installed before they were added to a RejectedPatches list. If
	// ALLOW_AS_DEPENDENCY is the specified option for RejectedPatchesAction, the value
	// of InstancesWithInstalledRejectedPatches will always be 0 (zero).
	InstancesWithInstalledRejectedPatches *int32

	// The number of instances with missing patches from the patch baseline.
	InstancesWithMissingPatches *int32

	// The number of instances with patches that aren't applicable.
	InstancesWithNotApplicablePatches *int32

	// The number of instances with NotApplicable patches beyond the supported limit,
	// which are not reported by name to Systems Manager Inventory.
	InstancesWithUnreportedNotApplicablePatches *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribePatchGroupStateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpDescribePatchGroupState{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpDescribePatchGroupState{})
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
	if err := addOpDescribePatchGroupStateValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDescribePatchGroupState(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDescribePatchGroupState(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "DescribePatchGroupState",
	}
}
