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

// Describes one or more of your instances, including information about the
// operating system platform, the version of SSM Agent installed on the instance,
// instance status, and so on. If you specify one or more instance IDs, it returns
// information for those instances. If you do not specify instance IDs, it returns
// information for all your instances. If you specify an instance ID that is not
// valid or an instance that you do not own, you receive an error. The IamRole
// field for this API action is the Amazon Identity and Access Management (IAM)
// role assigned to on-premises instances. This call does not return the IAM role
// for EC2 instances.
func (c *Client) DescribeInstanceInformation(ctx context.Context, params *DescribeInstanceInformationInput, optFns ...func(*Options)) (*DescribeInstanceInformationOutput, error) {
	if params == nil {
		params = &DescribeInstanceInformationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeInstanceInformation", params, optFns, addOperationDescribeInstanceInformationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeInstanceInformationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeInstanceInformationInput struct {

	// One or more filters. Use a filter to return a more specific list of instances.
	// You can filter based on tags applied to EC2 instances. Use this Filters data
	// type instead of InstanceInformationFilterList, which is deprecated.
	Filters []*types.InstanceInformationStringFilter

	// This is a legacy method. We recommend that you don't use this method. Instead,
	// use the Filters data type. Filters enables you to return instance information by
	// filtering based on tags applied to managed instances. Attempting to use
	// InstanceInformationFilterList and Filters leads to an exception error.
	InstanceInformationFilterList []*types.InstanceInformationFilter

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults *int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string
}

type DescribeInstanceInformationOutput struct {

	// The instance information list.
	InstanceInformationList []*types.InstanceInformation

	// The token to use when requesting the next set of items. If there are no
	// additional items to return, the string is empty.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeInstanceInformationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpDescribeInstanceInformation{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpDescribeInstanceInformation{})
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
	if err := addOpDescribeInstanceInformationValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDescribeInstanceInformation(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDescribeInstanceInformation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "DescribeInstanceInformation",
	}
}
