// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Submits a job to a cluster.
 */
export class Job extends pulumi.CustomResource {
    /**
     * Get an existing Job resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Job {
        return new Job(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:dataproc/v1beta2:Job';

    /**
     * Returns true if the given object is an instance of Job.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Job {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Job.__pulumiType;
    }

    /**
     * Indicates whether the job is completed. If the value is false, the job is still in progress. If true, the job is completed, and status.state field will indicate if it was successful, failed, or cancelled.
     */
    public /*out*/ readonly done!: pulumi.Output<boolean>;
    /**
     * If present, the location of miscellaneous control files which may be used as part of job setup and handling. If not present, control files may be placed in the same location as driver_output_uri.
     */
    public /*out*/ readonly driverControlFilesUri!: pulumi.Output<string>;
    /**
     * A URI pointing to the location of the stdout of the job's driver program.
     */
    public /*out*/ readonly driverOutputResourceUri!: pulumi.Output<string>;
    /**
     * Optional. Job is a Hadoop job.
     */
    public readonly hadoopJob!: pulumi.Output<outputs.dataproc.v1beta2.HadoopJobResponse>;
    /**
     * Optional. Job is a Hive job.
     */
    public readonly hiveJob!: pulumi.Output<outputs.dataproc.v1beta2.HiveJobResponse>;
    /**
     * A UUID that uniquely identifies a job within the project over time. This is in contrast to a user-settable reference.job_id that may be reused over time.
     */
    public /*out*/ readonly jobUuid!: pulumi.Output<string>;
    /**
     * Optional. The labels to associate with this job. Label keys must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). Label values may be empty, but, if present, must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). No more than 32 labels can be associated with a job.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Optional. Job is a Pig job.
     */
    public readonly pigJob!: pulumi.Output<outputs.dataproc.v1beta2.PigJobResponse>;
    /**
     * Required. Job information, including how, when, and where to run the job.
     */
    public readonly placement!: pulumi.Output<outputs.dataproc.v1beta2.JobPlacementResponse>;
    /**
     * Optional. Job is a Presto job.
     */
    public readonly prestoJob!: pulumi.Output<outputs.dataproc.v1beta2.PrestoJobResponse>;
    /**
     * Optional. Job is a PySpark job.
     */
    public readonly pysparkJob!: pulumi.Output<outputs.dataproc.v1beta2.PySparkJobResponse>;
    /**
     * Optional. The fully qualified reference to the job, which can be used to obtain the equivalent REST path of the job resource. If this property is not specified when a job is created, the server generates a job_id.
     */
    public readonly reference!: pulumi.Output<outputs.dataproc.v1beta2.JobReferenceResponse>;
    /**
     * Optional. Job scheduling configuration.
     */
    public readonly scheduling!: pulumi.Output<outputs.dataproc.v1beta2.JobSchedulingResponse>;
    /**
     * Optional. Job is a Spark job.
     */
    public readonly sparkJob!: pulumi.Output<outputs.dataproc.v1beta2.SparkJobResponse>;
    /**
     * Optional. Job is a SparkR job.
     */
    public readonly sparkRJob!: pulumi.Output<outputs.dataproc.v1beta2.SparkRJobResponse>;
    /**
     * Optional. Job is a SparkSql job.
     */
    public readonly sparkSqlJob!: pulumi.Output<outputs.dataproc.v1beta2.SparkSqlJobResponse>;
    /**
     * The job status. Additional application-specific status information may be contained in the type_job and yarn_applications fields.
     */
    public /*out*/ readonly status!: pulumi.Output<outputs.dataproc.v1beta2.JobStatusResponse>;
    /**
     * The previous job status.
     */
    public /*out*/ readonly statusHistory!: pulumi.Output<outputs.dataproc.v1beta2.JobStatusResponse[]>;
    /**
     * The email address of the user submitting the job. For jobs submitted on the cluster, the address is username@hostname.
     */
    public /*out*/ readonly submittedBy!: pulumi.Output<string>;
    /**
     * The collection of YARN applications spun up by this job.Beta Feature: This report is available for testing purposes only. It may be changed before final release.
     */
    public /*out*/ readonly yarnApplications!: pulumi.Output<outputs.dataproc.v1beta2.YarnApplicationResponse[]>;

    /**
     * Create a Job resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: JobArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            inputs["hadoopJob"] = args ? args.hadoopJob : undefined;
            inputs["hiveJob"] = args ? args.hiveJob : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["pigJob"] = args ? args.pigJob : undefined;
            inputs["placement"] = args ? args.placement : undefined;
            inputs["prestoJob"] = args ? args.prestoJob : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["pysparkJob"] = args ? args.pysparkJob : undefined;
            inputs["reference"] = args ? args.reference : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["requestId"] = args ? args.requestId : undefined;
            inputs["scheduling"] = args ? args.scheduling : undefined;
            inputs["sparkJob"] = args ? args.sparkJob : undefined;
            inputs["sparkRJob"] = args ? args.sparkRJob : undefined;
            inputs["sparkSqlJob"] = args ? args.sparkSqlJob : undefined;
            inputs["done"] = undefined /*out*/;
            inputs["driverControlFilesUri"] = undefined /*out*/;
            inputs["driverOutputResourceUri"] = undefined /*out*/;
            inputs["jobUuid"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["statusHistory"] = undefined /*out*/;
            inputs["submittedBy"] = undefined /*out*/;
            inputs["yarnApplications"] = undefined /*out*/;
        } else {
            inputs["done"] = undefined /*out*/;
            inputs["driverControlFilesUri"] = undefined /*out*/;
            inputs["driverOutputResourceUri"] = undefined /*out*/;
            inputs["hadoopJob"] = undefined /*out*/;
            inputs["hiveJob"] = undefined /*out*/;
            inputs["jobUuid"] = undefined /*out*/;
            inputs["labels"] = undefined /*out*/;
            inputs["pigJob"] = undefined /*out*/;
            inputs["placement"] = undefined /*out*/;
            inputs["prestoJob"] = undefined /*out*/;
            inputs["pysparkJob"] = undefined /*out*/;
            inputs["reference"] = undefined /*out*/;
            inputs["scheduling"] = undefined /*out*/;
            inputs["sparkJob"] = undefined /*out*/;
            inputs["sparkRJob"] = undefined /*out*/;
            inputs["sparkSqlJob"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["statusHistory"] = undefined /*out*/;
            inputs["submittedBy"] = undefined /*out*/;
            inputs["yarnApplications"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Job.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Job resource.
 */
export interface JobArgs {
    /**
     * Optional. Job is a Hadoop job.
     */
    readonly hadoopJob?: pulumi.Input<inputs.dataproc.v1beta2.HadoopJobArgs>;
    /**
     * Optional. Job is a Hive job.
     */
    readonly hiveJob?: pulumi.Input<inputs.dataproc.v1beta2.HiveJobArgs>;
    /**
     * Optional. The labels to associate with this job. Label keys must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). Label values may be empty, but, if present, must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). No more than 32 labels can be associated with a job.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Optional. Job is a Pig job.
     */
    readonly pigJob?: pulumi.Input<inputs.dataproc.v1beta2.PigJobArgs>;
    /**
     * Required. Job information, including how, when, and where to run the job.
     */
    readonly placement?: pulumi.Input<inputs.dataproc.v1beta2.JobPlacementArgs>;
    /**
     * Optional. Job is a Presto job.
     */
    readonly prestoJob?: pulumi.Input<inputs.dataproc.v1beta2.PrestoJobArgs>;
    readonly project: pulumi.Input<string>;
    /**
     * Optional. Job is a PySpark job.
     */
    readonly pysparkJob?: pulumi.Input<inputs.dataproc.v1beta2.PySparkJobArgs>;
    /**
     * Optional. The fully qualified reference to the job, which can be used to obtain the equivalent REST path of the job resource. If this property is not specified when a job is created, the server generates a job_id.
     */
    readonly reference?: pulumi.Input<inputs.dataproc.v1beta2.JobReferenceArgs>;
    readonly region: pulumi.Input<string>;
    /**
     * Optional. A unique id used to identify the request. If the server receives two SubmitJobRequest (https://cloud.google.com/dataproc/docs/reference/rpc/google.cloud.dataproc.v1beta2#google.cloud.dataproc.v1.SubmitJobRequest)s with the same id, then the second request will be ignored and the first Job created and stored in the backend is returned.It is recommended to always set this value to a UUID (https://en.wikipedia.org/wiki/Universally_unique_identifier).The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). The maximum length is 40 characters.
     */
    readonly requestId?: pulumi.Input<string>;
    /**
     * Optional. Job scheduling configuration.
     */
    readonly scheduling?: pulumi.Input<inputs.dataproc.v1beta2.JobSchedulingArgs>;
    /**
     * Optional. Job is a Spark job.
     */
    readonly sparkJob?: pulumi.Input<inputs.dataproc.v1beta2.SparkJobArgs>;
    /**
     * Optional. Job is a SparkR job.
     */
    readonly sparkRJob?: pulumi.Input<inputs.dataproc.v1beta2.SparkRJobArgs>;
    /**
     * Optional. Job is a SparkSql job.
     */
    readonly sparkSqlJob?: pulumi.Input<inputs.dataproc.v1beta2.SparkSqlJobArgs>;
}