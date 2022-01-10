// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.SQLAdmin.V1Beta4.Inputs
{

    /// <summary>
    /// Reference to another Cloud SQL instance.
    /// </summary>
    public sealed class InstanceReferenceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Cloud SQL instance being referenced. This does not include the project ID.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The project ID of the Cloud SQL instance being referenced. The default is the same project ID as the instance references it.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The region of the Cloud SQL instance being referenced.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public InstanceReferenceArgs()
        {
        }
    }
}