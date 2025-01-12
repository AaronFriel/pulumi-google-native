// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Gets the specified AndroidApp.
 */
export function getAndroidApp(args: GetAndroidAppArgs, opts?: pulumi.InvokeOptions): Promise<GetAndroidAppResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("google-native:firebase/v1beta1:getAndroidApp", {
        "androidAppId": args.androidAppId,
        "project": args.project,
    }, opts);
}

export interface GetAndroidAppArgs {
    androidAppId: string;
    project?: string;
}

export interface GetAndroidAppResult {
    /**
     * Immutable. The globally unique, Firebase-assigned identifier for the `AndroidApp`. This identifier should be treated as an opaque token, as the data format is not specified.
     */
    readonly appId: string;
    /**
     * The user-assigned display name for the `AndroidApp`.
     */
    readonly displayName: string;
    /**
     * The resource name of the AndroidApp, in the format: projects/ PROJECT_IDENTIFIER/androidApps/APP_ID * PROJECT_IDENTIFIER: the parent Project's [`ProjectNumber`](../projects#FirebaseProject.FIELDS.project_number) ***(recommended)*** or its [`ProjectId`](../projects#FirebaseProject.FIELDS.project_id). Learn more about using project identifiers in Google's [AIP 2510 standard](https://google.aip.dev/cloud/2510). Note that the value for PROJECT_IDENTIFIER in any response body will be the `ProjectId`. * APP_ID: the globally unique, Firebase-assigned identifier for the App (see [`appId`](../projects.androidApps#AndroidApp.FIELDS.app_id)).
     */
    readonly name: string;
    /**
     * Immutable. The canonical package name of the Android app as would appear in the Google Play Developer Console.
     */
    readonly packageName: string;
    /**
     * Immutable. A user-assigned unique identifier of the parent FirebaseProject for the `AndroidApp`.
     */
    readonly project: string;
}

export function getAndroidAppOutput(args: GetAndroidAppOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAndroidAppResult> {
    return pulumi.output(args).apply(a => getAndroidApp(a, opts))
}

export interface GetAndroidAppOutputArgs {
    androidAppId: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
