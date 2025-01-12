// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a global PublicDelegatedPrefix in the specified project using the parameters that are included in the request.
type GlobalPublicDelegatedPrefix struct {
	pulumi.CustomResourceState

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a new PublicDelegatedPrefix. An up-to-date fingerprint must be provided in order to update the PublicDelegatedPrefix, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a PublicDelegatedPrefix.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// The IPv4 address range, in CIDR format, represented by this public delegated prefix.
	IpCidrRange pulumi.StringOutput `pulumi:"ipCidrRange"`
	// If true, the prefix will be live migrated.
	IsLiveMigration pulumi.BoolOutput `pulumi:"isLiveMigration"`
	// Type of the resource. Always compute#publicDelegatedPrefix for public delegated prefixes.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// The URL of parent prefix. Either PublicAdvertisedPrefix or PublicDelegatedPrefix.
	ParentPrefix pulumi.StringOutput `pulumi:"parentPrefix"`
	// The list of sub public delegated prefixes that exist for this public delegated prefix.
	PublicDelegatedSubPrefixs PublicDelegatedPrefixPublicDelegatedSubPrefixResponseArrayOutput `pulumi:"publicDelegatedSubPrefixs"`
	// URL of the region where the public delegated prefix resides. This field applies only to the region resource. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
	Region pulumi.StringOutput `pulumi:"region"`
	// Server-defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// The status of the public delegated prefix, which can be one of following values: - `INITIALIZING` The public delegated prefix is being initialized and addresses cannot be created yet. - `READY_TO_ANNOUNCE` The public delegated prefix is a live migration prefix and is active. - `ANNOUNCED` The public delegated prefix is active. - `DELETING` The public delegated prefix is being deprovsioned.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewGlobalPublicDelegatedPrefix registers a new resource with the given unique name, arguments, and options.
func NewGlobalPublicDelegatedPrefix(ctx *pulumi.Context,
	name string, args *GlobalPublicDelegatedPrefixArgs, opts ...pulumi.ResourceOption) (*GlobalPublicDelegatedPrefix, error) {
	if args == nil {
		args = &GlobalPublicDelegatedPrefixArgs{}
	}

	var resource GlobalPublicDelegatedPrefix
	err := ctx.RegisterResource("google-native:compute/v1:GlobalPublicDelegatedPrefix", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGlobalPublicDelegatedPrefix gets an existing GlobalPublicDelegatedPrefix resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGlobalPublicDelegatedPrefix(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GlobalPublicDelegatedPrefixState, opts ...pulumi.ResourceOption) (*GlobalPublicDelegatedPrefix, error) {
	var resource GlobalPublicDelegatedPrefix
	err := ctx.ReadResource("google-native:compute/v1:GlobalPublicDelegatedPrefix", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GlobalPublicDelegatedPrefix resources.
type globalPublicDelegatedPrefixState struct {
}

type GlobalPublicDelegatedPrefixState struct {
}

func (GlobalPublicDelegatedPrefixState) ElementType() reflect.Type {
	return reflect.TypeOf((*globalPublicDelegatedPrefixState)(nil)).Elem()
}

type globalPublicDelegatedPrefixArgs struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// The IPv4 address range, in CIDR format, represented by this public delegated prefix.
	IpCidrRange *string `pulumi:"ipCidrRange"`
	// If true, the prefix will be live migrated.
	IsLiveMigration *bool `pulumi:"isLiveMigration"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The URL of parent prefix. Either PublicAdvertisedPrefix or PublicDelegatedPrefix.
	ParentPrefix *string `pulumi:"parentPrefix"`
	Project      *string `pulumi:"project"`
	// The list of sub public delegated prefixes that exist for this public delegated prefix.
	PublicDelegatedSubPrefixs []PublicDelegatedPrefixPublicDelegatedSubPrefix `pulumi:"publicDelegatedSubPrefixs"`
	RequestId                 *string                                         `pulumi:"requestId"`
}

// The set of arguments for constructing a GlobalPublicDelegatedPrefix resource.
type GlobalPublicDelegatedPrefixArgs struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// The IPv4 address range, in CIDR format, represented by this public delegated prefix.
	IpCidrRange pulumi.StringPtrInput
	// If true, the prefix will be live migrated.
	IsLiveMigration pulumi.BoolPtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The URL of parent prefix. Either PublicAdvertisedPrefix or PublicDelegatedPrefix.
	ParentPrefix pulumi.StringPtrInput
	Project      pulumi.StringPtrInput
	// The list of sub public delegated prefixes that exist for this public delegated prefix.
	PublicDelegatedSubPrefixs PublicDelegatedPrefixPublicDelegatedSubPrefixArrayInput
	RequestId                 pulumi.StringPtrInput
}

func (GlobalPublicDelegatedPrefixArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*globalPublicDelegatedPrefixArgs)(nil)).Elem()
}

type GlobalPublicDelegatedPrefixInput interface {
	pulumi.Input

	ToGlobalPublicDelegatedPrefixOutput() GlobalPublicDelegatedPrefixOutput
	ToGlobalPublicDelegatedPrefixOutputWithContext(ctx context.Context) GlobalPublicDelegatedPrefixOutput
}

func (*GlobalPublicDelegatedPrefix) ElementType() reflect.Type {
	return reflect.TypeOf((**GlobalPublicDelegatedPrefix)(nil)).Elem()
}

func (i *GlobalPublicDelegatedPrefix) ToGlobalPublicDelegatedPrefixOutput() GlobalPublicDelegatedPrefixOutput {
	return i.ToGlobalPublicDelegatedPrefixOutputWithContext(context.Background())
}

func (i *GlobalPublicDelegatedPrefix) ToGlobalPublicDelegatedPrefixOutputWithContext(ctx context.Context) GlobalPublicDelegatedPrefixOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalPublicDelegatedPrefixOutput)
}

type GlobalPublicDelegatedPrefixOutput struct{ *pulumi.OutputState }

func (GlobalPublicDelegatedPrefixOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GlobalPublicDelegatedPrefix)(nil)).Elem()
}

func (o GlobalPublicDelegatedPrefixOutput) ToGlobalPublicDelegatedPrefixOutput() GlobalPublicDelegatedPrefixOutput {
	return o
}

func (o GlobalPublicDelegatedPrefixOutput) ToGlobalPublicDelegatedPrefixOutputWithContext(ctx context.Context) GlobalPublicDelegatedPrefixOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GlobalPublicDelegatedPrefixInput)(nil)).Elem(), &GlobalPublicDelegatedPrefix{})
	pulumi.RegisterOutputType(GlobalPublicDelegatedPrefixOutput{})
}
