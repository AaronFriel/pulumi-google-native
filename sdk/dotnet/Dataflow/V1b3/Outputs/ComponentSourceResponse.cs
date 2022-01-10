// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataflow.V1b3.Outputs
{

    /// <summary>
    /// Description of an interstitial value between transforms in an execution stage.
    /// </summary>
    [OutputType]
    public sealed class ComponentSourceResponse
    {
        /// <summary>
        /// Dataflow service generated name for this source.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// User name for the original user transform or collection with which this source is most closely associated.
        /// </summary>
        public readonly string OriginalTransformOrCollection;
        /// <summary>
        /// Human-readable name for this transform; may be user or system generated.
        /// </summary>
        public readonly string UserName;

        [OutputConstructor]
        private ComponentSourceResponse(
            string name,

            string originalTransformOrCollection,

            string userName)
        {
            Name = name;
            OriginalTransformOrCollection = originalTransformOrCollection;
            UserName = userName;
        }
    }
}