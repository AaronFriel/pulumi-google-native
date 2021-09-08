// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the specified ArchiveDeployment.
func LookupArchiveDeployment(ctx *pulumi.Context, args *LookupArchiveDeploymentArgs, opts ...pulumi.InvokeOption) (*LookupArchiveDeploymentResult, error) {
	var rv LookupArchiveDeploymentResult
	err := ctx.Invoke("google-native:apigee/v1:getArchiveDeployment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupArchiveDeploymentArgs struct {
	ArchiveDeploymentId string `pulumi:"archiveDeploymentId"`
	EnvironmentId       string `pulumi:"environmentId"`
	OrganizationId      string `pulumi:"organizationId"`
}

type LookupArchiveDeploymentResult struct {
	// The time at which the Archive Deployment was created in milliseconds since the epoch.
	CreatedAt string `pulumi:"createdAt"`
	// Input only. The Google Cloud Storage signed URL returned from GenerateUploadUrl and used to upload the Archive zip file.
	GcsUri string `pulumi:"gcsUri"`
	// User-supplied key-value pairs used to organize ArchiveDeployments. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: \p{Ll}\p{Lo}{0,62} Label values must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given store.
	Labels map[string]string `pulumi:"labels"`
	// Name of the Archive Deployment in the following format: `organizations/{org}/environments/{env}/archiveDeployments/{id}`.
	Name string `pulumi:"name"`
	// A reference to the LRO that created this Archive Deployment in the following format: `organizations/{org}/operations/{id}`
	Operation string `pulumi:"operation"`
	// The time at which the Archive Deployment was updated in milliseconds since the epoch.
	UpdatedAt string `pulumi:"updatedAt"`
}