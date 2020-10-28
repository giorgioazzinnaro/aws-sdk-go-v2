// Code generated by smithy-go-codegen DO NOT EDIT.

package route53domains

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53domains/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// This operation returns detailed information about a specified domain that is
// associated with the current AWS account. Contact information for the domain is
// also returned as part of the output.
func (c *Client) GetDomainDetail(ctx context.Context, params *GetDomainDetailInput, optFns ...func(*Options)) (*GetDomainDetailOutput, error) {
	if params == nil {
		params = &GetDomainDetailInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDomainDetail", params, optFns, addOperationGetDomainDetailMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDomainDetailOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The GetDomainDetail request includes the following element.
type GetDomainDetailInput struct {

	// The name of the domain that you want to get detailed information about.
	//
	// This member is required.
	DomainName *string
}

// The GetDomainDetail response includes the following elements.
type GetDomainDetailOutput struct {

	// Provides details about the domain administrative contact.
	//
	// This member is required.
	AdminContact *types.ContactDetail

	// The name of a domain.
	//
	// This member is required.
	DomainName *string

	// The name of the domain.
	//
	// This member is required.
	Nameservers []*types.Nameserver

	// Provides details about the domain registrant.
	//
	// This member is required.
	RegistrantContact *types.ContactDetail

	// Provides details about the domain technical contact.
	//
	// This member is required.
	TechContact *types.ContactDetail

	// Email address to contact to report incorrect contact information for a domain,
	// to report that the domain is being used to send spam, to report that someone is
	// cybersquatting on a domain name, or report some other type of abuse.
	AbuseContactEmail *string

	// Phone number for reporting abuse.
	AbuseContactPhone *string

	// Specifies whether contact information is concealed from WHOIS queries. If the
	// value is true, WHOIS ("who is") queries return contact information either for
	// Amazon Registrar (for .com, .net, and .org domains) or for our registrar
	// associate, Gandi (for all other TLDs). If the value is false, WHOIS queries
	// return the information that you entered for the admin contact.
	AdminPrivacy *bool

	// Specifies whether the domain registration is set to renew automatically.
	AutoRenew *bool

	// The date when the domain was created as found in the response to a WHOIS query.
	// The date and time is in Unix time format and Coordinated Universal time (UTC).
	CreationDate *time.Time

	// Reserved for future use.
	DnsSec *string

	// The date when the registration for the domain is set to expire. The date and
	// time is in Unix time format and Coordinated Universal time (UTC).
	ExpirationDate *time.Time

	// Specifies whether contact information is concealed from WHOIS queries. If the
	// value is true, WHOIS ("who is") queries return contact information either for
	// Amazon Registrar (for .com, .net, and .org domains) or for our registrar
	// associate, Gandi (for all other TLDs). If the value is false, WHOIS queries
	// return the information that you entered for the registrant contact (domain
	// owner).
	RegistrantPrivacy *bool

	// Name of the registrar of the domain as identified in the registry. Domains with
	// a .com, .net, or .org TLD are registered by Amazon Registrar. All other domains
	// are registered by our registrar associate, Gandi. The value for domains that are
	// registered by Gandi is "GANDI SAS".
	RegistrarName *string

	// Web address of the registrar.
	RegistrarUrl *string

	// Reserved for future use.
	RegistryDomainId *string

	// Reseller of the domain. Domains registered or transferred using Route 53 domains
	// will have "Amazon" as the reseller.
	Reseller *string

	// An array of domain name status codes, also known as Extensible Provisioning
	// Protocol (EPP) status codes. ICANN, the organization that maintains a central
	// database of domain names, has developed a set of domain name status codes that
	// tell you the status of a variety of operations on a domain name, for example,
	// registering a domain name, transferring a domain name to another registrar,
	// renewing the registration for a domain name, and so on. All registrars use this
	// same set of status codes. For a current list of domain name status codes and an
	// explanation of what each code means, go to the ICANN website
	// (https://www.icann.org/) and search for epp status codes. (Search on the ICANN
	// website; web searches sometimes return an old version of the document.)
	StatusList []*string

	// Specifies whether contact information is concealed from WHOIS queries. If the
	// value is true, WHOIS ("who is") queries return contact information either for
	// Amazon Registrar (for .com, .net, and .org domains) or for our registrar
	// associate, Gandi (for all other TLDs). If the value is false, WHOIS queries
	// return the information that you entered for the technical contact.
	TechPrivacy *bool

	// The last updated date of the domain as found in the response to a WHOIS query.
	// The date and time is in Unix time format and Coordinated Universal time (UTC).
	UpdatedDate *time.Time

	// The fully qualified name of the WHOIS server that can answer the WHOIS query for
	// the domain.
	WhoIsServer *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetDomainDetailMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsjson11_serializeOpGetDomainDetail{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsjson11_deserializeOpGetDomainDetail{})
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
	if err := addOpGetDomainDetailValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opGetDomainDetail(options.Region)); err != nil {
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

func newServiceMetadataMiddleware_opGetDomainDetail(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53domains",
		OperationName: "GetDomainDetail",
	}
}
