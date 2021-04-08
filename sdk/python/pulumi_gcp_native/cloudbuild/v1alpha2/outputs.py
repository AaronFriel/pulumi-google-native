# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables

__all__ = [
    'NetworkConfigResponse',
    'WorkerConfigResponse',
]

@pulumi.output_type
class NetworkConfigResponse(dict):
    """
    Network describes the network configuration for a `WorkerPool`.
    """
    def __init__(__self__, *,
                 peered_network: str):
        """
        Network describes the network configuration for a `WorkerPool`.
        :param str peered_network: Required. Immutable. The network definition that the workers are peered to. If this section is left empty, the workers will be peered to WorkerPool.project_id on the default network. Must be in the format `projects/{project}/global/networks/{network}`, where {project} is a project number, such as `12345`, and {network} is the name of a VPC network in the project.
        """
        pulumi.set(__self__, "peered_network", peered_network)

    @property
    @pulumi.getter(name="peeredNetwork")
    def peered_network(self) -> str:
        """
        Required. Immutable. The network definition that the workers are peered to. If this section is left empty, the workers will be peered to WorkerPool.project_id on the default network. Must be in the format `projects/{project}/global/networks/{network}`, where {project} is a project number, such as `12345`, and {network} is the name of a VPC network in the project.
        """
        return pulumi.get(self, "peered_network")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class WorkerConfigResponse(dict):
    """
    WorkerConfig defines the configuration to be used for a creating workers in the pool.
    """
    def __init__(__self__, *,
                 disk_size_gb: str,
                 machine_type: str):
        """
        WorkerConfig defines the configuration to be used for a creating workers in the pool.
        :param str disk_size_gb: Size of the disk attached to the worker, in GB. See https://cloud.google.com/compute/docs/disks/ If `0` is specified, Cloud Build will use a standard disk size.
        :param str machine_type: Machine Type of the worker, such as n1-standard-1. See https://cloud.google.com/compute/docs/machine-types. If left blank, Cloud Build will use a standard unspecified machine to create the worker pool.
        """
        pulumi.set(__self__, "disk_size_gb", disk_size_gb)
        pulumi.set(__self__, "machine_type", machine_type)

    @property
    @pulumi.getter(name="diskSizeGb")
    def disk_size_gb(self) -> str:
        """
        Size of the disk attached to the worker, in GB. See https://cloud.google.com/compute/docs/disks/ If `0` is specified, Cloud Build will use a standard disk size.
        """
        return pulumi.get(self, "disk_size_gb")

    @property
    @pulumi.getter(name="machineType")
    def machine_type(self) -> str:
        """
        Machine Type of the worker, such as n1-standard-1. See https://cloud.google.com/compute/docs/machine-types. If left blank, Cloud Build will use a standard unspecified machine to create the worker pool.
        """
        return pulumi.get(self, "machine_type")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

