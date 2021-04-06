// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.ContainerAnalysis.V1Alpha1.Outputs
{

    [OutputType]
    public sealed class DetailResponse
    {
        /// <summary>
        /// The cpe_uri in [cpe format] (https://cpe.mitre.org/specification/) in which the vulnerability manifests. Examples include distro or storage location for vulnerable jar. This field can be used as a filter in list requests.
        /// </summary>
        public readonly string CpeUri;
        /// <summary>
        /// A vendor-specific description of this note.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The fix for this specific package version.
        /// </summary>
        public readonly Outputs.VulnerabilityLocationResponse FixedLocation;
        /// <summary>
        /// Whether this Detail is obsolete. Occurrences are expected not to point to obsolete details.
        /// </summary>
        public readonly bool IsObsolete;
        /// <summary>
        /// The max version of the package in which the vulnerability exists.
        /// </summary>
        public readonly Outputs.VersionResponse MaxAffectedVersion;
        /// <summary>
        /// The min version of the package in which the vulnerability exists.
        /// </summary>
        public readonly Outputs.VersionResponse MinAffectedVersion;
        /// <summary>
        /// The name of the package where the vulnerability was found. This field can be used as a filter in list requests.
        /// </summary>
        public readonly string Package;
        /// <summary>
        /// The type of package; whether native or non native(ruby gems, node.js packages etc)
        /// </summary>
        public readonly string PackageType;
        /// <summary>
        /// The severity (eg: distro assigned severity) for this vulnerability.
        /// </summary>
        public readonly string SeverityName;
        /// <summary>
        /// The source from which the information in this Detail was obtained.
        /// </summary>
        public readonly string Source;

        [OutputConstructor]
        private DetailResponse(
            string cpeUri,

            string description,

            Outputs.VulnerabilityLocationResponse fixedLocation,

            bool isObsolete,

            Outputs.VersionResponse maxAffectedVersion,

            Outputs.VersionResponse minAffectedVersion,

            string package,

            string packageType,

            string severityName,

            string source)
        {
            CpeUri = cpeUri;
            Description = description;
            FixedLocation = fixedLocation;
            IsObsolete = isObsolete;
            MaxAffectedVersion = maxAffectedVersion;
            MinAffectedVersion = minAffectedVersion;
            Package = package;
            PackageType = packageType;
            SeverityName = severityName;
            Source = source;
        }
    }
}