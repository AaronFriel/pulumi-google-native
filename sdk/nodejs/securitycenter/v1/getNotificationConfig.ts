// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Gets a notification config.
 */
export function getNotificationConfig(args: GetNotificationConfigArgs, opts?: pulumi.InvokeOptions): Promise<GetNotificationConfigResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("google-native:securitycenter/v1:getNotificationConfig", {
        "notificationConfigId": args.notificationConfigId,
        "organizationId": args.organizationId,
    }, opts);
}

export interface GetNotificationConfigArgs {
    notificationConfigId: string;
    organizationId: string;
}

export interface GetNotificationConfigResult {
    /**
     * The description of the notification config (max of 1024 characters).
     */
    readonly description: string;
    /**
     * The relative resource name of this notification config. See: https://cloud.google.com/apis/design/resource_names#relative_resource_name Example: "organizations/{organization_id}/notificationConfigs/notify_public_bucket".
     */
    readonly name: string;
    /**
     * The Pub/Sub topic to send notifications to. Its format is "projects/[project_id]/topics/[topic]".
     */
    readonly pubsubTopic: string;
    /**
     * The service account that needs "pubsub.topics.publish" permission to publish to the Pub/Sub topic.
     */
    readonly serviceAccount: string;
    /**
     * The config for triggering streaming-based notifications.
     */
    readonly streamingConfig: outputs.securitycenter.v1.StreamingConfigResponse;
}

export function getNotificationConfigOutput(args: GetNotificationConfigOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNotificationConfigResult> {
    return pulumi.output(args).apply(a => getNotificationConfig(a, opts))
}

export interface GetNotificationConfigOutputArgs {
    notificationConfigId: pulumi.Input<string>;
    organizationId: pulumi.Input<string>;
}