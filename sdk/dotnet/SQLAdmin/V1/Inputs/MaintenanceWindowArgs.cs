// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.SQLAdmin.V1.Inputs
{

    /// <summary>
    /// Maintenance window. This specifies when a Cloud SQL instance is restarted for system maintenance purposes.
    /// </summary>
    public sealed class MaintenanceWindowArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// day of week (1-7), starting on Monday.
        /// </summary>
        [Input("day")]
        public Input<int>? Day { get; set; }

        /// <summary>
        /// hour of day - 0 to 23.
        /// </summary>
        [Input("hour")]
        public Input<int>? Hour { get; set; }

        /// <summary>
        /// This is always **sql#maintenanceWindow**.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// Maintenance timing setting: **canary** (Earlier) or **stable** (Later). [Learn more] (https://cloud.google.com/sql/docs/mysql/instance-settings#maintenance-timing-2ndgen).
        /// </summary>
        [Input("updateTrack")]
        public Input<Pulumi.GoogleNative.SQLAdmin.V1.MaintenanceWindowUpdateTrack>? UpdateTrack { get; set; }

        public MaintenanceWindowArgs()
        {
        }
    }
}