// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataproc.V1Beta2.Outputs
{

    /// <summary>
    /// Dataproc job status.
    /// </summary>
    [OutputType]
    public sealed class JobStatusResponse
    {
        /// <summary>
        /// Optional Job state details, such as an error description if the state is ERROR.
        /// </summary>
        public readonly string Details;
        /// <summary>
        /// A state message specifying the overall job state.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The time when this state was entered.
        /// </summary>
        public readonly string StateStartTime;
        /// <summary>
        /// Additional state information, which includes status reported by the agent.
        /// </summary>
        public readonly string Substate;

        [OutputConstructor]
        private JobStatusResponse(
            string details,

            string state,

            string stateStartTime,

            string substate)
        {
            Details = details;
            State = state;
            StateStartTime = stateStartTime;
            Substate = substate;
        }
    }
}