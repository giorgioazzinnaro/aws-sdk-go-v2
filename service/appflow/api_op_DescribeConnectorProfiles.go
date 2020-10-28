// Code generated by smithy-go-codegen DO NOT EDIT.

package appflow

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appflow/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of connector-profile details matching the provided
// connector-profile names and connector-types. Both input lists are optional, and
// you can use them to filter the result. If no names or connector-types are
// provided, returns all connector profiles in a paginated form. If there is no
// match, this operation returns an empty list.
func (c *Client) DescribeConnectorProfiles(ctx context.Context, params *DescribeConnectorProfilesInput, optFns ...func(*Options)) (*DescribeConnectorProfilesOutput, error) {
	if params == nil {
		params = &DescribeConnectorProfilesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeConnectorProfiles", params, optFns, addOperationDescribeConnectorProfilesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeConnectorProfilesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeConnectorProfilesInput struct {

	// The name of the connector profile. The name is unique for each ConnectorProfile
	// in the AWS account.
	ConnectorProfileNames []*string

	// The type of connector, such as Salesforce, Amplitude, and so on.
	ConnectorType types.ConnectorType

	// Specifies the maximum number of items that should be returned in the result set.
	// The default for maxResults is 20 (for all paginated API operations).
	MaxResults *int32

	// The pagination token for the next page of data.
	NextToken *string
}

type DescribeConnectorProfilesOutput struct {

	// Returns information about the connector profiles associated with the flow.
	ConnectorProfileDetails []*types.ConnectorProfile

	// The pagination token for the next page of data. If nextToken=null, this means
	// that all records have been fetched.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeConnectorProfilesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpDescribeConnectorProfiles{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpDescribeConnectorProfiles{})
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
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDescribeConnectorProfiles(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDescribeConnectorProfiles(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appflow",
		OperationName: "DescribeConnectorProfiles",
	}
}
