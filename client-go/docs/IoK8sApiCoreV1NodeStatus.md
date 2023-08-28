# IoK8sApiCoreV1NodeStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | Pointer to [**[]IoK8sApiCoreV1NodeAddress**](IoK8sApiCoreV1NodeAddress.md) | List of addresses reachable to the node. Queried from cloud provider, if available. More info: https://kubernetes.io/docs/concepts/nodes/node/#addresses Note: This field is declared as mergeable, but the merge key is not sufficiently unique, which can cause data corruption when it is merged. Callers should instead use a full-replacement patch. See https://pr.k8s.io/79391 for an example. Consumers should assume that addresses can change during the lifetime of a Node. However, there are some exceptions where this may not be possible, such as Pods that inherit a Node&#39;s address in its own status or consumers of the downward API (status.hostIP). | [optional] 
**Allocatable** | Pointer to **map[string]string** | Allocatable represents the resources of a node that are available for scheduling. Defaults to Capacity. | [optional] 
**Capacity** | Pointer to **map[string]string** | Capacity represents the total resources of a node. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#capacity | [optional] 
**Conditions** | Pointer to [**[]IoK8sApiCoreV1NodeCondition**](IoK8sApiCoreV1NodeCondition.md) | Conditions is an array of current observed node conditions. More info: https://kubernetes.io/docs/concepts/nodes/node/#condition | [optional] 
**Config** | Pointer to [**IoK8sApiCoreV1NodeConfigStatus**](IoK8sApiCoreV1NodeConfigStatus.md) |  | [optional] 
**DaemonEndpoints** | Pointer to [**IoK8sApiCoreV1NodeDaemonEndpoints**](IoK8sApiCoreV1NodeDaemonEndpoints.md) |  | [optional] 
**Images** | Pointer to [**[]IoK8sApiCoreV1ContainerImage**](IoK8sApiCoreV1ContainerImage.md) | List of container images on this node | [optional] 
**NodeInfo** | Pointer to [**IoK8sApiCoreV1NodeSystemInfo**](IoK8sApiCoreV1NodeSystemInfo.md) |  | [optional] 
**Phase** | Pointer to **string** | NodePhase is the recently observed lifecycle phase of the node. More info: https://kubernetes.io/docs/concepts/nodes/node/#phase The field is never populated, and now is deprecated.  Possible enum values:  - &#x60;\&quot;Pending\&quot;&#x60; means the node has been created/added by the system, but not configured.  - &#x60;\&quot;Running\&quot;&#x60; means the node has been configured and has Kubernetes components running.  - &#x60;\&quot;Terminated\&quot;&#x60; means the node has been removed from the cluster. | [optional] 
**VolumesAttached** | Pointer to [**[]IoK8sApiCoreV1AttachedVolume**](IoK8sApiCoreV1AttachedVolume.md) | List of volumes that are attached to the node. | [optional] 
**VolumesInUse** | Pointer to **[]string** | List of attachable volumes in use (mounted) by the node. | [optional] 

## Methods

### NewIoK8sApiCoreV1NodeStatus

`func NewIoK8sApiCoreV1NodeStatus() *IoK8sApiCoreV1NodeStatus`

NewIoK8sApiCoreV1NodeStatus instantiates a new IoK8sApiCoreV1NodeStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1NodeStatusWithDefaults

`func NewIoK8sApiCoreV1NodeStatusWithDefaults() *IoK8sApiCoreV1NodeStatus`

NewIoK8sApiCoreV1NodeStatusWithDefaults instantiates a new IoK8sApiCoreV1NodeStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *IoK8sApiCoreV1NodeStatus) GetAddresses() []IoK8sApiCoreV1NodeAddress`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *IoK8sApiCoreV1NodeStatus) GetAddressesOk() (*[]IoK8sApiCoreV1NodeAddress, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *IoK8sApiCoreV1NodeStatus) SetAddresses(v []IoK8sApiCoreV1NodeAddress)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *IoK8sApiCoreV1NodeStatus) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetAllocatable

`func (o *IoK8sApiCoreV1NodeStatus) GetAllocatable() map[string]string`

GetAllocatable returns the Allocatable field if non-nil, zero value otherwise.

### GetAllocatableOk

`func (o *IoK8sApiCoreV1NodeStatus) GetAllocatableOk() (*map[string]string, bool)`

GetAllocatableOk returns a tuple with the Allocatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatable

`func (o *IoK8sApiCoreV1NodeStatus) SetAllocatable(v map[string]string)`

SetAllocatable sets Allocatable field to given value.

### HasAllocatable

`func (o *IoK8sApiCoreV1NodeStatus) HasAllocatable() bool`

HasAllocatable returns a boolean if a field has been set.

### GetCapacity

`func (o *IoK8sApiCoreV1NodeStatus) GetCapacity() map[string]string`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *IoK8sApiCoreV1NodeStatus) GetCapacityOk() (*map[string]string, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *IoK8sApiCoreV1NodeStatus) SetCapacity(v map[string]string)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *IoK8sApiCoreV1NodeStatus) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetConditions

`func (o *IoK8sApiCoreV1NodeStatus) GetConditions() []IoK8sApiCoreV1NodeCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *IoK8sApiCoreV1NodeStatus) GetConditionsOk() (*[]IoK8sApiCoreV1NodeCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *IoK8sApiCoreV1NodeStatus) SetConditions(v []IoK8sApiCoreV1NodeCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *IoK8sApiCoreV1NodeStatus) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetConfig

`func (o *IoK8sApiCoreV1NodeStatus) GetConfig() IoK8sApiCoreV1NodeConfigStatus`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *IoK8sApiCoreV1NodeStatus) GetConfigOk() (*IoK8sApiCoreV1NodeConfigStatus, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *IoK8sApiCoreV1NodeStatus) SetConfig(v IoK8sApiCoreV1NodeConfigStatus)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *IoK8sApiCoreV1NodeStatus) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetDaemonEndpoints

`func (o *IoK8sApiCoreV1NodeStatus) GetDaemonEndpoints() IoK8sApiCoreV1NodeDaemonEndpoints`

GetDaemonEndpoints returns the DaemonEndpoints field if non-nil, zero value otherwise.

### GetDaemonEndpointsOk

`func (o *IoK8sApiCoreV1NodeStatus) GetDaemonEndpointsOk() (*IoK8sApiCoreV1NodeDaemonEndpoints, bool)`

GetDaemonEndpointsOk returns a tuple with the DaemonEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaemonEndpoints

`func (o *IoK8sApiCoreV1NodeStatus) SetDaemonEndpoints(v IoK8sApiCoreV1NodeDaemonEndpoints)`

SetDaemonEndpoints sets DaemonEndpoints field to given value.

### HasDaemonEndpoints

`func (o *IoK8sApiCoreV1NodeStatus) HasDaemonEndpoints() bool`

HasDaemonEndpoints returns a boolean if a field has been set.

### GetImages

`func (o *IoK8sApiCoreV1NodeStatus) GetImages() []IoK8sApiCoreV1ContainerImage`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *IoK8sApiCoreV1NodeStatus) GetImagesOk() (*[]IoK8sApiCoreV1ContainerImage, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *IoK8sApiCoreV1NodeStatus) SetImages(v []IoK8sApiCoreV1ContainerImage)`

SetImages sets Images field to given value.

### HasImages

`func (o *IoK8sApiCoreV1NodeStatus) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetNodeInfo

`func (o *IoK8sApiCoreV1NodeStatus) GetNodeInfo() IoK8sApiCoreV1NodeSystemInfo`

GetNodeInfo returns the NodeInfo field if non-nil, zero value otherwise.

### GetNodeInfoOk

`func (o *IoK8sApiCoreV1NodeStatus) GetNodeInfoOk() (*IoK8sApiCoreV1NodeSystemInfo, bool)`

GetNodeInfoOk returns a tuple with the NodeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeInfo

`func (o *IoK8sApiCoreV1NodeStatus) SetNodeInfo(v IoK8sApiCoreV1NodeSystemInfo)`

SetNodeInfo sets NodeInfo field to given value.

### HasNodeInfo

`func (o *IoK8sApiCoreV1NodeStatus) HasNodeInfo() bool`

HasNodeInfo returns a boolean if a field has been set.

### GetPhase

`func (o *IoK8sApiCoreV1NodeStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *IoK8sApiCoreV1NodeStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *IoK8sApiCoreV1NodeStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *IoK8sApiCoreV1NodeStatus) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetVolumesAttached

`func (o *IoK8sApiCoreV1NodeStatus) GetVolumesAttached() []IoK8sApiCoreV1AttachedVolume`

GetVolumesAttached returns the VolumesAttached field if non-nil, zero value otherwise.

### GetVolumesAttachedOk

`func (o *IoK8sApiCoreV1NodeStatus) GetVolumesAttachedOk() (*[]IoK8sApiCoreV1AttachedVolume, bool)`

GetVolumesAttachedOk returns a tuple with the VolumesAttached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumesAttached

`func (o *IoK8sApiCoreV1NodeStatus) SetVolumesAttached(v []IoK8sApiCoreV1AttachedVolume)`

SetVolumesAttached sets VolumesAttached field to given value.

### HasVolumesAttached

`func (o *IoK8sApiCoreV1NodeStatus) HasVolumesAttached() bool`

HasVolumesAttached returns a boolean if a field has been set.

### GetVolumesInUse

`func (o *IoK8sApiCoreV1NodeStatus) GetVolumesInUse() []string`

GetVolumesInUse returns the VolumesInUse field if non-nil, zero value otherwise.

### GetVolumesInUseOk

`func (o *IoK8sApiCoreV1NodeStatus) GetVolumesInUseOk() (*[]string, bool)`

GetVolumesInUseOk returns a tuple with the VolumesInUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumesInUse

`func (o *IoK8sApiCoreV1NodeStatus) SetVolumesInUse(v []string)`

SetVolumesInUse sets VolumesInUse field to given value.

### HasVolumesInUse

`func (o *IoK8sApiCoreV1NodeStatus) HasVolumesInUse() bool`

HasVolumesInUse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


