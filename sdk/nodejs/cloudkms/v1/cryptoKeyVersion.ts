// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Create a new CryptoKeyVersion in a CryptoKey. The server will assign the next sequential id. If unset, state will be set to ENABLED.
 */
export class CryptoKeyVersion extends pulumi.CustomResource {
    /**
     * Get an existing CryptoKeyVersion resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): CryptoKeyVersion {
        return new CryptoKeyVersion(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:cloudkms/v1:CryptoKeyVersion';

    /**
     * Returns true if the given object is an instance of CryptoKeyVersion.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CryptoKeyVersion {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CryptoKeyVersion.__pulumiType;
    }

    /**
     * The CryptoKeyVersionAlgorithm that this CryptoKeyVersion supports.
     */
    public /*out*/ readonly algorithm!: pulumi.Output<string>;
    /**
     * Statement that was generated and signed by the HSM at key creation time. Use this statement to verify attributes of the key as stored on the HSM, independently of Google. Only provided for key versions with protection_level HSM.
     */
    public /*out*/ readonly attestation!: pulumi.Output<outputs.cloudkms.v1.KeyOperationAttestationResponse>;
    /**
     * The time at which this CryptoKeyVersion was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The time this CryptoKeyVersion's key material was destroyed. Only present if state is DESTROYED.
     */
    public /*out*/ readonly destroyEventTime!: pulumi.Output<string>;
    /**
     * The time this CryptoKeyVersion's key material is scheduled for destruction. Only present if state is DESTROY_SCHEDULED.
     */
    public /*out*/ readonly destroyTime!: pulumi.Output<string>;
    /**
     * ExternalProtectionLevelOptions stores a group of additional fields for configuring a CryptoKeyVersion that are specific to the EXTERNAL protection level.
     */
    public readonly externalProtectionLevelOptions!: pulumi.Output<outputs.cloudkms.v1.ExternalProtectionLevelOptionsResponse>;
    /**
     * The time this CryptoKeyVersion's key material was generated.
     */
    public /*out*/ readonly generateTime!: pulumi.Output<string>;
    /**
     * The root cause of an import failure. Only present if state is IMPORT_FAILED.
     */
    public /*out*/ readonly importFailureReason!: pulumi.Output<string>;
    /**
     * The name of the ImportJob used to import this CryptoKeyVersion. Only present if the underlying key material was imported.
     */
    public /*out*/ readonly importJob!: pulumi.Output<string>;
    /**
     * The time at which this CryptoKeyVersion's key material was imported.
     */
    public /*out*/ readonly importTime!: pulumi.Output<string>;
    /**
     * The resource name for this CryptoKeyVersion in the format `projects/*&#47;locations/*&#47;keyRings/*&#47;cryptoKeys/*&#47;cryptoKeyVersions/*`.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The ProtectionLevel describing how crypto operations are performed with this CryptoKeyVersion.
     */
    public /*out*/ readonly protectionLevel!: pulumi.Output<string>;
    /**
     * The current state of the CryptoKeyVersion.
     */
    public readonly state!: pulumi.Output<string>;

    /**
     * Create a CryptoKeyVersion resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CryptoKeyVersionArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.cryptoKeyId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cryptoKeyId'");
            }
            if ((!args || args.keyRingId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'keyRingId'");
            }
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            inputs["cryptoKeyId"] = args ? args.cryptoKeyId : undefined;
            inputs["externalProtectionLevelOptions"] = args ? args.externalProtectionLevelOptions : undefined;
            inputs["keyRingId"] = args ? args.keyRingId : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["state"] = args ? args.state : undefined;
            inputs["algorithm"] = undefined /*out*/;
            inputs["attestation"] = undefined /*out*/;
            inputs["createTime"] = undefined /*out*/;
            inputs["destroyEventTime"] = undefined /*out*/;
            inputs["destroyTime"] = undefined /*out*/;
            inputs["generateTime"] = undefined /*out*/;
            inputs["importFailureReason"] = undefined /*out*/;
            inputs["importJob"] = undefined /*out*/;
            inputs["importTime"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["protectionLevel"] = undefined /*out*/;
        } else {
            inputs["algorithm"] = undefined /*out*/;
            inputs["attestation"] = undefined /*out*/;
            inputs["createTime"] = undefined /*out*/;
            inputs["destroyEventTime"] = undefined /*out*/;
            inputs["destroyTime"] = undefined /*out*/;
            inputs["externalProtectionLevelOptions"] = undefined /*out*/;
            inputs["generateTime"] = undefined /*out*/;
            inputs["importFailureReason"] = undefined /*out*/;
            inputs["importJob"] = undefined /*out*/;
            inputs["importTime"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["protectionLevel"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(CryptoKeyVersion.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a CryptoKeyVersion resource.
 */
export interface CryptoKeyVersionArgs {
    readonly cryptoKeyId: pulumi.Input<string>;
    /**
     * ExternalProtectionLevelOptions stores a group of additional fields for configuring a CryptoKeyVersion that are specific to the EXTERNAL protection level.
     */
    readonly externalProtectionLevelOptions?: pulumi.Input<inputs.cloudkms.v1.ExternalProtectionLevelOptionsArgs>;
    readonly keyRingId: pulumi.Input<string>;
    readonly location: pulumi.Input<string>;
    readonly project: pulumi.Input<string>;
    /**
     * The current state of the CryptoKeyVersion.
     */
    readonly state?: pulumi.Input<string>;
}