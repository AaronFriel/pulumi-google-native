// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Retrieves a resource containing information about a Cloud SQL instance.
 */
export function getInstance(args: GetInstanceArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("google-native:sqladmin/v1:getInstance", {
        "instance": args.instance,
        "project": args.project,
    }, opts);
}

export interface GetInstanceArgs {
    instance: string;
    project?: string;
}

export interface GetInstanceResult {
    /**
     * The backend type. **SECOND_GEN**: Cloud SQL database instance. **EXTERNAL**: A database server that is not managed by Google. This property is read-only; use the **tier** property in the **settings** object to determine the database type.
     */
    readonly backendType: string;
    /**
     * Connection name of the Cloud SQL instance used in connection strings.
     */
    readonly connectionName: string;
    /**
     * The time when the instance was created in [RFC 3339](https://tools.ietf.org/html/rfc3339) format, for example **2012-11-15T16:19:00.094Z**.
     */
    readonly createTime: string;
    /**
     * The current disk usage of the instance in bytes. This property has been deprecated. Use the "cloudsql.googleapis.com/database/disk/bytes_used" metric in Cloud Monitoring API instead. Please see [this announcement](https://groups.google.com/d/msg/google-cloud-sql-announce/I_7-F9EBhT0/BtvFtdFeAgAJ) for details.
     */
    readonly currentDiskSize: string;
    /**
     * Stores the current database version running on the instance including minor version such as **MYSQL_8_0_18**.
     */
    readonly databaseInstalledVersion: string;
    /**
     * The database engine type and version. The **databaseVersion** field cannot be changed after instance creation.
     */
    readonly databaseVersion: string;
    /**
     * Disk encryption configuration specific to an instance.
     */
    readonly diskEncryptionConfiguration: outputs.sqladmin.v1.DiskEncryptionConfigurationResponse;
    /**
     * Disk encryption status specific to an instance.
     */
    readonly diskEncryptionStatus: outputs.sqladmin.v1.DiskEncryptionStatusResponse;
    /**
     * The name and status of the failover replica.
     */
    readonly failoverReplica: outputs.sqladmin.v1.InstanceFailoverReplicaResponse;
    /**
     * The Compute Engine zone that the instance is currently serving from. This value could be different from the zone that was specified when the instance was created if the instance has failed over to its secondary zone. WARNING: Changing this might restart the instance.
     */
    readonly gceZone: string;
    /**
     * The instance type.
     */
    readonly instanceType: string;
    /**
     * The assigned IP addresses for the instance.
     */
    readonly ipAddresses: outputs.sqladmin.v1.IpMappingResponse[];
    /**
     * This is always **sql#instance**.
     */
    readonly kind: string;
    /**
     * The name of the instance which will act as primary in the replication setup.
     */
    readonly masterInstanceName: string;
    /**
     * The maximum disk size of the instance in bytes.
     */
    readonly maxDiskSize: string;
    /**
     * Name of the Cloud SQL instance. This does not include the project ID.
     */
    readonly name: string;
    /**
     * Configuration specific to on-premises instances.
     */
    readonly onPremisesConfiguration: outputs.sqladmin.v1.OnPremisesConfigurationResponse;
    /**
     * This field represents the report generated by the proactive database wellness job for OutOfDisk issues. * Writers: * the proactive database wellness job for OOD. * Readers: * the proactive database wellness job
     */
    readonly outOfDiskReport: outputs.sqladmin.v1.SqlOutOfDiskReportResponse;
    /**
     * The project ID of the project containing the Cloud SQL instance. The Google apps domain is prefixed if applicable.
     */
    readonly project: string;
    /**
     * The geographical region. Can be: * **us-central** (**FIRST_GEN** instances only) * **us-central1** (**SECOND_GEN** instances only) * **asia-east1** or **europe-west1**. Defaults to **us-central** or **us-central1** depending on the instance type. The region cannot be changed after instance creation.
     */
    readonly region: string;
    /**
     * Configuration specific to failover replicas and read replicas.
     */
    readonly replicaConfiguration: outputs.sqladmin.v1.ReplicaConfigurationResponse;
    /**
     * The replicas of the instance.
     */
    readonly replicaNames: string[];
    /**
     * Initial root password. Use only on creation.
     */
    readonly rootPassword: string;
    /**
     * The status indicating if instance satisfiesPzs. Reserved for future use.
     */
    readonly satisfiesPzs: boolean;
    /**
     * The start time of any upcoming scheduled maintenance for this instance.
     */
    readonly scheduledMaintenance: outputs.sqladmin.v1.SqlScheduledMaintenanceResponse;
    /**
     * The Compute Engine zone that the failover instance is currently serving from for a regional instance. This value could be different from the zone that was specified when the instance was created if the instance has failed over to its secondary/failover zone. Reserved for future use.
     */
    readonly secondaryGceZone: string;
    /**
     * The URI of this resource.
     */
    readonly selfLink: string;
    /**
     * SSL configuration.
     */
    readonly serverCaCert: outputs.sqladmin.v1.SslCertResponse;
    /**
     * The service account email address assigned to the instance. This property is read-only.
     */
    readonly serviceAccountEmailAddress: string;
    /**
     * The user settings.
     */
    readonly settings: outputs.sqladmin.v1.SettingsResponse;
    /**
     * The current serving state of the Cloud SQL instance.
     */
    readonly state: string;
    /**
     * If the instance state is SUSPENDED, the reason for the suspension.
     */
    readonly suspensionReason: string[];
}

export function getInstanceOutput(args: GetInstanceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstanceResult> {
    return pulumi.output(args).apply(a => getInstance(a, opts))
}

export interface GetInstanceOutputArgs {
    instance: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}