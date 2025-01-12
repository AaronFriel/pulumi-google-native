// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Transcoder.V1Beta1.Outputs
{

    /// <summary>
    /// Ad break.
    /// </summary>
    [OutputType]
    public sealed class AdBreakResponse
    {
        /// <summary>
        /// Start time in seconds for the ad break, relative to the output file timeline. The default is `0s`.
        /// </summary>
        public readonly string StartTimeOffset;

        [OutputConstructor]
        private AdBreakResponse(string startTimeOffset)
        {
            StartTimeOffset = startTimeOffset;
        }
    }
}
