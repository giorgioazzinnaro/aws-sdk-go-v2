// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateIAMPolicyAssignmentInput struct {
	_ struct{} `type:"structure"`

	// The name of the assignment. It must be unique within an AWS account.
	//
	// AssignmentName is a required field
	AssignmentName *string `min:"1" type:"string" required:"true"`

	// The status of an assignment:
	//
	//    * ENABLED - Anything specified in this assignment is used while creating
	//    the data source.
	//
	//    * DISABLED - This assignment isn't used while creating the data source.
	//
	//    * DRAFT - Assignment is an unfinished draft and isn't used while creating
	//    the data source.
	//
	// AssignmentStatus is a required field
	AssignmentStatus AssignmentStatus `type:"string" required:"true" enum:"true"`

	// The AWS Account ID where you want to assign QuickSight users or groups to
	// an IAM policy.
	//
	// AwsAccountId is a required field
	AwsAccountId *string `location:"uri" locationName:"AwsAccountId" min:"12" type:"string" required:"true"`

	// QuickSight users and/or groups that you want to assign the policy to.
	Identities map[string][]string `type:"map"`

	// The namespace that contains the assignment.
	//
	// Namespace is a required field
	Namespace *string `location:"uri" locationName:"Namespace" type:"string" required:"true"`

	// An IAM policy ARN that you want to apply to the QuickSight users and groups
	// specified in this assignment.
	PolicyArn *string `type:"string"`
}

// String returns the string representation
func (s CreateIAMPolicyAssignmentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateIAMPolicyAssignmentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateIAMPolicyAssignmentInput"}

	if s.AssignmentName == nil {
		invalidParams.Add(aws.NewErrParamRequired("AssignmentName"))
	}
	if s.AssignmentName != nil && len(*s.AssignmentName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AssignmentName", 1))
	}
	if len(s.AssignmentStatus) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("AssignmentStatus"))
	}

	if s.AwsAccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AwsAccountId"))
	}
	if s.AwsAccountId != nil && len(*s.AwsAccountId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("AwsAccountId", 12))
	}

	if s.Namespace == nil {
		invalidParams.Add(aws.NewErrParamRequired("Namespace"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateIAMPolicyAssignmentInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AssignmentName != nil {
		v := *s.AssignmentName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "AssignmentName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.AssignmentStatus) > 0 {
		v := s.AssignmentStatus

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "AssignmentStatus", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Identities != nil {
		v := s.Identities

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "Identities", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ls1 := ms0.List(k1)
			ls1.Start()
			for _, v2 := range v1 {
				ls1.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v2)})
			}
			ls1.End()
		}
		ms0.End()

	}
	if s.PolicyArn != nil {
		v := *s.PolicyArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "PolicyArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AwsAccountId != nil {
		v := *s.AwsAccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "AwsAccountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Namespace != nil {
		v := *s.Namespace

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Namespace", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type CreateIAMPolicyAssignmentOutput struct {
	_ struct{} `type:"structure"`

	// An ID for the assignment.
	AssignmentId *string `type:"string"`

	// The name of the assignment. Must be unique within an AWS account.
	AssignmentName *string `min:"1" type:"string"`

	// The status of an assignment:
	//
	//    * ENABLED - Anything specified in this assignment is used while creating
	//    the data source.
	//
	//    * DISABLED - This assignment isn't used while creating the data source.
	//
	//    * DRAFT - Assignment is an unfinished draft and isn't used while creating
	//    the data source.
	AssignmentStatus AssignmentStatus `type:"string" enum:"true"`

	// QuickSight users and/or groups that are assigned to the IAM policy.
	Identities map[string][]string `type:"map"`

	// An IAM policy ARN that is applied to the QuickSight users and groups specified
	// in this assignment.
	PolicyArn *string `type:"string"`

	// The AWS request ID for this operation.
	RequestId *string `type:"string"`

	// The http status of the request.
	Status *int64 `location:"statusCode" type:"integer"`
}

// String returns the string representation
func (s CreateIAMPolicyAssignmentOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateIAMPolicyAssignmentOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.AssignmentId != nil {
		v := *s.AssignmentId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "AssignmentId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AssignmentName != nil {
		v := *s.AssignmentName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "AssignmentName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.AssignmentStatus) > 0 {
		v := s.AssignmentStatus

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "AssignmentStatus", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Identities != nil {
		v := s.Identities

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "Identities", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ls1 := ms0.List(k1)
			ls1.Start()
			for _, v2 := range v1 {
				ls1.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v2)})
			}
			ls1.End()
		}
		ms0.End()

	}
	if s.PolicyArn != nil {
		v := *s.PolicyArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "PolicyArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RequestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	// ignoring invalid encode state, StatusCode. Status
	return nil
}

const opCreateIAMPolicyAssignment = "CreateIAMPolicyAssignment"

// CreateIAMPolicyAssignmentRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Creates an assignment with one specified IAM policy ARN and will assigned
// to specified groups or users of QuickSight. Users and groups need to be in
// the same namespace.
//
// CLI syntax:
//
// aws quicksight create-iam-policy-assignment --aws-account-id=111122223333
// --assignment-name=helpAssignment --policy-arn=arn:aws:iam::aws:policy/AdministratorAccess
// --identities="user=user5,engineer123,group=QS-Admin" --namespace=default
// --region=us-west-2
//
//    // Example sending a request using CreateIAMPolicyAssignmentRequest.
//    req := client.CreateIAMPolicyAssignmentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/CreateIAMPolicyAssignment
func (c *Client) CreateIAMPolicyAssignmentRequest(input *CreateIAMPolicyAssignmentInput) CreateIAMPolicyAssignmentRequest {
	op := &aws.Operation{
		Name:       opCreateIAMPolicyAssignment,
		HTTPMethod: "POST",
		HTTPPath:   "/accounts/{AwsAccountId}/namespaces/{Namespace}/iam-policy-assignments/",
	}

	if input == nil {
		input = &CreateIAMPolicyAssignmentInput{}
	}

	req := c.newRequest(op, input, &CreateIAMPolicyAssignmentOutput{})
	return CreateIAMPolicyAssignmentRequest{Request: req, Input: input, Copy: c.CreateIAMPolicyAssignmentRequest}
}

// CreateIAMPolicyAssignmentRequest is the request type for the
// CreateIAMPolicyAssignment API operation.
type CreateIAMPolicyAssignmentRequest struct {
	*aws.Request
	Input *CreateIAMPolicyAssignmentInput
	Copy  func(*CreateIAMPolicyAssignmentInput) CreateIAMPolicyAssignmentRequest
}

// Send marshals and sends the CreateIAMPolicyAssignment API request.
func (r CreateIAMPolicyAssignmentRequest) Send(ctx context.Context) (*CreateIAMPolicyAssignmentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateIAMPolicyAssignmentResponse{
		CreateIAMPolicyAssignmentOutput: r.Request.Data.(*CreateIAMPolicyAssignmentOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateIAMPolicyAssignmentResponse is the response type for the
// CreateIAMPolicyAssignment API operation.
type CreateIAMPolicyAssignmentResponse struct {
	*CreateIAMPolicyAssignmentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateIAMPolicyAssignment request.
func (r *CreateIAMPolicyAssignmentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}