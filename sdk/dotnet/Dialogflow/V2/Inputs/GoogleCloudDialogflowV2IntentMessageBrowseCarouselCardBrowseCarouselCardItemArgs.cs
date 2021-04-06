// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.Dialogflow.V2.Inputs
{

    /// <summary>
    /// Browsing carousel tile
    /// </summary>
    public sealed class GoogleCloudDialogflowV2IntentMessageBrowseCarouselCardBrowseCarouselCardItemArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. Description of the carousel item. Maximum of four lines of text.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Optional. Text that appears at the bottom of the Browse Carousel Card. Maximum of one line of text.
        /// </summary>
        [Input("footer")]
        public Input<string>? Footer { get; set; }

        /// <summary>
        /// Optional. Hero image for the carousel item.
        /// </summary>
        [Input("image")]
        public Input<Inputs.GoogleCloudDialogflowV2IntentMessageImageArgs>? Image { get; set; }

        /// <summary>
        /// Required. Action to present to the user.
        /// </summary>
        [Input("openUriAction")]
        public Input<Inputs.GoogleCloudDialogflowV2IntentMessageBrowseCarouselCardBrowseCarouselCardItemOpenUrlActionArgs>? OpenUriAction { get; set; }

        /// <summary>
        /// Required. Title of the carousel item. Maximum of two lines of text.
        /// </summary>
        [Input("title")]
        public Input<string>? Title { get; set; }

        public GoogleCloudDialogflowV2IntentMessageBrowseCarouselCardBrowseCarouselCardItemArgs()
        {
        }
    }
}