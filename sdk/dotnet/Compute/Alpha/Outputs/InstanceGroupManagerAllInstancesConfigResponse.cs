// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Outputs
{

    [OutputType]
    public sealed class InstanceGroupManagerAllInstancesConfigResponse
    {
        /// <summary>
        /// Properties for instances that are created using this instances config. You can add or modify properties using the instanceGroupManagers.patch or regionInstanceGroupManagers.patch. After setting instances_config, you must update your instances to use it; for example, you can use the applyUpdatesToInstances method.
        /// </summary>
        public readonly Outputs.InstancePropertiesPatchResponse Properties;

        [OutputConstructor]
        private InstanceGroupManagerAllInstancesConfigResponse(Outputs.InstancePropertiesPatchResponse properties)
        {
            Properties = properties;
        }
    }
}