// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a debug session.
func LookupDebugSession(ctx *pulumi.Context, args *LookupDebugSessionArgs, opts ...pulumi.InvokeOption) (*LookupDebugSessionResult, error) {
	var rv LookupDebugSessionResult
	err := ctx.Invoke("google-native:apigee/v1:getDebugSession", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDebugSessionArgs struct {
	ApiId          string `pulumi:"apiId"`
	DebugsessionId string `pulumi:"debugsessionId"`
	EnvironmentId  string `pulumi:"environmentId"`
	OrganizationId string `pulumi:"organizationId"`
	RevisionId     string `pulumi:"revisionId"`
}

type LookupDebugSessionResult struct {
	// Optional. The number of request to be traced. Min = 1, Max = 15, Default = 10.
	Count int `pulumi:"count"`
	// Optional. A conditional statement which is evaluated against the request message to determine if it should be traced. Syntax matches that of on API Proxy bundle flow Condition.
	Filter string `pulumi:"filter"`
	// A unique ID for this DebugSession.
	Name string `pulumi:"name"`
	// Optional. The time in seconds after which this DebugSession should end. This value will override the value in query param, if both are provided.
	Timeout string `pulumi:"timeout"`
	// Optional. The maximum number of bytes captured from the response payload. Min = 0, Max = 5120, Default = 5120.
	Tracesize int `pulumi:"tracesize"`
	// Optional. The length of time, in seconds, that this debug session is valid, starting from when it's received in the control plane. Min = 1, Max = 15, Default = 10.
	Validity int `pulumi:"validity"`
}

func LookupDebugSessionOutput(ctx *pulumi.Context, args LookupDebugSessionOutputArgs, opts ...pulumi.InvokeOption) LookupDebugSessionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDebugSessionResult, error) {
			args := v.(LookupDebugSessionArgs)
			r, err := LookupDebugSession(ctx, &args, opts...)
			return *r, err
		}).(LookupDebugSessionResultOutput)
}

type LookupDebugSessionOutputArgs struct {
	ApiId          pulumi.StringInput `pulumi:"apiId"`
	DebugsessionId pulumi.StringInput `pulumi:"debugsessionId"`
	EnvironmentId  pulumi.StringInput `pulumi:"environmentId"`
	OrganizationId pulumi.StringInput `pulumi:"organizationId"`
	RevisionId     pulumi.StringInput `pulumi:"revisionId"`
}

func (LookupDebugSessionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDebugSessionArgs)(nil)).Elem()
}

type LookupDebugSessionResultOutput struct{ *pulumi.OutputState }

func (LookupDebugSessionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDebugSessionResult)(nil)).Elem()
}

func (o LookupDebugSessionResultOutput) ToLookupDebugSessionResultOutput() LookupDebugSessionResultOutput {
	return o
}

func (o LookupDebugSessionResultOutput) ToLookupDebugSessionResultOutputWithContext(ctx context.Context) LookupDebugSessionResultOutput {
	return o
}

// Optional. The number of request to be traced. Min = 1, Max = 15, Default = 10.
func (o LookupDebugSessionResultOutput) Count() pulumi.IntOutput {
	return o.ApplyT(func(v LookupDebugSessionResult) int { return v.Count }).(pulumi.IntOutput)
}

// Optional. A conditional statement which is evaluated against the request message to determine if it should be traced. Syntax matches that of on API Proxy bundle flow Condition.
func (o LookupDebugSessionResultOutput) Filter() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDebugSessionResult) string { return v.Filter }).(pulumi.StringOutput)
}

// A unique ID for this DebugSession.
func (o LookupDebugSessionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDebugSessionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional. The time in seconds after which this DebugSession should end. This value will override the value in query param, if both are provided.
func (o LookupDebugSessionResultOutput) Timeout() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDebugSessionResult) string { return v.Timeout }).(pulumi.StringOutput)
}

// Optional. The maximum number of bytes captured from the response payload. Min = 0, Max = 5120, Default = 5120.
func (o LookupDebugSessionResultOutput) Tracesize() pulumi.IntOutput {
	return o.ApplyT(func(v LookupDebugSessionResult) int { return v.Tracesize }).(pulumi.IntOutput)
}

// Optional. The length of time, in seconds, that this debug session is valid, starting from when it's received in the control plane. Min = 1, Max = 15, Default = 10.
func (o LookupDebugSessionResultOutput) Validity() pulumi.IntOutput {
	return o.ApplyT(func(v LookupDebugSessionResult) int { return v.Validity }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDebugSessionResultOutput{})
}