// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.Dialogflow.V2Beta1.Outputs
{

    [OutputType]
    public sealed class GoogleCloudDialogflowV2beta1IntentMessageCarouselSelectResponse
    {
        /// <summary>
        /// Required. Carousel items.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudDialogflowV2beta1IntentMessageCarouselSelectItemResponse> Items;

        [OutputConstructor]
        private GoogleCloudDialogflowV2beta1IntentMessageCarouselSelectResponse(ImmutableArray<Outputs.GoogleCloudDialogflowV2beta1IntentMessageCarouselSelectItemResponse> items)
        {
            Items = items;
        }
    }
}