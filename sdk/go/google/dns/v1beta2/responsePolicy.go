// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new Response Policy
// Auto-naming is currently not supported for this resource.
type ResponsePolicy struct {
	pulumi.CustomResourceState

	// User-provided description for this Response Policy.
	Description pulumi.StringOutput `pulumi:"description"`
	// The list of Google Kubernetes Engine clusters to which this response policy is applied.
	GkeClusters ResponsePolicyGKEClusterResponseArrayOutput `pulumi:"gkeClusters"`
	Kind        pulumi.StringOutput                         `pulumi:"kind"`
	// List of network names specifying networks to which this policy is applied.
	Networks ResponsePolicyNetworkResponseArrayOutput `pulumi:"networks"`
	// User assigned name for this Response Policy.
	ResponsePolicyName pulumi.StringOutput `pulumi:"responsePolicyName"`
}

// NewResponsePolicy registers a new resource with the given unique name, arguments, and options.
func NewResponsePolicy(ctx *pulumi.Context,
	name string, args *ResponsePolicyArgs, opts ...pulumi.ResourceOption) (*ResponsePolicy, error) {
	if args == nil {
		args = &ResponsePolicyArgs{}
	}

	var resource ResponsePolicy
	err := ctx.RegisterResource("google-native:dns/v1beta2:ResponsePolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResponsePolicy gets an existing ResponsePolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResponsePolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResponsePolicyState, opts ...pulumi.ResourceOption) (*ResponsePolicy, error) {
	var resource ResponsePolicy
	err := ctx.ReadResource("google-native:dns/v1beta2:ResponsePolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResponsePolicy resources.
type responsePolicyState struct {
}

type ResponsePolicyState struct {
}

func (ResponsePolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*responsePolicyState)(nil)).Elem()
}

type responsePolicyArgs struct {
	ClientOperationId *string `pulumi:"clientOperationId"`
	// User-provided description for this Response Policy.
	Description *string `pulumi:"description"`
	// The list of Google Kubernetes Engine clusters to which this response policy is applied.
	GkeClusters []ResponsePolicyGKECluster `pulumi:"gkeClusters"`
	// Unique identifier for the resource; defined by the server (output only).
	Id   *string `pulumi:"id"`
	Kind *string `pulumi:"kind"`
	// List of network names specifying networks to which this policy is applied.
	Networks []ResponsePolicyNetwork `pulumi:"networks"`
	Project  *string                 `pulumi:"project"`
	// User assigned name for this Response Policy.
	ResponsePolicyName *string `pulumi:"responsePolicyName"`
}

// The set of arguments for constructing a ResponsePolicy resource.
type ResponsePolicyArgs struct {
	ClientOperationId pulumi.StringPtrInput
	// User-provided description for this Response Policy.
	Description pulumi.StringPtrInput
	// The list of Google Kubernetes Engine clusters to which this response policy is applied.
	GkeClusters ResponsePolicyGKEClusterArrayInput
	// Unique identifier for the resource; defined by the server (output only).
	Id   pulumi.StringPtrInput
	Kind pulumi.StringPtrInput
	// List of network names specifying networks to which this policy is applied.
	Networks ResponsePolicyNetworkArrayInput
	Project  pulumi.StringPtrInput
	// User assigned name for this Response Policy.
	ResponsePolicyName pulumi.StringPtrInput
}

func (ResponsePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*responsePolicyArgs)(nil)).Elem()
}

type ResponsePolicyInput interface {
	pulumi.Input

	ToResponsePolicyOutput() ResponsePolicyOutput
	ToResponsePolicyOutputWithContext(ctx context.Context) ResponsePolicyOutput
}

func (*ResponsePolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ResponsePolicy)(nil)).Elem()
}

func (i *ResponsePolicy) ToResponsePolicyOutput() ResponsePolicyOutput {
	return i.ToResponsePolicyOutputWithContext(context.Background())
}

func (i *ResponsePolicy) ToResponsePolicyOutputWithContext(ctx context.Context) ResponsePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResponsePolicyOutput)
}

type ResponsePolicyOutput struct{ *pulumi.OutputState }

func (ResponsePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResponsePolicy)(nil)).Elem()
}

func (o ResponsePolicyOutput) ToResponsePolicyOutput() ResponsePolicyOutput {
	return o
}

func (o ResponsePolicyOutput) ToResponsePolicyOutputWithContext(ctx context.Context) ResponsePolicyOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResponsePolicyInput)(nil)).Elem(), &ResponsePolicy{})
	pulumi.RegisterOutputType(ResponsePolicyOutput{})
}
