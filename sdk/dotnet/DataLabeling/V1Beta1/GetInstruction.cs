// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.GoogleNative.DataLabeling.V1Beta1
{
    public static class GetInstruction
    {
        /// <summary>
        /// Gets an instruction by resource name.
        /// </summary>
        public static Task<GetInstructionResult> InvokeAsync(GetInstructionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetInstructionResult>("google-native:datalabeling/v1beta1:getInstruction", args ?? new GetInstructionArgs(), options.WithVersion());

        /// <summary>
        /// Gets an instruction by resource name.
        /// </summary>
        public static Output<GetInstructionResult> Invoke(GetInstructionInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetInstructionResult>("google-native:datalabeling/v1beta1:getInstruction", args ?? new GetInstructionInvokeArgs(), options.WithVersion());
    }


    public sealed class GetInstructionArgs : Pulumi.InvokeArgs
    {
        [Input("instructionId", required: true)]
        public string InstructionId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetInstructionArgs()
        {
        }
    }

    public sealed class GetInstructionInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("instructionId", required: true)]
        public Input<string> InstructionId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetInstructionInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetInstructionResult
    {
        /// <summary>
        /// The names of any related resources that are blocking changes to the instruction.
        /// </summary>
        public readonly ImmutableArray<string> BlockingResources;
        /// <summary>
        /// Creation time of instruction.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The data type of this instruction.
        /// </summary>
        public readonly string DataType;
        /// <summary>
        /// Optional. User-provided description of the instruction. The description can be up to 10000 characters long.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The display name of the instruction. Maximum of 64 characters.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Instruction resource name, format: projects/{project_id}/instructions/{instruction_id}
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Instruction from a PDF document. The PDF should be in a Cloud Storage bucket.
        /// </summary>
        public readonly Outputs.GoogleCloudDatalabelingV1beta1PdfInstructionResponse PdfInstruction;
        /// <summary>
        /// Last update time of instruction.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetInstructionResult(
            ImmutableArray<string> blockingResources,

            string createTime,

            string dataType,

            string description,

            string displayName,

            string name,

            Outputs.GoogleCloudDatalabelingV1beta1PdfInstructionResponse pdfInstruction,

            string updateTime)
        {
            BlockingResources = blockingResources;
            CreateTime = createTime;
            DataType = dataType;
            Description = description;
            DisplayName = displayName;
            Name = name;
            PdfInstruction = pdfInstruction;
            UpdateTime = updateTime;
        }
    }
}
