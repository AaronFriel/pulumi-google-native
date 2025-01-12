// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Creates a tag template. You must enable the Data Catalog API in the project identified by the `parent` parameter. For more information, see [Data Catalog resource project] (https://cloud.google.com/data-catalog/docs/concepts/resource-project).
 */
export class TagTemplate extends pulumi.CustomResource {
    /**
     * Get an existing TagTemplate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): TagTemplate {
        return new TagTemplate(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:datacatalog/v1:TagTemplate';

    /**
     * Returns true if the given object is an instance of TagTemplate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TagTemplate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TagTemplate.__pulumiType;
    }

    /**
     * Display name for this template. Defaults to an empty string. The name must contain only Unicode letters, numbers (0-9), underscores (_), dashes (-), spaces ( ), and can't start or end with spaces. The maximum length is 200 characters.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Map of tag template field IDs to the settings for the field. This map is an exhaustive list of the allowed fields. The map must contain at least one field and at most 500 fields. The keys to this map are tag template field IDs. The IDs have the following limitations: * Can contain uppercase and lowercase letters, numbers (0-9) and underscores (_). * Must be at least 1 character and at most 64 characters long. * Must start with a letter or underscore.
     */
    public readonly fields!: pulumi.Output<{[key: string]: string}>;
    /**
     * Indicates whether this is a public tag template. Every user has view access to a *public* tag template by default. This means that: * Every user can use this tag template to tag an entry. * If an entry is tagged using the tag template, the tag is always shown in the response to ``ListTags`` called on the entry. * To get the template using the GetTagTemplate method, you need view access either on the project or the organization the tag template resides in but no other permission is needed. * Operations on the tag template other than viewing (for example, editing IAM policies) follow standard IAM structures. Tags created with a public tag template are referred to as public tags. You can search for a public tag by value with a simple search query instead of using a ``tag:`` predicate. Public tag templates may not appear in search results depending on scope, see: include_public_tag_templates Note: If an [IAM domain restriction](https://cloud.google.com/resource-manager/docs/organization-policy/restricting-domains) is configured in the tag template's location, the public access will not be enabled but the simple search for tag values will still work.
     */
    public readonly isPubliclyReadable!: pulumi.Output<boolean>;
    /**
     * The resource name of the tag template in URL format. Note: The tag template itself and its child resources might not be stored in the location specified in its name.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a TagTemplate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TagTemplateArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.fields === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fields'");
            }
            if ((!args || args.tagTemplateId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tagTemplateId'");
            }
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["fields"] = args ? args.fields : undefined;
            resourceInputs["isPubliclyReadable"] = args ? args.isPubliclyReadable : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["tagTemplateId"] = args ? args.tagTemplateId : undefined;
        } else {
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["fields"] = undefined /*out*/;
            resourceInputs["isPubliclyReadable"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(TagTemplate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a TagTemplate resource.
 */
export interface TagTemplateArgs {
    /**
     * Display name for this template. Defaults to an empty string. The name must contain only Unicode letters, numbers (0-9), underscores (_), dashes (-), spaces ( ), and can't start or end with spaces. The maximum length is 200 characters.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Map of tag template field IDs to the settings for the field. This map is an exhaustive list of the allowed fields. The map must contain at least one field and at most 500 fields. The keys to this map are tag template field IDs. The IDs have the following limitations: * Can contain uppercase and lowercase letters, numbers (0-9) and underscores (_). * Must be at least 1 character and at most 64 characters long. * Must start with a letter or underscore.
     */
    fields: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Indicates whether this is a public tag template. Every user has view access to a *public* tag template by default. This means that: * Every user can use this tag template to tag an entry. * If an entry is tagged using the tag template, the tag is always shown in the response to ``ListTags`` called on the entry. * To get the template using the GetTagTemplate method, you need view access either on the project or the organization the tag template resides in but no other permission is needed. * Operations on the tag template other than viewing (for example, editing IAM policies) follow standard IAM structures. Tags created with a public tag template are referred to as public tags. You can search for a public tag by value with a simple search query instead of using a ``tag:`` predicate. Public tag templates may not appear in search results depending on scope, see: include_public_tag_templates Note: If an [IAM domain restriction](https://cloud.google.com/resource-manager/docs/organization-policy/restricting-domains) is configured in the tag template's location, the public access will not be enabled but the simple search for tag values will still work.
     */
    isPubliclyReadable?: pulumi.Input<boolean>;
    location?: pulumi.Input<string>;
    /**
     * The resource name of the tag template in URL format. Note: The tag template itself and its child resources might not be stored in the location specified in its name.
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    tagTemplateId: pulumi.Input<string>;
}
