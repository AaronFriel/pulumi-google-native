// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1Beta1.Inputs
{

    /// <summary>
    /// This represents a particular channel of distribution for a given package. E.g., Debian's jessie-backports dpkg mirror.
    /// </summary>
    public sealed class DistributionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The CPU architecture for which packages in this distribution channel were built.
        /// </summary>
        [Input("architecture")]
        public Input<Pulumi.GoogleNative.ContainerAnalysis.V1Beta1.DistributionArchitecture>? Architecture { get; set; }

        /// <summary>
        /// The cpe_uri in [CPE format](https://cpe.mitre.org/specification/) denoting the package manager version distributing a package.
        /// </summary>
        [Input("cpeUri", required: true)]
        public Input<string> CpeUri { get; set; } = null!;

        /// <summary>
        /// The distribution channel-specific description of this package.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The latest available version of this package in this distribution channel.
        /// </summary>
        [Input("latestVersion")]
        public Input<Inputs.VersionArgs>? LatestVersion { get; set; }

        /// <summary>
        /// A freeform string denoting the maintainer of this package.
        /// </summary>
        [Input("maintainer")]
        public Input<string>? Maintainer { get; set; }

        /// <summary>
        /// The distribution channel-specific homepage for this package.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public DistributionArgs()
        {
        }
    }
}
