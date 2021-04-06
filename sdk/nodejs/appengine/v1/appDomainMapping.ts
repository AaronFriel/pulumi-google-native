// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Maps a domain to an application. A user must be authorized to administer a domain in order to map it to an application. For a list of available authorized domains, see AuthorizedDomains.ListAuthorizedDomains.
 */
export class AppDomainMapping extends pulumi.CustomResource {
    /**
     * Get an existing AppDomainMapping resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): AppDomainMapping {
        return new AppDomainMapping(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-cloud:appengine/v1:AppDomainMapping';

    /**
     * Returns true if the given object is an instance of AppDomainMapping.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AppDomainMapping {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AppDomainMapping.__pulumiType;
    }

    /**
     * Full path to the DomainMapping resource in the API. Example: apps/myapp/domainMapping/example.com.@OutputOnly
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The resource records required to configure this domain mapping. These records must be added to the domain's DNS configuration in order to serve the application via this domain mapping.@OutputOnly
     */
    public readonly resourceRecords!: pulumi.Output<outputs.appengine.v1.ResourceRecordResponse[]>;
    /**
     * SSL configuration for this domain. If unconfigured, this domain will not serve with SSL.
     */
    public readonly sslSettings!: pulumi.Output<outputs.appengine.v1.SslSettingsResponse>;

    /**
     * Create a AppDomainMapping resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AppDomainMappingArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.appsId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'appsId'");
            }
            if ((!args || args.domainMappingsId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domainMappingsId'");
            }
            inputs["appsId"] = args ? args.appsId : undefined;
            inputs["domainMappingsId"] = args ? args.domainMappingsId : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceRecords"] = args ? args.resourceRecords : undefined;
            inputs["sslSettings"] = args ? args.sslSettings : undefined;
        } else {
            inputs["name"] = undefined /*out*/;
            inputs["resourceRecords"] = undefined /*out*/;
            inputs["sslSettings"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AppDomainMapping.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a AppDomainMapping resource.
 */
export interface AppDomainMappingArgs {
    readonly appsId: pulumi.Input<string>;
    readonly domainMappingsId: pulumi.Input<string>;
    /**
     * Relative name of the domain serving the application. Example: example.com.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * Full path to the DomainMapping resource in the API. Example: apps/myapp/domainMapping/example.com.@OutputOnly
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The resource records required to configure this domain mapping. These records must be added to the domain's DNS configuration in order to serve the application via this domain mapping.@OutputOnly
     */
    readonly resourceRecords?: pulumi.Input<pulumi.Input<inputs.appengine.v1.ResourceRecord>[]>;
    /**
     * SSL configuration for this domain. If unconfigured, this domain will not serve with SSL.
     */
    readonly sslSettings?: pulumi.Input<inputs.appengine.v1.SslSettings>;
}