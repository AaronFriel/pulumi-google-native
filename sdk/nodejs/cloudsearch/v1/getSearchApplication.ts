// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Gets the specified search application. **Note:** This API requires an admin account to execute.
 */
export function getSearchApplication(args: GetSearchApplicationArgs, opts?: pulumi.InvokeOptions): Promise<GetSearchApplicationResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("google-native:cloudsearch/v1:getSearchApplication", {
        "debugOptionsEnableDebugging": args.debugOptionsEnableDebugging,
        "searchapplicationId": args.searchapplicationId,
    }, opts);
}

export interface GetSearchApplicationArgs {
    debugOptionsEnableDebugging?: string;
    searchapplicationId: string;
}

export interface GetSearchApplicationResult {
    /**
     * Retrictions applied to the configurations. The maximum number of elements is 10.
     */
    readonly dataSourceRestrictions: outputs.cloudsearch.v1.DataSourceRestrictionResponse[];
    /**
     * The default fields for returning facet results. The sources specified here also have been included in data_source_restrictions above.
     */
    readonly defaultFacetOptions: outputs.cloudsearch.v1.FacetOptionsResponse[];
    /**
     * The default options for sorting the search results
     */
    readonly defaultSortOptions: outputs.cloudsearch.v1.SortOptionsResponse;
    /**
     * Display name of the Search Application. The maximum length is 300 characters.
     */
    readonly displayName: string;
    /**
     * Indicates whether audit logging is on/off for requests made for the search application in query APIs.
     */
    readonly enableAuditLog: boolean;
    /**
     * Name of the Search Application. Format: searchapplications/{application_id}.
     */
    readonly name: string;
    /**
     * IDs of the Long Running Operations (LROs) currently running for this schema. Output only field.
     */
    readonly operationIds: string[];
    /**
     * The default options for query interpretation
     */
    readonly queryInterpretationConfig: outputs.cloudsearch.v1.QueryInterpretationConfigResponse;
    /**
     * With each result we should return the URI for its thumbnail (when applicable)
     */
    readonly returnResultThumbnailUrls: boolean;
    /**
     * Configuration for ranking results.
     */
    readonly scoringConfig: outputs.cloudsearch.v1.ScoringConfigResponse;
    /**
     * Configuration for a sources specified in data_source_restrictions.
     */
    readonly sourceConfig: outputs.cloudsearch.v1.SourceConfigResponse[];
}

export function getSearchApplicationOutput(args: GetSearchApplicationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSearchApplicationResult> {
    return pulumi.output(args).apply(a => getSearchApplication(a, opts))
}

export interface GetSearchApplicationOutputArgs {
    debugOptionsEnableDebugging?: pulumi.Input<string>;
    searchapplicationId: pulumi.Input<string>;
}