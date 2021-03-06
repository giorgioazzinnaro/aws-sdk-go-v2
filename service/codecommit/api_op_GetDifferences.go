// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about the differences in a valid commit specifier (such as a
// branch, tag, HEAD, commit ID, or other fully qualified reference). Results can
// be limited to a specified path.
func (c *Client) GetDifferences(ctx context.Context, params *GetDifferencesInput, optFns ...func(*Options)) (*GetDifferencesOutput, error) {
	if params == nil {
		params = &GetDifferencesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDifferences", params, optFns, addOperationGetDifferencesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDifferencesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDifferencesInput struct {

	// The branch, tag, HEAD, or other fully qualified reference used to identify a
	// commit.
	//
	// This member is required.
	AfterCommitSpecifier *string

	// The name of the repository where you want to get differences.
	//
	// This member is required.
	RepositoryName *string

	// The file path in which to check differences. Limits the results to this path.
	// Can also be used to specify the changed name of a directory or folder, if it has
	// changed. If not specified, differences are shown for all paths.
	AfterPath *string

	// The branch, tag, HEAD, or other fully qualified reference used to identify a
	// commit (for example, the full commit ID). Optional. If not specified, all
	// changes before the afterCommitSpecifier value are shown. If you do not use
	// beforeCommitSpecifier in your request, consider limiting the results with
	// maxResults.
	BeforeCommitSpecifier *string

	// The file path in which to check for differences. Limits the results to this
	// path. Can also be used to specify the previous name of a directory or folder. If
	// beforePath and afterPath are not specified, differences are shown for all paths.
	BeforePath *string

	// A non-zero, non-negative integer used to limit the number of returned results.
	MaxResults *int32

	// An enumeration token that, when provided in a request, returns the next batch of
	// the results.
	NextToken *string
}

type GetDifferencesOutput struct {

	// A data type object that contains information about the differences, including
	// whether the difference is added, modified, or deleted (A, D, M).
	Differences []*types.Difference

	// An enumeration token that can be used in a request to return the next batch of
	// the results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetDifferencesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetDifferences{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetDifferences{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpGetDifferencesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetDifferences(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetDifferences(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "GetDifferences",
	}
}
