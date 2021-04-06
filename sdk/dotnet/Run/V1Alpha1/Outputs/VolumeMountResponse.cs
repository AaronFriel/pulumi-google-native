// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.Run.V1Alpha1.Outputs
{

    [OutputType]
    public sealed class VolumeMountResponse
    {
        /// <summary>
        /// Path within the container at which the volume should be mounted. Must not contain ':'.
        /// </summary>
        public readonly string MountPath;
        /// <summary>
        /// mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationHostToContainer is used. This field is beta in 1.10. +optional
        /// </summary>
        public readonly string MountPropagation;
        /// <summary>
        /// This must match the Name of a Volume.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false. +optional
        /// </summary>
        public readonly bool ReadOnly;
        /// <summary>
        /// Path within the volume from which the container's volume should be mounted. Defaults to "" (volume's root). +optional
        /// </summary>
        public readonly string SubPath;

        [OutputConstructor]
        private VolumeMountResponse(
            string mountPath,

            string mountPropagation,

            string name,

            bool readOnly,

            string subPath)
        {
            MountPath = mountPath;
            MountPropagation = mountPropagation;
            Name = name;
            ReadOnly = readOnly;
            SubPath = subPath;
        }
    }
}