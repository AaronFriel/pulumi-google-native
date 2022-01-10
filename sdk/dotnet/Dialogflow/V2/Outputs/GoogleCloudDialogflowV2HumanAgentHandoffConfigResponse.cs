// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V2.Outputs
{

    /// <summary>
    /// Defines the hand off to a live agent, typically on which external agent service provider to connect to a conversation. Currently, this feature is not general available, please contact Google to get access.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDialogflowV2HumanAgentHandoffConfigResponse
    {
        /// <summary>
        /// Uses LivePerson (https://www.liveperson.com).
        /// </summary>
        public readonly Outputs.GoogleCloudDialogflowV2HumanAgentHandoffConfigLivePersonConfigResponse LivePersonConfig;
        /// <summary>
        /// Uses Salesforce Live Agent.
        /// </summary>
        public readonly Outputs.GoogleCloudDialogflowV2HumanAgentHandoffConfigSalesforceLiveAgentConfigResponse SalesforceLiveAgentConfig;

        [OutputConstructor]
        private GoogleCloudDialogflowV2HumanAgentHandoffConfigResponse(
            Outputs.GoogleCloudDialogflowV2HumanAgentHandoffConfigLivePersonConfigResponse livePersonConfig,

            Outputs.GoogleCloudDialogflowV2HumanAgentHandoffConfigSalesforceLiveAgentConfigResponse salesforceLiveAgentConfig)
        {
            LivePersonConfig = livePersonConfig;
            SalesforceLiveAgentConfig = salesforceLiveAgentConfig;
        }
    }
}