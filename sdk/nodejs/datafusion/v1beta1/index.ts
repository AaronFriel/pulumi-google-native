// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./getInstance";
export * from "./getInstanceIamPolicy";
export * from "./getInstanceNamespaceIamPolicy";
export * from "./instance";
export * from "./instanceIamPolicy";
export * from "./instanceNamespaceIamPolicy";

// Export enums:
export * from "../../types/enums/datafusion/v1beta1";

// Import resources to register:
import { Instance } from "./instance";
import { InstanceIamPolicy } from "./instanceIamPolicy";
import { InstanceNamespaceIamPolicy } from "./instanceNamespaceIamPolicy";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "google-native:datafusion/v1beta1:Instance":
                return new Instance(name, <any>undefined, { urn })
            case "google-native:datafusion/v1beta1:InstanceIamPolicy":
                return new InstanceIamPolicy(name, <any>undefined, { urn })
            case "google-native:datafusion/v1beta1:InstanceNamespaceIamPolicy":
                return new InstanceNamespaceIamPolicy(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("google-native", "datafusion/v1beta1", _module)
