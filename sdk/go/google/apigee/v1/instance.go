// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates an Apigee runtime instance. The instance is accessible from the authorized network configured on the organization. **Note:** Not supported for Apigee hybrid.
type Instance struct {
	pulumi.CustomResourceState

	// Time the instance was created in milliseconds since epoch.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Optional. Description of the instance.
	Description pulumi.StringOutput `pulumi:"description"`
	// Customer Managed Encryption Key (CMEK) used for disk and volume encryption. Required for Apigee paid subscriptions only. Use the following format: `projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)`
	DiskEncryptionKeyName pulumi.StringOutput `pulumi:"diskEncryptionKeyName"`
	// Optional. Display name for the instance.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Internal hostname or IP address of the Apigee endpoint used by clients to connect to the service.
	Host pulumi.StringOutput `pulumi:"host"`
	// Time the instance was last modified in milliseconds since epoch.
	LastModifiedAt pulumi.StringOutput `pulumi:"lastModifiedAt"`
	// Compute Engine location where the instance resides.
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource ID of the instance. Values must match the regular expression `^a-z{0,30}[a-z\d]$`.
	Name pulumi.StringOutput `pulumi:"name"`
	// Optional. Size of the CIDR block range that will be reserved by the instance. PAID organizations support `SLASH_16` to `SLASH_20` and defaults to `SLASH_16`. Evaluation organizations support only `SLASH_23`.
	PeeringCidrRange pulumi.StringOutput `pulumi:"peeringCidrRange"`
	// Port number of the exposed Apigee endpoint.
	Port pulumi.StringOutput `pulumi:"port"`
	// Version of the runtime system running in the instance. The runtime system is the set of components that serve the API Proxy traffic in your Environments.
	RuntimeVersion pulumi.StringOutput `pulumi:"runtimeVersion"`
	// State of the instance. Values other than `ACTIVE` means the resource is not ready to use.
	State pulumi.StringOutput `pulumi:"state"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OrganizationId == nil {
		return nil, errors.New("invalid value for required argument 'OrganizationId'")
	}
	var resource Instance
	err := ctx.RegisterResource("google-native:apigee/v1:Instance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstance gets an existing Instance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceState, opts ...pulumi.ResourceOption) (*Instance, error) {
	var resource Instance
	err := ctx.ReadResource("google-native:apigee/v1:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
}

type InstanceState struct {
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// Optional. Description of the instance.
	Description *string `pulumi:"description"`
	// Customer Managed Encryption Key (CMEK) used for disk and volume encryption. Required for Apigee paid subscriptions only. Use the following format: `projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)`
	DiskEncryptionKeyName *string `pulumi:"diskEncryptionKeyName"`
	// Optional. Display name for the instance.
	DisplayName *string `pulumi:"displayName"`
	// Compute Engine location where the instance resides.
	Location *string `pulumi:"location"`
	// Resource ID of the instance. Values must match the regular expression `^a-z{0,30}[a-z\d]$`.
	Name           *string `pulumi:"name"`
	OrganizationId string  `pulumi:"organizationId"`
	// Optional. Size of the CIDR block range that will be reserved by the instance. PAID organizations support `SLASH_16` to `SLASH_20` and defaults to `SLASH_16`. Evaluation organizations support only `SLASH_23`.
	PeeringCidrRange *InstancePeeringCidrRange `pulumi:"peeringCidrRange"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// Optional. Description of the instance.
	Description pulumi.StringPtrInput
	// Customer Managed Encryption Key (CMEK) used for disk and volume encryption. Required for Apigee paid subscriptions only. Use the following format: `projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)`
	DiskEncryptionKeyName pulumi.StringPtrInput
	// Optional. Display name for the instance.
	DisplayName pulumi.StringPtrInput
	// Compute Engine location where the instance resides.
	Location pulumi.StringPtrInput
	// Resource ID of the instance. Values must match the regular expression `^a-z{0,30}[a-z\d]$`.
	Name           pulumi.StringPtrInput
	OrganizationId pulumi.StringInput
	// Optional. Size of the CIDR block range that will be reserved by the instance. PAID organizations support `SLASH_16` to `SLASH_20` and defaults to `SLASH_16`. Evaluation organizations support only `SLASH_23`.
	PeeringCidrRange InstancePeeringCidrRangePtrInput
}

func (InstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceArgs)(nil)).Elem()
}

type InstanceInput interface {
	pulumi.Input

	ToInstanceOutput() InstanceOutput
	ToInstanceOutputWithContext(ctx context.Context) InstanceOutput
}

func (*Instance) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (i *Instance) ToInstanceOutput() InstanceOutput {
	return i.ToInstanceOutputWithContext(context.Background())
}

func (i *Instance) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceOutput)
}

type InstanceOutput struct{ *pulumi.OutputState }

func (InstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (o InstanceOutput) ToInstanceOutput() InstanceOutput {
	return o
}

func (o InstanceOutput) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceInput)(nil)).Elem(), &Instance{})
	pulumi.RegisterOutputType(InstanceOutput{})
}
