// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.OSConfig.V1.Inputs
{

    /// <summary>
    /// A resource that manages the state of a file.
    /// </summary>
    public sealed class OSPolicyResourceFileResourceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A a file with this content. The size of the content is limited to 1024 characters.
        /// </summary>
        [Input("content")]
        public Input<string>? Content { get; set; }

        /// <summary>
        /// A remote or local source.
        /// </summary>
        [Input("file")]
        public Input<Inputs.OSPolicyResourceFileArgs>? File { get; set; }

        /// <summary>
        /// The absolute path of the file within the VM.
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        /// <summary>
        /// Consists of three octal digits which represent, in order, the permissions of the owner, group, and other users for the file (similarly to the numeric mode used in the linux chmod utility). Each digit represents a three bit number with the 4 bit corresponding to the read permissions, the 2 bit corresponds to the write bit, and the one bit corresponds to the execute permission. Default behavior is 755. Below are some examples of permissions and their associated values: read, write, and execute: 7 read and execute: 5 read and write: 6 read only: 4
        /// </summary>
        [Input("permissions")]
        public Input<string>? Permissions { get; set; }

        /// <summary>
        /// Desired state of the file.
        /// </summary>
        [Input("state", required: true)]
        public Input<Pulumi.GoogleNative.OSConfig.V1.OSPolicyResourceFileResourceState> State { get; set; } = null!;

        public OSPolicyResourceFileResourceArgs()
        {
        }
    }
}