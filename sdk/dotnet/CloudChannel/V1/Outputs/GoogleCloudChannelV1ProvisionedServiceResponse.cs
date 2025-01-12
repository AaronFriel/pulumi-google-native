// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudChannel.V1.Outputs
{

    /// <summary>
    /// Service provisioned for an entitlement.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudChannelV1ProvisionedServiceResponse
    {
        /// <summary>
        /// The product pertaining to the provisioning resource as specified in the Offer.
        /// </summary>
        public readonly string ProductId;
        /// <summary>
        /// Provisioning ID of the entitlement. For Google Workspace, this would be the underlying Subscription ID.
        /// </summary>
        public readonly string ProvisioningId;
        /// <summary>
        /// The SKU pertaining to the provisioning resource as specified in the Offer.
        /// </summary>
        public readonly string SkuId;

        [OutputConstructor]
        private GoogleCloudChannelV1ProvisionedServiceResponse(
            string productId,

            string provisioningId,

            string skuId)
        {
            ProductId = productId;
            ProvisioningId = provisioningId;
            SkuId = skuId;
        }
    }
}
