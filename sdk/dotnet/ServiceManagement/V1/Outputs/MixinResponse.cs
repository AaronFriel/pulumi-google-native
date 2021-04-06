// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.ServiceManagement.V1.Outputs
{

    [OutputType]
    public sealed class MixinResponse
    {
        /// <summary>
        /// The fully qualified name of the interface which is included.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// If non-empty specifies a path under which inherited HTTP paths are rooted.
        /// </summary>
        public readonly string Root;

        [OutputConstructor]
        private MixinResponse(
            string name,

            string root)
        {
            Name = name;
            Root = root;
        }
    }
}