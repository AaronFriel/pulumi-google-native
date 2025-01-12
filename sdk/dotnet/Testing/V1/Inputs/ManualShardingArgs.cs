// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Testing.V1.Inputs
{

    /// <summary>
    /// Shards test cases into the specified groups of packages, classes, and/or methods. With manual sharding enabled, specifying test targets via environment_variables or in InstrumentationTest is invalid.
    /// </summary>
    public sealed class ManualShardingArgs : Pulumi.ResourceArgs
    {
        [Input("testTargetsForShard", required: true)]
        private InputList<Inputs.TestTargetsForShardArgs>? _testTargetsForShard;

        /// <summary>
        /// Group of packages, classes, and/or test methods to be run for each shard. When any physical devices are selected, the number of test_targets_for_shard must be &gt;= 1 and &lt;= 50. When no physical devices are selected, the number must be &gt;= 1 and &lt;= 500.
        /// </summary>
        public InputList<Inputs.TestTargetsForShardArgs> TestTargetsForShard
        {
            get => _testTargetsForShard ?? (_testTargetsForShard = new InputList<Inputs.TestTargetsForShardArgs>());
            set => _testTargetsForShard = value;
        }

        public ManualShardingArgs()
        {
        }
    }
}
