// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.GoogleNative.BigQueryReservation.V1
{
    public static class GetCapacityCommitment
    {
        /// <summary>
        /// Returns information about the capacity commitment.
        /// </summary>
        public static Task<GetCapacityCommitmentResult> InvokeAsync(GetCapacityCommitmentArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCapacityCommitmentResult>("google-native:bigqueryreservation/v1:getCapacityCommitment", args ?? new GetCapacityCommitmentArgs(), options.WithVersion());

        /// <summary>
        /// Returns information about the capacity commitment.
        /// </summary>
        public static Output<GetCapacityCommitmentResult> Invoke(GetCapacityCommitmentInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetCapacityCommitmentResult>("google-native:bigqueryreservation/v1:getCapacityCommitment", args ?? new GetCapacityCommitmentInvokeArgs(), options.WithVersion());
    }


    public sealed class GetCapacityCommitmentArgs : Pulumi.InvokeArgs
    {
        [Input("capacityCommitmentId", required: true)]
        public string CapacityCommitmentId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetCapacityCommitmentArgs()
        {
        }
    }

    public sealed class GetCapacityCommitmentInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("capacityCommitmentId", required: true)]
        public Input<string> CapacityCommitmentId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetCapacityCommitmentInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetCapacityCommitmentResult
    {
        /// <summary>
        /// The end of the current commitment period. It is applicable only for ACTIVE capacity commitments.
        /// </summary>
        public readonly string CommitmentEndTime;
        /// <summary>
        /// The start of the current commitment period. It is applicable only for ACTIVE capacity commitments.
        /// </summary>
        public readonly string CommitmentStartTime;
        /// <summary>
        /// For FAILED commitment plan, provides the reason of failure.
        /// </summary>
        public readonly Outputs.StatusResponse FailureStatus;
        /// <summary>
        /// The resource name of the capacity commitment, e.g., `projects/myproject/locations/US/capacityCommitments/123`
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Capacity commitment commitment plan.
        /// </summary>
        public readonly string Plan;
        /// <summary>
        /// The plan this capacity commitment is converted to after commitment_end_time passes. Once the plan is changed, committed period is extended according to commitment plan. Only applicable for ANNUAL and TRIAL commitments.
        /// </summary>
        public readonly string RenewalPlan;
        /// <summary>
        /// Number of slots in this commitment.
        /// </summary>
        public readonly string SlotCount;
        /// <summary>
        /// State of the commitment.
        /// </summary>
        public readonly string State;

        [OutputConstructor]
        private GetCapacityCommitmentResult(
            string commitmentEndTime,

            string commitmentStartTime,

            Outputs.StatusResponse failureStatus,

            string name,

            string plan,

            string renewalPlan,

            string slotCount,

            string state)
        {
            CommitmentEndTime = commitmentEndTime;
            CommitmentStartTime = commitmentStartTime;
            FailureStatus = failureStatus;
            Name = name;
            Plan = plan;
            RenewalPlan = renewalPlan;
            SlotCount = slotCount;
            State = state;
        }
    }
}
