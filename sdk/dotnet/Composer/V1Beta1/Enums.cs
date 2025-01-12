// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.GoogleNative.Composer.V1Beta1
{
    /// <summary>
    /// Optional. The size of the Cloud Composer environment. This field is supported for Cloud Composer environments in versions composer-2.*.*-airflow-*.*.* and newer.
    /// </summary>
    [EnumType]
    public readonly struct EnvironmentConfigEnvironmentSize : IEquatable<EnvironmentConfigEnvironmentSize>
    {
        private readonly string _value;

        private EnvironmentConfigEnvironmentSize(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// The size of the environment is unspecified.
        /// </summary>
        public static EnvironmentConfigEnvironmentSize EnvironmentSizeUnspecified { get; } = new EnvironmentConfigEnvironmentSize("ENVIRONMENT_SIZE_UNSPECIFIED");
        /// <summary>
        /// The environment size is small.
        /// </summary>
        public static EnvironmentConfigEnvironmentSize EnvironmentSizeSmall { get; } = new EnvironmentConfigEnvironmentSize("ENVIRONMENT_SIZE_SMALL");
        /// <summary>
        /// The environment size is medium.
        /// </summary>
        public static EnvironmentConfigEnvironmentSize EnvironmentSizeMedium { get; } = new EnvironmentConfigEnvironmentSize("ENVIRONMENT_SIZE_MEDIUM");
        /// <summary>
        /// The environment size is large.
        /// </summary>
        public static EnvironmentConfigEnvironmentSize EnvironmentSizeLarge { get; } = new EnvironmentConfigEnvironmentSize("ENVIRONMENT_SIZE_LARGE");

        public static bool operator ==(EnvironmentConfigEnvironmentSize left, EnvironmentConfigEnvironmentSize right) => left.Equals(right);
        public static bool operator !=(EnvironmentConfigEnvironmentSize left, EnvironmentConfigEnvironmentSize right) => !left.Equals(right);

        public static explicit operator string(EnvironmentConfigEnvironmentSize value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is EnvironmentConfigEnvironmentSize other && Equals(other);
        public bool Equals(EnvironmentConfigEnvironmentSize other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The current state of the environment.
    /// </summary>
    [EnumType]
    public readonly struct EnvironmentState : IEquatable<EnvironmentState>
    {
        private readonly string _value;

        private EnvironmentState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// The state of the environment is unknown.
        /// </summary>
        public static EnvironmentState StateUnspecified { get; } = new EnvironmentState("STATE_UNSPECIFIED");
        /// <summary>
        /// The environment is in the process of being created.
        /// </summary>
        public static EnvironmentState Creating { get; } = new EnvironmentState("CREATING");
        /// <summary>
        /// The environment is currently running and healthy. It is ready for use.
        /// </summary>
        public static EnvironmentState Running { get; } = new EnvironmentState("RUNNING");
        /// <summary>
        /// The environment is being updated. It remains usable but cannot receive additional update requests or be deleted at this time.
        /// </summary>
        public static EnvironmentState Updating { get; } = new EnvironmentState("UPDATING");
        /// <summary>
        /// The environment is undergoing deletion. It cannot be used.
        /// </summary>
        public static EnvironmentState Deleting { get; } = new EnvironmentState("DELETING");
        /// <summary>
        /// The environment has encountered an error and cannot be used.
        /// </summary>
        public static EnvironmentState Error { get; } = new EnvironmentState("ERROR");

        public static bool operator ==(EnvironmentState left, EnvironmentState right) => left.Equals(right);
        public static bool operator !=(EnvironmentState left, EnvironmentState right) => !left.Equals(right);

        public static explicit operator string(EnvironmentState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is EnvironmentState other && Equals(other);
        public bool Equals(EnvironmentState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
