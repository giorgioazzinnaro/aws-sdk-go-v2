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

// Allows the update of one or more parameters of a database in Amazon Lightsail.
// Parameter updates don't cause outages; therefore, their application is not
// subject to the preferred maintenance window. However, there are two ways in
// which parameter updates are applied: dynamic or pending-reboot. Parameters
// marked with a dynamic apply type are applied immediately. Parameters marked with
// a pending-reboot apply type are applied only after the database is rebooted
// using the reboot relational database operation. The update relational database
// parameters operation supports tag-based access control via resource tags applied
// to the resource identified by relationalDatabaseName. For more information, see
// the Lightsail Dev Guide
// (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-controlling-access-using-tags).
func (c *Client) UpdateRelationalDatabaseParameters(ctx context.Context, params *UpdateRelationalDatabaseParametersInput, optFns ...func(*Options)) (*UpdateRelationalDatabaseParametersOutput, error) {
	if params == nil {
		params = &UpdateRelationalDatabaseParametersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateRelationalDatabaseParameters", params, optFns, addOperationUpdateRelationalDatabaseParametersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateRelationalDatabaseParametersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateRelationalDatabaseParametersInput struct {

	// The database parameters to update.
	//
	// This member is required.
	Parameters []*types.RelationalDatabaseParameter

	// The name of your database for which to update parameters.
	//
	// This member is required.
	RelationalDatabaseName *string
}

type UpdateRelationalDatabaseParametersOutput struct {

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request.
	Operations []*types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateRelationalDatabaseParametersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpUpdateRelationalDatabaseParameters{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpUpdateRelationalDatabaseParameters{})
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
	if err := addOpUpdateRelationalDatabaseParametersValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opUpdateRelationalDatabaseParameters(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opUpdateRelationalDatabaseParameters(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "UpdateRelationalDatabaseParameters",
	}
}
