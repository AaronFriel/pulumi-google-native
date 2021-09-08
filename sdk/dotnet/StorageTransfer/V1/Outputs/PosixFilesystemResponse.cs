// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.StorageTransfer.V1.Outputs
{

    [OutputType]
    public sealed class PosixFilesystemResponse
    {
        /// <summary>
        /// Root directory path to the filesystem.
        /// </summary>
        public readonly string RootDirectory;

        [OutputConstructor]
        private PosixFilesystemResponse(string rootDirectory)
        {
            RootDirectory = rootDirectory;
        }
    }
}