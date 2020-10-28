// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about all of your database snapshots in Amazon Lightsail.
func (c *Client) GetRelationalDatabaseSnapshots(ctx context.Context, params *GetRelationalDatabaseSnapshotsInput, optFns ...func(*Options)) (*GetRelationalDatabaseSnapshotsOutput, error) {
	if params == nil {
		params = &GetRelationalDatabaseSnapshotsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRelationalDatabaseSnapshots", params, optFns, addOperationGetRelationalDatabaseSnapshotsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRelationalDatabaseSnapshotsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRelationalDatabaseSnapshotsInput struct {

	// The token to advance to the next page of results from your request. To get a
	// page token, perform an initial GetRelationalDatabaseSnapshots request. If your
	// results are paginated, the response will return a next page token that you can
	// specify as the page token in a subsequent request.
	PageToken *string
}

type GetRelationalDatabaseSnapshotsOutput struct {

	// The token to advance to the next page of resutls from your request. A next page
	// token is not returned if there are no more results to display. To get the next
	// page of results, perform another GetRelationalDatabaseSnapshots request and
	// specify the next page token using the pageToken parameter.
	NextPageToken *string

	// An object describing the result of your get relational database snapshots
	// request.
	RelationalDatabaseSnapshots []*types.RelationalDatabaseSnapshot

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetRelationalDatabaseSnapshotsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpGetRelationalDatabaseSnapshots{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpGetRelationalDatabaseSnapshots{})
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
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opGetRelationalDatabaseSnapshots(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opGetRelationalDatabaseSnapshots(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "GetRelationalDatabaseSnapshots",
	}
}
