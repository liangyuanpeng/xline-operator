# IoK8sApiCoreV1ServiceSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllocateLoadBalancerNodePorts** | Pointer to **bool** | allocateLoadBalancerNodePorts defines if NodePorts will be automatically allocated for services with type LoadBalancer.  Default is \&quot;true\&quot;. It may be set to \&quot;false\&quot; if the cluster load-balancer does not rely on NodePorts.  If the caller requests specific NodePorts (by specifying a value), those requests will be respected, regardless of this field. This field may only be set for services with type LoadBalancer and will be cleared if the type is changed to any other type. | [optional] 
**ClusterIP** | Pointer to **string** | clusterIP is the IP address of the service and is usually assigned randomly. If an address is specified manually, is in-range (as per system configuration), and is not in use, it will be allocated to the service; otherwise creation of the service will fail. This field may not be changed through updates unless the type field is also being changed to ExternalName (which requires this field to be blank) or the type field is being changed from ExternalName (in which case this field may optionally be specified, as describe above).  Valid values are \&quot;None\&quot;, empty string (\&quot;\&quot;), or a valid IP address. Setting this to \&quot;None\&quot; makes a \&quot;headless service\&quot; (no virtual IP), which is useful when direct endpoint connections are preferred and proxying is not required.  Only applies to types ClusterIP, NodePort, and LoadBalancer. If this field is specified when creating a Service of type ExternalName, creation will fail. This field will be wiped when updating a Service to type ExternalName. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies | [optional] 
**ClusterIPs** | Pointer to **[]string** | ClusterIPs is a list of IP addresses assigned to this service, and are usually assigned randomly.  If an address is specified manually, is in-range (as per system configuration), and is not in use, it will be allocated to the service; otherwise creation of the service will fail. This field may not be changed through updates unless the type field is also being changed to ExternalName (which requires this field to be empty) or the type field is being changed from ExternalName (in which case this field may optionally be specified, as describe above).  Valid values are \&quot;None\&quot;, empty string (\&quot;\&quot;), or a valid IP address.  Setting this to \&quot;None\&quot; makes a \&quot;headless service\&quot; (no virtual IP), which is useful when direct endpoint connections are preferred and proxying is not required.  Only applies to types ClusterIP, NodePort, and LoadBalancer. If this field is specified when creating a Service of type ExternalName, creation will fail. This field will be wiped when updating a Service to type ExternalName.  If this field is not specified, it will be initialized from the clusterIP field.  If this field is specified, clients must ensure that clusterIPs[0] and clusterIP have the same value.  This field may hold a maximum of two entries (dual-stack IPs, in either order). These IPs must correspond to the values of the ipFamilies field. Both clusterIPs and ipFamilies are governed by the ipFamilyPolicy field. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies | [optional] 
**ExternalIPs** | Pointer to **[]string** | externalIPs is a list of IP addresses for which nodes in the cluster will also accept traffic for this service.  These IPs are not managed by Kubernetes.  The user is responsible for ensuring that traffic arrives at a node with this IP.  A common example is external load-balancers that are not part of the Kubernetes system. | [optional] 
**ExternalName** | Pointer to **string** | externalName is the external reference that discovery mechanisms will return as an alias for this service (e.g. a DNS CNAME record). No proxying will be involved.  Must be a lowercase RFC-1123 hostname (https://tools.ietf.org/html/rfc1123) and requires &#x60;type&#x60; to be \&quot;ExternalName\&quot;. | [optional] 
**ExternalTrafficPolicy** | Pointer to **string** | externalTrafficPolicy describes how nodes distribute service traffic they receive on one of the Service&#39;s \&quot;externally-facing\&quot; addresses (NodePorts, ExternalIPs, and LoadBalancer IPs). If set to \&quot;Local\&quot;, the proxy will configure the service in a way that assumes that external load balancers will take care of balancing the service traffic between nodes, and so each node will deliver traffic only to the node-local endpoints of the service, without masquerading the client source IP. (Traffic mistakenly sent to a node with no endpoints will be dropped.) The default value, \&quot;Cluster\&quot;, uses the standard behavior of routing to all endpoints evenly (possibly modified by topology and other features). Note that traffic sent to an External IP or LoadBalancer IP from within the cluster will always get \&quot;Cluster\&quot; semantics, but clients sending to a NodePort from within the cluster may need to take traffic policy into account when picking a node.  Possible enum values:  - &#x60;\&quot;Cluster\&quot;&#x60;  - &#x60;\&quot;Cluster\&quot;&#x60; routes traffic to all endpoints.  - &#x60;\&quot;Local\&quot;&#x60;  - &#x60;\&quot;Local\&quot;&#x60; preserves the source IP of the traffic by routing only to endpoints on the same node as the traffic was received on (dropping the traffic if there are no local endpoints). | [optional] 
**HealthCheckNodePort** | Pointer to **int32** | healthCheckNodePort specifies the healthcheck nodePort for the service. This only applies when type is set to LoadBalancer and externalTrafficPolicy is set to Local. If a value is specified, is in-range, and is not in use, it will be used.  If not specified, a value will be automatically allocated.  External systems (e.g. load-balancers) can use this port to determine if a given node holds endpoints for this service or not.  If this field is specified when creating a Service which does not need it, creation will fail. This field will be wiped when updating a Service to no longer need it (e.g. changing type). This field cannot be updated once set. | [optional] 
**InternalTrafficPolicy** | Pointer to **string** | InternalTrafficPolicy describes how nodes distribute service traffic they receive on the ClusterIP. If set to \&quot;Local\&quot;, the proxy will assume that pods only want to talk to endpoints of the service on the same node as the pod, dropping the traffic if there are no local endpoints. The default value, \&quot;Cluster\&quot;, uses the standard behavior of routing to all endpoints evenly (possibly modified by topology and other features).  Possible enum values:  - &#x60;\&quot;Cluster\&quot;&#x60; routes traffic to all endpoints.  - &#x60;\&quot;Local\&quot;&#x60; routes traffic only to endpoints on the same node as the client pod (dropping the traffic if there are no local endpoints). | [optional] 
**IpFamilies** | Pointer to **[]string** | IPFamilies is a list of IP families (e.g. IPv4, IPv6) assigned to this service. This field is usually assigned automatically based on cluster configuration and the ipFamilyPolicy field. If this field is specified manually, the requested family is available in the cluster, and ipFamilyPolicy allows it, it will be used; otherwise creation of the service will fail. This field is conditionally mutable: it allows for adding or removing a secondary IP family, but it does not allow changing the primary IP family of the Service. Valid values are \&quot;IPv4\&quot; and \&quot;IPv6\&quot;.  This field only applies to Services of types ClusterIP, NodePort, and LoadBalancer, and does apply to \&quot;headless\&quot; services. This field will be wiped when updating a Service to type ExternalName.  This field may hold a maximum of two entries (dual-stack families, in either order).  These families must correspond to the values of the clusterIPs field, if specified. Both clusterIPs and ipFamilies are governed by the ipFamilyPolicy field. | [optional] 
**IpFamilyPolicy** | Pointer to **string** | IPFamilyPolicy represents the dual-stack-ness requested or required by this Service. If there is no value provided, then this field will be set to SingleStack. Services can be \&quot;SingleStack\&quot; (a single IP family), \&quot;PreferDualStack\&quot; (two IP families on dual-stack configured clusters or a single IP family on single-stack clusters), or \&quot;RequireDualStack\&quot; (two IP families on dual-stack configured clusters, otherwise fail). The ipFamilies and clusterIPs fields depend on the value of this field. This field will be wiped when updating a service to type ExternalName.  Possible enum values:  - &#x60;\&quot;PreferDualStack\&quot;&#x60; indicates that this service prefers dual-stack when the cluster is configured for dual-stack. If the cluster is not configured for dual-stack the service will be assigned a single IPFamily. If the IPFamily is not set in service.spec.ipFamilies then the service will be assigned the default IPFamily configured on the cluster  - &#x60;\&quot;RequireDualStack\&quot;&#x60; indicates that this service requires dual-stack. Using IPFamilyPolicyRequireDualStack on a single stack cluster will result in validation errors. The IPFamilies (and their order) assigned to this service is based on service.spec.ipFamilies. If service.spec.ipFamilies was not provided then it will be assigned according to how they are configured on the cluster. If service.spec.ipFamilies has only one entry then the alternative IPFamily will be added by apiserver  - &#x60;\&quot;SingleStack\&quot;&#x60; indicates that this service is required to have a single IPFamily. The IPFamily assigned is based on the default IPFamily used by the cluster or as identified by service.spec.ipFamilies field | [optional] 
**LoadBalancerClass** | Pointer to **string** | loadBalancerClass is the class of the load balancer implementation this Service belongs to. If specified, the value of this field must be a label-style identifier, with an optional prefix, e.g. \&quot;internal-vip\&quot; or \&quot;example.com/internal-vip\&quot;. Unprefixed names are reserved for end-users. This field can only be set when the Service type is &#39;LoadBalancer&#39;. If not set, the default load balancer implementation is used, today this is typically done through the cloud provider integration, but should apply for any default implementation. If set, it is assumed that a load balancer implementation is watching for Services with a matching class. Any default load balancer implementation (e.g. cloud providers) should ignore Services that set this field. This field can only be set when creating or updating a Service to type &#39;LoadBalancer&#39;. Once set, it can not be changed. This field will be wiped when a service is updated to a non &#39;LoadBalancer&#39; type. | [optional] 
**LoadBalancerIP** | Pointer to **string** | Only applies to Service Type: LoadBalancer. This feature depends on whether the underlying cloud-provider supports specifying the loadBalancerIP when a load balancer is created. This field will be ignored if the cloud-provider does not support the feature. Deprecated: This field was under-specified and its meaning varies across implementations, and it cannot support dual-stack. As of Kubernetes v1.24, users are encouraged to use implementation-specific annotations when available. This field may be removed in a future API version. | [optional] 
**LoadBalancerSourceRanges** | Pointer to **[]string** | If specified and supported by the platform, this will restrict traffic through the cloud-provider load-balancer will be restricted to the specified client IPs. This field will be ignored if the cloud-provider does not support the feature.\&quot; More info: https://kubernetes.io/docs/tasks/access-application-cluster/create-external-load-balancer/ | [optional] 
**Ports** | Pointer to [**[]IoK8sApiCoreV1ServicePort**](IoK8sApiCoreV1ServicePort.md) | The list of ports that are exposed by this service. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies | [optional] 
**PublishNotReadyAddresses** | Pointer to **bool** | publishNotReadyAddresses indicates that any agent which deals with endpoints for this Service should disregard any indications of ready/not-ready. The primary use case for setting this field is for a StatefulSet&#39;s Headless Service to propagate SRV DNS records for its Pods for the purpose of peer discovery. The Kubernetes controllers that generate Endpoints and EndpointSlice resources for Services interpret this to mean that all endpoints are considered \&quot;ready\&quot; even if the Pods themselves are not. Agents which consume only Kubernetes generated endpoints through the Endpoints or EndpointSlice resources can safely assume this behavior. | [optional] 
**Selector** | Pointer to **map[string]string** | Route service traffic to pods with label keys and values matching this selector. If empty or not present, the service is assumed to have an external process managing its endpoints, which Kubernetes will not modify. Only applies to types ClusterIP, NodePort, and LoadBalancer. Ignored if type is ExternalName. More info: https://kubernetes.io/docs/concepts/services-networking/service/ | [optional] 
**SessionAffinity** | Pointer to **string** | Supports \&quot;ClientIP\&quot; and \&quot;None\&quot;. Used to maintain session affinity. Enable client IP based session affinity. Must be ClientIP or None. Defaults to None. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies  Possible enum values:  - &#x60;\&quot;ClientIP\&quot;&#x60; is the Client IP based.  - &#x60;\&quot;None\&quot;&#x60; - no session affinity. | [optional] 
**SessionAffinityConfig** | Pointer to [**IoK8sApiCoreV1SessionAffinityConfig**](IoK8sApiCoreV1SessionAffinityConfig.md) |  | [optional] 
**Type** | Pointer to **string** | type determines how the Service is exposed. Defaults to ClusterIP. Valid options are ExternalName, ClusterIP, NodePort, and LoadBalancer. \&quot;ClusterIP\&quot; allocates a cluster-internal IP address for load-balancing to endpoints. Endpoints are determined by the selector or if that is not specified, by manual construction of an Endpoints object or EndpointSlice objects. If clusterIP is \&quot;None\&quot;, no virtual IP is allocated and the endpoints are published as a set of endpoints rather than a virtual IP. \&quot;NodePort\&quot; builds on ClusterIP and allocates a port on every node which routes to the same endpoints as the clusterIP. \&quot;LoadBalancer\&quot; builds on NodePort and creates an external load-balancer (if supported in the current cloud) which routes to the same endpoints as the clusterIP. \&quot;ExternalName\&quot; aliases this service to the specified externalName. Several other fields do not apply to ExternalName services. More info: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types  Possible enum values:  - &#x60;\&quot;ClusterIP\&quot;&#x60; means a service will only be accessible inside the cluster, via the cluster IP.  - &#x60;\&quot;ExternalName\&quot;&#x60; means a service consists of only a reference to an external name that kubedns or equivalent will return as a CNAME record, with no exposing or proxying of any pods involved.  - &#x60;\&quot;LoadBalancer\&quot;&#x60; means a service will be exposed via an external load balancer (if the cloud provider supports it), in addition to &#39;NodePort&#39; type.  - &#x60;\&quot;NodePort\&quot;&#x60; means a service will be exposed on one port of every node, in addition to &#39;ClusterIP&#39; type. | [optional] 

## Methods

### NewIoK8sApiCoreV1ServiceSpec

`func NewIoK8sApiCoreV1ServiceSpec() *IoK8sApiCoreV1ServiceSpec`

NewIoK8sApiCoreV1ServiceSpec instantiates a new IoK8sApiCoreV1ServiceSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1ServiceSpecWithDefaults

`func NewIoK8sApiCoreV1ServiceSpecWithDefaults() *IoK8sApiCoreV1ServiceSpec`

NewIoK8sApiCoreV1ServiceSpecWithDefaults instantiates a new IoK8sApiCoreV1ServiceSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllocateLoadBalancerNodePorts

`func (o *IoK8sApiCoreV1ServiceSpec) GetAllocateLoadBalancerNodePorts() bool`

GetAllocateLoadBalancerNodePorts returns the AllocateLoadBalancerNodePorts field if non-nil, zero value otherwise.

### GetAllocateLoadBalancerNodePortsOk

`func (o *IoK8sApiCoreV1ServiceSpec) GetAllocateLoadBalancerNodePortsOk() (*bool, bool)`

GetAllocateLoadBalancerNodePortsOk returns a tuple with the AllocateLoadBalancerNodePorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocateLoadBalancerNodePorts

`func (o *IoK8sApiCoreV1ServiceSpec) SetAllocateLoadBalancerNodePorts(v bool)`

SetAllocateLoadBalancerNodePorts sets AllocateLoadBalancerNodePorts field to given value.

### HasAllocateLoadBalancerNodePorts

`func (o *IoK8sApiCoreV1ServiceSpec) HasAllocateLoadBalancerNodePorts() bool`

HasAllocateLoadBalancerNodePorts returns a boolean if a field has been set.

### GetClusterIP

`func (o *IoK8sApiCoreV1ServiceSpec) GetClusterIP() string`

GetClusterIP returns the ClusterIP field if non-nil, zero value otherwise.

### GetClusterIPOk

`func (o *IoK8sApiCoreV1ServiceSpec) GetClusterIPOk() (*string, bool)`

GetClusterIPOk returns a tuple with the ClusterIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIP

`func (o *IoK8sApiCoreV1ServiceSpec) SetClusterIP(v string)`

SetClusterIP sets ClusterIP field to given value.

### HasClusterIP

`func (o *IoK8sApiCoreV1ServiceSpec) HasClusterIP() bool`

HasClusterIP returns a boolean if a field has been set.

### GetClusterIPs

`func (o *IoK8sApiCoreV1ServiceSpec) GetClusterIPs() []string`

GetClusterIPs returns the ClusterIPs field if non-nil, zero value otherwise.

### GetClusterIPsOk

`func (o *IoK8sApiCoreV1ServiceSpec) GetClusterIPsOk() (*[]string, bool)`

GetClusterIPsOk returns a tuple with the ClusterIPs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIPs

`func (o *IoK8sApiCoreV1ServiceSpec) SetClusterIPs(v []string)`

SetClusterIPs sets ClusterIPs field to given value.

### HasClusterIPs

`func (o *IoK8sApiCoreV1ServiceSpec) HasClusterIPs() bool`

HasClusterIPs returns a boolean if a field has been set.

### GetExternalIPs

`func (o *IoK8sApiCoreV1ServiceSpec) GetExternalIPs() []string`

GetExternalIPs returns the ExternalIPs field if non-nil, zero value otherwise.

### GetExternalIPsOk

`func (o *IoK8sApiCoreV1ServiceSpec) GetExternalIPsOk() (*[]string, bool)`

GetExternalIPsOk returns a tuple with the ExternalIPs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIPs

`func (o *IoK8sApiCoreV1ServiceSpec) SetExternalIPs(v []string)`

SetExternalIPs sets ExternalIPs field to given value.

### HasExternalIPs

`func (o *IoK8sApiCoreV1ServiceSpec) HasExternalIPs() bool`

HasExternalIPs returns a boolean if a field has been set.

### GetExternalName

`func (o *IoK8sApiCoreV1ServiceSpec) GetExternalName() string`

GetExternalName returns the ExternalName field if non-nil, zero value otherwise.

### GetExternalNameOk

`func (o *IoK8sApiCoreV1ServiceSpec) GetExternalNameOk() (*string, bool)`

GetExternalNameOk returns a tuple with the ExternalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalName

`func (o *IoK8sApiCoreV1ServiceSpec) SetExternalName(v string)`

SetExternalName sets ExternalName field to given value.

### HasExternalName

`func (o *IoK8sApiCoreV1ServiceSpec) HasExternalName() bool`

HasExternalName returns a boolean if a field has been set.

### GetExternalTrafficPolicy

`func (o *IoK8sApiCoreV1ServiceSpec) GetExternalTrafficPolicy() string`

GetExternalTrafficPolicy returns the ExternalTrafficPolicy field if non-nil, zero value otherwise.

### GetExternalTrafficPolicyOk

`func (o *IoK8sApiCoreV1ServiceSpec) GetExternalTrafficPolicyOk() (*string, bool)`

GetExternalTrafficPolicyOk returns a tuple with the ExternalTrafficPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalTrafficPolicy

`func (o *IoK8sApiCoreV1ServiceSpec) SetExternalTrafficPolicy(v string)`

SetExternalTrafficPolicy sets ExternalTrafficPolicy field to given value.

### HasExternalTrafficPolicy

`func (o *IoK8sApiCoreV1ServiceSpec) HasExternalTrafficPolicy() bool`

HasExternalTrafficPolicy returns a boolean if a field has been set.

### GetHealthCheckNodePort

`func (o *IoK8sApiCoreV1ServiceSpec) GetHealthCheckNodePort() int32`

GetHealthCheckNodePort returns the HealthCheckNodePort field if non-nil, zero value otherwise.

### GetHealthCheckNodePortOk

`func (o *IoK8sApiCoreV1ServiceSpec) GetHealthCheckNodePortOk() (*int32, bool)`

GetHealthCheckNodePortOk returns a tuple with the HealthCheckNodePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckNodePort

`func (o *IoK8sApiCoreV1ServiceSpec) SetHealthCheckNodePort(v int32)`

SetHealthCheckNodePort sets HealthCheckNodePort field to given value.

### HasHealthCheckNodePort

`func (o *IoK8sApiCoreV1ServiceSpec) HasHealthCheckNodePort() bool`

HasHealthCheckNodePort returns a boolean if a field has been set.

### GetInternalTrafficPolicy

`func (o *IoK8sApiCoreV1ServiceSpec) GetInternalTrafficPolicy() string`

GetInternalTrafficPolicy returns the InternalTrafficPolicy field if non-nil, zero value otherwise.

### GetInternalTrafficPolicyOk

`func (o *IoK8sApiCoreV1ServiceSpec) GetInternalTrafficPolicyOk() (*string, bool)`

GetInternalTrafficPolicyOk returns a tuple with the InternalTrafficPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalTrafficPolicy

`func (o *IoK8sApiCoreV1ServiceSpec) SetInternalTrafficPolicy(v string)`

SetInternalTrafficPolicy sets InternalTrafficPolicy field to given value.

### HasInternalTrafficPolicy

`func (o *IoK8sApiCoreV1ServiceSpec) HasInternalTrafficPolicy() bool`

HasInternalTrafficPolicy returns a boolean if a field has been set.

### GetIpFamilies

`func (o *IoK8sApiCoreV1ServiceSpec) GetIpFamilies() []string`

GetIpFamilies returns the IpFamilies field if non-nil, zero value otherwise.

### GetIpFamiliesOk

`func (o *IoK8sApiCoreV1ServiceSpec) GetIpFamiliesOk() (*[]string, bool)`

GetIpFamiliesOk returns a tuple with the IpFamilies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpFamilies

`func (o *IoK8sApiCoreV1ServiceSpec) SetIpFamilies(v []string)`

SetIpFamilies sets IpFamilies field to given value.

### HasIpFamilies

`func (o *IoK8sApiCoreV1ServiceSpec) HasIpFamilies() bool`

HasIpFamilies returns a boolean if a field has been set.

### GetIpFamilyPolicy

`func (o *IoK8sApiCoreV1ServiceSpec) GetIpFamilyPolicy() string`

GetIpFamilyPolicy returns the IpFamilyPolicy field if non-nil, zero value otherwise.

### GetIpFamilyPolicyOk

`func (o *IoK8sApiCoreV1ServiceSpec) GetIpFamilyPolicyOk() (*string, bool)`

GetIpFamilyPolicyOk returns a tuple with the IpFamilyPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpFamilyPolicy

`func (o *IoK8sApiCoreV1ServiceSpec) SetIpFamilyPolicy(v string)`

SetIpFamilyPolicy sets IpFamilyPolicy field to given value.

### HasIpFamilyPolicy

`func (o *IoK8sApiCoreV1ServiceSpec) HasIpFamilyPolicy() bool`

HasIpFamilyPolicy returns a boolean if a field has been set.

### GetLoadBalancerClass

`func (o *IoK8sApiCoreV1ServiceSpec) GetLoadBalancerClass() string`

GetLoadBalancerClass returns the LoadBalancerClass field if non-nil, zero value otherwise.

### GetLoadBalancerClassOk

`func (o *IoK8sApiCoreV1ServiceSpec) GetLoadBalancerClassOk() (*string, bool)`

GetLoadBalancerClassOk returns a tuple with the LoadBalancerClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerClass

`func (o *IoK8sApiCoreV1ServiceSpec) SetLoadBalancerClass(v string)`

SetLoadBalancerClass sets LoadBalancerClass field to given value.

### HasLoadBalancerClass

`func (o *IoK8sApiCoreV1ServiceSpec) HasLoadBalancerClass() bool`

HasLoadBalancerClass returns a boolean if a field has been set.

### GetLoadBalancerIP

`func (o *IoK8sApiCoreV1ServiceSpec) GetLoadBalancerIP() string`

GetLoadBalancerIP returns the LoadBalancerIP field if non-nil, zero value otherwise.

### GetLoadBalancerIPOk

`func (o *IoK8sApiCoreV1ServiceSpec) GetLoadBalancerIPOk() (*string, bool)`

GetLoadBalancerIPOk returns a tuple with the LoadBalancerIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerIP

`func (o *IoK8sApiCoreV1ServiceSpec) SetLoadBalancerIP(v string)`

SetLoadBalancerIP sets LoadBalancerIP field to given value.

### HasLoadBalancerIP

`func (o *IoK8sApiCoreV1ServiceSpec) HasLoadBalancerIP() bool`

HasLoadBalancerIP returns a boolean if a field has been set.

### GetLoadBalancerSourceRanges

`func (o *IoK8sApiCoreV1ServiceSpec) GetLoadBalancerSourceRanges() []string`

GetLoadBalancerSourceRanges returns the LoadBalancerSourceRanges field if non-nil, zero value otherwise.

### GetLoadBalancerSourceRangesOk

`func (o *IoK8sApiCoreV1ServiceSpec) GetLoadBalancerSourceRangesOk() (*[]string, bool)`

GetLoadBalancerSourceRangesOk returns a tuple with the LoadBalancerSourceRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerSourceRanges

`func (o *IoK8sApiCoreV1ServiceSpec) SetLoadBalancerSourceRanges(v []string)`

SetLoadBalancerSourceRanges sets LoadBalancerSourceRanges field to given value.

### HasLoadBalancerSourceRanges

`func (o *IoK8sApiCoreV1ServiceSpec) HasLoadBalancerSourceRanges() bool`

HasLoadBalancerSourceRanges returns a boolean if a field has been set.

### GetPorts

`func (o *IoK8sApiCoreV1ServiceSpec) GetPorts() []IoK8sApiCoreV1ServicePort`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *IoK8sApiCoreV1ServiceSpec) GetPortsOk() (*[]IoK8sApiCoreV1ServicePort, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *IoK8sApiCoreV1ServiceSpec) SetPorts(v []IoK8sApiCoreV1ServicePort)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *IoK8sApiCoreV1ServiceSpec) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetPublishNotReadyAddresses

`func (o *IoK8sApiCoreV1ServiceSpec) GetPublishNotReadyAddresses() bool`

GetPublishNotReadyAddresses returns the PublishNotReadyAddresses field if non-nil, zero value otherwise.

### GetPublishNotReadyAddressesOk

`func (o *IoK8sApiCoreV1ServiceSpec) GetPublishNotReadyAddressesOk() (*bool, bool)`

GetPublishNotReadyAddressesOk returns a tuple with the PublishNotReadyAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishNotReadyAddresses

`func (o *IoK8sApiCoreV1ServiceSpec) SetPublishNotReadyAddresses(v bool)`

SetPublishNotReadyAddresses sets PublishNotReadyAddresses field to given value.

### HasPublishNotReadyAddresses

`func (o *IoK8sApiCoreV1ServiceSpec) HasPublishNotReadyAddresses() bool`

HasPublishNotReadyAddresses returns a boolean if a field has been set.

### GetSelector

`func (o *IoK8sApiCoreV1ServiceSpec) GetSelector() map[string]string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *IoK8sApiCoreV1ServiceSpec) GetSelectorOk() (*map[string]string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *IoK8sApiCoreV1ServiceSpec) SetSelector(v map[string]string)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *IoK8sApiCoreV1ServiceSpec) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### GetSessionAffinity

`func (o *IoK8sApiCoreV1ServiceSpec) GetSessionAffinity() string`

GetSessionAffinity returns the SessionAffinity field if non-nil, zero value otherwise.

### GetSessionAffinityOk

`func (o *IoK8sApiCoreV1ServiceSpec) GetSessionAffinityOk() (*string, bool)`

GetSessionAffinityOk returns a tuple with the SessionAffinity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionAffinity

`func (o *IoK8sApiCoreV1ServiceSpec) SetSessionAffinity(v string)`

SetSessionAffinity sets SessionAffinity field to given value.

### HasSessionAffinity

`func (o *IoK8sApiCoreV1ServiceSpec) HasSessionAffinity() bool`

HasSessionAffinity returns a boolean if a field has been set.

### GetSessionAffinityConfig

`func (o *IoK8sApiCoreV1ServiceSpec) GetSessionAffinityConfig() IoK8sApiCoreV1SessionAffinityConfig`

GetSessionAffinityConfig returns the SessionAffinityConfig field if non-nil, zero value otherwise.

### GetSessionAffinityConfigOk

`func (o *IoK8sApiCoreV1ServiceSpec) GetSessionAffinityConfigOk() (*IoK8sApiCoreV1SessionAffinityConfig, bool)`

GetSessionAffinityConfigOk returns a tuple with the SessionAffinityConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionAffinityConfig

`func (o *IoK8sApiCoreV1ServiceSpec) SetSessionAffinityConfig(v IoK8sApiCoreV1SessionAffinityConfig)`

SetSessionAffinityConfig sets SessionAffinityConfig field to given value.

### HasSessionAffinityConfig

`func (o *IoK8sApiCoreV1ServiceSpec) HasSessionAffinityConfig() bool`

HasSessionAffinityConfig returns a boolean if a field has been set.

### GetType

`func (o *IoK8sApiCoreV1ServiceSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IoK8sApiCoreV1ServiceSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IoK8sApiCoreV1ServiceSpec) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IoK8sApiCoreV1ServiceSpec) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


