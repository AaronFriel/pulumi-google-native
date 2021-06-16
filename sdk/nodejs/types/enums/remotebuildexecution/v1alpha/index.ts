// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const GoogleDevtoolsRemotebuildexecutionAdminV1alphaFeaturePolicyFeaturePolicy = {
    /**
     * Default value, if not explicitly set. Equivalent to FORBIDDEN, unless otherwise documented on a specific Feature.
     */
    PolicyUnspecified: "POLICY_UNSPECIFIED",
    /**
     * Feature is explicitly allowed.
     */
    Allowed: "ALLOWED",
    /**
     * Feature is forbidden. Requests attempting to leverage it will get an FailedPrecondition error, with a message like: "Feature forbidden by FeaturePolicy: Feature on instance "
     */
    Forbidden: "FORBIDDEN",
    /**
     * Only the values specified in the `allowed_values` are allowed.
     */
    Restricted: "RESTRICTED",
} as const;

/**
 * The policy of the feature.
 */
export type GoogleDevtoolsRemotebuildexecutionAdminV1alphaFeaturePolicyFeaturePolicy = (typeof GoogleDevtoolsRemotebuildexecutionAdminV1alphaFeaturePolicyFeaturePolicy)[keyof typeof GoogleDevtoolsRemotebuildexecutionAdminV1alphaFeaturePolicyFeaturePolicy];

export const GoogleDevtoolsRemotebuildexecutionAdminV1alphaFeaturePolicyLinuxIsolation = {
    /**
     * Default value. Will be using Linux default runtime.
     */
    LinuxIsolationUnspecified: "LINUX_ISOLATION_UNSPECIFIED",
    /**
     * Use gVisor runsc runtime.
     */
    Gvisor: "GVISOR",
    /**
     * Use stardard Linux runtime. This has the same behaviour as unspecified, but it can be used to revert back from gVisor.
     */
    Off: "OFF",
} as const;

/**
 * linux_isolation allows overriding the docker runtime used for containers started on Linux.
 */
export type GoogleDevtoolsRemotebuildexecutionAdminV1alphaFeaturePolicyLinuxIsolation = (typeof GoogleDevtoolsRemotebuildexecutionAdminV1alphaFeaturePolicyLinuxIsolation)[keyof typeof GoogleDevtoolsRemotebuildexecutionAdminV1alphaFeaturePolicyLinuxIsolation];

export const InstanceState = {
    /**
     * Not a valid state, but the default value of the enum.
     */
    StateUnspecified: "STATE_UNSPECIFIED",
    /**
     * The instance is in state `CREATING` once `CreateInstance` is called and before the instance is ready for use.
     */
    Creating: "CREATING",
    /**
     * The instance is in state `RUNNING` when it is ready for use.
     */
    Running: "RUNNING",
    /**
     * An `INACTIVE` instance indicates that there is a problem that needs to be fixed. Such instances cannot be used for execution and instances that remain in this state for a significant period of time will be removed permanently.
     */
    Inactive: "INACTIVE",
} as const;

/**
 * Output only. State of the instance.
 */
export type InstanceState = (typeof InstanceState)[keyof typeof InstanceState];

export const WorkerPoolState = {
    /**
     * Not a valid state, but the default value of the enum.
     */
    StateUnspecified: "STATE_UNSPECIFIED",
    /**
     * The worker pool is in state `CREATING` once `CreateWorkerPool` is called and before all requested workers are ready.
     */
    Creating: "CREATING",
    /**
     * The worker pool is in state `RUNNING` when all its workers are ready for use.
     */
    Running: "RUNNING",
    /**
     * The worker pool is in state `UPDATING` once `UpdateWorkerPool` is called and before the new configuration has all the requested workers ready for use, and no older configuration has any workers. At that point the state transitions to `RUNNING`.
     */
    Updating: "UPDATING",
    /**
     * The worker pool is in state `DELETING` once the `Delete` method is called and before the deletion completes.
     */
    Deleting: "DELETING",
    /**
     * The worker pool is in state `INACTIVE` when the instance hosting the worker pool in not running.
     */
    Inactive: "INACTIVE",
} as const;

/**
 * Output only. State of the worker pool.
 */
export type WorkerPoolState = (typeof WorkerPoolState)[keyof typeof WorkerPoolState];