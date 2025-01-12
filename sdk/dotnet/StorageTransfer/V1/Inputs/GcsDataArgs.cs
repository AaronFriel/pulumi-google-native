// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.StorageTransfer.V1.Inputs
{

    /// <summary>
    /// In a GcsData resource, an object's name is the Cloud Storage object's name and its "last modification time" refers to the object's `updated` property of Cloud Storage objects, which changes when the content or the metadata of the object is updated.
    /// </summary>
    public sealed class GcsDataArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cloud Storage bucket name. Must meet [Bucket Name Requirements](/storage/docs/naming#requirements).
        /// </summary>
        [Input("bucketName", required: true)]
        public Input<string> BucketName { get; set; } = null!;

        /// <summary>
        /// Root path to transfer objects. Must be an empty string or full path name that ends with a '/'. This field is treated as an object prefix. As such, it should generally not begin with a '/'. The root path value must meet [Object Name Requirements](/storage/docs/naming#objectnames).
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        public GcsDataArgs()
        {
        }
    }
}
