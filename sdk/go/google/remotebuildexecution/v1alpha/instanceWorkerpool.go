// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates a new worker pool with a specified size and configuration. Returns a long running operation which contains a worker pool on completion. While the long running operation is in progress, any call to `GetWorkerPool` returns a worker pool in state `CREATING`.
type InstanceWorkerpool struct {
	pulumi.CustomResourceState

	// The autoscale policy to apply on a pool.
	Autoscale GoogleDevtoolsRemotebuildexecutionAdminV1alphaAutoscaleResponseOutput `pulumi:"autoscale"`
	// Channel specifies the release channel of the pool.
	Channel pulumi.StringOutput `pulumi:"channel"`
	// WorkerPool resource name formatted as: `projects/[PROJECT_ID]/instances/[INSTANCE_ID]/workerpools/[POOL_ID]`. name should not be populated when creating a worker pool since it is provided in the `poolId` field.
	Name pulumi.StringOutput `pulumi:"name"`
	// State of the worker pool.
	State pulumi.StringOutput `pulumi:"state"`
	// Specifies the properties, such as machine type and disk size, used for creating workers in a worker pool.
	WorkerConfig GoogleDevtoolsRemotebuildexecutionAdminV1alphaWorkerConfigResponseOutput `pulumi:"workerConfig"`
	// The desired number of workers in the worker pool. Must be a value between 0 and 15000.
	WorkerCount pulumi.StringOutput `pulumi:"workerCount"`
}

// NewInstanceWorkerpool registers a new resource with the given unique name, arguments, and options.
func NewInstanceWorkerpool(ctx *pulumi.Context,
	name string, args *InstanceWorkerpoolArgs, opts ...pulumi.ResourceOption) (*InstanceWorkerpool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstancesId == nil {
		return nil, errors.New("invalid value for required argument 'InstancesId'")
	}
	if args.ProjectsId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectsId'")
	}
	if args.WorkerpoolsId == nil {
		return nil, errors.New("invalid value for required argument 'WorkerpoolsId'")
	}
	var resource InstanceWorkerpool
	err := ctx.RegisterResource("google-cloud:remotebuildexecution/v1alpha:InstanceWorkerpool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceWorkerpool gets an existing InstanceWorkerpool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceWorkerpool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceWorkerpoolState, opts ...pulumi.ResourceOption) (*InstanceWorkerpool, error) {
	var resource InstanceWorkerpool
	err := ctx.ReadResource("google-cloud:remotebuildexecution/v1alpha:InstanceWorkerpool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceWorkerpool resources.
type instanceWorkerpoolState struct {
	// The autoscale policy to apply on a pool.
	Autoscale *GoogleDevtoolsRemotebuildexecutionAdminV1alphaAutoscaleResponse `pulumi:"autoscale"`
	// Channel specifies the release channel of the pool.
	Channel *string `pulumi:"channel"`
	// WorkerPool resource name formatted as: `projects/[PROJECT_ID]/instances/[INSTANCE_ID]/workerpools/[POOL_ID]`. name should not be populated when creating a worker pool since it is provided in the `poolId` field.
	Name *string `pulumi:"name"`
	// State of the worker pool.
	State *string `pulumi:"state"`
	// Specifies the properties, such as machine type and disk size, used for creating workers in a worker pool.
	WorkerConfig *GoogleDevtoolsRemotebuildexecutionAdminV1alphaWorkerConfigResponse `pulumi:"workerConfig"`
	// The desired number of workers in the worker pool. Must be a value between 0 and 15000.
	WorkerCount *string `pulumi:"workerCount"`
}

type InstanceWorkerpoolState struct {
	// The autoscale policy to apply on a pool.
	Autoscale GoogleDevtoolsRemotebuildexecutionAdminV1alphaAutoscaleResponsePtrInput
	// Channel specifies the release channel of the pool.
	Channel pulumi.StringPtrInput
	// WorkerPool resource name formatted as: `projects/[PROJECT_ID]/instances/[INSTANCE_ID]/workerpools/[POOL_ID]`. name should not be populated when creating a worker pool since it is provided in the `poolId` field.
	Name pulumi.StringPtrInput
	// State of the worker pool.
	State pulumi.StringPtrInput
	// Specifies the properties, such as machine type and disk size, used for creating workers in a worker pool.
	WorkerConfig GoogleDevtoolsRemotebuildexecutionAdminV1alphaWorkerConfigResponsePtrInput
	// The desired number of workers in the worker pool. Must be a value between 0 and 15000.
	WorkerCount pulumi.StringPtrInput
}

func (InstanceWorkerpoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceWorkerpoolState)(nil)).Elem()
}

type instanceWorkerpoolArgs struct {
	// The autoscale policy to apply on a pool.
	Autoscale *GoogleDevtoolsRemotebuildexecutionAdminV1alphaAutoscale `pulumi:"autoscale"`
	// Channel specifies the release channel of the pool.
	Channel     *string `pulumi:"channel"`
	InstancesId string  `pulumi:"instancesId"`
	// WorkerPool resource name formatted as: `projects/[PROJECT_ID]/instances/[INSTANCE_ID]/workerpools/[POOL_ID]`. name should not be populated when creating a worker pool since it is provided in the `poolId` field.
	Name *string `pulumi:"name"`
	// Resource name of the instance in which to create the new worker pool. Format: `projects/[PROJECT_ID]/instances/[INSTANCE_ID]`.
	Parent *string `pulumi:"parent"`
	// ID of the created worker pool. A valid pool ID must: be 6-50 characters long, contain only lowercase letters, digits, hyphens and underscores, start with a lowercase letter, and end with a lowercase letter or a digit.
	PoolId     *string `pulumi:"poolId"`
	ProjectsId string  `pulumi:"projectsId"`
	// State of the worker pool.
	State *string `pulumi:"state"`
	// Specifies the properties, such as machine type and disk size, used for creating workers in a worker pool.
	WorkerConfig *GoogleDevtoolsRemotebuildexecutionAdminV1alphaWorkerConfig `pulumi:"workerConfig"`
	// The desired number of workers in the worker pool. Must be a value between 0 and 15000.
	WorkerCount   *string `pulumi:"workerCount"`
	WorkerpoolsId string  `pulumi:"workerpoolsId"`
}

// The set of arguments for constructing a InstanceWorkerpool resource.
type InstanceWorkerpoolArgs struct {
	// The autoscale policy to apply on a pool.
	Autoscale GoogleDevtoolsRemotebuildexecutionAdminV1alphaAutoscalePtrInput
	// Channel specifies the release channel of the pool.
	Channel     pulumi.StringPtrInput
	InstancesId pulumi.StringInput
	// WorkerPool resource name formatted as: `projects/[PROJECT_ID]/instances/[INSTANCE_ID]/workerpools/[POOL_ID]`. name should not be populated when creating a worker pool since it is provided in the `poolId` field.
	Name pulumi.StringPtrInput
	// Resource name of the instance in which to create the new worker pool. Format: `projects/[PROJECT_ID]/instances/[INSTANCE_ID]`.
	Parent pulumi.StringPtrInput
	// ID of the created worker pool. A valid pool ID must: be 6-50 characters long, contain only lowercase letters, digits, hyphens and underscores, start with a lowercase letter, and end with a lowercase letter or a digit.
	PoolId     pulumi.StringPtrInput
	ProjectsId pulumi.StringInput
	// State of the worker pool.
	State pulumi.StringPtrInput
	// Specifies the properties, such as machine type and disk size, used for creating workers in a worker pool.
	WorkerConfig GoogleDevtoolsRemotebuildexecutionAdminV1alphaWorkerConfigPtrInput
	// The desired number of workers in the worker pool. Must be a value between 0 and 15000.
	WorkerCount   pulumi.StringPtrInput
	WorkerpoolsId pulumi.StringInput
}

func (InstanceWorkerpoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceWorkerpoolArgs)(nil)).Elem()
}

type InstanceWorkerpoolInput interface {
	pulumi.Input

	ToInstanceWorkerpoolOutput() InstanceWorkerpoolOutput
	ToInstanceWorkerpoolOutputWithContext(ctx context.Context) InstanceWorkerpoolOutput
}

func (*InstanceWorkerpool) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceWorkerpool)(nil))
}

func (i *InstanceWorkerpool) ToInstanceWorkerpoolOutput() InstanceWorkerpoolOutput {
	return i.ToInstanceWorkerpoolOutputWithContext(context.Background())
}

func (i *InstanceWorkerpool) ToInstanceWorkerpoolOutputWithContext(ctx context.Context) InstanceWorkerpoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceWorkerpoolOutput)
}

type InstanceWorkerpoolOutput struct {
	*pulumi.OutputState
}

func (InstanceWorkerpoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceWorkerpool)(nil))
}

func (o InstanceWorkerpoolOutput) ToInstanceWorkerpoolOutput() InstanceWorkerpoolOutput {
	return o
}

func (o InstanceWorkerpoolOutput) ToInstanceWorkerpoolOutputWithContext(ctx context.Context) InstanceWorkerpoolOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(InstanceWorkerpoolOutput{})
}