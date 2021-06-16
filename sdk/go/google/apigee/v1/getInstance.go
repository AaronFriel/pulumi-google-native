// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
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
	// Required. Compute Engine location where the instance resides.
	Location string `pulumi:"location"`
	// Required. Resource ID of the instance. Values must match the regular expression `^a-z{0,30}[a-z\d]$`.
	Name string `pulumi:"name"`
	// Optional. Size of the CIDR block range that will be reserved by the instance. PAID organizations support `SLASH_16` to `SLASH_20` and defaults to `SLASH_16`. Evaluation organizations support only `SLASH_23`.
	PeeringCidrRange string `pulumi:"peeringCidrRange"`
	// Port number of the exposed Apigee endpoint.
	Port string `pulumi:"port"`
	// State of the instance. Values other than `ACTIVE` means the resource is not ready to use.
	State string `pulumi:"state"`
}