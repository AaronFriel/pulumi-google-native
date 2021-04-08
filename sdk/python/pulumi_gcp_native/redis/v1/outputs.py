# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables

__all__ = [
    'TlsCertificateResponse',
]

@pulumi.output_type
class TlsCertificateResponse(dict):
    """
    TlsCertificate Resource
    """
    def __init__(__self__, *,
                 cert: str,
                 create_time: str,
                 expire_time: str,
                 serial_number: str,
                 sha1_fingerprint: str):
        """
        TlsCertificate Resource
        :param str cert: PEM representation.
        :param str create_time: The time when the certificate was created in [RFC 3339](https://tools.ietf.org/html/rfc3339) format, for example `2020-05-18T00:00:00.094Z`.
        :param str expire_time: The time when the certificate expires in [RFC 3339](https://tools.ietf.org/html/rfc3339) format, for example `2020-05-18T00:00:00.094Z`.
        :param str serial_number: Serial number, as extracted from the certificate.
        :param str sha1_fingerprint: Sha1 Fingerprint of the certificate.
        """
        pulumi.set(__self__, "cert", cert)
        pulumi.set(__self__, "create_time", create_time)
        pulumi.set(__self__, "expire_time", expire_time)
        pulumi.set(__self__, "serial_number", serial_number)
        pulumi.set(__self__, "sha1_fingerprint", sha1_fingerprint)

    @property
    @pulumi.getter
    def cert(self) -> str:
        """
        PEM representation.
        """
        return pulumi.get(self, "cert")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The time when the certificate was created in [RFC 3339](https://tools.ietf.org/html/rfc3339) format, for example `2020-05-18T00:00:00.094Z`.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="expireTime")
    def expire_time(self) -> str:
        """
        The time when the certificate expires in [RFC 3339](https://tools.ietf.org/html/rfc3339) format, for example `2020-05-18T00:00:00.094Z`.
        """
        return pulumi.get(self, "expire_time")

    @property
    @pulumi.getter(name="serialNumber")
    def serial_number(self) -> str:
        """
        Serial number, as extracted from the certificate.
        """
        return pulumi.get(self, "serial_number")

    @property
    @pulumi.getter(name="sha1Fingerprint")
    def sha1_fingerprint(self) -> str:
        """
        Sha1 Fingerprint of the certificate.
        """
        return pulumi.get(self, "sha1_fingerprint")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

