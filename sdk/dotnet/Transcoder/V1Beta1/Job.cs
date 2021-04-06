// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.Transcoder.V1Beta1
{
    /// <summary>
    /// Creates a job in the specified region.
    /// </summary>
    [GoogleCloudResourceType("google-cloud:transcoder/v1beta1:Job")]
    public partial class Job : Pulumi.CustomResource
    {
        /// <summary>
        /// The configuration for this job.
        /// </summary>
        [Output("config")]
        public Output<Outputs.JobConfigResponse> Config { get; private set; } = null!;

        /// <summary>
        /// The time the job was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The time the transcoding finished.
        /// </summary>
        [Output("endTime")]
        public Output<string> EndTime { get; private set; } = null!;

        /// <summary>
        /// List of failure details. This property may contain additional information about the failure when `failure_reason` is present. *Note*: This feature is not yet available.
        /// </summary>
        [Output("failureDetails")]
        public Output<ImmutableArray<Outputs.FailureDetailResponse>> FailureDetails { get; private set; } = null!;

        /// <summary>
        /// A description of the reason for the failure. This property is always present when `state` is `FAILED`.
        /// </summary>
        [Output("failureReason")]
        public Output<string> FailureReason { get; private set; } = null!;

        /// <summary>
        /// Input only. Specify the `input_uri` to populate empty `uri` fields in each element of `Job.config.inputs` or `JobTemplate.config.inputs` when using template. URI of the media. Input files must be at least 5 seconds in duration and stored in Cloud Storage (for example, `gs://bucket/inputs/file.mp4`).
        /// </summary>
        [Output("inputUri")]
        public Output<string> InputUri { get; private set; } = null!;

        /// <summary>
        /// The resource name of the job. Format: `projects/{project}/locations/{location}/jobs/{job}`
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The origin URI. *Note*: This feature is not yet available.
        /// </summary>
        [Output("originUri")]
        public Output<Outputs.OriginUriResponse> OriginUri { get; private set; } = null!;

        /// <summary>
        /// Input only. Specify the `output_uri` to populate an empty `Job.config.output.uri` or `JobTemplate.config.output.uri` when using template. URI for the output file(s). For example, `gs://my-bucket/outputs/`.
        /// </summary>
        [Output("outputUri")]
        public Output<string> OutputUri { get; private set; } = null!;

        /// <summary>
        /// Specify the priority of the job. Enter a value between 0 and 100, where 0 is the lowest priority and 100 is the highest priority. The default is 0.
        /// </summary>
        [Output("priority")]
        public Output<int> Priority { get; private set; } = null!;

        /// <summary>
        /// Estimated fractional progress, from `0` to `1` for each step. *Note*: This feature is not yet available.
        /// </summary>
        [Output("progress")]
        public Output<Outputs.ProgressResponse> Progress { get; private set; } = null!;

        /// <summary>
        /// The time the transcoding started.
        /// </summary>
        [Output("startTime")]
        public Output<string> StartTime { get; private set; } = null!;

        /// <summary>
        /// The current state of the job.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Input only. Specify the `template_id` to use for populating `Job.config`. The default is `preset/web-hd`. Preset Transcoder templates: - `preset/{preset_id}` - User defined JobTemplate: `{job_template_id}`
        /// </summary>
        [Output("templateId")]
        public Output<string> TemplateId { get; private set; } = null!;

        /// <summary>
        /// Job time to live value in days, which will be effective after job completion. Job should be deleted automatically after the given TTL. Enter a value between 1 and 90. The default is 30.
        /// </summary>
        [Output("ttlAfterCompletionDays")]
        public Output<int> TtlAfterCompletionDays { get; private set; } = null!;


        /// <summary>
        /// Create a Job resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Job(string name, JobArgs args, CustomResourceOptions? options = null)
            : base("google-cloud:transcoder/v1beta1:Job", name, args ?? new JobArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Job(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-cloud:transcoder/v1beta1:Job", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Job resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Job Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Job(name, id, options);
        }
    }

    public sealed class JobArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The configuration for this job.
        /// </summary>
        [Input("config")]
        public Input<Inputs.JobConfigArgs>? Config { get; set; }

        /// <summary>
        /// Input only. Specify the `input_uri` to populate empty `uri` fields in each element of `Job.config.inputs` or `JobTemplate.config.inputs` when using template. URI of the media. Input files must be at least 5 seconds in duration and stored in Cloud Storage (for example, `gs://bucket/inputs/file.mp4`).
        /// </summary>
        [Input("inputUri")]
        public Input<string>? InputUri { get; set; }

        [Input("jobsId", required: true)]
        public Input<string> JobsId { get; set; } = null!;

        [Input("locationsId", required: true)]
        public Input<string> LocationsId { get; set; } = null!;

        /// <summary>
        /// The resource name of the job. Format: `projects/{project}/locations/{location}/jobs/{job}`
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Input only. Specify the `output_uri` to populate an empty `Job.config.output.uri` or `JobTemplate.config.output.uri` when using template. URI for the output file(s). For example, `gs://my-bucket/outputs/`.
        /// </summary>
        [Input("outputUri")]
        public Input<string>? OutputUri { get; set; }

        /// <summary>
        /// Specify the priority of the job. Enter a value between 0 and 100, where 0 is the lowest priority and 100 is the highest priority. The default is 0.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        [Input("projectsId", required: true)]
        public Input<string> ProjectsId { get; set; } = null!;

        /// <summary>
        /// Input only. Specify the `template_id` to use for populating `Job.config`. The default is `preset/web-hd`. Preset Transcoder templates: - `preset/{preset_id}` - User defined JobTemplate: `{job_template_id}`
        /// </summary>
        [Input("templateId")]
        public Input<string>? TemplateId { get; set; }

        /// <summary>
        /// Job time to live value in days, which will be effective after job completion. Job should be deleted automatically after the given TTL. Enter a value between 1 and 90. The default is 30.
        /// </summary>
        [Input("ttlAfterCompletionDays")]
        public Input<int>? TtlAfterCompletionDays { get; set; }

        public JobArgs()
        {
        }
    }
}