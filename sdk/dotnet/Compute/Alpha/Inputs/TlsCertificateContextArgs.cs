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
    /// [Deprecated] Defines the mechanism to obtain the client or server certificate. Defines the mechanism to obtain the client or server certificate.
    /// </summary>
    public sealed class TlsCertificateContextArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the certificate and private key paths. This field is applicable only if tlsCertificateSource is set to USE_PATH.
        /// </summary>
        [Input("certificatePaths")]
        public Input<Inputs.TlsCertificatePathsArgs>? CertificatePaths { get; set; }

        /// <summary>
        /// Defines how TLS certificates are obtained.
        /// </summary>
        [Input("certificateSource")]
        public Input<Pulumi.GoogleNative.Compute.Alpha.TlsCertificateContextCertificateSource>? CertificateSource { get; set; }

        /// <summary>
        /// Specifies the config to retrieve certificates through SDS. This field is applicable only if tlsCertificateSource is set to USE_SDS.
        /// </summary>
        [Input("sdsConfig")]
        public Input<Inputs.SdsConfigArgs>? SdsConfig { get; set; }

        public TlsCertificateContextArgs()
        {
        }
    }
}
