// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.GoogleNative.Run.V1
{
    public static class GetService
    {
        /// <summary>
        /// Get information about a service.
        /// </summary>
        public static Task<GetServiceResult> InvokeAsync(GetServiceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetServiceResult>("google-native:run/v1:getService", args ?? new GetServiceArgs(), options.WithVersion());

        /// <summary>
        /// Get information about a service.
        /// </summary>
        public static Output<GetServiceResult> Invoke(GetServiceInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetServiceResult>("google-native:run/v1:getService", args ?? new GetServiceInvokeArgs(), options.WithVersion());
    }


    public sealed class GetServiceArgs : Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("serviceId", required: true)]
        public string ServiceId { get; set; } = null!;

        public GetServiceArgs()
        {
        }
    }

    public sealed class GetServiceInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("serviceId", required: true)]
        public Input<string> ServiceId { get; set; } = null!;

        public GetServiceInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetServiceResult
    {
        /// <summary>
        /// The API version for this call such as "serving.knative.dev/v1".
        /// </summary>
        public readonly string ApiVersion;
        /// <summary>
        /// The kind of resource, in this case "Service".
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Metadata associated with this Service, including name, namespace, labels, and annotations. Cloud Run (fully managed) uses the following annotation keys to configure features on a Service: * `run.googleapis.com/ingress` sets the ingress settings for the Service. See [the ingress settings documentation](/run/docs/securing/ingress) for details on configuring ingress settings. * `run.googleapis.com/ingress-status` is output-only and contains the currently active ingress settings for the Service. `run.googleapis.com/ingress-status` may differ from `run.googleapis.com/ingress` while the system is processing a change to `run.googleapis.com/ingress` or if the system failed to process a change to `run.googleapis.com/ingress`. When the system has processed all changes successfully `run.googleapis.com/ingress-status` and `run.googleapis.com/ingress` are equal.
        /// </summary>
        public readonly Outputs.ObjectMetaResponse Metadata;
        /// <summary>
        /// Spec holds the desired state of the Service (from the client).
        /// </summary>
        public readonly Outputs.ServiceSpecResponse Spec;
        /// <summary>
        /// Status communicates the observed state of the Service (from the controller).
        /// </summary>
        public readonly Outputs.ServiceStatusResponse Status;

        [OutputConstructor]
        private GetServiceResult(
            string apiVersion,

            string kind,

            Outputs.ObjectMetaResponse metadata,

            Outputs.ServiceSpecResponse spec,

            Outputs.ServiceStatusResponse status)
        {
            ApiVersion = apiVersion;
            Kind = kind;
            Metadata = metadata;
            Spec = spec;
            Status = status;
        }
    }
}