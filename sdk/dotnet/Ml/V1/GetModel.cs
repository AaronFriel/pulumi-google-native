// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.GoogleNative.Ml.V1
{
    public static class GetModel
    {
        /// <summary>
        /// Gets information about a model, including its name, the description (if set), and the default version (if at least one version of the model has been deployed).
        /// </summary>
        public static Task<GetModelResult> InvokeAsync(GetModelArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetModelResult>("google-native:ml/v1:getModel", args ?? new GetModelArgs(), options.WithVersion());

        /// <summary>
        /// Gets information about a model, including its name, the description (if set), and the default version (if at least one version of the model has been deployed).
        /// </summary>
        public static Output<GetModelResult> Invoke(GetModelInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetModelResult>("google-native:ml/v1:getModel", args ?? new GetModelInvokeArgs(), options.WithVersion());
    }


    public sealed class GetModelArgs : Pulumi.InvokeArgs
    {
        [Input("modelId", required: true)]
        public string ModelId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetModelArgs()
        {
        }
    }

    public sealed class GetModelInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("modelId", required: true)]
        public Input<string> ModelId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetModelInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetModelResult
    {
        /// <summary>
        /// The default version of the model. This version will be used to handle prediction requests that do not specify a version. You can change the default version by calling projects.models.versions.setDefault.
        /// </summary>
        public readonly Outputs.GoogleCloudMlV1__VersionResponse DefaultVersion;
        /// <summary>
        /// Optional. The description specified for the model when it was created.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a model from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform model updates in order to avoid race conditions: An `etag` is returned in the response to `GetModel`, and systems are expected to put that etag in the request to `UpdateModel` to ensure that their change will be applied to the model as intended.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Optional. One or more labels that you can add, to organize your models. Each label is a key-value pair, where both the key and the value are arbitrary strings that you supply. For more information, see the documentation on using labels.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// The name specified for the model when it was created. The model name must be unique within the project it is created in.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Optional. If true, online prediction nodes send `stderr` and `stdout` streams to Cloud Logging. These can be more verbose than the standard access logs (see `onlinePredictionLogging`) and can incur higher cost. However, they are helpful for debugging. Note that [logs may incur a cost](/stackdriver/pricing), especially if your project receives prediction requests at a high QPS. Estimate your costs before enabling this option. Default is false.
        /// </summary>
        public readonly bool OnlinePredictionConsoleLogging;
        /// <summary>
        /// Optional. If true, online prediction access logs are sent to Cloud Logging. These logs are like standard server access logs, containing information like timestamp and latency for each request. Note that [logs may incur a cost](/stackdriver/pricing), especially if your project receives prediction requests at a high queries per second rate (QPS). Estimate your costs before enabling this option. Default is false.
        /// </summary>
        public readonly bool OnlinePredictionLogging;
        /// <summary>
        /// Optional. The list of regions where the model is going to be deployed. Only one region per model is supported. Defaults to 'us-central1' if nothing is set. See the available regions for AI Platform services. Note: * No matter where a model is deployed, it can always be accessed by users from anywhere, both for online and batch prediction. * The region for a batch prediction job is set by the region field when submitting the batch prediction job and does not take its value from this field.
        /// </summary>
        public readonly ImmutableArray<string> Regions;

        [OutputConstructor]
        private GetModelResult(
            Outputs.GoogleCloudMlV1__VersionResponse defaultVersion,

            string description,

            string etag,

            ImmutableDictionary<string, string> labels,

            string name,

            bool onlinePredictionConsoleLogging,

            bool onlinePredictionLogging,

            ImmutableArray<string> regions)
        {
            DefaultVersion = defaultVersion;
            Description = description;
            Etag = etag;
            Labels = labels;
            Name = name;
            OnlinePredictionConsoleLogging = onlinePredictionConsoleLogging;
            OnlinePredictionLogging = onlinePredictionLogging;
            Regions = regions;
        }
    }
}