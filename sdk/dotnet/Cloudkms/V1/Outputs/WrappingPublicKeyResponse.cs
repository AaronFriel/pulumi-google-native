// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Cloudkms.V1.Outputs
{

    /// <summary>
    /// The public key component of the wrapping key. For details of the type of key this public key corresponds to, see the ImportMethod.
    /// </summary>
    [OutputType]
    public sealed class WrappingPublicKeyResponse
    {
        /// <summary>
        /// The public key, encoded in PEM format. For more information, see the [RFC 7468](https://tools.ietf.org/html/rfc7468) sections for [General Considerations](https://tools.ietf.org/html/rfc7468#section-2) and [Textual Encoding of Subject Public Key Info] (https://tools.ietf.org/html/rfc7468#section-13).
        /// </summary>
        public readonly string Pem;

        [OutputConstructor]
        private WrappingPublicKeyResponse(string pem)
        {
            Pem = pem;
        }
    }
}
