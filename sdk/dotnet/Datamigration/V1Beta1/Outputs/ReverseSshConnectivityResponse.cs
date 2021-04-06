// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.Datamigration.V1Beta1.Outputs
{

    [OutputType]
    public sealed class ReverseSshConnectivityResponse
    {
        /// <summary>
        /// The name of the virtual machine (Compute Engine) used as the bastion server for the SSH tunnel.
        /// </summary>
        public readonly string Vm;
        /// <summary>
        /// Required. The IP of the virtual machine (Compute Engine) used as the bastion server for the SSH tunnel.
        /// </summary>
        public readonly string VmIp;
        /// <summary>
        /// Required. The forwarding port of the virtual machine (Compute Engine) used as the bastion server for the SSH tunnel.
        /// </summary>
        public readonly int VmPort;
        /// <summary>
        /// The name of the VPC to peer with the Cloud SQL private network.
        /// </summary>
        public readonly string Vpc;

        [OutputConstructor]
        private ReverseSshConnectivityResponse(
            string vm,

            string vmIp,

            int vmPort,

            string vpc)
        {
            Vm = vm;
            VmIp = vmIp;
            VmPort = vmPort;
            Vpc = vpc;
        }
    }
}