// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1Alpha1.Outputs
{

    /// <summary>
    /// Version contains structured information about the version of the package. For a discussion of this in Debian/Ubuntu: http://serverfault.com/questions/604541/debian-packages-version-convention For a discussion of this in Redhat/Fedora/Centos: http://blog.jasonantman.com/2014/07/how-yum-and-rpm-compare-versions/
    /// </summary>
    [OutputType]
    public sealed class VersionResponse
    {
        /// <summary>
        /// Used to correct mistakes in the version numbering scheme.
        /// </summary>
        public readonly int Epoch;
        /// <summary>
        /// Whether this version is vulnerable, when defining the version bounds. For example, if the minimum version is 2.0, inclusive=true would say 2.0 is vulnerable, while inclusive=false would say it's not
        /// </summary>
        public readonly bool Inclusive;
        /// <summary>
        /// Distinguish between sentinel MIN/MAX versions and normal versions. If kind is not NORMAL, then the other fields are ignored.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// The main part of the version name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The iteration of the package build from the above version.
        /// </summary>
        public readonly string Revision;

        [OutputConstructor]
        private VersionResponse(
            int epoch,

            bool inclusive,

            string kind,

            string name,

            string revision)
        {
            Epoch = epoch;
            Inclusive = inclusive;
            Kind = kind;
            Name = name;
            Revision = revision;
        }
    }
}