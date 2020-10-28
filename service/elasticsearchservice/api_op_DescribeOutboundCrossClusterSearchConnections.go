// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticsearchservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists all the outbound cross-cluster search connections for a source domain.
func (c *Client) DescribeOutboundCrossClusterSearchConnections(ctx context.Context, params *DescribeOutboundCrossClusterSearchConnectionsInput, optFns ...func(*Options)) (*DescribeOutboundCrossClusterSearchConnectionsOutput, error) {
	if params == nil {
		params = &DescribeOutboundCrossClusterSearchConnectionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeOutboundCrossClusterSearchConnections", params, optFns, addOperationDescribeOutboundCrossClusterSearchConnectionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeOutboundCrossClusterSearchConnectionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the
// DescribeOutboundCrossClusterSearchConnections operation.
type DescribeOutboundCrossClusterSearchConnectionsInput struct {

	// A list of filters used to match properties for outbound cross-cluster search
	// connection. Available Filter names for this operation are:
	//
	//     *
	// cross-cluster-search-connection-id
	//
	//     * destination-domain-info.domain-name
	//
	//
	// * destination-domain-info.owner-id
	//
	//     * destination-domain-info.region
	//
	//     *
	// source-domain-info.domain-name
	Filters []*types.Filter

	// Set this value to limit the number of results returned. If not specified,
	// defaults to 100.
	MaxResults *int32

	// NextToken is sent in case the earlier API call results contain the NextToken. It
	// is used for pagination.
	NextToken *string
}

// The result of a DescribeOutboundCrossClusterSearchConnections request. Contains
// the list of connections matching the filter criteria.
type DescribeOutboundCrossClusterSearchConnectionsOutput struct {

	// Consists of list of OutboundCrossClusterSearchConnection matching the specified
	// filter criteria.
	CrossClusterSearchConnections []*types.OutboundCrossClusterSearchConnection

	// If more results are available and NextToken is present, make the next request to
	// the same API with the received NextToken to paginate the remaining results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeOutboundCrossClusterSearchConnectionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpDescribeOutboundCrossClusterSearchConnections{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpDescribeOutboundCrossClusterSearchConnections{})
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
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDescribeOutboundCrossClusterSearchConnections(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDescribeOutboundCrossClusterSearchConnections(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "DescribeOutboundCrossClusterSearchConnections",
	}
}
