// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.SQLAdmin.V1.Outputs
{

    [OutputType]
    public sealed class InstanceFailoverReplicaResponse
    {
        /// <summary>
        /// The availability status of the failover replica. A false status indicates that the failover replica is out of sync. The primary instance can only failover to the failover replica when the status is true.
        /// </summary>
        public readonly bool Available;
        /// <summary>
        /// The name of the failover replica. If specified at instance creation, a failover replica is created for the instance. The name doesn't include the project ID.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private InstanceFailoverReplicaResponse(
            bool available,

            string name)
        {
            Available = available;
            Name = name;
        }
    }
}