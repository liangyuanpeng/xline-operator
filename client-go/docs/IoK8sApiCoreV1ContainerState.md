# IoK8sApiCoreV1ContainerState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Running** | Pointer to [**IoK8sApiCoreV1ContainerStateRunning**](IoK8sApiCoreV1ContainerStateRunning.md) |  | [optional] 
**Terminated** | Pointer to [**IoK8sApiCoreV1ContainerStateTerminated**](IoK8sApiCoreV1ContainerStateTerminated.md) |  | [optional] 
**Waiting** | Pointer to [**IoK8sApiCoreV1ContainerStateWaiting**](IoK8sApiCoreV1ContainerStateWaiting.md) |  | [optional] 

## Methods

### NewIoK8sApiCoreV1ContainerState

`func NewIoK8sApiCoreV1ContainerState() *IoK8sApiCoreV1ContainerState`

NewIoK8sApiCoreV1ContainerState instantiates a new IoK8sApiCoreV1ContainerState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1ContainerStateWithDefaults

`func NewIoK8sApiCoreV1ContainerStateWithDefaults() *IoK8sApiCoreV1ContainerState`

NewIoK8sApiCoreV1ContainerStateWithDefaults instantiates a new IoK8sApiCoreV1ContainerState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRunning

`func (o *IoK8sApiCoreV1ContainerState) GetRunning() IoK8sApiCoreV1ContainerStateRunning`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *IoK8sApiCoreV1ContainerState) GetRunningOk() (*IoK8sApiCoreV1ContainerStateRunning, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *IoK8sApiCoreV1ContainerState) SetRunning(v IoK8sApiCoreV1ContainerStateRunning)`

SetRunning sets Running field to given value.

### HasRunning

`func (o *IoK8sApiCoreV1ContainerState) HasRunning() bool`

HasRunning returns a boolean if a field has been set.

### GetTerminated

`func (o *IoK8sApiCoreV1ContainerState) GetTerminated() IoK8sApiCoreV1ContainerStateTerminated`

GetTerminated returns the Terminated field if non-nil, zero value otherwise.

### GetTerminatedOk

`func (o *IoK8sApiCoreV1ContainerState) GetTerminatedOk() (*IoK8sApiCoreV1ContainerStateTerminated, bool)`

GetTerminatedOk returns a tuple with the Terminated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminated

`func (o *IoK8sApiCoreV1ContainerState) SetTerminated(v IoK8sApiCoreV1ContainerStateTerminated)`

SetTerminated sets Terminated field to given value.

### HasTerminated

`func (o *IoK8sApiCoreV1ContainerState) HasTerminated() bool`

HasTerminated returns a boolean if a field has been set.

### GetWaiting

`func (o *IoK8sApiCoreV1ContainerState) GetWaiting() IoK8sApiCoreV1ContainerStateWaiting`

GetWaiting returns the Waiting field if non-nil, zero value otherwise.

### GetWaitingOk

`func (o *IoK8sApiCoreV1ContainerState) GetWaitingOk() (*IoK8sApiCoreV1ContainerStateWaiting, bool)`

GetWaitingOk returns a tuple with the Waiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaiting

`func (o *IoK8sApiCoreV1ContainerState) SetWaiting(v IoK8sApiCoreV1ContainerStateWaiting)`

SetWaiting sets Waiting field to given value.

### HasWaiting

`func (o *IoK8sApiCoreV1ContainerState) HasWaiting() bool`

HasWaiting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


