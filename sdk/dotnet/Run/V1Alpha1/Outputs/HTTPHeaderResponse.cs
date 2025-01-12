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
    /// Not supported by Cloud Run HTTPHeader describes a custom header to be used in HTTP probes
    /// </summary>
    [OutputType]
    public sealed class HTTPHeaderResponse
    {
        /// <summary>
        /// The header field name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The header field value
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private HTTPHeaderResponse(
            string name,

            string value)
        {
            Name = name;
            Value = value;
        }
    }
}
