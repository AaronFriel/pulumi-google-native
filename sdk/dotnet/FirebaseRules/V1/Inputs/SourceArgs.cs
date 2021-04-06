// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.FirebaseRules.V1.Inputs
{

    /// <summary>
    /// `Source` is one or more `File` messages comprising a logical set of rules.
    /// </summary>
    public sealed class SourceArgs : Pulumi.ResourceArgs
    {
        [Input("files")]
        private InputList<Inputs.FileArgs>? _files;

        /// <summary>
        /// `File` set constituting the `Source` bundle.
        /// </summary>
        public InputList<Inputs.FileArgs> Files
        {
            get => _files ?? (_files = new InputList<Inputs.FileArgs>());
            set => _files = value;
        }

        public SourceArgs()
        {
        }
    }
}