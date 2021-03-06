// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the details of one or more retention configurations. If the retention
// configuration name is not specified, this action returns the details for all the
// retention configurations for that account. Currently, AWS Config supports only
// one retention configuration per region in your account.
func (c *Client) DescribeRetentionConfigurations(ctx context.Context, params *DescribeRetentionConfigurationsInput, optFns ...func(*Options)) (*DescribeRetentionConfigurationsOutput, error) {
	if params == nil {
		params = &DescribeRetentionConfigurationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeRetentionConfigurations", params, optFns, addOperationDescribeRetentionConfigurationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeRetentionConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeRetentionConfigurationsInput struct {

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	// A list of names of retention configurations for which you want details. If you
	// do not specify a name, AWS Config returns details for all the retention
	// configurations for that account. Currently, AWS Config supports only one
	// retention configuration per region in your account.
	RetentionConfigurationNames []*string
}

type DescribeRetentionConfigurationsOutput struct {

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	// Returns a retention configuration object.
	RetentionConfigurations []*types.RetentionConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeRetentionConfigurationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeRetentionConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeRetentionConfigurations{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeRetentionConfigurations(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeRetentionConfigurations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "DescribeRetentionConfigurations",
	}
}
