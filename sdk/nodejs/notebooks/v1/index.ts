// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./environment";
export * from "./execution";
export * from "./getEnvironment";
export * from "./getExecution";
export * from "./getInstance";
export * from "./getInstanceIamPolicy";
export * from "./getRuntime";
export * from "./getRuntimeIamPolicy";
export * from "./getSchedule";
export * from "./instance";
export * from "./instanceIamPolicy";
export * from "./runtime";
export * from "./runtimeIamPolicy";
export * from "./schedule";

// Export enums:
export * from "../../types/enums/notebooks/v1";

// Import resources to register:
import { Environment } from "./environment";
import { Execution } from "./execution";
import { Instance } from "./instance";
import { InstanceIamPolicy } from "./instanceIamPolicy";
import { Runtime } from "./runtime";
import { RuntimeIamPolicy } from "./runtimeIamPolicy";
import { Schedule } from "./schedule";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "google-native:notebooks/v1:Environment":
                return new Environment(name, <any>undefined, { urn })
            case "google-native:notebooks/v1:Execution":
                return new Execution(name, <any>undefined, { urn })
            case "google-native:notebooks/v1:Instance":
                return new Instance(name, <any>undefined, { urn })
            case "google-native:notebooks/v1:InstanceIamPolicy":
                return new InstanceIamPolicy(name, <any>undefined, { urn })
            case "google-native:notebooks/v1:Runtime":
                return new Runtime(name, <any>undefined, { urn })
            case "google-native:notebooks/v1:RuntimeIamPolicy":
                return new RuntimeIamPolicy(name, <any>undefined, { urn })
            case "google-native:notebooks/v1:Schedule":
                return new Schedule(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("google-native", "notebooks/v1", _module)
