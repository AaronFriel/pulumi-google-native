// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1Alpha1.Inputs
{

    /// <summary>
    /// A CloudRepoSourceContext denotes a particular revision in a Google Cloud Source Repo.
    /// </summary>
    public sealed class GoogleDevtoolsContaineranalysisV1alpha1CloudRepoSourceContextArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An alias, which may be a branch or tag.
        /// </summary>
        [Input("aliasContext")]
        public Input<Inputs.GoogleDevtoolsContaineranalysisV1alpha1AliasContextArgs>? AliasContext { get; set; }

        /// <summary>
        /// The ID of the repo.
        /// </summary>
        [Input("repoId")]
        public Input<Inputs.GoogleDevtoolsContaineranalysisV1alpha1RepoIdArgs>? RepoId { get; set; }

        /// <summary>
        /// A revision ID.
        /// </summary>
        [Input("revisionId")]
        public Input<string>? RevisionId { get; set; }

        public GoogleDevtoolsContaineranalysisV1alpha1CloudRepoSourceContextArgs()
        {
        }
    }
}
