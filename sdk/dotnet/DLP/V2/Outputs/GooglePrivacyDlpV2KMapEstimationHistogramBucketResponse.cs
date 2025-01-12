// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DLP.V2.Outputs
{

    /// <summary>
    /// A KMapEstimationHistogramBucket message with the following values: min_anonymity: 3 max_anonymity: 5 frequency: 42 means that there are 42 records whose quasi-identifier values correspond to 3, 4 or 5 people in the overlying population. An important particular case is when min_anonymity = max_anonymity = 1: the frequency field then corresponds to the number of uniquely identifiable records.
    /// </summary>
    [OutputType]
    public sealed class GooglePrivacyDlpV2KMapEstimationHistogramBucketResponse
    {
        /// <summary>
        /// Number of records within these anonymity bounds.
        /// </summary>
        public readonly string BucketSize;
        /// <summary>
        /// Total number of distinct quasi-identifier tuple values in this bucket.
        /// </summary>
        public readonly string BucketValueCount;
        /// <summary>
        /// Sample of quasi-identifier tuple values in this bucket. The total number of classes returned per bucket is capped at 20.
        /// </summary>
        public readonly ImmutableArray<Outputs.GooglePrivacyDlpV2KMapEstimationQuasiIdValuesResponse> BucketValues;
        /// <summary>
        /// Always greater than or equal to min_anonymity.
        /// </summary>
        public readonly string MaxAnonymity;
        /// <summary>
        /// Always positive.
        /// </summary>
        public readonly string MinAnonymity;

        [OutputConstructor]
        private GooglePrivacyDlpV2KMapEstimationHistogramBucketResponse(
            string bucketSize,

            string bucketValueCount,

            ImmutableArray<Outputs.GooglePrivacyDlpV2KMapEstimationQuasiIdValuesResponse> bucketValues,

            string maxAnonymity,

            string minAnonymity)
        {
            BucketSize = bucketSize;
            BucketValueCount = bucketValueCount;
            BucketValues = bucketValues;
            MaxAnonymity = maxAnonymity;
            MinAnonymity = minAnonymity;
        }
    }
}
