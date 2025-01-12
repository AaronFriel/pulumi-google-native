// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Create a new domain mapping.
 * Auto-naming is currently not supported for this resource.
 */
export class DomainMapping extends pulumi.CustomResource {
    /**
     * Get an existing DomainMapping resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DomainMapping {
        return new DomainMapping(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:run/v1:DomainMapping';

    /**
     * Returns true if the given object is an instance of DomainMapping.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DomainMapping {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DomainMapping.__pulumiType;
    }

    /**
     * The API version for this call such as "domains.cloudrun.com/v1".
     */
    public readonly apiVersion!: pulumi.Output<string>;
    /**
     * The kind of resource, in this case "DomainMapping".
     */
    public readonly kind!: pulumi.Output<string>;
    /**
     * Metadata associated with this BuildTemplate.
     */
    public readonly metadata!: pulumi.Output<outputs.run.v1.ObjectMetaResponse>;
    /**
     * The spec for this DomainMapping.
     */
    public readonly spec!: pulumi.Output<outputs.run.v1.DomainMappingSpecResponse>;
    /**
     * The current status of the DomainMapping.
     */
    public readonly status!: pulumi.Output<outputs.run.v1.DomainMappingStatusResponse>;

    /**
     * Create a DomainMapping resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: DomainMappingArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["apiVersion"] = args ? args.apiVersion : undefined;
            resourceInputs["dryRun"] = args ? args.dryRun : undefined;
            resourceInputs["kind"] = args ? args.kind : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["metadata"] = args ? args.metadata : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["spec"] = args ? args.spec : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
        } else {
            resourceInputs["apiVersion"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["metadata"] = undefined /*out*/;
            resourceInputs["spec"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(DomainMapping.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a DomainMapping resource.
 */
export interface DomainMappingArgs {
    /**
     * The API version for this call such as "domains.cloudrun.com/v1".
     */
    apiVersion?: pulumi.Input<string>;
    dryRun?: pulumi.Input<string>;
    /**
     * The kind of resource, in this case "DomainMapping".
     */
    kind?: pulumi.Input<string>;
    location?: pulumi.Input<string>;
    /**
     * Metadata associated with this BuildTemplate.
     */
    metadata?: pulumi.Input<inputs.run.v1.ObjectMetaArgs>;
    project?: pulumi.Input<string>;
    /**
     * The spec for this DomainMapping.
     */
    spec?: pulumi.Input<inputs.run.v1.DomainMappingSpecArgs>;
    /**
     * The current status of the DomainMapping.
     */
    status?: pulumi.Input<inputs.run.v1.DomainMappingStatusArgs>;
}
