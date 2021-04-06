# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .registry import *
from .registry_device import *
from .registry_group_iam_policy import *
from .registry_iam_policy import *
from ._inputs import *
from . import outputs

def _register_module():
    import pulumi
    from ... import _utilities


    class Module(pulumi.runtime.ResourceModule):
        _version = _utilities.get_semver_version()

        def version(self):
            return Module._version

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "google-cloud:cloudiot/v1:Registry":
                return Registry(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "google-cloud:cloudiot/v1:RegistryDevice":
                return RegistryDevice(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "google-cloud:cloudiot/v1:RegistryGroupIamPolicy":
                return RegistryGroupIamPolicy(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "google-cloud:cloudiot/v1:RegistryIamPolicy":
                return RegistryIamPolicy(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("google-cloud", "cloudiot/v1", _module_instance)

_register_module()