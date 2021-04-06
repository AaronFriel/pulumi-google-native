// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.Container.V1Beta1.Inputs
{

    /// <summary>
    /// AutoUpgradeOptions defines the set of options for the user to control how the Auto Upgrades will proceed.
    /// </summary>
    public sealed class AutoUpgradeOptionsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// [Output only] This field is set when upgrades are about to commence with the approximate start time for the upgrades, in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.
        /// </summary>
        [Input("autoUpgradeStartTime")]
        public Input<string>? AutoUpgradeStartTime { get; set; }

        /// <summary>
        /// [Output only] This field is set when upgrades are about to commence with the description of the upgrade.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        public AutoUpgradeOptionsArgs()
        {
        }
    }
}