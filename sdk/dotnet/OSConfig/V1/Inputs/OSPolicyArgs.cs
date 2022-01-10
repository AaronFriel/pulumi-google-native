// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.OSConfig.V1.Inputs
{

    /// <summary>
    /// An OS policy defines the desired state configuration for a VM.
    /// </summary>
    public sealed class OSPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// This flag determines the OS policy compliance status when none of the resource groups within the policy are applicable for a VM. Set this value to `true` if the policy needs to be reported as compliant even if the policy has nothing to validate or enforce.
        /// </summary>
        [Input("allowNoResourceGroupMatch")]
        public Input<bool>? AllowNoResourceGroupMatch { get; set; }

        /// <summary>
        /// Policy description. Length of the description is limited to 1024 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The id of the OS policy with the following restrictions: * Must contain only lowercase letters, numbers, and hyphens. * Must start with a letter. * Must be between 1-63 characters. * Must end with a number or a letter. * Must be unique within the assignment.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// Policy mode
        /// </summary>
        [Input("mode", required: true)]
        public Input<Pulumi.GoogleNative.OSConfig.V1.OSPolicyMode> Mode { get; set; } = null!;

        [Input("resourceGroups", required: true)]
        private InputList<Inputs.OSPolicyResourceGroupArgs>? _resourceGroups;

        /// <summary>
        /// List of resource groups for the policy. For a particular VM, resource groups are evaluated in the order specified and the first resource group that is applicable is selected and the rest are ignored. If none of the resource groups are applicable for a VM, the VM is considered to be non-compliant w.r.t this policy. This behavior can be toggled by the flag `allow_no_resource_group_match`
        /// </summary>
        public InputList<Inputs.OSPolicyResourceGroupArgs> ResourceGroups
        {
            get => _resourceGroups ?? (_resourceGroups = new InputList<Inputs.OSPolicyResourceGroupArgs>());
            set => _resourceGroups = value;
        }

        public OSPolicyArgs()
        {
        }
    }
}