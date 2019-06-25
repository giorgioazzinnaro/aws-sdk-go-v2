// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesis

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input for EnableEnhancedMonitoring.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesis-2013-12-02/EnableEnhancedMonitoringInput
type EnableEnhancedMonitoringInput struct {
	_ struct{} `type:"structure"`

	// List of shard-level metrics to enable.
	//
	// The following are the valid shard-level metrics. The value "ALL" enables
	// every metric.
	//
	//    * IncomingBytes
	//
	//    * IncomingRecords
	//
	//    * OutgoingBytes
	//
	//    * OutgoingRecords
	//
	//    * WriteProvisionedThroughputExceeded
	//
	//    * ReadProvisionedThroughputExceeded
	//
	//    * IteratorAgeMilliseconds
	//
	//    * ALL
	//
	// For more information, see Monitoring the Amazon Kinesis Data Streams Service
	// with Amazon CloudWatch (http://docs.aws.amazon.com/kinesis/latest/dev/monitoring-with-cloudwatch.html)
	// in the Amazon Kinesis Data Streams Developer Guide.
	//
	// ShardLevelMetrics is a required field
	ShardLevelMetrics []MetricsName `min:"1" type:"list" required:"true"`

	// The name of the stream for which to enable enhanced monitoring.
	//
	// StreamName is a required field
	StreamName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s EnableEnhancedMonitoringInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *EnableEnhancedMonitoringInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "EnableEnhancedMonitoringInput"}

	if s.ShardLevelMetrics == nil {
		invalidParams.Add(aws.NewErrParamRequired("ShardLevelMetrics"))
	}
	if s.ShardLevelMetrics != nil && len(s.ShardLevelMetrics) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ShardLevelMetrics", 1))
	}

	if s.StreamName == nil {
		invalidParams.Add(aws.NewErrParamRequired("StreamName"))
	}
	if s.StreamName != nil && len(*s.StreamName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StreamName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output for EnableEnhancedMonitoring and DisableEnhancedMonitoring.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesis-2013-12-02/EnhancedMonitoringOutput
type EnableEnhancedMonitoringOutput struct {
	_ struct{} `type:"structure"`

	// Represents the current state of the metrics that are in the enhanced state
	// before the operation.
	CurrentShardLevelMetrics []MetricsName `min:"1" type:"list"`

	// Represents the list of all the metrics that would be in the enhanced state
	// after the operation.
	DesiredShardLevelMetrics []MetricsName `min:"1" type:"list"`

	// The name of the Kinesis data stream.
	StreamName *string `min:"1" type:"string"`
}

// String returns the string representation
func (s EnableEnhancedMonitoringOutput) String() string {
	return awsutil.Prettify(s)
}

const opEnableEnhancedMonitoring = "EnableEnhancedMonitoring"

// EnableEnhancedMonitoringRequest returns a request value for making API operation for
// Amazon Kinesis.
//
// Enables enhanced Kinesis data stream monitoring for shard-level metrics.
//
//    // Example sending a request using EnableEnhancedMonitoringRequest.
//    req := client.EnableEnhancedMonitoringRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesis-2013-12-02/EnableEnhancedMonitoring
func (c *Client) EnableEnhancedMonitoringRequest(input *EnableEnhancedMonitoringInput) EnableEnhancedMonitoringRequest {
	op := &aws.Operation{
		Name:       opEnableEnhancedMonitoring,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &EnableEnhancedMonitoringInput{}
	}

	req := c.newRequest(op, input, &EnableEnhancedMonitoringOutput{})
	return EnableEnhancedMonitoringRequest{Request: req, Input: input, Copy: c.EnableEnhancedMonitoringRequest}
}

// EnableEnhancedMonitoringRequest is the request type for the
// EnableEnhancedMonitoring API operation.
type EnableEnhancedMonitoringRequest struct {
	*aws.Request
	Input *EnableEnhancedMonitoringInput
	Copy  func(*EnableEnhancedMonitoringInput) EnableEnhancedMonitoringRequest
}

// Send marshals and sends the EnableEnhancedMonitoring API request.
func (r EnableEnhancedMonitoringRequest) Send(ctx context.Context) (*EnableEnhancedMonitoringResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &EnableEnhancedMonitoringResponse{
		EnableEnhancedMonitoringOutput: r.Request.Data.(*EnableEnhancedMonitoringOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// EnableEnhancedMonitoringResponse is the response type for the
// EnableEnhancedMonitoring API operation.
type EnableEnhancedMonitoringResponse struct {
	*EnableEnhancedMonitoringOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// EnableEnhancedMonitoring request.
func (r *EnableEnhancedMonitoringResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}