// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a Custom Report for an Organization. A Custom Report provides Apigee Customers to create custom dashboards in addition to the standard dashboards which are provided. The Custom Report in its simplest form contains specifications about metrics, dimensions and filters. It is important to note that the custom report by itself does not provide an executable entity. The Edge UI converts the custom report definition into an analytics query and displays the result in a chart.
 */
export class Report extends pulumi.CustomResource {
    /**
     * Get an existing Report resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Report {
        return new Report(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:apigee/v1:Report';

    /**
     * Returns true if the given object is an instance of Report.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Report {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Report.__pulumiType;
    }

    /**
     * This field contains the chart type for the report
     */
    public readonly chartType!: pulumi.Output<string>;
    /**
     * Legacy field: not used. This field contains a list of comments associated with custom report
     */
    public readonly comments!: pulumi.Output<string[]>;
    /**
     * Unix time when the app was created json key: createdAt
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * This contains the list of dimensions for the report
     */
    public readonly dimensions!: pulumi.Output<string[]>;
    /**
     * This is the display name for the report
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Environment name
     */
    public /*out*/ readonly environment!: pulumi.Output<string>;
    /**
     * This field contains the filter expression
     */
    public readonly filter!: pulumi.Output<string>;
    /**
     * Legacy field: not used. Contains the from time for the report
     */
    public readonly fromTime!: pulumi.Output<string>;
    /**
     * Modified time of this entity as milliseconds since epoch. json key: lastModifiedAt
     */
    public /*out*/ readonly lastModifiedAt!: pulumi.Output<string>;
    /**
     * Last viewed time of this entity as milliseconds since epoch
     */
    public /*out*/ readonly lastViewedAt!: pulumi.Output<string>;
    /**
     * Legacy field: not used This field contains the limit for the result retrieved
     */
    public readonly limit!: pulumi.Output<string>;
    /**
     * Required. This contains the list of metrics
     */
    public readonly metrics!: pulumi.Output<outputs.apigee.v1.GoogleCloudApigeeV1CustomReportMetricResponse[]>;
    /**
     * Required. Unique identifier for the report T his is a legacy field used to encode custom report unique id
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Legacy field: not used. This field contains the offset for the data
     */
    public readonly offset!: pulumi.Output<string>;
    /**
     * Organization name
     */
    public /*out*/ readonly organization!: pulumi.Output<string>;
    /**
     * This field contains report properties such as ui metadata etc.
     */
    public readonly properties!: pulumi.Output<outputs.apigee.v1.GoogleCloudApigeeV1ReportPropertyResponse[]>;
    /**
     * Legacy field: not used much. Contains the list of sort by columns
     */
    public readonly sortByCols!: pulumi.Output<string[]>;
    /**
     * Legacy field: not used much. Contains the sort order for the sort columns
     */
    public readonly sortOrder!: pulumi.Output<string>;
    /**
     * Legacy field: not used. This field contains a list of tags associated with custom report
     */
    public readonly tags!: pulumi.Output<string[]>;
    /**
     * This field contains the time unit of aggregation for the report
     */
    public readonly timeUnit!: pulumi.Output<string>;
    /**
     * Legacy field: not used. Contains the end time for the report
     */
    public readonly toTime!: pulumi.Output<string>;
    /**
     * Legacy field: not used. This field contains the top k parameter value for restricting the result
     */
    public readonly topk!: pulumi.Output<string>;

    /**
     * Create a Report resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ReportArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.organizationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'organizationId'");
            }
            inputs["chartType"] = args ? args.chartType : undefined;
            inputs["comments"] = args ? args.comments : undefined;
            inputs["dimensions"] = args ? args.dimensions : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["filter"] = args ? args.filter : undefined;
            inputs["fromTime"] = args ? args.fromTime : undefined;
            inputs["limit"] = args ? args.limit : undefined;
            inputs["metrics"] = args ? args.metrics : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["offset"] = args ? args.offset : undefined;
            inputs["organizationId"] = args ? args.organizationId : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["sortByCols"] = args ? args.sortByCols : undefined;
            inputs["sortOrder"] = args ? args.sortOrder : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["timeUnit"] = args ? args.timeUnit : undefined;
            inputs["toTime"] = args ? args.toTime : undefined;
            inputs["topk"] = args ? args.topk : undefined;
            inputs["createdAt"] = undefined /*out*/;
            inputs["environment"] = undefined /*out*/;
            inputs["lastModifiedAt"] = undefined /*out*/;
            inputs["lastViewedAt"] = undefined /*out*/;
            inputs["organization"] = undefined /*out*/;
        } else {
            inputs["chartType"] = undefined /*out*/;
            inputs["comments"] = undefined /*out*/;
            inputs["createdAt"] = undefined /*out*/;
            inputs["dimensions"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["environment"] = undefined /*out*/;
            inputs["filter"] = undefined /*out*/;
            inputs["fromTime"] = undefined /*out*/;
            inputs["lastModifiedAt"] = undefined /*out*/;
            inputs["lastViewedAt"] = undefined /*out*/;
            inputs["limit"] = undefined /*out*/;
            inputs["metrics"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["offset"] = undefined /*out*/;
            inputs["organization"] = undefined /*out*/;
            inputs["properties"] = undefined /*out*/;
            inputs["sortByCols"] = undefined /*out*/;
            inputs["sortOrder"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["timeUnit"] = undefined /*out*/;
            inputs["toTime"] = undefined /*out*/;
            inputs["topk"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Report.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Report resource.
 */
export interface ReportArgs {
    /**
     * This field contains the chart type for the report
     */
    readonly chartType?: pulumi.Input<string>;
    /**
     * Legacy field: not used. This field contains a list of comments associated with custom report
     */
    readonly comments?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * This contains the list of dimensions for the report
     */
    readonly dimensions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * This is the display name for the report
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * This field contains the filter expression
     */
    readonly filter?: pulumi.Input<string>;
    /**
     * Legacy field: not used. Contains the from time for the report
     */
    readonly fromTime?: pulumi.Input<string>;
    /**
     * Legacy field: not used This field contains the limit for the result retrieved
     */
    readonly limit?: pulumi.Input<string>;
    /**
     * Required. This contains the list of metrics
     */
    readonly metrics?: pulumi.Input<pulumi.Input<inputs.apigee.v1.GoogleCloudApigeeV1CustomReportMetricArgs>[]>;
    /**
     * Required. Unique identifier for the report T his is a legacy field used to encode custom report unique id
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Legacy field: not used. This field contains the offset for the data
     */
    readonly offset?: pulumi.Input<string>;
    readonly organizationId: pulumi.Input<string>;
    /**
     * This field contains report properties such as ui metadata etc.
     */
    readonly properties?: pulumi.Input<pulumi.Input<inputs.apigee.v1.GoogleCloudApigeeV1ReportPropertyArgs>[]>;
    /**
     * Legacy field: not used much. Contains the list of sort by columns
     */
    readonly sortByCols?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Legacy field: not used much. Contains the sort order for the sort columns
     */
    readonly sortOrder?: pulumi.Input<string>;
    /**
     * Legacy field: not used. This field contains a list of tags associated with custom report
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * This field contains the time unit of aggregation for the report
     */
    readonly timeUnit?: pulumi.Input<string>;
    /**
     * Legacy field: not used. Contains the end time for the report
     */
    readonly toTime?: pulumi.Input<string>;
    /**
     * Legacy field: not used. This field contains the top k parameter value for restricting the result
     */
    readonly topk?: pulumi.Input<string>;
}