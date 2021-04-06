// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.Logging.V2.Outputs
{

    [OutputType]
    public sealed class LinearResponse
    {
        /// <summary>
        /// Must be greater than 0.
        /// </summary>
        public readonly int NumFiniteBuckets;
        /// <summary>
        /// Lower bound of the first bucket.
        /// </summary>
        public readonly double Offset;
        /// <summary>
        /// Must be greater than 0.
        /// </summary>
        public readonly double Width;

        [OutputConstructor]
        private LinearResponse(
            int numFiniteBuckets,

            double offset,

            double width)
        {
            NumFiniteBuckets = numFiniteBuckets;
            Offset = offset;
            Width = width;
        }
    }
}