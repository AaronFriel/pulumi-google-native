// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.GoogleNative.Monitoring.V3
{
    public static class GetGroup
    {
        /// <summary>
        /// Gets a single group.
        /// </summary>
        public static Task<GetGroupResult> InvokeAsync(GetGroupArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetGroupResult>("google-native:monitoring/v3:getGroup", args ?? new GetGroupArgs(), options.WithVersion());

        /// <summary>
        /// Gets a single group.
        /// </summary>
        public static Output<GetGroupResult> Invoke(GetGroupInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetGroupResult>("google-native:monitoring/v3:getGroup", args ?? new GetGroupInvokeArgs(), options.WithVersion());
    }


    public sealed class GetGroupArgs : Pulumi.InvokeArgs
    {
        [Input("groupId", required: true)]
        public string GroupId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetGroupArgs()
        {
        }
    }

    public sealed class GetGroupInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("groupId", required: true)]
        public Input<string> GroupId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetGroupInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetGroupResult
    {
        /// <summary>
        /// A user-assigned name for this group, used only for display purposes.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// The filter used to determine which monitored resources belong to this group.
        /// </summary>
        public readonly string Filter;
        /// <summary>
        /// If true, the members of this group are considered to be a cluster. The system can perform additional analysis on groups that are clusters.
        /// </summary>
        public readonly bool IsCluster;
        /// <summary>
        /// The name of this group. The format is: projects/[PROJECT_ID_OR_NUMBER]/groups/[GROUP_ID] When creating a group, this field is ignored and a new name is created consisting of the project specified in the call to CreateGroup and a unique [GROUP_ID] that is generated automatically.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The name of the group's parent, if it has one. The format is: projects/[PROJECT_ID_OR_NUMBER]/groups/[GROUP_ID] For groups with no parent, parent_name is the empty string, "".
        /// </summary>
        public readonly string ParentName;

        [OutputConstructor]
        private GetGroupResult(
            string displayName,

            string filter,

            bool isCluster,

            string name,

            string parentName)
        {
            DisplayName = displayName;
            Filter = filter;
            IsCluster = isCluster;
            Name = name;
            ParentName = parentName;
        }
    }
}