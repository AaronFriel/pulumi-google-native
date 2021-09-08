// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Contactcenterinsights.V1.Outputs
{

    [OutputType]
    public sealed class GoogleCloudContactcenterinsightsV1PhraseMatchRuleConfigResponse
    {
        /// <summary>
        /// The configuration for the exact match rule.
        /// </summary>
        public readonly Outputs.GoogleCloudContactcenterinsightsV1ExactMatchConfigResponse ExactMatchConfig;

        [OutputConstructor]
        private GoogleCloudContactcenterinsightsV1PhraseMatchRuleConfigResponse(Outputs.GoogleCloudContactcenterinsightsV1ExactMatchConfigResponse exactMatchConfig)
        {
            ExactMatchConfig = exactMatchConfig;
        }
    }
}