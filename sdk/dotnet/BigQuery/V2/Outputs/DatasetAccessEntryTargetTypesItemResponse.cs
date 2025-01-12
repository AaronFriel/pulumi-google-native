// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BigQuery.V2.Outputs
{

    [OutputType]
    public sealed class DatasetAccessEntryTargetTypesItemResponse
    {
        /// <summary>
        /// [Required] Which resources in the dataset this entry applies to. Currently, only views are supported, but additional target types may be added in the future. Possible values: VIEWS: This entry applies to all views in the dataset.
        /// </summary>
        public readonly string TargetType;

        [OutputConstructor]
        private DatasetAccessEntryTargetTypesItemResponse(string targetType)
        {
            TargetType = targetType;
        }
    }
}
