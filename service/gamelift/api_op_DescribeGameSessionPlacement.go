// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves properties and current status of a game session placement request. To
// get game session placement details, specify the placement ID. If successful, a
// GameSessionPlacement object is returned.
//
//     * CreateGameSession
//
//     *
// DescribeGameSessions
//
//     * DescribeGameSessionDetails
//
//     *
// SearchGameSessions
//
//     * UpdateGameSession
//
//     * GetGameSessionLogUrl
//
//     *
// Game session placements
//
//         * StartGameSessionPlacement
//
//         *
// DescribeGameSessionPlacement
//
//         * StopGameSessionPlacement
func (c *Client) DescribeGameSessionPlacement(ctx context.Context, params *DescribeGameSessionPlacementInput, optFns ...func(*Options)) (*DescribeGameSessionPlacementOutput, error) {
	if params == nil {
		params = &DescribeGameSessionPlacementInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeGameSessionPlacement", params, optFns, addOperationDescribeGameSessionPlacementMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeGameSessionPlacementOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request operation.
type DescribeGameSessionPlacementInput struct {

	// A unique identifier for a game session placement to retrieve.
	//
	// This member is required.
	PlacementId *string
}

// Represents the returned data in response to a request operation.
type DescribeGameSessionPlacementOutput struct {

	// Object that describes the requested game session placement.
	GameSessionPlacement *types.GameSessionPlacement

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeGameSessionPlacementMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpDescribeGameSessionPlacement{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpDescribeGameSessionPlacement{})
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
	if err := addOpDescribeGameSessionPlacementValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDescribeGameSessionPlacement(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDescribeGameSessionPlacement(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "DescribeGameSessionPlacement",
	}
}
