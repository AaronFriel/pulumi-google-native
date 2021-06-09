# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = ['EndpointArgs', 'Endpoint']

@pulumi.input_type
class EndpointArgs:
    def __init__(__self__, *,
                 endpoint_id: pulumi.Input[str],
                 location: pulumi.Input[str],
                 namespace_id: pulumi.Input[str],
                 project: pulumi.Input[str],
                 service_id: pulumi.Input[str],
                 address: Optional[pulumi.Input[str]] = None,
                 annotations: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a Endpoint resource.
        :param pulumi.Input[str] address: Optional. An IPv4 or IPv6 address. Service Directory rejects bad addresses like: * `8.8.8` * `8.8.8.8:53` * `test:bad:address` * `[::1]` * `[::1]:8080` Limited to 45 characters.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] annotations: Optional. Annotations for the endpoint. This data can be consumed by service clients. Restrictions: * The entire annotations dictionary may contain up to 512 characters, spread accoss all key-value pairs. Annotations that go beyond this limit are rejected * Valid annotation keys have two segments: an optional prefix and name, separated by a slash (/). The name segment is required and must be 63 characters or less, beginning and ending with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between. The prefix is optional. If specified, the prefix must be a DNS subdomain: a series of DNS labels separated by dots (.), not longer than 253 characters in total, followed by a slash (/) Annotations that fails to meet these requirements are rejected. * The `(*.)google.com/` and `(*.)googleapis.com/` prefixes are reserved for system annotations managed by Service Directory. If the user tries to write to these keyspaces, those entries are silently ignored by the system Note: This field is equivalent to the `metadata` field in the v1beta1 API. They have the same syntax and read/write to the same location in Service Directory.
        :param pulumi.Input[str] name: Immutable. The resource name for the endpoint in the format `projects/*/locations/*/namespaces/*/services/*/endpoints/*`.
        :param pulumi.Input[int] port: Optional. Service Directory rejects values outside of `[0, 65535]`.
        """
        pulumi.set(__self__, "endpoint_id", endpoint_id)
        pulumi.set(__self__, "location", location)
        pulumi.set(__self__, "namespace_id", namespace_id)
        pulumi.set(__self__, "project", project)
        pulumi.set(__self__, "service_id", service_id)
        if address is not None:
            pulumi.set(__self__, "address", address)
        if annotations is not None:
            pulumi.set(__self__, "annotations", annotations)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if port is not None:
            pulumi.set(__self__, "port", port)

    @property
    @pulumi.getter(name="endpointId")
    def endpoint_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "endpoint_id")

    @endpoint_id.setter
    def endpoint_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "endpoint_id", value)

    @property
    @pulumi.getter
    def location(self) -> pulumi.Input[str]:
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: pulumi.Input[str]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter(name="namespaceId")
    def namespace_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "namespace_id")

    @namespace_id.setter
    def namespace_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "namespace_id", value)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="serviceId")
    def service_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "service_id")

    @service_id.setter
    def service_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_id", value)

    @property
    @pulumi.getter
    def address(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. An IPv4 or IPv6 address. Service Directory rejects bad addresses like: * `8.8.8` * `8.8.8.8:53` * `test:bad:address` * `[::1]` * `[::1]:8080` Limited to 45 characters.
        """
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address", value)

    @property
    @pulumi.getter
    def annotations(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Optional. Annotations for the endpoint. This data can be consumed by service clients. Restrictions: * The entire annotations dictionary may contain up to 512 characters, spread accoss all key-value pairs. Annotations that go beyond this limit are rejected * Valid annotation keys have two segments: an optional prefix and name, separated by a slash (/). The name segment is required and must be 63 characters or less, beginning and ending with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between. The prefix is optional. If specified, the prefix must be a DNS subdomain: a series of DNS labels separated by dots (.), not longer than 253 characters in total, followed by a slash (/) Annotations that fails to meet these requirements are rejected. * The `(*.)google.com/` and `(*.)googleapis.com/` prefixes are reserved for system annotations managed by Service Directory. If the user tries to write to these keyspaces, those entries are silently ignored by the system Note: This field is equivalent to the `metadata` field in the v1beta1 API. They have the same syntax and read/write to the same location in Service Directory.
        """
        return pulumi.get(self, "annotations")

    @annotations.setter
    def annotations(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "annotations", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Immutable. The resource name for the endpoint in the format `projects/*/locations/*/namespaces/*/services/*/endpoints/*`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        Optional. Service Directory rejects values outside of `[0, 65535]`.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)


class Endpoint(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address: Optional[pulumi.Input[str]] = None,
                 annotations: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 endpoint_id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace_id: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 service_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates an endpoint, and returns the new endpoint.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address: Optional. An IPv4 or IPv6 address. Service Directory rejects bad addresses like: * `8.8.8` * `8.8.8.8:53` * `test:bad:address` * `[::1]` * `[::1]:8080` Limited to 45 characters.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] annotations: Optional. Annotations for the endpoint. This data can be consumed by service clients. Restrictions: * The entire annotations dictionary may contain up to 512 characters, spread accoss all key-value pairs. Annotations that go beyond this limit are rejected * Valid annotation keys have two segments: an optional prefix and name, separated by a slash (/). The name segment is required and must be 63 characters or less, beginning and ending with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between. The prefix is optional. If specified, the prefix must be a DNS subdomain: a series of DNS labels separated by dots (.), not longer than 253 characters in total, followed by a slash (/) Annotations that fails to meet these requirements are rejected. * The `(*.)google.com/` and `(*.)googleapis.com/` prefixes are reserved for system annotations managed by Service Directory. If the user tries to write to these keyspaces, those entries are silently ignored by the system Note: This field is equivalent to the `metadata` field in the v1beta1 API. They have the same syntax and read/write to the same location in Service Directory.
        :param pulumi.Input[str] name: Immutable. The resource name for the endpoint in the format `projects/*/locations/*/namespaces/*/services/*/endpoints/*`.
        :param pulumi.Input[int] port: Optional. Service Directory rejects values outside of `[0, 65535]`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EndpointArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates an endpoint, and returns the new endpoint.

        :param str resource_name: The name of the resource.
        :param EndpointArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EndpointArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address: Optional[pulumi.Input[str]] = None,
                 annotations: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 endpoint_id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace_id: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 service_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EndpointArgs.__new__(EndpointArgs)

            __props__.__dict__["address"] = address
            __props__.__dict__["annotations"] = annotations
            if endpoint_id is None and not opts.urn:
                raise TypeError("Missing required property 'endpoint_id'")
            __props__.__dict__["endpoint_id"] = endpoint_id
            if location is None and not opts.urn:
                raise TypeError("Missing required property 'location'")
            __props__.__dict__["location"] = location
            __props__.__dict__["name"] = name
            if namespace_id is None and not opts.urn:
                raise TypeError("Missing required property 'namespace_id'")
            __props__.__dict__["namespace_id"] = namespace_id
            __props__.__dict__["port"] = port
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            if service_id is None and not opts.urn:
                raise TypeError("Missing required property 'service_id'")
            __props__.__dict__["service_id"] = service_id
        super(Endpoint, __self__).__init__(
            'google-native:servicedirectory/v1:Endpoint',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Endpoint':
        """
        Get an existing Endpoint resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = EndpointArgs.__new__(EndpointArgs)

        __props__.__dict__["address"] = None
        __props__.__dict__["annotations"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["port"] = None
        return Endpoint(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def address(self) -> pulumi.Output[str]:
        """
        Optional. An IPv4 or IPv6 address. Service Directory rejects bad addresses like: * `8.8.8` * `8.8.8.8:53` * `test:bad:address` * `[::1]` * `[::1]:8080` Limited to 45 characters.
        """
        return pulumi.get(self, "address")

    @property
    @pulumi.getter
    def annotations(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Optional. Annotations for the endpoint. This data can be consumed by service clients. Restrictions: * The entire annotations dictionary may contain up to 512 characters, spread accoss all key-value pairs. Annotations that go beyond this limit are rejected * Valid annotation keys have two segments: an optional prefix and name, separated by a slash (/). The name segment is required and must be 63 characters or less, beginning and ending with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between. The prefix is optional. If specified, the prefix must be a DNS subdomain: a series of DNS labels separated by dots (.), not longer than 253 characters in total, followed by a slash (/) Annotations that fails to meet these requirements are rejected. * The `(*.)google.com/` and `(*.)googleapis.com/` prefixes are reserved for system annotations managed by Service Directory. If the user tries to write to these keyspaces, those entries are silently ignored by the system Note: This field is equivalent to the `metadata` field in the v1beta1 API. They have the same syntax and read/write to the same location in Service Directory.
        """
        return pulumi.get(self, "annotations")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Immutable. The resource name for the endpoint in the format `projects/*/locations/*/namespaces/*/services/*/endpoints/*`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[int]:
        """
        Optional. Service Directory rejects values outside of `[0, 65535]`.
        """
        return pulumi.get(self, "port")
