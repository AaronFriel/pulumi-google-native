// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates a domain mapping on the specified site.
type SiteDomain struct {
	pulumi.CustomResourceState

	// Required. The domain name of the association.
	DomainName pulumi.StringOutput `pulumi:"domainName"`
	// If set, the domain should redirect with the provided parameters.
	DomainRedirect DomainRedirectResponseOutput `pulumi:"domainRedirect"`
	// Information about the provisioning of certificates and the health of the DNS resolution for the domain.
	Provisioning DomainProvisioningResponseOutput `pulumi:"provisioning"`
	// Required. The site name of the association.
	Site pulumi.StringOutput `pulumi:"site"`
	// Additional status of the domain association.
	Status pulumi.StringOutput `pulumi:"status"`
	// The time at which the domain was last updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewSiteDomain registers a new resource with the given unique name, arguments, and options.
func NewSiteDomain(ctx *pulumi.Context,
	name string, args *SiteDomainArgs, opts ...pulumi.ResourceOption) (*SiteDomain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DomainsId == nil {
		return nil, errors.New("invalid value for required argument 'DomainsId'")
	}
	if args.SitesId == nil {
		return nil, errors.New("invalid value for required argument 'SitesId'")
	}
	var resource SiteDomain
	err := ctx.RegisterResource("google-cloud:firebasehosting/v1beta1:SiteDomain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSiteDomain gets an existing SiteDomain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSiteDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SiteDomainState, opts ...pulumi.ResourceOption) (*SiteDomain, error) {
	var resource SiteDomain
	err := ctx.ReadResource("google-cloud:firebasehosting/v1beta1:SiteDomain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SiteDomain resources.
type siteDomainState struct {
	// Required. The domain name of the association.
	DomainName *string `pulumi:"domainName"`
	// If set, the domain should redirect with the provided parameters.
	DomainRedirect *DomainRedirectResponse `pulumi:"domainRedirect"`
	// Information about the provisioning of certificates and the health of the DNS resolution for the domain.
	Provisioning *DomainProvisioningResponse `pulumi:"provisioning"`
	// Required. The site name of the association.
	Site *string `pulumi:"site"`
	// Additional status of the domain association.
	Status *string `pulumi:"status"`
	// The time at which the domain was last updated.
	UpdateTime *string `pulumi:"updateTime"`
}

type SiteDomainState struct {
	// Required. The domain name of the association.
	DomainName pulumi.StringPtrInput
	// If set, the domain should redirect with the provided parameters.
	DomainRedirect DomainRedirectResponsePtrInput
	// Information about the provisioning of certificates and the health of the DNS resolution for the domain.
	Provisioning DomainProvisioningResponsePtrInput
	// Required. The site name of the association.
	Site pulumi.StringPtrInput
	// Additional status of the domain association.
	Status pulumi.StringPtrInput
	// The time at which the domain was last updated.
	UpdateTime pulumi.StringPtrInput
}

func (SiteDomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*siteDomainState)(nil)).Elem()
}

type siteDomainArgs struct {
	// Required. The domain name of the association.
	DomainName *string `pulumi:"domainName"`
	// If set, the domain should redirect with the provided parameters.
	DomainRedirect *DomainRedirect `pulumi:"domainRedirect"`
	DomainsId      string          `pulumi:"domainsId"`
	// Information about the provisioning of certificates and the health of the DNS resolution for the domain.
	Provisioning *DomainProvisioning `pulumi:"provisioning"`
	// Required. The site name of the association.
	Site    *string `pulumi:"site"`
	SitesId string  `pulumi:"sitesId"`
	// Additional status of the domain association.
	Status *string `pulumi:"status"`
	// The time at which the domain was last updated.
	UpdateTime *string `pulumi:"updateTime"`
}

// The set of arguments for constructing a SiteDomain resource.
type SiteDomainArgs struct {
	// Required. The domain name of the association.
	DomainName pulumi.StringPtrInput
	// If set, the domain should redirect with the provided parameters.
	DomainRedirect DomainRedirectPtrInput
	DomainsId      pulumi.StringInput
	// Information about the provisioning of certificates and the health of the DNS resolution for the domain.
	Provisioning DomainProvisioningPtrInput
	// Required. The site name of the association.
	Site    pulumi.StringPtrInput
	SitesId pulumi.StringInput
	// Additional status of the domain association.
	Status pulumi.StringPtrInput
	// The time at which the domain was last updated.
	UpdateTime pulumi.StringPtrInput
}

func (SiteDomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*siteDomainArgs)(nil)).Elem()
}

type SiteDomainInput interface {
	pulumi.Input

	ToSiteDomainOutput() SiteDomainOutput
	ToSiteDomainOutputWithContext(ctx context.Context) SiteDomainOutput
}

func (*SiteDomain) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteDomain)(nil))
}

func (i *SiteDomain) ToSiteDomainOutput() SiteDomainOutput {
	return i.ToSiteDomainOutputWithContext(context.Background())
}

func (i *SiteDomain) ToSiteDomainOutputWithContext(ctx context.Context) SiteDomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteDomainOutput)
}

type SiteDomainOutput struct {
	*pulumi.OutputState
}

func (SiteDomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteDomain)(nil))
}

func (o SiteDomainOutput) ToSiteDomainOutput() SiteDomainOutput {
	return o
}

func (o SiteDomainOutput) ToSiteDomainOutputWithContext(ctx context.Context) SiteDomainOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SiteDomainOutput{})
}