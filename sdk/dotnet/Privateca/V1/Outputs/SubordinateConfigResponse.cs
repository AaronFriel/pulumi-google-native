// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Privateca.V1.Outputs
{

    /// <summary>
    /// Describes a subordinate CA's issuers. This is either a resource name to a known issuing CertificateAuthority, or a PEM issuer certificate chain.
    /// </summary>
    [OutputType]
    public sealed class SubordinateConfigResponse
    {
        /// <summary>
        /// This can refer to a CertificateAuthority that was used to create a subordinate CertificateAuthority. This field is used for information and usability purposes only. The resource name is in the format `projects/*/locations/*/caPools/*/certificateAuthorities/*`.
        /// </summary>
        public readonly string CertificateAuthority;
        /// <summary>
        /// Contains the PEM certificate chain for the issuers of this CertificateAuthority, but not pem certificate for this CA itself.
        /// </summary>
        public readonly Outputs.SubordinateConfigChainResponse PemIssuerChain;

        [OutputConstructor]
        private SubordinateConfigResponse(
            string certificateAuthority,

            Outputs.SubordinateConfigChainResponse pemIssuerChain)
        {
            CertificateAuthority = certificateAuthority;
            PemIssuerChain = pemIssuerChain;
        }
    }
}