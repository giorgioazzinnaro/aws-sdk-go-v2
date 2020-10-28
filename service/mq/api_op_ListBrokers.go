// Code generated by smithy-go-codegen DO NOT EDIT.

package mq

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mq/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of all brokers.
func (c *Client) ListBrokers(ctx context.Context, params *ListBrokersInput, optFns ...func(*Options)) (*ListBrokersOutput, error) {
	if params == nil {
		params = &ListBrokersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListBrokers", params, optFns, addOperationListBrokersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListBrokersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListBrokersInput struct {

	// The maximum number of brokers that Amazon MQ can return per page (20 by
	// default). This value must be an integer from 5 to 100.
	MaxResults *int32

	// The token that specifies the next page of results Amazon MQ should return. To
	// request the first page, leave nextToken empty.
	NextToken *string
}

type ListBrokersOutput struct {

	// A list of information about all brokers.
	BrokerSummaries []*types.BrokerSummary

	// The token that specifies the next page of results Amazon MQ should return. To
	// request the first page, leave nextToken empty.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListBrokersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpListBrokers{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpListBrokers{})
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
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opListBrokers(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opListBrokers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mq",
		OperationName: "ListBrokers",
	}
}
