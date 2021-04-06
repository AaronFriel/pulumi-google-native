// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates a new Spoke in a given project and location.
type Spoke struct {
	pulumi.CustomResourceState

	// The time when the Spoke was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Short description of the spoke resource
	Description pulumi.StringOutput `pulumi:"description"`
	// The resource URL of the hub resource that the spoke is attached to
	Hub pulumi.StringOutput `pulumi:"hub"`
	// User-defined labels.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The URIs of linked interconnect attachment resources
	LinkedInterconnectAttachments pulumi.StringArrayOutput `pulumi:"linkedInterconnectAttachments"`
	// The URIs of linked Router appliance resources
	LinkedRouterApplianceInstances RouterApplianceInstanceResponseArrayOutput `pulumi:"linkedRouterApplianceInstances"`
	// The URIs of linked VPN tunnel resources
	LinkedVpnTunnels pulumi.StringArrayOutput `pulumi:"linkedVpnTunnels"`
	// Immutable. The name of a Spoke resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The current lifecycle state of this Hub.
	State pulumi.StringOutput `pulumi:"state"`
	// Google-generated UUID for this resource. This is unique across all Spoke resources. If a Spoke resource is deleted and another with the same name is created, it gets a different unique_id.
	UniqueId pulumi.StringOutput `pulumi:"uniqueId"`
	// The time when the Spoke was updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewSpoke registers a new resource with the given unique name, arguments, and options.
func NewSpoke(ctx *pulumi.Context,
	name string, args *SpokeArgs, opts ...pulumi.ResourceOption) (*Spoke, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LocationsId == nil {
		return nil, errors.New("invalid value for required argument 'LocationsId'")
	}
	if args.ProjectsId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectsId'")
	}
	if args.SpokesId == nil {
		return nil, errors.New("invalid value for required argument 'SpokesId'")
	}
	var resource Spoke
	err := ctx.RegisterResource("google-cloud:networkconnectivity/v1alpha1:Spoke", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSpoke gets an existing Spoke resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSpoke(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SpokeState, opts ...pulumi.ResourceOption) (*Spoke, error) {
	var resource Spoke
	err := ctx.ReadResource("google-cloud:networkconnectivity/v1alpha1:Spoke", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Spoke resources.
type spokeState struct {
	// The time when the Spoke was created.
	CreateTime *string `pulumi:"createTime"`
	// Short description of the spoke resource
	Description *string `pulumi:"description"`
	// The resource URL of the hub resource that the spoke is attached to
	Hub *string `pulumi:"hub"`
	// User-defined labels.
	Labels map[string]string `pulumi:"labels"`
	// The URIs of linked interconnect attachment resources
	LinkedInterconnectAttachments []string `pulumi:"linkedInterconnectAttachments"`
	// The URIs of linked Router appliance resources
	LinkedRouterApplianceInstances []RouterApplianceInstanceResponse `pulumi:"linkedRouterApplianceInstances"`
	// The URIs of linked VPN tunnel resources
	LinkedVpnTunnels []string `pulumi:"linkedVpnTunnels"`
	// Immutable. The name of a Spoke resource.
	Name *string `pulumi:"name"`
	// The current lifecycle state of this Hub.
	State *string `pulumi:"state"`
	// Google-generated UUID for this resource. This is unique across all Spoke resources. If a Spoke resource is deleted and another with the same name is created, it gets a different unique_id.
	UniqueId *string `pulumi:"uniqueId"`
	// The time when the Spoke was updated.
	UpdateTime *string `pulumi:"updateTime"`
}

type SpokeState struct {
	// The time when the Spoke was created.
	CreateTime pulumi.StringPtrInput
	// Short description of the spoke resource
	Description pulumi.StringPtrInput
	// The resource URL of the hub resource that the spoke is attached to
	Hub pulumi.StringPtrInput
	// User-defined labels.
	Labels pulumi.StringMapInput
	// The URIs of linked interconnect attachment resources
	LinkedInterconnectAttachments pulumi.StringArrayInput
	// The URIs of linked Router appliance resources
	LinkedRouterApplianceInstances RouterApplianceInstanceResponseArrayInput
	// The URIs of linked VPN tunnel resources
	LinkedVpnTunnels pulumi.StringArrayInput
	// Immutable. The name of a Spoke resource.
	Name pulumi.StringPtrInput
	// The current lifecycle state of this Hub.
	State pulumi.StringPtrInput
	// Google-generated UUID for this resource. This is unique across all Spoke resources. If a Spoke resource is deleted and another with the same name is created, it gets a different unique_id.
	UniqueId pulumi.StringPtrInput
	// The time when the Spoke was updated.
	UpdateTime pulumi.StringPtrInput
}

func (SpokeState) ElementType() reflect.Type {
	return reflect.TypeOf((*spokeState)(nil)).Elem()
}

type spokeArgs struct {
	// The time when the Spoke was created.
	CreateTime *string `pulumi:"createTime"`
	// Short description of the spoke resource
	Description *string `pulumi:"description"`
	// The resource URL of the hub resource that the spoke is attached to
	Hub *string `pulumi:"hub"`
	// User-defined labels.
	Labels map[string]string `pulumi:"labels"`
	// The URIs of linked interconnect attachment resources
	LinkedInterconnectAttachments []string `pulumi:"linkedInterconnectAttachments"`
	// The URIs of linked Router appliance resources
	LinkedRouterApplianceInstances []RouterApplianceInstance `pulumi:"linkedRouterApplianceInstances"`
	// The URIs of linked VPN tunnel resources
	LinkedVpnTunnels []string `pulumi:"linkedVpnTunnels"`
	LocationsId      string   `pulumi:"locationsId"`
	// Immutable. The name of a Spoke resource.
	Name       *string `pulumi:"name"`
	ProjectsId string  `pulumi:"projectsId"`
	SpokesId   string  `pulumi:"spokesId"`
	// The time when the Spoke was updated.
	UpdateTime *string `pulumi:"updateTime"`
}

// The set of arguments for constructing a Spoke resource.
type SpokeArgs struct {
	// The time when the Spoke was created.
	CreateTime pulumi.StringPtrInput
	// Short description of the spoke resource
	Description pulumi.StringPtrInput
	// The resource URL of the hub resource that the spoke is attached to
	Hub pulumi.StringPtrInput
	// User-defined labels.
	Labels pulumi.StringMapInput
	// The URIs of linked interconnect attachment resources
	LinkedInterconnectAttachments pulumi.StringArrayInput
	// The URIs of linked Router appliance resources
	LinkedRouterApplianceInstances RouterApplianceInstanceArrayInput
	// The URIs of linked VPN tunnel resources
	LinkedVpnTunnels pulumi.StringArrayInput
	LocationsId      pulumi.StringInput
	// Immutable. The name of a Spoke resource.
	Name       pulumi.StringPtrInput
	ProjectsId pulumi.StringInput
	SpokesId   pulumi.StringInput
	// The time when the Spoke was updated.
	UpdateTime pulumi.StringPtrInput
}

func (SpokeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*spokeArgs)(nil)).Elem()
}

type SpokeInput interface {
	pulumi.Input

	ToSpokeOutput() SpokeOutput
	ToSpokeOutputWithContext(ctx context.Context) SpokeOutput
}

func (*Spoke) ElementType() reflect.Type {
	return reflect.TypeOf((*Spoke)(nil))
}

func (i *Spoke) ToSpokeOutput() SpokeOutput {
	return i.ToSpokeOutputWithContext(context.Background())
}

func (i *Spoke) ToSpokeOutputWithContext(ctx context.Context) SpokeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpokeOutput)
}

type SpokeOutput struct {
	*pulumi.OutputState
}

func (SpokeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Spoke)(nil))
}

func (o SpokeOutput) ToSpokeOutput() SpokeOutput {
	return o
}

func (o SpokeOutput) ToSpokeOutputWithContext(ctx context.Context) SpokeOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SpokeOutput{})
}