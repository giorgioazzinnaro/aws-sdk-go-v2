// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudhsm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudhsm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This is documentation for AWS CloudHSM Classic. For more information, see AWS
// CloudHSM Classic FAQs (http://aws.amazon.com/cloudhsm/faqs-classic/), the AWS
// CloudHSM Classic User Guide
// (https://docs.aws.amazon.com/cloudhsm/classic/userguide/), and the AWS CloudHSM
// Classic API Reference
// (https://docs.aws.amazon.com/cloudhsm/classic/APIReference/). For information
// about the current version of AWS CloudHSM, see AWS CloudHSM
// (http://aws.amazon.com/cloudhsm/), the AWS CloudHSM User Guide
// (https://docs.aws.amazon.com/cloudhsm/latest/userguide/), and the AWS CloudHSM
// API Reference (https://docs.aws.amazon.com/cloudhsm/latest/APIReference/).
// Retrieves information about an HSM. You can identify the HSM by its ARN or its
// serial number.
func (c *Client) DescribeHsm(ctx context.Context, params *DescribeHsmInput, optFns ...func(*Options)) (*DescribeHsmOutput, error) {
	if params == nil {
		params = &DescribeHsmInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeHsm", params, optFns, addOperationDescribeHsmMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeHsmOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the inputs for the DescribeHsm operation.
type DescribeHsmInput struct {

	// The ARN of the HSM. Either the HsmArn or the SerialNumber parameter must be
	// specified.
	HsmArn *string

	// The serial number of the HSM. Either the HsmArn or the HsmSerialNumber parameter
	// must be specified.
	HsmSerialNumber *string
}

// Contains the output of the DescribeHsm operation.
type DescribeHsmOutput struct {

	// The Availability Zone that the HSM is in.
	AvailabilityZone *string

	// The identifier of the elastic network interface (ENI) attached to the HSM.
	EniId *string

	// The IP address assigned to the HSM's ENI.
	EniIp *string

	// The ARN of the HSM.
	HsmArn *string

	// The HSM model type.
	HsmType *string

	// The ARN of the IAM role assigned to the HSM.
	IamRoleArn *string

	// The list of partitions on the HSM.
	Partitions []*string

	// The serial number of the HSM.
	SerialNumber *string

	// The date and time that the server certificate was last updated.
	ServerCertLastUpdated *string

	// The URI of the certificate server.
	ServerCertUri *string

	// The HSM software version.
	SoftwareVersion *string

	// The date and time that the SSH key was last updated.
	SshKeyLastUpdated *string

	// The public SSH key.
	SshPublicKey *string

	// The status of the HSM.
	Status types.HsmStatus

	// Contains additional information about the status of the HSM.
	StatusDetails *string

	// The identifier of the subnet that the HSM is in.
	SubnetId *string

	// The subscription end date.
	SubscriptionEndDate *string

	// The subscription start date.
	SubscriptionStartDate *string

	// Specifies the type of subscription for the HSM.
	//
	//     * PRODUCTION - The HSM is
	// being used in a production environment.
	//
	//     * TRIAL - The HSM is being used in
	// a product trial.
	SubscriptionType types.SubscriptionType

	// The name of the HSM vendor.
	VendorName *string

	// The identifier of the VPC that the HSM is in.
	VpcId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeHsmMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpDescribeHsm{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpDescribeHsm{})
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
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opDescribeHsm(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opDescribeHsm(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudhsm",
		OperationName: "DescribeHsm",
	}
}
