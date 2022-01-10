// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Monitoring.V3.Inputs
{

    /// <summary>
    /// A DistributionCut defines a TimeSeries and thresholds used for measuring good service and total service. The TimeSeries must have ValueType = DISTRIBUTION and MetricKind = DELTA or MetricKind = CUMULATIVE. The computed good_service will be the estimated count of values in the Distribution that fall within the specified min and max.
    /// </summary>
    public sealed class DistributionCutArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A monitoring filter (https://cloud.google.com/monitoring/api/v3/filters) specifying a TimeSeries aggregating values. Must have ValueType = DISTRIBUTION and MetricKind = DELTA or MetricKind = CUMULATIVE.
        /// </summary>
        [Input("distributionFilter")]
        public Input<string>? DistributionFilter { get; set; }

        /// <summary>
        /// Range of values considered "good." For a one-sided range, set one bound to an infinite value.
        /// </summary>
        [Input("range")]
        public Input<Inputs.GoogleMonitoringV3RangeArgs>? Range { get; set; }

        public DistributionCutArgs()
        {
        }
    }
}