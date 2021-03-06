// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the current status of a change batch request. The status is one of the
// following values:
//
// * PENDING indicates that the changes in this request have not
// propagated to all Amazon Route 53 DNS servers. This is the initial status of all
// change batch requests.
//
// * INSYNC indicates that the changes have propagated to
// all Route 53 DNS servers.
func (c *Client) GetChange(ctx context.Context, params *GetChangeInput, optFns ...func(*Options)) (*GetChangeOutput, error) {
	if params == nil {
		params = &GetChangeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetChange", params, optFns, addOperationGetChangeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetChangeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for a GetChange request.
type GetChangeInput struct {

	// The ID of the change batch request. The value that you specify here is the value
	// that ChangeResourceRecordSets returned in the Id element when you submitted the
	// request.
	//
	// This member is required.
	Id *string
}

// A complex type that contains the ChangeInfo element.
type GetChangeOutput struct {

	// A complex type that contains information about the specified change batch.
	//
	// This member is required.
	ChangeInfo *types.ChangeInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetChangeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetChange{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetChange{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpGetChangeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetChange(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	addSanitizeURLMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetChange(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "GetChange",
	}
}
