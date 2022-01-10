// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.APIKeys.V2.Outputs
{

    /// <summary>
    /// The iOS apps that are allowed to use the key.
    /// </summary>
    [OutputType]
    public sealed class V2IosKeyRestrictionsResponse
    {
        /// <summary>
        /// A list of bundle IDs that are allowed when making API calls with this key.
        /// </summary>
        public readonly ImmutableArray<string> AllowedBundleIds;

        [OutputConstructor]
        private V2IosKeyRestrictionsResponse(ImmutableArray<string> allowedBundleIds)
        {
            AllowedBundleIds = allowedBundleIds;
        }
    }
}