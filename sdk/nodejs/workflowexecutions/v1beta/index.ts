// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./execution";
export * from "./getExecution";

// Export enums:
export * from "../../types/enums/workflowexecutions/v1beta";

// Import resources to register:
import { Execution } from "./execution";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "google-native:workflowexecutions/v1beta:Execution":
                return new Execution(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("google-native", "workflowexecutions/v1beta", _module)
