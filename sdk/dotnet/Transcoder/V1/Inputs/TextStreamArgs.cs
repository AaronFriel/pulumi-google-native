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
    /// Encoding of a text stream. For example, closed captions or subtitles.
    /// </summary>
    public sealed class TextStreamArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The codec for this text stream. The default is `webvtt`. Supported text codecs: - `srt` - `ttml` - `cea608` - `cea708` - `webvtt`
        /// </summary>
        [Input("codec")]
        public Input<string>? Codec { get; set; }

        [Input("mapping")]
        private InputList<Inputs.TextMappingArgs>? _mapping;

        /// <summary>
        /// The mapping for the `Job.edit_list` atoms with text `EditAtom.inputs`.
        /// </summary>
        public InputList<Inputs.TextMappingArgs> Mapping
        {
            get => _mapping ?? (_mapping = new InputList<Inputs.TextMappingArgs>());
            set => _mapping = value;
        }

        public TextStreamArgs()
        {
        }
    }
}
