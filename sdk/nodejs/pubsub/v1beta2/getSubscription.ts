// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Gets the configuration details of a subscription.
 */
export function getSubscription(args: GetSubscriptionArgs, opts?: pulumi.InvokeOptions): Promise<GetSubscriptionResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("google-native:pubsub/v1beta2:getSubscription", {
        "project": args.project,
        "subscriptionId": args.subscriptionId,
    }, opts);
}

export interface GetSubscriptionArgs {
    project?: string;
    subscriptionId: string;
}

export interface GetSubscriptionResult {
    /**
     * This value is the maximum time after a subscriber receives a message before the subscriber should acknowledge the message. After message delivery but before the ack deadline expires and before the message is acknowledged, it is an outstanding message and will not be delivered again during that time (on a best-effort basis). For pull subscriptions, this value is used as the initial value for the ack deadline. To override this value for a given message, call `ModifyAckDeadline` with the corresponding `ack_id` if using pull. The maximum custom deadline you can specify is 600 seconds (10 minutes). For push delivery, this value is also used to set the request timeout for the call to the push endpoint. If the subscriber never acknowledges the message, the Pub/Sub system will eventually redeliver the message. If this parameter is 0, a default value of 10 seconds is used.
     */
    readonly ackDeadlineSeconds: number;
    /**
     * The name of the subscription. It must have the format `"projects/{project}/subscriptions/{subscription}"`. `{subscription}` must start with a letter, and contain only letters (`[A-Za-z]`), numbers (`[0-9]`), dashes (`-`), underscores (`_`), periods (`.`), tildes (`~`), plus (`+`) or percent signs (`%`). It must be between 3 and 255 characters in length, and it must not start with `"goog"`.
     */
    readonly name: string;
    /**
     * If push delivery is used with this subscription, this field is used to configure it. An empty `pushConfig` signifies that the subscriber will pull and ack messages using API methods.
     */
    readonly pushConfig: outputs.pubsub.v1beta2.PushConfigResponse;
    /**
     * The name of the topic from which this subscription is receiving messages. The value of this field will be `_deleted-topic_` if the topic has been deleted.
     */
    readonly topic: string;
}

export function getSubscriptionOutput(args: GetSubscriptionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSubscriptionResult> {
    return pulumi.output(args).apply(a => getSubscription(a, opts))
}

export interface GetSubscriptionOutputArgs {
    project?: pulumi.Input<string>;
    subscriptionId: pulumi.Input<string>;
}