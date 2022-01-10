// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DataCatalog.V1.Inputs
{

    /// <summary>
    /// Contact people for the entry.
    /// </summary>
    public sealed class GoogleCloudDatacatalogV1ContactsArgs : Pulumi.ResourceArgs
    {
        [Input("people")]
        private InputList<Inputs.GoogleCloudDatacatalogV1ContactsPersonArgs>? _people;

        /// <summary>
        /// The list of contact people for the entry.
        /// </summary>
        public InputList<Inputs.GoogleCloudDatacatalogV1ContactsPersonArgs> People
        {
            get => _people ?? (_people = new InputList<Inputs.GoogleCloudDatacatalogV1ContactsPersonArgs>());
            set => _people = value;
        }

        public GoogleCloudDatacatalogV1ContactsArgs()
        {
        }
    }
}