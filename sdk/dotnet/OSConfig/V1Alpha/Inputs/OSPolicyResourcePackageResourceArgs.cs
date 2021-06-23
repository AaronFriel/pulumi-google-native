// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.OSConfig.V1Alpha.Inputs
{

    /// <summary>
    /// A resource that manages a system package.
    /// </summary>
    public sealed class OSPolicyResourcePackageResourceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A package managed by Apt.
        /// </summary>
        [Input("apt")]
        public Input<Inputs.OSPolicyResourcePackageResourceAPTArgs>? Apt { get; set; }

        /// <summary>
        /// A deb package file.
        /// </summary>
        [Input("deb")]
        public Input<Inputs.OSPolicyResourcePackageResourceDebArgs>? Deb { get; set; }

        /// <summary>
        /// Required. The desired state the agent should maintain for this package.
        /// </summary>
        [Input("desiredState")]
        public Input<Pulumi.GoogleNative.OSConfig.V1Alpha.OSPolicyResourcePackageResourceDesiredState>? DesiredState { get; set; }

        /// <summary>
        /// A package managed by GooGet.
        /// </summary>
        [Input("googet")]
        public Input<Inputs.OSPolicyResourcePackageResourceGooGetArgs>? Googet { get; set; }

        /// <summary>
        /// An MSI package.
        /// </summary>
        [Input("msi")]
        public Input<Inputs.OSPolicyResourcePackageResourceMSIArgs>? Msi { get; set; }

        /// <summary>
        /// An rpm package file.
        /// </summary>
        [Input("rpm")]
        public Input<Inputs.OSPolicyResourcePackageResourceRPMArgs>? Rpm { get; set; }

        /// <summary>
        /// A package managed by YUM.
        /// </summary>
        [Input("yum")]
        public Input<Inputs.OSPolicyResourcePackageResourceYUMArgs>? Yum { get; set; }

        /// <summary>
        /// A package managed by Zypper.
        /// </summary>
        [Input("zypper")]
        public Input<Inputs.OSPolicyResourcePackageResourceZypperArgs>? Zypper { get; set; }

        public OSPolicyResourcePackageResourceArgs()
        {
        }
    }
}