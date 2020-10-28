// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the Dedicated Host reservations that are available to purchase. The
// results describe all of the Dedicated Host reservation offerings, including
// offerings that might not match the instance family and Region of your Dedicated
// Hosts. When purchasing an offering, ensure that the instance family and Region
// of the offering matches that of the Dedicated Hosts with which it is to be
// associated. For more information about supported instance types, see Dedicated
// Hosts Overview
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/dedicated-hosts-overview.html)
// in the Amazon Elastic Compute Cloud User Guide.
func (c *Client) DescribeHostReservationOfferings(ctx context.Context, params *DescribeHostReservationOfferingsInput, optFns ...func(*Options)) (*DescribeHostReservationOfferingsOutput, error) {
	if params == nil {
		params = &DescribeHostReservationOfferingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeHostReservationOfferings", params, optFns, addOperationDescribeHostReservationOfferingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeHostReservationOfferingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeHostReservationOfferingsInput struct {

	// The filters.
	//
	//     * instance-family - The instance family of the offering (for
	// example, m4).
	//
	//     * payment-option - The payment option (NoUpfront |
	// PartialUpfront | AllUpfront).
	Filter []*types.Filter

	// This is the maximum duration of the reservation to purchase, specified in
	// seconds. Reservations are available in one-year and three-year terms. The number
	// of seconds specified must be the number of seconds in a year (365x24x60x60)
	// times one of the supported durations (1 or 3). For example, specify 94608000 for
	// three years.
	MaxDuration *int32

	// The maximum number of results to return for the request in a single page. The
	// remaining results can be seen by sending another request with the returned
	// nextToken value. This value can be between 5 and 500. If maxResults is given a
	// larger value than 500, you receive an error.
	MaxResults *int32

	// This is the minimum duration of the reservation you'd like to purchase,
	// specified in seconds. Reservations are available in one-year and three-year
	// terms. The number of seconds specified must be the number of seconds in a year
	// (365x24x60x60) times one of the supported durations (1 or 3). For example,
	// specify 31536000 for one year.
	MinDuration *int32

	// The token to use to retrieve the next page of results.
	NextToken *string

	// The ID of the reservation offering.
	OfferingId *string
}

type DescribeHostReservationOfferingsOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Information about the offerings.
	OfferingSet []*types.HostOffering

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeHostReservationOfferingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsEc2query_serializeOpDescribeHostReservationOfferings{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsEc2query_deserializeOpDescribeHostReservationOfferings{})
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
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDescribeHostReservationOfferings(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDescribeHostReservationOfferings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeHostReservationOfferings",
	}
}
