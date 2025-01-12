// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Run.V1Alpha1.Outputs
{

    /// <summary>
    /// Not supported by Cloud Run HTTPGetAction describes an action based on HTTP Get requests.
    /// </summary>
    [OutputType]
    public sealed class HTTPGetActionResponse
    {
        /// <summary>
        /// (Optional) Host name to connect to, defaults to the pod IP. You probably want to set "Host" in httpHeaders instead.
        /// </summary>
        public readonly string Host;
        /// <summary>
        /// (Optional) Custom headers to set in the request. HTTP allows repeated headers.
        /// </summary>
        public readonly ImmutableArray<Outputs.HTTPHeaderResponse> HttpHeaders;
        /// <summary>
        /// (Optional) Path to access on the HTTP server.
        /// </summary>
        public readonly string Path;
        /// <summary>
        /// (Optional) Scheme to use for connecting to the host. Defaults to HTTP.
        /// </summary>
        public readonly string Scheme;

        [OutputConstructor]
        private HTTPGetActionResponse(
            string host,

            ImmutableArray<Outputs.HTTPHeaderResponse> httpHeaders,

            string path,

            string scheme)
        {
            Host = host;
            HttpHeaders = httpHeaders;
            Path = path;
            Scheme = scheme;
        }
    }
}
