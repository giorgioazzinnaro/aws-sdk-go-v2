// Code generated by smithy-go-codegen DO NOT EDIT.

package guardduty

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the ThreatIntelSets of the GuardDuty service specified by the detector ID.
// If you use this operation from a member account, the ThreatIntelSets associated
// with the master account are returned.
func (c *Client) ListThreatIntelSets(ctx context.Context, params *ListThreatIntelSetsInput, optFns ...func(*Options)) (*ListThreatIntelSetsOutput, error) {
	if params == nil {
		params = &ListThreatIntelSetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListThreatIntelSets", params, optFns, addOperationListThreatIntelSetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListThreatIntelSetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListThreatIntelSetsInput struct {

	// The unique ID of the detector that the threatIntelSet is associated with.
	//
	// This member is required.
	DetectorId *string

	// You can use this parameter to indicate the maximum number of items that you want
	// in the response. The default value is 50. The maximum value is 50.
	MaxResults *int32

	// You can use this parameter to paginate results in the response. Set the value of
	// this parameter to null on your first call to the list action. For subsequent
	// calls to the action, fill nextToken in the request with the value of NextToken
	// from the previous response to continue listing data.
	NextToken *string
}

type ListThreatIntelSetsOutput struct {

	// The IDs of the ThreatIntelSet resources.
	//
	// This member is required.
	ThreatIntelSetIds []*string

	// The pagination parameter to be used on the next list operation to retrieve more
	// items.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListThreatIntelSetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpListThreatIntelSets{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpListThreatIntelSets{})
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
	if err := addOpListThreatIntelSetsValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opListThreatIntelSets(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opListThreatIntelSets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "guardduty",
		OperationName: "ListThreatIntelSets",
	}
}
