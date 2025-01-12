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
    /// Encoding of an input file such as an audio, video, or text track. Elementary streams must be packaged before mapping and sharing between different output formats.
    /// </summary>
    [OutputType]
    public sealed class ElementaryStreamResponse
    {
        /// <summary>
        /// Encoding of an audio stream.
        /// </summary>
        public readonly Outputs.AudioStreamResponse AudioStream;
        /// <summary>
        /// A unique key for this elementary stream.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// Encoding of a text stream. For example, closed captions or subtitles.
        /// </summary>
        public readonly Outputs.TextStreamResponse TextStream;
        /// <summary>
        /// Encoding of a video stream.
        /// </summary>
        public readonly Outputs.VideoStreamResponse VideoStream;

        [OutputConstructor]
        private ElementaryStreamResponse(
            Outputs.AudioStreamResponse audioStream,

            string key,

            Outputs.TextStreamResponse textStream,

            Outputs.VideoStreamResponse videoStream)
        {
            AudioStream = audioStream;
            Key = key;
            TextStream = textStream;
            VideoStream = videoStream;
        }
    }
}
