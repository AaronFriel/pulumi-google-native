// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.FirebaseHosting.V1Beta1.Outputs
{

    /// <summary>
    /// A configured rewrite that directs requests to a Cloud Run service. If the Cloud Run service does not exist when setting or updating your Firebase Hosting configuration, then the request fails. Any errors from the Cloud Run service are passed to the end user (for example, if you delete a service, any requests directed to that service receive a `404` error).
    /// </summary>
    [OutputType]
    public sealed class CloudRunRewriteResponse
    {
        /// <summary>
        /// Optional. User-provided region where the Cloud Run service is hosted. Defaults to `us-central1` if not supplied.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// User-defined ID of the Cloud Run service.
        /// </summary>
        public readonly string ServiceId;

        [OutputConstructor]
        private CloudRunRewriteResponse(
            string region,

            string serviceId)
        {
            Region = region;
            ServiceId = serviceId;
        }
    }
}
