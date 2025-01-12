# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'CompositeFilterLogicOperator',
    'SortOptionsSortOrder',
    'SourcePredefinedSource',
    'SourceScoringConfigSourceImportance',
]


class CompositeFilterLogicOperator(str, Enum):
    """
    The logic operator of the sub filter.
    """
    AND_ = "AND"
    """
    Logical operators, which can only be applied to sub filters.
    """
    OR_ = "OR"
    NOT_ = "NOT"
    """
    NOT can only be applied on a single sub filter.
    """


class SortOptionsSortOrder(str, Enum):
    """
    Ascending is the default sort order
    """
    ASCENDING = "ASCENDING"
    DESCENDING = "DESCENDING"


class SourcePredefinedSource(str, Enum):
    """
    Predefined content source for Google Apps.
    """
    NONE = "NONE"
    QUERY_HISTORY = "QUERY_HISTORY"
    """
    Suggests queries issued by the user in the past. Only valid when used with the suggest API. Ignored when used in the query API.
    """
    PERSON = "PERSON"
    """
    Suggests people in the organization. Only valid when used with the suggest API. Results in an error when used in the query API.
    """
    GOOGLE_DRIVE = "GOOGLE_DRIVE"
    GOOGLE_GMAIL = "GOOGLE_GMAIL"
    GOOGLE_SITES = "GOOGLE_SITES"
    GOOGLE_GROUPS = "GOOGLE_GROUPS"
    GOOGLE_CALENDAR = "GOOGLE_CALENDAR"
    GOOGLE_KEEP = "GOOGLE_KEEP"


class SourceScoringConfigSourceImportance(str, Enum):
    """
    Importance of the source.
    """
    DEFAULT = "DEFAULT"
    LOW = "LOW"
    HIGH = "HIGH"
