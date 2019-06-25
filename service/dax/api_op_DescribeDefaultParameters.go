// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dax

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/dax-2017-04-19/DescribeDefaultParametersRequest
type DescribeDefaultParametersInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to include in the response. If more results
	// exist than the specified MaxResults value, a token is included in the response
	// so that the remaining results can be retrieved.
	//
	// The value for MaxResults must be between 20 and 100.
	MaxResults *int64 `type:"integer"`

	// An optional token returned from a prior request. Use this token for pagination
	// of results from this action. If this parameter is specified, the response
	// includes only results beyond the token, up to the value specified by MaxResults.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeDefaultParametersInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/dax-2017-04-19/DescribeDefaultParametersResponse
type DescribeDefaultParametersOutput struct {
	_ struct{} `type:"structure"`

	// Provides an identifier to allow retrieval of paginated results.
	NextToken *string `type:"string"`

	// A list of parameters. Each element in the list represents one parameter.
	Parameters []Parameter `type:"list"`
}

// String returns the string representation
func (s DescribeDefaultParametersOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeDefaultParameters = "DescribeDefaultParameters"

// DescribeDefaultParametersRequest returns a request value for making API operation for
// Amazon DynamoDB Accelerator (DAX).
//
// Returns the default system parameter information for the DAX caching software.
//
//    // Example sending a request using DescribeDefaultParametersRequest.
//    req := client.DescribeDefaultParametersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dax-2017-04-19/DescribeDefaultParameters
func (c *Client) DescribeDefaultParametersRequest(input *DescribeDefaultParametersInput) DescribeDefaultParametersRequest {
	op := &aws.Operation{
		Name:       opDescribeDefaultParameters,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeDefaultParametersInput{}
	}

	req := c.newRequest(op, input, &DescribeDefaultParametersOutput{})
	return DescribeDefaultParametersRequest{Request: req, Input: input, Copy: c.DescribeDefaultParametersRequest}
}

// DescribeDefaultParametersRequest is the request type for the
// DescribeDefaultParameters API operation.
type DescribeDefaultParametersRequest struct {
	*aws.Request
	Input *DescribeDefaultParametersInput
	Copy  func(*DescribeDefaultParametersInput) DescribeDefaultParametersRequest
}

// Send marshals and sends the DescribeDefaultParameters API request.
func (r DescribeDefaultParametersRequest) Send(ctx context.Context) (*DescribeDefaultParametersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeDefaultParametersResponse{
		DescribeDefaultParametersOutput: r.Request.Data.(*DescribeDefaultParametersOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeDefaultParametersResponse is the response type for the
// DescribeDefaultParameters API operation.
type DescribeDefaultParametersResponse struct {
	*DescribeDefaultParametersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeDefaultParameters request.
func (r *DescribeDefaultParametersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}