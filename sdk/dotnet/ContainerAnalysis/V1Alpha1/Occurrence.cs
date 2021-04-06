// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.ContainerAnalysis.V1Alpha1
{
    /// <summary>
    /// Creates a new `Occurrence`. Use this method to create `Occurrences` for a resource.
    /// </summary>
    [GoogleCloudResourceType("google-cloud:containeranalysis/v1alpha1:Occurrence")]
    public partial class Occurrence : Pulumi.CustomResource
    {
        /// <summary>
        /// Describes an attestation of an artifact.
        /// </summary>
        [Output("attestation")]
        public Output<Outputs.AttestationResponse> Attestation { get; private set; } = null!;

        /// <summary>
        /// Build details for a verifiable build.
        /// </summary>
        [Output("buildDetails")]
        public Output<Outputs.BuildDetailsResponse> BuildDetails { get; private set; } = null!;

        /// <summary>
        /// The time this `Occurrence` was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Describes the deployment of an artifact on a runtime.
        /// </summary>
        [Output("deployment")]
        public Output<Outputs.DeploymentResponse> Deployment { get; private set; } = null!;

        /// <summary>
        /// Describes how this resource derives from the basis in the associated note.
        /// </summary>
        [Output("derivedImage")]
        public Output<Outputs.DerivedResponse> DerivedImage { get; private set; } = null!;

        /// <summary>
        /// Describes the initial scan status for this resource.
        /// </summary>
        [Output("discovered")]
        public Output<Outputs.DiscoveredResponse> Discovered { get; private set; } = null!;

        /// <summary>
        /// Describes the installation of a package on the linked resource.
        /// </summary>
        [Output("installation")]
        public Output<Outputs.InstallationResponse> Installation { get; private set; } = null!;

        /// <summary>
        /// This explicitly denotes which of the `Occurrence` details are specified. This field can be used as a filter in list requests.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// The name of the `Occurrence` in the form "projects/{project_id}/occurrences/{OCCURRENCE_ID}"
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// An analysis note associated with this image, in the form "providers/{provider_id}/notes/{NOTE_ID}" This field can be used as a filter in list requests.
        /// </summary>
        [Output("noteName")]
        public Output<string> NoteName { get; private set; } = null!;

        /// <summary>
        /// A description of actions that can be taken to remedy the `Note`
        /// </summary>
        [Output("remediation")]
        public Output<string> Remediation { get; private set; } = null!;

        /// <summary>
        ///  The resource for which the `Occurrence` applies.
        /// </summary>
        [Output("resource")]
        public Output<Outputs.ResourceResponse> Resource { get; private set; } = null!;

        /// <summary>
        /// The unique URL of the image or the container for which the `Occurrence` applies. For example, https://gcr.io/project/image@sha256:foo This field can be used as a filter in list requests.
        /// </summary>
        [Output("resourceUrl")]
        public Output<string> ResourceUrl { get; private set; } = null!;

        /// <summary>
        /// The time this `Occurrence` was last updated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;

        /// <summary>
        /// Describes an upgrade.
        /// </summary>
        [Output("upgrade")]
        public Output<Outputs.UpgradeOccurrenceResponse> Upgrade { get; private set; } = null!;

        /// <summary>
        /// Details of a security vulnerability note.
        /// </summary>
        [Output("vulnerabilityDetails")]
        public Output<Outputs.VulnerabilityDetailsResponse> VulnerabilityDetails { get; private set; } = null!;


        /// <summary>
        /// Create a Occurrence resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Occurrence(string name, OccurrenceArgs args, CustomResourceOptions? options = null)
            : base("google-cloud:containeranalysis/v1alpha1:Occurrence", name, args ?? new OccurrenceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Occurrence(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-cloud:containeranalysis/v1alpha1:Occurrence", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Occurrence resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Occurrence Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Occurrence(name, id, options);
        }
    }

    public sealed class OccurrenceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Describes an attestation of an artifact.
        /// </summary>
        [Input("attestation")]
        public Input<Inputs.AttestationArgs>? Attestation { get; set; }

        /// <summary>
        /// Build details for a verifiable build.
        /// </summary>
        [Input("buildDetails")]
        public Input<Inputs.BuildDetailsArgs>? BuildDetails { get; set; }

        /// <summary>
        /// The time this `Occurrence` was created.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Describes the deployment of an artifact on a runtime.
        /// </summary>
        [Input("deployment")]
        public Input<Inputs.DeploymentArgs>? Deployment { get; set; }

        /// <summary>
        /// Describes how this resource derives from the basis in the associated note.
        /// </summary>
        [Input("derivedImage")]
        public Input<Inputs.DerivedArgs>? DerivedImage { get; set; }

        /// <summary>
        /// Describes the initial scan status for this resource.
        /// </summary>
        [Input("discovered")]
        public Input<Inputs.DiscoveredArgs>? Discovered { get; set; }

        /// <summary>
        /// Describes the installation of a package on the linked resource.
        /// </summary>
        [Input("installation")]
        public Input<Inputs.InstallationArgs>? Installation { get; set; }

        /// <summary>
        /// This explicitly denotes which of the `Occurrence` details are specified. This field can be used as a filter in list requests.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// The name of the `Occurrence` in the form "projects/{project_id}/occurrences/{OCCURRENCE_ID}"
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// An analysis note associated with this image, in the form "providers/{provider_id}/notes/{NOTE_ID}" This field can be used as a filter in list requests.
        /// </summary>
        [Input("noteName")]
        public Input<string>? NoteName { get; set; }

        [Input("occurrencesId", required: true)]
        public Input<string> OccurrencesId { get; set; } = null!;

        [Input("projectsId", required: true)]
        public Input<string> ProjectsId { get; set; } = null!;

        /// <summary>
        /// A description of actions that can be taken to remedy the `Note`
        /// </summary>
        [Input("remediation")]
        public Input<string>? Remediation { get; set; }

        /// <summary>
        ///  The resource for which the `Occurrence` applies.
        /// </summary>
        [Input("resource")]
        public Input<Inputs.ResourceArgs>? Resource { get; set; }

        /// <summary>
        /// The unique URL of the image or the container for which the `Occurrence` applies. For example, https://gcr.io/project/image@sha256:foo This field can be used as a filter in list requests.
        /// </summary>
        [Input("resourceUrl")]
        public Input<string>? ResourceUrl { get; set; }

        /// <summary>
        /// The time this `Occurrence` was last updated.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        /// <summary>
        /// Describes an upgrade.
        /// </summary>
        [Input("upgrade")]
        public Input<Inputs.UpgradeOccurrenceArgs>? Upgrade { get; set; }

        /// <summary>
        /// Details of a security vulnerability note.
        /// </summary>
        [Input("vulnerabilityDetails")]
        public Input<Inputs.VulnerabilityDetailsArgs>? VulnerabilityDetails { get; set; }

        public OccurrenceArgs()
        {
        }
    }
}