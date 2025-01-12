// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.GoogleNative.Pubsub.V1
{
    public static class GetSchema
    {
        /// <summary>
        /// Gets a schema.
        /// </summary>
        public static Task<GetSchemaResult> InvokeAsync(GetSchemaArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSchemaResult>("google-native:pubsub/v1:getSchema", args ?? new GetSchemaArgs(), options.WithVersion());

        /// <summary>
        /// Gets a schema.
        /// </summary>
        public static Output<GetSchemaResult> Invoke(GetSchemaInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetSchemaResult>("google-native:pubsub/v1:getSchema", args ?? new GetSchemaInvokeArgs(), options.WithVersion());
    }


    public sealed class GetSchemaArgs : Pulumi.InvokeArgs
    {
        [Input("project")]
        public string? Project { get; set; }

        [Input("schemaId", required: true)]
        public string SchemaId { get; set; } = null!;

        [Input("view")]
        public string? View { get; set; }

        public GetSchemaArgs()
        {
        }
    }

    public sealed class GetSchemaInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("schemaId", required: true)]
        public Input<string> SchemaId { get; set; } = null!;

        [Input("view")]
        public Input<string>? View { get; set; }

        public GetSchemaInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSchemaResult
    {
        /// <summary>
        /// The definition of the schema. This should contain a string representing the full definition of the schema that is a valid schema definition of the type specified in `type`.
        /// </summary>
        public readonly string Definition;
        /// <summary>
        /// Name of the schema. Format is `projects/{project}/schemas/{schema}`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The type of the schema definition.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetSchemaResult(
            string definition,

            string name,

            string type)
        {
            Definition = definition;
            Name = name;
            Type = type;
        }
    }
}
