// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the details for an Apigee runtime instance. **Note:** Not supported for Apigee hybrid.
func LookupInstance(ctx *pulumi.Context, args *LookupInstanceArgs, opts ...pulumi.InvokeOption) (*LookupInstanceResult, error) {
	var rv LookupInstanceResult
	err := ctx.Invoke("google-native:apigee/v1:getInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInstanceArgs struct {
	InstanceId     string `pulumi:"instanceId"`
	OrganizationId string `pulumi:"organizationId"`
}

type LookupInstanceResult struct {
	// Time the instance was created in milliseconds since epoch.
	CreatedAt string `pulumi:"createdAt"`
	// Optional. Description of the instance.
	Description string `pulumi:"description"`
	// Customer Managed Encryption Key (CMEK) used for disk and volume encryption. Required for Apigee paid subscriptions only. Use the following format: `projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)`
	DiskEncryptionKeyName string `pulumi:"diskEncryptionKeyName"`
	// Optional. Display name for the instance.
	DisplayName string `pulumi:"displayName"`
	// Internal hostname or IP address of the Apigee endpoint used by clients to connect to the service.
	Host string `pulumi:"host"`
	// Time the instance was last modified in milliseconds since epoch.
	LastModifiedAt string `pulumi:"lastModifiedAt"`
	// Compute Engine location where the instance resides.
	Location string `pulumi:"location"`
	// Resource ID of the instance. Values must match the regular expression `^a-z{0,30}[a-z\d]$`.
	Name string `pulumi:"name"`
	// Optional. Size of the CIDR block range that will be reserved by the instance. PAID organizations support `SLASH_16` to `SLASH_20` and defaults to `SLASH_16`. Evaluation organizations support only `SLASH_23`.
	PeeringCidrRange string `pulumi:"peeringCidrRange"`
	// Port number of the exposed Apigee endpoint.
	Port string `pulumi:"port"`
	// Version of the runtime system running in the instance. The runtime system is the set of components that serve the API Proxy traffic in your Environments.
	RuntimeVersion string `pulumi:"runtimeVersion"`
	// State of the instance. Values other than `ACTIVE` means the resource is not ready to use.
	State string `pulumi:"state"`
}

func LookupInstanceOutput(ctx *pulumi.Context, args LookupInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupInstanceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInstanceResult, error) {
			args := v.(LookupInstanceArgs)
			r, err := LookupInstance(ctx, &args, opts...)
			return *r, err
		}).(LookupInstanceResultOutput)
}

type LookupInstanceOutputArgs struct {
	InstanceId     pulumi.StringInput `pulumi:"instanceId"`
	OrganizationId pulumi.StringInput `pulumi:"organizationId"`
}

func (LookupInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceArgs)(nil)).Elem()
}

type LookupInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceResult)(nil)).Elem()
}

func (o LookupInstanceResultOutput) ToLookupInstanceResultOutput() LookupInstanceResultOutput {
	return o
}

func (o LookupInstanceResultOutput) ToLookupInstanceResultOutputWithContext(ctx context.Context) LookupInstanceResultOutput {
	return o
}

// Time the instance was created in milliseconds since epoch.
func (o LookupInstanceResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// Optional. Description of the instance.
func (o LookupInstanceResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Description }).(pulumi.StringOutput)
}

// Customer Managed Encryption Key (CMEK) used for disk and volume encryption. Required for Apigee paid subscriptions only. Use the following format: `projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)`
func (o LookupInstanceResultOutput) DiskEncryptionKeyName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.DiskEncryptionKeyName }).(pulumi.StringOutput)
}

// Optional. Display name for the instance.
func (o LookupInstanceResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Internal hostname or IP address of the Apigee endpoint used by clients to connect to the service.
func (o LookupInstanceResultOutput) Host() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Host }).(pulumi.StringOutput)
}

// Time the instance was last modified in milliseconds since epoch.
func (o LookupInstanceResultOutput) LastModifiedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.LastModifiedAt }).(pulumi.StringOutput)
}

// Compute Engine location where the instance resides.
func (o LookupInstanceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource ID of the instance. Values must match the regular expression `^a-z{0,30}[a-z\d]$`.
func (o LookupInstanceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional. Size of the CIDR block range that will be reserved by the instance. PAID organizations support `SLASH_16` to `SLASH_20` and defaults to `SLASH_16`. Evaluation organizations support only `SLASH_23`.
func (o LookupInstanceResultOutput) PeeringCidrRange() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.PeeringCidrRange }).(pulumi.StringOutput)
}

// Port number of the exposed Apigee endpoint.
func (o LookupInstanceResultOutput) Port() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Port }).(pulumi.StringOutput)
}

// Version of the runtime system running in the instance. The runtime system is the set of components that serve the API Proxy traffic in your Environments.
func (o LookupInstanceResultOutput) RuntimeVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.RuntimeVersion }).(pulumi.StringOutput)
}

// State of the instance. Values other than `ACTIVE` means the resource is not ready to use.
func (o LookupInstanceResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.State }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInstanceResultOutput{})
}