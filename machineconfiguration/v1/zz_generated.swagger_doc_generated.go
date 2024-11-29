package v1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE
var map_CertExpiry = map[string]string{
	"":        "ceryExpiry contains the bundle name and the expiry date",
	"bundle":  "bundle is the name of the bundle in which the subject certificate resides",
	"subject": "subject is the subject of the certificate",
	"expiry":  "expiry is the date after which the certificate will no longer be valid",
}

func (CertExpiry) SwaggerDoc() map[string]string {
	return map_CertExpiry
}

var map_ContainerRuntimeConfig = map[string]string{
	"":       "ContainerRuntimeConfig describes a customized Container Runtime configuration.\n\nCompatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
	"spec":   "spec contains the desired container runtime configuration.",
	"status": "status contains observed information about the container runtime configuration.",
}

func (ContainerRuntimeConfig) SwaggerDoc() map[string]string {
	return map_ContainerRuntimeConfig
}

var map_ContainerRuntimeConfigCondition = map[string]string{
	"":                   "ContainerRuntimeConfigCondition defines the state of the ContainerRuntimeConfig",
	"type":               "type specifies the state of the operator's reconciliation functionality.",
	"status":             "status of the condition, one of True, False, Unknown.",
	"lastTransitionTime": "lastTransitionTime is the time of the last update to the current status object.",
	"reason":             "reason is the reason for the condition's last transition.  Reasons are PascalCase",
	"message":            "message provides additional information about the current condition. This is only to be consumed by humans.",
}

func (ContainerRuntimeConfigCondition) SwaggerDoc() map[string]string {
	return map_ContainerRuntimeConfigCondition
}

var map_ContainerRuntimeConfigList = map[string]string{
	"": "ContainerRuntimeConfigList is a list of ContainerRuntimeConfig resources\n\nCompatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
}

func (ContainerRuntimeConfigList) SwaggerDoc() map[string]string {
	return map_ContainerRuntimeConfigList
}

var map_ContainerRuntimeConfigSpec = map[string]string{
	"":                          "ContainerRuntimeConfigSpec defines the desired state of ContainerRuntimeConfig",
	"machineConfigPoolSelector": "machineConfigPoolSelector selects which pools the ContainerRuntimeConfig shoud apply to. A nil selector will result in no pools being selected.",
	"containerRuntimeConfig":    "containerRuntimeConfig defines the tuneables of the container runtime.",
}

func (ContainerRuntimeConfigSpec) SwaggerDoc() map[string]string {
	return map_ContainerRuntimeConfigSpec
}

var map_ContainerRuntimeConfigStatus = map[string]string{
	"":                   "ContainerRuntimeConfigStatus defines the observed state of a ContainerRuntimeConfig",
	"observedGeneration": "observedGeneration represents the generation observed by the controller.",
	"conditions":         "conditions represents the latest available observations of current state.",
}

func (ContainerRuntimeConfigStatus) SwaggerDoc() map[string]string {
	return map_ContainerRuntimeConfigStatus
}

var map_ContainerRuntimeConfiguration = map[string]string{
	"":               "ContainerRuntimeConfiguration defines the tuneables of the container runtime",
	"pidsLimit":      "pidsLimit specifies the maximum number of processes allowed in a container",
	"logLevel":       "logLevel specifies the verbosity of the logs based on the level it is set to. Options are fatal, panic, error, warn, info, and debug.",
	"logSizeMax":     "logSizeMax specifies the Maximum size allowed for the container log file. Negative numbers indicate that no size limit is imposed. If it is positive, it must be >= 8192 to match/exceed conmon's read buffer.",
	"overlaySize":    "overlaySize specifies the maximum size of a container image. This flag can be used to set quota on the size of container images. (default: 10GB)",
	"defaultRuntime": "defaultRuntime is the name of the OCI runtime to be used as the default.",
}

func (ContainerRuntimeConfiguration) SwaggerDoc() map[string]string {
	return map_ContainerRuntimeConfiguration
}

var map_ControllerCertificate = map[string]string{
	"":           "ControllerCertificate contains info about a specific cert.",
	"subject":    "subject is the cert subject",
	"signer":     "signer is the  cert Issuer",
	"notBefore":  "notBefore is the lower boundary for validity",
	"notAfter":   "notAfter is the upper boundary for validity",
	"bundleFile": "bundleFile is the larger bundle a cert comes from",
}

func (ControllerCertificate) SwaggerDoc() map[string]string {
	return map_ControllerCertificate
}

var map_ControllerConfig = map[string]string{
	"":       "ControllerConfig describes configuration for MachineConfigController. This is currently only used to drive the MachineConfig objects generated by the TemplateController.\n\nCompatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
	"spec":   "spec contains the desired controller config configuration.",
	"status": "status contains observed information about the controller config.",
}

func (ControllerConfig) SwaggerDoc() map[string]string {
	return map_ControllerConfig
}

var map_ControllerConfigList = map[string]string{
	"": "ControllerConfigList is a list of ControllerConfig resources\n\nCompatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
}

func (ControllerConfigList) SwaggerDoc() map[string]string {
	return map_ControllerConfigList
}

var map_ControllerConfigSpec = map[string]string{
	"":                               "ControllerConfigSpec is the spec for ControllerConfig resource.",
	"clusterDNSIP":                   "clusterDNSIP is the cluster DNS IP address",
	"cloudProviderConfig":            "cloudProviderConfig is the configuration for the given cloud provider",
	"platform":                       "platform is deprecated, use Infra.Status.PlatformStatus.Type instead",
	"etcdDiscoveryDomain":            "etcdDiscoveryDomain is deprecated, use Infra.Status.EtcdDiscoveryDomain instead",
	"kubeAPIServerServingCAData":     "kubeAPIServerServingCAData managed Kubelet to API Server Cert... Rotated automatically",
	"rootCAData":                     "rootCAData specifies the root CA data",
	"cloudProviderCAData":            "cloudProviderCAData specifies the cloud provider CA data",
	"additionalTrustBundle":          "additionalTrustBundle is a certificate bundle that will be added to the nodes trusted certificate store.",
	"imageRegistryBundleUserData":    "imageRegistryBundleUserData is Image Registry Data provided by the user",
	"imageRegistryBundleData":        "imageRegistryBundleData is the ImageRegistryData",
	"pullSecret":                     "pullSecret is the default pull secret that needs to be installed on all machines.",
	"internalRegistryPullSecret":     "internalRegistryPullSecret is the pull secret for the internal registry, used by rpm-ostree to pull images from the internal registry if present",
	"images":                         "images is map of images that are used by the controller to render templates under ./templates/",
	"baseOSContainerImage":           "baseOSContainerImage is the new-format container image for operating system updates.",
	"baseOSExtensionsContainerImage": "baseOSExtensionsContainerImage is the matching extensions container for the new-format container",
	"osImageURL":                     "osImageURL is the old-format container image that contains the OS update payload.",
	"releaseImage":                   "releaseImage is the image used when installing the cluster",
	"proxy":                          "proxy holds the current proxy configuration for the nodes",
	"infra":                          "infra holds the infrastructure details",
	"dns":                            "dns holds the cluster dns details",
	"ipFamilies":                     "ipFamilies indicates the IP families in use by the cluster network",
	"networkType":                    "networkType holds the type of network the cluster is using XXX: this is temporary and will be dropped as soon as possible in favor of a better support to start network related services the proper way. Nobody is also changing this once the cluster is up and running the first time, so, disallow regeneration if this changes.",
	"network":                        "network contains additional network related information",
}

func (ControllerConfigSpec) SwaggerDoc() map[string]string {
	return map_ControllerConfigSpec
}

var map_ControllerConfigStatus = map[string]string{
	"":                       "ControllerConfigStatus is the status for ControllerConfig",
	"observedGeneration":     "observedGeneration represents the generation observed by the controller.",
	"conditions":             "conditions represents the latest available observations of current state.",
	"controllerCertificates": "controllerCertificates represents the latest available observations of the automatically rotating certificates in the MCO.",
}

func (ControllerConfigStatus) SwaggerDoc() map[string]string {
	return map_ControllerConfigStatus
}

var map_ControllerConfigStatusCondition = map[string]string{
	"":                   "ControllerConfigStatusCondition contains condition information for ControllerConfigStatus",
	"type":               "type specifies the state of the operator's reconciliation functionality.",
	"status":             "status of the condition, one of True, False, Unknown.",
	"lastTransitionTime": "lastTransitionTime is the time of the last update to the current status object.",
	"reason":             "reason is the reason for the condition's last transition.  Reasons are PascalCase",
	"message":            "message provides additional information about the current condition. This is only to be consumed by humans.",
}

func (ControllerConfigStatusCondition) SwaggerDoc() map[string]string {
	return map_ControllerConfigStatusCondition
}

var map_ImageRegistryBundle = map[string]string{
	"":     "ImageRegistryBundle contains information for writing image registry certificates",
	"file": "file holds the name of the file where the bundle will be written to disk",
	"data": "data holds the contents of the bundle that will be written to the file location",
}

func (ImageRegistryBundle) SwaggerDoc() map[string]string {
	return map_ImageRegistryBundle
}

var map_KubeletConfig = map[string]string{
	"":       "KubeletConfig describes a customized Kubelet configuration.\n\nCompatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
	"spec":   "spec contains the desired kubelet configuration.",
	"status": "status contains observed information about the kubelet configuration.",
}

func (KubeletConfig) SwaggerDoc() map[string]string {
	return map_KubeletConfig
}

var map_KubeletConfigCondition = map[string]string{
	"":                   "KubeletConfigCondition defines the state of the KubeletConfig",
	"type":               "type specifies the state of the operator's reconciliation functionality.",
	"status":             "status of the condition, one of True, False, Unknown.",
	"lastTransitionTime": "lastTransitionTime is the time of the last update to the current status object.",
	"reason":             "reason is the reason for the condition's last transition.  Reasons are PascalCase",
	"message":            "message provides additional information about the current condition. This is only to be consumed by humans.",
}

func (KubeletConfigCondition) SwaggerDoc() map[string]string {
	return map_KubeletConfigCondition
}

var map_KubeletConfigList = map[string]string{
	"": "KubeletConfigList is a list of KubeletConfig resources\n\nCompatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
}

func (KubeletConfigList) SwaggerDoc() map[string]string {
	return map_KubeletConfigList
}

var map_KubeletConfigSpec = map[string]string{
	"":                          "KubeletConfigSpec defines the desired state of KubeletConfig",
	"machineConfigPoolSelector": "machineConfigPoolSelector selects which pools the KubeletConfig shoud apply to. A nil selector will result in no pools being selected.",
	"kubeletConfig":             "kubeletConfig fields are defined in kubernetes upstream. Please refer to the types defined in the version/commit used by OpenShift of the upstream kubernetes. It's important to note that, since the fields of the kubelet configuration are directly fetched from upstream the validation of those values is handled directly by the kubelet. Please refer to the upstream version of the relevant kubernetes for the valid values of these fields. Invalid values of the kubelet configuration fields may render cluster nodes unusable.",
	"tlsSecurityProfile":        "If unset, the default is based on the apiservers.config.openshift.io/cluster resource. Note that only Old and Intermediate profiles are currently supported, and the maximum available minTLSVersion is VersionTLS12.",
}

func (KubeletConfigSpec) SwaggerDoc() map[string]string {
	return map_KubeletConfigSpec
}

var map_KubeletConfigStatus = map[string]string{
	"":                   "KubeletConfigStatus defines the observed state of a KubeletConfig",
	"observedGeneration": "observedGeneration represents the generation observed by the controller.",
	"conditions":         "conditions represents the latest available observations of current state.",
}

func (KubeletConfigStatus) SwaggerDoc() map[string]string {
	return map_KubeletConfigStatus
}

var map_MachineConfig = map[string]string{
	"": "MachineConfig defines the configuration for a machine\n\nCompatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
}

func (MachineConfig) SwaggerDoc() map[string]string {
	return map_MachineConfig
}

var map_MachineConfigList = map[string]string{
	"": "MachineConfigList is a list of MachineConfig resources\n\nCompatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
}

func (MachineConfigList) SwaggerDoc() map[string]string {
	return map_MachineConfigList
}

var map_MachineConfigPool = map[string]string{
	"":       "MachineConfigPool describes a pool of MachineConfigs.\n\nCompatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
	"spec":   "spec contains the desired machine config pool configuration.",
	"status": "status contains observed information about the machine config pool.",
}

func (MachineConfigPool) SwaggerDoc() map[string]string {
	return map_MachineConfigPool
}

var map_MachineConfigPoolCondition = map[string]string{
	"":                   "MachineConfigPoolCondition contains condition information for an MachineConfigPool.",
	"type":               "type of the condition, currently ('Done', 'Updating', 'Failed').",
	"status":             "status of the condition, one of ('True', 'False', 'Unknown').",
	"lastTransitionTime": "lastTransitionTime is the timestamp corresponding to the last status change of this condition.",
	"reason":             "reason is a brief machine readable explanation for the condition's last transition.",
	"message":            "message is a human readable description of the details of the last transition, complementing reason.",
}

func (MachineConfigPoolCondition) SwaggerDoc() map[string]string {
	return map_MachineConfigPoolCondition
}

var map_MachineConfigPoolList = map[string]string{
	"": "MachineConfigPoolList is a list of MachineConfigPool resources\n\nCompatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
}

func (MachineConfigPoolList) SwaggerDoc() map[string]string {
	return map_MachineConfigPoolList
}

var map_MachineConfigPoolSpec = map[string]string{
	"":                      "MachineConfigPoolSpec is the spec for MachineConfigPool resource.",
	"machineConfigSelector": "machineConfigSelector specifies a label selector for MachineConfigs. Refer https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ on how label and selectors work.",
	"nodeSelector":          "nodeSelector specifies a label selector for Machines",
	"paused":                "paused specifies whether or not changes to this machine config pool should be stopped. This includes generating new desiredMachineConfig and update of machines.",
	"maxUnavailable":        "maxUnavailable defines either an integer number or percentage of nodes in the pool that can go Unavailable during an update. This includes nodes Unavailable for any reason, including user initiated cordons, failing nodes, etc. The default value is 1.\n\nA value larger than 1 will mean multiple nodes going unavailable during the update, which may affect your workload stress on the remaining nodes. You cannot set this value to 0 to stop updates (it will default back to 1); to stop updates, use the 'paused' property instead. Drain will respect Pod Disruption Budgets (PDBs) such as etcd quorum guards, even if maxUnavailable is greater than one.",
	"configuration":         "The targeted MachineConfig object for the machine config pool.",
	"pinnedImageSets":       "pinnedImageSets specifies a sequence of PinnedImageSetRef objects for the pool. Nodes within this pool will preload and pin images defined in the PinnedImageSet. Before pulling images the MachineConfigDaemon will ensure the total uncompressed size of all the images does not exceed available resources. If the total size of the images exceeds the available resources the controller will report a Degraded status to the MachineConfigPool and not attempt to pull any images. Also to help ensure the kubelet can mitigate storage risk, the pinned_image configuration and subsequent service reload will happen only after all of the images have been pulled for each set. Images from multiple PinnedImageSets are loaded and pinned sequentially as listed. Duplicate and existing images will be skipped.\n\nAny failure to prefetch or pin images will result in a Degraded pool. Resolving these failures is the responsibility of the user. The admin should be proactive in ensuring adequate storage and proper image authentication exists in advance.",
}

func (MachineConfigPoolSpec) SwaggerDoc() map[string]string {
	return map_MachineConfigPoolSpec
}

var map_MachineConfigPoolStatus = map[string]string{
	"":                        "MachineConfigPoolStatus is the status for MachineConfigPool resource.",
	"observedGeneration":      "observedGeneration represents the generation observed by the controller.",
	"configuration":           "configuration represents the current MachineConfig object for the machine config pool.",
	"machineCount":            "machineCount represents the total number of machines in the machine config pool.",
	"updatedMachineCount":     "updatedMachineCount represents the total number of machines targeted by the pool that have the CurrentMachineConfig as their config.",
	"readyMachineCount":       "readyMachineCount represents the total number of ready machines targeted by the pool.",
	"unavailableMachineCount": "unavailableMachineCount represents the total number of unavailable (non-ready) machines targeted by the pool. A node is marked unavailable if it is in updating state or NodeReady condition is false.",
	"degradedMachineCount":    "degradedMachineCount represents the total number of machines marked degraded (or unreconcilable). A node is marked degraded if applying a configuration failed..",
	"conditions":              "conditions represents the latest available observations of current state.",
	"certExpirys":             "certExpirys keeps track of important certificate expiration data",
	"poolSynchronizersStatus": "poolSynchronizersStatus is the status of the machines managed by the pool synchronizers.",
}

func (MachineConfigPoolStatus) SwaggerDoc() map[string]string {
	return map_MachineConfigPoolStatus
}

var map_MachineConfigPoolStatusConfiguration = map[string]string{
	"":       "MachineConfigPoolStatusConfiguration stores the current configuration for the pool, and optionally also stores the list of MachineConfig objects used to generate the configuration.",
	"source": "source is the list of MachineConfig objects that were used to generate the single MachineConfig object specified in `content`.",
}

func (MachineConfigPoolStatusConfiguration) SwaggerDoc() map[string]string {
	return map_MachineConfigPoolStatusConfiguration
}

var map_MachineConfigSpec = map[string]string{
	"":                               "MachineConfigSpec is the spec for MachineConfig",
	"osImageURL":                     "osImageURL specifies the remote location that will be used to fetch the OS.",
	"baseOSExtensionsContainerImage": "baseOSExtensionsContainerImage specifies the remote location that will be used to fetch the extensions container matching a new-format OS image",
	"config":                         "config is a Ignition Config object.",
	"kernelArguments":                "kernelArguments contains a list of kernel arguments to be added",
	"extensions":                     "extensions contains a list of additional features that can be enabled on host",
	"fips":                           "fips controls FIPS mode",
	"kernelType":                     "kernelType contains which kernel we want to be running like default (traditional), realtime, 64k-pages (aarch64 only).",
}

func (MachineConfigSpec) SwaggerDoc() map[string]string {
	return map_MachineConfigSpec
}

var map_NetworkInfo = map[string]string{
	"":             "Network contains network related configuration",
	"mtuMigration": "mtuMigration contains the MTU migration configuration.",
}

func (NetworkInfo) SwaggerDoc() map[string]string {
	return map_NetworkInfo
}

var map_PinnedImageSetRef = map[string]string{
	"name": "name is a reference to the name of a PinnedImageSet.  Must adhere to RFC-1123 (https://tools.ietf.org/html/rfc1123). Made up of one of more period-separated (.) segments, where each segment consists of alphanumeric characters and hyphens (-), must begin and end with an alphanumeric character, and is at most 63 characters in length. The total length of the name must not exceed 253 characters.",
}

func (PinnedImageSetRef) SwaggerDoc() map[string]string {
	return map_PinnedImageSetRef
}

var map_PoolSynchronizerStatus = map[string]string{
	"poolSynchronizerType":    "poolSynchronizerType describes the type of the pool synchronizer.",
	"machineCount":            "machineCount is the number of machines that are managed by the node synchronizer.",
	"updatedMachineCount":     "updatedMachineCount is the number of machines that have been updated by the node synchronizer.",
	"readyMachineCount":       "readyMachineCount is the number of machines managed by the node synchronizer that are in a ready state.",
	"availableMachineCount":   "availableMachineCount is the number of machines managed by the node synchronizer which are available.",
	"unavailableMachineCount": "unavailableMachineCount is the number of machines managed by the node synchronizer but are unavailable.",
	"observedGeneration":      "observedGeneration is the last generation change that has been applied.",
}

func (PoolSynchronizerStatus) SwaggerDoc() map[string]string {
	return map_PoolSynchronizerStatus
}

var map_MachineOSBuild = map[string]string{
	"":       "MachineOSBuild describes a build process managed and deployed by the MCO Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
	"spec":   "spec describes the configuration of the machine os build",
	"status": "status describes the lst observed state of this machine os build",
}

func (MachineOSBuild) SwaggerDoc() map[string]string {
	return map_MachineOSBuild
}

var map_MachineOSBuildList = map[string]string{
	"":      "MachineOSBuildList describes all of the Builds on the system\n\nCompatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
	"items": "items contains a collection of MachineOSBuild resources.",
}

func (MachineOSBuildList) SwaggerDoc() map[string]string {
	return map_MachineOSBuildList
}

var map_MachineOSBuildSpec = map[string]string{
	"":                      "MachineOSBuildSpec describes information about a build process primarily populated from a MachineOSConfig object.",
	"desiredConfig":         "desiredConfig is the desired config we want to build an image for.",
	"machineOSConfig":       "machineOSConfig is the config object which the build is based off of",
	"renderedImagePushspec": "renderedImagePushspec is set from the MachineOSConfig The format of the image pullspec is: host[:port][/namespace]/name:<tag> or svc_name.namespace.svc[:port]/repository/name:<tag>",
}

func (MachineOSBuildSpec) SwaggerDoc() map[string]string {
	return map_MachineOSBuildSpec
}

var map_MachineOSBuildStatus = map[string]string{
	"":                   "MachineOSBuildStatus describes the state of a build and other helpful information.",
	"conditions":         "conditions are state related conditions for the build. Valid types are: Prepared, Building, Failed, Interrupted, and Succeeded. Once a Build is marked as Failed or Interrupted, no future conditions can be set.",
	"builderReference":   "BuilderReference describes which ImageBuilder backend to use for this build",
	"relatedObjects":     "relatedObjects is a list of references to ephemeral objects such as ConfigMaps or Secrets that are meant to be consumed while the build process runs. After a successful build or when this MachineOSBuild is deleted, these ephemeral objects should be deleted. However, in the event of a failed build, the objects will not be deleted to allow for inspection and debugging of the failed build process.",
	"buildStart":         "buildStart describes when the build started.",
	"buildEnd":           "buildEnd describes when the build ended. When omitted the build has either not been started, or is in progress. It will be populated once the build completes, fails or is interrupted.",
	"finalImagePushspec": "finalImagePushSpec describes the fully qualified pushspec produced by this build that the final image can be. Must end with a valid '@sha256:<digest>' suffix, where '<digest>' is 64 hexadecimal characters long",
}

func (MachineOSBuildStatus) SwaggerDoc() map[string]string {
	return map_MachineOSBuildStatus
}

var map_MachineOSBuilderReference = map[string]string{
	"":                 "MachineOSBuilderReference describes which ImageBuilder backend to use for this build",
	"imageBuilderType": "ImageBuilderType describes the image builder set in the MachineOSConfig, which in turn describes the builder that the cluster will attempt the build with. Currently only JobImageBuilder is supported, which will spin up a custom pod builder that uses buildah to build the specified image.",
	"ImageBuilderRef":  "ImageBuilderRef is a reference to the object that is managing the image build For example, if the imageBuilderType is JobImageBuilder, this will be a reference to the Job object managing the build",
}

func (MachineOSBuilderReference) SwaggerDoc() map[string]string {
	return map_MachineOSBuilderReference
}

var map_MachineOSConfigReference = map[string]string{
	"":     "MachineOSConfigReference refers to the MachineOSConfig this build is based off of",
	"name": "name of the MachineOSConfig. This value should consist of only lowercase alphanumeric characters and hyphens.",
}

func (MachineOSConfigReference) SwaggerDoc() map[string]string {
	return map_MachineOSConfigReference
}

var map_ObjectReference = map[string]string{
	"":          "ObjectReference contains enough information to let you inspect or modify the referred object.",
	"group":     "group of the referent. The name must contain only lowercase alphanumeric characters, '-' or '.' and start/end with an alphanumeric character Example: \"\", \"apps\", \"build.openshift.io\", etc.",
	"resource":  "resource of the referent. This value should consist of only lowercase alphanumeric characters and hyphens. Example: \"deployments\", \"deploymentconfigs\", \"pods\", etc.",
	"namespace": "namespace of the referent. This value should consist of only lowercase alphanumeric characters and hyphens.",
	"name":      "name of the referent. This value should consist of only lowercase alphanumeric characters and hyphens.",
}

func (ObjectReference) SwaggerDoc() map[string]string {
	return map_ObjectReference
}

var map_RenderedMachineConfigReference = map[string]string{
	"":     "Refers to the name of a rendered MachineConfig (e.g., \"rendered-worker-ec40d2965ff81bce7cd7a7e82a680739\", etc.): the build targets this MachineConfig, this is often used to tell us whether we need an update.",
	"name": "name is the name of the rendered MachineConfig object. The name must contain only lowercase alphanumeric characters, '-' or '.' and start/end with an alphanumeric character",
}

func (RenderedMachineConfigReference) SwaggerDoc() map[string]string {
	return map_RenderedMachineConfigReference
}

var map_BuildInputs = map[string]string{
	"":                              "BuildInputs holds all of the information needed to trigger a build",
	"baseOSExtensionsImagePullspec": "baseOSExtensionsImagePullspec is the base Extensions image used in the build process The MachineOSConfig object will use the in cluster image registry configuration. If you wish to use a mirror or any other settings specific to registries.conf, please specify those in the cluster wide registries.conf. The format of the image pullspec is: host[:port][/namespace]/name@sha256:<digest> The digest must be 64 characters long, and consist only of lowercase hexadecimal characters, a-f and 0-9.",
	"baseOSImagePullspec":           "baseOSImagePullspec is the base OSImage we use to build our custom image. The MachineOSConfig object will use the in cluster image registry configuration. If you wish to use a mirror or any other settings specific to registries.conf, please specify those in the cluster wide registries.conf. The format of the image pullspec is: host[:port][/namespace]/name@sha256:<digest> The digest must be 64 characters long, and consist only of lowercase hexadecimal characters, a-f and 0-9.",
	"baseImagePullSecret":           "baseImagePullSecret is the secret used to pull the base image. Must live in the openshift-machine-config-operator namespace if provided. Defaults to using the cluster-wide pull secret if not specified. This is provided during install time of the cluster, and lives in the openshift-config namespace as a secret.",
	"imageBuilder":                  "machineOSImageBuilder describes which image builder will be used in each build triggered by this MachineOSConfig. Currently supported type(s): JobImageBuilder",
	"renderedImagePushSecret":       "renderedImagePushSecret is the secret used to connect to a user registry. The final image push and pull secrets should be separate and assume the principal of least privilege. The push secret with write privilege is only required to be present on the node hosting the MachineConfigController pod. The pull secret with read only privileges is required on all nodes. By separating the two secrets, the risk of write credentials becoming compromised is reduced.",
	"renderedImagePushSpec":         "renderedImagePushSpec describes the location of the final image. The MachineOSConfig object will use the in cluster image registry configuration. If you wish to use a mirror or any other settings specific to registries.conf, please specify those in the cluster wide registries.conf via the cluster image.config, ImageContentSourcePolicies, ImageDigestMirrorSet, or ImageTagMirrorSet objects. The format of the image pushspec is: host[:port][/namespace]/name:<tag> or svc_name.namespace.svc[:port]/repository/name:<tag>",
	"releaseVersion":                "releaseVersion is an Openshift release version which the base OS image is associated with. This field is populated from the machine-config-osimageurl configmap in the openshift-machine-config-operator namespace. It will come in the format: 4.16.0-0.nightly-2024-04-03-065948 or any valid release. The MachineOSBuilder populates this field and validates that this is a valid stream. This is used as a label in the Containerfile that builds the OS image.",
	"containerFile":                 "containerFile describes the custom data the user has specified to build into the image. This is also commonly called a Dockerfile and you can treat it as such. The content is the content of your Dockerfile. See https://github.com/containers/common/blob/main/docs/Containerfile.5.md for the spec reference. you can specify up to 7 containerFiles",
}

func (BuildInputs) SwaggerDoc() map[string]string {
	return map_BuildInputs
}

var map_BuildOutputs = map[string]string{
	"":                       "BuildOutputs holds all information needed to handle booting the image after a build",
	"currentImagePullSecret": "currentImagePullSecret is the secret used to pull the final produced image. Must live in the openshift-machine-config-operator namespace, the final image push and pull secrets should be separate for security concerns. If the final image push secret is somehow exfiltrated, that gives someone the power to push images to the image repository. By comparison, if the final image pull secret gets exfiltrated, that only gives someone to pull images from the image repository. It's basically the principle of least permissions. This pull secret will be used on all nodes in the pool. These nodes will need to pull the final OS image and boot into it using rpm-ostree or bootc.",
}

func (BuildOutputs) SwaggerDoc() map[string]string {
	return map_BuildOutputs
}

var map_ImageSecretObjectReference = map[string]string{
	"":     "Refers to the name of an image registry push/pull secret needed in the build process.",
	"name": "name is the name of the secret used to push or pull this MachineOSConfig object. Must consist of lower case alphanumeric characters, '-' or '.', and must start and end with an alphanumeric character. This secret must be in the openshift-machine-config-operator namespace.",
}

func (ImageSecretObjectReference) SwaggerDoc() map[string]string {
	return map_ImageSecretObjectReference
}

var map_MachineConfigPoolReference = map[string]string{
	"":     "Refers to the name of a MachineConfigPool (e.g., \"worker\", \"infra\", etc.): the MachineOSBuilder pod validates that the user has provided a valid pool",
	"name": "name of the MachineConfigPool object. Must consist of lower case alphanumeric characters, '-' or '.', and must start and end with an alphanumeric character.",
}

func (MachineConfigPoolReference) SwaggerDoc() map[string]string {
	return map_MachineConfigPoolReference
}

var map_MachineOSConfig = map[string]string{
	"":       "MachineOSConfig describes the configuration for a build process managed by the MCO Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
	"spec":   "spec describes the configuration of the machineosconfig",
	"status": "status describes the status of the machineosconfig",
}

func (MachineOSConfig) SwaggerDoc() map[string]string {
	return map_MachineOSConfig
}

var map_MachineOSConfigList = map[string]string{
	"":      "MachineOSConfigList describes all configurations for image builds on the system\n\nCompatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
	"items": "items contains a collection of MachineOSConfig resources.",
}

func (MachineOSConfigList) SwaggerDoc() map[string]string {
	return map_MachineOSConfigList
}

var map_MachineOSConfigSpec = map[string]string{
	"":                  "MachineOSConfigSpec describes user-configurable options as well as information about a build process.",
	"machineConfigPool": "machineConfigPool is the pool which the build is for",
	"buildInputs":       "buildInputs is where user input options for the build live",
	"buildOutputs":      "buildOutputs holds all information needed to handle booting the image after a build This currently contains a currentImagePullSecret field, which should be provided if the final pull secret used to pull the image to nodes from the registry is different than the one used for pushing the image to the registry during the build.",
}

func (MachineOSConfigSpec) SwaggerDoc() map[string]string {
	return map_MachineOSConfigSpec
}

var map_MachineOSConfigStatus = map[string]string{
	"":                     "MachineOSConfigStatus describes the status this config object and relates it to the builds associated with this MachineOSConfig",
	"observedGeneration":   "observedGeneration represents the generation of the MachineOSConfig object observed by the Machine Config Operator's build controller.",
	"currentImagePullspec": "currentImagePullspec is the fully qualified image pull spec used by the MCO to pull down the new OSImage. This must include sha256. The format of the image pullspec is: host[:port][/namespace]/name@sha256:<digest> The digest must be 64 characters long, and consist only of lowercase hexadecimal characters, a-f and 0-9.",
	"machineOSBuild":       "machineOSBuild is a reference to the MachineOSBuild object for this MachineOSConfig, which contains the status for the image build",
}

func (MachineOSConfigStatus) SwaggerDoc() map[string]string {
	return map_MachineOSConfigStatus
}

var map_MachineOSContainerfile = map[string]string{
	"":                  "MachineOSContainerfile contains all custom content the user wants built into the image",
	"containerfileArch": "containerfileArch describes the architecture this containerfile is to be built for. This arch is optional. If the user does not specify an architecture, it is assumed that the content can be applied to all architectures, or in a single arch cluster: the only architecture.",
	"content":           "content is an embedded Containerfile/Dockerfile that defines the contents to be built into your image. See https://github.com/containers/common/blob/main/docs/Containerfile.5.md for the spec reference. for example, this would add the tree package to your hosts:\n  FROM configs AS final\n  RUN rpm-ostree install tree && \\n    ostree container commit",
}

func (MachineOSContainerfile) SwaggerDoc() map[string]string {
	return map_MachineOSContainerfile
}

var map_MachineOSImageBuilder = map[string]string{
	"imageBuilderType": "imageBuilderType specifies the backend to be used to build the image. Valid options are: JobImageBuilder",
}

func (MachineOSImageBuilder) SwaggerDoc() map[string]string {
	return map_MachineOSImageBuilder
}

// AUTO-GENERATED FUNCTIONS END HERE
