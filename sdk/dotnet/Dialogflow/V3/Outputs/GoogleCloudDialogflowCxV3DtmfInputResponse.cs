// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.Dialogflow.V3.Outputs
{

    [OutputType]
    public sealed class GoogleCloudDialogflowCxV3DtmfInputResponse
    {
        /// <summary>
        /// The dtmf digits.
        /// </summary>
        public readonly string Digits;
        /// <summary>
        /// The finish digit (if any).
        /// </summary>
        public readonly string FinishDigit;

        [OutputConstructor]
        private GoogleCloudDialogflowCxV3DtmfInputResponse(
            string digits,

            string finishDigit)
        {
            Digits = digits;
            FinishDigit = finishDigit;
        }
    }
}