# IoK8sApiCoreV1ContainerStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllocatedResources** | Pointer to **map[string]string** | AllocatedResources represents the compute resources allocated for this container by the node. Kubelet sets this value to Container.Resources.Requests upon successful pod admission and after successfully admitting desired pod resize. | [optional] 
**ContainerID** | Pointer to **string** | ContainerID is the ID of the container in the format &#39;&lt;type&gt;://&lt;container_id&gt;&#39;. Where type is a container runtime identifier, returned from Version call of CRI API (for example \&quot;containerd\&quot;). | [optional] 
**Image** | **string** | Image is the name of container image that the container is running. The container image may not match the image used in the PodSpec, as it may have been resolved by the runtime. More info: https://kubernetes.io/docs/concepts/containers/images. | 
**ImageID** | **string** | ImageID is the image ID of the container&#39;s image. The image ID may not match the image ID of the image used in the PodSpec, as it may have been resolved by the runtime. | 
**LastState** | Pointer to [**IoK8sApiCoreV1ContainerState**](IoK8sApiCoreV1ContainerState.md) |  | [optional] 
**Name** | **string** | Name is a DNS_LABEL representing the unique name of the container. Each container in a pod must have a unique name across all container types. Cannot be updated. | 
**Ready** | **bool** | Ready specifies whether the container is currently passing its readiness check. The value will change as readiness probes keep executing. If no readiness probes are specified, this field defaults to true once the container is fully started (see Started field).  The value is typically used to determine whether a container is ready to accept traffic. | 
**Resources** | Pointer to [**IoK8sApiCoreV1ResourceRequirements**](IoK8sApiCoreV1ResourceRequirements.md) |  | [optional] 
**RestartCount** | **int32** | RestartCount holds the number of times the container has been restarted. Kubelet makes an effort to always increment the value, but there are cases when the state may be lost due to node restarts and then the value may be reset to 0. The value is never negative. | 
**Started** | Pointer to **bool** | Started indicates whether the container has finished its postStart lifecycle hook and passed its startup probe. Initialized as false, becomes true after startupProbe is considered successful. Resets to false when the container is restarted, or if kubelet loses state temporarily. In both cases, startup probes will run again. Is always true when no startupProbe is defined and container is running and has passed the postStart lifecycle hook. The null value must be treated the same as false. | [optional] 
**State** | Pointer to [**IoK8sApiCoreV1ContainerState**](IoK8sApiCoreV1ContainerState.md) |  | [optional] 

## Methods

### NewIoK8sApiCoreV1ContainerStatus

`func NewIoK8sApiCoreV1ContainerStatus(image string, imageID string, name string, ready bool, restartCount int32, ) *IoK8sApiCoreV1ContainerStatus`

NewIoK8sApiCoreV1ContainerStatus instantiates a new IoK8sApiCoreV1ContainerStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1ContainerStatusWithDefaults

`func NewIoK8sApiCoreV1ContainerStatusWithDefaults() *IoK8sApiCoreV1ContainerStatus`

NewIoK8sApiCoreV1ContainerStatusWithDefaults instantiates a new IoK8sApiCoreV1ContainerStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllocatedResources

`func (o *IoK8sApiCoreV1ContainerStatus) GetAllocatedResources() map[string]string`

GetAllocatedResources returns the AllocatedResources field if non-nil, zero value otherwise.

### GetAllocatedResourcesOk

`func (o *IoK8sApiCoreV1ContainerStatus) GetAllocatedResourcesOk() (*map[string]string, bool)`

GetAllocatedResourcesOk returns a tuple with the AllocatedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedResources

`func (o *IoK8sApiCoreV1ContainerStatus) SetAllocatedResources(v map[string]string)`

SetAllocatedResources sets AllocatedResources field to given value.

### HasAllocatedResources

`func (o *IoK8sApiCoreV1ContainerStatus) HasAllocatedResources() bool`

HasAllocatedResources returns a boolean if a field has been set.

### GetContainerID

`func (o *IoK8sApiCoreV1ContainerStatus) GetContainerID() string`

GetContainerID returns the ContainerID field if non-nil, zero value otherwise.

### GetContainerIDOk

`func (o *IoK8sApiCoreV1ContainerStatus) GetContainerIDOk() (*string, bool)`

GetContainerIDOk returns a tuple with the ContainerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerID

`func (o *IoK8sApiCoreV1ContainerStatus) SetContainerID(v string)`

SetContainerID sets ContainerID field to given value.

### HasContainerID

`func (o *IoK8sApiCoreV1ContainerStatus) HasContainerID() bool`

HasContainerID returns a boolean if a field has been set.

### GetImage

`func (o *IoK8sApiCoreV1ContainerStatus) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *IoK8sApiCoreV1ContainerStatus) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *IoK8sApiCoreV1ContainerStatus) SetImage(v string)`

SetImage sets Image field to given value.


### GetImageID

`func (o *IoK8sApiCoreV1ContainerStatus) GetImageID() string`

GetImageID returns the ImageID field if non-nil, zero value otherwise.

### GetImageIDOk

`func (o *IoK8sApiCoreV1ContainerStatus) GetImageIDOk() (*string, bool)`

GetImageIDOk returns a tuple with the ImageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageID

`func (o *IoK8sApiCoreV1ContainerStatus) SetImageID(v string)`

SetImageID sets ImageID field to given value.


### GetLastState

`func (o *IoK8sApiCoreV1ContainerStatus) GetLastState() IoK8sApiCoreV1ContainerState`

GetLastState returns the LastState field if non-nil, zero value otherwise.

### GetLastStateOk

`func (o *IoK8sApiCoreV1ContainerStatus) GetLastStateOk() (*IoK8sApiCoreV1ContainerState, bool)`

GetLastStateOk returns a tuple with the LastState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastState

`func (o *IoK8sApiCoreV1ContainerStatus) SetLastState(v IoK8sApiCoreV1ContainerState)`

SetLastState sets LastState field to given value.

### HasLastState

`func (o *IoK8sApiCoreV1ContainerStatus) HasLastState() bool`

HasLastState returns a boolean if a field has been set.

### GetName

`func (o *IoK8sApiCoreV1ContainerStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IoK8sApiCoreV1ContainerStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IoK8sApiCoreV1ContainerStatus) SetName(v string)`

SetName sets Name field to given value.


### GetReady

`func (o *IoK8sApiCoreV1ContainerStatus) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *IoK8sApiCoreV1ContainerStatus) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *IoK8sApiCoreV1ContainerStatus) SetReady(v bool)`

SetReady sets Ready field to given value.


### GetResources

`func (o *IoK8sApiCoreV1ContainerStatus) GetResources() IoK8sApiCoreV1ResourceRequirements`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *IoK8sApiCoreV1ContainerStatus) GetResourcesOk() (*IoK8sApiCoreV1ResourceRequirements, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *IoK8sApiCoreV1ContainerStatus) SetResources(v IoK8sApiCoreV1ResourceRequirements)`

SetResources sets Resources field to given value.

### HasResources

`func (o *IoK8sApiCoreV1ContainerStatus) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetRestartCount

`func (o *IoK8sApiCoreV1ContainerStatus) GetRestartCount() int32`

GetRestartCount returns the RestartCount field if non-nil, zero value otherwise.

### GetRestartCountOk

`func (o *IoK8sApiCoreV1ContainerStatus) GetRestartCountOk() (*int32, bool)`

GetRestartCountOk returns a tuple with the RestartCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartCount

`func (o *IoK8sApiCoreV1ContainerStatus) SetRestartCount(v int32)`

SetRestartCount sets RestartCount field to given value.


### GetStarted

`func (o *IoK8sApiCoreV1ContainerStatus) GetStarted() bool`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *IoK8sApiCoreV1ContainerStatus) GetStartedOk() (*bool, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *IoK8sApiCoreV1ContainerStatus) SetStarted(v bool)`

SetStarted sets Started field to given value.

### HasStarted

`func (o *IoK8sApiCoreV1ContainerStatus) HasStarted() bool`

HasStarted returns a boolean if a field has been set.

### GetState

`func (o *IoK8sApiCoreV1ContainerStatus) GetState() IoK8sApiCoreV1ContainerState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *IoK8sApiCoreV1ContainerStatus) GetStateOk() (*IoK8sApiCoreV1ContainerState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *IoK8sApiCoreV1ContainerStatus) SetState(v IoK8sApiCoreV1ContainerState)`

SetState sets State field to given value.

### HasState

`func (o *IoK8sApiCoreV1ContainerStatus) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


