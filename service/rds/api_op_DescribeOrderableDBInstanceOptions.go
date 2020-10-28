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

// Returns a list of orderable DB instance options for the specified engine.
func (c *Client) DescribeOrderableDBInstanceOptions(ctx context.Context, params *DescribeOrderableDBInstanceOptionsInput, optFns ...func(*Options)) (*DescribeOrderableDBInstanceOptionsOutput, error) {
	if params == nil {
		params = &DescribeOrderableDBInstanceOptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeOrderableDBInstanceOptions", params, optFns, addOperationDescribeOrderableDBInstanceOptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeOrderableDBInstanceOptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DescribeOrderableDBInstanceOptionsInput struct {

	// The name of the engine to retrieve DB instance options for.
	//
	// This member is required.
	Engine *string

	// The Availability Zone group associated with a Local Zone. Specify this parameter
	// to retrieve available offerings for the Local Zones in the group. Omit this
	// parameter to show the available offerings in the specified AWS Region.
	AvailabilityZoneGroup *string

	// The DB instance class filter value. Specify this parameter to show only the
	// available offerings matching the specified DB instance class.
	DBInstanceClass *string

	// The engine version filter value. Specify this parameter to show only the
	// available offerings matching the specified engine version.
	EngineVersion *string

	// This parameter isn't currently supported.
	Filters []*types.Filter

	// The license model filter value. Specify this parameter to show only the
	// available offerings matching the specified license model.
	LicenseModel *string

	// An optional pagination token provided by a previous
	// DescribeOrderableDBInstanceOptions request. If this parameter is specified, the
	// response includes only records beyond the marker, up to the value specified by
	// MaxRecords .
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that you can retrieve the remaining results.
	// Default: 100 Constraints: Minimum 20, maximum 100.
	MaxRecords *int32

	// A value that indicates whether to show only VPC or non-VPC offerings.
	Vpc *bool
}

// Contains the result of a successful invocation of the
// DescribeOrderableDBInstanceOptions action.
type DescribeOrderableDBInstanceOptionsOutput struct {

	// An optional pagination token provided by a previous OrderableDBInstanceOptions
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords .
	Marker *string

	// An OrderableDBInstanceOption structure containing information about orderable
	// options for the DB instance.
	OrderableDBInstanceOptions []*types.OrderableDBInstanceOption

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeOrderableDBInstanceOptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsquery_serializeOpDescribeOrderableDBInstanceOptions{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsquery_deserializeOpDescribeOrderableDBInstanceOptions{})
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
	if err := addOpDescribeOrderableDBInstanceOptionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDescribeOrderableDBInstanceOptions(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDescribeOrderableDBInstanceOptions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeOrderableDBInstanceOptions",
	}
}
