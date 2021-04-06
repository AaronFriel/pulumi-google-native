// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.Dataproc.V1Beta2.Inputs
{

    /// <summary>
    /// Job scheduling options.
    /// </summary>
    public sealed class JobSchedulingArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. Maximum number of times per hour a driver may be restarted as a result of driver terminating with non-zero code before job is reported failed.A job may be reported as thrashing if driver exits with non-zero code 4 times within 10 minute window.Maximum value is 10.
        /// </summary>
        [Input("maxFailuresPerHour")]
        public Input<int>? MaxFailuresPerHour { get; set; }

        /// <summary>
        /// Optional. Maximum number of times in total a driver may be restarted as a result of driver exiting with non-zero code before job is reported failed. Maximum value is 240.
        /// </summary>
        [Input("maxFailuresTotal")]
        public Input<int>? MaxFailuresTotal { get; set; }

        public JobSchedulingArgs()
        {
        }
    }
}