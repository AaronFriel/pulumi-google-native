// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Returns information about a `WorkerPool`.
 */
export function getWorkerPool(args: GetWorkerPoolArgs, opts?: pulumi.InvokeOptions): Promise<GetWorkerPoolResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("google-native:cloudbuild/v1alpha1:getWorkerPool", {
        "project": args.project,
        "workerPoolId": args.workerPoolId,
    }, opts);
}

export interface GetWorkerPoolArgs {
    project?: string;
    workerPoolId: string;
}

export interface GetWorkerPoolResult {
    /**
     * Time at which the request to create the `WorkerPool` was received.
     */
    readonly createTime: string;
    /**
     * Time at which the request to delete the `WorkerPool` was received.
     */
    readonly deleteTime: string;
    /**
     * User-defined name of the `WorkerPool`.
     */
    readonly name: string;
    /**
     * The project ID of the GCP project for which the `WorkerPool` is created.
     */
    readonly project: string;
    /**
     * List of regions to create the `WorkerPool`. Regions can't be empty. If Cloud Build adds a new GCP region in the future, the existing `WorkerPool` will not be enabled in the new region automatically; you must add the new region to the `regions` field to enable the `WorkerPool` in that region.
     */
    readonly regions: string[];
    /**
     * The service account used to manage the `WorkerPool`. The service account must have the Compute Instance Admin (Beta) permission at the project level.
     */
    readonly serviceAccountEmail: string;
    /**
     * WorkerPool Status.
     */
    readonly status: string;
    /**
     * Time at which the request to update the `WorkerPool` was received.
     */
    readonly updateTime: string;
    /**
     * Configuration to be used for a creating workers in the `WorkerPool`.
     */
    readonly workerConfig: outputs.cloudbuild.v1alpha1.WorkerConfigResponse;
    /**
     * Total number of workers to be created across all requested regions.
     */
    readonly workerCount: string;
}

export function getWorkerPoolOutput(args: GetWorkerPoolOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetWorkerPoolResult> {
    return pulumi.output(args).apply(a => getWorkerPool(a, opts))
}

export interface GetWorkerPoolOutputArgs {
    project?: pulumi.Input<string>;
    workerPoolId: pulumi.Input<string>;
}