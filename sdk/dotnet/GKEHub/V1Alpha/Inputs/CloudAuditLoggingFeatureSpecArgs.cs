// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.GKEHub.V1Alpha.Inputs
{

    /// <summary>
    /// Spec for Audit Logging Allowlisting.
    /// </summary>
    public sealed class CloudAuditLoggingFeatureSpecArgs : Pulumi.ResourceArgs
    {
        [Input("allowlistedServiceAccounts")]
        private InputList<string>? _allowlistedServiceAccounts;

        /// <summary>
        /// Service account that should be allowlisted to send the audit logs; eg cloudauditlogging@gcp-project.iam.gserviceaccount.com. These accounts must already exist, but do not need to have any permissions granted to them. The customer's entitlements will be checked prior to allowlisting (i.e. the customer must be an Anthos customer.)
        /// </summary>
        public InputList<string> AllowlistedServiceAccounts
        {
            get => _allowlistedServiceAccounts ?? (_allowlistedServiceAccounts = new InputList<string>());
            set => _allowlistedServiceAccounts = value;
        }

        public CloudAuditLoggingFeatureSpecArgs()
        {
        }
    }
}