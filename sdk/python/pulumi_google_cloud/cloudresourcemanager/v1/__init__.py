# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .lien import *
from .organization_iam_policy import *
from .project import *
from .project_iam_policy import *
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
            if typ == "google-cloud:cloudresourcemanager/v1:Lien":
                return Lien(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "google-cloud:cloudresourcemanager/v1:OrganizationIamPolicy":
                return OrganizationIamPolicy(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "google-cloud:cloudresourcemanager/v1:Project":
                return Project(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "google-cloud:cloudresourcemanager/v1:ProjectIamPolicy":
                return ProjectIamPolicy(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("google-cloud", "cloudresourcemanager/v1", _module_instance)

_register_module()