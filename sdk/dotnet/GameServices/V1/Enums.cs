// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.GoogleNative.GameServices.V1
{
    /// <summary>
    /// The log type that this config enables.
    /// </summary>
    [EnumType]
    public readonly struct AuditLogConfigLogType : IEquatable<AuditLogConfigLogType>
    {
        private readonly string _value;

        private AuditLogConfigLogType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Default case. Should never be this.
        /// </summary>
        public static AuditLogConfigLogType LogTypeUnspecified { get; } = new AuditLogConfigLogType("LOG_TYPE_UNSPECIFIED");
        /// <summary>
        /// Admin reads. Example: CloudIAM getIamPolicy
        /// </summary>
        public static AuditLogConfigLogType AdminRead { get; } = new AuditLogConfigLogType("ADMIN_READ");
        /// <summary>
        /// Data writes. Example: CloudSQL Users create
        /// </summary>
        public static AuditLogConfigLogType DataWrite { get; } = new AuditLogConfigLogType("DATA_WRITE");
        /// <summary>
        /// Data reads. Example: CloudSQL Users list
        /// </summary>
        public static AuditLogConfigLogType DataRead { get; } = new AuditLogConfigLogType("DATA_READ");

        public static bool operator ==(AuditLogConfigLogType left, AuditLogConfigLogType right) => left.Equals(right);
        public static bool operator !=(AuditLogConfigLogType left, AuditLogConfigLogType right) => !left.Equals(right);

        public static explicit operator string(AuditLogConfigLogType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AuditLogConfigLogType other && Equals(other);
        public bool Equals(AuditLogConfigLogType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The type of the permission that was checked.
    /// </summary>
    [EnumType]
    public readonly struct AuthorizationLoggingOptionsPermissionType : IEquatable<AuthorizationLoggingOptionsPermissionType>
    {
        private readonly string _value;

        private AuthorizationLoggingOptionsPermissionType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Default. Should not be used.
        /// </summary>
        public static AuthorizationLoggingOptionsPermissionType PermissionTypeUnspecified { get; } = new AuthorizationLoggingOptionsPermissionType("PERMISSION_TYPE_UNSPECIFIED");
        /// <summary>
        /// A read of admin (meta) data.
        /// </summary>
        public static AuthorizationLoggingOptionsPermissionType AdminRead { get; } = new AuthorizationLoggingOptionsPermissionType("ADMIN_READ");
        /// <summary>
        /// A write of admin (meta) data.
        /// </summary>
        public static AuthorizationLoggingOptionsPermissionType AdminWrite { get; } = new AuthorizationLoggingOptionsPermissionType("ADMIN_WRITE");
        /// <summary>
        /// A read of standard data.
        /// </summary>
        public static AuthorizationLoggingOptionsPermissionType DataRead { get; } = new AuthorizationLoggingOptionsPermissionType("DATA_READ");
        /// <summary>
        /// A write of standard data.
        /// </summary>
        public static AuthorizationLoggingOptionsPermissionType DataWrite { get; } = new AuthorizationLoggingOptionsPermissionType("DATA_WRITE");

        public static bool operator ==(AuthorizationLoggingOptionsPermissionType left, AuthorizationLoggingOptionsPermissionType right) => left.Equals(right);
        public static bool operator !=(AuthorizationLoggingOptionsPermissionType left, AuthorizationLoggingOptionsPermissionType right) => !left.Equals(right);

        public static explicit operator string(AuthorizationLoggingOptionsPermissionType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AuthorizationLoggingOptionsPermissionType other && Equals(other);
        public bool Equals(AuthorizationLoggingOptionsPermissionType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The log_name to populate in the Cloud Audit Record.
    /// </summary>
    [EnumType]
    public readonly struct CloudAuditOptionsLogName : IEquatable<CloudAuditOptionsLogName>
    {
        private readonly string _value;

        private CloudAuditOptionsLogName(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Default. Should not be used.
        /// </summary>
        public static CloudAuditOptionsLogName UnspecifiedLogName { get; } = new CloudAuditOptionsLogName("UNSPECIFIED_LOG_NAME");
        /// <summary>
        /// Corresponds to "cloudaudit.googleapis.com/activity"
        /// </summary>
        public static CloudAuditOptionsLogName AdminActivity { get; } = new CloudAuditOptionsLogName("ADMIN_ACTIVITY");
        /// <summary>
        /// Corresponds to "cloudaudit.googleapis.com/data_access"
        /// </summary>
        public static CloudAuditOptionsLogName DataAccess { get; } = new CloudAuditOptionsLogName("DATA_ACCESS");

        public static bool operator ==(CloudAuditOptionsLogName left, CloudAuditOptionsLogName right) => left.Equals(right);
        public static bool operator !=(CloudAuditOptionsLogName left, CloudAuditOptionsLogName right) => !left.Equals(right);

        public static explicit operator string(CloudAuditOptionsLogName value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is CloudAuditOptionsLogName other && Equals(other);
        public bool Equals(CloudAuditOptionsLogName other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Trusted attributes supplied by the IAM system.
    /// </summary>
    [EnumType]
    public readonly struct ConditionIam : IEquatable<ConditionIam>
    {
        private readonly string _value;

        private ConditionIam(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Default non-attribute.
        /// </summary>
        public static ConditionIam NoAttr { get; } = new ConditionIam("NO_ATTR");
        /// <summary>
        /// Either principal or (if present) authority selector.
        /// </summary>
        public static ConditionIam Authority { get; } = new ConditionIam("AUTHORITY");
        /// <summary>
        /// The principal (even if an authority selector is present), which must only be used for attribution, not authorization.
        /// </summary>
        public static ConditionIam Attribution { get; } = new ConditionIam("ATTRIBUTION");
        /// <summary>
        /// Any of the security realms in the IAMContext (go/security-realms). When used with IN, the condition indicates "any of the request's realms match one of the given values; with NOT_IN, "none of the realms match any of the given values". Note that a value can be: - 'self' (i.e., allow connections from clients that are in the same security realm, which is currently but not guaranteed to be campus-sized) - 'self:metro' (i.e., clients that are in the same metro) - 'self:cloud-region' (i.e., allow connections from clients that are in the same cloud region) - 'self:prod-region' (i.e., allow connections from clients that are in the same prod region) - 'guardians' (i.e., allow connections from its guardian realms. See go/security-realms-glossary#guardian for more information.) - a realm (e.g., 'campus-abc') - a realm group (e.g., 'realms-for-borg-cell-xx', see: go/realm-groups) A match is determined by a realm group membership check performed by a RealmAclRep object (go/realm-acl-howto). It is not permitted to grant access based on the *absence* of a realm, so realm conditions can only be used in a "positive" context (e.g., ALLOW/IN or DENY/NOT_IN).
        /// </summary>
        public static ConditionIam SecurityRealm { get; } = new ConditionIam("SECURITY_REALM");
        /// <summary>
        /// An approver (distinct from the requester) that has authorized this request. When used with IN, the condition indicates that one of the approvers associated with the request matches the specified principal, or is a member of the specified group. Approvers can only grant additional access, and are thus only used in a strictly positive context (e.g. ALLOW/IN or DENY/NOT_IN).
        /// </summary>
        public static ConditionIam Approver { get; } = new ConditionIam("APPROVER");
        /// <summary>
        /// What types of justifications have been supplied with this request. String values should match enum names from security.credentials.JustificationType, e.g. "MANUAL_STRING". It is not permitted to grant access based on the *absence* of a justification, so justification conditions can only be used in a "positive" context (e.g., ALLOW/IN or DENY/NOT_IN). Multiple justifications, e.g., a Buganizer ID and a manually-entered reason, are normal and supported.
        /// </summary>
        public static ConditionIam JustificationType { get; } = new ConditionIam("JUSTIFICATION_TYPE");
        /// <summary>
        /// What type of credentials have been supplied with this request. String values should match enum names from security_loas_l2.CredentialsType - currently, only CREDS_TYPE_EMERGENCY is supported. It is not permitted to grant access based on the *absence* of a credentials type, so the conditions can only be used in a "positive" context (e.g., ALLOW/IN or DENY/NOT_IN).
        /// </summary>
        public static ConditionIam CredentialsType { get; } = new ConditionIam("CREDENTIALS_TYPE");
        /// <summary>
        /// EXPERIMENTAL -- DO NOT USE. The conditions can only be used in a "positive" context (e.g., ALLOW/IN or DENY/NOT_IN).
        /// </summary>
        public static ConditionIam CredsAssertion { get; } = new ConditionIam("CREDS_ASSERTION");

        public static bool operator ==(ConditionIam left, ConditionIam right) => left.Equals(right);
        public static bool operator !=(ConditionIam left, ConditionIam right) => !left.Equals(right);

        public static explicit operator string(ConditionIam value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ConditionIam other && Equals(other);
        public bool Equals(ConditionIam other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// An operator to apply the subject with.
    /// </summary>
    [EnumType]
    public readonly struct ConditionOp : IEquatable<ConditionOp>
    {
        private readonly string _value;

        private ConditionOp(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Default no-op.
        /// </summary>
        public static ConditionOp NoOp { get; } = new ConditionOp("NO_OP");
        /// <summary>
        /// DEPRECATED. Use IN instead.
        /// </summary>
        public static ConditionOp EqualsValue { get; } = new ConditionOp("EQUALS");
        /// <summary>
        /// DEPRECATED. Use NOT_IN instead.
        /// </summary>
        public static ConditionOp NotEquals { get; } = new ConditionOp("NOT_EQUALS");
        /// <summary>
        /// The condition is true if the subject (or any element of it if it is a set) matches any of the supplied values.
        /// </summary>
        public static ConditionOp In { get; } = new ConditionOp("IN");
        /// <summary>
        /// The condition is true if the subject (or every element of it if it is a set) matches none of the supplied values.
        /// </summary>
        public static ConditionOp NotIn { get; } = new ConditionOp("NOT_IN");
        /// <summary>
        /// Subject is discharged
        /// </summary>
        public static ConditionOp Discharged { get; } = new ConditionOp("DISCHARGED");

        public static bool operator ==(ConditionOp left, ConditionOp right) => left.Equals(right);
        public static bool operator !=(ConditionOp left, ConditionOp right) => !left.Equals(right);

        public static explicit operator string(ConditionOp value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ConditionOp other && Equals(other);
        public bool Equals(ConditionOp other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Trusted attributes supplied by any service that owns resources and uses the IAM system for access control.
    /// </summary>
    [EnumType]
    public readonly struct ConditionSys : IEquatable<ConditionSys>
    {
        private readonly string _value;

        private ConditionSys(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Default non-attribute type
        /// </summary>
        public static ConditionSys NoAttr { get; } = new ConditionSys("NO_ATTR");
        /// <summary>
        /// Region of the resource
        /// </summary>
        public static ConditionSys Region { get; } = new ConditionSys("REGION");
        /// <summary>
        /// Service name
        /// </summary>
        public static ConditionSys Service { get; } = new ConditionSys("SERVICE");
        /// <summary>
        /// Resource name
        /// </summary>
        public static ConditionSys Name { get; } = new ConditionSys("NAME");
        /// <summary>
        /// IP address of the caller
        /// </summary>
        public static ConditionSys Ip { get; } = new ConditionSys("IP");

        public static bool operator ==(ConditionSys left, ConditionSys right) => left.Equals(right);
        public static bool operator !=(ConditionSys left, ConditionSys right) => !left.Equals(right);

        public static explicit operator string(ConditionSys value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ConditionSys other && Equals(other);
        public bool Equals(ConditionSys other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct DataAccessOptionsLogMode : IEquatable<DataAccessOptionsLogMode>
    {
        private readonly string _value;

        private DataAccessOptionsLogMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Client is not required to write a partial Gin log immediately after the authorization check. If client chooses to write one and it fails, client may either fail open (allow the operation to continue) or fail closed (handle as a DENY outcome).
        /// </summary>
        public static DataAccessOptionsLogMode LogModeUnspecified { get; } = new DataAccessOptionsLogMode("LOG_MODE_UNSPECIFIED");
        /// <summary>
        /// The application's operation in the context of which this authorization check is being made may only be performed if it is successfully logged to Gin. For instance, the authorization library may satisfy this obligation by emitting a partial log entry at authorization check time and only returning ALLOW to the application if it succeeds. If a matching Rule has this directive, but the client has not indicated that it will honor such requirements, then the IAM check will result in authorization failure by setting CheckPolicyResponse.success=false.
        /// </summary>
        public static DataAccessOptionsLogMode LogFailClosed { get; } = new DataAccessOptionsLogMode("LOG_FAIL_CLOSED");

        public static bool operator ==(DataAccessOptionsLogMode left, DataAccessOptionsLogMode right) => left.Equals(right);
        public static bool operator !=(DataAccessOptionsLogMode left, DataAccessOptionsLogMode right) => !left.Equals(right);

        public static explicit operator string(DataAccessOptionsLogMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DataAccessOptionsLogMode other && Equals(other);
        public bool Equals(DataAccessOptionsLogMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Required
    /// </summary>
    [EnumType]
    public readonly struct RuleAction : IEquatable<RuleAction>
    {
        private readonly string _value;

        private RuleAction(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Default no action.
        /// </summary>
        public static RuleAction NoAction { get; } = new RuleAction("NO_ACTION");
        /// <summary>
        /// Matching 'Entries' grant access.
        /// </summary>
        public static RuleAction Allow { get; } = new RuleAction("ALLOW");
        /// <summary>
        /// Matching 'Entries' grant access and the caller promises to log the request per the returned log_configs.
        /// </summary>
        public static RuleAction AllowWithLog { get; } = new RuleAction("ALLOW_WITH_LOG");
        /// <summary>
        /// Matching 'Entries' deny access.
        /// </summary>
        public static RuleAction Deny { get; } = new RuleAction("DENY");
        /// <summary>
        /// Matching 'Entries' deny access and the caller promises to log the request per the returned log_configs.
        /// </summary>
        public static RuleAction DenyWithLog { get; } = new RuleAction("DENY_WITH_LOG");
        /// <summary>
        /// Matching 'Entries' tell IAM.Check callers to generate logs.
        /// </summary>
        public static RuleAction Log { get; } = new RuleAction("LOG");

        public static bool operator ==(RuleAction left, RuleAction right) => left.Equals(right);
        public static bool operator !=(RuleAction left, RuleAction right) => !left.Equals(right);

        public static explicit operator string(RuleAction value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is RuleAction other && Equals(other);
        public bool Equals(RuleAction other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
