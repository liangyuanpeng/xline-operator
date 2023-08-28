# IoK8sApiCoreV1PodStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to [**[]IoK8sApiCoreV1PodCondition**](IoK8sApiCoreV1PodCondition.md) | Current service state of pod. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-conditions | [optional] 
**ContainerStatuses** | Pointer to [**[]IoK8sApiCoreV1ContainerStatus**](IoK8sApiCoreV1ContainerStatus.md) | The list has one entry per container in the manifest. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-and-container-status | [optional] 
**EphemeralContainerStatuses** | Pointer to [**[]IoK8sApiCoreV1ContainerStatus**](IoK8sApiCoreV1ContainerStatus.md) | Status for any ephemeral containers that have run in this pod. | [optional] 
**HostIP** | Pointer to **string** | IP address of the host to which the pod is assigned. Empty if not yet scheduled. | [optional] 
**InitContainerStatuses** | Pointer to [**[]IoK8sApiCoreV1ContainerStatus**](IoK8sApiCoreV1ContainerStatus.md) | The list has one entry per init container in the manifest. The most recent successful init container will have ready &#x3D; true, the most recently started container will have startTime set. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-and-container-status | [optional] 
**Message** | Pointer to **string** | A human readable message indicating details about why the pod is in this condition. | [optional] 
**NominatedNodeName** | Pointer to **string** | nominatedNodeName is set only when this pod preempts other pods on the node, but it cannot be scheduled right away as preemption victims receive their graceful termination periods. This field does not guarantee that the pod will be scheduled on this node. Scheduler may decide to place the pod elsewhere if other nodes become available sooner. Scheduler may also decide to give the resources on this node to a higher priority pod that is created after preemption. As a result, this field may be different than PodSpec.nodeName when the pod is scheduled. | [optional] 
**Phase** | Pointer to **string** | The phase of a Pod is a simple, high-level summary of where the Pod is in its lifecycle. The conditions array, the reason and message fields, and the individual container status arrays contain more detail about the pod&#39;s status. There are five possible phase values:  Pending: The pod has been accepted by the Kubernetes system, but one or more of the container images has not been created. This includes time before being scheduled as well as time spent downloading images over the network, which could take a while. Running: The pod has been bound to a node, and all of the containers have been created. At least one container is still running, or is in the process of starting or restarting. Succeeded: All containers in the pod have terminated in success, and will not be restarted. Failed: All containers in the pod have terminated, and at least one container has terminated in failure. The container either exited with non-zero status or was terminated by the system. Unknown: For some reason the state of the pod could not be obtained, typically due to an error in communicating with the host of the pod.  More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-phase  Possible enum values:  - &#x60;\&quot;Failed\&quot;&#x60; means that all containers in the pod have terminated, and at least one container has terminated in a failure (exited with a non-zero exit code or was stopped by the system).  - &#x60;\&quot;Pending\&quot;&#x60; means the pod has been accepted by the system, but one or more of the containers has not been started. This includes time before being bound to a node, as well as time spent pulling images onto the host.  - &#x60;\&quot;Running\&quot;&#x60; means the pod has been bound to a node and all of the containers have been started. At least one container is still running or is in the process of being restarted.  - &#x60;\&quot;Succeeded\&quot;&#x60; means that all containers in the pod have voluntarily terminated with a container exit code of 0, and the system is not going to restart any of these containers.  - &#x60;\&quot;Unknown\&quot;&#x60; means that for some reason the state of the pod could not be obtained, typically due to an error in communicating with the host of the pod. Deprecated: It isn&#39;t being set since 2015 (74da3b14b0c0f658b3bb8d2def5094686d0e9095) | [optional] 
**PodIP** | Pointer to **string** | IP address allocated to the pod. Routable at least within the cluster. Empty if not yet allocated. | [optional] 
**PodIPs** | Pointer to [**[]IoK8sApiCoreV1PodIP**](IoK8sApiCoreV1PodIP.md) | podIPs holds the IP addresses allocated to the pod. If this field is specified, the 0th entry must match the podIP field. Pods may be allocated at most 1 value for each of IPv4 and IPv6. This list is empty if no IPs have been allocated yet. | [optional] 
**QosClass** | Pointer to **string** | The Quality of Service (QOS) classification assigned to the pod based on resource requirements See PodQOSClass type for available QOS classes More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-qos/#quality-of-service-classes  Possible enum values:  - &#x60;\&quot;BestEffort\&quot;&#x60; is the BestEffort qos class.  - &#x60;\&quot;Burstable\&quot;&#x60; is the Burstable qos class.  - &#x60;\&quot;Guaranteed\&quot;&#x60; is the Guaranteed qos class. | [optional] 
**Reason** | Pointer to **string** | A brief CamelCase message indicating details about why the pod is in this state. e.g. &#39;Evicted&#39; | [optional] 
**Resize** | Pointer to **string** | Status of resources resize desired for pod&#39;s containers. It is empty if no resources resize is pending. Any changes to container resources will automatically set this to \&quot;Proposed\&quot; | [optional] 
**StartTime** | Pointer to **time.Time** | Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers. | [optional] 

## Methods

### NewIoK8sApiCoreV1PodStatus

`func NewIoK8sApiCoreV1PodStatus() *IoK8sApiCoreV1PodStatus`

NewIoK8sApiCoreV1PodStatus instantiates a new IoK8sApiCoreV1PodStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1PodStatusWithDefaults

`func NewIoK8sApiCoreV1PodStatusWithDefaults() *IoK8sApiCoreV1PodStatus`

NewIoK8sApiCoreV1PodStatusWithDefaults instantiates a new IoK8sApiCoreV1PodStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *IoK8sApiCoreV1PodStatus) GetConditions() []IoK8sApiCoreV1PodCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *IoK8sApiCoreV1PodStatus) GetConditionsOk() (*[]IoK8sApiCoreV1PodCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *IoK8sApiCoreV1PodStatus) SetConditions(v []IoK8sApiCoreV1PodCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *IoK8sApiCoreV1PodStatus) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetContainerStatuses

`func (o *IoK8sApiCoreV1PodStatus) GetContainerStatuses() []IoK8sApiCoreV1ContainerStatus`

GetContainerStatuses returns the ContainerStatuses field if non-nil, zero value otherwise.

### GetContainerStatusesOk

`func (o *IoK8sApiCoreV1PodStatus) GetContainerStatusesOk() (*[]IoK8sApiCoreV1ContainerStatus, bool)`

GetContainerStatusesOk returns a tuple with the ContainerStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerStatuses

`func (o *IoK8sApiCoreV1PodStatus) SetContainerStatuses(v []IoK8sApiCoreV1ContainerStatus)`

SetContainerStatuses sets ContainerStatuses field to given value.

### HasContainerStatuses

`func (o *IoK8sApiCoreV1PodStatus) HasContainerStatuses() bool`

HasContainerStatuses returns a boolean if a field has been set.

### GetEphemeralContainerStatuses

`func (o *IoK8sApiCoreV1PodStatus) GetEphemeralContainerStatuses() []IoK8sApiCoreV1ContainerStatus`

GetEphemeralContainerStatuses returns the EphemeralContainerStatuses field if non-nil, zero value otherwise.

### GetEphemeralContainerStatusesOk

`func (o *IoK8sApiCoreV1PodStatus) GetEphemeralContainerStatusesOk() (*[]IoK8sApiCoreV1ContainerStatus, bool)`

GetEphemeralContainerStatusesOk returns a tuple with the EphemeralContainerStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEphemeralContainerStatuses

`func (o *IoK8sApiCoreV1PodStatus) SetEphemeralContainerStatuses(v []IoK8sApiCoreV1ContainerStatus)`

SetEphemeralContainerStatuses sets EphemeralContainerStatuses field to given value.

### HasEphemeralContainerStatuses

`func (o *IoK8sApiCoreV1PodStatus) HasEphemeralContainerStatuses() bool`

HasEphemeralContainerStatuses returns a boolean if a field has been set.

### GetHostIP

`func (o *IoK8sApiCoreV1PodStatus) GetHostIP() string`

GetHostIP returns the HostIP field if non-nil, zero value otherwise.

### GetHostIPOk

`func (o *IoK8sApiCoreV1PodStatus) GetHostIPOk() (*string, bool)`

GetHostIPOk returns a tuple with the HostIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostIP

`func (o *IoK8sApiCoreV1PodStatus) SetHostIP(v string)`

SetHostIP sets HostIP field to given value.

### HasHostIP

`func (o *IoK8sApiCoreV1PodStatus) HasHostIP() bool`

HasHostIP returns a boolean if a field has been set.

### GetInitContainerStatuses

`func (o *IoK8sApiCoreV1PodStatus) GetInitContainerStatuses() []IoK8sApiCoreV1ContainerStatus`

GetInitContainerStatuses returns the InitContainerStatuses field if non-nil, zero value otherwise.

### GetInitContainerStatusesOk

`func (o *IoK8sApiCoreV1PodStatus) GetInitContainerStatusesOk() (*[]IoK8sApiCoreV1ContainerStatus, bool)`

GetInitContainerStatusesOk returns a tuple with the InitContainerStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitContainerStatuses

`func (o *IoK8sApiCoreV1PodStatus) SetInitContainerStatuses(v []IoK8sApiCoreV1ContainerStatus)`

SetInitContainerStatuses sets InitContainerStatuses field to given value.

### HasInitContainerStatuses

`func (o *IoK8sApiCoreV1PodStatus) HasInitContainerStatuses() bool`

HasInitContainerStatuses returns a boolean if a field has been set.

### GetMessage

`func (o *IoK8sApiCoreV1PodStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *IoK8sApiCoreV1PodStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *IoK8sApiCoreV1PodStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *IoK8sApiCoreV1PodStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetNominatedNodeName

`func (o *IoK8sApiCoreV1PodStatus) GetNominatedNodeName() string`

GetNominatedNodeName returns the NominatedNodeName field if non-nil, zero value otherwise.

### GetNominatedNodeNameOk

`func (o *IoK8sApiCoreV1PodStatus) GetNominatedNodeNameOk() (*string, bool)`

GetNominatedNodeNameOk returns a tuple with the NominatedNodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNominatedNodeName

`func (o *IoK8sApiCoreV1PodStatus) SetNominatedNodeName(v string)`

SetNominatedNodeName sets NominatedNodeName field to given value.

### HasNominatedNodeName

`func (o *IoK8sApiCoreV1PodStatus) HasNominatedNodeName() bool`

HasNominatedNodeName returns a boolean if a field has been set.

### GetPhase

`func (o *IoK8sApiCoreV1PodStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *IoK8sApiCoreV1PodStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *IoK8sApiCoreV1PodStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *IoK8sApiCoreV1PodStatus) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetPodIP

`func (o *IoK8sApiCoreV1PodStatus) GetPodIP() string`

GetPodIP returns the PodIP field if non-nil, zero value otherwise.

### GetPodIPOk

`func (o *IoK8sApiCoreV1PodStatus) GetPodIPOk() (*string, bool)`

GetPodIPOk returns a tuple with the PodIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodIP

`func (o *IoK8sApiCoreV1PodStatus) SetPodIP(v string)`

SetPodIP sets PodIP field to given value.

### HasPodIP

`func (o *IoK8sApiCoreV1PodStatus) HasPodIP() bool`

HasPodIP returns a boolean if a field has been set.

### GetPodIPs

`func (o *IoK8sApiCoreV1PodStatus) GetPodIPs() []IoK8sApiCoreV1PodIP`

GetPodIPs returns the PodIPs field if non-nil, zero value otherwise.

### GetPodIPsOk

`func (o *IoK8sApiCoreV1PodStatus) GetPodIPsOk() (*[]IoK8sApiCoreV1PodIP, bool)`

GetPodIPsOk returns a tuple with the PodIPs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodIPs

`func (o *IoK8sApiCoreV1PodStatus) SetPodIPs(v []IoK8sApiCoreV1PodIP)`

SetPodIPs sets PodIPs field to given value.

### HasPodIPs

`func (o *IoK8sApiCoreV1PodStatus) HasPodIPs() bool`

HasPodIPs returns a boolean if a field has been set.

### GetQosClass

`func (o *IoK8sApiCoreV1PodStatus) GetQosClass() string`

GetQosClass returns the QosClass field if non-nil, zero value otherwise.

### GetQosClassOk

`func (o *IoK8sApiCoreV1PodStatus) GetQosClassOk() (*string, bool)`

GetQosClassOk returns a tuple with the QosClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosClass

`func (o *IoK8sApiCoreV1PodStatus) SetQosClass(v string)`

SetQosClass sets QosClass field to given value.

### HasQosClass

`func (o *IoK8sApiCoreV1PodStatus) HasQosClass() bool`

HasQosClass returns a boolean if a field has been set.

### GetReason

`func (o *IoK8sApiCoreV1PodStatus) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *IoK8sApiCoreV1PodStatus) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *IoK8sApiCoreV1PodStatus) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *IoK8sApiCoreV1PodStatus) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetResize

`func (o *IoK8sApiCoreV1PodStatus) GetResize() string`

GetResize returns the Resize field if non-nil, zero value otherwise.

### GetResizeOk

`func (o *IoK8sApiCoreV1PodStatus) GetResizeOk() (*string, bool)`

GetResizeOk returns a tuple with the Resize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResize

`func (o *IoK8sApiCoreV1PodStatus) SetResize(v string)`

SetResize sets Resize field to given value.

### HasResize

`func (o *IoK8sApiCoreV1PodStatus) HasResize() bool`

HasResize returns a boolean if a field has been set.

### GetStartTime

`func (o *IoK8sApiCoreV1PodStatus) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *IoK8sApiCoreV1PodStatus) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *IoK8sApiCoreV1PodStatus) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *IoK8sApiCoreV1PodStatus) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


