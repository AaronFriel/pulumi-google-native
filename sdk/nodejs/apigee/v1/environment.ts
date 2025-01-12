// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates an environment in an organization.
 */
export class Environment extends pulumi.CustomResource {
    /**
     * Get an existing Environment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Environment {
        return new Environment(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:apigee/v1:Environment';

    /**
     * Returns true if the given object is an instance of Environment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Environment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Environment.__pulumiType;
    }

    /**
     * Optional. API Proxy type supported by the environment. The type can be set when creating the Environment and cannot be changed.
     */
    public readonly apiProxyType!: pulumi.Output<string>;
    /**
     * Creation time of this environment as milliseconds since epoch.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Optional. Deployment type supported by the environment. The deployment type can be set when creating the environment and cannot be changed. When you enable archive deployment, you will be **prevented from performing** a [subset of actions](/apigee/docs/api-platform/local-development/overview#prevented-actions) within the environment, including: * Managing the deployment of API proxy or shared flow revisions * Creating, updating, or deleting resource files * Creating, updating, or deleting target servers
     */
    public readonly deploymentType!: pulumi.Output<string>;
    /**
     * Optional. Description of the environment.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Optional. Display name for this environment.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Last modification time of this environment as milliseconds since epoch.
     */
    public /*out*/ readonly lastModifiedAt!: pulumi.Output<string>;
    /**
     * Name of the environment. Values must match the regular expression `^[.\\p{Alnum}-_]{1,255}$`
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Optional. Key-value pairs that may be used for customizing the environment.
     */
    public readonly properties!: pulumi.Output<outputs.apigee.v1.GoogleCloudApigeeV1PropertiesResponse>;
    /**
     * State of the environment. Values other than ACTIVE means the resource is not ready to use.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;

    /**
     * Create a Environment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EnvironmentArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.organizationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'organizationId'");
            }
            resourceInputs["apiProxyType"] = args ? args.apiProxyType : undefined;
            resourceInputs["deploymentType"] = args ? args.deploymentType : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["organizationId"] = args ? args.organizationId : undefined;
            resourceInputs["properties"] = args ? args.properties : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["lastModifiedAt"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
        } else {
            resourceInputs["apiProxyType"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["deploymentType"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["lastModifiedAt"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["properties"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Environment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Environment resource.
 */
export interface EnvironmentArgs {
    /**
     * Optional. API Proxy type supported by the environment. The type can be set when creating the Environment and cannot be changed.
     */
    apiProxyType?: pulumi.Input<enums.apigee.v1.EnvironmentApiProxyType>;
    /**
     * Optional. Deployment type supported by the environment. The deployment type can be set when creating the environment and cannot be changed. When you enable archive deployment, you will be **prevented from performing** a [subset of actions](/apigee/docs/api-platform/local-development/overview#prevented-actions) within the environment, including: * Managing the deployment of API proxy or shared flow revisions * Creating, updating, or deleting resource files * Creating, updating, or deleting target servers
     */
    deploymentType?: pulumi.Input<enums.apigee.v1.EnvironmentDeploymentType>;
    /**
     * Optional. Description of the environment.
     */
    description?: pulumi.Input<string>;
    /**
     * Optional. Display name for this environment.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Name of the environment. Values must match the regular expression `^[.\\p{Alnum}-_]{1,255}$`
     */
    name?: pulumi.Input<string>;
    organizationId: pulumi.Input<string>;
    /**
     * Optional. Key-value pairs that may be used for customizing the environment.
     */
    properties?: pulumi.Input<inputs.apigee.v1.GoogleCloudApigeeV1PropertiesArgs>;
}
