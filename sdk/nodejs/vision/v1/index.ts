// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./product";
export * from "./productReferenceImage";
export * from "./productSet";

// Import resources to register:
import { Product } from "./product";
import { ProductReferenceImage } from "./productReferenceImage";
import { ProductSet } from "./productSet";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "google-cloud:vision/v1:Product":
                return new Product(name, <any>undefined, { urn })
            case "google-cloud:vision/v1:ProductReferenceImage":
                return new ProductReferenceImage(name, <any>undefined, { urn })
            case "google-cloud:vision/v1:ProductSet":
                return new ProductSet(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("google-cloud", "vision/v1", _module)