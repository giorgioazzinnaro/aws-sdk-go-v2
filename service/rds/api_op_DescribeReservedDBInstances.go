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

// Returns information about reserved DB instances for this account, or about a
// specified reserved DB instance.
func (c *Client) DescribeReservedDBInstances(ctx context.Context, params *DescribeReservedDBInstancesInput, optFns ...func(*Options)) (*DescribeReservedDBInstancesOutput, error) {
	if params == nil {
		params = &DescribeReservedDBInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeReservedDBInstances", params, optFns, addOperationDescribeReservedDBInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeReservedDBInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DescribeReservedDBInstancesInput struct {

	// The DB instance class filter value. Specify this parameter to show only those
	// reservations matching the specified DB instances class.
	DBInstanceClass *string

	// The duration filter value, specified in years or seconds. Specify this parameter
	// to show only reservations for this duration. Valid Values: 1 | 3 | 31536000 |
	// 94608000
	Duration *string

	// This parameter isn't currently supported.
	Filters []*types.Filter

	// The lease identifier filter value. Specify this parameter to show only the
	// reservation that matches the specified lease ID. AWS Support might request the
	// lease ID for an issue related to a reserved DB instance.
	LeaseId *string

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords.
	Marker *string

	// The maximum number of records to include in the response. If more than the
	// MaxRecords value is available, a pagination token called a marker is included in
	// the response so you can retrieve the remaining results. Default: 100
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int32

	// A value that indicates whether to show only those reservations that support
	// Multi-AZ.
	MultiAZ *bool

	// The offering type filter value. Specify this parameter to show only the
	// available offerings matching the specified offering type. Valid Values: "Partial
	// Upfront" | "All Upfront" | "No Upfront"
	OfferingType *string

	// The product description filter value. Specify this parameter to show only those
	// reservations matching the specified product description.
	ProductDescription *string

	// The reserved DB instance identifier filter value. Specify this parameter to show
	// only the reservation that matches the specified reservation ID.
	ReservedDBInstanceId *string

	// The offering identifier filter value. Specify this parameter to show only
	// purchased reservations matching the specified offering identifier.
	ReservedDBInstancesOfferingId *string
}

// Contains the result of a successful invocation of the
// DescribeReservedDBInstances action.
type DescribeReservedDBInstancesOutput struct {

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords.
	Marker *string

	// A list of reserved DB instances.
	ReservedDBInstances []*types.ReservedDBInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeReservedDBInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsquery_serializeOpDescribeReservedDBInstances{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsquery_deserializeOpDescribeReservedDBInstances{})
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
	if err := addOpDescribeReservedDBInstancesValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDescribeReservedDBInstances(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDescribeReservedDBInstances(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeReservedDBInstances",
	}
}
