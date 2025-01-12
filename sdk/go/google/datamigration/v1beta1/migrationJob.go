// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new migration job in a given project and location.
type MigrationJob struct {
	pulumi.CustomResourceState

	// The timestamp when the migration job resource was created. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The resource name (URI) of the destination connection profile.
	Destination pulumi.StringOutput `pulumi:"destination"`
	// The database engine type and provider of the destination.
	DestinationDatabase DatabaseTypeResponseOutput `pulumi:"destinationDatabase"`
	// The migration job display name.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The path to the dump file in Google Cloud Storage, in the format: (gs://[BUCKET_NAME]/[OBJECT_NAME]).
	DumpPath pulumi.StringOutput `pulumi:"dumpPath"`
	// The duration of the migration job (in seconds). A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
	Duration pulumi.StringOutput `pulumi:"duration"`
	// If the migration job is completed, the time when it was completed.
	EndTime pulumi.StringOutput `pulumi:"endTime"`
	// The error details in case of state FAILED.
	Error StatusResponseOutput `pulumi:"error"`
	// The resource labels for migration job to use to annotate any related underlying resources such as Compute Engine VMs. An object containing a list of "key": "value" pairs. Example: `{ "name": "wrench", "mass": "1.3kg", "count": "3" }`.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The name (URI) of this migration job resource, in the form of: projects/{project}/locations/{location}/migrationJobs/{migrationJob}.
	Name pulumi.StringOutput `pulumi:"name"`
	// The current migration job phase.
	Phase pulumi.StringOutput `pulumi:"phase"`
	// The details needed to communicate to the source over Reverse SSH tunnel connectivity.
	ReverseSshConnectivity ReverseSshConnectivityResponseOutput `pulumi:"reverseSshConnectivity"`
	// The resource name (URI) of the source connection profile.
	Source pulumi.StringOutput `pulumi:"source"`
	// The database engine type and provider of the source.
	SourceDatabase DatabaseTypeResponseOutput `pulumi:"sourceDatabase"`
	// The current migration job state.
	State pulumi.StringOutput `pulumi:"state"`
	// static ip connectivity data (default, no additional details needed).
	StaticIpConnectivity StaticIpConnectivityResponseOutput `pulumi:"staticIpConnectivity"`
	// The migration job type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The timestamp when the migration job resource was last updated. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// The details of the VPC network that the source database is located in.
	VpcPeeringConnectivity VpcPeeringConnectivityResponseOutput `pulumi:"vpcPeeringConnectivity"`
}

// NewMigrationJob registers a new resource with the given unique name, arguments, and options.
func NewMigrationJob(ctx *pulumi.Context,
	name string, args *MigrationJobArgs, opts ...pulumi.ResourceOption) (*MigrationJob, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Destination == nil {
		return nil, errors.New("invalid value for required argument 'Destination'")
	}
	if args.MigrationJobId == nil {
		return nil, errors.New("invalid value for required argument 'MigrationJobId'")
	}
	if args.Source == nil {
		return nil, errors.New("invalid value for required argument 'Source'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource MigrationJob
	err := ctx.RegisterResource("google-native:datamigration/v1beta1:MigrationJob", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMigrationJob gets an existing MigrationJob resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMigrationJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MigrationJobState, opts ...pulumi.ResourceOption) (*MigrationJob, error) {
	var resource MigrationJob
	err := ctx.ReadResource("google-native:datamigration/v1beta1:MigrationJob", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MigrationJob resources.
type migrationJobState struct {
}

type MigrationJobState struct {
}

func (MigrationJobState) ElementType() reflect.Type {
	return reflect.TypeOf((*migrationJobState)(nil)).Elem()
}

type migrationJobArgs struct {
	// The resource name (URI) of the destination connection profile.
	Destination string `pulumi:"destination"`
	// The database engine type and provider of the destination.
	DestinationDatabase *DatabaseType `pulumi:"destinationDatabase"`
	// The migration job display name.
	DisplayName *string `pulumi:"displayName"`
	// The path to the dump file in Google Cloud Storage, in the format: (gs://[BUCKET_NAME]/[OBJECT_NAME]).
	DumpPath *string `pulumi:"dumpPath"`
	// The resource labels for migration job to use to annotate any related underlying resources such as Compute Engine VMs. An object containing a list of "key": "value" pairs. Example: `{ "name": "wrench", "mass": "1.3kg", "count": "3" }`.
	Labels         map[string]string `pulumi:"labels"`
	Location       *string           `pulumi:"location"`
	MigrationJobId string            `pulumi:"migrationJobId"`
	// The name (URI) of this migration job resource, in the form of: projects/{project}/locations/{location}/migrationJobs/{migrationJob}.
	Name      *string `pulumi:"name"`
	Project   *string `pulumi:"project"`
	RequestId *string `pulumi:"requestId"`
	// The details needed to communicate to the source over Reverse SSH tunnel connectivity.
	ReverseSshConnectivity *ReverseSshConnectivity `pulumi:"reverseSshConnectivity"`
	// The resource name (URI) of the source connection profile.
	Source string `pulumi:"source"`
	// The database engine type and provider of the source.
	SourceDatabase *DatabaseType `pulumi:"sourceDatabase"`
	// The current migration job state.
	State *MigrationJobStateEnum `pulumi:"state"`
	// static ip connectivity data (default, no additional details needed).
	StaticIpConnectivity *StaticIpConnectivity `pulumi:"staticIpConnectivity"`
	// The migration job type.
	Type MigrationJobType `pulumi:"type"`
	// The details of the VPC network that the source database is located in.
	VpcPeeringConnectivity *VpcPeeringConnectivity `pulumi:"vpcPeeringConnectivity"`
}

// The set of arguments for constructing a MigrationJob resource.
type MigrationJobArgs struct {
	// The resource name (URI) of the destination connection profile.
	Destination pulumi.StringInput
	// The database engine type and provider of the destination.
	DestinationDatabase DatabaseTypePtrInput
	// The migration job display name.
	DisplayName pulumi.StringPtrInput
	// The path to the dump file in Google Cloud Storage, in the format: (gs://[BUCKET_NAME]/[OBJECT_NAME]).
	DumpPath pulumi.StringPtrInput
	// The resource labels for migration job to use to annotate any related underlying resources such as Compute Engine VMs. An object containing a list of "key": "value" pairs. Example: `{ "name": "wrench", "mass": "1.3kg", "count": "3" }`.
	Labels         pulumi.StringMapInput
	Location       pulumi.StringPtrInput
	MigrationJobId pulumi.StringInput
	// The name (URI) of this migration job resource, in the form of: projects/{project}/locations/{location}/migrationJobs/{migrationJob}.
	Name      pulumi.StringPtrInput
	Project   pulumi.StringPtrInput
	RequestId pulumi.StringPtrInput
	// The details needed to communicate to the source over Reverse SSH tunnel connectivity.
	ReverseSshConnectivity ReverseSshConnectivityPtrInput
	// The resource name (URI) of the source connection profile.
	Source pulumi.StringInput
	// The database engine type and provider of the source.
	SourceDatabase DatabaseTypePtrInput
	// The current migration job state.
	State MigrationJobStateEnumPtrInput
	// static ip connectivity data (default, no additional details needed).
	StaticIpConnectivity StaticIpConnectivityPtrInput
	// The migration job type.
	Type MigrationJobTypeInput
	// The details of the VPC network that the source database is located in.
	VpcPeeringConnectivity VpcPeeringConnectivityPtrInput
}

func (MigrationJobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*migrationJobArgs)(nil)).Elem()
}

type MigrationJobInput interface {
	pulumi.Input

	ToMigrationJobOutput() MigrationJobOutput
	ToMigrationJobOutputWithContext(ctx context.Context) MigrationJobOutput
}

func (*MigrationJob) ElementType() reflect.Type {
	return reflect.TypeOf((**MigrationJob)(nil)).Elem()
}

func (i *MigrationJob) ToMigrationJobOutput() MigrationJobOutput {
	return i.ToMigrationJobOutputWithContext(context.Background())
}

func (i *MigrationJob) ToMigrationJobOutputWithContext(ctx context.Context) MigrationJobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrationJobOutput)
}

type MigrationJobOutput struct{ *pulumi.OutputState }

func (MigrationJobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MigrationJob)(nil)).Elem()
}

func (o MigrationJobOutput) ToMigrationJobOutput() MigrationJobOutput {
	return o
}

func (o MigrationJobOutput) ToMigrationJobOutputWithContext(ctx context.Context) MigrationJobOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MigrationJobInput)(nil)).Elem(), &MigrationJob{})
	pulumi.RegisterOutputType(MigrationJobOutput{})
}
