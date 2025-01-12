// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Transcoder.V1Beta1.Outputs
{

    /// <summary>
    /// Configuration for MPEG Common Encryption (MPEG-CENC).
    /// </summary>
    [OutputType]
    public sealed class MpegCommonEncryptionResponse
    {
        /// <summary>
        /// 128 bit Key ID represented as lowercase hexadecimal digits for use with common encryption.
        /// </summary>
        public readonly string KeyId;
        /// <summary>
        /// Specify the encryption scheme. Supported encryption schemes: - 'cenc' - 'cbcs'
        /// </summary>
        public readonly string Scheme;

        [OutputConstructor]
        private MpegCommonEncryptionResponse(
            string keyId,

            string scheme)
        {
            KeyId = keyId;
            Scheme = scheme;
        }
    }
}
