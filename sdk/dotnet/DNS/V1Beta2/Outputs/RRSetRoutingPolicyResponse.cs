// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DNS.V1Beta2.Outputs
{

    [OutputType]
    public sealed class RRSetRoutingPolicyResponse
    {
        public readonly Outputs.RRSetRoutingPolicyGeoPolicyResponse Geo;
        public readonly Outputs.RRSetRoutingPolicyGeoPolicyResponse GeoPolicy;
        public readonly string Kind;
        public readonly Outputs.RRSetRoutingPolicyWrrPolicyResponse Wrr;
        public readonly Outputs.RRSetRoutingPolicyWrrPolicyResponse WrrPolicy;

        [OutputConstructor]
        private RRSetRoutingPolicyResponse(
            Outputs.RRSetRoutingPolicyGeoPolicyResponse geo,

            Outputs.RRSetRoutingPolicyGeoPolicyResponse geoPolicy,

            string kind,

            Outputs.RRSetRoutingPolicyWrrPolicyResponse wrr,

            Outputs.RRSetRoutingPolicyWrrPolicyResponse wrrPolicy)
        {
            Geo = geo;
            GeoPolicy = geoPolicy;
            Kind = kind;
            Wrr = wrr;
            WrrPolicy = wrrPolicy;
        }
    }
}