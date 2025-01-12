// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Container.V1Beta1.Outputs
{

    /// <summary>
    /// Represents the Maintenance exclusion option.
    /// </summary>
    [OutputType]
    public sealed class MaintenanceExclusionOptionsResponse
    {
        /// <summary>
        /// Scope specifies the upgrade scope which upgrades are blocked by the exclusion.
        /// </summary>
        public readonly string Scope;

        [OutputConstructor]
        private MaintenanceExclusionOptionsResponse(string scope)
        {
            Scope = scope;
        }
    }
}
