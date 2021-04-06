// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.Genomics.V1Alpha2.Inputs
{

    /// <summary>
    /// The Docker execuctor specification.
    /// </summary>
    public sealed class DockerExecutorArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Required. The command or newline delimited script to run. The command string will be executed within a bash shell. If the command exits with a non-zero exit code, output parameter de-localization will be skipped and the pipeline operation's `error` field will be populated. Maximum command string length is 16384.
        /// </summary>
        [Input("cmd")]
        public Input<string>? Cmd { get; set; }

        /// <summary>
        /// Required. Image name from either Docker Hub or Google Container Registry. Users that run pipelines must have READ access to the image.
        /// </summary>
        [Input("imageName")]
        public Input<string>? ImageName { get; set; }

        public DockerExecutorArgs()
        {
        }
    }
}