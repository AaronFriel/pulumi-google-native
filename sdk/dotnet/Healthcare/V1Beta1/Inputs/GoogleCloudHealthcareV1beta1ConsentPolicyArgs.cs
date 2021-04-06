// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.Healthcare.V1Beta1.Inputs
{

    /// <summary>
    /// Represents a user's consent in terms of the resources that can be accessed and under what conditions.
    /// </summary>
    public sealed class GoogleCloudHealthcareV1beta1ConsentPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Required. The request conditions to meet to grant access. In addition to any supported comparison operators, authorization rules may have `IN` operator as well as at most 10 logical operators that are limited to `AND` (`&amp;&amp;`), `OR` (`||`).
        /// </summary>
        [Input("authorizationRule")]
        public Input<Inputs.ExprArgs>? AuthorizationRule { get; set; }

        [Input("resourceAttributes")]
        private InputList<Inputs.AttributeArgs>? _resourceAttributes;

        /// <summary>
        /// The resources that this policy applies to. A resource is a match if it matches all the attributes listed here. If empty, this policy applies to all User data mappings for the given user.
        /// </summary>
        public InputList<Inputs.AttributeArgs> ResourceAttributes
        {
            get => _resourceAttributes ?? (_resourceAttributes = new InputList<Inputs.AttributeArgs>());
            set => _resourceAttributes = value;
        }

        public GoogleCloudHealthcareV1beta1ConsentPolicyArgs()
        {
        }
    }
}