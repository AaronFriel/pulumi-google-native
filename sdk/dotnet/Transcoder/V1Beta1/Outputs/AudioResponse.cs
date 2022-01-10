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
    /// Audio preprocessing configuration.
    /// </summary>
    [OutputType]
    public sealed class AudioResponse
    {
        /// <summary>
        /// Enable boosting high frequency components. The default is `false`.
        /// </summary>
        public readonly bool HighBoost;
        /// <summary>
        /// Enable boosting low frequency components. The default is `false`.
        /// </summary>
        public readonly bool LowBoost;
        /// <summary>
        /// Specify audio loudness normalization in loudness units relative to full scale (LUFS). Enter a value between -24 and 0 (the default), where: * -24 is the Advanced Television Systems Committee (ATSC A/85) standard * -23 is the EU R128 broadcast standard * -19 is the prior standard for online mono audio * -18 is the ReplayGain standard * -16 is the prior standard for stereo audio * -14 is the new online audio standard recommended by Spotify, as well as Amazon Echo * 0 disables normalization
        /// </summary>
        public readonly double Lufs;

        [OutputConstructor]
        private AudioResponse(
            bool highBoost,

            bool lowBoost,

            double lufs)
        {
            HighBoost = highBoost;
            LowBoost = lowBoost;
            Lufs = lufs;
        }
    }
}