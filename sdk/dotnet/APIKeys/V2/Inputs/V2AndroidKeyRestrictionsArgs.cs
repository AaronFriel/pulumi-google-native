// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.APIKeys.V2.Inputs
{

    /// <summary>
    /// The Android apps that are allowed to use the key.
    /// </summary>
    public sealed class V2AndroidKeyRestrictionsArgs : Pulumi.ResourceArgs
    {
        [Input("allowedApplications")]
        private InputList<Inputs.V2AndroidApplicationArgs>? _allowedApplications;

        /// <summary>
        /// A list of Android applications that are allowed to make API calls with this key.
        /// </summary>
        public InputList<Inputs.V2AndroidApplicationArgs> AllowedApplications
        {
            get => _allowedApplications ?? (_allowedApplications = new InputList<Inputs.V2AndroidApplicationArgs>());
            set => _allowedApplications = value;
        }

        public V2AndroidKeyRestrictionsArgs()
        {
        }
    }
}
