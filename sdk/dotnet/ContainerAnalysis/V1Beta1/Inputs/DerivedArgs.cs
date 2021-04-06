// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.ContainerAnalysis.V1Beta1.Inputs
{

    /// <summary>
    /// Derived describes the derived image portion (Occurrence) of the DockerImage relationship. This image would be produced from a Dockerfile with FROM .
    /// </summary>
    public sealed class DerivedArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// This contains the base image URL for the derived image occurrence.
        /// </summary>
        [Input("baseResourceUrl")]
        public Input<string>? BaseResourceUrl { get; set; }

        /// <summary>
        /// The number of layers by which this image differs from the associated image basis.
        /// </summary>
        [Input("distance")]
        public Input<int>? Distance { get; set; }

        /// <summary>
        /// Required. The fingerprint of the derived image.
        /// </summary>
        [Input("fingerprint")]
        public Input<Inputs.FingerprintArgs>? Fingerprint { get; set; }

        [Input("layerInfo")]
        private InputList<Inputs.LayerArgs>? _layerInfo;

        /// <summary>
        /// This contains layer-specific metadata, if populated it has length "distance" and is ordered with [distance] being the layer immediately following the base image and [1] being the final layer.
        /// </summary>
        public InputList<Inputs.LayerArgs> LayerInfo
        {
            get => _layerInfo ?? (_layerInfo = new InputList<Inputs.LayerArgs>());
            set => _layerInfo = value;
        }

        public DerivedArgs()
        {
        }
    }
}