// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudAsset.V1.Inputs
{

    /// <summary>
    /// IAM policy analysis query message.
    /// </summary>
    public sealed class IamPolicyAnalysisQueryArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. Specifies roles or permissions for analysis. This is optional.
        /// </summary>
        [Input("accessSelector")]
        public Input<Inputs.AccessSelectorArgs>? AccessSelector { get; set; }

        /// <summary>
        /// Optional. The hypothetical context for IAM conditions evaluation.
        /// </summary>
        [Input("conditionContext")]
        public Input<Inputs.ConditionContextArgs>? ConditionContext { get; set; }

        /// <summary>
        /// Optional. Specifies an identity for analysis.
        /// </summary>
        [Input("identitySelector")]
        public Input<Inputs.IdentitySelectorArgs>? IdentitySelector { get; set; }

        /// <summary>
        /// Optional. The query options.
        /// </summary>
        [Input("options")]
        public Input<Inputs.OptionsArgs>? Options { get; set; }

        /// <summary>
        /// Optional. Specifies a resource for analysis.
        /// </summary>
        [Input("resourceSelector")]
        public Input<Inputs.ResourceSelectorArgs>? ResourceSelector { get; set; }

        /// <summary>
        /// The relative name of the root asset. Only resources and IAM policies within the scope will be analyzed. This can only be an organization number (such as "organizations/123"), a folder number (such as "folders/123"), a project ID (such as "projects/my-project-id"), or a project number (such as "projects/12345"). To know how to get organization id, visit [here ](https://cloud.google.com/resource-manager/docs/creating-managing-organization#retrieving_your_organization_id). To know how to get folder or project id, visit [here ](https://cloud.google.com/resource-manager/docs/creating-managing-folders#viewing_or_listing_folders_and_projects).
        /// </summary>
        [Input("scope", required: true)]
        public Input<string> Scope { get; set; } = null!;

        public IamPolicyAnalysisQueryArgs()
        {
        }
    }
}