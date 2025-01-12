// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataproc.V1.Outputs
{

    /// <summary>
    /// Specifies a Metastore configuration.
    /// </summary>
    [OutputType]
    public sealed class MetastoreConfigResponse
    {
        /// <summary>
        /// Resource name of an existing Dataproc Metastore service.Example: projects/[project_id]/locations/[dataproc_region]/services/[service-name]
        /// </summary>
        public readonly string DataprocMetastoreService;

        [OutputConstructor]
        private MetastoreConfigResponse(string dataprocMetastoreService)
        {
            DataprocMetastoreService = dataprocMetastoreService;
        }
    }
}
