// Code generated by smithy-go-codegen DO NOT EDIT.

package detective

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new behavior graph for the calling account, and sets that account as
// the master account. This operation is called by the account that is enabling
// Detective. Before you try to enable Detective, make sure that your account has
// been enrolled in Amazon GuardDuty for at least 48 hours. If you do not meet this
// requirement, you cannot enable Detective. If you do meet the GuardDuty
// prerequisite, then when you make the request to enable Detective, it checks
// whether your data volume is within the Detective quota. If it exceeds the quota,
// then you cannot enable Detective. The operation also enables Detective for the
// calling account in the currently selected Region. It returns the ARN of the new
// behavior graph. CreateGraph triggers a process to create the corresponding data
// tables for the new behavior graph. An account can only be the master account for
// one behavior graph within a Region. If the same account calls CreateGraph with
// the same master account, it always returns the same behavior graph ARN. It does
// not create a new behavior graph.
func (c *Client) CreateGraph(ctx context.Context, params *CreateGraphInput, optFns ...func(*Options)) (*CreateGraphOutput, error) {
	if params == nil {
		params = &CreateGraphInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateGraph", params, optFns, addOperationCreateGraphMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateGraphOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateGraphInput struct {
}

type CreateGraphOutput struct {

	// The ARN of the new behavior graph.
	GraphArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateGraphMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsRestjson1_serializeOpCreateGraph{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsRestjson1_deserializeOpCreateGraph{})
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
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opCreateGraph(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opCreateGraph(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "detective",
		OperationName: "CreateGraph",
	}
}
