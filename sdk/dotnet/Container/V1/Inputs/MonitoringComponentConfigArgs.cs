// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Container.V1.Inputs
{

    /// <summary>
    /// MonitoringComponentConfig is cluster monitoring component configuration.
    /// </summary>
    public sealed class MonitoringComponentConfigArgs : Pulumi.ResourceArgs
    {
        [Input("enableComponents")]
        private InputList<Pulumi.GoogleNative.Container.V1.MonitoringComponentConfigEnableComponentsItem>? _enableComponents;

        /// <summary>
        /// Select components to collect metrics. An empty set would disable all monitoring.
        /// </summary>
        public InputList<Pulumi.GoogleNative.Container.V1.MonitoringComponentConfigEnableComponentsItem> EnableComponents
        {
            get => _enableComponents ?? (_enableComponents = new InputList<Pulumi.GoogleNative.Container.V1.MonitoringComponentConfigEnableComponentsItem>());
            set => _enableComponents = value;
        }

        public MonitoringComponentConfigArgs()
        {
        }
    }
}
