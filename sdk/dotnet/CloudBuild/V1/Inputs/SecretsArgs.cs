// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudBuild.V1.Inputs
{

    /// <summary>
    /// Secrets and secret environment variables.
    /// </summary>
    public sealed class SecretsArgs : Pulumi.ResourceArgs
    {
        [Input("inline")]
        private InputList<Inputs.InlineSecretArgs>? _inline;

        /// <summary>
        /// Secrets encrypted with KMS key and the associated secret environment variable.
        /// </summary>
        public InputList<Inputs.InlineSecretArgs> Inline
        {
            get => _inline ?? (_inline = new InputList<Inputs.InlineSecretArgs>());
            set => _inline = value;
        }

        [Input("secretManager")]
        private InputList<Inputs.SecretManagerSecretArgs>? _secretManager;

        /// <summary>
        /// Secrets in Secret Manager and associated secret environment variable.
        /// </summary>
        public InputList<Inputs.SecretManagerSecretArgs> SecretManager
        {
            get => _secretManager ?? (_secretManager = new InputList<Inputs.SecretManagerSecretArgs>());
            set => _secretManager = value;
        }

        public SecretsArgs()
        {
        }
    }
}
