// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.GoogleNative.DataCatalog.V1
{
    /// <summary>
    /// The type of the entry. Only used for entries with types listed in the `EntryType` enum. Currently, only `FILESET` enum value is allowed. All other entries created in Data Catalog must use the `user_specified_type`.
    /// </summary>
    [EnumType]
    public readonly struct EntryType : IEquatable<EntryType>
    {
        private readonly string _value;

        private EntryType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Default unknown type.
        /// </summary>
        public static EntryType EntryTypeUnspecified { get; } = new EntryType("ENTRY_TYPE_UNSPECIFIED");
        /// <summary>
        /// Output only. The entry type that has a GoogleSQL schema, including logical views.
        /// </summary>
        public static EntryType Table { get; } = new EntryType("TABLE");
        /// <summary>
        /// Output only. The type of models. For more information, see [Supported models in BigQuery ML] (https://cloud.google.com/bigquery-ml/docs/introduction#supported_models_in).
        /// </summary>
        public static EntryType Model { get; } = new EntryType("MODEL");
        /// <summary>
        /// An entry type for streaming entries. For example, a Pub/Sub topic.
        /// </summary>
        public static EntryType DataStream { get; } = new EntryType("DATA_STREAM");
        /// <summary>
        /// An entry type for a set of files or objects. For example, a Cloud Storage fileset.
        /// </summary>
        public static EntryType Fileset { get; } = new EntryType("FILESET");
        /// <summary>
        /// A group of servers that work together. For example, a Kafka cluster.
        /// </summary>
        public static EntryType Cluster { get; } = new EntryType("CLUSTER");
        /// <summary>
        /// A database.
        /// </summary>
        public static EntryType Database { get; } = new EntryType("DATABASE");
        /// <summary>
        /// Output only. Connection to a data source. For example, a BigQuery connection.
        /// </summary>
        public static EntryType DataSourceConnection { get; } = new EntryType("DATA_SOURCE_CONNECTION");
        /// <summary>
        /// Output only. Routine, for example, a BigQuery routine.
        /// </summary>
        public static EntryType Routine { get; } = new EntryType("ROUTINE");
        /// <summary>
        /// A service, for example, a Dataproc Metastore service.
        /// </summary>
        public static EntryType Service { get; } = new EntryType("SERVICE");

        public static bool operator ==(EntryType left, EntryType right) => left.Equals(right);
        public static bool operator !=(EntryType left, EntryType right) => !left.Equals(right);

        public static explicit operator string(EntryType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is EntryType other && Equals(other);
        public bool Equals(EntryType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The type of the BigQuery connection.
    /// </summary>
    [EnumType]
    public readonly struct GoogleCloudDatacatalogV1BigQueryConnectionSpecConnectionType : IEquatable<GoogleCloudDatacatalogV1BigQueryConnectionSpecConnectionType>
    {
        private readonly string _value;

        private GoogleCloudDatacatalogV1BigQueryConnectionSpecConnectionType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Unspecified type.
        /// </summary>
        public static GoogleCloudDatacatalogV1BigQueryConnectionSpecConnectionType ConnectionTypeUnspecified { get; } = new GoogleCloudDatacatalogV1BigQueryConnectionSpecConnectionType("CONNECTION_TYPE_UNSPECIFIED");
        /// <summary>
        /// Cloud SQL connection.
        /// </summary>
        public static GoogleCloudDatacatalogV1BigQueryConnectionSpecConnectionType CloudSql { get; } = new GoogleCloudDatacatalogV1BigQueryConnectionSpecConnectionType("CLOUD_SQL");

        public static bool operator ==(GoogleCloudDatacatalogV1BigQueryConnectionSpecConnectionType left, GoogleCloudDatacatalogV1BigQueryConnectionSpecConnectionType right) => left.Equals(right);
        public static bool operator !=(GoogleCloudDatacatalogV1BigQueryConnectionSpecConnectionType left, GoogleCloudDatacatalogV1BigQueryConnectionSpecConnectionType right) => !left.Equals(right);

        public static explicit operator string(GoogleCloudDatacatalogV1BigQueryConnectionSpecConnectionType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is GoogleCloudDatacatalogV1BigQueryConnectionSpecConnectionType other && Equals(other);
        public bool Equals(GoogleCloudDatacatalogV1BigQueryConnectionSpecConnectionType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Type of the Cloud SQL database.
    /// </summary>
    [EnumType]
    public readonly struct GoogleCloudDatacatalogV1CloudSqlBigQueryConnectionSpecType : IEquatable<GoogleCloudDatacatalogV1CloudSqlBigQueryConnectionSpecType>
    {
        private readonly string _value;

        private GoogleCloudDatacatalogV1CloudSqlBigQueryConnectionSpecType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Unspecified database type.
        /// </summary>
        public static GoogleCloudDatacatalogV1CloudSqlBigQueryConnectionSpecType DatabaseTypeUnspecified { get; } = new GoogleCloudDatacatalogV1CloudSqlBigQueryConnectionSpecType("DATABASE_TYPE_UNSPECIFIED");
        /// <summary>
        /// Cloud SQL for PostgreSQL.
        /// </summary>
        public static GoogleCloudDatacatalogV1CloudSqlBigQueryConnectionSpecType Postgres { get; } = new GoogleCloudDatacatalogV1CloudSqlBigQueryConnectionSpecType("POSTGRES");
        /// <summary>
        /// Cloud SQL for MySQL.
        /// </summary>
        public static GoogleCloudDatacatalogV1CloudSqlBigQueryConnectionSpecType Mysql { get; } = new GoogleCloudDatacatalogV1CloudSqlBigQueryConnectionSpecType("MYSQL");

        public static bool operator ==(GoogleCloudDatacatalogV1CloudSqlBigQueryConnectionSpecType left, GoogleCloudDatacatalogV1CloudSqlBigQueryConnectionSpecType right) => left.Equals(right);
        public static bool operator !=(GoogleCloudDatacatalogV1CloudSqlBigQueryConnectionSpecType left, GoogleCloudDatacatalogV1CloudSqlBigQueryConnectionSpecType right) => !left.Equals(right);

        public static explicit operator string(GoogleCloudDatacatalogV1CloudSqlBigQueryConnectionSpecType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is GoogleCloudDatacatalogV1CloudSqlBigQueryConnectionSpecType other && Equals(other);
        public bool Equals(GoogleCloudDatacatalogV1CloudSqlBigQueryConnectionSpecType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Type of this table.
    /// </summary>
    [EnumType]
    public readonly struct GoogleCloudDatacatalogV1DatabaseTableSpecType : IEquatable<GoogleCloudDatacatalogV1DatabaseTableSpecType>
    {
        private readonly string _value;

        private GoogleCloudDatacatalogV1DatabaseTableSpecType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Default unknown table type.
        /// </summary>
        public static GoogleCloudDatacatalogV1DatabaseTableSpecType TableTypeUnspecified { get; } = new GoogleCloudDatacatalogV1DatabaseTableSpecType("TABLE_TYPE_UNSPECIFIED");
        /// <summary>
        /// Native table.
        /// </summary>
        public static GoogleCloudDatacatalogV1DatabaseTableSpecType Native { get; } = new GoogleCloudDatacatalogV1DatabaseTableSpecType("NATIVE");
        /// <summary>
        /// External table.
        /// </summary>
        public static GoogleCloudDatacatalogV1DatabaseTableSpecType External { get; } = new GoogleCloudDatacatalogV1DatabaseTableSpecType("EXTERNAL");

        public static bool operator ==(GoogleCloudDatacatalogV1DatabaseTableSpecType left, GoogleCloudDatacatalogV1DatabaseTableSpecType right) => left.Equals(right);
        public static bool operator !=(GoogleCloudDatacatalogV1DatabaseTableSpecType left, GoogleCloudDatacatalogV1DatabaseTableSpecType right) => !left.Equals(right);

        public static explicit operator string(GoogleCloudDatacatalogV1DatabaseTableSpecType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is GoogleCloudDatacatalogV1DatabaseTableSpecType other && Equals(other);
        public bool Equals(GoogleCloudDatacatalogV1DatabaseTableSpecType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Specifies whether the argument is input or output.
    /// </summary>
    [EnumType]
    public readonly struct GoogleCloudDatacatalogV1RoutineSpecArgumentMode : IEquatable<GoogleCloudDatacatalogV1RoutineSpecArgumentMode>
    {
        private readonly string _value;

        private GoogleCloudDatacatalogV1RoutineSpecArgumentMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Unspecified mode.
        /// </summary>
        public static GoogleCloudDatacatalogV1RoutineSpecArgumentMode ModeUnspecified { get; } = new GoogleCloudDatacatalogV1RoutineSpecArgumentMode("MODE_UNSPECIFIED");
        /// <summary>
        /// The argument is input-only.
        /// </summary>
        public static GoogleCloudDatacatalogV1RoutineSpecArgumentMode In { get; } = new GoogleCloudDatacatalogV1RoutineSpecArgumentMode("IN");
        /// <summary>
        /// The argument is output-only.
        /// </summary>
        public static GoogleCloudDatacatalogV1RoutineSpecArgumentMode Out { get; } = new GoogleCloudDatacatalogV1RoutineSpecArgumentMode("OUT");
        /// <summary>
        /// The argument is both an input and an output.
        /// </summary>
        public static GoogleCloudDatacatalogV1RoutineSpecArgumentMode Inout { get; } = new GoogleCloudDatacatalogV1RoutineSpecArgumentMode("INOUT");

        public static bool operator ==(GoogleCloudDatacatalogV1RoutineSpecArgumentMode left, GoogleCloudDatacatalogV1RoutineSpecArgumentMode right) => left.Equals(right);
        public static bool operator !=(GoogleCloudDatacatalogV1RoutineSpecArgumentMode left, GoogleCloudDatacatalogV1RoutineSpecArgumentMode right) => !left.Equals(right);

        public static explicit operator string(GoogleCloudDatacatalogV1RoutineSpecArgumentMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is GoogleCloudDatacatalogV1RoutineSpecArgumentMode other && Equals(other);
        public bool Equals(GoogleCloudDatacatalogV1RoutineSpecArgumentMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The type of the routine.
    /// </summary>
    [EnumType]
    public readonly struct GoogleCloudDatacatalogV1RoutineSpecRoutineType : IEquatable<GoogleCloudDatacatalogV1RoutineSpecRoutineType>
    {
        private readonly string _value;

        private GoogleCloudDatacatalogV1RoutineSpecRoutineType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Unspecified type.
        /// </summary>
        public static GoogleCloudDatacatalogV1RoutineSpecRoutineType RoutineTypeUnspecified { get; } = new GoogleCloudDatacatalogV1RoutineSpecRoutineType("ROUTINE_TYPE_UNSPECIFIED");
        /// <summary>
        /// Non-builtin permanent scalar function.
        /// </summary>
        public static GoogleCloudDatacatalogV1RoutineSpecRoutineType ScalarFunction { get; } = new GoogleCloudDatacatalogV1RoutineSpecRoutineType("SCALAR_FUNCTION");
        /// <summary>
        /// Stored procedure.
        /// </summary>
        public static GoogleCloudDatacatalogV1RoutineSpecRoutineType Procedure { get; } = new GoogleCloudDatacatalogV1RoutineSpecRoutineType("PROCEDURE");

        public static bool operator ==(GoogleCloudDatacatalogV1RoutineSpecRoutineType left, GoogleCloudDatacatalogV1RoutineSpecRoutineType right) => left.Equals(right);
        public static bool operator !=(GoogleCloudDatacatalogV1RoutineSpecRoutineType left, GoogleCloudDatacatalogV1RoutineSpecRoutineType right) => !left.Equals(right);

        public static explicit operator string(GoogleCloudDatacatalogV1RoutineSpecRoutineType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is GoogleCloudDatacatalogV1RoutineSpecRoutineType other && Equals(other);
        public bool Equals(GoogleCloudDatacatalogV1RoutineSpecRoutineType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct TaxonomyActivatedPolicyTypesItem : IEquatable<TaxonomyActivatedPolicyTypesItem>
    {
        private readonly string _value;

        private TaxonomyActivatedPolicyTypesItem(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Unspecified policy type.
        /// </summary>
        public static TaxonomyActivatedPolicyTypesItem PolicyTypeUnspecified { get; } = new TaxonomyActivatedPolicyTypesItem("POLICY_TYPE_UNSPECIFIED");
        /// <summary>
        /// Fine-grained access control policy that enables access control on tagged sub-resources.
        /// </summary>
        public static TaxonomyActivatedPolicyTypesItem FineGrainedAccessControl { get; } = new TaxonomyActivatedPolicyTypesItem("FINE_GRAINED_ACCESS_CONTROL");

        public static bool operator ==(TaxonomyActivatedPolicyTypesItem left, TaxonomyActivatedPolicyTypesItem right) => left.Equals(right);
        public static bool operator !=(TaxonomyActivatedPolicyTypesItem left, TaxonomyActivatedPolicyTypesItem right) => !left.Equals(right);

        public static explicit operator string(TaxonomyActivatedPolicyTypesItem value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is TaxonomyActivatedPolicyTypesItem other && Equals(other);
        public bool Equals(TaxonomyActivatedPolicyTypesItem other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}