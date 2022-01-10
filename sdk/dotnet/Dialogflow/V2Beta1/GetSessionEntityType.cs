// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.GoogleNative.Dialogflow.V2Beta1
{
    public static class GetSessionEntityType
    {
        /// <summary>
        /// Retrieves the specified session entity type. This method doesn't work with Google Assistant integration. Contact Dialogflow support if you need to use session entities with Google Assistant integration.
        /// </summary>
        public static Task<GetSessionEntityTypeResult> InvokeAsync(GetSessionEntityTypeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSessionEntityTypeResult>("google-native:dialogflow/v2beta1:getSessionEntityType", args ?? new GetSessionEntityTypeArgs(), options.WithVersion());

        /// <summary>
        /// Retrieves the specified session entity type. This method doesn't work with Google Assistant integration. Contact Dialogflow support if you need to use session entities with Google Assistant integration.
        /// </summary>
        public static Output<GetSessionEntityTypeResult> Invoke(GetSessionEntityTypeInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetSessionEntityTypeResult>("google-native:dialogflow/v2beta1:getSessionEntityType", args ?? new GetSessionEntityTypeInvokeArgs(), options.WithVersion());
    }


    public sealed class GetSessionEntityTypeArgs : Pulumi.InvokeArgs
    {
        [Input("entityTypeId", required: true)]
        public string EntityTypeId { get; set; } = null!;

        [Input("environmentId", required: true)]
        public string EnvironmentId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("sessionId", required: true)]
        public string SessionId { get; set; } = null!;

        [Input("userId", required: true)]
        public string UserId { get; set; } = null!;

        public GetSessionEntityTypeArgs()
        {
        }
    }

    public sealed class GetSessionEntityTypeInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("entityTypeId", required: true)]
        public Input<string> EntityTypeId { get; set; } = null!;

        [Input("environmentId", required: true)]
        public Input<string> EnvironmentId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("sessionId", required: true)]
        public Input<string> SessionId { get; set; } = null!;

        [Input("userId", required: true)]
        public Input<string> UserId { get; set; } = null!;

        public GetSessionEntityTypeInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSessionEntityTypeResult
    {
        /// <summary>
        /// The collection of entities associated with this session entity type.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudDialogflowV2beta1EntityTypeEntityResponse> Entities;
        /// <summary>
        /// Indicates whether the additional data should override or supplement the custom entity type definition.
        /// </summary>
        public readonly string EntityOverrideMode;
        /// <summary>
        /// The unique identifier of this session entity type. Supported formats: - `projects//agent/sessions//entityTypes/` - `projects//locations//agent/sessions//entityTypes/` - `projects//agent/environments//users//sessions//entityTypes/` - `projects//locations//agent/environments/ /users//sessions//entityTypes/` If `Location ID` is not specified we assume default 'us' location. If `Environment ID` is not specified, we assume default 'draft' environment. If `User ID` is not specified, we assume default '-' user. `` must be the display name of an existing entity type in the same agent that will be overridden or supplemented.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetSessionEntityTypeResult(
            ImmutableArray<Outputs.GoogleCloudDialogflowV2beta1EntityTypeEntityResponse> entities,

            string entityOverrideMode,

            string name)
        {
            Entities = entities;
            EntityOverrideMode = entityOverrideMode;
            Name = name;
        }
    }
}