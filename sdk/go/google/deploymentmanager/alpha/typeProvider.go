// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package alpha

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a type provider.
type TypeProvider struct {
	pulumi.CustomResourceState

	// Allows resource handling overrides for specific collections
	CollectionOverrides CollectionOverrideResponseArrayOutput `pulumi:"collectionOverrides"`
	// Credential used when interacting with this type.
	Credential CredentialResponseOutput `pulumi:"credential"`
	// List of up to 2 custom certificate authority roots to use for TLS authentication when making calls on behalf of this type provider. If set, TLS authentication will exclusively use these roots instead of relying on publicly trusted certificate authorities when validating TLS certificate authenticity. The certificates must be in base64-encoded PEM format. The maximum size of each certificate must not exceed 10KB.
	CustomCertificateAuthorityRoots pulumi.StringArrayOutput `pulumi:"customCertificateAuthorityRoots"`
	// An optional textual description of the resource; provided by the client when the resource is created.
	Description pulumi.StringOutput `pulumi:"description"`
	// Descriptor Url for the this type provider.
	DescriptorUrl pulumi.StringOutput `pulumi:"descriptorUrl"`
	// Creation timestamp in RFC3339 text format.
	InsertTime pulumi.StringOutput `pulumi:"insertTime"`
	// Map of One Platform labels; provided by the client when the resource is created or updated. Specifically: Label keys must be between 1 and 63 characters long and must conform to the following regular expression: `[a-z]([-a-z0-9]*[a-z0-9])?` Label values must be between 0 and 63 characters long and must conform to the regular expression `([a-z]([-a-z0-9]*[a-z0-9])?)?`
	Labels TypeProviderLabelEntryResponseArrayOutput `pulumi:"labels"`
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// The Operation that most recently ran, or is currently running, on this type provider.
	Operation OperationResponseOutput `pulumi:"operation"`
	// Options to apply when handling any resources in this service.
	Options OptionsResponseOutput `pulumi:"options"`
	// Self link for the type provider.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
}

// NewTypeProvider registers a new resource with the given unique name, arguments, and options.
func NewTypeProvider(ctx *pulumi.Context,
	name string, args *TypeProviderArgs, opts ...pulumi.ResourceOption) (*TypeProvider, error) {
	if args == nil {
		args = &TypeProviderArgs{}
	}

	var resource TypeProvider
	err := ctx.RegisterResource("google-native:deploymentmanager/alpha:TypeProvider", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTypeProvider gets an existing TypeProvider resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTypeProvider(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TypeProviderState, opts ...pulumi.ResourceOption) (*TypeProvider, error) {
	var resource TypeProvider
	err := ctx.ReadResource("google-native:deploymentmanager/alpha:TypeProvider", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TypeProvider resources.
type typeProviderState struct {
}

type TypeProviderState struct {
}

func (TypeProviderState) ElementType() reflect.Type {
	return reflect.TypeOf((*typeProviderState)(nil)).Elem()
}

type typeProviderArgs struct {
	// Allows resource handling overrides for specific collections
	CollectionOverrides []CollectionOverride `pulumi:"collectionOverrides"`
	// Credential used when interacting with this type.
	Credential *Credential `pulumi:"credential"`
	// List of up to 2 custom certificate authority roots to use for TLS authentication when making calls on behalf of this type provider. If set, TLS authentication will exclusively use these roots instead of relying on publicly trusted certificate authorities when validating TLS certificate authenticity. The certificates must be in base64-encoded PEM format. The maximum size of each certificate must not exceed 10KB.
	CustomCertificateAuthorityRoots []string `pulumi:"customCertificateAuthorityRoots"`
	// An optional textual description of the resource; provided by the client when the resource is created.
	Description *string `pulumi:"description"`
	// Descriptor Url for the this type provider.
	DescriptorUrl *string `pulumi:"descriptorUrl"`
	// Map of One Platform labels; provided by the client when the resource is created or updated. Specifically: Label keys must be between 1 and 63 characters long and must conform to the following regular expression: `[a-z]([-a-z0-9]*[a-z0-9])?` Label values must be between 0 and 63 characters long and must conform to the regular expression `([a-z]([-a-z0-9]*[a-z0-9])?)?`
	Labels []TypeProviderLabelEntry `pulumi:"labels"`
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// Options to apply when handling any resources in this service.
	Options *Options `pulumi:"options"`
	Project *string  `pulumi:"project"`
}

// The set of arguments for constructing a TypeProvider resource.
type TypeProviderArgs struct {
	// Allows resource handling overrides for specific collections
	CollectionOverrides CollectionOverrideArrayInput
	// Credential used when interacting with this type.
	Credential CredentialPtrInput
	// List of up to 2 custom certificate authority roots to use for TLS authentication when making calls on behalf of this type provider. If set, TLS authentication will exclusively use these roots instead of relying on publicly trusted certificate authorities when validating TLS certificate authenticity. The certificates must be in base64-encoded PEM format. The maximum size of each certificate must not exceed 10KB.
	CustomCertificateAuthorityRoots pulumi.StringArrayInput
	// An optional textual description of the resource; provided by the client when the resource is created.
	Description pulumi.StringPtrInput
	// Descriptor Url for the this type provider.
	DescriptorUrl pulumi.StringPtrInput
	// Map of One Platform labels; provided by the client when the resource is created or updated. Specifically: Label keys must be between 1 and 63 characters long and must conform to the following regular expression: `[a-z]([-a-z0-9]*[a-z0-9])?` Label values must be between 0 and 63 characters long and must conform to the regular expression `([a-z]([-a-z0-9]*[a-z0-9])?)?`
	Labels TypeProviderLabelEntryArrayInput
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// Options to apply when handling any resources in this service.
	Options OptionsPtrInput
	Project pulumi.StringPtrInput
}

func (TypeProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*typeProviderArgs)(nil)).Elem()
}

type TypeProviderInput interface {
	pulumi.Input

	ToTypeProviderOutput() TypeProviderOutput
	ToTypeProviderOutputWithContext(ctx context.Context) TypeProviderOutput
}

func (*TypeProvider) ElementType() reflect.Type {
	return reflect.TypeOf((**TypeProvider)(nil)).Elem()
}

func (i *TypeProvider) ToTypeProviderOutput() TypeProviderOutput {
	return i.ToTypeProviderOutputWithContext(context.Background())
}

func (i *TypeProvider) ToTypeProviderOutputWithContext(ctx context.Context) TypeProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TypeProviderOutput)
}

type TypeProviderOutput struct{ *pulumi.OutputState }

func (TypeProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TypeProvider)(nil)).Elem()
}

func (o TypeProviderOutput) ToTypeProviderOutput() TypeProviderOutput {
	return o
}

func (o TypeProviderOutput) ToTypeProviderOutputWithContext(ctx context.Context) TypeProviderOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TypeProviderInput)(nil)).Elem(), &TypeProvider{})
	pulumi.RegisterOutputType(TypeProviderOutput{})
}
