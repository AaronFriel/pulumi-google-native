// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V2.Inputs
{

    /// <summary>
    /// Knowledge base source settings. Supported features: ARTICLE_SUGGESTION, FAQ.
    /// </summary>
    public sealed class GoogleCloudDialogflowV2HumanAgentAssistantConfigSuggestionQueryConfigKnowledgeBaseQuerySourceArgs : Pulumi.ResourceArgs
    {
        [Input("knowledgeBases", required: true)]
        private InputList<string>? _knowledgeBases;

        /// <summary>
        /// Knowledge bases to query. Format: `projects//locations//knowledgeBases/`. Currently, at most 5 knowledge bases are supported.
        /// </summary>
        public InputList<string> KnowledgeBases
        {
            get => _knowledgeBases ?? (_knowledgeBases = new InputList<string>());
            set => _knowledgeBases = value;
        }

        public GoogleCloudDialogflowV2HumanAgentAssistantConfigSuggestionQueryConfigKnowledgeBaseQuerySourceArgs()
        {
        }
    }
}