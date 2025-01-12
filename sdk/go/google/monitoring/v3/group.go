// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new group.
// Auto-naming is currently not supported for this resource.
type Group struct {
	pulumi.CustomResourceState

	// A user-assigned name for this group, used only for display purposes.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The filter used to determine which monitored resources belong to this group.
	Filter pulumi.StringOutput `pulumi:"filter"`
	// If true, the members of this group are considered to be a cluster. The system can perform additional analysis on groups that are clusters.
	IsCluster pulumi.BoolOutput `pulumi:"isCluster"`
	// The name of this group. The format is: projects/[PROJECT_ID_OR_NUMBER]/groups/[GROUP_ID] When creating a group, this field is ignored and a new name is created consisting of the project specified in the call to CreateGroup and a unique [GROUP_ID] that is generated automatically.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the group's parent, if it has one. The format is: projects/[PROJECT_ID_OR_NUMBER]/groups/[GROUP_ID] For groups with no parent, parent_name is the empty string, "".
	ParentName pulumi.StringOutput `pulumi:"parentName"`
}

// NewGroup registers a new resource with the given unique name, arguments, and options.
func NewGroup(ctx *pulumi.Context,
	name string, args *GroupArgs, opts ...pulumi.ResourceOption) (*Group, error) {
	if args == nil {
		args = &GroupArgs{}
	}

	var resource Group
	err := ctx.RegisterResource("google-native:monitoring/v3:Group", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroup gets an existing Group resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupState, opts ...pulumi.ResourceOption) (*Group, error) {
	var resource Group
	err := ctx.ReadResource("google-native:monitoring/v3:Group", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Group resources.
type groupState struct {
}

type GroupState struct {
}

func (GroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupState)(nil)).Elem()
}

type groupArgs struct {
	// A user-assigned name for this group, used only for display purposes.
	DisplayName *string `pulumi:"displayName"`
	// The filter used to determine which monitored resources belong to this group.
	Filter *string `pulumi:"filter"`
	// If true, the members of this group are considered to be a cluster. The system can perform additional analysis on groups that are clusters.
	IsCluster *bool `pulumi:"isCluster"`
	// The name of the group's parent, if it has one. The format is: projects/[PROJECT_ID_OR_NUMBER]/groups/[GROUP_ID] For groups with no parent, parent_name is the empty string, "".
	ParentName   *string `pulumi:"parentName"`
	Project      *string `pulumi:"project"`
	ValidateOnly *string `pulumi:"validateOnly"`
}

// The set of arguments for constructing a Group resource.
type GroupArgs struct {
	// A user-assigned name for this group, used only for display purposes.
	DisplayName pulumi.StringPtrInput
	// The filter used to determine which monitored resources belong to this group.
	Filter pulumi.StringPtrInput
	// If true, the members of this group are considered to be a cluster. The system can perform additional analysis on groups that are clusters.
	IsCluster pulumi.BoolPtrInput
	// The name of the group's parent, if it has one. The format is: projects/[PROJECT_ID_OR_NUMBER]/groups/[GROUP_ID] For groups with no parent, parent_name is the empty string, "".
	ParentName   pulumi.StringPtrInput
	Project      pulumi.StringPtrInput
	ValidateOnly pulumi.StringPtrInput
}

func (GroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupArgs)(nil)).Elem()
}

type GroupInput interface {
	pulumi.Input

	ToGroupOutput() GroupOutput
	ToGroupOutputWithContext(ctx context.Context) GroupOutput
}

func (*Group) ElementType() reflect.Type {
	return reflect.TypeOf((**Group)(nil)).Elem()
}

func (i *Group) ToGroupOutput() GroupOutput {
	return i.ToGroupOutputWithContext(context.Background())
}

func (i *Group) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupOutput)
}

type GroupOutput struct{ *pulumi.OutputState }

func (GroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Group)(nil)).Elem()
}

func (o GroupOutput) ToGroupOutput() GroupOutput {
	return o
}

func (o GroupOutput) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupInput)(nil)).Elem(), &Group{})
	pulumi.RegisterOutputType(GroupOutput{})
}
