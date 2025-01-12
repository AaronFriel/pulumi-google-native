// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.OSConfig.V1.Outputs
{

    /// <summary>
    /// A package managed by APT. - install: `apt-get update &amp;&amp; apt-get -y install [name]` - remove: `apt-get -y remove [name]`
    /// </summary>
    [OutputType]
    public sealed class OSPolicyResourcePackageResourceAPTResponse
    {
        /// <summary>
        /// Package name.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private OSPolicyResourcePackageResourceAPTResponse(string name)
        {
            Name = name;
        }
    }
}
