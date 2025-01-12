// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.GoogleNative.Orgpolicy.V2
{
    public static class GetPolicy
    {
        /// <summary>
        /// Gets a `Policy` on a resource. If no `Policy` is set on the resource, NOT_FOUND is returned. The `etag` value can be used with `UpdatePolicy()` to update a `Policy` during read-modify-write.
        /// </summary>
        public static Task<GetPolicyResult> InvokeAsync(GetPolicyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPolicyResult>("google-native:orgpolicy/v2:getPolicy", args ?? new GetPolicyArgs(), options.WithVersion());

        /// <summary>
        /// Gets a `Policy` on a resource. If no `Policy` is set on the resource, NOT_FOUND is returned. The `etag` value can be used with `UpdatePolicy()` to update a `Policy` during read-modify-write.
        /// </summary>
        public static Output<GetPolicyResult> Invoke(GetPolicyInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetPolicyResult>("google-native:orgpolicy/v2:getPolicy", args ?? new GetPolicyInvokeArgs(), options.WithVersion());
    }


    public sealed class GetPolicyArgs : Pulumi.InvokeArgs
    {
        [Input("policyId", required: true)]
        public string PolicyId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetPolicyArgs()
        {
        }
    }

    public sealed class GetPolicyInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("policyId", required: true)]
        public Input<string> PolicyId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetPolicyInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPolicyResult
    {
        /// <summary>
        /// Immutable. The resource name of the Policy. Must be one of the following forms, where constraint_name is the name of the constraint which this Policy configures: * `projects/{project_number}/policies/{constraint_name}` * `folders/{folder_id}/policies/{constraint_name}` * `organizations/{organization_id}/policies/{constraint_name}` For example, "projects/123/policies/compute.disableSerialPortAccess". Note: `projects/{project_id}/policies/{constraint_name}` is also an acceptable name for API requests, but responses will return the name using the equivalent project number.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Basic information about the Organization Policy.
        /// </summary>
        public readonly Outputs.GoogleCloudOrgpolicyV2PolicySpecResponse Spec;

        [OutputConstructor]
        private GetPolicyResult(
            string name,

            Outputs.GoogleCloudOrgpolicyV2PolicySpecResponse spec)
        {
            Name = name;
            Spec = spec;
        }
    }
}
