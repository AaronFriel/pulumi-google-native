// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.Compute.Alpha.Inputs
{

    /// <summary>
    /// Guest OS features.
    /// </summary>
    public sealed class GuestOsFeatureArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of a supported feature. Read  Enabling guest operating system features to see a list of available options.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public GuestOsFeatureArgs()
        {
        }
    }
}