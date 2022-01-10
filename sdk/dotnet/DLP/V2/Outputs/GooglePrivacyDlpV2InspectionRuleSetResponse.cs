// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DLP.V2.Outputs
{

    /// <summary>
    /// Rule set for modifying a set of infoTypes to alter behavior under certain circumstances, depending on the specific details of the rules within the set.
    /// </summary>
    [OutputType]
    public sealed class GooglePrivacyDlpV2InspectionRuleSetResponse
    {
        /// <summary>
        /// List of infoTypes this rule set is applied to.
        /// </summary>
        public readonly ImmutableArray<Outputs.GooglePrivacyDlpV2InfoTypeResponse> InfoTypes;
        /// <summary>
        /// Set of rules to be applied to infoTypes. The rules are applied in order.
        /// </summary>
        public readonly ImmutableArray<Outputs.GooglePrivacyDlpV2InspectionRuleResponse> Rules;

        [OutputConstructor]
        private GooglePrivacyDlpV2InspectionRuleSetResponse(
            ImmutableArray<Outputs.GooglePrivacyDlpV2InfoTypeResponse> infoTypes,

            ImmutableArray<Outputs.GooglePrivacyDlpV2InspectionRuleResponse> rules)
        {
            InfoTypes = infoTypes;
            Rules = rules;
        }
    }
}