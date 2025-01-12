// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a TargetHttpProxy resource in the specified project using the data included in the request.
type TargetHttpProxy struct {
	pulumi.CustomResourceState

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a TargetHttpProxy. An up-to-date fingerprint must be provided in order to patch/update the TargetHttpProxy; otherwise, the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve the TargetHttpProxy.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// Type of resource. Always compute#targetHttpProxy for target HTTP proxies.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// This field only applies when the forwarding rule that references this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED. When this field is set to true, Envoy proxies set up inbound traffic interception and bind to the IP address and port specified in the forwarding rule. This is generally useful when using Traffic Director to configure Envoy as a gateway or middle proxy (in other words, not a sidecar proxy). The Envoy proxy listens for inbound requests and handles requests when it receives them. The default is false.
	ProxyBind pulumi.BoolOutput `pulumi:"proxyBind"`
	// URL of the region where the regional Target HTTP Proxy resides. This field is not applicable to global Target HTTP Proxies.
	Region pulumi.StringOutput `pulumi:"region"`
	// Server-defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// URL to the UrlMap resource that defines the mapping from URL to the BackendService.
	UrlMap pulumi.StringOutput `pulumi:"urlMap"`
}

// NewTargetHttpProxy registers a new resource with the given unique name, arguments, and options.
func NewTargetHttpProxy(ctx *pulumi.Context,
	name string, args *TargetHttpProxyArgs, opts ...pulumi.ResourceOption) (*TargetHttpProxy, error) {
	if args == nil {
		args = &TargetHttpProxyArgs{}
	}

	var resource TargetHttpProxy
	err := ctx.RegisterResource("google-native:compute/v1:TargetHttpProxy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTargetHttpProxy gets an existing TargetHttpProxy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTargetHttpProxy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TargetHttpProxyState, opts ...pulumi.ResourceOption) (*TargetHttpProxy, error) {
	var resource TargetHttpProxy
	err := ctx.ReadResource("google-native:compute/v1:TargetHttpProxy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TargetHttpProxy resources.
type targetHttpProxyState struct {
}

type TargetHttpProxyState struct {
}

func (TargetHttpProxyState) ElementType() reflect.Type {
	return reflect.TypeOf((*targetHttpProxyState)(nil)).Elem()
}

type targetHttpProxyArgs struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// This field only applies when the forwarding rule that references this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED. When this field is set to true, Envoy proxies set up inbound traffic interception and bind to the IP address and port specified in the forwarding rule. This is generally useful when using Traffic Director to configure Envoy as a gateway or middle proxy (in other words, not a sidecar proxy). The Envoy proxy listens for inbound requests and handles requests when it receives them. The default is false.
	ProxyBind *bool   `pulumi:"proxyBind"`
	RequestId *string `pulumi:"requestId"`
	// URL to the UrlMap resource that defines the mapping from URL to the BackendService.
	UrlMap *string `pulumi:"urlMap"`
}

// The set of arguments for constructing a TargetHttpProxy resource.
type TargetHttpProxyArgs struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// This field only applies when the forwarding rule that references this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED. When this field is set to true, Envoy proxies set up inbound traffic interception and bind to the IP address and port specified in the forwarding rule. This is generally useful when using Traffic Director to configure Envoy as a gateway or middle proxy (in other words, not a sidecar proxy). The Envoy proxy listens for inbound requests and handles requests when it receives them. The default is false.
	ProxyBind pulumi.BoolPtrInput
	RequestId pulumi.StringPtrInput
	// URL to the UrlMap resource that defines the mapping from URL to the BackendService.
	UrlMap pulumi.StringPtrInput
}

func (TargetHttpProxyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*targetHttpProxyArgs)(nil)).Elem()
}

type TargetHttpProxyInput interface {
	pulumi.Input

	ToTargetHttpProxyOutput() TargetHttpProxyOutput
	ToTargetHttpProxyOutputWithContext(ctx context.Context) TargetHttpProxyOutput
}

func (*TargetHttpProxy) ElementType() reflect.Type {
	return reflect.TypeOf((**TargetHttpProxy)(nil)).Elem()
}

func (i *TargetHttpProxy) ToTargetHttpProxyOutput() TargetHttpProxyOutput {
	return i.ToTargetHttpProxyOutputWithContext(context.Background())
}

func (i *TargetHttpProxy) ToTargetHttpProxyOutputWithContext(ctx context.Context) TargetHttpProxyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetHttpProxyOutput)
}

type TargetHttpProxyOutput struct{ *pulumi.OutputState }

func (TargetHttpProxyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TargetHttpProxy)(nil)).Elem()
}

func (o TargetHttpProxyOutput) ToTargetHttpProxyOutput() TargetHttpProxyOutput {
	return o
}

func (o TargetHttpProxyOutput) ToTargetHttpProxyOutputWithContext(ctx context.Context) TargetHttpProxyOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TargetHttpProxyInput)(nil)).Elem(), &TargetHttpProxy{})
	pulumi.RegisterOutputType(TargetHttpProxyOutput{})
}
