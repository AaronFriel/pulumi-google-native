// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.AppEngine.V1Beta.Inputs
{

    /// <summary>
    /// Allows autoscaling based on Stackdriver metrics.
    /// </summary>
    public sealed class CustomMetricArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Allows filtering on the metric's fields.
        /// </summary>
        [Input("filter")]
        public Input<string>? Filter { get; set; }

        /// <summary>
        /// The name of the metric.
        /// </summary>
        [Input("metricName")]
        public Input<string>? MetricName { get; set; }

        /// <summary>
        /// May be used instead of target_utilization when an instance can handle a specific amount of work/resources and the metric value is equal to the current amount of work remaining. The autoscaler will try to keep the number of instances equal to the metric value divided by single_instance_assignment.
        /// </summary>
        [Input("singleInstanceAssignment")]
        public Input<double>? SingleInstanceAssignment { get; set; }

        /// <summary>
        /// The type of the metric. Must be a string representing a Stackdriver metric type e.g. GAGUE, DELTA_PER_SECOND, etc.
        /// </summary>
        [Input("targetType")]
        public Input<string>? TargetType { get; set; }

        /// <summary>
        /// The target value for the metric.
        /// </summary>
        [Input("targetUtilization")]
        public Input<double>? TargetUtilization { get; set; }

        public CustomMetricArgs()
        {
        }
    }
}