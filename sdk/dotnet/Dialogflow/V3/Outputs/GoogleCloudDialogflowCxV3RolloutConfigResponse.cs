// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V3.Outputs
{

    [OutputType]
    public sealed class GoogleCloudDialogflowCxV3RolloutConfigResponse
    {
        /// <summary>
        /// The conditions that are used to evaluate the failure of a rollout step. If not specified, no rollout steps will fail. E.g. "containment_rate &lt; 10% OR average_turn_count &lt; 3". See the [conditions reference](https://cloud.google.com/dialogflow/cx/docs/reference/condition).
        /// </summary>
        public readonly string FailureCondition;
        /// <summary>
        /// The conditions that are used to evaluate the success of a rollout step. If not specified, all rollout steps will proceed to the next one unless failure conditions are met. E.g. "containment_rate &gt; 60% AND callback_rate &lt; 20%". See the [conditions reference](https://cloud.google.com/dialogflow/cx/docs/reference/condition).
        /// </summary>
        public readonly string RolloutCondition;
        /// <summary>
        /// Steps to roll out a flow version. Steps should be sorted by percentage in ascending order.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudDialogflowCxV3RolloutConfigRolloutStepResponse> RolloutSteps;

        [OutputConstructor]
        private GoogleCloudDialogflowCxV3RolloutConfigResponse(
            string failureCondition,

            string rolloutCondition,

            ImmutableArray<Outputs.GoogleCloudDialogflowCxV3RolloutConfigRolloutStepResponse> rolloutSteps)
        {
            FailureCondition = failureCondition;
            RolloutCondition = rolloutCondition;
            RolloutSteps = rolloutSteps;
        }
    }
}