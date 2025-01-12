// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Run.V1.Inputs
{

    /// <summary>
    /// ContainerPort represents a network port in a single container.
    /// </summary>
    public sealed class ContainerPortArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Optional) Port number the container listens on. This must be a valid port number, 0 &lt; x &lt; 65536.
        /// </summary>
        [Input("containerPort")]
        public Input<int>? ContainerPort { get; set; }

        /// <summary>
        /// (Optional) If specified, used to specify which protocol to use. Allowed values are "http1" and "h2c".
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// (Optional) Protocol for port. Must be "TCP". Defaults to "TCP".
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        public ContainerPortArgs()
        {
        }
    }
}
