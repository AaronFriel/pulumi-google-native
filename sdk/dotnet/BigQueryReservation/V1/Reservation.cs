// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.BigQueryReservation.V1
{
    /// <summary>
    /// Creates a new reservation resource.
    /// </summary>
    [GoogleCloudResourceType("google-cloud:bigqueryreservation/v1:Reservation")]
    public partial class Reservation : Pulumi.CustomResource
    {
        /// <summary>
        /// Creation time of the reservation.
        /// </summary>
        [Output("creationTime")]
        public Output<string> CreationTime { get; private set; } = null!;

        /// <summary>
        /// If false, any query using this reservation will use idle slots from other reservations within the same admin project. If true, a query using this reservation will execute with the slot capacity specified above at most.
        /// </summary>
        [Output("ignoreIdleSlots")]
        public Output<bool> IgnoreIdleSlots { get; private set; } = null!;

        /// <summary>
        /// The resource name of the reservation, e.g., `projects/*/locations/*/reservations/team1-prod`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Minimum slots available to this reservation. A slot is a unit of computational power in BigQuery, and serves as the unit of parallelism. Queries using this reservation might use more slots during runtime if ignore_idle_slots is set to false. If the new reservation's slot capacity exceed the parent's slot capacity or if total slot capacity of the new reservation and its siblings exceeds the parent's slot capacity, the request will fail with `google.rpc.Code.RESOURCE_EXHAUSTED`.
        /// </summary>
        [Output("slotCapacity")]
        public Output<string> SlotCapacity { get; private set; } = null!;

        /// <summary>
        /// Last update time of the reservation.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a Reservation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Reservation(string name, ReservationArgs args, CustomResourceOptions? options = null)
            : base("google-cloud:bigqueryreservation/v1:Reservation", name, args ?? new ReservationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Reservation(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-cloud:bigqueryreservation/v1:Reservation", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Reservation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Reservation Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Reservation(name, id, options);
        }
    }

    public sealed class ReservationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// If false, any query using this reservation will use idle slots from other reservations within the same admin project. If true, a query using this reservation will execute with the slot capacity specified above at most.
        /// </summary>
        [Input("ignoreIdleSlots")]
        public Input<bool>? IgnoreIdleSlots { get; set; }

        [Input("locationsId", required: true)]
        public Input<string> LocationsId { get; set; } = null!;

        /// <summary>
        /// The resource name of the reservation, e.g., `projects/*/locations/*/reservations/team1-prod`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("projectsId", required: true)]
        public Input<string> ProjectsId { get; set; } = null!;

        [Input("reservationsId", required: true)]
        public Input<string> ReservationsId { get; set; } = null!;

        /// <summary>
        /// Minimum slots available to this reservation. A slot is a unit of computational power in BigQuery, and serves as the unit of parallelism. Queries using this reservation might use more slots during runtime if ignore_idle_slots is set to false. If the new reservation's slot capacity exceed the parent's slot capacity or if total slot capacity of the new reservation and its siblings exceeds the parent's slot capacity, the request will fail with `google.rpc.Code.RESOURCE_EXHAUSTED`.
        /// </summary>
        [Input("slotCapacity")]
        public Input<string>? SlotCapacity { get; set; }

        public ReservationArgs()
        {
        }
    }
}