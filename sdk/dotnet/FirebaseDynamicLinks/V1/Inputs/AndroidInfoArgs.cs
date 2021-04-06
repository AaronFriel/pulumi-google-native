// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.FirebaseDynamicLinks.V1.Inputs
{

    /// <summary>
    /// Android related attributes to the Dynamic Link.
    /// </summary>
    public sealed class AndroidInfoArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Link to open on Android if the app is not installed.
        /// </summary>
        [Input("androidFallbackLink")]
        public Input<string>? AndroidFallbackLink { get; set; }

        /// <summary>
        /// If specified, this overrides the ‘link’ parameter on Android.
        /// </summary>
        [Input("androidLink")]
        public Input<string>? AndroidLink { get; set; }

        /// <summary>
        /// Minimum version code for the Android app. If the installed app’s version code is lower, then the user is taken to the Play Store.
        /// </summary>
        [Input("androidMinPackageVersionCode")]
        public Input<string>? AndroidMinPackageVersionCode { get; set; }

        /// <summary>
        /// Android package name of the app.
        /// </summary>
        [Input("androidPackageName")]
        public Input<string>? AndroidPackageName { get; set; }

        public AndroidInfoArgs()
        {
        }
    }
}