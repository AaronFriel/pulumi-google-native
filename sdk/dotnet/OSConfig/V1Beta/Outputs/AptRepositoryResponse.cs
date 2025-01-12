// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.OSConfig.V1Beta.Outputs
{

    /// <summary>
    /// Represents a single Apt package repository. This repository is added to a repo file that is stored at `/etc/apt/sources.list.d/google_osconfig.list`.
    /// </summary>
    [OutputType]
    public sealed class AptRepositoryResponse
    {
        /// <summary>
        /// Type of archive files in this repository. The default behavior is DEB.
        /// </summary>
        public readonly string ArchiveType;
        /// <summary>
        /// List of components for this repository. Must contain at least one item.
        /// </summary>
        public readonly ImmutableArray<string> Components;
        /// <summary>
        /// Distribution of this repository.
        /// </summary>
        public readonly string Distribution;
        /// <summary>
        /// URI of the key file for this repository. The agent maintains a keyring at `/etc/apt/trusted.gpg.d/osconfig_agent_managed.gpg` containing all the keys in any applied guest policy.
        /// </summary>
        public readonly string GpgKey;
        /// <summary>
        /// URI for this repository.
        /// </summary>
        public readonly string Uri;

        [OutputConstructor]
        private AptRepositoryResponse(
            string archiveType,

            ImmutableArray<string> components,

            string distribution,

            string gpgKey,

            string uri)
        {
            ArchiveType = archiveType;
            Components = components;
            Distribution = distribution;
            GpgKey = gpgKey;
            Uri = uri;
        }
    }
}
