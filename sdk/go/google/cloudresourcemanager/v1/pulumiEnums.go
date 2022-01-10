// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The log type that this config enables.
type AuditLogConfigLogType string

const (
	// Default case. Should never be this.
	AuditLogConfigLogTypeLogTypeUnspecified = AuditLogConfigLogType("LOG_TYPE_UNSPECIFIED")
	// Admin reads. Example: CloudIAM getIamPolicy
	AuditLogConfigLogTypeAdminRead = AuditLogConfigLogType("ADMIN_READ")
	// Data writes. Example: CloudSQL Users create
	AuditLogConfigLogTypeDataWrite = AuditLogConfigLogType("DATA_WRITE")
	// Data reads. Example: CloudSQL Users list
	AuditLogConfigLogTypeDataRead = AuditLogConfigLogType("DATA_READ")
)

func (AuditLogConfigLogType) ElementType() reflect.Type {
	return reflect.TypeOf((*AuditLogConfigLogType)(nil)).Elem()
}

func (e AuditLogConfigLogType) ToAuditLogConfigLogTypeOutput() AuditLogConfigLogTypeOutput {
	return pulumi.ToOutput(e).(AuditLogConfigLogTypeOutput)
}

func (e AuditLogConfigLogType) ToAuditLogConfigLogTypeOutputWithContext(ctx context.Context) AuditLogConfigLogTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AuditLogConfigLogTypeOutput)
}

func (e AuditLogConfigLogType) ToAuditLogConfigLogTypePtrOutput() AuditLogConfigLogTypePtrOutput {
	return e.ToAuditLogConfigLogTypePtrOutputWithContext(context.Background())
}

func (e AuditLogConfigLogType) ToAuditLogConfigLogTypePtrOutputWithContext(ctx context.Context) AuditLogConfigLogTypePtrOutput {
	return AuditLogConfigLogType(e).ToAuditLogConfigLogTypeOutputWithContext(ctx).ToAuditLogConfigLogTypePtrOutputWithContext(ctx)
}

func (e AuditLogConfigLogType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AuditLogConfigLogType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AuditLogConfigLogType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AuditLogConfigLogType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AuditLogConfigLogTypeOutput struct{ *pulumi.OutputState }

func (AuditLogConfigLogTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuditLogConfigLogType)(nil)).Elem()
}

func (o AuditLogConfigLogTypeOutput) ToAuditLogConfigLogTypeOutput() AuditLogConfigLogTypeOutput {
	return o
}

func (o AuditLogConfigLogTypeOutput) ToAuditLogConfigLogTypeOutputWithContext(ctx context.Context) AuditLogConfigLogTypeOutput {
	return o
}

func (o AuditLogConfigLogTypeOutput) ToAuditLogConfigLogTypePtrOutput() AuditLogConfigLogTypePtrOutput {
	return o.ToAuditLogConfigLogTypePtrOutputWithContext(context.Background())
}

func (o AuditLogConfigLogTypeOutput) ToAuditLogConfigLogTypePtrOutputWithContext(ctx context.Context) AuditLogConfigLogTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AuditLogConfigLogType) *AuditLogConfigLogType {
		return &v
	}).(AuditLogConfigLogTypePtrOutput)
}

func (o AuditLogConfigLogTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AuditLogConfigLogTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AuditLogConfigLogType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AuditLogConfigLogTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AuditLogConfigLogTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AuditLogConfigLogType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AuditLogConfigLogTypePtrOutput struct{ *pulumi.OutputState }

func (AuditLogConfigLogTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuditLogConfigLogType)(nil)).Elem()
}

func (o AuditLogConfigLogTypePtrOutput) ToAuditLogConfigLogTypePtrOutput() AuditLogConfigLogTypePtrOutput {
	return o
}

func (o AuditLogConfigLogTypePtrOutput) ToAuditLogConfigLogTypePtrOutputWithContext(ctx context.Context) AuditLogConfigLogTypePtrOutput {
	return o
}

func (o AuditLogConfigLogTypePtrOutput) Elem() AuditLogConfigLogTypeOutput {
	return o.ApplyT(func(v *AuditLogConfigLogType) AuditLogConfigLogType {
		if v != nil {
			return *v
		}
		var ret AuditLogConfigLogType
		return ret
	}).(AuditLogConfigLogTypeOutput)
}

func (o AuditLogConfigLogTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AuditLogConfigLogTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AuditLogConfigLogType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// AuditLogConfigLogTypeInput is an input type that accepts AuditLogConfigLogTypeArgs and AuditLogConfigLogTypeOutput values.
// You can construct a concrete instance of `AuditLogConfigLogTypeInput` via:
//
//          AuditLogConfigLogTypeArgs{...}
type AuditLogConfigLogTypeInput interface {
	pulumi.Input

	ToAuditLogConfigLogTypeOutput() AuditLogConfigLogTypeOutput
	ToAuditLogConfigLogTypeOutputWithContext(context.Context) AuditLogConfigLogTypeOutput
}

var auditLogConfigLogTypePtrType = reflect.TypeOf((**AuditLogConfigLogType)(nil)).Elem()

type AuditLogConfigLogTypePtrInput interface {
	pulumi.Input

	ToAuditLogConfigLogTypePtrOutput() AuditLogConfigLogTypePtrOutput
	ToAuditLogConfigLogTypePtrOutputWithContext(context.Context) AuditLogConfigLogTypePtrOutput
}

type auditLogConfigLogTypePtr string

func AuditLogConfigLogTypePtr(v string) AuditLogConfigLogTypePtrInput {
	return (*auditLogConfigLogTypePtr)(&v)
}

func (*auditLogConfigLogTypePtr) ElementType() reflect.Type {
	return auditLogConfigLogTypePtrType
}

func (in *auditLogConfigLogTypePtr) ToAuditLogConfigLogTypePtrOutput() AuditLogConfigLogTypePtrOutput {
	return pulumi.ToOutput(in).(AuditLogConfigLogTypePtrOutput)
}

func (in *auditLogConfigLogTypePtr) ToAuditLogConfigLogTypePtrOutputWithContext(ctx context.Context) AuditLogConfigLogTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AuditLogConfigLogTypePtrOutput)
}

// The Project lifecycle state. Read-only.
type ProjectLifecycleState string

const (
	// Unspecified state. This is only used/useful for distinguishing unset values.
	ProjectLifecycleStateLifecycleStateUnspecified = ProjectLifecycleState("LIFECYCLE_STATE_UNSPECIFIED")
	// The normal and active state.
	ProjectLifecycleStateActive = ProjectLifecycleState("ACTIVE")
	// The project has been marked for deletion by the user (by invoking DeleteProject) or by the system (Google Cloud Platform). This can generally be reversed by invoking UndeleteProject.
	ProjectLifecycleStateDeleteRequested = ProjectLifecycleState("DELETE_REQUESTED")
	// This lifecycle state is no longer used and not returned by the API.
	ProjectLifecycleStateDeleteInProgress = ProjectLifecycleState("DELETE_IN_PROGRESS")
)

func (ProjectLifecycleState) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectLifecycleState)(nil)).Elem()
}

func (e ProjectLifecycleState) ToProjectLifecycleStateOutput() ProjectLifecycleStateOutput {
	return pulumi.ToOutput(e).(ProjectLifecycleStateOutput)
}

func (e ProjectLifecycleState) ToProjectLifecycleStateOutputWithContext(ctx context.Context) ProjectLifecycleStateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ProjectLifecycleStateOutput)
}

func (e ProjectLifecycleState) ToProjectLifecycleStatePtrOutput() ProjectLifecycleStatePtrOutput {
	return e.ToProjectLifecycleStatePtrOutputWithContext(context.Background())
}

func (e ProjectLifecycleState) ToProjectLifecycleStatePtrOutputWithContext(ctx context.Context) ProjectLifecycleStatePtrOutput {
	return ProjectLifecycleState(e).ToProjectLifecycleStateOutputWithContext(ctx).ToProjectLifecycleStatePtrOutputWithContext(ctx)
}

func (e ProjectLifecycleState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProjectLifecycleState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProjectLifecycleState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ProjectLifecycleState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ProjectLifecycleStateOutput struct{ *pulumi.OutputState }

func (ProjectLifecycleStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectLifecycleState)(nil)).Elem()
}

func (o ProjectLifecycleStateOutput) ToProjectLifecycleStateOutput() ProjectLifecycleStateOutput {
	return o
}

func (o ProjectLifecycleStateOutput) ToProjectLifecycleStateOutputWithContext(ctx context.Context) ProjectLifecycleStateOutput {
	return o
}

func (o ProjectLifecycleStateOutput) ToProjectLifecycleStatePtrOutput() ProjectLifecycleStatePtrOutput {
	return o.ToProjectLifecycleStatePtrOutputWithContext(context.Background())
}

func (o ProjectLifecycleStateOutput) ToProjectLifecycleStatePtrOutputWithContext(ctx context.Context) ProjectLifecycleStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProjectLifecycleState) *ProjectLifecycleState {
		return &v
	}).(ProjectLifecycleStatePtrOutput)
}

func (o ProjectLifecycleStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ProjectLifecycleStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProjectLifecycleState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ProjectLifecycleStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProjectLifecycleStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProjectLifecycleState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ProjectLifecycleStatePtrOutput struct{ *pulumi.OutputState }

func (ProjectLifecycleStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectLifecycleState)(nil)).Elem()
}

func (o ProjectLifecycleStatePtrOutput) ToProjectLifecycleStatePtrOutput() ProjectLifecycleStatePtrOutput {
	return o
}

func (o ProjectLifecycleStatePtrOutput) ToProjectLifecycleStatePtrOutputWithContext(ctx context.Context) ProjectLifecycleStatePtrOutput {
	return o
}

func (o ProjectLifecycleStatePtrOutput) Elem() ProjectLifecycleStateOutput {
	return o.ApplyT(func(v *ProjectLifecycleState) ProjectLifecycleState {
		if v != nil {
			return *v
		}
		var ret ProjectLifecycleState
		return ret
	}).(ProjectLifecycleStateOutput)
}

func (o ProjectLifecycleStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProjectLifecycleStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ProjectLifecycleState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ProjectLifecycleStateInput is an input type that accepts ProjectLifecycleStateArgs and ProjectLifecycleStateOutput values.
// You can construct a concrete instance of `ProjectLifecycleStateInput` via:
//
//          ProjectLifecycleStateArgs{...}
type ProjectLifecycleStateInput interface {
	pulumi.Input

	ToProjectLifecycleStateOutput() ProjectLifecycleStateOutput
	ToProjectLifecycleStateOutputWithContext(context.Context) ProjectLifecycleStateOutput
}

var projectLifecycleStatePtrType = reflect.TypeOf((**ProjectLifecycleState)(nil)).Elem()

type ProjectLifecycleStatePtrInput interface {
	pulumi.Input

	ToProjectLifecycleStatePtrOutput() ProjectLifecycleStatePtrOutput
	ToProjectLifecycleStatePtrOutputWithContext(context.Context) ProjectLifecycleStatePtrOutput
}

type projectLifecycleStatePtr string

func ProjectLifecycleStatePtr(v string) ProjectLifecycleStatePtrInput {
	return (*projectLifecycleStatePtr)(&v)
}

func (*projectLifecycleStatePtr) ElementType() reflect.Type {
	return projectLifecycleStatePtrType
}

func (in *projectLifecycleStatePtr) ToProjectLifecycleStatePtrOutput() ProjectLifecycleStatePtrOutput {
	return pulumi.ToOutput(in).(ProjectLifecycleStatePtrOutput)
}

func (in *projectLifecycleStatePtr) ToProjectLifecycleStatePtrOutputWithContext(ctx context.Context) ProjectLifecycleStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ProjectLifecycleStatePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AuditLogConfigLogTypeInput)(nil)).Elem(), AuditLogConfigLogType("LOG_TYPE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*AuditLogConfigLogTypePtrInput)(nil)).Elem(), AuditLogConfigLogType("LOG_TYPE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectLifecycleStateInput)(nil)).Elem(), ProjectLifecycleState("LIFECYCLE_STATE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectLifecycleStatePtrInput)(nil)).Elem(), ProjectLifecycleState("LIFECYCLE_STATE_UNSPECIFIED"))
	pulumi.RegisterOutputType(AuditLogConfigLogTypeOutput{})
	pulumi.RegisterOutputType(AuditLogConfigLogTypePtrOutput{})
	pulumi.RegisterOutputType(ProjectLifecycleStateOutput{})
	pulumi.RegisterOutputType(ProjectLifecycleStatePtrOutput{})
}