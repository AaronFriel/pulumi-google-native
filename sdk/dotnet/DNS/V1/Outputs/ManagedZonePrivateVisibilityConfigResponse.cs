// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DNS.V1.Outputs
{

    [OutputType]
    public sealed class ManagedZonePrivateVisibilityConfigResponse
    {
        public readonly string Kind;
        /// <summary>
        /// The list of VPC networks that can see this zone.
        /// </summary>
        public readonly ImmutableArray<Outputs.ManagedZonePrivateVisibilityConfigNetworkResponse> Networks;

        [OutputConstructor]
        private ManagedZonePrivateVisibilityConfigResponse(
            string kind,

            ImmutableArray<Outputs.ManagedZonePrivateVisibilityConfigNetworkResponse> networks)
        {
            Kind = kind;
            Networks = networks;
        }
    }
}
