// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Inputs
{

    /// <summary>
    /// [Deprecated] The TLS settings for the client or server. The TLS settings for the client or server.
    /// </summary>
    public sealed class TlsContextArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Defines the mechanism to obtain the client or server certificate.
        /// </summary>
        [Input("certificateContext")]
        public Input<Inputs.TlsCertificateContextArgs>? CertificateContext { get; set; }

        /// <summary>
        /// Defines the mechanism to obtain the Certificate Authority certificate to validate the client/server certificate. If omitted, the proxy will not validate the server or client certificate.
        /// </summary>
        [Input("validationContext")]
        public Input<Inputs.TlsValidationContextArgs>? ValidationContext { get; set; }

        public TlsContextArgs()
        {
        }
    }
}