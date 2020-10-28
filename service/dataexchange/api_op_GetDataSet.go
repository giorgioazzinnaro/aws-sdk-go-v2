// Code generated by smithy-go-codegen DO NOT EDIT.

package dataexchange

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/dataexchange/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// This operation returns information about a data set.
func (c *Client) GetDataSet(ctx context.Context, params *GetDataSetInput, optFns ...func(*Options)) (*GetDataSetOutput, error) {
	if params == nil {
		params = &GetDataSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDataSet", params, optFns, addOperationGetDataSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDataSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDataSetInput struct {

	// The unique identifier for a data set.
	//
	// This member is required.
	DataSetId *string
}

type GetDataSetOutput struct {

	// The ARN for the data set.
	Arn *string

	// The type of file your data is stored in. Currently, the supported asset type is
	// S3_SNAPSHOT.
	AssetType types.AssetType

	// The date and time that the data set was created, in ISO 8601 format.
	CreatedAt *time.Time

	// The description for the data set.
	Description *string

	// The unique identifier for the data set.
	Id *string

	// The name of the data set.
	Name *string

	// A property that defines the data set as OWNED by the account (for providers) or
	// ENTITLED to the account (for subscribers).
	Origin types.Origin

	// If the origin of this data set is ENTITLED, includes the details for the product
	// on AWS Marketplace.
	OriginDetails *types.OriginDetails

	// The data set ID of the owned data set corresponding to the entitled data set
	// being viewed. This parameter is returned when a data set owner is viewing the
	// entitled copy of its owned data set.
	SourceId *string

	// The tags for the data set.
	Tags map[string]*string

	// The date and time that the data set was last updated, in ISO 8601 format.
	UpdatedAt *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetDataSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpGetDataSet{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpGetDataSet{})
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
	if err := addOpGetDataSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opGetDataSet(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opGetDataSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dataexchange",
		OperationName: "GetDataSet",
	}
}
