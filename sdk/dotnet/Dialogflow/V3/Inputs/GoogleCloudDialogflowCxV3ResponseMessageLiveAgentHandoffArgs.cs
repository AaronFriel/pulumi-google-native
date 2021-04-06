// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.Dialogflow.V3.Inputs
{

    /// <summary>
    /// Indicates that the conversation should be handed off to a live agent. Dialogflow only uses this to determine which conversations were handed off to a human agent for measurement purposes. What else to do with this signal is up to you and your handoff procedures. You may set this, for example: * In the entry_fulfillment of a Page if entering the page indicates something went extremely wrong in the conversation. * In a webhook response when you determine that the customer issue can only be handled by a human.
    /// </summary>
    public sealed class GoogleCloudDialogflowCxV3ResponseMessageLiveAgentHandoffArgs : Pulumi.ResourceArgs
    {
        [Input("metadata")]
        private InputMap<string>? _metadata;

        /// <summary>
        /// Custom metadata for your handoff procedure. Dialogflow doesn't impose any structure on this.
        /// </summary>
        public InputMap<string> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<string>());
            set => _metadata = value;
        }

        public GoogleCloudDialogflowCxV3ResponseMessageLiveAgentHandoffArgs()
        {
        }
    }
}