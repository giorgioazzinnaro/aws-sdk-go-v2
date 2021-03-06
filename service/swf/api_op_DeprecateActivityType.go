// Code generated by smithy-go-codegen DO NOT EDIT.

package swf

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/swf/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deprecates the specified activity type. After an activity type has been
// deprecated, you cannot create new tasks of that activity type. Tasks of this
// type that were scheduled before the type was deprecated continue to run. This
// operation is eventually consistent. The results are best effort and may not
// exactly reflect recent updates and changes. Access Control You can use IAM
// policies to control this action's access to Amazon SWF resources as follows:
//
// *
// Use a Resource element with the domain name to limit the action to only
// specified domains.
//
// * Use an Action element to allow or deny permission to call
// this action.
//
// * Constrain the following parameters by using a Condition element
// with the appropriate keys.
//
// * activityType.name: String constraint. The key is
// swf:activityType.name.
//
// * activityType.version: String constraint. The key is
// swf:activityType.version.
//
// If the caller doesn't have sufficient permissions to
// invoke the action, or the parameter values fall outside the specified
// constraints, the action fails. The associated event attribute's cause parameter
// is set to OPERATION_NOT_PERMITTED. For details and example IAM policies, see
// Using IAM to Manage Access to Amazon SWF Workflows
// (https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html)
// in the Amazon SWF Developer Guide.
func (c *Client) DeprecateActivityType(ctx context.Context, params *DeprecateActivityTypeInput, optFns ...func(*Options)) (*DeprecateActivityTypeOutput, error) {
	if params == nil {
		params = &DeprecateActivityTypeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeprecateActivityType", params, optFns, addOperationDeprecateActivityTypeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeprecateActivityTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeprecateActivityTypeInput struct {

	// The activity type to deprecate.
	//
	// This member is required.
	ActivityType *types.ActivityType

	// The name of the domain in which the activity type is registered.
	//
	// This member is required.
	Domain *string
}

type DeprecateActivityTypeOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeprecateActivityTypeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDeprecateActivityType{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDeprecateActivityType{}, middleware.After)
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
	addOpDeprecateActivityTypeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeprecateActivityType(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeprecateActivityType(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "swf",
		OperationName: "DeprecateActivityType",
	}
}
