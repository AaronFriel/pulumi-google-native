// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.DLP.V2.Inputs
{

    /// <summary>
    /// Set of primitive values supported by the system. Note that for the purposes of inspection or transformation, the number of bytes considered to comprise a 'Value' is based on its representation as a UTF-8 encoded string. For example, if 'integer_value' is set to 123456789, the number of bytes would be counted as 9, even though an int64 only holds up to 8 bytes of data.
    /// </summary>
    public sealed class GooglePrivacyDlpV2ValueArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// boolean
        /// </summary>
        [Input("booleanValue")]
        public Input<bool>? BooleanValue { get; set; }

        /// <summary>
        /// date
        /// </summary>
        [Input("dateValue")]
        public Input<Inputs.GoogleTypeDateArgs>? DateValue { get; set; }

        /// <summary>
        /// day of week
        /// </summary>
        [Input("dayOfWeekValue")]
        public Input<string>? DayOfWeekValue { get; set; }

        /// <summary>
        /// float
        /// </summary>
        [Input("floatValue")]
        public Input<double>? FloatValue { get; set; }

        /// <summary>
        /// integer
        /// </summary>
        [Input("integerValue")]
        public Input<string>? IntegerValue { get; set; }

        /// <summary>
        /// string
        /// </summary>
        [Input("stringValue")]
        public Input<string>? StringValue { get; set; }

        /// <summary>
        /// time of day
        /// </summary>
        [Input("timeValue")]
        public Input<Inputs.GoogleTypeTimeOfDayArgs>? TimeValue { get; set; }

        /// <summary>
        /// timestamp
        /// </summary>
        [Input("timestampValue")]
        public Input<string>? TimestampValue { get; set; }

        public GooglePrivacyDlpV2ValueArgs()
        {
        }
    }
}