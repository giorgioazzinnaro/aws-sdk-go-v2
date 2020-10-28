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

// Creates a traffic policy, which you use to create multiple DNS resource record
// sets for one domain name (such as example.com) or one subdomain name (such as
// www.example.com).
func (c *Client) CreateTrafficPolicy(ctx context.Context, params *CreateTrafficPolicyInput, optFns ...func(*Options)) (*CreateTrafficPolicyOutput, error) {
	if params == nil {
		params = &CreateTrafficPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateTrafficPolicy", params, optFns, addOperationCreateTrafficPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateTrafficPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A complex type that contains information about the traffic policy that you want
// to create.
type CreateTrafficPolicyInput struct {

	// The definition of this traffic policy in JSON format. For more information, see
	// Traffic Policy Document Format
	// (https://docs.aws.amazon.com/Route53/latest/APIReference/api-policies-traffic-policy-document-format.html).
	//
	// This member is required.
	Document *string

	// The name of the traffic policy.
	//
	// This member is required.
	Name *string

	// (Optional) Any comments that you want to include about the traffic policy.
	Comment *string
}

// A complex type that contains the response information for the
// CreateTrafficPolicy request.
type CreateTrafficPolicyOutput struct {

	// A unique URL that represents a new traffic policy.
	//
	// This member is required.
	Location *string

	// A complex type that contains settings for the new traffic policy.
	//
	// This member is required.
	TrafficPolicy *types.TrafficPolicy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateTrafficPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestxml_serializeOpCreateTrafficPolicy{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestxml_deserializeOpCreateTrafficPolicy{})
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
	if err := addOpCreateTrafficPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opCreateTrafficPolicy(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opCreateTrafficPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "CreateTrafficPolicy",
	}
}
