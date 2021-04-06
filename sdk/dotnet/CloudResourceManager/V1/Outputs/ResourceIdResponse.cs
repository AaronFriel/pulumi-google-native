// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.CloudResourceManager.V1.Outputs
{

    [OutputType]
    public sealed class ResourceIdResponse
    {
        /// <summary>
        /// The resource type this id is for. At present, the valid types are: "organization", "folder", and "project".
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private ResourceIdResponse(string type)
        {
            Type = type;
        }
    }
}