// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.Vision.V1.Outputs
{

    [OutputType]
    public sealed class NormalizedVertexResponse
    {
        /// <summary>
        /// X coordinate.
        /// </summary>
        public readonly double X;
        /// <summary>
        /// Y coordinate.
        /// </summary>
        public readonly double Y;

        [OutputConstructor]
        private NormalizedVertexResponse(
            double x,

            double y)
        {
            X = x;
            Y = y;
        }
    }
}