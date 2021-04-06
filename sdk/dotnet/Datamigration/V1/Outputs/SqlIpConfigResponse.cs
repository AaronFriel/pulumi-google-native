// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.Datamigration.V1.Outputs
{

    [OutputType]
    public sealed class SqlIpConfigResponse
    {
        /// <summary>
        /// The list of external networks that are allowed to connect to the instance using the IP. See https://en.wikipedia.org/wiki/CIDR_notation#CIDR_notation, also known as 'slash' notation (e.g. `192.168.100.0/24`).
        /// </summary>
        public readonly ImmutableArray<Outputs.SqlAclEntryResponse> AuthorizedNetworks;
        /// <summary>
        /// Whether the instance should be assigned an IPv4 address or not.
        /// </summary>
        public readonly bool EnableIpv4;
        /// <summary>
        /// The resource link for the VPC network from which the Cloud SQL instance is accessible for private IP. For example, `/projects/myProject/global/networks/default`. This setting can be updated, but it cannot be removed after it is set.
        /// </summary>
        public readonly string PrivateNetwork;
        /// <summary>
        /// Whether SSL connections over IP should be enforced or not.
        /// </summary>
        public readonly bool RequireSsl;

        [OutputConstructor]
        private SqlIpConfigResponse(
            ImmutableArray<Outputs.SqlAclEntryResponse> authorizedNetworks,

            bool enableIpv4,

            string privateNetwork,

            bool requireSsl)
        {
            AuthorizedNetworks = authorizedNetworks;
            EnableIpv4 = enableIpv4;
            PrivateNetwork = privateNetwork;
            RequireSsl = requireSsl;
        }
    }
}