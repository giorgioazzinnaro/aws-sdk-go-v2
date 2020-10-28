// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new custom endpoint and associates it with an Amazon Aurora DB
// cluster. This action only applies to Aurora DB clusters.
func (c *Client) CreateDBClusterEndpoint(ctx context.Context, params *CreateDBClusterEndpointInput, optFns ...func(*Options)) (*CreateDBClusterEndpointOutput, error) {
	if params == nil {
		params = &CreateDBClusterEndpointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDBClusterEndpoint", params, optFns, addOperationCreateDBClusterEndpointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDBClusterEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDBClusterEndpointInput struct {

	// The identifier to use for the new endpoint. This parameter is stored as a
	// lowercase string.
	//
	// This member is required.
	DBClusterEndpointIdentifier *string

	// The DB cluster identifier of the DB cluster associated with the endpoint. This
	// parameter is stored as a lowercase string.
	//
	// This member is required.
	DBClusterIdentifier *string

	// The type of the endpoint. One of: READER, WRITER, ANY.
	//
	// This member is required.
	EndpointType *string

	// List of DB instance identifiers that aren't part of the custom endpoint group.
	// All other eligible instances are reachable through the custom endpoint. Only
	// relevant if the list of static members is empty.
	ExcludedMembers []*string

	// List of DB instance identifiers that are part of the custom endpoint group.
	StaticMembers []*string

	// The tags to be assigned to the Amazon RDS resource.
	Tags []*types.Tag
}

// This data type represents the information you need to connect to an Amazon
// Aurora DB cluster. This data type is used as a response element in the following
// actions:
//
//     * CreateDBClusterEndpoint
//
//     * DescribeDBClusterEndpoints
//
//     *
// ModifyDBClusterEndpoint
//
//     * DeleteDBClusterEndpoint
//
// For the data structure
// that represents Amazon RDS DB instance endpoints, see Endpoint.
type CreateDBClusterEndpointOutput struct {

	// The type associated with a custom endpoint. One of: READER, WRITER, ANY.
	CustomEndpointType *string

	// The Amazon Resource Name (ARN) for the endpoint.
	DBClusterEndpointArn *string

	// The identifier associated with the endpoint. This parameter is stored as a
	// lowercase string.
	DBClusterEndpointIdentifier *string

	// A unique system-generated identifier for an endpoint. It remains the same for
	// the whole life of the endpoint.
	DBClusterEndpointResourceIdentifier *string

	// The DB cluster identifier of the DB cluster associated with the endpoint. This
	// parameter is stored as a lowercase string.
	DBClusterIdentifier *string

	// The DNS address of the endpoint.
	Endpoint *string

	// The type of the endpoint. One of: READER, WRITER, CUSTOM.
	EndpointType *string

	// List of DB instance identifiers that aren't part of the custom endpoint group.
	// All other eligible instances are reachable through the custom endpoint. Only
	// relevant if the list of static members is empty.
	ExcludedMembers []*string

	// List of DB instance identifiers that are part of the custom endpoint group.
	StaticMembers []*string

	// The current status of the endpoint. One of: creating, available, deleting,
	// inactive, modifying. The inactive state applies to an endpoint that can't be
	// used for a certain kind of cluster, such as a writer endpoint for a read-only
	// secondary cluster in a global database.
	Status *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateDBClusterEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsquery_serializeOpCreateDBClusterEndpoint{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsquery_deserializeOpCreateDBClusterEndpoint{})
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
	if err := addOpCreateDBClusterEndpointValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opCreateDBClusterEndpoint(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opCreateDBClusterEndpoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "CreateDBClusterEndpoint",
	}
}
