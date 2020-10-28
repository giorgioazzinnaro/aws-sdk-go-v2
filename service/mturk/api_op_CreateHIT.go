// Code generated by smithy-go-codegen DO NOT EDIT.

package mturk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mturk/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The CreateHIT operation creates a new Human Intelligence Task (HIT). The new HIT
// is made available for Workers to find and accept on the Amazon Mechanical Turk
// website. This operation allows you to specify a new HIT by passing in values for
// the properties of the HIT, such as its title, reward amount and number of
// assignments. When you pass these values to CreateHIT, a new HIT is created for
// you, with a new HITTypeID. The HITTypeID can be used to create additional HITs
// in the future without needing to specify common parameters such as the title,
// description and reward amount each time. An alternative way to create HITs is to
// first generate a HITTypeID using the CreateHITType operation and then call the
// CreateHITWithHITType operation. This is the recommended best practice for
// Requesters who are creating large numbers of HITs. CreateHIT also supports
// several ways to provide question data: by providing a value for the Question
// parameter that fully specifies the contents of the HIT, or by providing a
// HitLayoutId and associated HitLayoutParameters. If a HIT is created with 10 or
// more maximum assignments, there is an additional fee. For more information, see
// Amazon Mechanical Turk Pricing (https://requester.mturk.com/pricing).
func (c *Client) CreateHIT(ctx context.Context, params *CreateHITInput, optFns ...func(*Options)) (*CreateHITOutput, error) {
	if params == nil {
		params = &CreateHITInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateHIT", params, optFns, addOperationCreateHITMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateHITOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateHITInput struct {

	// The amount of time, in seconds, that a Worker has to complete the HIT after
	// accepting it. If a Worker does not complete the assignment within the specified
	// duration, the assignment is considered abandoned. If the HIT is still active
	// (that is, its lifetime has not elapsed), the assignment becomes available for
	// other users to find and accept.
	//
	// This member is required.
	AssignmentDurationInSeconds *int64

	// A general description of the HIT. A description includes detailed information
	// about the kind of task the HIT contains. On the Amazon Mechanical Turk web site,
	// the HIT description appears in the expanded view of search results, and in the
	// HIT and assignment screens. A good description gives the user enough information
	// to evaluate the HIT before accepting it.
	//
	// This member is required.
	Description *string

	// An amount of time, in seconds, after which the HIT is no longer available for
	// users to accept. After the lifetime of the HIT elapses, the HIT no longer
	// appears in HIT searches, even if not all of the assignments for the HIT have
	// been accepted.
	//
	// This member is required.
	LifetimeInSeconds *int64

	// The amount of money the Requester will pay a Worker for successfully completing
	// the HIT.
	//
	// This member is required.
	Reward *string

	// The title of the HIT. A title should be short and descriptive about the kind of
	// task the HIT contains. On the Amazon Mechanical Turk web site, the HIT title
	// appears in search results, and everywhere the HIT is mentioned.
	//
	// This member is required.
	Title *string

	// The Assignment-level Review Policy applies to the assignments under the HIT. You
	// can specify for Mechanical Turk to take various actions based on the policy.
	AssignmentReviewPolicy *types.ReviewPolicy

	// The number of seconds after an assignment for the HIT has been submitted, after
	// which the assignment is considered Approved automatically unless the Requester
	// explicitly rejects it.
	AutoApprovalDelayInSeconds *int64

	// The HITLayoutId allows you to use a pre-existing HIT design with placeholder
	// values and create an additional HIT by providing those values as
	// HITLayoutParameters. Constraints: Either a Question parameter or a HITLayoutId
	// parameter must be provided.
	HITLayoutId *string

	// If the HITLayoutId is provided, any placeholder values must be filled in with
	// values using the HITLayoutParameter structure. For more information, see
	// HITLayout.
	HITLayoutParameters []*types.HITLayoutParameter

	// The HIT-level Review Policy applies to the HIT. You can specify for Mechanical
	// Turk to take various actions based on the policy.
	HITReviewPolicy *types.ReviewPolicy

	// One or more words or phrases that describe the HIT, separated by commas. These
	// words are used in searches to find HITs.
	Keywords *string

	// The number of times the HIT can be accepted and completed before the HIT becomes
	// unavailable.
	MaxAssignments *int32

	// Conditions that a Worker's Qualifications must meet in order to accept the HIT.
	// A HIT can have between zero and ten Qualification requirements. All requirements
	// must be met in order for a Worker to accept the HIT. Additionally, other actions
	// can be restricted using the ActionsGuarded field on each
	// QualificationRequirement structure.
	QualificationRequirements []*types.QualificationRequirement

	// The data the person completing the HIT uses to produce the results. Constraints:
	// Must be a QuestionForm data structure, an ExternalQuestion data structure, or an
	// HTMLQuestion data structure. The XML question data must not be larger than 64
	// kilobytes (65,535 bytes) in size, including whitespace. Either a Question
	// parameter or a HITLayoutId parameter must be provided.
	Question *string

	// An arbitrary data field. The RequesterAnnotation parameter lets your application
	// attach arbitrary data to the HIT for tracking purposes. For example, this
	// parameter could be an identifier internal to the Requester's application that
	// corresponds with the HIT. The RequesterAnnotation parameter for a HIT is only
	// visible to the Requester who created the HIT. It is not shown to the Worker, or
	// any other Requester. The RequesterAnnotation parameter may be different for each
	// HIT you submit. It does not affect how your HITs are grouped.
	RequesterAnnotation *string

	// A unique identifier for this request which allows you to retry the call on error
	// without creating duplicate HITs. This is useful in cases such as network
	// timeouts where it is unclear whether or not the call succeeded on the server. If
	// the HIT already exists in the system from a previous call using the same
	// UniqueRequestToken, subsequent calls will return a
	// AWS.MechanicalTurk.HitAlreadyExists error with a message containing the HITId.
	// Note: It is your responsibility to ensure uniqueness of the token. The unique
	// token expires after 24 hours. Subsequent calls using the same UniqueRequestToken
	// made after the 24 hour limit could create duplicate HITs.
	UniqueRequestToken *string
}

type CreateHITOutput struct {

	// Contains the newly created HIT data. For a description of the HIT data structure
	// as it appears in responses, see the HIT Data Structure documentation.
	HIT *types.HIT

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateHITMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpCreateHIT{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpCreateHIT{})
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
	if err := addOpCreateHITValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opCreateHIT(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opCreateHIT(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mturk-requester",
		OperationName: "CreateHIT",
	}
}
