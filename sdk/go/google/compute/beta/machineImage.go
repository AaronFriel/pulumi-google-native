// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package beta

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a machine image in the specified project using the data that is included in the request. If you are creating a new machine image to update an existing instance, your new machine image should use the same network or, if applicable, the same subnetwork as the original instance.
type MachineImage struct {
	pulumi.CustomResourceState

	// The creation timestamp for this machine image in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// [Input Only] Whether to attempt an application consistent machine image by informing the OS to prepare for the snapshot process. Currently only supported on Windows instances using the Volume Shadow Copy Service (VSS).
	GuestFlush pulumi.BoolOutput `pulumi:"guestFlush"`
	// Properties of source instance
	InstanceProperties InstancePropertiesResponseOutput `pulumi:"instanceProperties"`
	// The resource type, which is always compute#machineImage for machine image.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Encrypts the machine image using a customer-supplied encryption key. After you encrypt a machine image using a customer-supplied key, you must provide the same key if you use the machine image later. For example, you must provide the encryption key when you create an instance from the encrypted machine image in a future request. Customer-supplied encryption keys do not protect access to metadata of the machine image. If you do not provide an encryption key when creating the machine image, then the machine image will be encrypted using an automatically generated key and you do not need to provide a key to use the machine image later.
	MachineImageEncryptionKey CustomerEncryptionKeyResponseOutput `pulumi:"machineImageEncryptionKey"`
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// Reserved for future use.
	SatisfiesPzs pulumi.BoolOutput `pulumi:"satisfiesPzs"`
	// An array of Machine Image specific properties for disks attached to the source instance
	SavedDisks SavedDiskResponseArrayOutput `pulumi:"savedDisks"`
	// The URL for this machine image. The server defines this URL.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// [Input Only] The customer-supplied encryption key of the disks attached to the source instance. Required if the source disk is protected by a customer-supplied encryption key.
	SourceDiskEncryptionKeys SourceDiskEncryptionKeyResponseArrayOutput `pulumi:"sourceDiskEncryptionKeys"`
	// The source instance used to create the machine image. You can provide this as a partial or full URL to the resource. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /instances/instance - projects/project/zones/zone/instances/instance
	SourceInstance pulumi.StringOutput `pulumi:"sourceInstance"`
	// DEPRECATED: Please use instance_properties instead for source instance related properties. New properties will not be added to this field.
	SourceInstanceProperties SourceInstancePropertiesResponseOutput `pulumi:"sourceInstanceProperties"`
	// The status of the machine image. One of the following values: INVALID, CREATING, READY, DELETING, and UPLOADING.
	Status pulumi.StringOutput `pulumi:"status"`
	// The regional or multi-regional Cloud Storage bucket location where the machine image is stored.
	StorageLocations pulumi.StringArrayOutput `pulumi:"storageLocations"`
	// Total size of the storage used by the machine image.
	TotalStorageBytes pulumi.StringOutput `pulumi:"totalStorageBytes"`
}

// NewMachineImage registers a new resource with the given unique name, arguments, and options.
func NewMachineImage(ctx *pulumi.Context,
	name string, args *MachineImageArgs, opts ...pulumi.ResourceOption) (*MachineImage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SourceInstance == nil {
		return nil, errors.New("invalid value for required argument 'SourceInstance'")
	}
	var resource MachineImage
	err := ctx.RegisterResource("google-native:compute/beta:MachineImage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMachineImage gets an existing MachineImage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMachineImage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MachineImageState, opts ...pulumi.ResourceOption) (*MachineImage, error) {
	var resource MachineImage
	err := ctx.ReadResource("google-native:compute/beta:MachineImage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MachineImage resources.
type machineImageState struct {
}

type MachineImageState struct {
}

func (MachineImageState) ElementType() reflect.Type {
	return reflect.TypeOf((*machineImageState)(nil)).Elem()
}

type machineImageArgs struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// [Input Only] Whether to attempt an application consistent machine image by informing the OS to prepare for the snapshot process. Currently only supported on Windows instances using the Volume Shadow Copy Service (VSS).
	GuestFlush *bool `pulumi:"guestFlush"`
	// Encrypts the machine image using a customer-supplied encryption key. After you encrypt a machine image using a customer-supplied key, you must provide the same key if you use the machine image later. For example, you must provide the encryption key when you create an instance from the encrypted machine image in a future request. Customer-supplied encryption keys do not protect access to metadata of the machine image. If you do not provide an encryption key when creating the machine image, then the machine image will be encrypted using an automatically generated key and you do not need to provide a key to use the machine image later.
	MachineImageEncryptionKey *CustomerEncryptionKey `pulumi:"machineImageEncryptionKey"`
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name      *string `pulumi:"name"`
	Project   *string `pulumi:"project"`
	RequestId *string `pulumi:"requestId"`
	// An array of Machine Image specific properties for disks attached to the source instance
	SavedDisks []SavedDisk `pulumi:"savedDisks"`
	// [Input Only] The customer-supplied encryption key of the disks attached to the source instance. Required if the source disk is protected by a customer-supplied encryption key.
	SourceDiskEncryptionKeys []SourceDiskEncryptionKey `pulumi:"sourceDiskEncryptionKeys"`
	// The source instance used to create the machine image. You can provide this as a partial or full URL to the resource. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /instances/instance - projects/project/zones/zone/instances/instance
	SourceInstance string `pulumi:"sourceInstance"`
	// The regional or multi-regional Cloud Storage bucket location where the machine image is stored.
	StorageLocations []string `pulumi:"storageLocations"`
}

// The set of arguments for constructing a MachineImage resource.
type MachineImageArgs struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// [Input Only] Whether to attempt an application consistent machine image by informing the OS to prepare for the snapshot process. Currently only supported on Windows instances using the Volume Shadow Copy Service (VSS).
	GuestFlush pulumi.BoolPtrInput
	// Encrypts the machine image using a customer-supplied encryption key. After you encrypt a machine image using a customer-supplied key, you must provide the same key if you use the machine image later. For example, you must provide the encryption key when you create an instance from the encrypted machine image in a future request. Customer-supplied encryption keys do not protect access to metadata of the machine image. If you do not provide an encryption key when creating the machine image, then the machine image will be encrypted using an automatically generated key and you do not need to provide a key to use the machine image later.
	MachineImageEncryptionKey CustomerEncryptionKeyPtrInput
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name      pulumi.StringPtrInput
	Project   pulumi.StringPtrInput
	RequestId pulumi.StringPtrInput
	// An array of Machine Image specific properties for disks attached to the source instance
	SavedDisks SavedDiskArrayInput
	// [Input Only] The customer-supplied encryption key of the disks attached to the source instance. Required if the source disk is protected by a customer-supplied encryption key.
	SourceDiskEncryptionKeys SourceDiskEncryptionKeyArrayInput
	// The source instance used to create the machine image. You can provide this as a partial or full URL to the resource. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /instances/instance - projects/project/zones/zone/instances/instance
	SourceInstance pulumi.StringInput
	// The regional or multi-regional Cloud Storage bucket location where the machine image is stored.
	StorageLocations pulumi.StringArrayInput
}

func (MachineImageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*machineImageArgs)(nil)).Elem()
}

type MachineImageInput interface {
	pulumi.Input

	ToMachineImageOutput() MachineImageOutput
	ToMachineImageOutputWithContext(ctx context.Context) MachineImageOutput
}

func (*MachineImage) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineImage)(nil)).Elem()
}

func (i *MachineImage) ToMachineImageOutput() MachineImageOutput {
	return i.ToMachineImageOutputWithContext(context.Background())
}

func (i *MachineImage) ToMachineImageOutputWithContext(ctx context.Context) MachineImageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineImageOutput)
}

type MachineImageOutput struct{ *pulumi.OutputState }

func (MachineImageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineImage)(nil)).Elem()
}

func (o MachineImageOutput) ToMachineImageOutput() MachineImageOutput {
	return o
}

func (o MachineImageOutput) ToMachineImageOutputWithContext(ctx context.Context) MachineImageOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MachineImageInput)(nil)).Elem(), &MachineImage{})
	pulumi.RegisterOutputType(MachineImageOutput{})
}
