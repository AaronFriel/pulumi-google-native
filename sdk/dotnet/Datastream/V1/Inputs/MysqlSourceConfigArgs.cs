// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Datastream.V1.Inputs
{

    /// <summary>
    /// MySQL source configuration
    /// </summary>
    public sealed class MysqlSourceConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// MySQL objects to exclude from the stream.
        /// </summary>
        [Input("excludeObjects")]
        public Input<Inputs.MysqlRdbmsArgs>? ExcludeObjects { get; set; }

        /// <summary>
        /// MySQL objects to retrieve from the source.
        /// </summary>
        [Input("includeObjects")]
        public Input<Inputs.MysqlRdbmsArgs>? IncludeObjects { get; set; }

        public MysqlSourceConfigArgs()
        {
        }
    }
}
