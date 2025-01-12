// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.NetworkServices.V1.Inputs
{

    /// <summary>
    /// A definition of a matcher that selects endpoints to which the policies should be applied.
    /// </summary>
    public sealed class EndpointMatcherArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The matcher is based on node metadata presented by xDS clients.
        /// </summary>
        [Input("metadataLabelMatcher")]
        public Input<Inputs.EndpointMatcherMetadataLabelMatcherArgs>? MetadataLabelMatcher { get; set; }

        public EndpointMatcherArgs()
        {
        }
    }
}
