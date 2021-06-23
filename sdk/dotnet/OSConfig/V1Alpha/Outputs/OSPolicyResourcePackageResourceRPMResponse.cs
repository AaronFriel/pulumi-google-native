// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.OSConfig.V1Alpha.Outputs
{

    [OutputType]
    public sealed class OSPolicyResourcePackageResourceRPMResponse
    {
        /// <summary>
        /// Whether dependencies should also be installed. - install when false: `rpm --upgrade --replacepkgs package.rpm` - install when true: `yum -y install package.rpm` or `zypper -y install package.rpm`
        /// </summary>
        public readonly bool PullDeps;
        /// <summary>
        /// Required. An rpm package.
        /// </summary>
        public readonly Outputs.OSPolicyResourceFileResponse Source;

        [OutputConstructor]
        private OSPolicyResourcePackageResourceRPMResponse(
            bool pullDeps,

            Outputs.OSPolicyResourceFileResponse source)
        {
            PullDeps = pullDeps;
            Source = source;
        }
    }
}