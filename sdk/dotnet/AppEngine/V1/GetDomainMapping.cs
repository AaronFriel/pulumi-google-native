// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.AppEngine.V1
{
    public static class GetDomainMapping
    {
        /// <summary>
        /// Gets the specified domain mapping.
        /// </summary>
        public static Task<GetDomainMappingResult> InvokeAsync(GetDomainMappingArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDomainMappingResult>("google-native:appengine/v1:getDomainMapping", args ?? new GetDomainMappingArgs(), options.WithVersion());
    }


    public sealed class GetDomainMappingArgs : Pulumi.InvokeArgs
    {
        [Input("appId", required: true)]
        public string AppId { get; set; } = null!;

        [Input("domainMappingId", required: true)]
        public string DomainMappingId { get; set; } = null!;

        public GetDomainMappingArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDomainMappingResult
    {
        /// <summary>
        /// Full path to the DomainMapping resource in the API. Example: apps/myapp/domainMapping/example.com.@OutputOnly
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The resource records required to configure this domain mapping. These records must be added to the domain's DNS configuration in order to serve the application via this domain mapping.@OutputOnly
        /// </summary>
        public readonly ImmutableArray<Outputs.ResourceRecordResponse> ResourceRecords;
        /// <summary>
        /// SSL configuration for this domain. If unconfigured, this domain will not serve with SSL.
        /// </summary>
        public readonly Outputs.SslSettingsResponse SslSettings;

        [OutputConstructor]
        private GetDomainMappingResult(
            string name,

            ImmutableArray<Outputs.ResourceRecordResponse> resourceRecords,

            Outputs.SslSettingsResponse sslSettings)
        {
            Name = name;
            ResourceRecords = resourceRecords;
            SslSettings = sslSettings;
        }
    }
}