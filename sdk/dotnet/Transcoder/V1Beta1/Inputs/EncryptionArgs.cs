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
    /// Encryption settings.
    /// </summary>
    public sealed class EncryptionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configuration for AES-128 encryption.
        /// </summary>
        [Input("aes128")]
        public Input<Inputs.Aes128EncryptionArgs>? Aes128 { get; set; }

        /// <summary>
        /// 128 bit Initialization Vector (IV) represented as lowercase hexadecimal digits.
        /// </summary>
        [Input("iv", required: true)]
        public Input<string> Iv { get; set; } = null!;

        /// <summary>
        /// 128 bit encryption key represented as lowercase hexadecimal digits.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// Configuration for MPEG Common Encryption (MPEG-CENC).
        /// </summary>
        [Input("mpegCenc")]
        public Input<Inputs.MpegCommonEncryptionArgs>? MpegCenc { get; set; }

        /// <summary>
        /// Configuration for SAMPLE-AES encryption.
        /// </summary>
        [Input("sampleAes")]
        public Input<Inputs.SampleAesEncryptionArgs>? SampleAes { get; set; }

        public EncryptionArgs()
        {
        }
    }
}
