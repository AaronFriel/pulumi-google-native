// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Contactcenterinsights.V1.Inputs
{

    /// <summary>
    /// A message representing a rule in the phrase matcher.
    /// </summary>
    public sealed class GoogleCloudContactcenterinsightsV1PhraseMatchRuleGroupArgs : Pulumi.ResourceArgs
    {
        [Input("phraseMatchRules")]
        private InputList<Inputs.GoogleCloudContactcenterinsightsV1PhraseMatchRuleArgs>? _phraseMatchRules;

        /// <summary>
        /// A list of phase match rules that are included in this group.
        /// </summary>
        public InputList<Inputs.GoogleCloudContactcenterinsightsV1PhraseMatchRuleArgs> PhraseMatchRules
        {
            get => _phraseMatchRules ?? (_phraseMatchRules = new InputList<Inputs.GoogleCloudContactcenterinsightsV1PhraseMatchRuleArgs>());
            set => _phraseMatchRules = value;
        }

        /// <summary>
        /// The type of this phrase match rule group.
        /// </summary>
        [Input("type", required: true)]
        public Input<Pulumi.GoogleNative.Contactcenterinsights.V1.GoogleCloudContactcenterinsightsV1PhraseMatchRuleGroupType> Type { get; set; } = null!;

        public GoogleCloudContactcenterinsightsV1PhraseMatchRuleGroupArgs()
        {
        }
    }
}
