// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.GoogleNative.DLP.V2
{
    public static class GetStoredInfoType
    {
        /// <summary>
        /// Gets a stored infoType. See https://cloud.google.com/dlp/docs/creating-stored-infotypes to learn more.
        /// </summary>
        public static Task<GetStoredInfoTypeResult> InvokeAsync(GetStoredInfoTypeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetStoredInfoTypeResult>("google-native:dlp/v2:getStoredInfoType", args ?? new GetStoredInfoTypeArgs(), options.WithVersion());

        /// <summary>
        /// Gets a stored infoType. See https://cloud.google.com/dlp/docs/creating-stored-infotypes to learn more.
        /// </summary>
        public static Output<GetStoredInfoTypeResult> Invoke(GetStoredInfoTypeInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetStoredInfoTypeResult>("google-native:dlp/v2:getStoredInfoType", args ?? new GetStoredInfoTypeInvokeArgs(), options.WithVersion());
    }


    public sealed class GetStoredInfoTypeArgs : Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("storedInfoTypeId", required: true)]
        public string StoredInfoTypeId { get; set; } = null!;

        public GetStoredInfoTypeArgs()
        {
        }
    }

    public sealed class GetStoredInfoTypeInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("storedInfoTypeId", required: true)]
        public Input<string> StoredInfoTypeId { get; set; } = null!;

        public GetStoredInfoTypeInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetStoredInfoTypeResult
    {
        /// <summary>
        /// Current version of the stored info type.
        /// </summary>
        public readonly Outputs.GooglePrivacyDlpV2StoredInfoTypeVersionResponse CurrentVersion;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Pending versions of the stored info type. Empty if no versions are pending.
        /// </summary>
        public readonly ImmutableArray<Outputs.GooglePrivacyDlpV2StoredInfoTypeVersionResponse> PendingVersions;

        [OutputConstructor]
        private GetStoredInfoTypeResult(
            Outputs.GooglePrivacyDlpV2StoredInfoTypeVersionResponse currentVersion,

            string name,

            ImmutableArray<Outputs.GooglePrivacyDlpV2StoredInfoTypeVersionResponse> pendingVersions)
        {
            CurrentVersion = currentVersion;
            Name = name;
            PendingVersions = pendingVersions;
        }
    }
}