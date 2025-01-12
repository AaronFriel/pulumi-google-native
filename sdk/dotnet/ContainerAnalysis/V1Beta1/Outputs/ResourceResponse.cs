// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1Beta1.Outputs
{

    /// <summary>
    /// An entity that can have metadata. For example, a Docker image.
    /// </summary>
    [OutputType]
    public sealed class ResourceResponse
    {
        /// <summary>
        /// The unique URI of the resource. For example, `https://gcr.io/project/image@sha256:foo` for a Docker image.
        /// </summary>
        public readonly string Uri;

        [OutputConstructor]
        private ResourceResponse(string uri)
        {
            Uri = uri;
        }
    }
}
