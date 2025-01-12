// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a Backup for a domain.
// Auto-naming is currently not supported for this resource.
type Backup struct {
	pulumi.CustomResourceState

	// The time the backups was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Optional. Resource labels to represent user provided metadata.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The unique name of the Backup in the form of projects/{project_id}/locations/global/domains/{domain_name}/backups/{name}
	Name pulumi.StringOutput `pulumi:"name"`
	// The current state of the backup.
	State pulumi.StringOutput `pulumi:"state"`
	// Additional information about the current status of this backup, if available.
	StatusMessage pulumi.StringOutput `pulumi:"statusMessage"`
	// Indicates whether it’s an on-demand backup or scheduled.
	Type pulumi.StringOutput `pulumi:"type"`
	// Last update time.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewBackup registers a new resource with the given unique name, arguments, and options.
func NewBackup(ctx *pulumi.Context,
	name string, args *BackupArgs, opts ...pulumi.ResourceOption) (*Backup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BackupId == nil {
		return nil, errors.New("invalid value for required argument 'BackupId'")
	}
	if args.DomainId == nil {
		return nil, errors.New("invalid value for required argument 'DomainId'")
	}
	var resource Backup
	err := ctx.RegisterResource("google-native:managedidentities/v1beta1:Backup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackup gets an existing Backup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackupState, opts ...pulumi.ResourceOption) (*Backup, error) {
	var resource Backup
	err := ctx.ReadResource("google-native:managedidentities/v1beta1:Backup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Backup resources.
type backupState struct {
}

type BackupState struct {
}

func (BackupState) ElementType() reflect.Type {
	return reflect.TypeOf((*backupState)(nil)).Elem()
}

type backupArgs struct {
	BackupId string `pulumi:"backupId"`
	DomainId string `pulumi:"domainId"`
	// Optional. Resource labels to represent user provided metadata.
	Labels  map[string]string `pulumi:"labels"`
	Project *string           `pulumi:"project"`
}

// The set of arguments for constructing a Backup resource.
type BackupArgs struct {
	BackupId pulumi.StringInput
	DomainId pulumi.StringInput
	// Optional. Resource labels to represent user provided metadata.
	Labels  pulumi.StringMapInput
	Project pulumi.StringPtrInput
}

func (BackupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backupArgs)(nil)).Elem()
}

type BackupInput interface {
	pulumi.Input

	ToBackupOutput() BackupOutput
	ToBackupOutputWithContext(ctx context.Context) BackupOutput
}

func (*Backup) ElementType() reflect.Type {
	return reflect.TypeOf((**Backup)(nil)).Elem()
}

func (i *Backup) ToBackupOutput() BackupOutput {
	return i.ToBackupOutputWithContext(context.Background())
}

func (i *Backup) ToBackupOutputWithContext(ctx context.Context) BackupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupOutput)
}

type BackupOutput struct{ *pulumi.OutputState }

func (BackupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Backup)(nil)).Elem()
}

func (o BackupOutput) ToBackupOutput() BackupOutput {
	return o
}

func (o BackupOutput) ToBackupOutputWithContext(ctx context.Context) BackupOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BackupInput)(nil)).Elem(), &Backup{})
	pulumi.RegisterOutputType(BackupOutput{})
}
