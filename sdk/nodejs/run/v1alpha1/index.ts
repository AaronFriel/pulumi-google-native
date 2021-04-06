// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./domainmapping";
export * from "./namespaceDomainmapping";
export * from "./namespaceJob";
export * from "./namespaceService";
export * from "./service";
export * from "./serviceIamPolicy";

// Import resources to register:
import { Domainmapping } from "./domainmapping";
import { NamespaceDomainmapping } from "./namespaceDomainmapping";
import { NamespaceJob } from "./namespaceJob";
import { NamespaceService } from "./namespaceService";
import { Service } from "./service";
import { ServiceIamPolicy } from "./serviceIamPolicy";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "google-cloud:run/v1alpha1:Domainmapping":
                return new Domainmapping(name, <any>undefined, { urn })
            case "google-cloud:run/v1alpha1:NamespaceDomainmapping":
                return new NamespaceDomainmapping(name, <any>undefined, { urn })
            case "google-cloud:run/v1alpha1:NamespaceJob":
                return new NamespaceJob(name, <any>undefined, { urn })
            case "google-cloud:run/v1alpha1:NamespaceService":
                return new NamespaceService(name, <any>undefined, { urn })
            case "google-cloud:run/v1alpha1:Service":
                return new Service(name, <any>undefined, { urn })
            case "google-cloud:run/v1alpha1:ServiceIamPolicy":
                return new ServiceIamPolicy(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("google-cloud", "run/v1alpha1", _module)