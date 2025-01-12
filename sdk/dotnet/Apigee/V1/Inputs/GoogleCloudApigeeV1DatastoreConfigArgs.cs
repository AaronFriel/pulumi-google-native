// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Apigee.V1.Inputs
{

    /// <summary>
    /// Configuration detail for datastore
    /// </summary>
    public sealed class GoogleCloudApigeeV1DatastoreConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the Cloud Storage bucket. Required for `gcs` target_type.
        /// </summary>
        [Input("bucketName")]
        public Input<string>? BucketName { get; set; }

        /// <summary>
        /// BigQuery dataset name Required for `bigquery` target_type.
        /// </summary>
        [Input("datasetName")]
        public Input<string>? DatasetName { get; set; }

        /// <summary>
        /// Path of Cloud Storage bucket Required for `gcs` target_type.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// GCP project in which the datastore exists
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// Prefix of BigQuery table Required for `bigquery` target_type.
        /// </summary>
        [Input("tablePrefix")]
        public Input<string>? TablePrefix { get; set; }

        public GoogleCloudApigeeV1DatastoreConfigArgs()
        {
        }
    }
}
