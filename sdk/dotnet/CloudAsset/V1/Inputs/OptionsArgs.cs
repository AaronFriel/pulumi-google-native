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
    /// Contains query options.
    /// </summary>
    public sealed class OptionsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. If true, the response will include access analysis from identities to resources via service account impersonation. This is a very expensive operation, because many derived queries will be executed. We highly recommend you use AssetService.AnalyzeIamPolicyLongrunning rpc instead. For example, if the request analyzes for which resources user A has permission P, and there's an IAM policy states user A has iam.serviceAccounts.getAccessToken permission to a service account SA, and there's another IAM policy states service account SA has permission P to a GCP folder F, then user A potentially has access to the GCP folder F. And those advanced analysis results will be included in AnalyzeIamPolicyResponse.service_account_impersonation_analysis. Another example, if the request analyzes for who has permission P to a GCP folder F, and there's an IAM policy states user A has iam.serviceAccounts.actAs permission to a service account SA, and there's another IAM policy states service account SA has permission P to the GCP folder F, then user A potentially has access to the GCP folder F. And those advanced analysis results will be included in AnalyzeIamPolicyResponse.service_account_impersonation_analysis. Default is false.
        /// </summary>
        [Input("analyzeServiceAccountImpersonation")]
        public Input<bool>? AnalyzeServiceAccountImpersonation { get; set; }

        /// <summary>
        /// Optional. If true, the identities section of the result will expand any Google groups appearing in an IAM policy binding. If IamPolicyAnalysisQuery.identity_selector is specified, the identity in the result will be determined by the selector, and this flag is not allowed to set. Default is false.
        /// </summary>
        [Input("expandGroups")]
        public Input<bool>? ExpandGroups { get; set; }

        /// <summary>
        /// Optional. If true and IamPolicyAnalysisQuery.resource_selector is not specified, the resource section of the result will expand any resource attached to an IAM policy to include resources lower in the resource hierarchy. For example, if the request analyzes for which resources user A has permission P, and the results include an IAM policy with P on a GCP folder, the results will also include resources in that folder with permission P. If true and IamPolicyAnalysisQuery.resource_selector is specified, the resource section of the result will expand the specified resource to include resources lower in the resource hierarchy. Only project or lower resources are supported. Folder and organization resource cannot be used together with this option. For example, if the request analyzes for which users have permission P on a GCP project with this option enabled, the results will include all users who have permission P on that project or any lower resource. Default is false.
        /// </summary>
        [Input("expandResources")]
        public Input<bool>? ExpandResources { get; set; }

        /// <summary>
        /// Optional. If true, the access section of result will expand any roles appearing in IAM policy bindings to include their permissions. If IamPolicyAnalysisQuery.access_selector is specified, the access section of the result will be determined by the selector, and this flag is not allowed to set. Default is false.
        /// </summary>
        [Input("expandRoles")]
        public Input<bool>? ExpandRoles { get; set; }

        /// <summary>
        /// Optional. If true, the result will output the relevant membership relationships between groups and other groups, and between groups and principals. Default is false.
        /// </summary>
        [Input("outputGroupEdges")]
        public Input<bool>? OutputGroupEdges { get; set; }

        /// <summary>
        /// Optional. If true, the result will output the relevant parent/child relationships between resources. Default is false.
        /// </summary>
        [Input("outputResourceEdges")]
        public Input<bool>? OutputResourceEdges { get; set; }

        public OptionsArgs()
        {
        }
    }
}
