// Code generated by smithy-go-codegen DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/alexaforbusiness/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates room skill parameter details by room, skill, and parameter key ID. Not
// all skills have a room skill parameter.
func (c *Client) PutRoomSkillParameter(ctx context.Context, params *PutRoomSkillParameterInput, optFns ...func(*Options)) (*PutRoomSkillParameterOutput, error) {
	if params == nil {
		params = &PutRoomSkillParameterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutRoomSkillParameter", params, optFns, addOperationPutRoomSkillParameterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutRoomSkillParameterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutRoomSkillParameterInput struct {

	// The updated room skill parameter. Required.
	//
	// This member is required.
	RoomSkillParameter *types.RoomSkillParameter

	// The ARN of the skill associated with the room skill parameter. Required.
	//
	// This member is required.
	SkillId *string

	// The ARN of the room associated with the room skill parameter. Required.
	RoomArn *string
}

type PutRoomSkillParameterOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutRoomSkillParameterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpPutRoomSkillParameter{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpPutRoomSkillParameter{})
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
	if err := addOpPutRoomSkillParameterValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opPutRoomSkillParameter(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opPutRoomSkillParameter(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "PutRoomSkillParameter",
	}
}
