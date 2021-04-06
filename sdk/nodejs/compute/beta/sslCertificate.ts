// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a SslCertificate resource in the specified project using the data included in the request.
 */
export class SslCertificate extends pulumi.CustomResource {
    /**
     * Get an existing SslCertificate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): SslCertificate {
        return new SslCertificate(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-cloud:compute/beta:SslCertificate';

    /**
     * Returns true if the given object is an instance of SslCertificate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SslCertificate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SslCertificate.__pulumiType;
    }

    /**
     * A value read into memory from a certificate file. The certificate file must be in PEM format. The certificate chain must be no greater than 5 certs long. The chain must include at least one intermediate cert.
     */
    public readonly certificate!: pulumi.Output<string>;
    /**
     * [Output Only] Creation timestamp in RFC3339 text format.
     */
    public readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * [Output Only] Expire time of the certificate. RFC3339
     */
    public readonly expireTime!: pulumi.Output<string>;
    /**
     * [Output Only] Type of the resource. Always compute#sslCertificate for SSL certificates.
     */
    public readonly kind!: pulumi.Output<string>;
    /**
     * Configuration and status of a managed SSL certificate.
     */
    public readonly managed!: pulumi.Output<outputs.compute.beta.SslCertificateManagedSslCertificateResponse>;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A value read into memory from a write-only private key file. The private key file must be in PEM format. For security, only insert requests include this field.
     */
    public readonly privateKey!: pulumi.Output<string>;
    /**
     * [Output Only] URL of the region where the regional SSL Certificate resides. This field is not applicable to global SSL Certificate.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * [Output only] Server-defined URL for the resource.
     */
    public readonly selfLink!: pulumi.Output<string>;
    /**
     * Configuration and status of a self-managed SSL certificate.
     */
    public readonly selfManaged!: pulumi.Output<outputs.compute.beta.SslCertificateSelfManagedSslCertificateResponse>;
    /**
     * [Output Only] Domains associated with the certificate via Subject Alternative Name.
     */
    public readonly subjectAlternativeNames!: pulumi.Output<string[]>;
    /**
     * (Optional) Specifies the type of SSL certificate, either "SELF_MANAGED" or "MANAGED". If not specified, the certificate is self-managed and the fields certificate and private_key are used.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a SslCertificate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SslCertificateArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.sslCertificate === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sslCertificate'");
            }
            inputs["certificate"] = args ? args.certificate : undefined;
            inputs["creationTimestamp"] = args ? args.creationTimestamp : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["expireTime"] = args ? args.expireTime : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["managed"] = args ? args.managed : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["privateKey"] = args ? args.privateKey : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["selfLink"] = args ? args.selfLink : undefined;
            inputs["selfManaged"] = args ? args.selfManaged : undefined;
            inputs["sslCertificate"] = args ? args.sslCertificate : undefined;
            inputs["subjectAlternativeNames"] = args ? args.subjectAlternativeNames : undefined;
            inputs["type"] = args ? args.type : undefined;
        } else {
            inputs["certificate"] = undefined /*out*/;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["expireTime"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["managed"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["privateKey"] = undefined /*out*/;
            inputs["region"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
            inputs["selfManaged"] = undefined /*out*/;
            inputs["subjectAlternativeNames"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(SslCertificate.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a SslCertificate resource.
 */
export interface SslCertificateArgs {
    /**
     * A value read into memory from a certificate file. The certificate file must be in PEM format. The certificate chain must be no greater than 5 certs long. The chain must include at least one intermediate cert.
     */
    readonly certificate?: pulumi.Input<string>;
    /**
     * [Output Only] Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp?: pulumi.Input<string>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * [Output Only] Expire time of the certificate. RFC3339
     */
    readonly expireTime?: pulumi.Input<string>;
    /**
     * [Output Only] The unique identifier for the resource. This identifier is defined by the server.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * [Output Only] Type of the resource. Always compute#sslCertificate for SSL certificates.
     */
    readonly kind?: pulumi.Input<string>;
    /**
     * Configuration and status of a managed SSL certificate.
     */
    readonly managed?: pulumi.Input<inputs.compute.beta.SslCertificateManagedSslCertificate>;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A value read into memory from a write-only private key file. The private key file must be in PEM format. For security, only insert requests include this field.
     */
    readonly privateKey?: pulumi.Input<string>;
    readonly project: pulumi.Input<string>;
    /**
     * [Output Only] URL of the region where the regional SSL Certificate resides. This field is not applicable to global SSL Certificate.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * [Output only] Server-defined URL for the resource.
     */
    readonly selfLink?: pulumi.Input<string>;
    /**
     * Configuration and status of a self-managed SSL certificate.
     */
    readonly selfManaged?: pulumi.Input<inputs.compute.beta.SslCertificateSelfManagedSslCertificate>;
    readonly sslCertificate: pulumi.Input<string>;
    /**
     * [Output Only] Domains associated with the certificate via Subject Alternative Name.
     */
    readonly subjectAlternativeNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * (Optional) Specifies the type of SSL certificate, either "SELF_MANAGED" or "MANAGED". If not specified, the certificate is self-managed and the fields certificate and private_key are used.
     */
    readonly type?: pulumi.Input<string>;
}