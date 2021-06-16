// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v4

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves specified tenant.
func LookupTenant(ctx *pulumi.Context, args *LookupTenantArgs, opts ...pulumi.InvokeOption) (*LookupTenantResult, error) {
	var rv LookupTenantResult
	err := ctx.Invoke("google-native:jobs/v4:getTenant", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTenantArgs struct {
	Project  string `pulumi:"project"`
	TenantId string `pulumi:"tenantId"`
}

type LookupTenantResult struct {
	// Required. Client side tenant identifier, used to uniquely identify the tenant. The maximum number of allowed characters is 255.
	ExternalId string `pulumi:"externalId"`
	// Required during tenant update. The resource name for a tenant. This is generated by the service when a tenant is created. The format is "projects/{project_id}/tenants/{tenant_id}", for example, "projects/foo/tenants/bar".
	Name string `pulumi:"name"`
}