// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V3.Outputs
{

    /// <summary>
    /// Represents configurations for a test case.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDialogflowCxV3TestConfigResponse
    {
        /// <summary>
        /// Flow name. If not set, default start flow is assumed. Format: `projects//locations//agents//flows/`.
        /// </summary>
        public readonly string Flow;
        /// <summary>
        /// Session parameters to be compared when calculating differences.
        /// </summary>
        public readonly ImmutableArray<string> TrackingParameters;

        [OutputConstructor]
        private GoogleCloudDialogflowCxV3TestConfigResponse(
            string flow,

            ImmutableArray<string> trackingParameters)
        {
            Flow = flow;
            TrackingParameters = trackingParameters;
        }
    }
}
