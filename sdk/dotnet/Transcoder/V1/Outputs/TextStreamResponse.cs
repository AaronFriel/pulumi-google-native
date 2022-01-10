// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Transcoder.V1.Outputs
{

    /// <summary>
    /// Encoding of a text stream. For example, closed captions or subtitles.
    /// </summary>
    [OutputType]
    public sealed class TextStreamResponse
    {
        /// <summary>
        /// The codec for this text stream. The default is `webvtt`. Supported text codecs: - `srt` - `ttml` - `cea608` - `cea708` - `webvtt`
        /// </summary>
        public readonly string Codec;
        /// <summary>
        /// The mapping for the `Job.edit_list` atoms with text `EditAtom.inputs`.
        /// </summary>
        public readonly ImmutableArray<Outputs.TextMappingResponse> Mapping;

        [OutputConstructor]
        private TextStreamResponse(
            string codec,

            ImmutableArray<Outputs.TextMappingResponse> mapping)
        {
            Codec = codec;
            Mapping = mapping;
        }
    }
}