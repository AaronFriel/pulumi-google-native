// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Apigee.V1
{
    /// <summary>
    /// Create a Datastore for an org
    /// </summary>
    [GoogleNativeResourceType("google-native:apigee/v1:Datastore")]
    public partial class Datastore : Pulumi.CustomResource
    {
        /// <summary>
        /// Datastore create time, in milliseconds since the epoch of 1970-01-01T00:00:00Z
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Datastore Configurations.
        /// </summary>
        [Output("datastoreConfig")]
        public Output<Outputs.GoogleCloudApigeeV1DatastoreConfigResponse> DatastoreConfig { get; private set; } = null!;

        /// <summary>
        /// Required. Display name in UI
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Datastore last update time, in milliseconds since the epoch of 1970-01-01T00:00:00Z
        /// </summary>
        [Output("lastUpdateTime")]
        public Output<string> LastUpdateTime { get; private set; } = null!;

        /// <summary>
        /// Organization that the datastore belongs to
        /// </summary>
        [Output("org")]
        public Output<string> Org { get; private set; } = null!;

        /// <summary>
        /// Resource link of Datastore. Example: `/organizations/{org}/analytics/datastores/{uuid}`
        /// </summary>
        [Output("self")]
        public Output<string> Self { get; private set; } = null!;

        /// <summary>
        /// Destination storage type. Supported types `gcs` or `bigquery`.
        /// </summary>
        [Output("targetType")]
        public Output<string> TargetType { get; private set; } = null!;


        /// <summary>
        /// Create a Datastore resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Datastore(string name, DatastoreArgs args, CustomResourceOptions? options = null)
            : base("google-native:apigee/v1:Datastore", name, args ?? new DatastoreArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Datastore(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:apigee/v1:Datastore", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Datastore resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Datastore Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Datastore(name, id, options);
        }
    }

    public sealed class DatastoreArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Datastore Configurations.
        /// </summary>
        [Input("datastoreConfig")]
        public Input<Inputs.GoogleCloudApigeeV1DatastoreConfigArgs>? DatastoreConfig { get; set; }

        /// <summary>
        /// Required. Display name in UI
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("organizationId", required: true)]
        public Input<string> OrganizationId { get; set; } = null!;

        /// <summary>
        /// Destination storage type. Supported types `gcs` or `bigquery`.
        /// </summary>
        [Input("targetType")]
        public Input<string>? TargetType { get; set; }

        public DatastoreArgs()
        {
        }
    }
}