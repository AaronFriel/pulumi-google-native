// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Transcoder.V1.Inputs
{

    /// <summary>
    /// Video stream resource.
    /// </summary>
    public sealed class VideoStreamArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// H264 codec settings.
        /// </summary>
        [Input("h264")]
        public Input<Inputs.H264CodecSettingsArgs>? H264 { get; set; }

        /// <summary>
        /// H265 codec settings.
        /// </summary>
        [Input("h265")]
        public Input<Inputs.H265CodecSettingsArgs>? H265 { get; set; }

        /// <summary>
        /// VP9 codec settings.
        /// </summary>
        [Input("vp9")]
        public Input<Inputs.Vp9CodecSettingsArgs>? Vp9 { get; set; }

        public VideoStreamArgs()
        {
        }
    }
}
