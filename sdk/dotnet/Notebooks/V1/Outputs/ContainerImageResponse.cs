// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.Notebooks.V1.Outputs
{

    [OutputType]
    public sealed class ContainerImageResponse
    {
        /// <summary>
        /// Required. The path to the container image repository. For example: `gcr.io/{project_id}/{image_name}`
        /// </summary>
        public readonly string Repository;
        /// <summary>
        /// The tag of the container image. If not specified, this defaults to the latest tag.
        /// </summary>
        public readonly string Tag;

        [OutputConstructor]
        private ContainerImageResponse(
            string repository,

            string tag)
        {
            Repository = repository;
            Tag = tag;
        }
    }
}