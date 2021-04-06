// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.NetworkManagement.V1Beta1.Outputs
{

    [OutputType]
    public sealed class ForwardingRuleInfoResponse
    {
        /// <summary>
        /// Name of a Compute Engine forwarding rule.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Port range defined in the forwarding rule that matches the test.
        /// </summary>
        public readonly string MatchedPortRange;
        /// <summary>
        /// Protocol defined in the forwarding rule that matches the test.
        /// </summary>
        public readonly string MatchedProtocol;
        /// <summary>
        /// Network URI. Only valid for Internal Load Balancer.
        /// </summary>
        public readonly string NetworkUri;
        /// <summary>
        /// Target type of the forwarding rule.
        /// </summary>
        public readonly string Target;
        /// <summary>
        /// URI of a Compute Engine forwarding rule.
        /// </summary>
        public readonly string Uri;
        /// <summary>
        /// VIP of the forwarding rule.
        /// </summary>
        public readonly string Vip;

        [OutputConstructor]
        private ForwardingRuleInfoResponse(
            string displayName,

            string matchedPortRange,

            string matchedProtocol,

            string networkUri,

            string target,

            string uri,

            string vip)
        {
            DisplayName = displayName;
            MatchedPortRange = matchedPortRange;
            MatchedProtocol = matchedProtocol;
            NetworkUri = networkUri;
            Target = target;
            Uri = uri;
            Vip = vip;
        }
    }
}