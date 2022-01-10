# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs

__all__ = [
    'GetVersionResult',
    'AwaitableGetVersionResult',
    'get_version',
    'get_version_output',
]

@pulumi.output_type
class GetVersionResult:
    def __init__(__self__, api_config=None, automatic_scaling=None, basic_scaling=None, beta_settings=None, build_env_variables=None, create_time=None, created_by=None, default_expiration=None, deployment=None, disk_usage_bytes=None, endpoints_api_service=None, entrypoint=None, env=None, env_variables=None, error_handlers=None, handlers=None, health_check=None, inbound_services=None, instance_class=None, libraries=None, liveness_check=None, manual_scaling=None, name=None, network=None, nobuild_files_regex=None, readiness_check=None, resources=None, runtime=None, runtime_api_version=None, runtime_channel=None, runtime_main_executable_path=None, service_account=None, serving_status=None, threadsafe=None, version_url=None, vm=None, vpc_access_connector=None):
        if api_config and not isinstance(api_config, dict):
            raise TypeError("Expected argument 'api_config' to be a dict")
        pulumi.set(__self__, "api_config", api_config)
        if automatic_scaling and not isinstance(automatic_scaling, dict):
            raise TypeError("Expected argument 'automatic_scaling' to be a dict")
        pulumi.set(__self__, "automatic_scaling", automatic_scaling)
        if basic_scaling and not isinstance(basic_scaling, dict):
            raise TypeError("Expected argument 'basic_scaling' to be a dict")
        pulumi.set(__self__, "basic_scaling", basic_scaling)
        if beta_settings and not isinstance(beta_settings, dict):
            raise TypeError("Expected argument 'beta_settings' to be a dict")
        pulumi.set(__self__, "beta_settings", beta_settings)
        if build_env_variables and not isinstance(build_env_variables, dict):
            raise TypeError("Expected argument 'build_env_variables' to be a dict")
        pulumi.set(__self__, "build_env_variables", build_env_variables)
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if created_by and not isinstance(created_by, str):
            raise TypeError("Expected argument 'created_by' to be a str")
        pulumi.set(__self__, "created_by", created_by)
        if default_expiration and not isinstance(default_expiration, str):
            raise TypeError("Expected argument 'default_expiration' to be a str")
        pulumi.set(__self__, "default_expiration", default_expiration)
        if deployment and not isinstance(deployment, dict):
            raise TypeError("Expected argument 'deployment' to be a dict")
        pulumi.set(__self__, "deployment", deployment)
        if disk_usage_bytes and not isinstance(disk_usage_bytes, str):
            raise TypeError("Expected argument 'disk_usage_bytes' to be a str")
        pulumi.set(__self__, "disk_usage_bytes", disk_usage_bytes)
        if endpoints_api_service and not isinstance(endpoints_api_service, dict):
            raise TypeError("Expected argument 'endpoints_api_service' to be a dict")
        pulumi.set(__self__, "endpoints_api_service", endpoints_api_service)
        if entrypoint and not isinstance(entrypoint, dict):
            raise TypeError("Expected argument 'entrypoint' to be a dict")
        pulumi.set(__self__, "entrypoint", entrypoint)
        if env and not isinstance(env, str):
            raise TypeError("Expected argument 'env' to be a str")
        pulumi.set(__self__, "env", env)
        if env_variables and not isinstance(env_variables, dict):
            raise TypeError("Expected argument 'env_variables' to be a dict")
        pulumi.set(__self__, "env_variables", env_variables)
        if error_handlers and not isinstance(error_handlers, list):
            raise TypeError("Expected argument 'error_handlers' to be a list")
        pulumi.set(__self__, "error_handlers", error_handlers)
        if handlers and not isinstance(handlers, list):
            raise TypeError("Expected argument 'handlers' to be a list")
        pulumi.set(__self__, "handlers", handlers)
        if health_check and not isinstance(health_check, dict):
            raise TypeError("Expected argument 'health_check' to be a dict")
        pulumi.set(__self__, "health_check", health_check)
        if inbound_services and not isinstance(inbound_services, list):
            raise TypeError("Expected argument 'inbound_services' to be a list")
        pulumi.set(__self__, "inbound_services", inbound_services)
        if instance_class and not isinstance(instance_class, str):
            raise TypeError("Expected argument 'instance_class' to be a str")
        pulumi.set(__self__, "instance_class", instance_class)
        if libraries and not isinstance(libraries, list):
            raise TypeError("Expected argument 'libraries' to be a list")
        pulumi.set(__self__, "libraries", libraries)
        if liveness_check and not isinstance(liveness_check, dict):
            raise TypeError("Expected argument 'liveness_check' to be a dict")
        pulumi.set(__self__, "liveness_check", liveness_check)
        if manual_scaling and not isinstance(manual_scaling, dict):
            raise TypeError("Expected argument 'manual_scaling' to be a dict")
        pulumi.set(__self__, "manual_scaling", manual_scaling)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if network and not isinstance(network, dict):
            raise TypeError("Expected argument 'network' to be a dict")
        pulumi.set(__self__, "network", network)
        if nobuild_files_regex and not isinstance(nobuild_files_regex, str):
            raise TypeError("Expected argument 'nobuild_files_regex' to be a str")
        pulumi.set(__self__, "nobuild_files_regex", nobuild_files_regex)
        if readiness_check and not isinstance(readiness_check, dict):
            raise TypeError("Expected argument 'readiness_check' to be a dict")
        pulumi.set(__self__, "readiness_check", readiness_check)
        if resources and not isinstance(resources, dict):
            raise TypeError("Expected argument 'resources' to be a dict")
        pulumi.set(__self__, "resources", resources)
        if runtime and not isinstance(runtime, str):
            raise TypeError("Expected argument 'runtime' to be a str")
        pulumi.set(__self__, "runtime", runtime)
        if runtime_api_version and not isinstance(runtime_api_version, str):
            raise TypeError("Expected argument 'runtime_api_version' to be a str")
        pulumi.set(__self__, "runtime_api_version", runtime_api_version)
        if runtime_channel and not isinstance(runtime_channel, str):
            raise TypeError("Expected argument 'runtime_channel' to be a str")
        pulumi.set(__self__, "runtime_channel", runtime_channel)
        if runtime_main_executable_path and not isinstance(runtime_main_executable_path, str):
            raise TypeError("Expected argument 'runtime_main_executable_path' to be a str")
        pulumi.set(__self__, "runtime_main_executable_path", runtime_main_executable_path)
        if service_account and not isinstance(service_account, str):
            raise TypeError("Expected argument 'service_account' to be a str")
        pulumi.set(__self__, "service_account", service_account)
        if serving_status and not isinstance(serving_status, str):
            raise TypeError("Expected argument 'serving_status' to be a str")
        pulumi.set(__self__, "serving_status", serving_status)
        if threadsafe and not isinstance(threadsafe, bool):
            raise TypeError("Expected argument 'threadsafe' to be a bool")
        pulumi.set(__self__, "threadsafe", threadsafe)
        if version_url and not isinstance(version_url, str):
            raise TypeError("Expected argument 'version_url' to be a str")
        pulumi.set(__self__, "version_url", version_url)
        if vm and not isinstance(vm, bool):
            raise TypeError("Expected argument 'vm' to be a bool")
        pulumi.set(__self__, "vm", vm)
        if vpc_access_connector and not isinstance(vpc_access_connector, dict):
            raise TypeError("Expected argument 'vpc_access_connector' to be a dict")
        pulumi.set(__self__, "vpc_access_connector", vpc_access_connector)

    @property
    @pulumi.getter(name="apiConfig")
    def api_config(self) -> 'outputs.ApiConfigHandlerResponse':
        """
        Serving configuration for Google Cloud Endpoints (https://cloud.google.com/appengine/docs/python/endpoints/).Only returned in GET requests if view=FULL is set.
        """
        return pulumi.get(self, "api_config")

    @property
    @pulumi.getter(name="automaticScaling")
    def automatic_scaling(self) -> 'outputs.AutomaticScalingResponse':
        """
        Automatic scaling is based on request rate, response latencies, and other application metrics. Instances are dynamically created and destroyed as needed in order to handle traffic.
        """
        return pulumi.get(self, "automatic_scaling")

    @property
    @pulumi.getter(name="basicScaling")
    def basic_scaling(self) -> 'outputs.BasicScalingResponse':
        """
        A service with basic scaling will create an instance when the application receives a request. The instance will be turned down when the app becomes idle. Basic scaling is ideal for work that is intermittent or driven by user activity.
        """
        return pulumi.get(self, "basic_scaling")

    @property
    @pulumi.getter(name="betaSettings")
    def beta_settings(self) -> Mapping[str, str]:
        """
        Metadata settings that are supplied to this version to enable beta runtime features.
        """
        return pulumi.get(self, "beta_settings")

    @property
    @pulumi.getter(name="buildEnvVariables")
    def build_env_variables(self) -> Mapping[str, str]:
        """
        Environment variables available to the build environment.Only returned in GET requests if view=FULL is set.
        """
        return pulumi.get(self, "build_env_variables")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        Time that this version was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="createdBy")
    def created_by(self) -> str:
        """
        Email address of the user who created this version.
        """
        return pulumi.get(self, "created_by")

    @property
    @pulumi.getter(name="defaultExpiration")
    def default_expiration(self) -> str:
        """
        Duration that static files should be cached by web proxies and browsers. Only applicable if the corresponding StaticFilesHandler (https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions#StaticFilesHandler) does not specify its own expiration time.Only returned in GET requests if view=FULL is set.
        """
        return pulumi.get(self, "default_expiration")

    @property
    @pulumi.getter
    def deployment(self) -> 'outputs.DeploymentResponse':
        """
        Code and application artifacts that make up this version.Only returned in GET requests if view=FULL is set.
        """
        return pulumi.get(self, "deployment")

    @property
    @pulumi.getter(name="diskUsageBytes")
    def disk_usage_bytes(self) -> str:
        """
        Total size in bytes of all the files that are included in this version and currently hosted on the App Engine disk.
        """
        return pulumi.get(self, "disk_usage_bytes")

    @property
    @pulumi.getter(name="endpointsApiService")
    def endpoints_api_service(self) -> 'outputs.EndpointsApiServiceResponse':
        """
        Cloud Endpoints configuration.If endpoints_api_service is set, the Cloud Endpoints Extensible Service Proxy will be provided to serve the API implemented by the app.
        """
        return pulumi.get(self, "endpoints_api_service")

    @property
    @pulumi.getter
    def entrypoint(self) -> 'outputs.EntrypointResponse':
        """
        The entrypoint for the application.
        """
        return pulumi.get(self, "entrypoint")

    @property
    @pulumi.getter
    def env(self) -> str:
        """
        App Engine execution environment for this version.Defaults to standard.
        """
        return pulumi.get(self, "env")

    @property
    @pulumi.getter(name="envVariables")
    def env_variables(self) -> Mapping[str, str]:
        """
        Environment variables available to the application.Only returned in GET requests if view=FULL is set.
        """
        return pulumi.get(self, "env_variables")

    @property
    @pulumi.getter(name="errorHandlers")
    def error_handlers(self) -> Sequence['outputs.ErrorHandlerResponse']:
        """
        Custom static error pages. Limited to 10KB per page.Only returned in GET requests if view=FULL is set.
        """
        return pulumi.get(self, "error_handlers")

    @property
    @pulumi.getter
    def handlers(self) -> Sequence['outputs.UrlMapResponse']:
        """
        An ordered list of URL-matching patterns that should be applied to incoming requests. The first matching URL handles the request and other request handlers are not attempted.Only returned in GET requests if view=FULL is set.
        """
        return pulumi.get(self, "handlers")

    @property
    @pulumi.getter(name="healthCheck")
    def health_check(self) -> 'outputs.HealthCheckResponse':
        """
        Configures health checking for instances. Unhealthy instances are stopped and replaced with new instances. Only applicable in the App Engine flexible environment.Only returned in GET requests if view=FULL is set.
        """
        return pulumi.get(self, "health_check")

    @property
    @pulumi.getter(name="inboundServices")
    def inbound_services(self) -> Sequence[str]:
        """
        Before an application can receive email or XMPP messages, the application must be configured to enable the service.
        """
        return pulumi.get(self, "inbound_services")

    @property
    @pulumi.getter(name="instanceClass")
    def instance_class(self) -> str:
        """
        Instance class that is used to run this version. Valid values are: AutomaticScaling: F1, F2, F4, F4_1G ManualScaling or BasicScaling: B1, B2, B4, B8, B4_1GDefaults to F1 for AutomaticScaling and B1 for ManualScaling or BasicScaling.
        """
        return pulumi.get(self, "instance_class")

    @property
    @pulumi.getter
    def libraries(self) -> Sequence['outputs.LibraryResponse']:
        """
        Configuration for third-party Python runtime libraries that are required by the application.Only returned in GET requests if view=FULL is set.
        """
        return pulumi.get(self, "libraries")

    @property
    @pulumi.getter(name="livenessCheck")
    def liveness_check(self) -> 'outputs.LivenessCheckResponse':
        """
        Configures liveness health checking for instances. Unhealthy instances are stopped and replaced with new instancesOnly returned in GET requests if view=FULL is set.
        """
        return pulumi.get(self, "liveness_check")

    @property
    @pulumi.getter(name="manualScaling")
    def manual_scaling(self) -> 'outputs.ManualScalingResponse':
        """
        A service with manual scaling runs continuously, allowing you to perform complex initialization and rely on the state of its memory over time. Manually scaled versions are sometimes referred to as "backends".
        """
        return pulumi.get(self, "manual_scaling")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Full path to the Version resource in the API. Example: apps/myapp/services/default/versions/v1.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def network(self) -> 'outputs.NetworkResponse':
        """
        Extra network settings. Only applicable in the App Engine flexible environment.
        """
        return pulumi.get(self, "network")

    @property
    @pulumi.getter(name="nobuildFilesRegex")
    def nobuild_files_regex(self) -> str:
        """
        Files that match this pattern will not be built into this version. Only applicable for Go runtimes.Only returned in GET requests if view=FULL is set.
        """
        return pulumi.get(self, "nobuild_files_regex")

    @property
    @pulumi.getter(name="readinessCheck")
    def readiness_check(self) -> 'outputs.ReadinessCheckResponse':
        """
        Configures readiness health checking for instances. Unhealthy instances are not put into the backend traffic rotation.Only returned in GET requests if view=FULL is set.
        """
        return pulumi.get(self, "readiness_check")

    @property
    @pulumi.getter
    def resources(self) -> 'outputs.ResourcesResponse':
        """
        Machine resources for this version. Only applicable in the App Engine flexible environment.
        """
        return pulumi.get(self, "resources")

    @property
    @pulumi.getter
    def runtime(self) -> str:
        """
        Desired runtime. Example: python27.
        """
        return pulumi.get(self, "runtime")

    @property
    @pulumi.getter(name="runtimeApiVersion")
    def runtime_api_version(self) -> str:
        """
        The version of the API in the given runtime environment. Please see the app.yaml reference for valid values at https://cloud.google.com/appengine/docs/standard//config/appref
        """
        return pulumi.get(self, "runtime_api_version")

    @property
    @pulumi.getter(name="runtimeChannel")
    def runtime_channel(self) -> str:
        """
        The channel of the runtime to use. Only available for some runtimes. Defaults to the default channel.
        """
        return pulumi.get(self, "runtime_channel")

    @property
    @pulumi.getter(name="runtimeMainExecutablePath")
    def runtime_main_executable_path(self) -> str:
        """
        The path or name of the app's main executable.
        """
        return pulumi.get(self, "runtime_main_executable_path")

    @property
    @pulumi.getter(name="serviceAccount")
    def service_account(self) -> str:
        """
        The identity that the deployed version will run as. Admin API will use the App Engine Appspot service account as default if this field is neither provided in app.yaml file nor through CLI flag.
        """
        return pulumi.get(self, "service_account")

    @property
    @pulumi.getter(name="servingStatus")
    def serving_status(self) -> str:
        """
        Current serving status of this version. Only the versions with a SERVING status create instances and can be billed.SERVING_STATUS_UNSPECIFIED is an invalid value. Defaults to SERVING.
        """
        return pulumi.get(self, "serving_status")

    @property
    @pulumi.getter
    def threadsafe(self) -> bool:
        """
        Whether multiple requests can be dispatched to this version at once.
        """
        return pulumi.get(self, "threadsafe")

    @property
    @pulumi.getter(name="versionUrl")
    def version_url(self) -> str:
        """
        Serving URL for this version. Example: "https://myversion-dot-myservice-dot-myapp.appspot.com"
        """
        return pulumi.get(self, "version_url")

    @property
    @pulumi.getter
    def vm(self) -> bool:
        """
        Whether to deploy this version in a container on a virtual machine.
        """
        return pulumi.get(self, "vm")

    @property
    @pulumi.getter(name="vpcAccessConnector")
    def vpc_access_connector(self) -> 'outputs.VpcAccessConnectorResponse':
        """
        Enables VPC connectivity for standard apps.
        """
        return pulumi.get(self, "vpc_access_connector")


class AwaitableGetVersionResult(GetVersionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVersionResult(
            api_config=self.api_config,
            automatic_scaling=self.automatic_scaling,
            basic_scaling=self.basic_scaling,
            beta_settings=self.beta_settings,
            build_env_variables=self.build_env_variables,
            create_time=self.create_time,
            created_by=self.created_by,
            default_expiration=self.default_expiration,
            deployment=self.deployment,
            disk_usage_bytes=self.disk_usage_bytes,
            endpoints_api_service=self.endpoints_api_service,
            entrypoint=self.entrypoint,
            env=self.env,
            env_variables=self.env_variables,
            error_handlers=self.error_handlers,
            handlers=self.handlers,
            health_check=self.health_check,
            inbound_services=self.inbound_services,
            instance_class=self.instance_class,
            libraries=self.libraries,
            liveness_check=self.liveness_check,
            manual_scaling=self.manual_scaling,
            name=self.name,
            network=self.network,
            nobuild_files_regex=self.nobuild_files_regex,
            readiness_check=self.readiness_check,
            resources=self.resources,
            runtime=self.runtime,
            runtime_api_version=self.runtime_api_version,
            runtime_channel=self.runtime_channel,
            runtime_main_executable_path=self.runtime_main_executable_path,
            service_account=self.service_account,
            serving_status=self.serving_status,
            threadsafe=self.threadsafe,
            version_url=self.version_url,
            vm=self.vm,
            vpc_access_connector=self.vpc_access_connector)


def get_version(app_id: Optional[str] = None,
                service_id: Optional[str] = None,
                version_id: Optional[str] = None,
                view: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVersionResult:
    """
    Gets the specified Version resource. By default, only a BASIC_VIEW will be returned. Specify the FULL_VIEW parameter to get the full resource.
    """
    __args__ = dict()
    __args__['appId'] = app_id
    __args__['serviceId'] = service_id
    __args__['versionId'] = version_id
    __args__['view'] = view
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:appengine/v1:getVersion', __args__, opts=opts, typ=GetVersionResult).value

    return AwaitableGetVersionResult(
        api_config=__ret__.api_config,
        automatic_scaling=__ret__.automatic_scaling,
        basic_scaling=__ret__.basic_scaling,
        beta_settings=__ret__.beta_settings,
        build_env_variables=__ret__.build_env_variables,
        create_time=__ret__.create_time,
        created_by=__ret__.created_by,
        default_expiration=__ret__.default_expiration,
        deployment=__ret__.deployment,
        disk_usage_bytes=__ret__.disk_usage_bytes,
        endpoints_api_service=__ret__.endpoints_api_service,
        entrypoint=__ret__.entrypoint,
        env=__ret__.env,
        env_variables=__ret__.env_variables,
        error_handlers=__ret__.error_handlers,
        handlers=__ret__.handlers,
        health_check=__ret__.health_check,
        inbound_services=__ret__.inbound_services,
        instance_class=__ret__.instance_class,
        libraries=__ret__.libraries,
        liveness_check=__ret__.liveness_check,
        manual_scaling=__ret__.manual_scaling,
        name=__ret__.name,
        network=__ret__.network,
        nobuild_files_regex=__ret__.nobuild_files_regex,
        readiness_check=__ret__.readiness_check,
        resources=__ret__.resources,
        runtime=__ret__.runtime,
        runtime_api_version=__ret__.runtime_api_version,
        runtime_channel=__ret__.runtime_channel,
        runtime_main_executable_path=__ret__.runtime_main_executable_path,
        service_account=__ret__.service_account,
        serving_status=__ret__.serving_status,
        threadsafe=__ret__.threadsafe,
        version_url=__ret__.version_url,
        vm=__ret__.vm,
        vpc_access_connector=__ret__.vpc_access_connector)


@_utilities.lift_output_func(get_version)
def get_version_output(app_id: Optional[pulumi.Input[str]] = None,
                       service_id: Optional[pulumi.Input[str]] = None,
                       version_id: Optional[pulumi.Input[str]] = None,
                       view: Optional[pulumi.Input[Optional[str]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVersionResult]:
    """
    Gets the specified Version resource. By default, only a BASIC_VIEW will be returned. Specify the FULL_VIEW parameter to get the full resource.
    """
    ...