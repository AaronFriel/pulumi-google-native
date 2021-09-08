// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.NetworkServices.V1Beta1.Outputs
{

    [OutputType]
    public sealed class TrafficPortSelectorResponse
    {
        /// <summary>
        /// Optional. A list of ports. Can be port numbers or port range (example, [80-90] specifies all ports from 80 to 90, including 80 and 90) or named ports or * to specify all ports. If the list is empty, all ports are selected.
        /// </summary>
        public readonly ImmutableArray<string> Ports;

        [OutputConstructor]
        private TrafficPortSelectorResponse(ImmutableArray<string> ports)
        {
            Ports = ports;
        }
    }
}