// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const AuditLogConfigLogType = {
    /**
     * Default case. Should never be this.
     */
    LogTypeUnspecified: "LOG_TYPE_UNSPECIFIED",
    /**
     * Admin reads. Example: CloudIAM getIamPolicy
     */
    AdminRead: "ADMIN_READ",
    /**
     * Data writes. Example: CloudSQL Users create
     */
    DataWrite: "DATA_WRITE",
    /**
     * Data reads. Example: CloudSQL Users list
     */
    DataRead: "DATA_READ",
} as const;

/**
 * The log type that this config enables.
 */
export type AuditLogConfigLogType = (typeof AuditLogConfigLogType)[keyof typeof AuditLogConfigLogType];

export const FunctionIngressSettings = {
    /**
     * Unspecified.
     */
    IngressSettingsUnspecified: "INGRESS_SETTINGS_UNSPECIFIED",
    /**
     * Allow HTTP traffic from public and private sources.
     */
    AllowAll: "ALLOW_ALL",
    /**
     * Allow HTTP traffic from only private VPC sources.
     */
    AllowInternalOnly: "ALLOW_INTERNAL_ONLY",
    /**
     * Allow HTTP traffic from private VPC sources and through GCLB.
     */
    AllowInternalAndGclb: "ALLOW_INTERNAL_AND_GCLB",
} as const;

/**
 * The ingress settings for the function, controlling what traffic can reach it.
 */
export type FunctionIngressSettings = (typeof FunctionIngressSettings)[keyof typeof FunctionIngressSettings];

export const FunctionVpcConnectorEgressSettings = {
    /**
     * Unspecified.
     */
    VpcConnectorEgressSettingsUnspecified: "VPC_CONNECTOR_EGRESS_SETTINGS_UNSPECIFIED",
    /**
     * Use the VPC Access Connector only for private IP space from RFC1918.
     */
    PrivateRangesOnly: "PRIVATE_RANGES_ONLY",
    /**
     * Force the use of VPC Access Connector for all egress traffic from the function.
     */
    AllTraffic: "ALL_TRAFFIC",
} as const;

/**
 * The egress settings for the connector, controlling what traffic is diverted through it.
 */
export type FunctionVpcConnectorEgressSettings = (typeof FunctionVpcConnectorEgressSettings)[keyof typeof FunctionVpcConnectorEgressSettings];

export const HttpsTriggerSecurityLevel = {
    /**
     * Unspecified.
     */
    SecurityLevelUnspecified: "SECURITY_LEVEL_UNSPECIFIED",
    /**
     * Requests for a URL that match this handler that do not use HTTPS are automatically redirected to the HTTPS URL with the same path. Query parameters are reserved for the redirect.
     */
    SecureAlways: "SECURE_ALWAYS",
    /**
     * Both HTTP and HTTPS requests with URLs that match the handler succeed without redirects. The application can examine the request to determine which protocol was used and respond accordingly.
     */
    SecureOptional: "SECURE_OPTIONAL",
} as const;

/**
 * The security level for the function.
 */
export type HttpsTriggerSecurityLevel = (typeof HttpsTriggerSecurityLevel)[keyof typeof HttpsTriggerSecurityLevel];