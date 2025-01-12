// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const LoggingConfigLogActionStatesItem = {
    /**
     * Default value. This value is unused.
     */
    LoggableActionStateUnspecified: "LOGGABLE_ACTION_STATE_UNSPECIFIED",
    /**
     * `LoggableAction` completed successfully. `SUCCEEDED` actions are logged as INFO.
     */
    Succeeded: "SUCCEEDED",
    /**
     * `LoggableAction` terminated in an error state. `FAILED` actions are logged as ERROR.
     */
    Failed: "FAILED",
} as const;

export type LoggingConfigLogActionStatesItem = (typeof LoggingConfigLogActionStatesItem)[keyof typeof LoggingConfigLogActionStatesItem];

export const LoggingConfigLogActionsItem = {
    /**
     * Default value. This value is unused.
     */
    LoggableActionUnspecified: "LOGGABLE_ACTION_UNSPECIFIED",
    /**
     * Listing objects in a bucket.
     */
    Find: "FIND",
    /**
     * Deleting objects at the source or the destination.
     */
    Delete: "DELETE",
    /**
     * Copying objects to Google Cloud Storage.
     */
    Copy: "COPY",
} as const;

export type LoggingConfigLogActionsItem = (typeof LoggingConfigLogActionsItem)[keyof typeof LoggingConfigLogActionsItem];

export const MetadataOptionsGid = {
    /**
     * GID behavior is unspecified.
     */
    GidUnspecified: "GID_UNSPECIFIED",
    /**
     * Skip GID during a transfer job.
     */
    GidSkip: "GID_SKIP",
    /**
     * Preserve GID during a transfer job.
     */
    GidNumber: "GID_NUMBER",
} as const;

/**
 * Specifies how each file's GID attribute should be handled by the transfer. If unspecified, the default behavior is the same as GID_SKIP when the source is a POSIX file system.
 */
export type MetadataOptionsGid = (typeof MetadataOptionsGid)[keyof typeof MetadataOptionsGid];

export const MetadataOptionsMode = {
    /**
     * Mode behavior is unspecified.
     */
    ModeUnspecified: "MODE_UNSPECIFIED",
    /**
     * Skip mode during a transfer job.
     */
    ModeSkip: "MODE_SKIP",
    /**
     * Preserve mode during a transfer job.
     */
    ModePreserve: "MODE_PRESERVE",
} as const;

/**
 * Specifies how each file's mode attribute should be handled by the transfer. If unspecified, the default behavior is the same as MODE_SKIP when the source is a POSIX file system.
 */
export type MetadataOptionsMode = (typeof MetadataOptionsMode)[keyof typeof MetadataOptionsMode];

export const MetadataOptionsSymlink = {
    /**
     * Symlink behavior is unspecified. The default behavior is to skip symlinks during a transfer job.
     */
    SymlinkUnspecified: "SYMLINK_UNSPECIFIED",
    /**
     * Skip symlinks during a transfer job.
     */
    SymlinkSkip: "SYMLINK_SKIP",
    /**
     * Preserve symlinks during a transfer job.
     */
    SymlinkPreserve: "SYMLINK_PRESERVE",
} as const;

/**
 * Specifies how symlinks should be handled by the transfer. If unspecified, the default behavior is the same as SYMLINK_SKIP when the source is a POSIX file system.
 */
export type MetadataOptionsSymlink = (typeof MetadataOptionsSymlink)[keyof typeof MetadataOptionsSymlink];

export const MetadataOptionsUid = {
    /**
     * UID behavior is unspecified.
     */
    UidUnspecified: "UID_UNSPECIFIED",
    /**
     * Skip UID during a transfer job.
     */
    UidSkip: "UID_SKIP",
    /**
     * Preserve UID during a transfer job.
     */
    UidNumber: "UID_NUMBER",
} as const;

/**
 * Specifies how each file's UID attribute should be handled by the transfer. If unspecified, the default behavior is the same as UID_SKIP when the source is a POSIX file system.
 */
export type MetadataOptionsUid = (typeof MetadataOptionsUid)[keyof typeof MetadataOptionsUid];

export const NotificationConfigEventTypesItem = {
    /**
     * Illegal value, to avoid allowing a default.
     */
    EventTypeUnspecified: "EVENT_TYPE_UNSPECIFIED",
    /**
     * `TransferOperation` completed with status SUCCESS.
     */
    TransferOperationSuccess: "TRANSFER_OPERATION_SUCCESS",
    /**
     * `TransferOperation` completed with status FAILED.
     */
    TransferOperationFailed: "TRANSFER_OPERATION_FAILED",
    /**
     * `TransferOperation` completed with status ABORTED.
     */
    TransferOperationAborted: "TRANSFER_OPERATION_ABORTED",
} as const;

export type NotificationConfigEventTypesItem = (typeof NotificationConfigEventTypesItem)[keyof typeof NotificationConfigEventTypesItem];

export const NotificationConfigPayloadFormat = {
    /**
     * Illegal value, to avoid allowing a default.
     */
    PayloadFormatUnspecified: "PAYLOAD_FORMAT_UNSPECIFIED",
    /**
     * No payload is included with the notification.
     */
    None: "NONE",
    /**
     * `TransferOperation` is [formatted as a JSON response](https://developers.google.com/protocol-buffers/docs/proto3#json), in application/json.
     */
    Json: "JSON",
} as const;

/**
 * Required. The desired format of the notification message payloads.
 */
export type NotificationConfigPayloadFormat = (typeof NotificationConfigPayloadFormat)[keyof typeof NotificationConfigPayloadFormat];

export const TransferJobStatus = {
    /**
     * Zero is an illegal value.
     */
    StatusUnspecified: "STATUS_UNSPECIFIED",
    /**
     * New transfers are performed based on the schedule.
     */
    Enabled: "ENABLED",
    /**
     * New transfers are not scheduled.
     */
    Disabled: "DISABLED",
    /**
     * This is a soft delete state. After a transfer job is set to this state, the job and all the transfer executions are subject to garbage collection. Transfer jobs become eligible for garbage collection 30 days after their status is set to `DELETED`.
     */
    Deleted: "DELETED",
} as const;

/**
 * Status of the job. This value MUST be specified for `CreateTransferJobRequests`. **Note:** The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.
 */
export type TransferJobStatus = (typeof TransferJobStatus)[keyof typeof TransferJobStatus];
