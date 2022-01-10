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
    /// Time window specified for daily maintenance operations.
    /// </summary>
    [OutputType]
    public sealed class DailyMaintenanceWindowResponse
    {
        /// <summary>
        /// [Output only] Duration of the time window, automatically chosen to be smallest possible in the given scenario. Duration will be in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) format "PTnHnMnS".
        /// </summary>
        public readonly string Duration;
        /// <summary>
        /// Time within the maintenance window to start the maintenance operations. Time format should be in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) format "HH:MM", where HH : [00-23] and MM : [00-59] GMT.
        /// </summary>
        public readonly string StartTime;

        [OutputConstructor]
        private DailyMaintenanceWindowResponse(
            string duration,

            string startTime)
        {
            Duration = duration;
            StartTime = startTime;
        }
    }
}