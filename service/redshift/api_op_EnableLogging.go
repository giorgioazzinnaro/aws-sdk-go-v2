// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Starts logging information, such as queries and connection attempts, for the
// specified Amazon Redshift cluster.
func (c *Client) EnableLogging(ctx context.Context, params *EnableLoggingInput, optFns ...func(*Options)) (*EnableLoggingOutput, error) {
	if params == nil {
		params = &EnableLoggingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "EnableLogging", params, optFns, addOperationEnableLoggingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*EnableLoggingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type EnableLoggingInput struct {

	// The name of an existing S3 bucket where the log files are to be stored.
	// Constraints:
	//
	//     * Must be in the same region as the cluster
	//
	//     * The cluster
	// must have read bucket and put object permissions
	//
	// This member is required.
	BucketName *string

	// The identifier of the cluster on which logging is to be started. Example:
	// examplecluster
	//
	// This member is required.
	ClusterIdentifier *string

	// The prefix applied to the log file names. Constraints:
	//
	//     * Cannot exceed 512
	// characters
	//
	//     * Cannot contain spaces( ), double quotes ("), single quotes
	// ('), a backslash (\), or control characters. The hexadecimal codes for invalid
	// characters are:
	//
	//         * x00 to x20
	//
	//         * x22
	//
	//         * x27
	//
	//         *
	// x5c
	//
	//         * x7f or larger
	S3KeyPrefix *string
}

// Describes the status of logging for a cluster.
type EnableLoggingOutput struct {

	// The name of the S3 bucket where the log files are stored.
	BucketName *string

	// The message indicating that logs failed to be delivered.
	LastFailureMessage *string

	// The last time when logs failed to be delivered.
	LastFailureTime *time.Time

	// The last time that logs were delivered.
	LastSuccessfulDeliveryTime *time.Time

	// true if logging is on, false if logging is off.
	LoggingEnabled *bool

	// The prefix applied to the log file names.
	S3KeyPrefix *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationEnableLoggingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsquery_serializeOpEnableLogging{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsquery_deserializeOpEnableLogging{})
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
	if err := addOpEnableLoggingValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opEnableLogging(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opEnableLogging(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "EnableLogging",
	}
}
