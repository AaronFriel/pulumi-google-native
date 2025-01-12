// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1Alpha1.Outputs
{

    /// <summary>
    /// This represents how a particular software package may be installed on a system.
    /// </summary>
    [OutputType]
    public sealed class InstallationResponse
    {
        /// <summary>
        /// All of the places within the filesystem versions of this package have been found.
        /// </summary>
        public readonly ImmutableArray<Outputs.LocationResponse> Location;
        /// <summary>
        /// The name of the installed package.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private InstallationResponse(
            ImmutableArray<Outputs.LocationResponse> location,

            string name)
        {
            Location = location;
            Name = name;
        }
    }
}
