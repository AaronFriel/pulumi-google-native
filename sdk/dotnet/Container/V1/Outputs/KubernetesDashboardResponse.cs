// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Container.V1.Outputs
{

    /// <summary>
    /// Configuration for the Kubernetes Dashboard.
    /// </summary>
    [OutputType]
    public sealed class KubernetesDashboardResponse
    {
        /// <summary>
        /// Whether the Kubernetes Dashboard is enabled for this cluster.
        /// </summary>
        public readonly bool Disabled;

        [OutputConstructor]
        private KubernetesDashboardResponse(bool disabled)
        {
            Disabled = disabled;
        }
    }
}