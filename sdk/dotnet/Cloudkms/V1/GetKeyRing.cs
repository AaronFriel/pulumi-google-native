// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.GoogleNative.Cloudkms.V1
{
    public static class GetKeyRing
    {
        /// <summary>
        /// Returns metadata for a given KeyRing.
        /// </summary>
        public static Task<GetKeyRingResult> InvokeAsync(GetKeyRingArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetKeyRingResult>("google-native:cloudkms/v1:getKeyRing", args ?? new GetKeyRingArgs(), options.WithVersion());

        /// <summary>
        /// Returns metadata for a given KeyRing.
        /// </summary>
        public static Output<GetKeyRingResult> Invoke(GetKeyRingInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetKeyRingResult>("google-native:cloudkms/v1:getKeyRing", args ?? new GetKeyRingInvokeArgs(), options.WithVersion());
    }


    public sealed class GetKeyRingArgs : Pulumi.InvokeArgs
    {
        [Input("keyRingId", required: true)]
        public string KeyRingId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetKeyRingArgs()
        {
        }
    }

    public sealed class GetKeyRingInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("keyRingId", required: true)]
        public Input<string> KeyRingId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetKeyRingInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetKeyRingResult
    {
        /// <summary>
        /// The time at which this KeyRing was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The resource name for the KeyRing in the format `projects/*/locations/*/keyRings/*`.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetKeyRingResult(
            string createTime,

            string name)
        {
            CreateTime = createTime;
            Name = name;
        }
    }
}