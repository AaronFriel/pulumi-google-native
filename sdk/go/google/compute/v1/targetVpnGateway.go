// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a target VPN gateway in the specified project and region using the data included in the request.
type TargetVpnGateway struct {
	pulumi.CustomResourceState

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// A list of URLs to the ForwardingRule resources. ForwardingRules are created using compute.forwardingRules.insert and associated with a VPN gateway.
	ForwardingRules pulumi.StringArrayOutput `pulumi:"forwardingRules"`
	// Type of resource. Always compute#targetVpnGateway for target VPN gateways.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// URL of the network to which this VPN gateway is attached. Provided by the client when the VPN gateway is created.
	Network pulumi.StringOutput `pulumi:"network"`
	// URL of the region where the target VPN gateway resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
	Region pulumi.StringOutput `pulumi:"region"`
	// Server-defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// The status of the VPN gateway, which can be one of the following: CREATING, READY, FAILED, or DELETING.
	Status pulumi.StringOutput `pulumi:"status"`
	// A list of URLs to VpnTunnel resources. VpnTunnels are created using the compute.vpntunnels.insert method and associated with a VPN gateway.
	Tunnels pulumi.StringArrayOutput `pulumi:"tunnels"`
}

// NewTargetVpnGateway registers a new resource with the given unique name, arguments, and options.
func NewTargetVpnGateway(ctx *pulumi.Context,
	name string, args *TargetVpnGatewayArgs, opts ...pulumi.ResourceOption) (*TargetVpnGateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	var resource TargetVpnGateway
	err := ctx.RegisterResource("google-native:compute/v1:TargetVpnGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTargetVpnGateway gets an existing TargetVpnGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTargetVpnGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TargetVpnGatewayState, opts ...pulumi.ResourceOption) (*TargetVpnGateway, error) {
	var resource TargetVpnGateway
	err := ctx.ReadResource("google-native:compute/v1:TargetVpnGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TargetVpnGateway resources.
type targetVpnGatewayState struct {
}

type TargetVpnGatewayState struct {
}

func (TargetVpnGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*targetVpnGatewayState)(nil)).Elem()
}

type targetVpnGatewayArgs struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// URL of the network to which this VPN gateway is attached. Provided by the client when the VPN gateway is created.
	Network   *string `pulumi:"network"`
	Project   *string `pulumi:"project"`
	Region    string  `pulumi:"region"`
	RequestId *string `pulumi:"requestId"`
}

// The set of arguments for constructing a TargetVpnGateway resource.
type TargetVpnGatewayArgs struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// URL of the network to which this VPN gateway is attached. Provided by the client when the VPN gateway is created.
	Network   pulumi.StringPtrInput
	Project   pulumi.StringPtrInput
	Region    pulumi.StringInput
	RequestId pulumi.StringPtrInput
}

func (TargetVpnGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*targetVpnGatewayArgs)(nil)).Elem()
}

type TargetVpnGatewayInput interface {
	pulumi.Input

	ToTargetVpnGatewayOutput() TargetVpnGatewayOutput
	ToTargetVpnGatewayOutputWithContext(ctx context.Context) TargetVpnGatewayOutput
}

func (*TargetVpnGateway) ElementType() reflect.Type {
	return reflect.TypeOf((**TargetVpnGateway)(nil)).Elem()
}

func (i *TargetVpnGateway) ToTargetVpnGatewayOutput() TargetVpnGatewayOutput {
	return i.ToTargetVpnGatewayOutputWithContext(context.Background())
}

func (i *TargetVpnGateway) ToTargetVpnGatewayOutputWithContext(ctx context.Context) TargetVpnGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetVpnGatewayOutput)
}

type TargetVpnGatewayOutput struct{ *pulumi.OutputState }

func (TargetVpnGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TargetVpnGateway)(nil)).Elem()
}

func (o TargetVpnGatewayOutput) ToTargetVpnGatewayOutput() TargetVpnGatewayOutput {
	return o
}

func (o TargetVpnGatewayOutput) ToTargetVpnGatewayOutputWithContext(ctx context.Context) TargetVpnGatewayOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TargetVpnGatewayInput)(nil)).Elem(), &TargetVpnGateway{})
	pulumi.RegisterOutputType(TargetVpnGatewayOutput{})
}
