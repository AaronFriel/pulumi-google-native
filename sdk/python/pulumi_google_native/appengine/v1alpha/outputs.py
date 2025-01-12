# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = [
    'CertificateRawDataResponse',
    'ManagedCertificateResponse',
    'ResourceRecordResponse',
    'SslSettingsResponse',
]

@pulumi.output_type
class CertificateRawDataResponse(dict):
    """
    An SSL certificate obtained from a certificate authority.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "privateKey":
            suggest = "private_key"
        elif key == "publicCertificate":
            suggest = "public_certificate"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in CertificateRawDataResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        CertificateRawDataResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        CertificateRawDataResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 private_key: str,
                 public_certificate: str):
        """
        An SSL certificate obtained from a certificate authority.
        :param str private_key: Unencrypted PEM encoded RSA private key. This field is set once on certificate creation and then encrypted. The key size must be 2048 bits or fewer. Must include the header and footer. Example: -----BEGIN RSA PRIVATE KEY----- -----END RSA PRIVATE KEY----- @InputOnly
        :param str public_certificate: PEM encoded x.509 public key certificate. This field is set once on certificate creation. Must include the header and footer. Example: -----BEGIN CERTIFICATE----- -----END CERTIFICATE----- 
        """
        pulumi.set(__self__, "private_key", private_key)
        pulumi.set(__self__, "public_certificate", public_certificate)

    @property
    @pulumi.getter(name="privateKey")
    def private_key(self) -> str:
        """
        Unencrypted PEM encoded RSA private key. This field is set once on certificate creation and then encrypted. The key size must be 2048 bits or fewer. Must include the header and footer. Example: -----BEGIN RSA PRIVATE KEY----- -----END RSA PRIVATE KEY----- @InputOnly
        """
        return pulumi.get(self, "private_key")

    @property
    @pulumi.getter(name="publicCertificate")
    def public_certificate(self) -> str:
        """
        PEM encoded x.509 public key certificate. This field is set once on certificate creation. Must include the header and footer. Example: -----BEGIN CERTIFICATE----- -----END CERTIFICATE----- 
        """
        return pulumi.get(self, "public_certificate")


@pulumi.output_type
class ManagedCertificateResponse(dict):
    """
    A certificate managed by App Engine.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "lastRenewalTime":
            suggest = "last_renewal_time"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ManagedCertificateResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ManagedCertificateResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ManagedCertificateResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 last_renewal_time: str,
                 status: str):
        """
        A certificate managed by App Engine.
        :param str last_renewal_time: Time at which the certificate was last renewed. The renewal process is fully managed. Certificate renewal will automatically occur before the certificate expires. Renewal errors can be tracked via ManagementStatus.
        :param str status: Status of certificate management. Refers to the most recent certificate acquisition or renewal attempt.
        """
        pulumi.set(__self__, "last_renewal_time", last_renewal_time)
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="lastRenewalTime")
    def last_renewal_time(self) -> str:
        """
        Time at which the certificate was last renewed. The renewal process is fully managed. Certificate renewal will automatically occur before the certificate expires. Renewal errors can be tracked via ManagementStatus.
        """
        return pulumi.get(self, "last_renewal_time")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Status of certificate management. Refers to the most recent certificate acquisition or renewal attempt.
        """
        return pulumi.get(self, "status")


@pulumi.output_type
class ResourceRecordResponse(dict):
    """
    A DNS resource record.
    """
    def __init__(__self__, *,
                 name: str,
                 rrdata: str,
                 type: str):
        """
        A DNS resource record.
        :param str name: Relative name of the object affected by this record. Only applicable for CNAME records. Example: 'www'.
        :param str rrdata: Data for this record. Values vary by record type, as defined in RFC 1035 (section 5) and RFC 1034 (section 3.6.1).
        :param str type: Resource record type. Example: AAAA.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "rrdata", rrdata)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Relative name of the object affected by this record. Only applicable for CNAME records. Example: 'www'.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def rrdata(self) -> str:
        """
        Data for this record. Values vary by record type, as defined in RFC 1035 (section 5) and RFC 1034 (section 3.6.1).
        """
        return pulumi.get(self, "rrdata")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Resource record type. Example: AAAA.
        """
        return pulumi.get(self, "type")


@pulumi.output_type
class SslSettingsResponse(dict):
    """
    SSL configuration for a DomainMapping resource.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "certificateId":
            suggest = "certificate_id"
        elif key == "isManagedCertificate":
            suggest = "is_managed_certificate"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in SslSettingsResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        SslSettingsResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        SslSettingsResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 certificate_id: str,
                 is_managed_certificate: bool):
        """
        SSL configuration for a DomainMapping resource.
        :param str certificate_id: ID of the AuthorizedCertificate resource configuring SSL for the application. Clearing this field will remove SSL support.By default, a managed certificate is automatically created for every domain mapping. To omit SSL support or to configure SSL manually, specify no_managed_certificate on a CREATE or UPDATE request. You must be authorized to administer the AuthorizedCertificate resource to manually map it to a DomainMapping resource. Example: 12345.
        :param bool is_managed_certificate: Whether the mapped certificate is an App Engine managed certificate. Managed certificates are created by default with a domain mapping. To opt out, specify no_managed_certificate on a CREATE or UPDATE request.
        """
        pulumi.set(__self__, "certificate_id", certificate_id)
        pulumi.set(__self__, "is_managed_certificate", is_managed_certificate)

    @property
    @pulumi.getter(name="certificateId")
    def certificate_id(self) -> str:
        """
        ID of the AuthorizedCertificate resource configuring SSL for the application. Clearing this field will remove SSL support.By default, a managed certificate is automatically created for every domain mapping. To omit SSL support or to configure SSL manually, specify no_managed_certificate on a CREATE or UPDATE request. You must be authorized to administer the AuthorizedCertificate resource to manually map it to a DomainMapping resource. Example: 12345.
        """
        return pulumi.get(self, "certificate_id")

    @property
    @pulumi.getter(name="isManagedCertificate")
    def is_managed_certificate(self) -> bool:
        """
        Whether the mapped certificate is an App Engine managed certificate. Managed certificates are created by default with a domain mapping. To opt out, specify no_managed_certificate on a CREATE or UPDATE request.
        """
        return pulumi.get(self, "is_managed_certificate")


