// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.WebSecurityScanner.V1Alpha.Inputs
{

    /// <summary>
    /// Scan schedule configuration.
    /// </summary>
    public sealed class ScheduleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The duration of time between executions in days.
        /// </summary>
        [Input("intervalDurationDays", required: true)]
        public Input<int> IntervalDurationDays { get; set; } = null!;

        /// <summary>
        /// A timestamp indicates when the next run will be scheduled. The value is refreshed by the server after each run. If unspecified, it will default to current server time, which means the scan will be scheduled to start immediately.
        /// </summary>
        [Input("scheduleTime")]
        public Input<string>? ScheduleTime { get; set; }

        public ScheduleArgs()
        {
        }
    }
}