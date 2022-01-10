# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs

__all__ = [
    'GetCompanyResult',
    'AwaitableGetCompanyResult',
    'get_company',
    'get_company_output',
]

@pulumi.output_type
class GetCompanyResult:
    def __init__(__self__, career_site_uri=None, derived_info=None, display_name=None, eeo_text=None, external_id=None, headquarters_address=None, hiring_agency=None, image_uri=None, keyword_searchable_job_custom_attributes=None, name=None, size=None, suspended=None, website_uri=None):
        if career_site_uri and not isinstance(career_site_uri, str):
            raise TypeError("Expected argument 'career_site_uri' to be a str")
        pulumi.set(__self__, "career_site_uri", career_site_uri)
        if derived_info and not isinstance(derived_info, dict):
            raise TypeError("Expected argument 'derived_info' to be a dict")
        pulumi.set(__self__, "derived_info", derived_info)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if eeo_text and not isinstance(eeo_text, str):
            raise TypeError("Expected argument 'eeo_text' to be a str")
        pulumi.set(__self__, "eeo_text", eeo_text)
        if external_id and not isinstance(external_id, str):
            raise TypeError("Expected argument 'external_id' to be a str")
        pulumi.set(__self__, "external_id", external_id)
        if headquarters_address and not isinstance(headquarters_address, str):
            raise TypeError("Expected argument 'headquarters_address' to be a str")
        pulumi.set(__self__, "headquarters_address", headquarters_address)
        if hiring_agency and not isinstance(hiring_agency, bool):
            raise TypeError("Expected argument 'hiring_agency' to be a bool")
        pulumi.set(__self__, "hiring_agency", hiring_agency)
        if image_uri and not isinstance(image_uri, str):
            raise TypeError("Expected argument 'image_uri' to be a str")
        pulumi.set(__self__, "image_uri", image_uri)
        if keyword_searchable_job_custom_attributes and not isinstance(keyword_searchable_job_custom_attributes, list):
            raise TypeError("Expected argument 'keyword_searchable_job_custom_attributes' to be a list")
        pulumi.set(__self__, "keyword_searchable_job_custom_attributes", keyword_searchable_job_custom_attributes)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if size and not isinstance(size, str):
            raise TypeError("Expected argument 'size' to be a str")
        pulumi.set(__self__, "size", size)
        if suspended and not isinstance(suspended, bool):
            raise TypeError("Expected argument 'suspended' to be a bool")
        pulumi.set(__self__, "suspended", suspended)
        if website_uri and not isinstance(website_uri, str):
            raise TypeError("Expected argument 'website_uri' to be a str")
        pulumi.set(__self__, "website_uri", website_uri)

    @property
    @pulumi.getter(name="careerSiteUri")
    def career_site_uri(self) -> str:
        """
        Optional. The URI to employer's career site or careers page on the employer's web site, for example, "https://careers.google.com".
        """
        return pulumi.get(self, "career_site_uri")

    @property
    @pulumi.getter(name="derivedInfo")
    def derived_info(self) -> 'outputs.CompanyDerivedInfoResponse':
        """
        Derived details about the company.
        """
        return pulumi.get(self, "derived_info")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        The display name of the company, for example, "Google LLC".
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="eeoText")
    def eeo_text(self) -> str:
        """
        Optional. Equal Employment Opportunity legal disclaimer text to be associated with all jobs, and typically to be displayed in all roles. The maximum number of allowed characters is 500.
        """
        return pulumi.get(self, "eeo_text")

    @property
    @pulumi.getter(name="externalId")
    def external_id(self) -> str:
        """
        Client side company identifier, used to uniquely identify the company. The maximum number of allowed characters is 255.
        """
        return pulumi.get(self, "external_id")

    @property
    @pulumi.getter(name="headquartersAddress")
    def headquarters_address(self) -> str:
        """
        Optional. The street address of the company's main headquarters, which may be different from the job location. The service attempts to geolocate the provided address, and populates a more specific location wherever possible in DerivedInfo.headquarters_location.
        """
        return pulumi.get(self, "headquarters_address")

    @property
    @pulumi.getter(name="hiringAgency")
    def hiring_agency(self) -> bool:
        """
        Optional. Set to true if it is the hiring agency that post jobs for other employers. Defaults to false if not provided.
        """
        return pulumi.get(self, "hiring_agency")

    @property
    @pulumi.getter(name="imageUri")
    def image_uri(self) -> str:
        """
        Optional. A URI that hosts the employer's company logo.
        """
        return pulumi.get(self, "image_uri")

    @property
    @pulumi.getter(name="keywordSearchableJobCustomAttributes")
    def keyword_searchable_job_custom_attributes(self) -> Sequence[str]:
        """
        Optional. A list of keys of filterable Job.custom_attributes, whose corresponding `string_values` are used in keyword search. Jobs with `string_values` under these specified field keys are returned if any of the values matches the search keyword. Custom field values with parenthesis, brackets and special symbols won't be properly searchable, and those keyword queries need to be surrounded by quotes.
        """
        return pulumi.get(self, "keyword_searchable_job_custom_attributes")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Required during company update. The resource name for a company. This is generated by the service when a company is created. The format is "projects/{project_id}/companies/{company_id}", for example, "projects/api-test-project/companies/foo".
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def size(self) -> str:
        """
        Optional. The employer's company size.
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter
    def suspended(self) -> bool:
        """
        Indicates whether a company is flagged to be suspended from public availability by the service when job content appears suspicious, abusive, or spammy.
        """
        return pulumi.get(self, "suspended")

    @property
    @pulumi.getter(name="websiteUri")
    def website_uri(self) -> str:
        """
        Optional. The URI representing the company's primary web site or home page, for example, "https://www.google.com". The maximum number of allowed characters is 255.
        """
        return pulumi.get(self, "website_uri")


class AwaitableGetCompanyResult(GetCompanyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCompanyResult(
            career_site_uri=self.career_site_uri,
            derived_info=self.derived_info,
            display_name=self.display_name,
            eeo_text=self.eeo_text,
            external_id=self.external_id,
            headquarters_address=self.headquarters_address,
            hiring_agency=self.hiring_agency,
            image_uri=self.image_uri,
            keyword_searchable_job_custom_attributes=self.keyword_searchable_job_custom_attributes,
            name=self.name,
            size=self.size,
            suspended=self.suspended,
            website_uri=self.website_uri)


def get_company(company_id: Optional[str] = None,
                project: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCompanyResult:
    """
    Retrieves specified company.
    """
    __args__ = dict()
    __args__['companyId'] = company_id
    __args__['project'] = project
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:jobs/v3:getCompany', __args__, opts=opts, typ=GetCompanyResult).value

    return AwaitableGetCompanyResult(
        career_site_uri=__ret__.career_site_uri,
        derived_info=__ret__.derived_info,
        display_name=__ret__.display_name,
        eeo_text=__ret__.eeo_text,
        external_id=__ret__.external_id,
        headquarters_address=__ret__.headquarters_address,
        hiring_agency=__ret__.hiring_agency,
        image_uri=__ret__.image_uri,
        keyword_searchable_job_custom_attributes=__ret__.keyword_searchable_job_custom_attributes,
        name=__ret__.name,
        size=__ret__.size,
        suspended=__ret__.suspended,
        website_uri=__ret__.website_uri)


@_utilities.lift_output_func(get_company)
def get_company_output(company_id: Optional[pulumi.Input[str]] = None,
                       project: Optional[pulumi.Input[Optional[str]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetCompanyResult]:
    """
    Retrieves specified company.
    """
    ...