// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataproc.V1.Outputs
{

    /// <summary>
    /// The cluster's GKE config.
    /// </summary>
    [OutputType]
    public sealed class GkeClusterConfigResponse
    {
        /// <summary>
        /// Optional. A target for the deployment.
        /// </summary>
        public readonly Outputs.NamespacedGkeDeploymentTargetResponse NamespacedGkeDeploymentTarget;

        [OutputConstructor]
        private GkeClusterConfigResponse(Outputs.NamespacedGkeDeploymentTargetResponse namespacedGkeDeploymentTarget)
        {
            NamespacedGkeDeploymentTarget = namespacedGkeDeploymentTarget;
        }
    }
}
