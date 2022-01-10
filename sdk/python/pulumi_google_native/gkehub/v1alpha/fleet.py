# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = ['FleetArgs', 'Fleet']

@pulumi.input_type
class FleetArgs:
    def __init__(__self__, *,
                 display_name: Optional[pulumi.Input[str]] = None,
                 fleet_name: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Fleet resource.
        :param pulumi.Input[str] display_name: Optional. A user-assigned display name of the Fleet. When present, it must be between 4 to 30 characters. Allowed characters are: lowercase and uppercase letters, numbers, hyphen, single-quote, double-quote, space, and exclamation point. Example: `Production Fleet`
        :param pulumi.Input[str] fleet_name: The name for the fleet. The name must meet the following constraints: + The name of a fleet should be unique within the organization; + It must consist of lower case alphanumeric characters or `-`; + The length of the name must be less than or equal to 63; + Unicode names must be expressed in Punycode format (rfc3492). Examples: + prod-fleet + xn--wlq33vhyw9jb （Punycode form for "生产环境")
        """
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if fleet_name is not None:
            pulumi.set(__self__, "fleet_name", fleet_name)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. A user-assigned display name of the Fleet. When present, it must be between 4 to 30 characters. Allowed characters are: lowercase and uppercase letters, numbers, hyphen, single-quote, double-quote, space, and exclamation point. Example: `Production Fleet`
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="fleetName")
    def fleet_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name for the fleet. The name must meet the following constraints: + The name of a fleet should be unique within the organization; + It must consist of lower case alphanumeric characters or `-`; + The length of the name must be less than or equal to 63; + Unicode names must be expressed in Punycode format (rfc3492). Examples: + prod-fleet + xn--wlq33vhyw9jb （Punycode form for "生产环境")
        """
        return pulumi.get(self, "fleet_name")

    @fleet_name.setter
    def fleet_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fleet_name", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)


class Fleet(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 fleet_name: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a fleet.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] display_name: Optional. A user-assigned display name of the Fleet. When present, it must be between 4 to 30 characters. Allowed characters are: lowercase and uppercase letters, numbers, hyphen, single-quote, double-quote, space, and exclamation point. Example: `Production Fleet`
        :param pulumi.Input[str] fleet_name: The name for the fleet. The name must meet the following constraints: + The name of a fleet should be unique within the organization; + It must consist of lower case alphanumeric characters or `-`; + The length of the name must be less than or equal to 63; + Unicode names must be expressed in Punycode format (rfc3492). Examples: + prod-fleet + xn--wlq33vhyw9jb （Punycode form for "生产环境")
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[FleetArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a fleet.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param FleetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FleetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 fleet_name: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
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
            __props__ = FleetArgs.__new__(FleetArgs)

            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["fleet_name"] = fleet_name
            __props__.__dict__["location"] = location
            __props__.__dict__["project"] = project
            __props__.__dict__["create_time"] = None
            __props__.__dict__["delete_time"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["uid"] = None
            __props__.__dict__["update_time"] = None
        super(Fleet, __self__).__init__(
            'google-native:gkehub/v1alpha:Fleet',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Fleet':
        """
        Get an existing Fleet resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = FleetArgs.__new__(FleetArgs)

        __props__.__dict__["create_time"] = None
        __props__.__dict__["delete_time"] = None
        __props__.__dict__["display_name"] = None
        __props__.__dict__["fleet_name"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["uid"] = None
        __props__.__dict__["update_time"] = None
        return Fleet(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        When the Fleet was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="deleteTime")
    def delete_time(self) -> pulumi.Output[str]:
        """
        When the Fleet was deleted.
        """
        return pulumi.get(self, "delete_time")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        Optional. A user-assigned display name of the Fleet. When present, it must be between 4 to 30 characters. Allowed characters are: lowercase and uppercase letters, numbers, hyphen, single-quote, double-quote, space, and exclamation point. Example: `Production Fleet`
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="fleetName")
    def fleet_name(self) -> pulumi.Output[str]:
        """
        The name for the fleet. The name must meet the following constraints: + The name of a fleet should be unique within the organization; + It must consist of lower case alphanumeric characters or `-`; + The length of the name must be less than or equal to 63; + Unicode names must be expressed in Punycode format (rfc3492). Examples: + prod-fleet + xn--wlq33vhyw9jb （Punycode form for "生产环境")
        """
        return pulumi.get(self, "fleet_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The full, unique resource name of this fleet in the format of `projects/{project}/locations/{location}/fleets/{fleet}`. Each GCP project can have at most one fleet resource, named "default".
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def uid(self) -> pulumi.Output[str]:
        """
        Google-generated UUID for this resource. This is unique across all Fleet resources. If a Fleet resource is deleted and another resource with the same name is created, it gets a different uid.
        """
        return pulumi.get(self, "uid")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        When the Fleet was last updated.
        """
        return pulumi.get(self, "update_time")
