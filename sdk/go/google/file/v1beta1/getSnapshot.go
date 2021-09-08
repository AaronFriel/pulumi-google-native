// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the details of a specific snapshot.
func LookupSnapshot(ctx *pulumi.Context, args *LookupSnapshotArgs, opts ...pulumi.InvokeOption) (*LookupSnapshotResult, error) {
	var rv LookupSnapshotResult
	err := ctx.Invoke("google-native:file/v1beta1:getSnapshot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSnapshotArgs struct {
	InstanceId string  `pulumi:"instanceId"`
	Location   string  `pulumi:"location"`
	Project    *string `pulumi:"project"`
	SnapshotId string  `pulumi:"snapshotId"`
}

type LookupSnapshotResult struct {
	// The time when the snapshot was created.
	CreateTime string `pulumi:"createTime"`
	// A description of the snapshot with 2048 characters or less. Requests with longer descriptions will be rejected.
	Description string `pulumi:"description"`
	// The amount of bytes needed to allocate a full copy of the snapshot content
	FilesystemUsedBytes string `pulumi:"filesystemUsedBytes"`
	// Resource labels to represent user provided metadata.
	Labels map[string]string `pulumi:"labels"`
	// The resource name of the snapshot, in the format `projects/{project_id}/locations/{location_id}/instances/{instance_id}/snapshots/{snapshot_id}`.
	Name string `pulumi:"name"`
	// The snapshot state.
	State string `pulumi:"state"`
}