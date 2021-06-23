// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DataCatalog.V1.Inputs
{

    /// <summary>
    /// A column within a schema. Columns can be nested inside other columns.
    /// </summary>
    public sealed class GoogleCloudDatacatalogV1ColumnSchemaArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Required. Name of the column. Must be a UTF-8 string without dots (.). The maximum size is 64 bytes.
        /// </summary>
        [Input("column")]
        public Input<string>? Column { get; set; }

        /// <summary>
        /// Optional. Description of the column. Default value is an empty string. The description must be a UTF-8 string with the maximum size of 2000 bytes.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Optional. A column's mode indicates whether values in this column are required, nullable, or repeated. Only `NULLABLE`, `REQUIRED`, and `REPEATED` values are supported. Default mode is `NULLABLE`.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        [Input("subcolumns")]
        private InputList<Inputs.GoogleCloudDatacatalogV1ColumnSchemaArgs>? _subcolumns;

        /// <summary>
        /// Optional. Schema of sub-columns. A column can have zero or more sub-columns.
        /// </summary>
        public InputList<Inputs.GoogleCloudDatacatalogV1ColumnSchemaArgs> Subcolumns
        {
            get => _subcolumns ?? (_subcolumns = new InputList<Inputs.GoogleCloudDatacatalogV1ColumnSchemaArgs>());
            set => _subcolumns = value;
        }

        /// <summary>
        /// Required. Type of the column. Must be a UTF-8 string with the maximum size of 128 bytes.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public GoogleCloudDatacatalogV1ColumnSchemaArgs()
        {
        }
    }
}