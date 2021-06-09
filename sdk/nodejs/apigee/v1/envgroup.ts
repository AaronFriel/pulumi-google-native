// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Creates a new environment group.
 */
export class Envgroup extends pulumi.CustomResource {
    /**
     * Get an existing Envgroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Envgroup {
        return new Envgroup(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:apigee/v1:Envgroup';

    /**
     * Returns true if the given object is an instance of Envgroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Envgroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Envgroup.__pulumiType;
    }

    /**
     * The time at which the environment group was created as milliseconds since epoch.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Required. Host names for this environment group.
     */
    public readonly hostnames!: pulumi.Output<string[]>;
    /**
     * The time at which the environment group was last updated as milliseconds since epoch.
     */
    public /*out*/ readonly lastModifiedAt!: pulumi.Output<string>;
    /**
     * ID of the environment group.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * State of the environment group. Values other than ACTIVE means the resource is not ready to use.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;

    /**
     * Create a Envgroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EnvgroupArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.organizationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'organizationId'");
            }
            inputs["hostnames"] = args ? args.hostnames : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["organizationId"] = args ? args.organizationId : undefined;
            inputs["createdAt"] = undefined /*out*/;
            inputs["lastModifiedAt"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
        } else {
            inputs["createdAt"] = undefined /*out*/;
            inputs["hostnames"] = undefined /*out*/;
            inputs["lastModifiedAt"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Envgroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Envgroup resource.
 */
export interface EnvgroupArgs {
    /**
     * Required. Host names for this environment group.
     */
    readonly hostnames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * ID of the environment group.
     */
    readonly name?: pulumi.Input<string>;
    readonly organizationId: pulumi.Input<string>;
}