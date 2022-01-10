// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Transcoder.V1Beta1.Inputs
{

    /// <summary>
    /// Deblock preprocessing configuration.
    /// </summary>
    public sealed class DeblockArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable deblocker. The default is `false`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Set strength of the deblocker. Enter a value between 0 and 1. The higher the value, the stronger the block removal. 0 is no deblocking. The default is 0.
        /// </summary>
        [Input("strength")]
        public Input<double>? Strength { get; set; }

        public DeblockArgs()
        {
        }
    }
}