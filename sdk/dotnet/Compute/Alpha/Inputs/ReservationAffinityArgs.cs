// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.Compute.Alpha.Inputs
{

    /// <summary>
    /// Specifies the reservations that this instance can consume from.
    /// </summary>
    public sealed class ReservationAffinityArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the type of reservation from which this instance can consume resources: ANY_RESERVATION (default), SPECIFIC_RESERVATION, or NO_RESERVATION. See  Consuming reserved instances for examples.
        /// </summary>
        [Input("consumeReservationType")]
        public Input<string>? ConsumeReservationType { get; set; }

        /// <summary>
        /// Corresponds to the label key of a reservation resource. To target a SPECIFIC_RESERVATION by name, specify googleapis.com/reservation-name as the key and specify the name of your reservation as its value.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        [Input("values")]
        private InputList<string>? _values;

        /// <summary>
        /// Corresponds to the label values of a reservation resource.
        /// </summary>
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        public ReservationAffinityArgs()
        {
        }
    }
}