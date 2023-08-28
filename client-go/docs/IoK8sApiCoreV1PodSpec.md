# IoK8sApiCoreV1PodSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveDeadlineSeconds** | Pointer to **int64** | Optional duration in seconds the pod may be active on the node relative to StartTime before the system will actively try to mark it failed and kill associated containers. Value must be a positive integer. | [optional] 
**Affinity** | Pointer to [**IoK8sApiCoreV1Affinity**](IoK8sApiCoreV1Affinity.md) |  | [optional] 
**AutomountServiceAccountToken** | Pointer to **bool** | AutomountServiceAccountToken indicates whether a service account token should be automatically mounted. | [optional] 
**Containers** | [**[]IoK8sApiCoreV1Container**](IoK8sApiCoreV1Container.md) | List of containers belonging to the pod. Containers cannot currently be added or removed. There must be at least one container in a Pod. Cannot be updated. | 
**DnsConfig** | Pointer to [**IoK8sApiCoreV1PodDNSConfig**](IoK8sApiCoreV1PodDNSConfig.md) |  | [optional] 
**DnsPolicy** | Pointer to **string** | Set DNS policy for the pod. Defaults to \&quot;ClusterFirst\&quot;. Valid values are &#39;ClusterFirstWithHostNet&#39;, &#39;ClusterFirst&#39;, &#39;Default&#39; or &#39;None&#39;. DNS parameters given in DNSConfig will be merged with the policy selected with DNSPolicy. To have DNS options set along with hostNetwork, you have to specify DNS policy explicitly to &#39;ClusterFirstWithHostNet&#39;.  Possible enum values:  - &#x60;\&quot;ClusterFirst\&quot;&#x60; indicates that the pod should use cluster DNS first unless hostNetwork is true, if it is available, then fall back on the default (as determined by kubelet) DNS settings.  - &#x60;\&quot;ClusterFirstWithHostNet\&quot;&#x60; indicates that the pod should use cluster DNS first, if it is available, then fall back on the default (as determined by kubelet) DNS settings.  - &#x60;\&quot;Default\&quot;&#x60; indicates that the pod should use the default (as determined by kubelet) DNS settings.  - &#x60;\&quot;None\&quot;&#x60; indicates that the pod should use empty DNS settings. DNS parameters such as nameservers and search paths should be defined via DNSConfig. | [optional] 
**EnableServiceLinks** | Pointer to **bool** | EnableServiceLinks indicates whether information about services should be injected into pod&#39;s environment variables, matching the syntax of Docker links. Optional: Defaults to true. | [optional] 
**EphemeralContainers** | Pointer to [**[]IoK8sApiCoreV1EphemeralContainer**](IoK8sApiCoreV1EphemeralContainer.md) | List of ephemeral containers run in this pod. Ephemeral containers may be run in an existing pod to perform user-initiated actions such as debugging. This list cannot be specified when creating a pod, and it cannot be modified by updating the pod spec. In order to add an ephemeral container to an existing pod, use the pod&#39;s ephemeralcontainers subresource. | [optional] 
**HostAliases** | Pointer to [**[]IoK8sApiCoreV1HostAlias**](IoK8sApiCoreV1HostAlias.md) | HostAliases is an optional list of hosts and IPs that will be injected into the pod&#39;s hosts file if specified. This is only valid for non-hostNetwork pods. | [optional] 
**HostIPC** | Pointer to **bool** | Use the host&#39;s ipc namespace. Optional: Default to false. | [optional] 
**HostNetwork** | Pointer to **bool** | Host networking requested for this pod. Use the host&#39;s network namespace. If this option is set, the ports that will be used must be specified. Default to false. | [optional] 
**HostPID** | Pointer to **bool** | Use the host&#39;s pid namespace. Optional: Default to false. | [optional] 
**HostUsers** | Pointer to **bool** | Use the host&#39;s user namespace. Optional: Default to true. If set to true or not present, the pod will be run in the host user namespace, useful for when the pod needs a feature only available to the host user namespace, such as loading a kernel module with CAP_SYS_MODULE. When set to false, a new userns is created for the pod. Setting false is useful for mitigating container breakout vulnerabilities even allowing users to run their containers as root without actually having root privileges on the host. This field is alpha-level and is only honored by servers that enable the UserNamespacesSupport feature. | [optional] 
**Hostname** | Pointer to **string** | Specifies the hostname of the Pod If not specified, the pod&#39;s hostname will be set to a system-defined value. | [optional] 
**ImagePullSecrets** | Pointer to [**[]IoK8sApiCoreV1LocalObjectReference**](IoK8sApiCoreV1LocalObjectReference.md) | ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec. If specified, these secrets will be passed to individual puller implementations for them to use. More info: https://kubernetes.io/docs/concepts/containers/images#specifying-imagepullsecrets-on-a-pod | [optional] 
**InitContainers** | Pointer to [**[]IoK8sApiCoreV1Container**](IoK8sApiCoreV1Container.md) | List of initialization containers belonging to the pod. Init containers are executed in order prior to containers being started. If any init container fails, the pod is considered to have failed and is handled according to its restartPolicy. The name for an init container or normal container must be unique among all containers. Init containers may not have Lifecycle actions, Readiness probes, Liveness probes, or Startup probes. The resourceRequirements of an init container are taken into account during scheduling by finding the highest request/limit for each resource type, and then using the max of of that value or the sum of the normal containers. Limits are applied to init containers in a similar fashion. Init containers cannot currently be added or removed. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/ | [optional] 
**NodeName** | Pointer to **string** | NodeName is a request to schedule this pod onto a specific node. If it is non-empty, the scheduler simply schedules this pod onto that node, assuming that it fits resource requirements. | [optional] 
**NodeSelector** | Pointer to **map[string]string** | NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node&#39;s labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/ | [optional] 
**Os** | Pointer to [**IoK8sApiCoreV1PodOS**](IoK8sApiCoreV1PodOS.md) |  | [optional] 
**Overhead** | Pointer to **map[string]string** | Overhead represents the resource overhead associated with running a pod for a given RuntimeClass. This field will be autopopulated at admission time by the RuntimeClass admission controller. If the RuntimeClass admission controller is enabled, overhead must not be set in Pod create requests. The RuntimeClass admission controller will reject Pod create requests which have the overhead already set. If RuntimeClass is configured and selected in the PodSpec, Overhead will be set to the value defined in the corresponding RuntimeClass, otherwise it will remain unset and treated as zero. More info: https://git.k8s.io/enhancements/keps/sig-node/688-pod-overhead/README.md | [optional] 
**PreemptionPolicy** | Pointer to **string** | PreemptionPolicy is the Policy for preempting pods with lower priority. One of Never, PreemptLowerPriority. Defaults to PreemptLowerPriority if unset.  Possible enum values:  - &#x60;\&quot;Never\&quot;&#x60; means that pod never preempts other pods with lower priority.  - &#x60;\&quot;PreemptLowerPriority\&quot;&#x60; means that pod can preempt other pods with lower priority. | [optional] 
**Priority** | Pointer to **int32** | The priority value. Various system components use this field to find the priority of the pod. When Priority Admission Controller is enabled, it prevents users from setting this field. The admission controller populates this field from PriorityClassName. The higher the value, the higher the priority. | [optional] 
**PriorityClassName** | Pointer to **string** | If specified, indicates the pod&#39;s priority. \&quot;system-node-critical\&quot; and \&quot;system-cluster-critical\&quot; are two special keywords which indicate the highest priorities with the former being the highest priority. Any other name must be defined by creating a PriorityClass object with that name. If not specified, the pod priority will be default or zero if there is no default. | [optional] 
**ReadinessGates** | Pointer to [**[]IoK8sApiCoreV1PodReadinessGate**](IoK8sApiCoreV1PodReadinessGate.md) | If specified, all readiness gates will be evaluated for pod readiness. A pod is ready when all its containers are ready AND all conditions specified in the readiness gates have status equal to \&quot;True\&quot; More info: https://git.k8s.io/enhancements/keps/sig-network/580-pod-readiness-gates | [optional] 
**ResourceClaims** | Pointer to [**[]IoK8sApiCoreV1PodResourceClaim**](IoK8sApiCoreV1PodResourceClaim.md) | ResourceClaims defines which ResourceClaims must be allocated and reserved before the Pod is allowed to start. The resources will be made available to those containers which consume them by name.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. | [optional] 
**RestartPolicy** | Pointer to **string** | Restart policy for all containers within the pod. One of Always, OnFailure, Never. In some contexts, only a subset of those values may be permitted. Default to Always. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy  Possible enum values:  - &#x60;\&quot;Always\&quot;&#x60;  - &#x60;\&quot;Never\&quot;&#x60;  - &#x60;\&quot;OnFailure\&quot;&#x60; | [optional] 
**RuntimeClassName** | Pointer to **string** | RuntimeClassName refers to a RuntimeClass object in the node.k8s.io group, which should be used to run this pod.  If no RuntimeClass resource matches the named class, the pod will not be run. If unset or empty, the \&quot;legacy\&quot; RuntimeClass will be used, which is an implicit class with an empty definition that uses the default runtime handler. More info: https://git.k8s.io/enhancements/keps/sig-node/585-runtime-class | [optional] 
**SchedulerName** | Pointer to **string** | If specified, the pod will be dispatched by specified scheduler. If not specified, the pod will be dispatched by default scheduler. | [optional] 
**SchedulingGates** | Pointer to [**[]IoK8sApiCoreV1PodSchedulingGate**](IoK8sApiCoreV1PodSchedulingGate.md) | SchedulingGates is an opaque list of values that if specified will block scheduling the pod. If schedulingGates is not empty, the pod will stay in the SchedulingGated state and the scheduler will not attempt to schedule the pod.  SchedulingGates can only be set at pod creation time, and be removed only afterwards.  This is a beta feature enabled by the PodSchedulingReadiness feature gate. | [optional] 
**SecurityContext** | Pointer to [**IoK8sApiCoreV1PodSecurityContext**](IoK8sApiCoreV1PodSecurityContext.md) |  | [optional] 
**ServiceAccount** | Pointer to **string** | DeprecatedServiceAccount is a depreciated alias for ServiceAccountName. Deprecated: Use serviceAccountName instead. | [optional] 
**ServiceAccountName** | Pointer to **string** | ServiceAccountName is the name of the ServiceAccount to use to run this pod. More info: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/ | [optional] 
**SetHostnameAsFQDN** | Pointer to **bool** | If true the pod&#39;s hostname will be configured as the pod&#39;s FQDN, rather than the leaf name (the default). In Linux containers, this means setting the FQDN in the hostname field of the kernel (the nodename field of struct utsname). In Windows containers, this means setting the registry value of hostname for the registry key HKEY_LOCAL_MACHINE\\SYSTEM\\CurrentControlSet\\Services\\Tcpip\\Parameters to FQDN. If a pod does not have FQDN, this has no effect. Default to false. | [optional] 
**ShareProcessNamespace** | Pointer to **bool** | Share a single process namespace between all of the containers in a pod. When this is set containers will be able to view and signal processes from other containers in the same pod, and the first process in each container will not be assigned PID 1. HostPID and ShareProcessNamespace cannot both be set. Optional: Default to false. | [optional] 
**Subdomain** | Pointer to **string** | If specified, the fully qualified Pod hostname will be \&quot;&lt;hostname&gt;.&lt;subdomain&gt;.&lt;pod namespace&gt;.svc.&lt;cluster domain&gt;\&quot;. If not specified, the pod will not have a domainname at all. | [optional] 
**TerminationGracePeriodSeconds** | Pointer to **int64** | Optional duration in seconds the pod needs to terminate gracefully. May be decreased in delete request. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). If this value is nil, the default grace period will be used instead. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. Defaults to 30 seconds. | [optional] 
**Tolerations** | Pointer to [**[]IoK8sApiCoreV1Toleration**](IoK8sApiCoreV1Toleration.md) | If specified, the pod&#39;s tolerations. | [optional] 
**TopologySpreadConstraints** | Pointer to [**[]IoK8sApiCoreV1TopologySpreadConstraint**](IoK8sApiCoreV1TopologySpreadConstraint.md) | TopologySpreadConstraints describes how a group of pods ought to spread across topology domains. Scheduler will schedule pods in a way which abides by the constraints. All topologySpreadConstraints are ANDed. | [optional] 
**Volumes** | Pointer to [**[]IoK8sApiCoreV1Volume**](IoK8sApiCoreV1Volume.md) | List of volumes that can be mounted by containers belonging to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes | [optional] 

## Methods

### NewIoK8sApiCoreV1PodSpec

`func NewIoK8sApiCoreV1PodSpec(containers []IoK8sApiCoreV1Container, ) *IoK8sApiCoreV1PodSpec`

NewIoK8sApiCoreV1PodSpec instantiates a new IoK8sApiCoreV1PodSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1PodSpecWithDefaults

`func NewIoK8sApiCoreV1PodSpecWithDefaults() *IoK8sApiCoreV1PodSpec`

NewIoK8sApiCoreV1PodSpecWithDefaults instantiates a new IoK8sApiCoreV1PodSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveDeadlineSeconds

`func (o *IoK8sApiCoreV1PodSpec) GetActiveDeadlineSeconds() int64`

GetActiveDeadlineSeconds returns the ActiveDeadlineSeconds field if non-nil, zero value otherwise.

### GetActiveDeadlineSecondsOk

`func (o *IoK8sApiCoreV1PodSpec) GetActiveDeadlineSecondsOk() (*int64, bool)`

GetActiveDeadlineSecondsOk returns a tuple with the ActiveDeadlineSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDeadlineSeconds

`func (o *IoK8sApiCoreV1PodSpec) SetActiveDeadlineSeconds(v int64)`

SetActiveDeadlineSeconds sets ActiveDeadlineSeconds field to given value.

### HasActiveDeadlineSeconds

`func (o *IoK8sApiCoreV1PodSpec) HasActiveDeadlineSeconds() bool`

HasActiveDeadlineSeconds returns a boolean if a field has been set.

### GetAffinity

`func (o *IoK8sApiCoreV1PodSpec) GetAffinity() IoK8sApiCoreV1Affinity`

GetAffinity returns the Affinity field if non-nil, zero value otherwise.

### GetAffinityOk

`func (o *IoK8sApiCoreV1PodSpec) GetAffinityOk() (*IoK8sApiCoreV1Affinity, bool)`

GetAffinityOk returns a tuple with the Affinity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffinity

`func (o *IoK8sApiCoreV1PodSpec) SetAffinity(v IoK8sApiCoreV1Affinity)`

SetAffinity sets Affinity field to given value.

### HasAffinity

`func (o *IoK8sApiCoreV1PodSpec) HasAffinity() bool`

HasAffinity returns a boolean if a field has been set.

### GetAutomountServiceAccountToken

`func (o *IoK8sApiCoreV1PodSpec) GetAutomountServiceAccountToken() bool`

GetAutomountServiceAccountToken returns the AutomountServiceAccountToken field if non-nil, zero value otherwise.

### GetAutomountServiceAccountTokenOk

`func (o *IoK8sApiCoreV1PodSpec) GetAutomountServiceAccountTokenOk() (*bool, bool)`

GetAutomountServiceAccountTokenOk returns a tuple with the AutomountServiceAccountToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomountServiceAccountToken

`func (o *IoK8sApiCoreV1PodSpec) SetAutomountServiceAccountToken(v bool)`

SetAutomountServiceAccountToken sets AutomountServiceAccountToken field to given value.

### HasAutomountServiceAccountToken

`func (o *IoK8sApiCoreV1PodSpec) HasAutomountServiceAccountToken() bool`

HasAutomountServiceAccountToken returns a boolean if a field has been set.

### GetContainers

`func (o *IoK8sApiCoreV1PodSpec) GetContainers() []IoK8sApiCoreV1Container`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *IoK8sApiCoreV1PodSpec) GetContainersOk() (*[]IoK8sApiCoreV1Container, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *IoK8sApiCoreV1PodSpec) SetContainers(v []IoK8sApiCoreV1Container)`

SetContainers sets Containers field to given value.


### GetDnsConfig

`func (o *IoK8sApiCoreV1PodSpec) GetDnsConfig() IoK8sApiCoreV1PodDNSConfig`

GetDnsConfig returns the DnsConfig field if non-nil, zero value otherwise.

### GetDnsConfigOk

`func (o *IoK8sApiCoreV1PodSpec) GetDnsConfigOk() (*IoK8sApiCoreV1PodDNSConfig, bool)`

GetDnsConfigOk returns a tuple with the DnsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsConfig

`func (o *IoK8sApiCoreV1PodSpec) SetDnsConfig(v IoK8sApiCoreV1PodDNSConfig)`

SetDnsConfig sets DnsConfig field to given value.

### HasDnsConfig

`func (o *IoK8sApiCoreV1PodSpec) HasDnsConfig() bool`

HasDnsConfig returns a boolean if a field has been set.

### GetDnsPolicy

`func (o *IoK8sApiCoreV1PodSpec) GetDnsPolicy() string`

GetDnsPolicy returns the DnsPolicy field if non-nil, zero value otherwise.

### GetDnsPolicyOk

`func (o *IoK8sApiCoreV1PodSpec) GetDnsPolicyOk() (*string, bool)`

GetDnsPolicyOk returns a tuple with the DnsPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsPolicy

`func (o *IoK8sApiCoreV1PodSpec) SetDnsPolicy(v string)`

SetDnsPolicy sets DnsPolicy field to given value.

### HasDnsPolicy

`func (o *IoK8sApiCoreV1PodSpec) HasDnsPolicy() bool`

HasDnsPolicy returns a boolean if a field has been set.

### GetEnableServiceLinks

`func (o *IoK8sApiCoreV1PodSpec) GetEnableServiceLinks() bool`

GetEnableServiceLinks returns the EnableServiceLinks field if non-nil, zero value otherwise.

### GetEnableServiceLinksOk

`func (o *IoK8sApiCoreV1PodSpec) GetEnableServiceLinksOk() (*bool, bool)`

GetEnableServiceLinksOk returns a tuple with the EnableServiceLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableServiceLinks

`func (o *IoK8sApiCoreV1PodSpec) SetEnableServiceLinks(v bool)`

SetEnableServiceLinks sets EnableServiceLinks field to given value.

### HasEnableServiceLinks

`func (o *IoK8sApiCoreV1PodSpec) HasEnableServiceLinks() bool`

HasEnableServiceLinks returns a boolean if a field has been set.

### GetEphemeralContainers

`func (o *IoK8sApiCoreV1PodSpec) GetEphemeralContainers() []IoK8sApiCoreV1EphemeralContainer`

GetEphemeralContainers returns the EphemeralContainers field if non-nil, zero value otherwise.

### GetEphemeralContainersOk

`func (o *IoK8sApiCoreV1PodSpec) GetEphemeralContainersOk() (*[]IoK8sApiCoreV1EphemeralContainer, bool)`

GetEphemeralContainersOk returns a tuple with the EphemeralContainers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEphemeralContainers

`func (o *IoK8sApiCoreV1PodSpec) SetEphemeralContainers(v []IoK8sApiCoreV1EphemeralContainer)`

SetEphemeralContainers sets EphemeralContainers field to given value.

### HasEphemeralContainers

`func (o *IoK8sApiCoreV1PodSpec) HasEphemeralContainers() bool`

HasEphemeralContainers returns a boolean if a field has been set.

### GetHostAliases

`func (o *IoK8sApiCoreV1PodSpec) GetHostAliases() []IoK8sApiCoreV1HostAlias`

GetHostAliases returns the HostAliases field if non-nil, zero value otherwise.

### GetHostAliasesOk

`func (o *IoK8sApiCoreV1PodSpec) GetHostAliasesOk() (*[]IoK8sApiCoreV1HostAlias, bool)`

GetHostAliasesOk returns a tuple with the HostAliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostAliases

`func (o *IoK8sApiCoreV1PodSpec) SetHostAliases(v []IoK8sApiCoreV1HostAlias)`

SetHostAliases sets HostAliases field to given value.

### HasHostAliases

`func (o *IoK8sApiCoreV1PodSpec) HasHostAliases() bool`

HasHostAliases returns a boolean if a field has been set.

### GetHostIPC

`func (o *IoK8sApiCoreV1PodSpec) GetHostIPC() bool`

GetHostIPC returns the HostIPC field if non-nil, zero value otherwise.

### GetHostIPCOk

`func (o *IoK8sApiCoreV1PodSpec) GetHostIPCOk() (*bool, bool)`

GetHostIPCOk returns a tuple with the HostIPC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostIPC

`func (o *IoK8sApiCoreV1PodSpec) SetHostIPC(v bool)`

SetHostIPC sets HostIPC field to given value.

### HasHostIPC

`func (o *IoK8sApiCoreV1PodSpec) HasHostIPC() bool`

HasHostIPC returns a boolean if a field has been set.

### GetHostNetwork

`func (o *IoK8sApiCoreV1PodSpec) GetHostNetwork() bool`

GetHostNetwork returns the HostNetwork field if non-nil, zero value otherwise.

### GetHostNetworkOk

`func (o *IoK8sApiCoreV1PodSpec) GetHostNetworkOk() (*bool, bool)`

GetHostNetworkOk returns a tuple with the HostNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostNetwork

`func (o *IoK8sApiCoreV1PodSpec) SetHostNetwork(v bool)`

SetHostNetwork sets HostNetwork field to given value.

### HasHostNetwork

`func (o *IoK8sApiCoreV1PodSpec) HasHostNetwork() bool`

HasHostNetwork returns a boolean if a field has been set.

### GetHostPID

`func (o *IoK8sApiCoreV1PodSpec) GetHostPID() bool`

GetHostPID returns the HostPID field if non-nil, zero value otherwise.

### GetHostPIDOk

`func (o *IoK8sApiCoreV1PodSpec) GetHostPIDOk() (*bool, bool)`

GetHostPIDOk returns a tuple with the HostPID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostPID

`func (o *IoK8sApiCoreV1PodSpec) SetHostPID(v bool)`

SetHostPID sets HostPID field to given value.

### HasHostPID

`func (o *IoK8sApiCoreV1PodSpec) HasHostPID() bool`

HasHostPID returns a boolean if a field has been set.

### GetHostUsers

`func (o *IoK8sApiCoreV1PodSpec) GetHostUsers() bool`

GetHostUsers returns the HostUsers field if non-nil, zero value otherwise.

### GetHostUsersOk

`func (o *IoK8sApiCoreV1PodSpec) GetHostUsersOk() (*bool, bool)`

GetHostUsersOk returns a tuple with the HostUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostUsers

`func (o *IoK8sApiCoreV1PodSpec) SetHostUsers(v bool)`

SetHostUsers sets HostUsers field to given value.

### HasHostUsers

`func (o *IoK8sApiCoreV1PodSpec) HasHostUsers() bool`

HasHostUsers returns a boolean if a field has been set.

### GetHostname

`func (o *IoK8sApiCoreV1PodSpec) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *IoK8sApiCoreV1PodSpec) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *IoK8sApiCoreV1PodSpec) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *IoK8sApiCoreV1PodSpec) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetImagePullSecrets

`func (o *IoK8sApiCoreV1PodSpec) GetImagePullSecrets() []IoK8sApiCoreV1LocalObjectReference`

GetImagePullSecrets returns the ImagePullSecrets field if non-nil, zero value otherwise.

### GetImagePullSecretsOk

`func (o *IoK8sApiCoreV1PodSpec) GetImagePullSecretsOk() (*[]IoK8sApiCoreV1LocalObjectReference, bool)`

GetImagePullSecretsOk returns a tuple with the ImagePullSecrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagePullSecrets

`func (o *IoK8sApiCoreV1PodSpec) SetImagePullSecrets(v []IoK8sApiCoreV1LocalObjectReference)`

SetImagePullSecrets sets ImagePullSecrets field to given value.

### HasImagePullSecrets

`func (o *IoK8sApiCoreV1PodSpec) HasImagePullSecrets() bool`

HasImagePullSecrets returns a boolean if a field has been set.

### GetInitContainers

`func (o *IoK8sApiCoreV1PodSpec) GetInitContainers() []IoK8sApiCoreV1Container`

GetInitContainers returns the InitContainers field if non-nil, zero value otherwise.

### GetInitContainersOk

`func (o *IoK8sApiCoreV1PodSpec) GetInitContainersOk() (*[]IoK8sApiCoreV1Container, bool)`

GetInitContainersOk returns a tuple with the InitContainers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitContainers

`func (o *IoK8sApiCoreV1PodSpec) SetInitContainers(v []IoK8sApiCoreV1Container)`

SetInitContainers sets InitContainers field to given value.

### HasInitContainers

`func (o *IoK8sApiCoreV1PodSpec) HasInitContainers() bool`

HasInitContainers returns a boolean if a field has been set.

### GetNodeName

`func (o *IoK8sApiCoreV1PodSpec) GetNodeName() string`

GetNodeName returns the NodeName field if non-nil, zero value otherwise.

### GetNodeNameOk

`func (o *IoK8sApiCoreV1PodSpec) GetNodeNameOk() (*string, bool)`

GetNodeNameOk returns a tuple with the NodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeName

`func (o *IoK8sApiCoreV1PodSpec) SetNodeName(v string)`

SetNodeName sets NodeName field to given value.

### HasNodeName

`func (o *IoK8sApiCoreV1PodSpec) HasNodeName() bool`

HasNodeName returns a boolean if a field has been set.

### GetNodeSelector

`func (o *IoK8sApiCoreV1PodSpec) GetNodeSelector() map[string]string`

GetNodeSelector returns the NodeSelector field if non-nil, zero value otherwise.

### GetNodeSelectorOk

`func (o *IoK8sApiCoreV1PodSpec) GetNodeSelectorOk() (*map[string]string, bool)`

GetNodeSelectorOk returns a tuple with the NodeSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSelector

`func (o *IoK8sApiCoreV1PodSpec) SetNodeSelector(v map[string]string)`

SetNodeSelector sets NodeSelector field to given value.

### HasNodeSelector

`func (o *IoK8sApiCoreV1PodSpec) HasNodeSelector() bool`

HasNodeSelector returns a boolean if a field has been set.

### GetOs

`func (o *IoK8sApiCoreV1PodSpec) GetOs() IoK8sApiCoreV1PodOS`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *IoK8sApiCoreV1PodSpec) GetOsOk() (*IoK8sApiCoreV1PodOS, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *IoK8sApiCoreV1PodSpec) SetOs(v IoK8sApiCoreV1PodOS)`

SetOs sets Os field to given value.

### HasOs

`func (o *IoK8sApiCoreV1PodSpec) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetOverhead

`func (o *IoK8sApiCoreV1PodSpec) GetOverhead() map[string]string`

GetOverhead returns the Overhead field if non-nil, zero value otherwise.

### GetOverheadOk

`func (o *IoK8sApiCoreV1PodSpec) GetOverheadOk() (*map[string]string, bool)`

GetOverheadOk returns a tuple with the Overhead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverhead

`func (o *IoK8sApiCoreV1PodSpec) SetOverhead(v map[string]string)`

SetOverhead sets Overhead field to given value.

### HasOverhead

`func (o *IoK8sApiCoreV1PodSpec) HasOverhead() bool`

HasOverhead returns a boolean if a field has been set.

### GetPreemptionPolicy

`func (o *IoK8sApiCoreV1PodSpec) GetPreemptionPolicy() string`

GetPreemptionPolicy returns the PreemptionPolicy field if non-nil, zero value otherwise.

### GetPreemptionPolicyOk

`func (o *IoK8sApiCoreV1PodSpec) GetPreemptionPolicyOk() (*string, bool)`

GetPreemptionPolicyOk returns a tuple with the PreemptionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreemptionPolicy

`func (o *IoK8sApiCoreV1PodSpec) SetPreemptionPolicy(v string)`

SetPreemptionPolicy sets PreemptionPolicy field to given value.

### HasPreemptionPolicy

`func (o *IoK8sApiCoreV1PodSpec) HasPreemptionPolicy() bool`

HasPreemptionPolicy returns a boolean if a field has been set.

### GetPriority

`func (o *IoK8sApiCoreV1PodSpec) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *IoK8sApiCoreV1PodSpec) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *IoK8sApiCoreV1PodSpec) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *IoK8sApiCoreV1PodSpec) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetPriorityClassName

`func (o *IoK8sApiCoreV1PodSpec) GetPriorityClassName() string`

GetPriorityClassName returns the PriorityClassName field if non-nil, zero value otherwise.

### GetPriorityClassNameOk

`func (o *IoK8sApiCoreV1PodSpec) GetPriorityClassNameOk() (*string, bool)`

GetPriorityClassNameOk returns a tuple with the PriorityClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityClassName

`func (o *IoK8sApiCoreV1PodSpec) SetPriorityClassName(v string)`

SetPriorityClassName sets PriorityClassName field to given value.

### HasPriorityClassName

`func (o *IoK8sApiCoreV1PodSpec) HasPriorityClassName() bool`

HasPriorityClassName returns a boolean if a field has been set.

### GetReadinessGates

`func (o *IoK8sApiCoreV1PodSpec) GetReadinessGates() []IoK8sApiCoreV1PodReadinessGate`

GetReadinessGates returns the ReadinessGates field if non-nil, zero value otherwise.

### GetReadinessGatesOk

`func (o *IoK8sApiCoreV1PodSpec) GetReadinessGatesOk() (*[]IoK8sApiCoreV1PodReadinessGate, bool)`

GetReadinessGatesOk returns a tuple with the ReadinessGates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadinessGates

`func (o *IoK8sApiCoreV1PodSpec) SetReadinessGates(v []IoK8sApiCoreV1PodReadinessGate)`

SetReadinessGates sets ReadinessGates field to given value.

### HasReadinessGates

`func (o *IoK8sApiCoreV1PodSpec) HasReadinessGates() bool`

HasReadinessGates returns a boolean if a field has been set.

### GetResourceClaims

`func (o *IoK8sApiCoreV1PodSpec) GetResourceClaims() []IoK8sApiCoreV1PodResourceClaim`

GetResourceClaims returns the ResourceClaims field if non-nil, zero value otherwise.

### GetResourceClaimsOk

`func (o *IoK8sApiCoreV1PodSpec) GetResourceClaimsOk() (*[]IoK8sApiCoreV1PodResourceClaim, bool)`

GetResourceClaimsOk returns a tuple with the ResourceClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceClaims

`func (o *IoK8sApiCoreV1PodSpec) SetResourceClaims(v []IoK8sApiCoreV1PodResourceClaim)`

SetResourceClaims sets ResourceClaims field to given value.

### HasResourceClaims

`func (o *IoK8sApiCoreV1PodSpec) HasResourceClaims() bool`

HasResourceClaims returns a boolean if a field has been set.

### GetRestartPolicy

`func (o *IoK8sApiCoreV1PodSpec) GetRestartPolicy() string`

GetRestartPolicy returns the RestartPolicy field if non-nil, zero value otherwise.

### GetRestartPolicyOk

`func (o *IoK8sApiCoreV1PodSpec) GetRestartPolicyOk() (*string, bool)`

GetRestartPolicyOk returns a tuple with the RestartPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartPolicy

`func (o *IoK8sApiCoreV1PodSpec) SetRestartPolicy(v string)`

SetRestartPolicy sets RestartPolicy field to given value.

### HasRestartPolicy

`func (o *IoK8sApiCoreV1PodSpec) HasRestartPolicy() bool`

HasRestartPolicy returns a boolean if a field has been set.

### GetRuntimeClassName

`func (o *IoK8sApiCoreV1PodSpec) GetRuntimeClassName() string`

GetRuntimeClassName returns the RuntimeClassName field if non-nil, zero value otherwise.

### GetRuntimeClassNameOk

`func (o *IoK8sApiCoreV1PodSpec) GetRuntimeClassNameOk() (*string, bool)`

GetRuntimeClassNameOk returns a tuple with the RuntimeClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeClassName

`func (o *IoK8sApiCoreV1PodSpec) SetRuntimeClassName(v string)`

SetRuntimeClassName sets RuntimeClassName field to given value.

### HasRuntimeClassName

`func (o *IoK8sApiCoreV1PodSpec) HasRuntimeClassName() bool`

HasRuntimeClassName returns a boolean if a field has been set.

### GetSchedulerName

`func (o *IoK8sApiCoreV1PodSpec) GetSchedulerName() string`

GetSchedulerName returns the SchedulerName field if non-nil, zero value otherwise.

### GetSchedulerNameOk

`func (o *IoK8sApiCoreV1PodSpec) GetSchedulerNameOk() (*string, bool)`

GetSchedulerNameOk returns a tuple with the SchedulerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulerName

`func (o *IoK8sApiCoreV1PodSpec) SetSchedulerName(v string)`

SetSchedulerName sets SchedulerName field to given value.

### HasSchedulerName

`func (o *IoK8sApiCoreV1PodSpec) HasSchedulerName() bool`

HasSchedulerName returns a boolean if a field has been set.

### GetSchedulingGates

`func (o *IoK8sApiCoreV1PodSpec) GetSchedulingGates() []IoK8sApiCoreV1PodSchedulingGate`

GetSchedulingGates returns the SchedulingGates field if non-nil, zero value otherwise.

### GetSchedulingGatesOk

`func (o *IoK8sApiCoreV1PodSpec) GetSchedulingGatesOk() (*[]IoK8sApiCoreV1PodSchedulingGate, bool)`

GetSchedulingGatesOk returns a tuple with the SchedulingGates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulingGates

`func (o *IoK8sApiCoreV1PodSpec) SetSchedulingGates(v []IoK8sApiCoreV1PodSchedulingGate)`

SetSchedulingGates sets SchedulingGates field to given value.

### HasSchedulingGates

`func (o *IoK8sApiCoreV1PodSpec) HasSchedulingGates() bool`

HasSchedulingGates returns a boolean if a field has been set.

### GetSecurityContext

`func (o *IoK8sApiCoreV1PodSpec) GetSecurityContext() IoK8sApiCoreV1PodSecurityContext`

GetSecurityContext returns the SecurityContext field if non-nil, zero value otherwise.

### GetSecurityContextOk

`func (o *IoK8sApiCoreV1PodSpec) GetSecurityContextOk() (*IoK8sApiCoreV1PodSecurityContext, bool)`

GetSecurityContextOk returns a tuple with the SecurityContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityContext

`func (o *IoK8sApiCoreV1PodSpec) SetSecurityContext(v IoK8sApiCoreV1PodSecurityContext)`

SetSecurityContext sets SecurityContext field to given value.

### HasSecurityContext

`func (o *IoK8sApiCoreV1PodSpec) HasSecurityContext() bool`

HasSecurityContext returns a boolean if a field has been set.

### GetServiceAccount

`func (o *IoK8sApiCoreV1PodSpec) GetServiceAccount() string`

GetServiceAccount returns the ServiceAccount field if non-nil, zero value otherwise.

### GetServiceAccountOk

`func (o *IoK8sApiCoreV1PodSpec) GetServiceAccountOk() (*string, bool)`

GetServiceAccountOk returns a tuple with the ServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccount

`func (o *IoK8sApiCoreV1PodSpec) SetServiceAccount(v string)`

SetServiceAccount sets ServiceAccount field to given value.

### HasServiceAccount

`func (o *IoK8sApiCoreV1PodSpec) HasServiceAccount() bool`

HasServiceAccount returns a boolean if a field has been set.

### GetServiceAccountName

`func (o *IoK8sApiCoreV1PodSpec) GetServiceAccountName() string`

GetServiceAccountName returns the ServiceAccountName field if non-nil, zero value otherwise.

### GetServiceAccountNameOk

`func (o *IoK8sApiCoreV1PodSpec) GetServiceAccountNameOk() (*string, bool)`

GetServiceAccountNameOk returns a tuple with the ServiceAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountName

`func (o *IoK8sApiCoreV1PodSpec) SetServiceAccountName(v string)`

SetServiceAccountName sets ServiceAccountName field to given value.

### HasServiceAccountName

`func (o *IoK8sApiCoreV1PodSpec) HasServiceAccountName() bool`

HasServiceAccountName returns a boolean if a field has been set.

### GetSetHostnameAsFQDN

`func (o *IoK8sApiCoreV1PodSpec) GetSetHostnameAsFQDN() bool`

GetSetHostnameAsFQDN returns the SetHostnameAsFQDN field if non-nil, zero value otherwise.

### GetSetHostnameAsFQDNOk

`func (o *IoK8sApiCoreV1PodSpec) GetSetHostnameAsFQDNOk() (*bool, bool)`

GetSetHostnameAsFQDNOk returns a tuple with the SetHostnameAsFQDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetHostnameAsFQDN

`func (o *IoK8sApiCoreV1PodSpec) SetSetHostnameAsFQDN(v bool)`

SetSetHostnameAsFQDN sets SetHostnameAsFQDN field to given value.

### HasSetHostnameAsFQDN

`func (o *IoK8sApiCoreV1PodSpec) HasSetHostnameAsFQDN() bool`

HasSetHostnameAsFQDN returns a boolean if a field has been set.

### GetShareProcessNamespace

`func (o *IoK8sApiCoreV1PodSpec) GetShareProcessNamespace() bool`

GetShareProcessNamespace returns the ShareProcessNamespace field if non-nil, zero value otherwise.

### GetShareProcessNamespaceOk

`func (o *IoK8sApiCoreV1PodSpec) GetShareProcessNamespaceOk() (*bool, bool)`

GetShareProcessNamespaceOk returns a tuple with the ShareProcessNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareProcessNamespace

`func (o *IoK8sApiCoreV1PodSpec) SetShareProcessNamespace(v bool)`

SetShareProcessNamespace sets ShareProcessNamespace field to given value.

### HasShareProcessNamespace

`func (o *IoK8sApiCoreV1PodSpec) HasShareProcessNamespace() bool`

HasShareProcessNamespace returns a boolean if a field has been set.

### GetSubdomain

`func (o *IoK8sApiCoreV1PodSpec) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *IoK8sApiCoreV1PodSpec) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *IoK8sApiCoreV1PodSpec) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *IoK8sApiCoreV1PodSpec) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetTerminationGracePeriodSeconds

`func (o *IoK8sApiCoreV1PodSpec) GetTerminationGracePeriodSeconds() int64`

GetTerminationGracePeriodSeconds returns the TerminationGracePeriodSeconds field if non-nil, zero value otherwise.

### GetTerminationGracePeriodSecondsOk

`func (o *IoK8sApiCoreV1PodSpec) GetTerminationGracePeriodSecondsOk() (*int64, bool)`

GetTerminationGracePeriodSecondsOk returns a tuple with the TerminationGracePeriodSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationGracePeriodSeconds

`func (o *IoK8sApiCoreV1PodSpec) SetTerminationGracePeriodSeconds(v int64)`

SetTerminationGracePeriodSeconds sets TerminationGracePeriodSeconds field to given value.

### HasTerminationGracePeriodSeconds

`func (o *IoK8sApiCoreV1PodSpec) HasTerminationGracePeriodSeconds() bool`

HasTerminationGracePeriodSeconds returns a boolean if a field has been set.

### GetTolerations

`func (o *IoK8sApiCoreV1PodSpec) GetTolerations() []IoK8sApiCoreV1Toleration`

GetTolerations returns the Tolerations field if non-nil, zero value otherwise.

### GetTolerationsOk

`func (o *IoK8sApiCoreV1PodSpec) GetTolerationsOk() (*[]IoK8sApiCoreV1Toleration, bool)`

GetTolerationsOk returns a tuple with the Tolerations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTolerations

`func (o *IoK8sApiCoreV1PodSpec) SetTolerations(v []IoK8sApiCoreV1Toleration)`

SetTolerations sets Tolerations field to given value.

### HasTolerations

`func (o *IoK8sApiCoreV1PodSpec) HasTolerations() bool`

HasTolerations returns a boolean if a field has been set.

### GetTopologySpreadConstraints

`func (o *IoK8sApiCoreV1PodSpec) GetTopologySpreadConstraints() []IoK8sApiCoreV1TopologySpreadConstraint`

GetTopologySpreadConstraints returns the TopologySpreadConstraints field if non-nil, zero value otherwise.

### GetTopologySpreadConstraintsOk

`func (o *IoK8sApiCoreV1PodSpec) GetTopologySpreadConstraintsOk() (*[]IoK8sApiCoreV1TopologySpreadConstraint, bool)`

GetTopologySpreadConstraintsOk returns a tuple with the TopologySpreadConstraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologySpreadConstraints

`func (o *IoK8sApiCoreV1PodSpec) SetTopologySpreadConstraints(v []IoK8sApiCoreV1TopologySpreadConstraint)`

SetTopologySpreadConstraints sets TopologySpreadConstraints field to given value.

### HasTopologySpreadConstraints

`func (o *IoK8sApiCoreV1PodSpec) HasTopologySpreadConstraints() bool`

HasTopologySpreadConstraints returns a boolean if a field has been set.

### GetVolumes

`func (o *IoK8sApiCoreV1PodSpec) GetVolumes() []IoK8sApiCoreV1Volume`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *IoK8sApiCoreV1PodSpec) GetVolumesOk() (*[]IoK8sApiCoreV1Volume, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *IoK8sApiCoreV1PodSpec) SetVolumes(v []IoK8sApiCoreV1Volume)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *IoK8sApiCoreV1PodSpec) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


