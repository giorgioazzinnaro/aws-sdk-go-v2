// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisanalytics

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kinesisanalytics/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This documentation is for version 1 of the Amazon Kinesis Data Analytics API,
// which only supports SQL applications. Version 2 of the API supports SQL and Java
// applications. For more information about version 2, see Amazon Kinesis Data
// Analytics API V2 Documentation. Creates an Amazon Kinesis Analytics application.
// You can configure each application with one streaming source as input,
// application code to process the input, and up to three destinations where you
// want Amazon Kinesis Analytics to write the output data from your application.
// For an overview, see How it Works
// (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/how-it-works.html). In
// the input configuration, you map the streaming source to an in-application
// stream, which you can think of as a constantly updating table. In the mapping,
// you must provide a schema for the in-application stream and map each data column
// in the in-application stream to a data element in the streaming source. Your
// application code is one or more SQL statements that read input data, transform
// it, and generate output. Your application code can create one or more SQL
// artifacts like SQL streams or pumps. In the output configuration, you can
// configure the application to write data from in-application streams created in
// your applications to up to three destinations. To read data from your source
// stream or write data to destination streams, Amazon Kinesis Analytics needs your
// permissions. You grant these permissions by creating IAM roles. This operation
// requires permissions to perform the kinesisanalytics:CreateApplication action.
// For introductory exercises to create an Amazon Kinesis Analytics application,
// see Getting Started
// (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/getting-started.html).
func (c *Client) CreateApplication(ctx context.Context, params *CreateApplicationInput, optFns ...func(*Options)) (*CreateApplicationOutput, error) {
	if params == nil {
		params = &CreateApplicationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateApplication", params, optFns, addOperationCreateApplicationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateApplicationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// TBD
type CreateApplicationInput struct {

	// Name of your Amazon Kinesis Analytics application (for example, sample-app).
	//
	// This member is required.
	ApplicationName *string

	// One or more SQL statements that read input data, transform it, and generate
	// output. For example, you can write a SQL statement that reads data from one
	// in-application stream, generates a running average of the number of
	// advertisement clicks by vendor, and insert resulting rows in another
	// in-application stream using pumps. For more information about the typical
	// pattern, see Application Code
	// (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/how-it-works-app-code.html).
	// You can provide such series of SQL statements, where output of one statement can
	// be used as the input for the next statement. You store intermediate results by
	// creating in-application streams and pumps. Note that the application code must
	// create the streams with names specified in the Outputs. For example, if your
	// Outputs defines output streams named ExampleOutputStream1 and
	// ExampleOutputStream2, then your application code must create these streams.
	ApplicationCode *string

	// Summary description of the application.
	ApplicationDescription *string

	// Use this parameter to configure a CloudWatch log stream to monitor application
	// configuration errors. For more information, see Working with Amazon CloudWatch
	// Logs
	// (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/cloudwatch-logs.html).
	CloudWatchLoggingOptions []*types.CloudWatchLoggingOption

	// Use this parameter to configure the application input. You can configure your
	// application to receive input from a single streaming source. In this
	// configuration, you map this streaming source to an in-application stream that is
	// created. Your application code can then query the in-application stream like a
	// table (you can think of it as a constantly updating table). For the streaming
	// source, you provide its Amazon Resource Name (ARN) and format of data on the
	// stream (for example, JSON, CSV, etc.). You also must provide an IAM role that
	// Amazon Kinesis Analytics can assume to read this stream on your behalf. To
	// create the in-application stream, you need to specify a schema to transform your
	// data into a schematized version used in SQL. In the schema, you provide the
	// necessary mapping of the data elements in the streaming source to record columns
	// in the in-app stream.
	Inputs []*types.Input

	// You can configure application output to write data from any of the
	// in-application streams to up to three destinations. These destinations can be
	// Amazon Kinesis streams, Amazon Kinesis Firehose delivery streams, AWS Lambda
	// destinations, or any combination of the three. In the configuration, you specify
	// the in-application stream name, the destination stream or Lambda function Amazon
	// Resource Name (ARN), and the format to use when writing data. You must also
	// provide an IAM role that Amazon Kinesis Analytics can assume to write to the
	// destination stream or Lambda function on your behalf. In the output
	// configuration, you also provide the output stream or Lambda function ARN. For
	// stream destinations, you provide the format of data in the stream (for example,
	// JSON, CSV). You also must provide an IAM role that Amazon Kinesis Analytics can
	// assume to write to the stream or Lambda function on your behalf.
	Outputs []*types.Output

	// A list of one or more tags to assign to the application. A tag is a key-value
	// pair that identifies an application. Note that the maximum number of application
	// tags includes system tags. The maximum number of user-defined application tags
	// is 50. For more information, see Using Tagging
	// (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/how-tagging.html).
	Tags []*types.Tag
}

// TBD
type CreateApplicationOutput struct {

	// In response to your CreateApplication request, Amazon Kinesis Analytics returns
	// a response with a summary of the application it created, including the
	// application Amazon Resource Name (ARN), name, and status.
	//
	// This member is required.
	ApplicationSummary *types.ApplicationSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateApplicationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpCreateApplication{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpCreateApplication{})
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
	if err := addOpCreateApplicationValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opCreateApplication(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opCreateApplication(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisanalytics",
		OperationName: "CreateApplication",
	}
}
