// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.Testing.V1.Outputs
{

    [OutputType]
    public sealed class AccountResponse
    {
        /// <summary>
        /// An automatic google login account.
        /// </summary>
        public readonly Outputs.GoogleAutoResponse GoogleAuto;

        [OutputConstructor]
        private AccountResponse(Outputs.GoogleAutoResponse googleAuto)
        {
            GoogleAuto = googleAuto;
        }
    }
}