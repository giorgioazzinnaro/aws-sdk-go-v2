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

// This operation is used with the Amazon GameLift FleetIQ solution and game server
// groups. Retrieves information for a registered game server. Information includes
// game server status, health check info, and the instance that the game server is
// running on. To retrieve game server information, specify the game server ID. If
// successful, the requested game server object is returned. Learn more GameLift
// FleetIQ Guide
// (https://docs.aws.amazon.com/gamelift/latest/fleetiqguide/gsg-intro.html)
// Related operations
//
//     * RegisterGameServer
//
//     * ListGameServers
//
//     *
// ClaimGameServer
//
//     * DescribeGameServer
//
//     * UpdateGameServer
//
//     *
// DeregisterGameServer
func (c *Client) DescribeGameServer(ctx context.Context, params *DescribeGameServerInput, optFns ...func(*Options)) (*DescribeGameServerOutput, error) {
	if params == nil {
		params = &DescribeGameServerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeGameServer", params, optFns, addOperationDescribeGameServerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeGameServerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeGameServerInput struct {

	// A unique identifier for the game server group where the game server is running.
	// Use either the GameServerGroup name or ARN value.
	//
	// This member is required.
	GameServerGroupName *string

	// A custom string that uniquely identifies the game server information to be
	// retrieved.
	//
	// This member is required.
	GameServerId *string
}

type DescribeGameServerOutput struct {

	// Object that describes the requested game server.
	GameServer *types.GameServer

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeGameServerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpDescribeGameServer{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpDescribeGameServer{})
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
	if err := addOpDescribeGameServerValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDescribeGameServer(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDescribeGameServer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "DescribeGameServer",
	}
}
