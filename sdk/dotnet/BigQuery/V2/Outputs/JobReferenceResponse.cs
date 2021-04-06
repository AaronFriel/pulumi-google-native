// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.BigQuery.V2.Outputs
{

    [OutputType]
    public sealed class JobReferenceResponse
    {
        /// <summary>
        /// [Required] The ID of the job. The ID must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), or dashes (-). The maximum length is 1,024 characters.
        /// </summary>
        public readonly string JobId;
        /// <summary>
        /// The geographic location of the job. See details at https://cloud.google.com/bigquery/docs/locations#specifying_your_location.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// [Required] The ID of the project containing this job.
        /// </summary>
        public readonly string ProjectId;

        [OutputConstructor]
        private JobReferenceResponse(
            string jobId,

            string location,

            string projectId)
        {
            JobId = jobId;
            Location = location;
            ProjectId = projectId;
        }
    }
}