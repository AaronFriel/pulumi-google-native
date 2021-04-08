# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['DatasetAnnotationStoreAnnotation']


class DatasetAnnotationStoreAnnotation(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 annotation_source: Optional[pulumi.Input[pulumi.InputType['AnnotationSourceArgs']]] = None,
                 annotation_stores_id: Optional[pulumi.Input[str]] = None,
                 annotations_id: Optional[pulumi.Input[str]] = None,
                 custom_data: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 datasets_id: Optional[pulumi.Input[str]] = None,
                 image_annotation: Optional[pulumi.Input[pulumi.InputType['ImageAnnotationArgs']]] = None,
                 locations_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 projects_id: Optional[pulumi.Input[str]] = None,
                 resource_annotation: Optional[pulumi.Input[pulumi.InputType['ResourceAnnotationArgs']]] = None,
                 text_annotation: Optional[pulumi.Input[pulumi.InputType['SensitiveTextAnnotationArgs']]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Creates a new Annotation record. It is valid to create Annotation objects for the same source more than once since a unique ID is assigned to each record by this service.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['AnnotationSourceArgs']] annotation_source: Details of the source.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] custom_data: Additional information for this annotation record, such as annotator and verifier information or study campaign.
        :param pulumi.Input[pulumi.InputType['ImageAnnotationArgs']] image_annotation: Annotations for images. For example, bounding polygons.
        :param pulumi.Input[str] name: Resource name of the Annotation, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/annotationStores/{annotation_store_id}/annotations/{annotation_id}`.
        :param pulumi.Input[pulumi.InputType['ResourceAnnotationArgs']] resource_annotation: Annotations for resource. For example, classification tags.
        :param pulumi.Input[pulumi.InputType['SensitiveTextAnnotationArgs']] text_annotation: Annotations for sensitive texts. For example, a range that describes the location of sensitive text.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['annotation_source'] = annotation_source
            if annotation_stores_id is None and not opts.urn:
                raise TypeError("Missing required property 'annotation_stores_id'")
            __props__['annotation_stores_id'] = annotation_stores_id
            if annotations_id is None and not opts.urn:
                raise TypeError("Missing required property 'annotations_id'")
            __props__['annotations_id'] = annotations_id
            __props__['custom_data'] = custom_data
            if datasets_id is None and not opts.urn:
                raise TypeError("Missing required property 'datasets_id'")
            __props__['datasets_id'] = datasets_id
            __props__['image_annotation'] = image_annotation
            if locations_id is None and not opts.urn:
                raise TypeError("Missing required property 'locations_id'")
            __props__['locations_id'] = locations_id
            __props__['name'] = name
            if projects_id is None and not opts.urn:
                raise TypeError("Missing required property 'projects_id'")
            __props__['projects_id'] = projects_id
            __props__['resource_annotation'] = resource_annotation
            __props__['text_annotation'] = text_annotation
        super(DatasetAnnotationStoreAnnotation, __self__).__init__(
            'gcp-native:healthcare/v1beta1:DatasetAnnotationStoreAnnotation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'DatasetAnnotationStoreAnnotation':
        """
        Get an existing DatasetAnnotationStoreAnnotation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["annotation_source"] = None
        __props__["custom_data"] = None
        __props__["image_annotation"] = None
        __props__["name"] = None
        __props__["resource_annotation"] = None
        __props__["text_annotation"] = None
        return DatasetAnnotationStoreAnnotation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="annotationSource")
    def annotation_source(self) -> pulumi.Output['outputs.AnnotationSourceResponse']:
        """
        Details of the source.
        """
        return pulumi.get(self, "annotation_source")

    @property
    @pulumi.getter(name="customData")
    def custom_data(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Additional information for this annotation record, such as annotator and verifier information or study campaign.
        """
        return pulumi.get(self, "custom_data")

    @property
    @pulumi.getter(name="imageAnnotation")
    def image_annotation(self) -> pulumi.Output['outputs.ImageAnnotationResponse']:
        """
        Annotations for images. For example, bounding polygons.
        """
        return pulumi.get(self, "image_annotation")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name of the Annotation, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/annotationStores/{annotation_store_id}/annotations/{annotation_id}`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceAnnotation")
    def resource_annotation(self) -> pulumi.Output['outputs.ResourceAnnotationResponse']:
        """
        Annotations for resource. For example, classification tags.
        """
        return pulumi.get(self, "resource_annotation")

    @property
    @pulumi.getter(name="textAnnotation")
    def text_annotation(self) -> pulumi.Output['outputs.SensitiveTextAnnotationResponse']:
        """
        Annotations for sensitive texts. For example, a range that describes the location of sensitive text.
        """
        return pulumi.get(self, "text_annotation")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
