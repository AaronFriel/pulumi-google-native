// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./deliveryPipeline";
export * from "./deliveryPipelineIamPolicy";
export * from "./getDeliveryPipeline";
export * from "./getDeliveryPipelineIamPolicy";
export * from "./getRelease";
export * from "./getRollout";
export * from "./getTarget";
export * from "./getTargetIamPolicy";
export * from "./release";
export * from "./rollout";
export * from "./target";
export * from "./targetIamPolicy";

// Export enums:
export * from "../../types/enums/clouddeploy/v1";

// Import resources to register:
import { DeliveryPipeline } from "./deliveryPipeline";
import { DeliveryPipelineIamPolicy } from "./deliveryPipelineIamPolicy";
import { Release } from "./release";
import { Rollout } from "./rollout";
import { Target } from "./target";
import { TargetIamPolicy } from "./targetIamPolicy";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "google-native:clouddeploy/v1:DeliveryPipeline":
                return new DeliveryPipeline(name, <any>undefined, { urn })
            case "google-native:clouddeploy/v1:DeliveryPipelineIamPolicy":
                return new DeliveryPipelineIamPolicy(name, <any>undefined, { urn })
            case "google-native:clouddeploy/v1:Release":
                return new Release(name, <any>undefined, { urn })
            case "google-native:clouddeploy/v1:Rollout":
                return new Rollout(name, <any>undefined, { urn })
            case "google-native:clouddeploy/v1:Target":
                return new Target(name, <any>undefined, { urn })
            case "google-native:clouddeploy/v1:TargetIamPolicy":
                return new TargetIamPolicy(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("google-native", "clouddeploy/v1", _module)