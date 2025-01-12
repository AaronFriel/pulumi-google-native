// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ToolResults.V1Beta3.Inputs
{

    /// <summary>
    /// A test of an Android application that can control an Android component independently of its normal lifecycle. See for more information on types of Android tests.
    /// </summary>
    public sealed class AndroidInstrumentationTestArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The java package for the test to be executed. Required
        /// </summary>
        [Input("testPackageId")]
        public Input<string>? TestPackageId { get; set; }

        /// <summary>
        /// The InstrumentationTestRunner class. Required
        /// </summary>
        [Input("testRunnerClass")]
        public Input<string>? TestRunnerClass { get; set; }

        [Input("testTargets")]
        private InputList<string>? _testTargets;

        /// <summary>
        /// Each target must be fully qualified with the package name or class name, in one of these formats: - "package package_name" - "class package_name.class_name" - "class package_name.class_name#method_name" If empty, all targets in the module will be run.
        /// </summary>
        public InputList<string> TestTargets
        {
            get => _testTargets ?? (_testTargets = new InputList<string>());
            set => _testTargets = value;
        }

        /// <summary>
        /// The flag indicates whether Android Test Orchestrator will be used to run test or not.
        /// </summary>
        [Input("useOrchestrator")]
        public Input<bool>? UseOrchestrator { get; set; }

        public AndroidInstrumentationTestArgs()
        {
        }
    }
}
